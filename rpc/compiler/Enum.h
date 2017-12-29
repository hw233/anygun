#ifndef __Enum_h__
#define __Enum_h__

#include <vector>
#include "Definition.h"
#include "Field.h"

/** 枚举定义类.
	Enum 表示一个和c++ enum类似的枚举定义。
*/
class Enum : public Definition
{
public:
	Enum():Definition(){}
	Enum(const std::string& f, const std::string& n)
		:Definition(f, n)
	{
		super_.setType(FT_INT32);
	}

	/** 查找一个枚举项是否存在.
	*/
	bool findItem( const std::string& item );
	Field &getSuperType(){return super_;}
	virtual Enum* getEnum() { return this; }

	Field						super_;
	std::vector< std::string >	items_;	///< 枚举项列表.
};


#endif//__Enum_h__