

/* Declarations */

%{

#include <string>
#define YYSTYPE std::string

void yyerror (const char *);
int yylex (void);

#include "Compiler.h"

%}

%error-verbose


/*
 * Token types: These are returned by the lexer
 */
%token				TOKEN_IDENTIFIER
%token				TOKEN_ENUM
%token				TOKEN_STRUCT
%token				TOKEN_SERVICE
%token				TOKEN_INT64
%token				TOKEN_UINT64
%token				TOKEN_DOUBLE
%token				TOKEN_FLOAT
%token				TOKEN_INT32
%token				TOKEN_UINT32
%token				TOKEN_INT16
%token				TOKEN_UINT16
%token				TOKEN_INT8
%token				TOKEN_UINT8
%token				TOKEN_BOOL
%token				TOKEN_STRING
%token				TOKEN_ARRAY
%token				TOKEN_BYTES
%token				TOKEN_UINTEGER_LITERAL


%%

/*
 * Production starts here.
 */
start: definitions;

/* definitions */
definitions: 
	definitions definition
	| 
	/* empty */
	;

/*definition*/
definition:
	enumeration ';'
	|
	structure ';'
	|
	service ';'
    ;

/*enumeration*/
enumeration:
	TOKEN_ENUM
	TOKEN_IDENTIFIER ':' enumeration_super 
	{
		// Check enum name.
		if( Compiler::inst().findDefinition( $2 ) )
		{
			Compiler::inst().outputErrorFL("duplicated definition \"%s\".\n", $2.c_str()); 
			YYERROR; 
		};
		
		// Init current enum.
		Compiler::inst().curEnum_ = Enum(Compiler::inst().curFilename_, $2);
	}
    |
    TOKEN_ENUM 
    TOKEN_IDENTIFIER
    {
    	// Check enum name.
		if( Compiler::inst().findDefinition( $2 ) )
		{
			Compiler::inst().outputErrorFL("duplicated definition \"%s\".\n", $2.c_str()); 
			YYERROR; 
		};
		
		// Init current enum.
		Compiler::inst().curEnum_ = Enum(Compiler::inst().curFilename_, $2);
    }
    '{' enum_items '}'
	{
		// Add this definition.
		Enum* e = new Enum( Compiler::inst().curEnum_ );
		Compiler::inst().definitions_.push_back(e);
	}
	;

/*enum super type*/
enumeration_super:
	TOKEN_INT64	    { Compiler::inst().curEnum_.super_.setType(FT_INT64); }
	|
	TOKEN_UINT64	{ Compiler::inst().curEnum_.super_.setType( FT_UINT64); }
	|
	TOKEN_INT32	    { Compiler::inst().curEnum_.super_.setType(FT_INT32); }
	|
	TOKEN_UINT32	{ Compiler::inst().curEnum_.super_.setType(FT_UINT32); }
	|
	TOKEN_INT16	    { Compiler::inst().curEnum_.super_.setType(FT_INT16); }
	|
	TOKEN_UINT16	{ Compiler::inst().curEnum_.super_.setType(FT_UINT16); }
	|   
	TOKEN_INT8	    { Compiler::inst().curEnum_.super_.setType(FT_INT8); }
	|
	TOKEN_UINT8	    { Compiler::inst().curEnum_.super_.setType(FT_UINT8); }
	|
    ;

/*enum_items*/
enum_items:
	enum_items enum_item
	|
	/*empty*/
	;

/*enum_item*/	
enum_item:
	TOKEN_IDENTIFIER ','
	{
		// Check enum item name.
		if( Compiler::inst().findDefinition( $1 ) )
		{
			Compiler::inst().outputErrorFL("invalid enum item \"%s\"", $1.c_str());
			YYERROR;
		}
		if( Compiler::inst().curEnum_.findItem( $1 ) )
		{
			Compiler::inst().outputErrorFL("duplicated enum item \"%s\"", $1.c_str());
			YYERROR;
		}
		Compiler::inst().curEnum_.items_.push_back($1);
	}
	;

/*structure*/
structure:
	TOKEN_STRUCT
	TOKEN_IDENTIFIER
	{
		// Check struct name.
		if( Compiler::inst().findDefinition( $2 ) )
		{ 
			Compiler::inst().outputErrorFL("duplicated definition \"%s\".\n", $2.c_str()); 
			YYERROR; 
		};
		
		// Init current struct.
		Compiler::inst().curStruct_ = Struct(Compiler::inst().curFilename_, $2);
	}
	opt_super_structure
	opt_struct_flags
	'{' struct_fields '}'
	{
		// Add this struct definition.
		Struct* s = new Struct( Compiler::inst().curStruct_ );
		Compiler::inst().definitions_.push_back( s );
	}
	;

/*opt_struct_flags */
opt_struct_flags:
	'(' struct_flags ')'
	|
	/*empty*/
	;

/* struct_flags */
struct_flags:
	struct_flags struct_flag
	|
	/*empty*/
	;

/* struct_flag */
struct_flag:
	TOKEN_IDENTIFIER
	{
		// skip serializtion compress.
		if( $1 == "skipcomp" )
			Compiler::inst().curStruct_.skipComp_ = true;
		else
		{
			Compiler::inst().outputErrorFL("invalid struct flag \"%s\"", $1.c_str());
			YYERROR;
		}
	}
	;

/*opt_super_structure*/
opt_super_structure:
	super_structure
	|
	/*empty*/
	;

/*super_structure*/
super_structure:
	':' TOKEN_IDENTIFIER
	{
		Definition* d = Compiler::inst().findDefinition( $2 );
		if( !d || !d->getStruct() || Compiler::inst().curStruct_.getName() == $2 )
		{
			Compiler::inst().outputErrorFL("invalid super struct \"%s\"", $2.c_str());
			YYERROR;
		}
		Compiler::inst().curStruct_.super_ = d->getStruct();
	}
	;

/*struct_fields*/
struct_fields:
	struct_fields struct_field
	|
	/* empty */
	;
	
/*struct_field*/
struct_field:
	field_type TOKEN_IDENTIFIER ';'
	{
		// Check field name.
		if( Compiler::inst().findDefinition( $2 ) )
		{
			Compiler::inst().outputErrorFL("invalid field name\"%s\"", $2.c_str());
			YYERROR;
		}
		if( Compiler::inst().curStruct_.findField( $2 ) )
		{
			Compiler::inst().outputErrorFL("duplicated field name\"%s\"", $2.c_str());
			YYERROR;
		}

		Compiler::inst().curField_.setName($2);
		Compiler::inst().curStruct_.fields_.push_back( Compiler::inst().curField_ );
		Compiler::inst().curField_ = Field();
	}
	;
	
/*service*/
service:
	TOKEN_SERVICE
	TOKEN_IDENTIFIER
	{
		// Check service name.
		if( Compiler::inst().findDefinition( $2 ) )
		{ 
			Compiler::inst().outputErrorFL("duplicated definition \"%s\".\n", $2.c_str()); 
			YYERROR; 
		};
		
		// Init current service.
		Compiler::inst().curService_ = Service(Compiler::inst().curFilename_, $2);
		Compiler::inst().curMethod_ = Method();
	}
	opt_super_service
	'{' service_methods '}'
	{
		// Add this service.
		Service* s = new Service( Compiler::inst().curService_ );
		Compiler::inst().definitions_.push_back( s );
	}
	;

/*opt_super_service*/
opt_super_service:
	super_service
	|
	/*empty*/
	;

/*super_service*/
super_service:
	':' TOKEN_IDENTIFIER
	{
		Definition* d = Compiler::inst().findDefinition( $2 );
		if( !d || !d->getService() || Compiler::inst().curService_.getName() == $2 )
		{
			Compiler::inst().outputErrorFL("invalid super service \"%s\"", $2.c_str());
			YYERROR;
		}
		Compiler::inst().curService_.super_ = d->getService();
	}
	;

/*service_methods*/
service_methods:
	service_methods service_method
	|
	/*empty*/
	;

/*service_method*/
service_method:
	TOKEN_IDENTIFIER '(' service_method_params ')' ';'
	{
		if( Compiler::inst().curService_.findMethod( $1 ) )
		{
			Compiler::inst().outputErrorFL("duplicated method name\"%s\"", $1.c_str());
			YYERROR;
		}
		Compiler::inst().curMethod_.setName($1);
		Compiler::inst().curService_.methods_.push_back( Compiler::inst().curMethod_ );
		Compiler::inst().curMethod_ = Method();
	}
	;

/*service_method_params*/
service_method_params:
	service_method_params ',' service_method_param
	|
	service_method_param
	|
	/*empty*/
	;

/*service_method_param*/
service_method_param:
	field_type 	TOKEN_IDENTIFIER
	{
		if( Compiler::inst().curMethod_.findField( $2 ) )
		{
			Compiler::inst().outputErrorFL("duplicated param name\"%s\"", $2.c_str());
			YYERROR;
		}
		Compiler::inst().curField_.setName($2);
		Compiler::inst().curMethod_.fields_.push_back( Compiler::inst().curField_ );
		Compiler::inst().curField_ = Field();
	}
	;
	
/*field_type*/
field_type:
	field_inner_type
	{
		Compiler::inst().curField_.setArray(false);
	}
	|
	TOKEN_ARRAY opt_array_max_size '<' field_inner_type '>'
	{
		Compiler::inst().curField_.setArray(true);
	}
	|
	TOKEN_BYTES opt_array_max_size
	{
		// bytes == array<uint8>
		Compiler::inst().curField_.setArray(true);
		Compiler::inst().curField_.setType(FT_UINT8);
	}
	;

/* opt_array_max_size. */
opt_array_max_size:
	'[' TOKEN_UINTEGER_LITERAL ']'
	{
		Compiler::inst().curField_.setMaxArrLength(::atoi($2.c_str()));
	}
	|
	/*empty*/
	;

/*field_inner_type*/
field_inner_type:
	TOKEN_INT64	{ Compiler::inst().curField_.setType(FT_INT64); }
	|
	TOKEN_UINT64	{ Compiler::inst().curField_.setType(FT_UINT64); }
	|
	TOKEN_DOUBLE	{ Compiler::inst().curField_.setType(FT_DOUBLE); }
	|
	TOKEN_FLOAT	{ Compiler::inst().curField_.setType(FT_FLOAT); }
	|
	TOKEN_INT32	{ Compiler::inst().curField_.setType(FT_INT32); }
	|
	TOKEN_UINT32	{ Compiler::inst().curField_.setType(FT_UINT32); }
	|
	TOKEN_INT16	{ Compiler::inst().curField_.setType(FT_INT16); }
	|
	TOKEN_UINT16	{ Compiler::inst().curField_.setType(FT_UINT16); }
	|
	TOKEN_INT8	{ Compiler::inst().curField_.setType(FT_INT8); }
	|
	TOKEN_UINT8	{ Compiler::inst().curField_.setType(FT_UINT8); }
	|
	TOKEN_BOOL	{ Compiler::inst().curField_.setType(FT_BOOL); }
	|
	TOKEN_STRING	opt_string_max_length { Compiler::inst().curField_.setType(FT_STRING); }
	|
	TOKEN_IDENTIFIER	
	{
		Definition* d = Compiler::inst().findDefinition( $1 );
		if( !d || d->getService() )
		{
			Compiler::inst().outputErrorFL("invalid type \"%s\"", $1.c_str());
			YYERROR;
		}
		if( d->getEnum() )
			Compiler::inst().curField_.setType(FT_ENUM); 
		else
			Compiler::inst().curField_.setType(FT_USER); 
		Compiler::inst().curField_.setUserType(d);
	}
	;

/* opt_string_max_length. */
opt_string_max_length:
	'[' TOKEN_UINTEGER_LITERAL ']'
	{
		Compiler::inst().curField_.setMaxStrLength(::atoi($2.c_str()));
	}
	|
	/*empty*/
	;
		
%%


int yywrap (void)
{
  return 1;
}

void yyerror (const char *msg)
{
	Compiler::inst().outputErrorFL(msg);
}
