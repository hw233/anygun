#include "config.h"
#include "AnySDKWrapper.h"
#include "json/json.h"
#include "HttpParser.h"
#include "player.h"
#include "Shop.h"
#include "account.h"
#include "dbhandler.h"

std::vector<AnyOauth*> LoginOauth::infos_;



LoginOauth::LoginOauth(){
}

LoginOauth::~LoginOauth(){
}

int
LoginOauth::open(void* p)
{
	int r = Connection::open(p);

	ACE_INET_Addr const &remote = getRemoteAddr();

	return r;
}

int 
LoginOauth::handleReceived(void* data, size_t size)
{
	ACE_Message_Block mb((char*)data,size);
	mb.wr_ptr(size);
	HTTP_Request req;
	req.parse_request(mb);
	proc(req.data());
	//mb.release();
	return size;
}

int
LoginOauth::handle_close(ACE_HANDLE handle, ACE_Reactor_Mask close_mask)
{
	if(ACE_OS::last_error() == EWOULDBLOCK){
		return 0;
	}
	status_ = Connection::RemoteClosed;

	delete this;
	return 0;
}

void LoginOauth::proc( const char* post )
{
	ACE_DEBUG((LM_DEBUG,"%s\n",post));
	Json::Value json;
	Json::Reader reader;
	if(reader.parse(post,json)){
	}

	if(!json.isObject())
	{
		ACE_DEBUG((LM_ERROR,"if(!json.isObject())\n"));
		close();
		return;
	}

	AnyOauth info;

	if(json["status"] != "ok")
	{
		ACE_DEBUG((LM_ERROR,"if(json[status] != ok)\n"));
		close();
		return;
	}

	if(!json["common"].isObject()){
		ACE_DEBUG((LM_ERROR,"if(!json[common].isObject()){\n"));
		close();
		return;
	}

	info.pluginId_ = json["common"]["plugin_id"].asString();

	if(info.pluginId_.empty()){
		ACE_DEBUG((LM_ERROR,"if(plugin_id.empty())\n"));
		close();
		return;
	}

	info.userId_ = json["common"]["uid"].asString();

	if(info.userId_.empty()){
		ACE_DEBUG((LM_ERROR,"if(userid.empty())\n"));
		close();
		return;
	}
	
	info.pfId_ = json["common"]["channel"].asString();
	if(info.pfId_.empty()){
		close();
		return;
	}
	
	info.pfName_ = json["common"]["user_sdk"].asString();
	if(info.pfName_.empty())
	{
		ACE_DEBUG((LM_ERROR,"if(pfname.empty())\n"));	
		close();
		return;
	}

	info.timeout_ = 60 * 2;
	AnyOauth *pOauth = getOauthByUsername2(info.userId_);//NEW_MEM(AnyOauth,info);
	if (pOauth != NULL){
		pOauth->timeout_ = info.timeout_;
	}
	else {
		pOauth = NEW_MEM(AnyOauth, info);
		infos_.push_back(pOauth);
	}
	ACE_DEBUG((LM_INFO,"Login oauth %s %s %s \n",info.pluginId_.c_str(),info.userId_.c_str(),info.pfName_.c_str()));
	const char *  HTTP_HEADER = "HTTP/1.0 200 OK\r\nContent-type: text/html\r\n\r\n";	
	std::string strHeader= HTTP_HEADER;
	strHeader += "ok";
	fill((void*)strHeader.c_str(),strHeader.size());
	flush();
	close();
}

AnyOauth* LoginOauth::getOauthByUsername(std::string & username){
	for (size_t i = 0; i < infos_.size(); ++i){
		if ((infos_[i]->pluginId_ + "=" + infos_[i]->userId_) == username){
			return infos_[i];
		}
	}

	ACE_DEBUG((LM_ERROR, "Can not find oauth cache %s \n", username.c_str()));
	return NULL;
}

AnyOauth* LoginOauth::getOauthByUsername2(std::string & username){
	for (size_t i = 0; i < infos_.size(); ++i){
		if ( infos_[i]->userId_ == username){
			return infos_[i];
		}
	}

	ACE_DEBUG((LM_ERROR, "Can not find oauth cache %s \n", username.c_str()));
	return NULL;
}


void LoginOauth::removeOauthByUsername(std::string const & username){
	for(size_t i=0; i<infos_.size(); ++i){
		if((infos_[i]->pluginId_ + "=" + infos_[i]->userId_) == username){
			AnyOauth *p = infos_[i];
			DEL_MEM(p);
			infos_.erase(infos_.begin() + i--);
		}
	}
}

void LoginOauth::timeOutOauths(float dt){
	if(Env::get<int32>(V_UsedAnySDK)){
		for(size_t i=0; i<infos_.size(); ++i){
			infos_[i]->timeout_ -= dt;
			if(infos_[i]->timeout_ <= 0.F){
				AnyOauth *p = infos_[i];
				DEL_MEM(p);
				infos_.erase(infos_.begin() + i--);
			}
		}
	}
}

PayNotify::PayNotify(){
}

PayNotify::~PayNotify(){
}

int
PayNotify::open(void* p){
	int r = Connection::open(p);
	ACE_INET_Addr const &remote = getRemoteAddr();
	ACE_DEBUG((LM_INFO,ACE_TEXT("One client conneted at address(%s:%d)\n"),remote.get_host_addr(),remote.get_port_number()));
	return r;
}

int 
PayNotify::handleReceived(void* data, size_t size)
{
	ACE_Message_Block mb((char*)data,size);
	mb.wr_ptr(size);
	HTTP_Request req;
	req.parse_request(mb);
	proc(req.data());
	return size;
}

int
PayNotify::handle_close(ACE_HANDLE handle, ACE_Reactor_Mask close_mask)
{
	if(ACE_OS::last_error() == EWOULDBLOCK)
	{
		ACE_DEBUG((LM_INFO,ACE_TEXT("ACE_OS::last_error() == EWOULDBLOCK , Status = %d , Peer = %s\n"),status_,remoteAddr_.get_host_addr()));
		return 0;
	}
	ACE_DEBUG((LM_INFO,ACE_TEXT("Client close %s:%d Total Incoming=%Q,Outgoing-%Q\n"),getRemoteAddr().get_host_addr(),getRemoteAddr().get_port_number(),getTotalReadBytes(),getTotalWriteBytes()));
	status_ = Connection::RemoteClosed;

	delete this;
	return 0;
}

void 
PayNotify::proc( const char* post )
{
	Json::Value json;
	Json::Reader reader;
	if(reader.parse(post,json)){
	}

	if(!json.isObject()){
		close();
		return;
	}
	std::string orderId;		//订单号
	std::string orderType;		//支付渠道编号
	std::string payTime;		//支付时间
	int32 payStatus;			//支付状态
	float amount;				//支付金额
	int32 productCount;
	int32 productId = 0;
	int32 gameUserId;
	std::string userid;
	std::string pluginid;

	if(json["order_id"].isArray()){
		Json::Value tmp = json["order_id"];
		orderId = tmp[(Json::Value::UInt)0].asString();
	}

	if(json["order_type"].isArray()){
		Json::Value tmp = json["order_type"];
		orderType = tmp[(Json::Value::UInt)0].asString();
	}

	if(json["pay_time"].isArray()){
		Json::Value tmp = json["pay_time"];
		payTime =  tmp[(Json::Value::UInt)0].asString();
	}

	if(json["product_count"].isArray()){
		Json::Value tmp = json["product_count"];
		productCount = ACE_OS::atoi(tmp[(Json::Value::UInt)0].asString().c_str());
	}

	if(json["product_id"].isArray()){
		Json::Value tmp = json["product_id"];
		productId = ACE_OS::atoi(tmp[(Json::Value::UInt)0].asString().c_str());
		
		if(0 == productId){
			ACE_DEBUG((LM_INFO,"RMB by product code = %s\n",tmp[(Json::Value::UInt)0].asString().c_str()));
			const Shop::Record *record = Shop::getRecordByName(tmp[(Json::Value::UInt)0].asString());
			if(record){
				productId = record->id_;
			}
		}
	}

	if(json["game_user_id"].isArray()){
		Json::Value tmp = json["game_user_id"];
		gameUserId = ACE_OS::atoi(tmp[(Json::Value::UInt)0].asString().c_str());
	}

	if(json["pay_status"].isArray()){
		Json::Value tmp = json["pay_status"];
		payStatus = ACE_OS::atoi(tmp[(Json::Value::UInt)0].asString().c_str());
	}

	if(json["user_id"].isArray()){
		Json::Value tmp = json["user_id"];
		userid = tmp[(Json::Value::UInt)0].asString().c_str();
	}

	if(json["plugin_id"].isArray()){
		Json::Value tmp = json["plugin_id"];
		pluginid = tmp[(Json::Value::UInt)0].asString().c_str();
	}
	
	if(json["amount"].isArray()){
		Json::Value tmp = json["amount"];
		amount = ACE_OS::atof(tmp[(Json::Value::UInt)0].asString().c_str());
	}

	if(payStatus != 1){
		close();
		return;
	}
	
	Player* player = Player::getPlayerByInstId(gameUserId);

	if(player != NULL){
		if(!player->orderFromSDK(productId,productCount,orderId,payTime,amount)){
			ACE_DEBUG((LM_ERROR,"Rechage error player %d ==>? shop %d \n",gameUserId,productId));
		}	
	}else{
		SGE_OrderInfo info;
		info.productId_ = productId;
		info.productCount_ = productCount;
		info.amount_ = amount;
		info.orderId_ = orderId;
		info.payTime_ = payTime;
		
		std::string accName = pluginid + "=" + userid;

		Account* pAccount = Account::getAccountByName(accName);
		if(pAccount){
			SGE_DBPlayerData * pPlayerData = pAccount->findDBPlayerById(gameUserId);
			if(pPlayerData){
				pPlayerData->orders_.push_back(info);
			}else {
				ACE_DEBUG((LM_ERROR,"Charge cannot find dbplayer %s, %d\n",accName.c_str(),gameUserId));
			}
		}else {
			ACE_DEBUG((LM_INFO,"Charge account is offline %s\n",accName.c_str()));
		}
		
		DBHandler::instance()->insertLoseCharge(gameUserId,info);
	}

	const char *  HTTP_HEADER = "HTTP/1.0 200 OK\r\nContent-type: text/html\r\n\r\n";	
	std::string strHeader= HTTP_HEADER;
	strHeader += "ok";
	fill((void*)strHeader.c_str(),strHeader.size());
	flush();
	close();

}