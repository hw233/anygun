#ifndef __Service_h__
#define __Service_h__

#include "Definition.h"
#include "Field.h"

/** Method 表示一个 service method 定义，类似一个c function declaration. */
class Method : public FieldContainer
{
public:
	Method()
	{}

	const std::string& getName()	{ return name_; }
	const char* getNameC()			{ return name_.c_str(); }
	void setName(const std::string& n) { name_ = n; }
private:
	std::string				name_;				///< 方法名称.
};


class Service : public Definition
{
public:
	Service():
	Definition(),
	super_(NULL)
	{}
	Service(const std::string&f, const std::string& n):
	Definition(f, n),
	super_(NULL)
	{}

	virtual Service* getService()	{ return this; }

	/** 查找一个方法是否存在. */
	bool findMethod( const std::string& name );

	/** 将自己的所有父对象按照顺序推入一个数组. */
	void getParents(std::vector<Service*>& parents);

	/** 获得方法数量. */
	size_t getMethodNum();

	/** 获得所有method */
	void getAllMethods(std::vector<Method*>& methods);

	Service*				super_;
	std::vector<Method>	methods_;
};


#endif//__Service_h__