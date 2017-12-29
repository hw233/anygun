#include "config.h"
#include "../sqltask.h"
#include "../routine.h"
#include "../handler.h"
#include "../server.h"

U32
QueryPlayer::go(SQLTask *pTask)
{
	hasPlayer_ = false;
	players_.clear();
	std::stringstream sstream;
	sstream << "SELECT * FROM Player WHERE UserName = \"" << username_ << "\" AND InDoorId = " << serverId_  ;
	
	DBC *dbc = pTask->getDBC();
	SRV_ASSERT(dbc);
DB_EXEC_GUARD
#if defined(USE_SQLITE)
	CppSQLite3Query q = dbc->execQuery(sstream.str().c_str());

	if(!q.eof())
	{
		hasPlayer_ = true;

		while(!q.eof())
		{
			U32 instId = q.getIntField("PlayerGuid");
			S32 len=0;
			const unsigned char* pCacheBlob= q.getBlobField("BinData",len);
			SRV_ASSERT(len);
			ProtocolMemReader mr(pCacheBlob,len);
			SGE_DBPlayerData inst;
			inst.instId_ = instId;
			inst.freeze_ =  q.getIntField("Freeze");
			inst.seal_ =  q.getIntField("Seal");
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
					instbaby.ownerName_ = inst.instName_;
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
					instemployee.ownerName_ = inst.instName_;
					inst.employees_.push_back(instemployee);
					qemployees.nextRow();
				}
			}

			players_.push_back(inst);
			q.nextRow();
		}
	}
#elif  defined(USE_MYSQL)
	std::auto_ptr< sql::Statement > stmt(dbc->createStatement());
	std::auto_ptr< sql::ResultSet > res(stmt->executeQuery(sstream.str().c_str()));
	
	while(res->next())
	{
		hasPlayer_ = true;
		sql::SQLString pCacheBlob= res->getString("BinData");
		ProtocolMemReader mr(pCacheBlob->c_str(),pCacheBlob->length());
		SGE_DBPlayerData inst;
		inst.deserialize(&mr);

		inst.instId_ =  res->getInt("PlayerGuid");
		inst.instName_	= res->getString("PlayerName").c_str();
		inst.freeze_ = res->getInt("Freeze");
		inst.seal_ =  res->getInt("Seal");
		{
			std::stringstream ssbaby;
			ssbaby << "SELECT * FROM Baby WHERE OwnerName = \"" << inst.instName_ << "\"";
			std::auto_ptr< sql::Statement > stmt1(dbc->createStatement());
			std::auto_ptr< sql::ResultSet > res1(stmt1->executeQuery(ssbaby.str().c_str()));
			inst.babies_.clear();
			for(size_t i=0; i<inst.babyStorage_.size(); ++i){
				if(inst.babyStorage_[i].instId_ == 0)
					continue;
				ACE_DEBUG((LM_DEBUG, "It's in the babystorage but not in the table babyStorage.ownerName[%s] inst.instName[%s] inst.table %d\n",inst.babyStorage_[i].ownerName_.c_str(),inst.instName_.c_str(),inst.babyStorage_[i].tableId_));
				inst.babyStorage_[i].instId_ = 0;
			}
			while(res1->next())
			{
				sql::SQLString pCacheBlob1= res1->getString("BinData");
				ProtocolMemReader mrbaby(pCacheBlob1->c_str(),pCacheBlob1->length());
				COM_BabyInst instbaby;
				instbaby.deserialize(&mrbaby);
				instbaby.instId_ = res1->getInt("BabyGuid");
				instbaby.ownerName_ = inst.instName_;

				ACE_DEBUG((LM_INFO,ACE_TEXT("baby[%d]===babyowner[%s]===instbaby.slot_==[%d]===table[%d] back\n"),instbaby.instId_,inst.instName_.c_str(),instbaby.slot_,instbaby.tableId_));
				
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
				instemployee.instId_ = res1->getInt("EmployeeGuid");
				instemployee.ownerName_ = inst.instName_;
				inst.employees_.push_back(instemployee);
			}
		}

		players_.push_back(inst);
	}

#endif
DB_EXEC_UNGUARD_RETURN
	return 0;
}

U32
QueryPlayer::back()
{
	ACE_DEBUG((LM_INFO,ACE_TEXT("Query player %s back\n"),username_.c_str()));
	WorldHandler::instance()->queryPlayerOk(username_,players_,serverId_);
	return 0;
}