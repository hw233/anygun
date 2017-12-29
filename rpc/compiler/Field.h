#ifndef __Field_h__
#define __Field_h__

#include <string>
#include <vector>

#include "Definition.h"

/** Field Type. */
enum EFieldType
{
	FT_INT64,	///< 8 bytes
	FT_UINT64,	///< 8 bytes.
	FT_DOUBLE,	///< 8 bytes.
	FT_FLOAT,	///< 4 bytes
	FT_INT32,	///< 4 bytes
	FT_UINT32,	///< 4 bytes
	FT_INT16,	///< 2 bytes
	FT_UINT16,	///< 2 bytes
	FT_INT8,	///< 1 bytes
	FT_UINT8,	///< 1 bytes
	FT_BOOL,	///< 1 bytes
	FT_ENUM,	///< 1 bytes
	FT_STRING,	///< len * 1 bytes.
	FT_USER,	///< getSize()
	//FT_BYTES,	///< getSize()
};

/** Field 代表一个 type name 表示的字段对象.
	Field 用来表示 Struct 中的成员变量以及 Method 的参数.
	每个 Field 实例都有类型和名称。
*/
class Field
{
public:
	Field():
	type_(FT_USER),
	userType_(NULL),
	array_(false),
	maxArrLength_(0XFFFF),
	maxStrLength_(0XFFFF)
	{}

	void setName(const std::string& n) { name_ = n; }
	const std::string& getName(){ return name_; }
	const char* getNameC()		{ return name_.c_str(); }
	void setType(EFieldType t) { type_ = t;}
	EFieldType getType()		{ return type_; }
	void setUserType(Definition* ut)	{ userType_ = ut; }
	Definition* getUserType()			{ return userType_; }
	void setArray(bool a)	{ array_ = a; }
	bool getArray()			{ return array_; }
	void setMaxArrLength(size_t l)	{ maxArrLength_ = l; }
	size_t getMaxArrLength()		{ return maxArrLength_; }
	void setMaxStrLength(size_t l)	{ maxStrLength_ = l; }
	size_t getMaxStrLength()		{ return maxStrLength_; }
private:
	std::string				name_;			///< 字段名称.
	EFieldType			type_;			///< 字段类型.
	Definition*		userType_;		///< 如果是FT_USER或FT_ENUM.
	bool					array_;			///< 是否是array
	size_t					maxArrLength_;	///< 接收端允许的最大数组长度.
	size_t					maxStrLength_;	///< 接收端允许的最大字符串长度.
};

/** Field 容器类.
	FieldContainer 包含了一个 Field 对象数组，用来管理这些fields.
*/
class FieldContainer
{
public:
	/** 查找一个field是否存在. */
	bool findField(const std::string& name);
	/** How many bytes will be used for the field mask of this field container. */
	int getFMByteNum() { return (fields_.size()-1)/8+1; }

	std::vector<Field>	fields_;
};

#endif//__Field_h__