#include "config.h"
#include "../routine.h"
#include "../handler.h"
#include "../server.h"
#include "../sqltask.h"

//U32
//Queryidgen::go(SQLTask *pTask)
//{
//	ACE_DEBUG((LM_DEBUG ,"Queryidgen go\n"));
//	std::stringstream sstream;
//	sstream << "SELECT * FROM KeyGiftTable WHERE cdkey = \"" << idgen_ << "\"";
//
//	DBC *dbc = pTask->getDBC();
//	SRV_ASSERT(dbc);
//	hasIdgen_ = false;
//	DB_EXEC_GUARD
//#if defined(USE_SQLITE)
//	CppSQLite3Query q = dbc->execQuery(sstream.str().c_str());
//
//	if(!q.eof())
//	{
//		while(!q.eof())
//		{
//			content_.key_		 = q.getStringField("cdkey");
//			/*if(content_.key_ != idgen_)
//				hasIdgen_ = true;*/
//			content_.playerName_ = q.getStringField("playerName");
//			content_.giftname_	 = q.getStringField("giftName");
//			int len=0;
//			const unsigned char* pBlob = q.getBlobField("rewardItem" , len);
//			ProtocolMemReader protocol2(pBlob , len);
//			protocol2.readVector(content_.items_);
//			q.nextRow();
//		}
//	}
//#elif  defined(USE_MYSQL)
//	std::auto_ptr< sql::Statement > stmt(dbc->createStatement());
//	std::auto_ptr< sql::ResultSet > res(stmt->executeQuery(sstream.str().c_str()));
//	while(res->next())
//	{
//		content_.key_ = res->getString("cdkey");
//		/*if(content_.key_ != idgen_)
//			hasIdgen_ = true;*/
//		content_.playerName_ = res->getString("playerName");
//		content_.giftname_ = res->getString("giftName");
//		sql::SQLString tmpBlob = res->getString("rewardItem");
//		const  char * pBlob= tmpBlob.c_str();
//		ProtocolMemReader protocol1((S8*)pBlob , tmpBlob.length());
//		protocol1.readVector(content_.items_);
//	}
//
//#endif
//
//	DB_EXEC_UNGUARD_RETURN
//		return 0;
//}
//
//U32
//Queryidgen::back()
//{
//	WorldHandler::instance()->queryIdgenOK(playerName_,content_,hasIdgen_);
//	return 0;
//}

//////////////////////////////////////////////////////////////////////////
//U32
//UpdateKeyGift::go(SQLTask *pTask)
//{
//	enum {BUFFER_SIZE = 1024*8};
//	
//	static const char * pCode =
//		"UPDATE KeyGiftTable SET playerName=?,useTime=? WHERE cdkey=?;";
//
//	DBC *dbc = pTask->getDBC();
//	SRV_ASSERT(dbc);
//	DB_EXEC_GUARD
//#if defined(USE_SQLITE)
//	CppSQLite3Statement stmt =dbc->compileStatement(pCode);
//	stmt.bind( 1, giftdata_.playerName_.c_str());
//	stmt.bind( 2, int(giftdata_.usetime_));
//	stmt.bind( 3, giftdata_.key_.c_str());
//	stmt.execDML();
//#elif  defined(USE_MYSQL)
//		std::auto_ptr< sql::PreparedStatement > prep_stmt(dbc->prepareStatement(pCode));
//	prep_stmt->setString( 1 , giftdata_.playerName_.c_str());
//	prep_stmt->setInt( 2, int(giftdata_.usetime_));
//	prep_stmt->setString( 3, giftdata_.key_.c_str());
//
//	prep_stmt->executeUpdate();
//#endif
//
//	return 0;
//	DB_EXEC_UNGUARD_RETURN		
//}
//
//U32
//QueryKeyGift::go(SQLTask *pTask)
//{
//	DBC *dbc = pTask->getDBC();
//	SRV_ASSERT(dbc);
//	std::vector<COM_KeyContent> contents;
//	DB_EXEC_GUARD
//
//	std::string str="SELECT * FROM KeyGiftTable";
//#if defined(USE_SQLITE)
//	CppSQLite3Query q = dbc->execQuery(str.c_str());
//	while (!q.eof())
//	{
//		COM_KeyContent   gift;
//		gift.pfid_ = q.getStringField("pfIp");
//		gift.giftname_ = q.getStringField("giftName");
//		gift.key_ = q.getStringField("cdkey");
//		gift.playerName_ = q.getStringField("playerName");
//		gift.usetime_ = q.getInt64Field("useTime");
//
//		int len=0;
//		const unsigned char* pBlob = q.getBlobField("rewardItem" , len);
//		ProtocolMemReader protocol1((S8*)pBlob , len);
//		protocol1.readVector(gift.items_);
//		Server::instance()->giftCache_.push_back(gift);
//		q.nextRow();
//	}
//#elif  defined(USE_MYSQL)
//	std::auto_ptr< sql::Statement > stmt(dbc->createStatement());
//	std::auto_ptr< sql::ResultSet > res(stmt->executeQuery(str.c_str()));
//	while (res->next()) 
//	{
//		COM_KeyContent   gift;
//		gift.pfid_		= res->getString("pfIp");
//		gift.giftname_	= res->getString("giftName");
//		gift.key_		= res->getString("cdkey");
//		gift.playerName_= res->getString("playerName");
//		gift.usetime_	= res->getInt64("useTime");
//		sql::SQLString tmpBlob1 = res->getString("rewardItem");
//		const  char * pBlob1= tmpBlob1.c_str();
//		ProtocolMemReader protocol2(pBlob1 , tmpBlob1.length());
//		protocol2.readVector(gift.items_);
//		Server::instance()->giftCache_.push_back(gift);
//	}
//#endif
//	return 0;
//	DB_EXEC_UNGUARD_RETURN
//}