#ifndef __CodeFile_h__
#define __CodeFile_h__

#include <string>
/** Representing a code file being generated. */
class CodeFile
{
public:
	CodeFile(const std::string& filename);
	~CodeFile();
	void indent();
	void recover();
	/** Output a line of indented code. */
	void output(const char* s, ...);
	/** Begin a new line. */
	void begin();
	/** Append to current line. */
	void append(const char* s, ...);
	/** Finish a line. */
	void end();

private:
	FILE*	file_;
	unsigned int	indentation_;
};

#endif// __CodeFile_h__