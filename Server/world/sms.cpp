#include "config.h"
#include "sms.h"
#include "curl/curl.h"
#include <iomanip>
#include "MD5.h"
#include "json/json.h"

inline std::string base64(std::string code){
	char _[512] = {'\0'};
	int len = Base64encode(_,code.c_str(),code.length());
	return std::string(_,len);
}

inline std::string randomcode(){
	std::stringstream ss;
	ss << std::setfill('0') << std::setw(6) << UtlMath::randNM(0,999999);
	return ss.str();
}


size_t smsback(void *buffer , size_t size , size_t nmemb , void *user_p){
	ACE_DEBUG((LM_INFO,"%s\n",(char*)buffer));
	return size * nmemb;
}

SMSTask::SMSTask()
:ACE_Task<ACE_MT_SYNCH>()
,running_(false)
,curl_(NULL)
{

}
int SMSTask::init(std::string username, std::string authtoken, std::string appId, std::string tmpId, std::string timeout){

	userName_ = username;
	authToken_ = authtoken;
	appId_ = appId;
	tmpId_ = tmpId;
	timeout_ = timeout;
	if(NULL == curl_)
		curl_ = curl_easy_init();
	curl_easy_setopt(curl_, CURLOPT_TIMEOUT, 2L);    /*timeout 30s,add by edgeyang*/
	curl_easy_setopt(curl_, CURLOPT_NOSIGNAL, 1L);    /*no signal,add by edgeyang*/
	running_ = true;
	return activate();
}
int SMSTask::fini(){
	if(curl_){
		curl_easy_cleanup(curl_);
	}
	running_ = false;
	return 0;
}
int SMSTask::svc(void){
	Logger::instance()->init();

	enum{
		POST_INTERVAL = 5
	};
	static time_t last = ACE_OS::gettimeofday().sec();
	while(running_){

		time_t curtime = ACE_OS::gettimeofday().sec();
		if(curtime - last >= POST_INTERVAL)
		{
			post();
			last = curtime;
		}
		ACE_OS::sleep(1);
	}
	return 0;
}
void SMSTask::post(){
	ACE_Guard<ACE_Recursive_Thread_Mutex> guard(mtx_);
	if(prepare_.empty())
		return;
	std::string nowtm = TimeFormat::StrLocalTimeNow("%04d%02d%02d%02d%02d%02d");
	MD5 sig(userName_ + authToken_ + nowtm);
	
	std::stringstream sshost;
	sshost << "https://app.cloopen.com:8883/2013-12-26/Accounts/" << userName_ << "/SMS/TemplateSMS?sig="  << sig.toStr();
	
	ACE_DEBUG((LM_INFO,"%s\n",sshost.str().c_str()));
	
	std::string req;

	{
		std::string code = randomcode();
		Json::Value reqbody(Json::objectValue);
		reqbody["appId"] = appId_;
		reqbody["templateId"] = tmpId_;
		Json::Value datas(Json::arrayValue);
		datas.append(Json::Value(code));
		datas.append(Json::Value("2"));
		reqbody["datas"] = datas;
		
		std::stringstream ssphone;

		enum{
			ONESENDMAX = 100
		};
		
		for(size_t i=0; i<ONESENDMAX&&(!prepare_.empty()); ++i){
			SMSContent sc = prepare_.back();
			prepare_.pop_back();
			sc.smsCode_ = code;
			complate_.push_back(sc);
			ssphone << sc.phoneNumber_ << ( ((i+1)<ONESENDMAX&&(!prepare_.empty())) ? ",":"" );
		}

		reqbody["to"] = ssphone.str();
		req = Json::FastWriter().write(reqbody);
	}
	

	struct curl_slist *headers = NULL;
	headers = curl_slist_append(headers,"Accept:application/json;");
	headers = curl_slist_append(headers,"Content-Type:application/json;charset=utf-8;");
	{
		std::stringstream ss;
		ss << "Content-Length:" << req.size();
		headers = curl_slist_append(headers,ss.str().c_str());
	}
	{
		std::stringstream ss;
		ss << "Authorization:" << base64(userName_ + ":" + nowtm);
		headers = curl_slist_append(headers,ss.str().c_str());
	}


	curl_easy_setopt(curl_, CURLOPT_HTTPHEADER, headers);
	curl_easy_setopt(curl_, CURLOPT_SSL_VERIFYPEER, 0L);  
	curl_easy_setopt(curl_, CURLOPT_SSL_VERIFYHOST, 0L);  
	curl_easy_setopt(curl_, CURLOPT_URL, sshost.str().c_str());  
	curl_easy_setopt(curl_, CURLOPT_POSTFIELDS,req.c_str());
	curl_easy_setopt(curl_, CURLOPT_POST, 1); 
	curl_easy_setopt(curl_, CURLOPT_VERBOSE, 1);
	curl_easy_setopt(curl_, CURLOPT_CONNECTTIMEOUT , 5);		///<设置连接超时时间 10 秒
	curl_easy_setopt(curl_, CURLOPT_WRITEFUNCTION,&smsback);///<设置数据回调 

	CURLcode result = curl_easy_perform(curl_);
	curl_slist_free_all(headers);
	switch(result)
	{
	case CURLE_OK:						///<一切安好 

		break;
	case CURLE_UNSUPPORTED_PROTOCOL:	///<不支持协议URL头决定
		//ACE_DEBUG((LM_ERROR , ACE_TEXT("curl perform CURLE_UNSUPPORTED_PROTOCOL \n")));
		return ;
	case CURLE_COULDNT_CONNECT:			///<不能连接到主机或者代理
		//ACE_DEBUG((LM_ERROR , ACE_TEXT("curl perform CURLE_COULDNT_CONNECT \n")));
		return ;
	case CURLE_REMOTE_ACCESS_DENIED:	///<访问被拒绝
		//ACE_DEBUG((LM_ERROR , ACE_TEXT("curl perform CURLE_REMOTE_ACCESS_DENIED \n")));
		return ;
	case CURLE_HTTP_RETURNED_ERROR:		///<http返回错误
		//ACE_DEBUG((LM_ERROR , ACE_TEXT("curl perform CURLE_HTTP_RETURNED_ERROR \n")));
		return ;
	default:
		//ACE_DEBUG((LM_ERROR , ACE_TEXT("curl perform unkown error code(%d) \n"), (int)result));
		return ;
	}

	enum { HTTP_RESPONSE_SUCCESS = 200U };

	long httpCode = 0;
	curl_easy_getinfo(curl_,CURLINFO_RESPONSE_CODE,&httpCode);	///<获得HTTP 返回值,目前只要不等于 200 就返回-1再次请求

	if(HTTP_RESPONSE_SUCCESS!=httpCode)
	{
		//ACE_DEBUG((LM_ERROR , ACE_TEXT("curl CURLINFO_RESPONSE_CODE (%q)\n"),httpCode));
		return ;
	}
}


void SMSTask::inout(std::vector<SMSContent> &_in, std::vector<SMSContent> &_out){
	ACE_Guard<ACE_Recursive_Thread_Mutex> guard(mtx_);
	while(!_in.empty()){
		SMSContent sc = _in.back();
		_in.pop_back();
		bool maked = false;
		for(size_t i=0; i<prepare_.size(); ++i){
			if(prepare_[i].playerId_ == sc.playerId_){
				prepare_[i].phoneNumber_ = sc.phoneNumber_;
				maked = true;
				break;
			}
		}
		if(!maked){
			prepare_.push_back(sc);
		}
	}
	if(!complate_.empty()){
		_out.insert(_out.end(),complate_.begin(),complate_.end());
		complate_.clear();
	}
}

