#ifndef __PRODUCT_H__
#define __PRODUCT_H__

#include "config.h"
#include <mysql_connection.h>
#include <mysql_driver.h>
#include <cppconn/resultset.h>
#include <cppconn/statement.h>
#include <cppconn/prepared_statement.h>
#include <ace/Date_Time.h>

namespace sql{ class Connection;}

class RefreshSql
{
	enum {OLD_VERSION = 2};
	enum
	{
		HOST,	///<地址
		USER,	///用户名
		PWD,	///用户密码
		SCHEMA,	///<数据库名
		MAX
	};

	enum
	{
		PLAYERMAXID,	
		BABYMAXID,	
		EMPLOYEEMAXID,	
		GUILDMAXID,	
		MAILMAXID,
		MAXID
	};

public:
	RefreshSql();
	void init();
	void fini();
	void run();
	void playerRun();
	void process(std::string& tableName,int current , int total , int sec);
private:
	bool open_mysql();
	bool close_mysql();
private:
	bool checkSamePlayerName(std::string& name);
	bool checkSamePlayerInstID(U32 instid);
	bool checkSameBabyInstID(U32 instid);
	bool checkSameEmployeeInstID(U32 instid);
	void insertPlayer(std::string &username,SGE_DBPlayerData &data);
	void insertBaby(COM_BabyInst &inst);
	void insertEmployee(COM_EmployeeInst &inst);
private:
	bool checkSameGuildName(std::string &name);
	bool checkSameGuildID(U32 instid);
	void mergeGuild();
	void insertGuild(COM_Guild &guild);
	void mergeGuildMember();
	void insertGuildMember(COM_GuildMember &member);
	void updateGuildMember(U32 oldGuildid,U32 newGuildid);
private:
	void mergeMail();
	void insertMail(COM_Mail &mail);
	bool checkSameMailID(U32 mailid);
private:
	void mergeMall();
	void insertMall(COM_SellItem &item);
	void mergeMallSell();
	void insertMallSell(COM_SelledItem &item);
private:
	void mergeAccount();
	void insertAccount(COM_AccountInfo & info,bool isSeal);
	void checkAccount();
private:
	RefreshSql(RefreshSql const &){}
	RefreshSql const & operator=(RefreshSql const &){}
	int							total_rows_;
	int							trans_rows_;
	sql::Connection*			formmysql_;
	sql::Connection*			tomysql_;
	std::vector<std::string>	fparams_;
	std::vector<std::string>	tparams_;
	std::vector<std::string>	catchPlayerName_;
	std::vector<U32>			catchPlayerinstid_;
	std::vector<U32>			catchGuildinstid_;
	std::vector<U32>			catchMaxid_;
};
#endif ///\end  #ifndef __PRODUCT_H__