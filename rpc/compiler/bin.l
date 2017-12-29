

/* IMPORT START CONDITION */
%x IMPORT

%{

#include <stdio.h>

/* bison value type. */
#include <string>
#define YYSTYPE std::string
#include "bin.tab.hpp"

#include "Compiler.h"

/* Save current file state.
	To import a new file, we muse push current parsing file
	state to fileStack, then begin to parse new file.
*/
struct FileStackNode
{
	std::string		filename_;
	int				lineno_;
	YY_BUFFER_STATE	bs_;
};
std::vector<FileStackNode>	fileStack;

%}

/* SO we don't choke on files that use \r\n */
NL [\r?\n]

%array
%option never-interactive
%option noyywrap


%%

enum            return TOKEN_ENUM;
struct			return TOKEN_STRUCT;
service			return TOKEN_SERVICE;
int64			return TOKEN_INT64;
uint64			return TOKEN_UINT64;
double			return TOKEN_DOUBLE;
float			return TOKEN_FLOAT;
int32			return TOKEN_INT32;
uint32			return TOKEN_UINT32;
int16			return TOKEN_INT16;
uint16			return TOKEN_UINT16;
int8			return TOKEN_INT8;
uint8			return TOKEN_UINT8;
bool			return TOKEN_BOOL;
string			return TOKEN_STRING;
array			return TOKEN_ARRAY;
bytes			return TOKEN_BYTES;

#import			{	
					/* Begin 'IMPORT' start condition if we met "import" keyword */
					BEGIN(IMPORT);
				}
<IMPORT>[ \t]*		/* eat the white space */
<IMPORT>[^ \t\n]+	{ 
					
					/* got the import file name */
					std::string importFilename;
					if( yytext[yyleng-1] == '\r' )
						importFilename = std::string( yytext, yyleng-1 );
					else
						importFilename = std::string( yytext, yyleng );
					if( Compiler::inst().importedFiles_.find( importFilename ) != Compiler::inst().importedFiles_.end() )
					{
						// File is imported by other file, skip it.
						BEGIN( INITIAL );
					}
					else
					{
						// save this filename to imported files.
						Compiler::inst().importedFiles_.insert( importFilename );

						// Save current file state.
						FileStackNode n;
						n.filename_ = Compiler::inst().curFilename_;
						n.lineno_	= Compiler::inst().curLineno_ + 1;	
						n.bs_		= YY_CURRENT_BUFFER;
						fileStack.push_back( n );
						// Set new file state.!!!Remove the last \n
						Compiler::inst().curFilename_	= importFilename;
						Compiler::inst().curLineno_	= 1;
						if( fileStack.size() == 1 )
						{
							/* root import */
							Compiler::inst().imports_.push_back( importFilename );
						}
						
						// Open new file to parse.
						yyin = fopen( importFilename.c_str(), "r" );
						if( !yyin )
						{
							// Iterate all import path to find file.
							for( size_t i = 0; i < Compiler::inst().importPaths_.size(); i++ )
							{
								std::string tempfilename = Compiler::inst().importPaths_[i];
								tempfilename += importFilename;
								yyin = fopen( tempfilename.c_str(), "r" );
								if( yyin )
									break;
							}
							if( !yyin )
							{
								// Can not found imported file.
								fprintf( stderr, "failed to import file \"%s\".\n", importFilename.c_str() );
								exit(1);
							}
						}
						yy_switch_to_buffer( yy_create_buffer( yyin, YY_BUF_SIZE ) );
						BEGIN( INITIAL );
					}
				}
<<EOF>>			{
					/* End of file */
					if( fileStack.size() == 0 )
					{
						// Root file, teminated.
						yyterminate();
					}
					else
					{
						// Recover file state.
						FileStackNode n = fileStack.back();
						fileStack.pop_back();
						Compiler::inst().curFilename_	= n.filename_;
						Compiler::inst().curLineno_	= n.lineno_;
						yy_delete_buffer( YY_CURRENT_BUFFER );
						yy_switch_to_buffer( n.bs_ ); 
					}
				}

[a-ij-rs-zA-IJ-RS-Z_][a-ij-rs-zA-IJ-RS-Z0-9_]* {
					/* Identifier */
					yylval = std::string( yytext, yyleng );
					return TOKEN_IDENTIFIER;
				}
				
[1-9][0-9]*		{
					/* unsigned int */
					yylval = std::string( yytext, yyleng );
					return TOKEN_UINTEGER_LITERAL;
				}
				
\/\/.*{NL}      {
					/* C++ Comment */
					Compiler::inst().curLineno_++;
                }

"#<"			{
					/* start cpp code. */
					for(;;) 
					{
						char const c = yyinput();
						if (c == '#') 
						{
							char const next = yyinput();
							if (next == '>')
								break;
							else
								unput(next);
						}
						else if (c == '\n')
							Compiler::inst().curLineno_++;
						// 仅存储根文件的cppcode.
						if( fileStack.size() == 0 )
							Compiler::inst().cppcode_.push_back(c);
					}
				}

"#{"			{
					/* struct cpp code. */
					for(;;) 
					{
						char const c = yyinput();
						if (c == '#') 
						{
							char const next = yyinput();
							if (next == '}')
								break;
							else
								unput(next);
						}
						else if (c == '\n')
							Compiler::inst().curLineno_++;
						Compiler::inst().curStruct_.cppcode_.push_back(c);
					}
				}

"/*"			{
					/* C Comment. */
					for(;;) {
						char const c = yyinput();
						if (c == '*') {
							char const next = yyinput();
							if (next == '/')
								break;
							else
								unput(next);
						}
						else if (c == '\n') {
							Compiler::inst().curLineno_++;
						}
					}
				}

[ \t]*          { /*White space.*/ }

{NL}            { /*New Line.*/ Compiler::inst().curLineno_++;}

.               return yytext[0];

%%

/* subroutines */
