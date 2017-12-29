#include "config.h"
#include "product.h"

int ACE_TMAIN(int argc, ACE_TCHAR *argv[])
{

	Logger::instance()->init();
	Logger::instance()->enableFileOut("mergeSql.log","log");
	///∂¡»°LUA
	ACE_DEBUG((LM_INFO,ACE_TEXT("Begin load lua...\n")));
	ScriptEnv::init();
#include "ComScriptRegster.h"
#include "ComScriptApi.h"
	std::string err;
	if( !ScriptEnv::loadFile( "env.lua", err ) )
	{
		ACE_DEBUG( (LM_ERROR, ACE_TEXT("load env.lua failed:%s\n"), err.c_str() ) );
		SRV_ASSERT(0);
	}


	RefreshSql rsql;
	rsql.init();
	rsql.run();
	rsql.fini();
	return 0;
}