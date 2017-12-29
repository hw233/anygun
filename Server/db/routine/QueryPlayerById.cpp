
#include "config.h"
#include "../sqltask.h"
#include "../routine.h"
#include "../handler.h"
#include "../server.h"

U32
QueryPlayerById::go(SQLTask *pTask)
{
	ACE_DEBUG((LM_INFO,ACE_TEXT("QueryPlayerById player %d \n"),playerGuid_));
	ACE_DEBUG((LM_DEBUG ,"QueryPlayerById go\n"));
	std::stringstream sstream;
	sstream << "SELECT * FROM Player WHERE playerGuid = \"" << playerGuid_ << "\"";

	DBC *dbc = pTask->getDBC();
	SRV_ASSERT(dbc);
	DB_EXEC_GUARD
#if defined(USE_SQLITE)
		CppSQLite3Query q = dbc->execQuery(sstream.str().c_str());

	if(!q.eof())
	{
		while(!q.eof())
		{
			U32 instId = q.getIntField("PlayerGuid");
			S32 len=0;
			const unsigned char* pCacheBlob= q.getBlobField("BinData",len);
			SRV_ASSERT(len);
			ProtocolMemReader mr(pCacheBlob,len);
			SGE_DBPlayerData inst;
			inst.instId_ = instId;
			inst.deserialize(&mr);

			{
				std::stringstream ssbaby;
				ssbaby << "SELECT * FROM Baby WHERE OwnerName = \"" << inst.instName_ << "\"";
				CppSQLite3Query qbaby = dbc->execQuery(ssbaby.str().c_str());
				inst.babies_.clear();
				while(!qbaby.eof())
				{
					const unsigned char* pCacheBlobBaby= qbaby.getBlobField("BinData",len);
					ProtocolMemReader mrbaby(pCacheBlobBaby,len);
					COM_BabyInst instbaby;
					instbaby.deserialize(&mrbaby);
					inst.babies_.push_back(instbaby);
					qbaby.nextRow();
				}
			}

			{
				std::stringstream ssemployee;
				ssemployee << "SELECT * FROM Employee WHERE OwnerName = \"" << inst.instName_ << "\"";
				CppSQLite3Query qemployees = dbc->execQuery(ssemployee.str().c_str());
				inst.employees_.clear();
				while(!qemployees.eof())
				{
					const unsigned char* pCacheBlobEmployee= qemployees.getBlobField("BinData",len);
					ProtocolMemReader mremployee(pCacheBlobEmployee,len);
					COM_EmployeeInst instemployee;
					instemployee.deserialize(&mremployee);
					inst.employees_.push_back(instemployee);
					qemployees.nextRow();
				}
			}

			player_ = inst;
			q.nextRow();
		}
	}
#elif  defined(USE_MYSQL)
	std::auto_ptr< sql::Statement > stmt(dbc->createStatement());
	std::auto_ptr< sql::ResultSet > res(stmt->executeQuery(sstream.str().c_str()));

	while(res->next())
	{
		U32 instId = res->getInt("PlayerGuid");
		S32 len=0;
		sql::SQLString pCacheBlob= res->getString("BinData");
		ProtocolMemReader mr(pCacheBlob->c_str(),pCacheBlob->length());
		SGE_DBPlayerData inst;
		inst.instId_ = instId;
		inst.deserialize(&mr);

		{
			std::stringstream ssbaby;
			ssbaby << "SELECT * FROM Baby WHERE OwnerName = \"" << inst.instName_ << "\"";
			std::auto_ptr< sql::Statement > stmt1(dbc->createStatement());
			std::auto_ptr< sql::ResultSet > res1(stmt1->executeQuery(ssbaby.str().c_str()));
			inst.babies_.clear();
			while(res1->next())
			{
				sql::SQLString pCacheBlob1= res1->getString("BinData");
				ProtocolMemReader mrbaby(pCacheBlob1->c_str(),pCacheBlob1->length());
				COM_BabyInst instbaby;
				instbaby.deserialize(&mrbaby);
				inst.babies_.push_back(instbaby);
			}
		}

		{
			std::stringstream ssemployee;
			ssemployee << "SELECT * FROM Employee WHERE OwnerName = \"" << inst.instName_ << "\"";
			std::auto_ptr< sql::Statement > stmt1(dbc->createStatement());
			std::auto_ptr< sql::ResultSet > res1(stmt1->executeQuery(ssemployee.str().c_str()));
			inst.employees_.clear();
			while(res1->next())
			{
				sql::SQLString pCacheBlob1= res1->getString("BinData");
				ProtocolMemReader mremployee(pCacheBlob1->c_str(),pCacheBlob1->length());
				COM_EmployeeInst instemployee;
				instemployee.deserialize(&mremployee);
				inst.employees_.push_back(instemployee);
			}
		}

		player_ = inst;
	}
#endif

	DB_EXEC_UNGUARD_RETURN
		return 0;
}

U32
QueryPlayerById::back()
{
	ACE_DEBUG((LM_INFO,ACE_TEXT("QueryPlayerById player %d back\n"),playerGuid_));
	WorldHandler::instance()->queryPlayerByIdOK(initiator_,player_,hasPlayer_);
	
	return 0;
}