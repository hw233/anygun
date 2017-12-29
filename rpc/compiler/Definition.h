#ifndef __Definition_h__
#define __Definition_h__

#include <stdio.h>
#include <string>

// Forward declaration.
class Enum;
class Struct;
class Service;

/** RPC 语言定义基础类.
	Definition 的列表组成了 aprc 语言的基础单位。
*/
class Definition
{
public:
	Definition(){}
	Definition(const std::string& f, const std::string& n):
	file_(f),
	name_(n)
	{}
	virtual ~Definition() {};

	const char* getFileC()	{ return file_.c_str(); }
	const std::string& getFile(){ return file_; }

	const char* getNameC()	{ return name_.c_str(); }
	const std::string& getName(){ return name_; }

	/** @name Type Query. */
	//@{
	virtual Enum*		getEnum()		{ return NULL; }
	virtual Struct*	getStruct()		{ return NULL; }
	virtual Service*	getService()	{ return NULL; }
	//@}

private:
	std::string					file_;	///< 这个定义所在文件名称.
	std::string					name_;	///< 定义名称.
};


#endif//__Definition_h__