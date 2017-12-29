#include "Compiler.h"
#include <string>
#include <iostream>
#include <iterator>
#include "Args.h"

/** arpcc编译器入口. */
int main( int argc, char *argv[] )
{

	try 
	{
		Args args(argc, argv, ";i;o;g;");

		Compiler::inst().generator_ = args.GetCString('g');
		Compiler::inst().inputFileName_ = args.GetCString('i');
		Compiler::inst().outputDir_ = args.GetCString('o');
		
	}
	catch(...) 
	{
		return 1;
	}

	// 编译源文件
	return Compiler::inst().compile();
}
