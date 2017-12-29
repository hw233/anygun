#include "CppGenerator.h"
#include "Compiler.h"
#include "CodeFile.h"


static const char* getFieldCppType(Field& f, bool withArray = true)
{
	static std::string name;
	if(f.getArray() && withArray)
		name = "std::vector< ";
	else
		name = "";
	switch(f.getType())
	{
	case FT_INT64:
		name += "S64";
		break;
	case FT_UINT64:
		name += "U64";
		break;
	case FT_DOUBLE:
		name += "F64";
		break;
	case FT_FLOAT:
		name += "F32";
		break;
	case FT_INT32:
		name += "S32";
		break;
	case FT_UINT32:
		name += "U32";
		break;
	case FT_INT16:
		name += "S16";
		break;
	case FT_UINT16:
		name += "U16";
		break;
	case FT_INT8:
		name += "S8";
		break;
	case FT_UINT8:
		name += "U8";
		break;
	case FT_BOOL:
		name += "bool";
		break;
	case FT_STRING:
		name += "std::string";
		break;
	case FT_USER:
	case FT_ENUM:
		name += f.getUserType()->getName();
		break;
	default:
		throw "Invalid field type.";
	}

	if(f.getArray() && withArray)
		name += " >";
	return name.c_str();
}


static void generateEnumDecl(CodeFile& f, Enum* e)
{
	f.output("// enum %s", e->getNameC());
	f.output("enum %s : %s", e->getNameC(),getFieldCppType(e->getSuperType()));
	f.output("{");
	f.indent();
	for(size_t i = 0; i < e->items_.size(); i++)
		f.output("%s,", e->items_[i].c_str());
	f.recover();
	f.output("};");
	f.output("extern EnumInfo enum%s;", e->getNameC());
}

static const char* getFieldCppDefault(Field& f)
{
	static std::string name;
	if( f.getArray() )
		return NULL;
	name = "";
	switch( f.getType() )
	{
	case FT_INT64:
	case FT_UINT64:
	case FT_DOUBLE:
	case FT_FLOAT:
	case FT_INT32:
	case FT_UINT32:
	case FT_INT16:
	case FT_UINT16:
	case FT_INT8:
	case FT_UINT8:
		name = "0";
		break;
	case FT_BOOL:
		name = "false";
		break;
	case FT_ENUM:
		name = "(" + f.getUserType()->getName() + ")(0)";
		break;
	case FT_STRING:
		return NULL;
	case FT_USER:
		return NULL;
	default:
		throw "Invalid field type.";
	}

	return name.c_str();
}

static bool useParamReference(Field& f)
{
	if(f.getArray() || f.getType() == FT_STRING || f.getType() == FT_USER)
		return true;
	return false;
}

static void generateStubMethodDecl(CodeFile& f, Method& m)
{
	f.begin();
	f.append("void %s(", m.getNameC());
	for(size_t i = 0; i < m.fields_.size(); i++)
	{
		Field& field = m.fields_[i];
		f.append("%s%s%s %s%s", 
			useParamReference(field)?"const ":"",
			getFieldCppType(field),
			useParamReference(field)?"&":"",
			field.getNameC(),
			(i == m.fields_.size()-1)?"":",");
	}
	f.append(");");
	f.end();
}

static void generateProxyMethodDecl(CodeFile& f, Method& m, bool pureVirtual = true)
{
	f.begin();
	f.append("virtual bool %s(", m.getNameC());
	for(size_t i = 0; i < m.fields_.size(); i++)
	{
		Field& field = m.fields_[i];
		f.append("%s%s %s%s", 
			getFieldCppType(field),
			useParamReference(field)?"&":"",
			field.getNameC(),
			(i == m.fields_.size()-1)?"":", ");
	}
	f.append(")%s;", pureVirtual?" = 0":"");
	f.end();
}

static void generateStructDecl(CodeFile& f, Struct* s)
{
	/** 生成struct定义开头部分.处理struct的派生.
	*/
	f.output("// struct %s", s->getNameC());
	if(s->super_)
		f.output("struct %s : public %s", s->getNameC(), s->super_->getNameC());
	else
		f.output("struct %s", s->getNameC());
	f.output("{");
	f.indent();

	/** 生成struct的成员变量列表.
	*/
	f.output("// member list.");
	for(size_t i = 0; i < s->fields_.size(); i++)
	{
		Field& field = s->fields_[i];
		f.output("%s %s;", getFieldCppType(field), field.getNameC());
	}

	/** 生成field id列表. 
		enum
		{
			FID_...,
			FID_...,
			FIDMAX,
		}
	*/
	f.output("// field ids.");
	f.output("enum");
	f.output("{");
	f.indent();
	size_t fid = s->super_?s->super_->getFieldNum():0;
	for(size_t i = 0; i < s->fields_.size(); i++)
	{
		Field& field = s->fields_[i];
		f.output("FID_%s = %d,", field.getNameC(), fid);
		fid++;
	}
	f.output("FIDMAX = %d,", fid);
	f.recover();
	f.output("};");

	/** 自动生成struct的构造函数.
		检查是否需要constructor.如果没有需要产生default value的field，则不需要constructor.
	*/
	bool needCtor = false;
	for(size_t i = 0; i < s->fields_.size(); i++)
	{
		Field& field = s->fields_[i];
		if(getFieldCppDefault(field))
		{
			needCtor = true;
			break;
		}
	}
	if(needCtor)
	{
		f.output("// constructor.");
		f.output("%s();", s->getNameC());
	}

	/** 生成序列化代码.
		void serialize(ProtocolWriter* s) const;
		bool deserialize(ProtocolReader* r);
	*/
	f.output("// serialization.");
	f.output("void serialize(ProtocolWriter* s) const;");
	f.output("// deserialization.");
	f.output("bool deserialize(ProtocolReader* r);");
	f.output("void serializeJson(std::ostream& ss, bool needBracket = true)const;");


	/** cppcode. */
	if(s->cppcode_.length())
	{
		f.output("// cppcode.");
		f.output("%s", s->cppcode_.c_str());
	}
	/** 结束生成.
	*/
	f.recover();
	f.output("};");
}

static void generateStubDecl(CodeFile& f, Service* s)
{
	/** 生成service stub class定义开头部分.
	*/
	f.output("// service stub %s", s->getNameC());
	if(s->super_)
		f.output("class %sStub : public %sStub", s->getNameC(), s->super_->getNameC());
	else
		f.output("class %sStub", s->getNameC());
	f.output("{");
	f.output("public:");
	f.indent();

	/** Methods */
	f.output("// methods.");
	for(size_t i = 0; i < s->methods_.size(); i++)
		generateStubMethodDecl(f, s->methods_[i]);

	if(!s->super_)
	{
		f.recover();
		f.output("protected:");
		f.indent();
		f.output("// events to be processed.");
		f.output("virtual ProtocolWriter* methodBegin() = 0;");
		f.output("virtual void methodEnd() = 0;");
	}
	f.recover();
	f.output("};");
}

static void generateProxyDecl(CodeFile& f, Service* s)
{
	/** 生成service Proxy class定义开头部分.
	*/
	f.output("// service proxy %s", s->getNameC());
	if(s->super_)
		f.output("class %sProxy : public %sProxy", s->getNameC(), s->super_->getNameC());
	else
		f.output("class %sProxy", s->getNameC());
	f.output("{");
	f.output("public:");
	f.indent();

	// methods纯虚函数列表定义.
	f.output("// methods.");
	for(size_t i = 0; i < s->methods_.size(); i++)	
	{
		Method& method = s->methods_[i];
		generateProxyMethodDecl(f, method);
	}
	f.output("");
	
	// dispatch 
	f.output("// dispatch.");
	f.output("bool dispatch(ProtocolReader* reader);");

	// deserialization.
	f.recover();
	f.output("protected:");
	f.indent();

	f.output("// deserialization.");
	for(size_t i = 0; i < s->methods_.size(); i++)
		f.output("bool %s(ProtocolReader* r);", s->methods_[i].getNameC());

	f.recover();
	f.output("};");

}

static void generateEnumDef(CodeFile& f, Enum* e)
{
	f.output("static void initFunc%s(EnumInfo* e)", e->getNameC());
	f.output("{");
	f.indent();
	for(size_t i = 0; i < e->items_.size(); i++)
		f.output("	e->items_.push_back(\"%s\");", e->items_[i].c_str());
	f.recover();
	f.output("}");
	f.output("EnumInfo enum%s(\"%s\", initFunc%s);", 
		e->getNameC(),
		e->getNameC(),
		e->getNameC());
}

static void generateFieldSerialize(CodeFile& f, Field& field, const char* senderName, bool skipComp)
{
	f.output("// serialize %s",field.getNameC());
	if(field.getArray())
	{
		// 数组属性，先序列化数组大小
		if(!skipComp)		f.output("if(%s.size())", field.getNameC());
							f.output("{");
							f.indent();
							f.output("U32 __len__ = (U32)%s.size();", field.getNameC());
							f.output("%s->writeDynSize(__len__); ", senderName);
							f.output("for(U32 i = 0; i < __len__; i++)");
							f.output("{");
							f.indent();
		// 遍历序列化每个元素进行序列化处理.
		if(field.getType() == FT_USER)
		{
							f.output("%s[i].serialize(%s);", field.getNameC(), senderName);
		}
		else if(field.getType() == FT_ENUM)
		{
							f.output("EnumSize __e__ = (EnumSize)%s[i];", field.getNameC());
							f.output("%s->writeType(__e__);", senderName );
		}
		else if(field.getType() == FT_STRING)
		{
							f.output("%s->writeType(%s[i]);", senderName, field.getNameC());
		}
		else
		{
							f.output("%s->writeType(%s[i]);", senderName, field.getNameC());
		}
							f.recover();
							f.output("}");
							f.recover();
							f.output("}");
	}
	else
	{
							f.output("{");
							f.indent();
		if(field.getType() == FT_USER)
		{
							f.output("%s.serialize(%s);", field.getNameC(), senderName);
		}
		else if(field.getType() == FT_STRING)
		{
			if(!skipComp)	f.output("if(%s.length()){", field.getNameC());
							f.output("%s->writeType(%s);", senderName, field.getNameC());
			if(!skipComp)	f.output("}");
		}
		else if(field.getType() == FT_BOOL)
		{ 
			/* bool类型使用fieldmask就可以表示了，这里省略.*/ 
			if(skipComp) 
			{
							f.output("%s->writeType(%s);", senderName, field.getNameC());
			}
		}
		else if(field.getType() == FT_ENUM)
		{
							f.output("EnumSize __e__ = (EnumSize)%s;", field.getNameC());
			if(!skipComp)	f.output("if(__e__){");
							f.output("%s->writeType(__e__);", senderName );
			if(!skipComp)	f.output("}");
		}
		else
		{
			if(!skipComp)	f.output("if(%s != %s){", field.getNameC(), getFieldCppDefault(field));
							f.output("%s->writeType(%s);", senderName, field.getNameC());
			if(!skipComp)	f.output("}");
		}
							f.recover();
							f.output("}");
	}
}

static void generateFieldSerializeJson(CodeFile& f, Field& field)
{
	f.output("// serialize %s", field.getNameC());
	f.output("ss << \"\\\"%s\\\":\";", field.getNameC());
	if (field.getArray())
	{
		// 数组属性，先序列化数组大小
		f.output("{");
		f.indent();
		f.output("ss << \"[\";");
		f.output("U32 __len__ = (U32)%s.size();", field.getNameC());
		f.output("for(U32 i = 0; i < __len__; i++)");
		f.output("{");
		f.indent();
		// 遍历序列化每个元素进行序列化处理.
		if (field.getType() == FT_USER)
		{
			f.output("%s[i].serializeJson(ss);", field.getNameC());
		}
		else if (field.getType() == FT_ENUM)
		{
			f.output("ss << \"\\\"\" << ENUM(%s).getItemName(%s[i]) << \"\\\"\";", field.getUserType()->getNameC(), field.getNameC());
		}
		else if (field.getType() == FT_STRING)
		{
			f.output("ss << \"\\\"\" << %s[i] << \"\\\"\";", field.getNameC());
		}
		else if (field.getType() == FT_BOOL)
		{
			f.output("ss << %s[i] ? \"true\" : \"false\";", field.getNameC());
		}
		else if (field.getType() == FT_FLOAT || field.getType() == FT_DOUBLE)
		{
			f.output("ss << (double)%s[i];", field.getNameC());
		}
		else
		{
			f.output("ss << (S64)%s[i];", field.getNameC());
		}
		
		f.output("ss <<(((i+1) == __len__) ? \"\":\",\")<<\"\";");
		f.recover();
		f.output("}");
		f.output("ss << \"]\";");
		f.recover();
		f.output("}");
	}
	else
	{
		f.output("{");
		f.indent();
		if (field.getType() == FT_USER)
		{
			f.output("%s.serializeJson(ss);", field.getNameC());
		}
		else if (field.getType() == FT_STRING)
		{
			f.output("ss << \"\\\"\" << %s << \"\\\"\";", field.getNameC());
		}
		else if (field.getType() == FT_BOOL)
		{
			f.output("ss << (%s ? \"true\" : \"false\");", field.getNameC());
			
		}
		else if (field.getType() == FT_ENUM)
		{
			f.output("ss << \"\\\"\" << ENUM(%s).getItemName(%s) << \"\\\"\";",field.getUserType()->getNameC() ,field.getNameC());
		}
		else if (field.getType() == FT_FLOAT || field.getType() == FT_DOUBLE)
		{
			f.output("ss << (double)%s;", field.getNameC());
		}
		else 
		{
			f.output("ss << (S64)%s;", field.getNameC());
		}
		f.recover();
		f.output("}");
	}
}

static void generateFieldContainerSerialize(CodeFile& f, FieldContainer* fc, const char* senderName, bool skipComp = false)
{
	if(!fc->fields_.size())
		return;

	if(!skipComp)
	{
		// 首先为这些fields生成 FieldMask 设置代码.
		f.output("//field mask");
		f.output("FieldMask<%d> __fm__;", fc->getFMByteNum());
		for(size_t i = 0; i <fc-> fields_.size(); i++)
		{
			Field& field = fc->fields_[i];
			if(field.getArray())
				f.output("__fm__.writeBit(%s.size()?true:false);", field.getNameC());
			else
			{
				if(field.getType() == FT_USER)
					f.output("__fm__.writeBit(true);");
				else if(field.getType() == FT_STRING)
					f.output("__fm__.writeBit(%s.length()?true:false);", field.getNameC());
				else if(field.getType() == FT_BOOL)
					f.output("__fm__.writeBit(%s);", field.getNameC());
				else
					f.output("__fm__.writeBit((%s==%s)?false:true);", field.getNameC(), getFieldCppDefault(field));
			}
		}
		f.output("%s->write(__fm__.masks_, %d);", senderName, fc->getFMByteNum());
	}

	// 序列化所有的field的内容.
	for(size_t i = 0; i < fc->fields_.size(); i++)
		generateFieldSerialize(f, fc->fields_[i], senderName, skipComp);
}

static void generateFieldContainerSerializeJson(CodeFile& f, FieldContainer* fc)
{
	if (!fc->fields_.size())
		return;

	// 序列化所有的field的内容.
	for (size_t i = 0; i < fc->fields_.size(); i++)
	{
		generateFieldSerializeJson(f, fc->fields_[i]);
		f.output(((i + 1) == fc->fields_.size()) ? "ss<<\"\\n\";" : " ss << \",\\n\";");
	}
}

static void generateFieldDeserialize(CodeFile& f, Field& field, const char* recvName, bool skipComp)
{
	f.output("// deserialize %s", field.getNameC());
	if(field.getArray())
	{
		// 数组属性，先反序列化数组大小
		if(!skipComp)		f.output("if(__fm__.readBit())");
							f.output("{");
							f.indent();
							f.output("U32 __len__;");
							f.output("if(!%s->readDynSize(__len__) || __len__ > %d) return false;", recvName, field.getMaxArrLength());
							f.output("%s.resize(__len__);", field.getNameC());
							// 遍历反序列化每个元素.
							f.output("for(U32 i = 0; i < __len__; i++)");
							f.output("{");
							f.indent();
		if(field.getType() == FT_USER)
		{
							f.output("if(!%s[i].deserialize(%s)) return false;", field.getNameC(), recvName);
		}
		else if(field.getType() == FT_STRING)
		{
							f.output("if(!%s->readType(%s[i], %d)) return false;", recvName, field.getNameC(), field.getMaxStrLength());
		}
		else if(field.getType() == FT_ENUM)
		{
							f.output("EnumSize __e__;");
							f.output("if(!%s->readType(__e__) || __e__ >= %d) return false;", recvName, field.getUserType()->getEnum()->items_.size());
							f.output("%s[i] = (%s)__e__;", field.getNameC(), getFieldCppType(field, false));
		}
		else
		{	
							f.output("if(!%s->readType(%s[i])) return false;", recvName, field.getNameC());
		}
							f.recover();
							f.output("}");
							f.recover();
							f.output("}");
	}
	else
	{
		f.output("{");
		f.indent();
		if(field.getType() == FT_USER)
		{
			if(!skipComp)	f.output("if(__fm__.readBit()){");
							f.output("if(!%s.deserialize(%s)) return false;", field.getNameC(), recvName);
			if(!skipComp)	f.output("}");
		}
		else if(field.getType() == FT_STRING)
		{
			if(!skipComp)	f.output("if(__fm__.readBit()){");
							f.output("if(!%s->readType(%s, %d)) return false;", recvName, field.getNameC(), field.getMaxStrLength());
			if(!skipComp)	f.output("}");
		}
		else if(field.getType() == FT_BOOL)
		{
			if(!skipComp)
			{
							f.output("%s = __fm__.readBit();", field.getNameC());
			}
			else
			{
							f.output("if(!%s->readType(%s)) return false;", recvName, field.getNameC());
			}
		}
		else if(field.getType() == FT_ENUM)
		{
							f.output("EnumSize __e__ = 0;");	//如果readBit为0，则__e__当作0来用
			if(!skipComp)	f.output("if(__fm__.readBit()){");
							f.output("if(!%s->readType(__e__) || __e__ >= %d) return false;", recvName, field.getUserType()->getEnum()->items_.size()); 
							f.output("%s = (%s)__e__;", field.getNameC(), getFieldCppType(field, false));
			if(!skipComp)	f.output("}");
		}
		else
		{
			if(!skipComp)	f.output("if(__fm__.readBit()){");
							f.output("if(!%s->readType(%s)) return false;", recvName, field.getNameC());
			if(!skipComp)	f.output("}");
		}
		f.recover();
		f.output("}");
	}

}

static void generateFieldContainerDeserialize(CodeFile& f, FieldContainer* fc, const char* recvName, bool skipComp = false)
{
	if(!fc->fields_.size())
		return;

	if(!skipComp)
	{
		// 首先读取这些field的 field mask.
		f.output("//field mask");
		f.output("FieldMask<%d> __fm__;", fc->getFMByteNum());
		f.output("if(!%s->read(__fm__.masks_, %d)) return false;", recvName, fc->getFMByteNum());
	}

	// 读取每个field的序列化信息.
	for(size_t i = 0; i < fc->fields_.size(); i++)
		generateFieldDeserialize(f, fc->fields_[i], recvName, skipComp);
}

static void generateStructDef(CodeFile& f, Struct* s)
{
	/** 自动生成struct的构造函数.
		检查是否需要constructor.如果没有需要产生default value的field，则不需要constructor.
	*/
	bool needCtor = false;
	for(size_t i = 0; i < s->fields_.size(); i++)
	{
		Field& field = s->fields_[i];
		if(getFieldCppDefault(field))
		{
			needCtor = true;
			break;
		}
	}
	if(needCtor)
	{
		f.output("%s::%s():", s->getNameC(), s->getNameC());
		// defaults.
		bool needComma = false;
		for(size_t i = 0; i < s->fields_.size(); i++)
		{
			Field& field = s->fields_[i];
			if(!getFieldCppDefault(field))
				continue;
			f.output("%s%s(%s)", needComma?",":"", field.getNameC(), getFieldCppDefault(field));
			needComma = true;
		}
		f.output("{}");
	}

	/** 自动生成序列化代码.
		void serialize(ProtocolWriter* s) const;
	*/
	f.output("void %s::serialize(ProtocolWriter* __s__) const", s->getNameC());
	f.output("{");
	f.indent();
	if(s->super_)
		f.output("%s::serialize(__s__);", s->super_->getNameC());
	generateFieldContainerSerialize(f, s, "__s__", s->skipComp_);
	f.recover();
	f.output("}");

	/** 自动生成反序列化代码.
		bool deserialize(ProtocolReader* r);
	*/
	f.output("bool %s::deserialize(ProtocolReader* __r__)", s->getNameC());
	f.output("{");
	f.indent();
	if(s->super_)
		f.output("if(!%s::deserialize(__r__)) return false;",s-> super_->getNameC());
	generateFieldContainerDeserialize(f, s, "__r__", s->skipComp_);
	f.output("	return true;");
	f.recover();
	f.output("}");

	f.output("void %s::serializeJson(std::ostream& ss, bool needBracket)const", s->getNameC());
	f.output("{");
	f.indent();
	f.output("if(needBracket){ ss << \"{\"; }");
	if (s->super_)
		f.output("%s::serializeJson(ss,false);", s->super_->getNameC());
	generateFieldContainerSerializeJson(f, s);
	f.output("if(needBracket){ ss << \"}\"; }");
	f.recover();
	f.output("}");
}

static void generateStubMethodDef(CodeFile& f, Service*s, Method& m, size_t pid)
{
	f.begin();
	f.append("void %sStub::%s(", s->getNameC(), m.getNameC());
	for(size_t i = 0; i < m.fields_.size(); i++)
	{
		Field& field = m.fields_[i];
		f.append("%s%s%s %s%s", 
			useParamReference(field)?"const ":"",
			getFieldCppType(field),
			useParamReference(field)?"&":"",
			field.getNameC(),
			(i == m.fields_.size()-1)?"":",");
	}
	f.append(")");
	f.end();
	f.output("{");
	f.indent();
	f.output("ProtocolWriter* w = methodBegin();");
	f.output("if(!w) return;");
	f.output("U16 pid = %d;", pid);
	f.output("w->writeType(pid);");
	generateFieldContainerSerialize(f, &m, "w", true);
	f.output("methodEnd();");
	f.recover();
	f.output("}");
}

static void generateStubDef(CodeFile& f, Service* s)
{
	/** Methods */
	size_t methodStartId = s->super_?s->super_->getMethodNum():0;
	for(size_t i = 0; i < s->methods_.size(); i++)
		generateStubMethodDef(f, s, s->methods_[i], methodStartId+i);
}

static void generateProxyMethodDef(CodeFile&f, Method& m, const std::string& svcName)
{
	f.output("bool %sProxy::%s(ProtocolReader* __r__)", svcName.c_str(), m.getNameC());
	f.output("{");
	f.indent();
	for(size_t i = 0; i < m.fields_.size(); i++)
	{
		Field& field = m.fields_[i];
		if(getFieldCppDefault(field))
			f.output("%s %s=%s;", getFieldCppType(field), field.getNameC(), getFieldCppDefault(field));
		else
			f.output("%s %s;", getFieldCppType(field), field.getNameC());
	}
	generateFieldContainerDeserialize(f, &m, "__r__", true);
	f.begin();
	f.append("return %s(", m.getNameC());
	for(size_t i = 0; i < m.fields_.size(); i++)
	{
		Field& field = m.fields_[i];
		f.append("%s%s", field.getNameC(), (i == m.fields_.size()-1)?"":",");
	}
	f.append(");");
	f.end();
	f.recover();
	f.output("}");
}

static void generateProxyDef(CodeFile& f, Service* s)
{

	// methods解编代码.
	for(size_t i = 0; i < s->methods_.size(); i++)
		generateProxyMethodDef(f, s->methods_[i], s->getName());

	// Dispatch 函数定义.
				f.output("bool %sProxy::dispatch(ProtocolReader* r)", s->getNameC());
				f.output("{");
				f.indent();
				f.output("U16 pid;");
				f.output("if(!r->readType(pid)) return false;");
				f.output("switch(pid)");
				f.output("{");
				f.indent();
	std::vector<Method*> methods;
	s->getAllMethods(methods);
	for(size_t i = 0; i < methods.size(); i++)
	{
		Method* method = methods[i];
			f.output("case %d:", i);
			f.output("{");
			f.indent();
			f.output("if(!%s(r)) return false;", method->getNameC());
			f.recover();
			f.output("}");
			f.output("break;");
	}
				f.output("default: return false;");
				f.recover();
				f.output("}");
				f.output("return true;");
				f.recover();
				f.output("}");
}

void CppGenerator::generate()
{
	// H File.
	{
		// 打开文件.
		std::string fn = 
			Compiler::inst().outputDir_ + 
			Compiler::inst().fileStem_ + ".h";
		CodeFile f(fn);

		// 写文件头.
		f.output("/* This file is generated by arpcc, do not change manually! */");
		f.output("#ifndef __%s_h__", Compiler::inst().fileStem_.c_str());
		f.output("#define __%s_h__", Compiler::inst().fileStem_.c_str());
		f.output("#include \"ProtocolWriter.h\"");
		f.output("#include \"ProtocolReader.h\"");
		f.output("#include \"EnumInfo.h\"");
		for(size_t i = 0; i < Compiler::inst().imports_.size(); i++)
		{
			std::string incFilename = Compiler::inst().imports_[i];
			incFilename = incFilename.substr(0,incFilename.find('.'));
			f.output("#include \"%s.h\"", incFilename.c_str());

		}

		// 开始处的cppcode.
		if(Compiler::inst().cppcode_.length())
			f.output("%s", Compiler::inst().cppcode_.c_str());

		// 遍历所有的定义.
		for(size_t i = 0; i < Compiler::inst().definitions_.size(); i++)
		{
			Definition* definition = Compiler::inst().definitions_[i];
			if(definition->getFile() != Compiler::inst().filename_)
				continue;
			f.output("//=============================================================");
			if (definition->getEnum())
				generateEnumDecl(f, definition->getEnum());
			else if (definition->getStruct())
				generateStructDecl(f, definition->getStruct());
			else if (definition->getService())
			{
				Service* s = definition->getService();
				generateStubDecl(f, s);
				generateProxyDecl(f, s);

				/* Methods file.
				methods文件是一个service接口需要被重载的函数声明，派生类在内
				部可以直接包含这个文件，而不用再手动维护与service接口的一致性。
				methods文件并不是必须使用，只是为了维护方便，也可以不使用methods文件，
				而手动的维护这些接口函数。
				*/
				std::string mfn = Compiler::inst().outputDir_ + s->getName() + "Methods.h";
				CodeFile mf(mfn);
				if(s->super_)
					mf.output("#include \"%sMethods.h\"", s->super_->getNameC());
				for(size_t i = 0; i < s->methods_.size(); i++)
					generateProxyMethodDecl(mf, s->methods_[i], false);
			}

		}

		f.output("#endif");
	}
	// Cpp File.
	{
		// 打开文件.
		std::string fn = 
			Compiler::inst().outputDir_ + 
			Compiler::inst().fileStem_ + ".cpp";
		CodeFile f(fn);

		// 写文件头.
		f.output("/* arpcc auto generated cpp file. */");
		f.output("#include \"FieldMask.h\"");
		f.output("#include \"%s.h\"", Compiler::inst().fileStem_.c_str());

		// 遍历所有的定义.
		for(size_t i = 0; i < Compiler::inst().definitions_.size(); i++)
		{
			Definition* definition = Compiler::inst().definitions_[i];
			if(definition->getFile() != Compiler::inst().filename_)
				continue;
			f.output("//=============================================================");
			if (definition->getEnum())
				generateEnumDef(f, definition->getEnum());
			else if (definition->getStruct())
				generateStructDef(f, definition->getStruct());
			else if (definition->getService())
			{
				generateStubDef(f, definition->getService());
				generateProxyDef(f, definition->getService());
			}
		}
	}
}
