
#include "config.h"
#include "dbhandler.h"
#include "account.h"
#include "worldserv.h"
#include "EndlessStair.h"
#include "Guild.h"
#include "loghandler.h"
#include "Activity.h"
#include "GameEvent.h"
#include "LotteryTable.h"
#include "EmployeeQuestSystem.h"

int DBHandler::open(void * p){
	int r = Connection::open(p);
	if(Env::get<int>(V_IsNewServer)){
		GameEvent::procGameEvent(GET_NewServer,0,0,0);
	}
	return r;
}

int 
DBHandler::handle_close(ACE_HANDLE handle, ACE_Reactor_Mask close_mask)
{
	ACE_DEBUG((LM_INFO,ACE_TEXT("DB serv closed...!!!\n")));
	int ret = Connection::handle_close(handle,close_mask);
	SRV_ASSERT(0);
	return ret;
}

bool
DBHandler::syncGlobalGuid(U32 id){
	if(WorldServ::instance()->maxGuid_ < id){
		ACE_DEBUG((LM_INFO,"SYNC Max Guid %d \n",id));
		WorldServ::instance()->maxGuid_ = id;
	}
	return true;
}

bool
DBHandler::syncEmployeeQuest(std::vector< SGE_PlayerEmployeeQuest >& info){
	EmployeeQuestSystem::initEmployeeQuest(info);
	return true;
}

bool
DBHandler::syncContactInfo(std::vector< SGE_ContactInfoExt >& info)
{
	ACE_DEBUG((LM_DEBUG,"Sync player infos \n"));
	WorldServ::instance()->addContactInfo(info);
	return true;
}

bool 
DBHandler::queryPlayerOk(STRING& username, std::vector< SGE_DBPlayerData >& insts, int32 inDoorId)
{
	Account *account = Account::getAccountByName(username);
	if(NULL == account)
	{
		ACE_DEBUG((LM_INFO,ACE_TEXT("Account is offline  %s...!!!\n"),username.c_str()));
		return true;
	}
	account->setDBPlayers(insts);
	return true;
}


bool 
DBHandler::createPlayerSameName(STRING& username){
	Account *account = Account::getAccountByName(username);
	if(NULL == account)
	{
		ACE_DEBUG((LM_INFO,ACE_TEXT("Account is offline  %s...!!!\n"),username.c_str()));
		return true;
	}
	account->createPlayerSameName();
	return true;
}

bool 
DBHandler::createPlayerOk(STRING& username,SGE_DBPlayerData &inst, int32 inDoorId)
{
	Account *account = Account::getAccountByName(username);
	if(NULL == account)
	{
		ACE_DEBUG((LM_INFO,ACE_TEXT("Account is offline  %s...!!!\n"),username.c_str()));
		return true;
	}
	SGE_ContactInfoExt info;
	info.accName_ = username;
	info.instId_ = inst.instId_;
	info.name_   = inst.instName_;
	info.level_  = inst.properties_[PT_Level];
	info.assetId_= inst.properties_[PT_AssetId];
	info.job_ = JT_Newbie;
	info.jobLevel_ = 1;
	info.rolefirst_ = inst.createTime_;
	info.rolelast_ = inst.loginTime_;
	info.serverid_ = inDoorId; 
	WorldServ::instance()->addContactInfo(info);
	account->createPlayerOk(inst);

	return true;
}

bool
DBHandler::queryPlayerByIdOK(std::string& playername,SGE_DBPlayerData& inst, bool isFlag)
{
	if(isFlag)
		EndlessStair::instance()->getPlayerDBData(inst);
	else
	{
		Player* player = Player::getPlayerByName(playername);
		if(NULL == player)
			return true;
		player->queryPlayerInstOK(inst);
	}
		
	return true;
}

bool
DBHandler::queryEndlessStairAllDateOK(std::vector< STRING >& name)
{
	EndlessStair::instance()->getEndlessStairDate(name);
	return true;
}

bool
DBHandler::queryPlayerByLevelOK(std::vector< COM_ContactInfo >& info)
{
	WorldServ::instance()->syncPlayerLevelRank(info);
	return true;
}

bool
DBHandler::queryPlayerByFFOK(std::vector< COM_ContactInfo >& info)
{
	WorldServ::instance()->syncPlayerFFRank(info);
	return true;
}

bool
DBHandler::createBabyOK(STRING& playername,COM_BabyInst& inst,bool isToStorage)
{
	ACE_DEBUG((LM_DEBUG,"DB INSERT BABY BK 0\n"));
	Player* player = Player::getPlayerByName(playername);
	if(NULL == player)
		return true;
	
	if(isToStorage)
	{
			ACE_DEBUG((LM_DEBUG,"DB INSERT BABY BK 1\n"));
		player->depositBaby(inst);
	}else
	{
			ACE_DEBUG((LM_DEBUG,"DB INSERT BABY BK 2\n"));
		player->addBaby(inst);
	}
	return true;
}

bool
DBHandler::deleteBabyOK(STRING& playername,S32 babyInstId)
{
	Player* player = Player::getPlayerByName(playername);
	if(NULL == player)
		return true;
	player->delBabyOk(babyInstId);
	return true;
}

bool
DBHandler::queryBabyByFFOK(std::vector< COM_BabyRankData >& infos)
{
	WorldServ::instance()->fatchBabyRankOK(infos);
	return true;
}

bool
DBHandler::queryBabyByIdOK(std::string& name, COM_BabyInst& inst)
{
	Player* player = Player::getPlayerByName(name);
	if(NULL == player)
		return true;
	CALL_CLIENT(player,queryBabyOK(inst));
	return true;
}

bool
DBHandler::UpdateBabysOK(std::string& playername){
	Player* player = Player::getPlayerByName(playername);
	if(NULL == player)
		return true;
	player->isSorting_ = false;
	return true;
}

bool
DBHandler::createEmployeeOK(STRING& playername,COM_EmployeeInst& inst){
	return true;
}

bool
DBHandler::deleteEmployeeOK(STRING& playername,std::vector< U32 >& instIds)
{
	Player* player = Player::getPlayerByName(playername);
	if(NULL == player)
		return true;
	player->delEmployeeOk(instIds);
	return true;
}

bool
DBHandler::queryEmployeeByFFOK(std::vector< COM_EmployeeRankData >& infos)
{
	WorldServ::instance()->fatchEmpRankOK(infos);
	return true;
}

bool
DBHandler::queryEmployeeByIdOK(std::string& name, COM_EmployeeInst& inst)
{
	Player* player = Player::getPlayerByName(name);
	if(NULL == player)
		return true;
	CALL_CLIENT(player,queryEmployeeOK(inst));
	return true;
}

bool
DBHandler::appendMail(std::string& playerName, std::vector< COM_Mail >& mails){
	Player* player = Player::getPlayerByName(playerName);
	if(NULL == player)
		return true;
	player->appendMail(mails);
	return true;
}

bool
DBHandler::syncGuild(std::vector< COM_Guild >& guilds)
{
	for (size_t i=0;i<guilds.size();i++)
	{
		COM_Guild&	guild=guilds[i];

		if(NULL!=Guild::findGuildById(guild.guildId_))
		{
			ACE_DEBUG((LM_ERROR ," guild id confict %d \n",guild.guildId_));
			SRV_ASSERT(0);
		}
		else
		{
			Guild::addGuild(guild);
		}
	}
	return true;
}

bool
DBHandler::syncGuildMember(std::vector< COM_GuildMember >& guildMember)
{
	for (size_t i=0;i<guildMember.size();i++)
	{
		Guild::addGuildMember(guildMember[i]);
	}
	return true;
}

bool
DBHandler::insertGuildOK(COM_Guild& guild, COM_GuildMember& guildMember)
{
	Player* player = Player::getPlayerByName(guild.masterName_);
	if(NULL == player)
		return true;
	//Çå³ýÕ¼Î» 
	Guild::delCritical(guild.guildName_ ,guildMember.roleId_ );
	Guild* newGuild=Guild::addGuild(guild);
	SRV_ASSERT(newGuild);
	Guild::addGuildMember(guildMember);
	newGuild->updateMemberList(player);
	newGuild->updateGuild();
	CALL_CLIENT(player,createGuildOK());
	
	return true;
}

bool
DBHandler::updateMemberJobOk(S32 roleId, GuildJob job)
{
	Guild* pGuild = Guild::findGuildByPlayer(roleId);
	SRV_ASSERT(pGuild);
	COM_GuildMember* pMember = pGuild->findMember(roleId);
	SRV_ASSERT(pMember);
	pMember->job_ = job;
	pGuild->updateMember(*pMember,MLF_ChangePosition,true);
	return true;
}

bool
DBHandler::queryIdgenOK(std::string& playerName, COM_KeyContent& content, bool isHas)
{
	return true;
}

bool
DBHandler::fatchActivity(SGE_SysActivity& date)
{
	if(!Env::get<int>(V_IsNewServer))
		Activity::sycnallActData(date);
	return true;
}