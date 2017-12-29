#ifndef __SMS_H__
#define __SMS_H__
#include "config.h"
#include "curl/curl.h"
struct SMSContent{
	uint32 playerId_;
	std::string phoneNumber_;
	std::string smsCode_;
};
class SMSTask: public ACE_Task<ACE_MT_SYNCH>{
public:
	SMSTask();
	virtual int init(std::string username, std::string authtoken, std::string appId, std::string tmpId, std::string timeout);
	virtual int fini();
	virtual int svc(void);
	void post();
	void inout(std::vector<SMSContent> &_in, std::vector<SMSContent> &_out);
private:
	bool running_;
	CURL * curl_;
	std::string userName_;
	std::string authToken_;
	std::string appId_;
	std::string tmpId_;
	std::string timeout_;
	std::vector<SMSContent> complate_;
	std::vector<SMSContent> prepare_;
	
	ACE_Recursive_Thread_Mutex  mtx_;
};

#endif