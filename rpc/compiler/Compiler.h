#ifndef __Compiler_h__
#define __Compiler_h__
#include <cstring>
#include <stdlib.h>
#include <algorithm>
#include <string>
#include <vector>
#include <set>
#include <map>
#include <string>

namespace String
{
	static void Trim(std::string& str, bool left, bool right)
	{
		static const std::string delims = " \t\r\n";
		if (right)
			str.erase(str.find_last_not_of(delims) + 1); // trim right
		if (left)
			str.erase(0, str.find_first_not_of(delims)); // trim left
	}

	static void ToLowerCase(std::string& str)
	{
		std::transform(
			str.begin(),
			str.end(),
			str.begin(),
			tolower);
	}

	static void ToUpperCase(std::string& str)
	{
		std::transform(
			str.begin(),
			str.end(),
			str.begin(),
			toupper);
	}

	static std::vector<std::string > Split(std::string const &str, std::string const&delims, unsigned maxSplits = 0)
	{
		std::vector<std::string> ret;
		ret.reserve(maxSplits ? maxSplits + 1 : 1000);

		unsigned int numSplits = 0;
		size_t start, pos;
		start = 0;
		do
		{
			pos = str.find_first_of(delims, start);
			if (pos == start)
			{
				start = pos + 1;
			}
			else if (pos == std::string::npos || (maxSplits && numSplits == maxSplits))
			{
				ret.push_back(str.substr(start));
				break;
			}
			else
			{
				ret.push_back(str.substr(start, pos - start));
				start = pos + 1;
			}
			start = str.find_first_not_of(delims, start);
			++numSplits;

		} while (pos != std::string::npos);
		return ret;
	}

	static void Split(const char* line, char delims, std::vector<std::string>& out, bool skip = true)
	{
		size_t i = 0, p = 0, s = strlen(line);
		char word[1024] = { '\0' };
		if (0 == s)
			return;
		do
		{
			if (line[i] == delims)
			{
				word[p] = '\0';
				out.push_back(word);

				if (!skip)
				{
					word[0] = delims;
					word[1] = '\0';
					out.push_back(word);
				}

				p = 0;
			}
			else
			{
				word[p++] = line[i];
			}
			++i;
		} while (i < s);

		word[p] = '\0';
		out.push_back(word);
	}

	static std::string	Join(std::vector< std::string > arr, std::string const&delims)
	{
		std::string ret = "";
		if (arr.empty())
			return ret;

		for (size_t i = 0, len = arr.size(); i < len; ++i)
		{
			ret += arr[i];
			if (1 + i != len)
			{
				ret += delims;
			}
		}

		return ret;
	}

	static int IndexOf(std::string const& str, char ch)
	{
		size_t i = str.find_first_of(ch);
		if (i == std::string::npos)
			return -1;
		return i;
	}
}

namespace File
{
	static std::string GetParentPath(const std::string fn)
	{
		if (fn.empty())
			return "./";

		size_t idx = fn.find_last_of('/');
		if (idx != std::string::npos)
		{
			return fn.substr(0, idx);
		}
		else
		{
			return "./";
		}
	}

	static std::string GetFileName(const std::string& fn)
	{
		if (fn.empty())
			return "";

		std::string _fn = fn;
		if (_fn.find_last_of('/') != std::string::npos)
		{
			std::vector<std::string> pathpart = String::Split(_fn, "/");
			_fn = pathpart.back();
		}
		return _fn;
	}

	static std::string GetFileExt(const std::string& fn)
	{
		if (fn.empty())
			return "";

		std::string _fn = fn;

		if (_fn.find_last_of('.') != std::string::npos)
		{
			std::vector<std::string> parts = String::Split(_fn, ".");
			return parts.back();
		}
		else
			return "";
	}

	static std::string GetFileBaseName(const std::string& fn)
	{
		if (fn.empty())
			return "";

		std::string _fn = fn;
		if (_fn.find_last_of('/') != std::string::npos)
		{
			std::vector<std::string> pathpart = String::Split(_fn, "/");
			_fn = pathpart.back();
		}
		if (_fn.find_last_of('.') != std::string::npos)
		{
			std::vector<std::string> part = String::Split(_fn, ".");
			part.pop_back();
			return part.back();
		}
		else
		{
			return _fn;
		}
	}
}

#include "Definition.h"
#include "Field.h"
#include "Enum.h"
#include "Struct.h"
#include "Service.h"

/** arpc语言编译器类.
	Compiler 通过flex和bison对源文件进行分析，构建符合arpc语法规范的抽象语法
	树，然后生成对应的代码文件.
*/
class Generator;
class Compiler
{
public:
	/** 获得全局引用. */
	static Compiler& inst();

	Compiler();
	~Compiler();

	/** 编译源文件. */
	int compile();

	/** 根据名称查找一个定义. 
	*/
	Definition* findDefinition( const std::string& name );

	/** 输出编译错误. */
	void outputError(const char* e, ...);
	/** 输出当前文件行编译错误. */
	void outputErrorFL(const char* e, ...);

	/** @name 配置信息. */
	//@{
	/** 源文件完整路径名称. */
	std::string						inputFileName_;
	/** 生成代码文件的目录. */
	std::string						outputDir_;
	/** 用于寻找import文件的目录. */
	std::vector<std::string>		importPaths_;
	/** 代码生成器名称 .*/
	std::string						generator_;
	/** 当前编译的根文件名称. */
	std::string						filename_;
	/** 当前编译的文件名stem. */
	std::string						fileStem_;
	//@}

	/** 根文件直接包含的文件名称. */
	std::vector<std::string>		imports_;
	/** 所有已经被包含过的文件名称，防止重复包含. */
	std::set<std::string>			importedFiles_;
	/** 当前被解析的文件名称. */
	std::string						curFilename_;
	/** 当前被解析文件的行号. */
	int								curLineno_;
	/** 开始处的cppcode. */
	std::string						cppcode_;
	/** 这个源文件包含的所有arpc定义. */
	std::vector<Definition*>	definitions_;
	// 当前编译状态.
	Enum						curEnum_;
	Struct						curStruct_;
	Service					curService_;
	Method						curMethod_;
	Field						curField_;
};

#endif//__Compiler_h__