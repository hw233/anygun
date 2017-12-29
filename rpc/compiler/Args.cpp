#include "Args.h"
#include <cstring>
#include<stdlib.h>
static const char Deli = ';';

Args::Args(int argc, char* argv[], const char* shortopts)
{
	SetShortKey(shortopts);
	Scan(argc, argv);
}

void Args::SetShortKey(const char* shortopts)
{
	for (size_t i = 0, l = strlen(shortopts); i < l; ++i)
	{
		if (shortopts[i] == Deli)
		{
			continue;
		}

		Token token;
		token.ShortKey = shortopts[i];
		Tokens.push_back(token);
	}
}

void Args::SetLongKey(char shortkey, const char* longkey)
{
	for (size_t i = 0; i < Tokens.size(); ++i)
	{
		if (Tokens[i].ShortKey == shortkey)
		{
			Tokens[i].LongKey = longkey;
			return;
		}
	}
}

void Args::Scan(int argc, char* argv[])
{
	char shortkey = '\0';
	std::string longkey = "  ";
	for (int i = 0; i < argc; ++i)
	{
		size_t argsize = strlen(argv[i]);
		if (argv[i][0] == '-')
		{
			shortkey = '\0';
			longkey = "  ";

			if (argv[i][1] == '-' && argsize > 2)
			{
				longkey = (argv[i] + 2);
			}
			else
			{
				shortkey = argv[i][1];
			}
		}
		else
		{
			Token* ptoken = GetToken(shortkey);
			if (NULL == ptoken)
				ptoken = GetToken(longkey.c_str());
			if (ptoken != NULL)
			{
				ptoken->Values.push_back(argv[i]);
			}
		}
	}
}

Args::Token* Args::GetToken(char shortkey)
{
	for (size_t i = 0; i < Tokens.size(); ++i)
	{
		if (Tokens[i].ShortKey == shortkey)
		{
			return &Tokens[i];
		}
	}
	return NULL;
}

Args::Token* Args::GetToken(const char* longkey)
{
	for (size_t i = 0; i < Tokens.size(); ++i)
	{
		if (Tokens[i].LongKey == longkey)
		{
			return &Tokens[i];
		}
	}
	return NULL;
}

const char* Args::GetCString(char shortkey, int index)
{
	Token* ptoken = GetToken(shortkey);
	if (NULL == ptoken)
		return NULL;
	if (index >= ptoken->Values.size())
		return NULL;
	return ptoken->Values[index].c_str();
}

const char* Args::GetCString(const char* longkey, int index)
{
	Token* ptoken = GetToken(longkey);
	if (NULL == ptoken)
		return NULL;
	if (index >= ptoken->Values.size())
		return NULL;
	return ptoken->Values[index].c_str();
}

int Args::GetInteger(char shortkey, int index)
{
	const char* c = GetCString(shortkey, index);
	if (NULL == c)
		return 0;
	return atoi(c);
}

int Args::GetInteger(const char* longkey, int index)
{
	const char* c = GetCString(longkey, index);
	if (NULL == c)
		return 0;
	return atoi(c);
}

float Args::GetFloat(char shortkey, int index)
{
	const char* c = GetCString(shortkey, index);
	if (NULL == c)
		return 0.F;
	return (float)atof(c);
}

float Args::GetFloat(const char* longkey, int index)
{
	const char* c = GetCString(longkey, index);
	if (NULL == c)
		return 0.F;
	return (float)atof(c);
}