#ifndef __Generator_h__
#define __Generator_h__

#include <string>

/** CodeGenerator is a base class which is used to generate real code of some kind.
*/
class CodeGenerator
{
public:
	virtual void generate()=0;
};

#endif // __Generator_h__
