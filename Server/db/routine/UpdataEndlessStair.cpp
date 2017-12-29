#include "config.h"
#include "../routine.h"
#include "../handler.h"
#include "../sqltask.h"

U32 
UpdateEndlsee::go(SQLTask *pTask)
{
	static const char * pCode =
		"UPDATE EndlessStair SET PlayerName=? WHERE Rank=?;";

	DBC *dbc = pTask->getDBC();
	SRV_ASSERT(dbc);
	DB_EXEC_GUARD
#if defined(USE_SQLITE)
	CppSQLite3Statement stmt =dbc->compileStatement(pCode);
	stmt.bind(1,playerName_.c_str());
	stmt.bind(2,rank_);
	stmt.execDML();
#elif  defined(USE_MYSQL)
	std::auto_ptr< sql::PreparedStatement > prep_stmt(dbc->prepareStatement(pCode));
	prep_stmt->setString(1,playerName_.c_str());
	prep_stmt->setInt(2,rank_);
	prep_stmt->executeUpdate();
#endif

	DB_EXEC_UNGUARD_RETURN		
		return 0;
}

U32 
UpdateEndlsee::back()
{
	ACE_DEBUG((LM_INFO,ACE_TEXT("Update EndlessStair ok!!! [%s]\n"),playerName_.c_str()));
	return 0;
}
