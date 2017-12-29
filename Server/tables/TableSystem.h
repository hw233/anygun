#ifndef __TABLE_SYSTEM_H__
#define __TABLE_SYSTEM_H__
#include "config.h"

class TableSystem{
public:
	SINGLETON_FUNCTION(TableSystem);
public:
	bool Load();
	bool Check();
};

#endif