#include "ERLGenerator.h"
#include "Compiler.h"
#include "CodeFile.h"

static void generateEnumHRL(CodeFile& f, Enum* e)
{
	f.output("%%%%==================================================================");
	f.output("%%%% enum %s", e->getNameC());
	for(size_t i = 0; i < e->items_.size(); i++)
		f.output("-define(%s_%s, %d).", e->getNameC(), e->items_[i].c_str(), i);
}

static void generateStructRecord(CodeFile& f, Struct* s)
{
	f.output("-record(\'%s\', {", s->getNameC());
	f.indent();
	bool needComma = false;
	if(s->super_)
	{
		f.output("super = #\'%s\'{}", s->super_->getNameC());
		needComma = true;
	}
	for(size_t i = 0; i < s->fields_.size(); i++)
	{
		Field& field = s->fields_[i];
		f.output("%s \'%s\'", needComma?",":"", field.getNameC());
		needComma = true;
	}
	f.recover();
	f.output("}).");
}

static void generateStructFieldMacro(CodeFile& f, Struct* s)
{
	for(size_t i = 0; i < s->fields_.size(); i++)
	{
		Field& field = s->fields_[i];
		f.output("-define(%s_%s, i).", s->getNameC(), field.getNameC());
	}
}

static void generateStructHRL(CodeFile& f, Struct* s)
{
	f.output("%%%%==================================================================");
	f.output("%%%% struct %s", s->getNameC());
	generateStructRecord(f, s);
	generateStructFieldMacro(f, s);
}

static const char* getFieldTypeName(Field& f)
{
	switch(f.getType())
	{
	case FT_INT64:	return "int64";
	case FT_UINT64: return "uint64";
	case FT_DOUBLE: return "double";
	case FT_FLOAT:	return "float";
	case FT_INT32:	return "int32";
	case FT_UINT32: return "uint32";
	case FT_INT16:	return "int16";
	case FT_UINT16: return "uint16";
	case FT_INT8:	return "int8";
	case FT_UINT8:	return "uint8";
	case FT_BOOL:	return "bool";
	case FT_STRING:	return "string";
	case FT_USER:	return f.getUserType()->getNameC();
	case FT_ENUM:	return "enum";
	}
	return "";
}

static void generateStructSerializeCode(CodeFile& f, Struct* s)
{
	f.output("%% serialize");
	f.output("serialize(S = #\'%s\'{}) ->", s->getNameC());
	f.indent();
	if(s->super_)
		f.output("Super = \'%s\':serialize(S#\'%s\'.super),", 
		s->super_->getNameC(), 
		s->getNameC());
	for(size_t i = 0; i < s->fields_.size(); i++)
	{
		Field& field = s->fields_[i];
		if(s->skipComp_)
			f.output("FB_%s = writer:wr_type(\'%s\', %s, S#\'%s\'.\'%s\', false),",
				field.getNameC(),
				getFieldTypeName(field),
				field.getArray()?"true":"false",
				s->getNameC(),
				field.getNameC());
		else
			f.output("{FM_%s, FB_%s} = writer:wr_type(\'%s\', %s, S#\'%s\'.\'%s\', true),",
				field.getNameC(),
				field.getNameC(),
				getFieldTypeName(field),
				field.getArray()?"true":"false",
				s->getNameC(),
				field.getNameC());
	}
	f.output("iolist_to_binary([");
	f.indent();
	if(s->super_)
		f.output("Super,");
	if(!s->skipComp_)
	{
		f.output("<<");
		f.indent();
		for(size_t i = 0; i < s->fields_.size(); i++)
		{
			Field& field = s->fields_[i];
			f.output("FM_%s:1,", field.getNameC());
		}
		f.output("0:%d", s->getFMByteNum()*8 - s->fields_.size());
		f.recover();
		f.output(">>,");
	}
	for(size_t i = 0; i < s->fields_.size(); i++)
	{
		Field& field = s->fields_[i];
		f.output("FB_%s,", field.getNameC());
	}
	f.output("<<>>");
	f.recover();
	f.output("]).");
	f.recover();
}

static size_t getFieldMax(Field& f)
{
	if(f.getType() == FT_STRING)
		return f.getMaxStrLength();
	else if(f.getType() == FT_ENUM)
		return f.getUserType()->getEnum()->items_.size();
	return 0;
}

static void generateStructDeserializeCode(CodeFile& f, Struct* s)
{
	f.output("%% deserialize");
	f.output("deserialize(B0) when is_binary(B0) ->");
	int bNum = 0;
	f.indent();
	if(s->super_)
	{
		f.output("{Super, B%d} = \'%s\':deserialize(B%d),", 
			bNum + 1,
			s->super_->getNameC(), 
			bNum);
		bNum++;
	}
	if(!s->skipComp_)
	{
		f.output("<<");
		f.indent();
		for(size_t i = 0; i < s->fields_.size(); i++)
		{
			Field& field  = s->fields_[i];
			f.output("FM_%s:1,", field.getNameC());
		}
		f.output("_:%d, B%d/binary", s->getFMByteNum()*8 - s->fields_.size(), bNum + 1);
		f.recover();
		f.output(">> = B%d,",bNum);
		bNum++;
	}
	for(size_t i = 0; i < s->fields_.size(); i++)
	{
		Field& field = s->fields_[i];
		f.begin();
		f.append("{FV_%s, B%d} = reader:rd_type(\'%s\', %d, %d, B%d, ",
			field.getNameC(),
			bNum + 1,
			getFieldTypeName(field),
			field.getArray()?field.getMaxArrLength():0,
			getFieldMax(field),
			bNum);
		if(s->skipComp_)
			f.append("false),");
		else
			f.append("FM_%s),", field.getNameC());
		f.end();
		bNum++;
	}
	f.output("{#\'%s\'{", s->getNameC());
	f.indent();
	bool needComma = false;
	if(s->super_)
	{
		f.output("super = Super");
		needComma = true;
	}
	for(size_t i = 0; i < s->fields_.size(); i++)
	{
		Field& field = s->fields_[i];
		f.output("%s\'%s\' = FV_%s", needComma?",":"", field.getNameC(), field.getNameC());
		needComma = true;
	}
	f.recover();
	f.output("}, B%d}.", bNum);
	f.recover();
}

static void generateStructModule(Struct* s)
{
	std::string fn = Compiler::inst().outputDir_ + s->getName() + ".erl";
	CodeFile f(fn);
	f.output("%%%% This file is generated by arpcc, do not change manually!");
	f.output("-module(\'%s\').", s->getNameC());
	f.output("-include(\"%s.hrl\").", Compiler::inst().fileStem_.c_str());
	f.output("-export([default/0, serialize/1, deserialize/1]).");
	// default.
	f.output("%% default");
	f.output("default() -> #?MODULE{}.");

	// serialize.
	generateStructSerializeCode(f, s);
	// deserialize.
	generateStructDeserializeCode(f, s);
}

static void generateServiceStubMethod(CodeFile& f, size_t id, Method* m)
{
	f.begin();
	f.append("\'%s\'(", m->getNameC());
	bool needComma = false;
	for(size_t i = 0; i < m->fields_.size(); i++)
	{
		Field& field = m->fields_[i];
		f.append("%sFV_%s", needComma?",":"", field.getNameC());
		needComma = true;
	}
	f.append(") ->");
	f.end();
	f.indent();
	f.output("iolist_to_binary([");
	f.indent();
	f.output("<<%d:16>>", id);
	for(size_t i = 0; i < m->fields_.size(); i++)
	{
		Field& field = m->fields_[i];
		f.output(",writer:wr_type(\'%s\', %s, FV_%s, false)",
			getFieldTypeName(field),
			field.getArray()?"true":"false",
			field.getNameC());
	}
	f.recover();
	f.output("]).");
	f.recover();
}

static void generateServiceStubModule(Service* s)
{
	std::vector<Method*> methods;
	s->getAllMethods(methods);

	std::string fn = Compiler::inst().outputDir_ + s->getName() + "Stub.erl";
	CodeFile f(fn);
	f.output("%%%% This file is generated by arpcc, do not change manually!");
	f.output("-module(\'%sStub\').", s->getNameC());
	f.output("-include(\"%s.hrl\").", Compiler::inst().fileStem_.c_str());
	f.output("-export([");
	f.indent();
	bool needComma = false;
	for(size_t i = 0; i < methods.size(); i++)
	{
		Method* method = methods[i];
		f.output("%s\'%s\'/%d", needComma?",":"", method->getNameC(), method->fields_.size());
		needComma = true;
	}
	f.recover();
	f.output("]).");

	for(size_t i = 0; i < methods.size(); i++)
		generateServiceStubMethod(f, i, methods[i]);
}

static void generateServiceProxyMethod(CodeFile& f, size_t id, Method* m, bool isLast)
{
	f.output("dispatch(%d, B0, M) ->", id);
	f.indent();
	int bNum = 0;
	for(size_t i = 0; i < m->fields_.size(); i++)
	{
		Field& field = m->fields_[i];
		f.output("{FV_%s, B%d} = reader:rd_type(\'%s\', %d, %d, B%d, false),",
			field.getNameC(),
			bNum + 1,
			getFieldTypeName(field),
			field.getArray()?field.getMaxArrLength():0,
			getFieldMax(field),
			bNum);
		bNum++;
	}
	f.output("M:\'%s\'(", m->getNameC());
	f.indent();
	bool needComma = false;
	for(size_t i = 0; i < m->fields_.size(); i++)
	{
		Field& field = m->fields_[i];
		f.output("%sFV_%s", needComma?",":"", field.getNameC());
		needComma = true;
	}
	f.recover();
	f.output("),");
	f.output("B%d%s", bNum, isLast?".":";");
	f.recover();
}

static void generateServiceProxyModule(Service* s)
{
	std::vector<Method*> methods;
	s->getAllMethods(methods);

	std::string fn = Compiler::inst().outputDir_ + s->getName() + "Proxy.erl";
	CodeFile f(fn);
	f.output("%%%% This file is generated by arpcc, do not change manually!");
	f.output("-module(\'%sProxy\').", s->getNameC());
	f.output("-include(\"%s.hrl\").", Compiler::inst().fileStem_.c_str());
	f.output("-export([behaviour_info/1, dispatch/2]).");

	f.output("%%%% callbacks.");
	f.output("behaviour_info(callbacks) ->");
	f.indent();
	f.output("[");
	bool needComma = false;
	for(size_t i = 0; i < methods.size(); i++)
	{
		Method* method = methods[i];
		f.output("%s{%s, %d}", 
			needComma?",":"", 
			method->getNameC(), 
			method->fields_.size());
		needComma = true;
	}
	f.output("];");
	f.recover();
	f.output("behaviour_info(_) -> undefined.");

	f.output("%%%% dispatch.");
	f.output("dispatch(B, M) when is_binary(B) ->");
	f.indent();
	f.output("<<ID:16, BR/binary>> = B,");
	f.output("dispatch(ID, BR, M).");
	f.recover();

	for(size_t i = 0; i < methods.size(); i++)
		generateServiceProxyMethod(f, i, methods[i], i == methods.size()-1);
}

void ERLGenerator::generate()
{
	// HRL File.
	std::string fn = Compiler::inst().outputDir_ + Compiler::inst().fileStem_ + ".hrl";
	CodeFile f(fn);

	// 写文件头.
	f.output("%%%% This file is generated by arpcc, do not change manually!");
	f.output("-ifndef(%s_HRL).", Compiler::inst().fileStem_.c_str());
	f.output("-define(%s_HRL, true).", Compiler::inst().fileStem_.c_str());
	for(size_t i = 0; i < Compiler::inst().imports_.size(); i++)
	{
		std::string incFilename = Compiler::inst().imports_[i];
		incFilename = incFilename.substr(0,incFilename.find('.'));
		f.output("-include(\"%s.hrl\").", incFilename.c_str());
	}
	// 遍历所有的定义.
	for(size_t i = 0; i < Compiler::inst().definitions_.size(); i++)
	{
		Definition* definition = Compiler::inst().definitions_[i];
		if(definition->getFile() != Compiler::inst().filename_)
			continue;
		if (definition->getEnum())
			generateEnumHRL(f, definition->getEnum());
		else if (definition->getStruct())
		{
			generateStructHRL(f, definition->getStruct());
			generateStructModule(definition->getStruct());
		}
		else if (definition->getService())
		{
			generateServiceStubModule(definition->getService());
			generateServiceProxyModule(definition->getService());
		}
	}

	f.output("-endif.");
}