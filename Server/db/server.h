
#ifndef __SERVER_H__
#define __SERVER_H__
#include "config.h"
#include "sqltask.h"
#include "handler.h"

class Server 
	: public ACE_Service_Object
{
public:
	SINGLETON_FUNCTION(Server)
public:
	int init (int argc, ACE_TCHAR *argv[]);
	int fini (void);
	int handle_timeout (const ACE_Time_Value &current_time, const void *act);
	void initData();
	void initRank();
public:
	U32 maxGuid_;
	std::map<U32,COM_ActivityTable>		tables_;
	std::vector<SGE_ContactInfoExt>		contactCache_;
	std::vector<COM_ContactInfo>		playerFFrankCache_;
	std::vector<COM_ContactInfo>		playerlevelrankCache_;
	std::vector<COM_BabyRankData>		babyFFrankCache_;
	std::vector<COM_EmployeeRankData>	employeeFFrankCache_;
	std::vector<SGE_PlayerEmployeeQuest> employeeQuestCache_;  
	std::vector<COM_KeyContent>			giftCache_;
};

#endif