#ifndef __ASGenerator_h__
#define __ASGenerator_h__

#include "CodeGenerator.h"

class ASGenerator :public CodeGenerator
{
public:
	virtual void generate();
};

#endif// __ASGenerator_h__