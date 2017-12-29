
#ifndef __Struct_h__
#define __Struct_h__

#include <vector>

#include "Definition.h"
#include "Field.h"

/** Struct 代表一个用户自定义的数据类型. 
	Struct 类似一个 c++ struct 定义，由成员变量列表组成。
	APRC_Struct 支持派生.
*/
class Struct : 
	public Definition,
	public FieldContainer
{
public:
	Struct():
	Definition(),
	super_(NULL),
	skipComp_(false)
	{}
	Struct(const std::string& f, const std::string& n):
	Definition(f, n),
	super_(NULL),
	skipComp_(false)
	{}

	virtual Struct*	getStruct()	{ return this; }

	/** 查找一个成员属性定义是否存在. */
	bool findField( const std::string& name );

	/** 获得struct field数量. */
	size_t getFieldNum();

	Struct*				super_;		///< 继承自
	std::string					cppcode_;	///< c++ code.
	bool						skipComp_;	///< 忽略压缩方式.
};


#endif//__Struct_h__