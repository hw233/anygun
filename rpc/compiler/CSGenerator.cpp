#include "CSGenerator.h"
#include "Compiler.h"
#include "CodeFile.h"



static const char* getFieldTypeName(Field& f)
{
	switch(f.getType())
	{
	case FT_INT64:
		return "long";
	case FT_UINT64:
		return "ulong";
	case FT_DOUBLE:
		return "double";
	case FT_FLOAT:
		return "float";
	case FT_INT32:
		return "int";
	case FT_UINT32:
		return "uint";
	case FT_INT16:
		return "short";
	case FT_UINT16:
		return "ushort";
	case FT_INT8:
		return "sbyte";
	case FT_UINT8:
		return "byte";
	case FT_BOOL:
		return "bool";
	case FT_STRING:
		return "string";
	case FT_USER:
	case FT_ENUM:
		return f.getUserType()->getNameC();
	default:
		throw "Invalid field type.";
	}
}

static void generateEnum(CodeFile& f, Enum* e)
{
	f.output("public enum %s : %s", e->getNameC(),getFieldTypeName(e->getSuperType()));
	f.output("{");
	f.indent();
	for(size_t i = 0; i < e->items_.size(); i++)
		f.output("%s,", e->items_[i].c_str());
	f.recover();
	f.output("}");
}

static void generateSingleFieldSerializeCode(CodeFile& f, Field& field, const char* wn, bool skipComp)
{
	if(field.getType() == FT_USER)
		f.output("%s.serialize(%s);", field.getNameC(), wn);
	else if(field.getType() == FT_STRING)
	{
		if(!skipComp)
		{
			f.output("if(%s.Length > 0)", field.getNameC());
			f.indent();
		}
		f.output("bin.ProtocolWriter.writeType(%s, %s);", wn, field.getNameC());
		if(!skipComp)
			f.recover();
	}
	else if(field.getType() == FT_BOOL)
	{
		if(skipComp)
			f.output("bin.ProtocolWriter.writeType(%s, %s);", wn, field.getNameC());
	}
	else if(field.getType() == FT_ENUM)
	{
		f.output("byte __e__ = (byte)%s;", field.getNameC());
		if(!skipComp)
		{
			f.output("if(__e__ != 0)");
			f.indent();
		}
		f.output("bin.ProtocolWriter.writeType(%s, __e__);", wn);
		if(!skipComp)
			f.recover();
	}
	else
	{
		if(!skipComp)
		{
			f.output("if(%s != 0)", field.getNameC());
			f.indent();
		}
		f.output("bin.ProtocolWriter.writeType(%s, %s);", wn, field.getNameC());
		if(!skipComp)
			f.recover();
	}
}

static void generateArrayFieldSerializeCode(CodeFile& f, Field& field, const char* wn, bool skipComp)
{
	// Array size.
	if(!skipComp)
	{
		f.output("if(%s != null && %s.Length > 0)", field.getNameC(), field.getNameC());
		f.output("{");
		f.indent();
	}
	f.output("uint __len__ = (%s == null)?0:(uint)%s.Length;", field.getNameC(), field.getNameC());
	f.output("bin.ProtocolWriter.writeDynSize(%s, __len__);", wn);
	f.output("for(uint i = 0; i < __len__; i++)");
	f.output("{");
	f.indent();
	if(field.getType() == FT_USER)
		f.output("%s[i].serialize(%s);", field.getNameC(), wn);
	else if(field.getType() == FT_ENUM)
		f.output("bin.ProtocolWriter.writeType(%s, (byte)%s[i]);", wn, field.getNameC());
	else
		f.output("bin.ProtocolWriter.writeType(%s, %s[i]);", wn, field.getNameC());
	f.recover();
	f.output("}");
	if(!skipComp)
	{
		f.recover();
		f.output("}");
	}
}

static void generateFieldSerializeCode(CodeFile& f, Field& field, const char* wn, bool skipComp)
{
	f.output("{");
	f.indent();
	if(field.getArray())
		generateArrayFieldSerializeCode(f, field, wn, skipComp);
	else
		generateSingleFieldSerializeCode(f, field, wn, skipComp);
	f.recover();
	f.output("}");
}

static void generateFieldContainerSerializeCode(CodeFile& f, FieldContainer* fc, const char* wn, bool skipComp)
{
	if(!fc->fields_.size())
		return;
	if(!skipComp)
	{
		f.output("bin.FieldMask __fm__ = new bin.FieldMask(new byte[%d]);", fc->getFMByteNum());
		for(size_t i = 0; i < fc->fields_.size(); i++)
		{
			Field& field = fc->fields_[i];
			if(field.getArray())
				f.output("__fm__.writeBit((%s==null)?false:(%s.Length>0?true:false));",
				field.getNameC(), field.getNameC());
			else
			{
				if(field.getType() == FT_USER)
					f.output("__fm__.writeBit(true);");
				else if(field.getType() == FT_STRING)
					f.output("__fm__.writeBit(%s.Length>0?true:false);", field.getNameC());
				else if(field.getType() == FT_BOOL)
					f.output("__fm__.writeBit(%s);", field.getNameC());
				else
					f.output("__fm__.writeBit(%s==0?false:true);", field.getNameC());
			}
		}
		f.output("bin.ProtocolWriter.writeType(%s, __fm__.getBits());", wn);
	}
	for(size_t i = 0; i < fc->fields_.size(); i++)
		generateFieldSerializeCode(f, fc->fields_[i], wn, skipComp);
}

static void generateArrayFieldDeserializeCode(CodeFile& f, Field& field, const char* rn, bool skipComp)
{
	if(!skipComp)
	{
		f.output("if(__fm__.readBit())");
		f.output("{");
		f.indent();
	}
	f.output("uint __len__;");
	f.output("if(!bin.ProtocolReader.readDynSize(%s, out __len__) || __len__ > %d) return false;", rn, field.getMaxArrLength());
	f.output("%s = new %s[__len__];", field.getNameC(), getFieldTypeName(field));
	f.output("for(uint i = 0; i < __len__; i++)");
	f.output("{");
	f.indent();
	if(field.getType() == FT_USER)
	{
		f.output("%s[i] = new %s();", field.getNameC(), getFieldTypeName(field));
		f.output("if(!%s[i].deserialize(%s)) return false;", field.getNameC(), rn);
	}
	else if(field.getType() == FT_STRING)
	{
		f.output("if(!bin.ProtocolReader.readType(%s, out %s[i], %d)) return false;", rn, field.getNameC(), field.getMaxStrLength());
	}
	else if(field.getType() == FT_ENUM)
	{
		f.output("byte __e__;");
		f.output("if(!bin.ProtocolReader.readType(%s, out __e__) || __e__ >= %d) return false;", rn, field.getUserType()->getEnum()->items_.size());
		f.output("%s[i] = (%s)__e__;", field.getNameC(), getFieldTypeName(field));
	}
	else
	{	
		f.output("if(!bin.ProtocolReader.readType(%s, out %s[i])) return false;", rn, field.getNameC());
	}
	f.recover();
	f.output("}");
	if(!skipComp)
	{
		f.recover();
		f.output("}");
	}
}

static void generateSingleFieldDeserializeCode(CodeFile& f, Field& field, const char* rn, bool skipComp)
{
	if(field.getType() == FT_USER)
	{
		if(!skipComp) 
			f.output("__fm__.readBit();");
		f.output("if(!%s.deserialize(%s)) return false;", field.getNameC(), rn);
	}
	else if(field.getType() == FT_STRING)
	{
		if(!skipComp)	
		{
			f.output("if(__fm__.readBit())");
			f.output("{");
			f.indent();
		}
		f.output("if(!bin.ProtocolReader.readType(%s, out %s, %d)) return false;", rn, field.getNameC(), field.getMaxStrLength());
		if(!skipComp)	
		{
			f.recover();
			f.output("}");
		}
	}
	else if(field.getType() == FT_BOOL)
	{
		if(!skipComp)
			f.output("%s = __fm__.readBit();", field.getNameC());
		else
			f.output("if(!bin.ProtocolReader.readType(%s, out %s)) return false;", rn, field.getNameC());
	}
	else if(field.getType() == FT_ENUM)
	{
		if(!skipComp)	
		{
			f.output("if(__fm__.readBit())");
			f.output("{");
			f.indent();
		}
		f.output("byte __e__ = 0;");
		f.output("if(!bin.ProtocolReader.readType(%s, out __e__) || __e__ >= %d) return false;", rn, field.getUserType()->getEnum()->items_.size()); 
		f.output("%s = (%s)__e__;", field.getNameC(), getFieldTypeName(field));
		if(!skipComp)	
		{
			f.recover();
			f.output("}");
		}
	}
	else
	{
		if(!skipComp)	
		{
			f.output("if(__fm__.readBit())");
			f.output("{");
			f.indent();
		}
		f.output("if(!bin.ProtocolReader.readType(%s, out %s)) return false;", rn, field.getNameC());
		if(!skipComp)	
		{
			f.recover();
			f.output("}");
		}
	}
}

static void generateFieldDeserializeCode(CodeFile& f, Field& field, const char* rn, bool skipComp)
{
	f.output("{");
	f.indent();
	if(field.getArray())
		generateArrayFieldDeserializeCode(f, field, rn, skipComp);
	else
		generateSingleFieldDeserializeCode(f, field, rn, skipComp);
	f.recover();
	f.output("}");
}

static void generateFieldContainerDeserializeCode(CodeFile& f, FieldContainer* fc, const char* rn, bool skipComp)
{
	if(!fc->fields_.size())
		return;
	if(!skipComp)
	{
		f.output("byte[] __fmbits__;");
		f.output("if(!bin.ProtocolReader.readType(%s, out __fmbits__, %d)) return false;", rn, fc->getFMByteNum());
		f.output("bin.FieldMask __fm__ = new bin.FieldMask(__fmbits__);");
	}
	for(size_t i = 0; i < fc->fields_.size(); i++)
		generateFieldDeserializeCode(f, fc->fields_[i], rn, skipComp);
}

static void generateStruct(CodeFile& f, Struct* s)
{
	if(s->super_)
		f.output("public class %s : %s", s->getNameC(), s->super_->getNameC());
	else
		f.output("public class %s", s->getNameC());
	f.output("{");
	f.indent();
	// fields.
	for(size_t i = 0; i < s->fields_.size(); i++)
	{
		Field& field = s->fields_[i];
		if(field.getArray())
			f.output("public %s[] %s = new %s[0];", 
			getFieldTypeName(field), 
			field.getNameC(),
			getFieldTypeName(field));
		else
		{
			if(field.getType() == FT_USER)
				f.output("public %s %s = new %s();", 
				getFieldTypeName(field), 
				field.getNameC(), 
				getFieldTypeName(field));
			else if(field.getType() == FT_STRING)
				f.output("public string %s = \"\";", field.getNameC());
			else
				f.output("public %s %s;", getFieldTypeName(field), field.getNameC());
		}
	}
	/** Field ids. */
	f.output("// member ids.");
	f.output("public enum FID");
	f.output("{");
	f.indent();
	size_t fid = s->super_?s->super_->getFieldNum():0;
	for(size_t i = 0; i < s->fields_.size(); i++)
	{
		Field& field = s->fields_[i];
		f.output("%s = %d,", field.getNameC(), fid);
		fid++;
	}
	f.output("MAX = %d,", fid);
	f.recover();
	f.output("}");

	// serialize code.
	f.output("public %s void serialize(bin.IWriter w)", s->super_?"new":"");
	f.output("{");
	f.indent();
	if(s->super_)
		f.output("base.serialize(w);");
	generateFieldContainerSerializeCode(f, s, "w", s->skipComp_);
	f.recover();
	f.output("}");
	// deserialize code.
	f.output("public %s bool deserialize(bin.IReader r)", s->super_?"new":"");
	f.output("{");
	f.indent();
	if(s->super_)
		f.output("base.deserialize(r);");
	generateFieldContainerDeserializeCode(f, s, "r", s->skipComp_);
	f.output("return true;");
	f.recover();
	f.output("}");
	// field serialize.
	f.output("public %s bool serializeField(uint fid, bin.IWriter w)", s->super_?"new":"");
	f.output("{");
	f.indent();
	f.output("switch(fid)");
	f.output("{");
	f.indent();
	for(size_t i = 0; i < s->fields_.size(); i++)
	{
		Field& field = s->fields_[i];
		f.output("case (uint)FID.%s:", field.getNameC());
		f.output("{");
		f.indent();
		generateFieldSerializeCode(f, field, "w", true);
		f.recover();
		f.output("}");
		f.output("return true;");
	}
	f.recover();
	f.output("}");
	if(s->super_)
		f.output("return base.serializeField(fid, w);");
	else
		f.output("return false;");
	f.recover();
	f.output("}");
	// field deserialize.
	f.output("public %s bool deserializeField(uint fid, bin.IReader r)", s->super_?"new":"");
	f.output("{");
	f.indent();
	f.output("switch(fid)");
	f.output("{");
	f.indent();
	for(size_t i = 0; i < s->fields_.size(); i++)
	{
		Field& field = s->fields_[i];
		f.output("case (uint)FID.%s:", field.getNameC());
		f.output("{");
		f.indent();
		generateFieldDeserializeCode(f, field, "r", true);
		f.recover();
		f.output("}");
		f.output("return true;");
	}
	f.recover();
	f.output("}");
	if(s->super_)
		f.output("return base.deserializeField(fid, r);");
	else
		f.output("return false;");
	f.recover();
	f.output("}");

	f.recover();
	f.output("}");
}

static void generateStubMethod(CodeFile& f, Service* s, Method& m, size_t mid)
{
	f.begin();
	f.append("public void %s(", m.getNameC());
	for(size_t i = 0; i < m.fields_.size(); i++)
	{
		Field& field = m.fields_[i];
		f.append("%s%s %s%s", 
			getFieldTypeName(field), 
			field.getArray()?"[]":"",
			field.getNameC(),
			(i == m.fields_.size()-1)?"":",");
	}
	f.append(")");
	f.end();
	f.output("{");
	f.indent();

	f.output("bin.IWriter w = methodBegin();");
	f.output("if(w == null) return;");
	f.output("ushort __pid__ = %d;", mid);
	f.output("bin.ProtocolWriter.writeType(w, __pid__);");
	generateFieldContainerSerializeCode(f, &m, "w", true);
	f.output("methodEnd();");

	f.recover();
	f.output("}");
}

static void generateServiceStub(CodeFile& f, Service* s)
{
	if(s->super_)
		f.output("public abstract class %sStub : %sStub", s->getNameC(), s->super_->getNameC());
	else
		f.output("public abstract class %sStub", s->getNameC());
	f.output("{");
	f.indent();
	if(!s->super_)
	{
		f.output("protected abstract bin.IWriter methodBegin();");
		f.output("protected abstract void methodEnd();");
	}
	// methods.
	size_t methodStartId = s->super_?s->super_->getMethodNum():0;
	for(size_t i = 0; i < s->methods_.size(); i++)
		generateStubMethod(f, s, s->methods_[i], methodStartId + i);
	f.recover();
	f.output("}");
}

static void generateProxyAbstractMethod(CodeFile& f, Method& m)
{
	f.begin();
	f.append("bool %s(", m.getNameC());
	for(size_t i = 0; i < m.fields_.size(); i++)
	{
		Field& field = m.fields_[i];
		f.append("%s%s %s%s",
			getFieldTypeName(field),
			field.getArray()?"[]":"",
			field.getNameC(),
			(i == m.fields_.size()-1)?"":",");
	}
	f.append(");");
	f.end();
}

static void generateMethodDispatcher(CodeFile& f, Service* s, Method& m)
{
	f.output("public static bool %s(bin.IReader __r__, %sProxy __p__)", m.getNameC(), s->getNameC());
	f.output("{");
	f.indent();
	for(size_t i = 0; i < m.fields_.size(); i++)
	{
		Field& field = m.fields_[i];
		if(field.getType() == FT_USER && !field.getArray())
			f.output("%s %s = new %s();", getFieldTypeName(field), field.getNameC(), getFieldTypeName(field));
		else
			f.output("%s%s %s;", getFieldTypeName(field), field.getArray()?"[]":"", field.getNameC());
	}
	generateFieldContainerDeserializeCode(f, &m, "__r__", true);
	f.begin();
	f.append("return __p__.%s(", m.getNameC());
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

static void generateServiceProxy(CodeFile& f, Service* s)
{
	if(s->super_)
		f.output("public interface %sProxy : %sProxy", s->getNameC(), s->super_->getNameC());
	else
		f.output("public interface %sProxy", s->getNameC());
	f.output("{");
	f.indent();

	// methods.
	for(size_t i = 0; i < s->methods_.size(); i++)
		generateProxyAbstractMethod(f, s->methods_[i]);
	f.recover();
	f.output("}");
}

static void generateServiceDispatcher(CodeFile& f, Service* s)
{
	f.output("public static class %sDispatcher", s->getNameC());
	f.output("{");
	f.indent();

	// deserializations.
	for(size_t i = 0; i < s->methods_.size(); i++)
		generateMethodDispatcher(f, s, s->methods_[i]);

	// dispatch function.
	f.output("public static bool dispatch(bin.IReader r, %sProxy p)", s->getNameC());
	f.output("{");
	f.indent();
	f.output("ushort pid;");
	f.output("if(!bin.ProtocolReader.readType(r, out pid)) return false;");
	f.output("switch(pid)");
	f.output("{");
	f.indent();
	std::vector<Service*> parents;
	s->getParents(parents);
	size_t mid = 0;
	for(size_t i = 0; i < parents.size(); i++)
	{
		Service* parent = parents[i];
		for(size_t m = 0; m < parent->methods_.size(); m++)
		{
			Method& method = parent->methods_[m];
			f.output("case %d:", mid++);
			f.output("{");
			f.indent();
			f.output("if(!%sDispatcher.%s(r, p)) return false;", parent->getNameC(), method.getNameC());
			f.recover();
			f.output("}");
			f.output("break;");
		}
	}
	f.output("default: return false;");
	f.recover();
	f.output("}");
	f.output("return true;");
	f.recover();
	f.output("}");
	f.recover();
	f.output("}");
}

static void generateService(CodeFile& f, Service* s)
{
	generateServiceStub(f, s);
	f.output("//=============================================================");
	generateServiceProxy(f, s);
	f.output("//=============================================================");
	generateServiceDispatcher(f, s);
}

void CSGenerator::generate()
{
	std::string fn = 
		Compiler::inst().outputDir_ + 
		Compiler::inst().fileStem_ + ".cs";
	CodeFile f(fn);

	f.output("/* This file is generated by arpcc, do not change manually! */");

	// iterate all definations.
	for(size_t i = 0; i < Compiler::inst().definitions_.size(); i++)
	{
		Definition* definition = Compiler::inst().definitions_[i];
		if(definition->getFile() != Compiler::inst().filename_)
			continue;
		f.output("//=============================================================");
		if (definition->getEnum())
			generateEnum(f, definition->getEnum());
		else if (definition->getStruct())
			generateStruct(f, definition->getStruct());
		else if (definition->getService())
			generateService(f, definition->getService());
	}
}
