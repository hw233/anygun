#include <stdio.h>
#include <stdarg.h>
#include "CodeFile.h"

CodeFile::CodeFile(const std::string& filename):
file_(NULL),
indentation_(0)
{
	file_ = fopen(filename.c_str(), "w");
	if(!file_)
	{
		std::string es = "failed to open file ";
		es += filename;
		throw es.c_str();
	}
}

CodeFile::~CodeFile()
{
	fclose(file_);
}

void CodeFile::indent()
{
	indentation_++;
}

void CodeFile::recover()
{
	indentation_--;
}

void CodeFile::output(const char* s, ...)
{
	for(unsigned int i = 0; i < indentation_; i++)
		fprintf(file_, "\t");
	va_list arg_ptr;
	va_start(arg_ptr, s);
	vfprintf(file_, s, arg_ptr);
	va_end(arg_ptr);
	fprintf(file_, "\n");
}

void CodeFile::begin()
{
	for(unsigned int i = 0; i < indentation_; i++)
		fprintf(file_, "\t");
}

void CodeFile::append(const char* s, ...)
{
	va_list arg_ptr;
	va_start(arg_ptr, s);
	vfprintf(file_, s, arg_ptr);
	va_end(arg_ptr);
}

void CodeFile::end()
{
	fprintf(file_, "\n");
}
