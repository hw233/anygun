#ifndef __PYGenerator_h__
#define __PYGenerator_h__

#include "CodeGenerator.h"

class PYGenerator : public CodeGenerator
{
public:
	virtual void generate();
};

#endif// __PYGenerator_h__