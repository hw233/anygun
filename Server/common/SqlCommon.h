#ifndef __SQL_COMMON_H__
#define __SQL_COMMON_H__

#ifdef USE_SQLITE
#	include "CppSQLite3.h"
typedef class CppSQLite3DB DBC;
#elif defined( USE_MYSQL )

#	include <mysql_connection.h>
#	include <mysql_driver.h>
#	include <cppconn/resultset.h>
#	include <cppconn/statement.h>
#	include <cppconn/prepared_statement.h>

namespace sql{ class Connection;}
typedef sql::Connection DBC;
#else
#endif 

#endif