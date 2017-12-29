#include "PYGenerator.h"
#include "Compiler.h"
#include "CodeFile.h"

static void generateEnum(CodeFile& f, Enum* e)
{
	f.output("# %s", e->getNameC());
	f.output("class %s(object):", e->getNameC());
	f.indent();
	for(size_t i = 0; i < e->items_.size(); i++)
		f.output("%s = %d", e->items_[i].c_str(), i);
	f.recover();
}
static const char* getFieldDefault(Field& f)
{
	if(f.getArray())
		return "[]";
	if(f.getType() == FT_BOOL)
		return "False";
	else if(f.getType() == FT_STRING)
		return "\"\"";
	else if(f.getType() == FT_USER)
	{
		static std::string un;
		un = f.getUserType()->getName() + "()";
		return un.c_str();
	}
	else
		return "0";
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

static size_t getFieldValMax(Field& f)
{
	if(f.getType() == FT_STRING)
		return f.getMaxStrLength();
	else if(f.getType() == FT_ENUM)
		return f.getUserType()->getEnum()->items_.size();
	return 0;
}

static void generateStruct(CodeFile& f, Struct* s)
{
	f.output("# %s", s->getNameC());
	f.output("class %s(%s):", s->getNameC(), s->super_?s->super_->getNameC():"object");
	f.indent();
	// __init__
	f.output("def __init__(self):");
	f.indent();
	if(s->super_)
		f.output("%s.__init__(self)", s->super_->getNameC());
	for(size_t i = 0; i < s->fields_.size(); i++)
	{
		Field& field = s->fields_[i];
		f.output("self.%s = %s", field.getNameC(), getFieldDefault(field));
	}
	f.recover();
	// serialize
	f.output("def serialize(self, _b_):");
	f.indent();
	if(s->super_)
		f.output("%s.serialize(self, _b_)", s->super_->getNameC());
	if(!s->skipComp_)
	{
		f.output("_fm_ = FieldMaskWriter(%d)", s->getFMByteNum());
		f.output("_pfm_ = len(_b_)");
		f.output("_b_.append(\'\')");
	}
	else
		f.output("_fm_ = None");
	for(size_t i = 0; i < s->fields_.size(); i++)
	{
		Field& field = s->fields_[i];
		f.output("write(%sWriter, %s, _b_, self.%s, _fm_)",
			getFieldTypeName(field),
			field.getArray()?"True":"False",
			field.getNameC()
			);
	}
	if(!s->skipComp_)
		f.output("_b_[_pfm_] = _fm_.write()");
	f.recover();
	// deserialize
	f.output("def deserialize(self, _b_, _p_):");
	f.indent();
	if(s->super_)
		f.output("_p_ = %s.deserialize(self, _b_, _p_)", s->super_->getNameC());
	if(!s->skipComp_)
	{
		f.output("_fm_ = FieldMaskReader(_b_, _p_, %d)", s->getFMByteNum());
		f.output("_p_ += %d", s->getFMByteNum());
	}
	else
		f.output("_fm_ = None");
	for(size_t i = 0; i < s->fields_.size(); i++)
	{
		Field& field = s->fields_[i];
		f.output("self.%s, _p_= read(%sReader, _b_, _p_, %d, %d, _fm_)",
			field.getNameC(),
			getFieldTypeName(field),
			field.getArray()?field.getMaxArrLength():0,
			getFieldValMax(field)
			);
	}
	f.output("return _p_");
	f.recover();
	f.recover();
	// Writer & Reader.
	f.output("def %sWriter(b, v, fm):", s->getNameC());
	f.indent();
	f.output("v.serialize(b)");
	f.recover();
	f.output("def %sReader(b, p, valMax, fm):", s->getNameC());
	f.indent();
	f.output("v = %s()", s->getNameC());
	f.output("p = v.deserialize(b, p)");
	f.output("return v, p");
	f.recover();
}

static void generateServiceStubMethod(CodeFile& f, size_t id, Method& m)
{
	f.begin();
	f.append("def %s(self", m.getNameC());
	for(size_t i = 0; i < m.fields_.size(); i++)
	{
		Field& field = m.fields_[i];
		f.append(",%s", field.getNameC());
	}
	f.append("):");
	f.end();
	f.indent();
	f.output("_b_ = []", id);
	f.output("uint16Writer(_b_, %d, None)", id);
	for(size_t i = 0; i < m.fields_.size(); i++)
	{
		Field& field = m.fields_[i];
		f.output("write(%sWriter, %s, _b_, %s, None)",
			getFieldTypeName(field),
			field.getArray()?"True":"False",
			field.getNameC()
			);
	}
	f.output("self.call(_b_)");
	f.recover();
}

static void generateServiceStub(CodeFile& f, Service* s)
{
	f.output("# %s stub", s->getNameC());
	if(s->super_)
		f.output("class %sStub(%sStub):", s->getNameC(), s->super_->getNameC());
	else
		f.output("class %sStub(object):", s->getNameC());
	f.indent();
	size_t mid = s->super_?s->super_->getMethodNum():0;
	for(size_t i = 0; i < s->methods_.size(); i++, mid++)
		generateServiceStubMethod(f, mid, s->methods_[i]);
	f.recover();
}

static void generateServiceProxy(CodeFile& f, Service* s)
{
	f.output("# %s proxy", s->getNameC());
	if(s->super_)
		f.output("class %sProxy(%sProxy):", s->getNameC(), s->super_->getNameC());
	else
		f.output("class %sProxy(object):", s->getNameC());
	f.indent();
	f.output("def dispatch(self, _b_):");
	f.indent();
	f.output("_id_, _p_ = uint16Reader(_b_, 0, 0, None)");
	f.output("self.dispatchID(_id_, _b_, _p_)");
	f.recover();
	f.output("def dispatchID(self, _id_, _b_, _p_):");
	f.indent();
	size_t mid = s->super_?s->super_->getMethodNum():0;
	for(size_t m = 0; m < s->methods_.size(); m++, mid++)
	{
		Method& method = s->methods_[m];
		f.output("if _id_ == %d:", mid);
		f.indent();
		for(size_t fid = 0; fid < method.fields_.size(); fid++)
		{
			Field& field = method.fields_[fid];
			f.output("%s, _p_= read(%sReader, _b_, _p_, %d, %d, None)",
				field.getNameC(),
				getFieldTypeName(field),
				field.getArray()?field.getMaxArrLength():0,
				getFieldValMax(field)
				);
		}
		f.begin();
		f.append("self.%s(", method.getNameC());
		bool needComma = false;
		for(size_t fid = 0; fid < method.fields_.size(); fid++)
		{
			Field& field = method.fields_[fid];
			f.append("%s%s", needComma?",":"", field.getNameC());
			needComma = true;
		}
		f.append(")");
		f.end();
		f.output("return");
		f.recover();
	}
	if(s->super_)
		f.output("%sProxy.dispatchID(self, _id_, _b_, _p_)", s->super_->getNameC());
	f.recover();
	f.recover();
}

static void generateService(CodeFile& f, Service* s)
{
	generateServiceStub(f, s);
	generateServiceProxy(f, s);
}

void PYGenerator::generate()
{
	// Python File.
	std::string fn = Compiler::inst().outputDir_ + Compiler::inst().fileStem_ + ".py";
	CodeFile f(fn);

	// import.
	f.output("from arpc.writer import *");
	f.output("from arpc.reader import *");
	for(std::set<std::string>::iterator iter = Compiler::inst().importedFiles_.begin();
		iter != Compiler::inst().importedFiles_.end(); ++iter)
	{
		std::string incFilename = *iter;
		incFilename = incFilename.substr(0,incFilename.find('.'));
		f.output("from %s import *", incFilename.c_str());
	}

	// 遍历所有的定义.
	for(size_t i = 0; i < Compiler::inst().definitions_.size(); i++)
	{
		Definition* definition = Compiler::inst().definitions_[i];
		if(definition->getFile() != Compiler::inst().filename_)
			continue;
		if (definition->getEnum())
			generateEnum(f, definition->getEnum());
		else if (definition->getStruct())
			generateStruct(f, definition->getStruct());
		else if (definition->getService())
			generateService(f, definition->getService());
	}
}