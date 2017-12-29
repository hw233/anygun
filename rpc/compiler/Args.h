#ifndef __OPTIONS_H__
#define __OPTIONS_H__
#include <string>
#include <vector>

class Args
{
	struct Token
	{
		char					 ShortKey;
		std::string				 LongKey;
		std::vector<std::string> Values;
	};
public:
	Args(int argc, char* argv[], const char* shortopts);

	void SetShortKey(const char* shortopts);
	
	void SetLongKey(char shortkey, const char* longkey);

	void Scan(int argc, char* argv[]);

	Token* GetToken(char shortkey);
	Token* GetToken(const char* longkey);

	const char* GetCString(char shortkey, int index = 0);
	const char* GetCString(const char* longkey, int index = 0);

	int GetInteger(char shortkey, int index = 0);
	int GetInteger(const char* longkey, int index = 0);

	float GetFloat(char shortkey, int index = 0);
	float GetFloat(const char* longkey, int index = 0);

private:
	std::vector<Token> Tokens;
};

#endif