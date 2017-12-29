#include "product.h"

RefreshSql::RefreshSql()
:total_rows_(0)
,trans_rows_(0)
{
	fparams_.resize(RefreshSql::MAX);
	tparams_.resize(RefreshSql::MAX);
	catchMaxid_.resize(RefreshSql::MAXID);
	catchPlayerName_.clear();
	catchPlayerinstid_.clear();
}

void RefreshSql::init()
{
	//fparams_[HOST]		= Env::get<const char*>(V_FromSqlHost);
	//fparams_[USER]		= Env::get<const char*>(V_FromSqlUser);
	//fparams_[PWD]		= Env::get<const char*>(V_FromSqlPwd);
	//fparams_[SCHEMA]	= Env::get<const char*>(V_FromSqlSchema);
	//tparams_[HOST]		= Env::get<const char*>(V_ToSqlHost);
	//tparams_[USER]		= Env::get<const char*>(V_ToSqlUser);
	//tparams_[PWD]		= Env::get<const char*>(V_ToSqlPwd);
	//tparams_[SCHEMA]	= Env::get<const char*>(V_ToSqlSchema);

	SRV_ASSERT(true==open_mysql());
}

void RefreshSql::fini()
{
	SRV_ASSERT(true==close_mysql());
}

void RefreshSql::run(){
DB_EXEC_GUARD
	{
		char * select_code = "SELECT * FROM Player";

		std::auto_ptr< sql::Statement > stmt(tomysql_->createStatement());
		std::auto_ptr< sql::ResultSet > res(stmt->executeQuery(select_code));

		while(res->next())
		{
			std::string playername = res->getString("PlayerName");
			U32 instId			   = res->getInt("PlayerGuid");
			catchPlayerName_.push_back(playername);
			catchPlayerinstid_.push_back(instId);
		}
	}
	
	{
		char * select_code = "SELECT * FROM Player WHERE PlayerGuid=(SELECT MAX(PlayerGuid) FROM Player);";

		std::auto_ptr< sql::Statement > stmt(tomysql_->createStatement());
		std::auto_ptr< sql::ResultSet > res(stmt->executeQuery(select_code));
		if(res->next()){
			catchMaxid_[PLAYERMAXID] = res->getInt("PlayerGuid");
		}
	}

	{
		char * select_code = "SELECT * FROM Baby WHERE BabyGuid=(SELECT MAX(BabyGuid) FROM Baby);";

		std::auto_ptr< sql::Statement > stmt(tomysql_->createStatement());
		std::auto_ptr< sql::ResultSet > res(stmt->executeQuery(select_code));
		if(res->next()){
			catchMaxid_[BABYMAXID] = res->getInt("BabyGuid");
		}
	}

	{
		char * select_code = "SELECT * FROM Employee WHERE EmployeeGuid=(SELECT MAX(EmployeeGuid) FROM Employee);";

		std::auto_ptr< sql::Statement > stmt(tomysql_->createStatement());
		std::auto_ptr< sql::ResultSet > res(stmt->executeQuery(select_code));
		if(res->next()){
			catchMaxid_[EMPLOYEEMAXID] = res->getInt("EmployeeGuid");
		}
	}

	{
		char * select_code = "SELECT * FROM Guild WHERE GuildId=(SELECT MAX(GuildId) FROM Guild);";

		std::auto_ptr< sql::Statement > stmt(tomysql_->createStatement());
		std::auto_ptr< sql::ResultSet > res(stmt->executeQuery(select_code));
		if(res->next()){
			catchMaxid_[GUILDMAXID] = res->getInt("GuildId");
		}
	}

	{
		char * select_code = "SELECT * FROM Mail WHERE MailGuid=(SELECT MAX(MailGuid) FROM Mail);";

		std::auto_ptr< sql::Statement > stmt(tomysql_->createStatement());
		std::auto_ptr< sql::ResultSet > res(stmt->executeQuery(select_code));
		if(res->next()){
			catchMaxid_[MAILMAXID] = res->getInt("MailGuid");
		}
	}
	mergeAccount();
	playerRun();

DB_EXEC_UNGUARD_RETURN
}

void RefreshSql::playerRun()
{
	ACE_DEBUG((LM_INFO,ACE_TEXT("Begin Merge Player...\n")));
DB_EXEC_GUARD
	static const char * select_code = "SELECT * FROM Player";
	
	std::auto_ptr< sql::Statement > stmt(formmysql_->createStatement());
	std::auto_ptr< sql::ResultSet > res(stmt->executeQuery(select_code));
	
	int sec = ACE_OS::gettimeofday().sec();

	total_rows_ = res->rowsCount();
	while(res->next())
	{
		U32 instId					= res->getInt("PlayerGuid");
		std::string username		= res->getString("UserName").c_str();
		std::string playername		= res->getString("PlayerName");
		U32 level					= res->getInt("PlayerLevel");
		U32 pf						= res->getInt("PlayerProfession");
		U32 grade					= res->getInt("PlayerGrade");
		U32 seal					= res->getInt("Seal");
		U32 freeze					= res->getInt("Freeze");
		std::stringstream ssname;
		ssname << playername;
		if(checkSamePlayerName(playername)){
			ssname << Env::get<int>(V_FromGameServId);
			ACE_DEBUG((LM_INFO,ACE_TEXT("Same PlayerName NEW Name[%s]...\n"),ssname.str().c_str()));
		}

		if(checkSamePlayerInstID(instId)){
			instId+=catchMaxid_[PLAYERMAXID];
			ACE_DEBUG((LM_INFO,ACE_TEXT("Same PlayerInstId NEWID[%d]...\n"),instId));
		}

		sql::SQLString pCacheBlob	= res->getString("BinData");
		ProtocolMemReader mr(pCacheBlob->c_str(),pCacheBlob->length());
		SGE_DBPlayerData inst;
		inst.deserialize(&mr);

		inst.instId_	= instId;
		inst.instName_	= ssname.str();

		{
			std::stringstream ssbaby;
			ssbaby << "SELECT * FROM Baby WHERE OwnerName = \"" << playername << "\"";
			std::auto_ptr< sql::Statement > stmt1(formmysql_->createStatement());
			std::auto_ptr< sql::ResultSet > res1(stmt1->executeQuery(ssbaby.str().c_str()));
			inst.babies_.clear();
			while(res1->next())
			{
				sql::SQLString pCacheBlob1= res1->getString("BinData");
				ProtocolMemReader mrbaby(pCacheBlob1->c_str(),pCacheBlob1->length());
				COM_BabyInst instbaby;
				instbaby.deserialize(&mrbaby);
				
				if(checkSameBabyInstID(instbaby.instId_))
					instbaby.instId_ += catchMaxid_[BABYMAXID];
				instbaby.ownerName_ = inst.instName_;
				inst.babies_.push_back(instbaby);

				insertBaby(instbaby);
			}
		}

		{
			std::stringstream ssemployee;
			ssemployee << "SELECT * FROM Employee WHERE OwnerName = \"" << playername << "\"";
			std::auto_ptr< sql::Statement > stmt1(formmysql_->createStatement());
			std::auto_ptr< sql::ResultSet > res1(stmt1->executeQuery(ssemployee.str().c_str()));
			inst.employees_.clear();
			while(res1->next())
			{
				sql::SQLString pCacheBlob1= res1->getString("BinData");
				ProtocolMemReader mremployee(pCacheBlob1->c_str(),pCacheBlob1->length());
				COM_EmployeeInst instemployee;
				instemployee.deserialize(&mremployee);

				if(checkSameEmployeeInstID(instId))
					instemployee.instId_ += catchMaxid_[EMPLOYEEMAXID];
				instemployee.ownerName_ = inst.instName_;
				inst.employees_.push_back(instemployee);

				insertEmployee(instemployee);
			}
		}

		insertPlayer(username,inst);
	}
	ACE_DEBUG((LM_INFO,ACE_TEXT("End Merge Player...\n")));
	mergeGuild();
	mergeGuildMember();
	mergeMail();
	mergeMall();
	//mergeMallSell();
DB_EXEC_UNGUARD_RETURN
};

void RefreshSql::process(std::string& tableName,int current,int total, int sec )
{
	
	float ratio = (float)current / float(total);	///<±ÈÀý
	
	enum{BAR_LENGTH=50};
	
	char buffer[BAR_LENGTH+1] = {'\0'};
	int i = 0;
	for(int length=BAR_LENGTH*ratio; i<length;++i)
	{
		buffer[i] = '=';
	}	
	
	for(;i<BAR_LENGTH;++i)
	{
		buffer[i] = ' ';
	}

	buffer[BAR_LENGTH] = '\0';

	ACE_OS::printf("\r[%s]%d%%[%s] time=%d",tableName.c_str(),(int)(ratio*100),(char*)buffer,sec);
	
	ACE_OS::fflush(stdout);
	
}

bool RefreshSql::open_mysql()
{
DB_EXEC_GUARD
	sql::Driver *pfdriver = sql::mysql::get_driver_instance();
	if(NULL==pfdriver)
	{
		ACE_DEBUG((LM_INFO,ACE_TEXT("pfdriver==NULL\n")));
		return false;
	}
	formmysql_ =pfdriver->connect(fparams_[HOST].c_str(),fparams_[USER].c_str(),fparams_[PWD].c_str());
	if(NULL==formmysql_)
	{
		ACE_DEBUG((LM_INFO,ACE_TEXT("oldmysql_==NULL\n")));
		return false;
	}
	formmysql_->setSchema(fparams_[SCHEMA].c_str());
	pfdriver=NULL;

	sql::Driver *ptdriver = sql::mysql::get_driver_instance();
	if(NULL==ptdriver)
	{
		ACE_DEBUG((LM_INFO,ACE_TEXT("pfdriver==NULL\n")));
		return false;
	}
	tomysql_ =ptdriver->connect(tparams_[HOST].c_str(),tparams_[USER].c_str(),tparams_[PWD].c_str());
	if(NULL==tomysql_)
	{
		ACE_DEBUG((LM_INFO,ACE_TEXT("newmysql_==NULL\n")));
		return false;
	}
	tomysql_->setSchema(tparams_[SCHEMA].c_str());
	ptdriver=NULL;

	return true;
DB_EXEC_UNGUARD_RETURN
}

bool RefreshSql::close_mysql()
{
DB_EXEC_GUARD
		if(NULL == formmysql_)
		{
			return true;
		}
		formmysql_->close();
		formmysql_=NULL;

		if(NULL == tomysql_)
		{
			return true;
		}
		tomysql_->close();
		tomysql_=NULL;

		return true;
DB_EXEC_UNGUARD_RETURN
}

bool RefreshSql::checkSamePlayerName(std::string& name){
	for (size_t i=0; i<catchPlayerName_.size();++i)
	{
		if(catchPlayerName_[i] == name)
			return true;
	}
	return false;
}

bool RefreshSql::checkSamePlayerInstID(U32 instid){
	for (size_t i=0;i<catchPlayerinstid_.size();++i)
	{
		if(catchPlayerinstid_[i] == instid)
			return true;
	}
	return false;
}

bool RefreshSql::checkSameBabyInstID(U32 instid){
	DB_EXEC_GUARD
	std::stringstream sstream;
	sstream << "SELECT * FROM Baby WHERE BabyGuid = \"" << instid << "\"";

	std::auto_ptr< sql::Statement > stmt(tomysql_->createStatement());
	std::auto_ptr< sql::ResultSet > res(stmt->executeQuery(sstream.str().c_str()));
	if(res->next()){
		return false;
	}

	return true;
	DB_EXEC_UNGUARD_RETURN
}

bool RefreshSql::checkSameEmployeeInstID(U32 instid){
	DB_EXEC_GUARD
	std::stringstream sstream;
	sstream << "SELECT * FROM Employee WHERE EmployeeGuid = \"" << instid << "\"";

	std::auto_ptr< sql::Statement > stmt(tomysql_->createStatement());
	std::auto_ptr< sql::ResultSet > res(stmt->executeQuery(sstream.str().c_str()));
	if(res->next()){
		return false;
	}

	return true;
	DB_EXEC_UNGUARD_RETURN
}

void RefreshSql::insertBaby(COM_BabyInst &inst){
	char * pCode =
		"INSERT INTO Baby(BabyGuid,BabyName,BinData,OwnerName,BabyGrade)"
		"VALUES(?,?,?,?,?);";
	std::auto_ptr< sql::PreparedStatement > res(tomysql_->prepareStatement(pCode));

	char buffer[20480] = {'\0'};
	ProtocolMemWriter mw(buffer,sizeof(buffer));
	inst.serialize(&mw);
	DB_EXEC_GUARD

	res->setInt(1,inst.instId_);
	res->setString(2,inst.instName_.c_str());
	sql::SQLString binString(buffer,mw.length());
	res->setString(3,binString);
	res->setString(4,inst.ownerName_.c_str());
	res->setInt(5,inst.properties_[PT_FightingForce]);
	res->executeUpdate();
	ACE_DEBUG((LM_INFO,ACE_TEXT("INSERTBABY ===>OWNER[%s]===>NAME[%s]===>INSTID[%d]\n"),inst.ownerName_.c_str(),inst.instName_.c_str(),inst.instId_));
	DB_EXEC_UNGUARD_RETURN
}

void RefreshSql::insertEmployee(COM_EmployeeInst &inst){
	char * pCode =
		"INSERT INTO Employee(EmployeeGuid,EmployeeName,BinData,OwnerName,EmployeeGrade)"
		"VALUES(?,?,?,?,?);";
	std::auto_ptr< sql::PreparedStatement > res(tomysql_->prepareStatement(pCode));
	char buffer[20480] = {'\0'};
	ProtocolMemWriter mw(buffer,sizeof(buffer));
	inst.serialize(&mw);
	DB_EXEC_GUARD

	res->setInt(1,inst.instId_);
	res->setString(2,inst.instName_);
	sql::SQLString binString(buffer,mw.length());
	res->setString(3,binString);
	res->setString(4,inst.ownerName_);
	res->setInt(5,inst.properties_[PT_FightingForce]);
	res->executeUpdate();

	ACE_DEBUG((LM_INFO,ACE_TEXT("INSERTEMPLOYEE ===>OWNER[%s]===>NAME[%s]===>INSTID[%d]\n"),inst.ownerName_.c_str(),inst.instName_.c_str(),inst.instId_));
	DB_EXEC_UNGUARD_RETURN
}

void RefreshSql::insertPlayer(std::string &username,SGE_DBPlayerData &data){
	static const char * pCode =
		"INSERT INTO Player(UserName,PlayerGuid,BinData,PlayerName,PlayerLevel,PlayerProfession,PlayerGrade,Seal,Freeze,InDoorId)"
		"VALUES(?,?,?,?,?,?,?,?,?,?);";

	std::auto_ptr< sql::PreparedStatement > res(tomysql_->prepareStatement(pCode));

	char buffer[20480] = {'\0'};
	ProtocolMemWriter mw(buffer,sizeof(buffer));
	data.serialize(&mw);
	DB_EXEC_GUARD

	res->setString(1,username.c_str());
	res->setInt(2,data.instId_);
	sql::SQLString binString(buffer,mw.length());
	res->setString(3,binString);
	res->setString(4,data.instName_);
	res->setInt(5,data.properties_[PT_Level]);
	res->setInt(6,data.properties_[PT_Profession]);
	res->setInt(7,data.properties_[PT_FightingForce]);
	res->setInt(8,data.seal_);
	res->setInt(9,data.freeze_);
	res->setInt(10,Env::get<int>(V_FromGameServId));
	res->executeUpdate();

	ACE_DEBUG((LM_INFO,ACE_TEXT("INSERTPLAYER ==TOMAX[%d]=>NAME[%s]===>INSTID[%d]\n"),catchMaxid_[PLAYERMAXID],data.instName_.c_str(),data.instId_));
	DB_EXEC_UNGUARD_RETURN		
}


//////////////////////////////////////////////////////////////////////////

void RefreshSql::mergeGuild(){
	ACE_DEBUG((LM_INFO,ACE_TEXT("Begin Merge Guild...\n")));
	DB_EXEC_GUARD
	std::string str="SELECT * FROM Guild";

	std::auto_ptr< sql::Statement > stmt(formmysql_->createStatement());
	std::auto_ptr< sql::ResultSet > res(stmt->executeQuery(str.c_str()));
	while (res->next()) 
	{
		COM_Guild	guild;
		guild.guildId_   = res->getInt("GuildId");
		guild.guildLevel_= res->getInt("GuildLevel");
		guild.guildName_ = res->getString("GuildName");
		guild.guildContribution_ = res->getInt("Contribution");
		guild.fundz_ = res->getInt("Fundz");
		guild.master_    = res->getInt64("Master");
		guild.masterName_=res->getString("MasterName");
		guild.notice_	 = res->getString("Notice");
		guild.createTime_=res->getInt("Credit");

		sql::SQLString tmpBlob = res->getString("RequestList");
		const  char * pBlob= tmpBlob.c_str();
		ProtocolMemReader protocol1((S8*)pBlob , tmpBlob.length());
		protocol1.readVector(guild.requestList_);

		tmpBlob = res->getString("Buildings" );
		pBlob= tmpBlob.c_str();
		ProtocolMemReader protocol2((S8*)pBlob , tmpBlob.length());
		protocol2.readVector(guild.buildings_);

		tmpBlob = res->getString("Progenitus");
		pBlob= tmpBlob.c_str();
		ProtocolMemReader protocol3((S8*)pBlob, tmpBlob.length());
		protocol3.readVector(guild.progenitus_);

		tmpBlob = res->getString("ProgenitusPos");
		pBlob= tmpBlob.c_str();
		ProtocolMemReader protocol4((S8*)pBlob , tmpBlob.length());
		protocol4.readVector(guild.progenitusPositions_);

		{
			std::stringstream sstream;
			sstream << guild.guildName_;
			if(checkSameGuildName(guild.guildName_))
				sstream << Env::get<int>(V_FromGameServId);
			guild.guildName_ = sstream.str();
		}

		{
			std::stringstream sstream;
			sstream << guild.masterName_;
			if(checkSamePlayerName(guild.masterName_))
				sstream << Env::get<int>(V_FromGameServId);
			guild.masterName_ = sstream.str();
		}
		if(checkSameGuildID(guild.guildId_)){
			U32 old = guild.guildId_;
			std::stringstream ss;
			ss << old << Env::get<int>(V_FromGameServId);
			guild.guildId_ = atoi(ss.str().c_str());
			updateGuildMember(old,guild.guildId_);
		}
		
		insertGuild(guild);
	}
	ACE_DEBUG((LM_INFO,ACE_TEXT("End Merge Guild...\n")));
	DB_EXEC_UNGUARD_RETURN
}

bool RefreshSql::checkSameGuildName(std::string &name){
	DB_EXEC_GUARD
		std::stringstream sstream;
	sstream << "SELECT * FROM Guild WHERE GuildName = \"" << name << "\"";

	std::auto_ptr< sql::Statement > stmt(tomysql_->createStatement());
	std::auto_ptr< sql::ResultSet > res(stmt->executeQuery(sstream.str().c_str()));
	if(res->next()){
		return true;
	}

	return false;
	DB_EXEC_UNGUARD_RETURN
}

bool RefreshSql::checkSameGuildID(U32 instid){
	DB_EXEC_GUARD
		std::stringstream sstream;
	sstream << "SELECT * FROM Guild WHERE GuildId = \"" << instid << "\"";

	std::auto_ptr< sql::Statement > stmt(tomysql_->createStatement());
	std::auto_ptr< sql::ResultSet > res(stmt->executeQuery(sstream.str().c_str()));
	if(res->next()){
		return true;
	}

	return false;
	DB_EXEC_UNGUARD_RETURN
}

void RefreshSql::insertGuild(COM_Guild &guild){
	DB_EXEC_GUARD
		const char* pGuildCode = "INSERT INTO Guild(GuildId,GuildName,GuildLevel,Contribution,Fundz,Credit,Master,MasterName,Notice,RequestList,Buildings,Progenitus,ProgenitusPos)"
		"values(?,?,?,?,?,?,?,?,?,?,?,?,?) ;";

	std::auto_ptr< sql::PreparedStatement > prep_stmt(tomysql_->prepareStatement(pGuildCode));
	prep_stmt->setInt( 1 , int(guild.guildId_) );
	prep_stmt->setString( 2 , guild.guildName_.c_str());
	prep_stmt->setInt( 3 , guild.guildLevel_);
	prep_stmt->setInt( 4 , guild.guildContribution_);
	prep_stmt->setInt( 5 , guild.fundz_);
	prep_stmt->setInt( 6 , guild.createTime_);
	prep_stmt->setInt64( 7 , guild.master_);
	prep_stmt->setString( 8 , guild.masterName_.c_str());
	prep_stmt->setString( 9 , guild.notice_.c_str());

	char buffer1[sizeof(COM_GuildRequestData)*32] = {0};
	ProtocolMemWriter protocol1((S8*)buffer1 , sizeof(buffer1));
	bool checker = protocol1.writeVector(guild.requestList_);
	SRV_ASSERT(checker);
	sql::SQLString tmpStr1(buffer1,protocol1.length());
	prep_stmt->setString( 10 ,tmpStr1 );

	char buffer2[sizeof(COM_GuildBuilding)*32] = {0};
	ProtocolMemWriter protocol2((S8*)buffer2 , sizeof(buffer2));
	checker = protocol2.writeVector(guild.buildings_);
	SRV_ASSERT(checker);
	sql::SQLString tmpStr2(buffer2,protocol2.length());
	prep_stmt->setString( 11 , tmpStr2);

	char buffer3[sizeof(COM_GuildProgen)*32] = {0};
	ProtocolMemWriter protocol3((S8*)buffer3 , sizeof(buffer3));
	checker = protocol3.writeVector(guild.progenitus_);
	SRV_ASSERT(checker);
	sql::SQLString tmpStr3(buffer3,protocol3.length());
	prep_stmt->setString( 12 , tmpStr3);

	char buffer4[sizeof(S32)*32] = {0};
	ProtocolMemWriter protocol4((S8*)buffer4 , sizeof(buffer4));
	checker = protocol4.writeVector(guild.progenitusPositions_);
	SRV_ASSERT(checker);
	sql::SQLString tmpStr4(buffer4,protocol4.length());
	prep_stmt->setString( 13 , tmpStr4);

	prep_stmt->execute();
	ACE_DEBUG((LM_INFO,ACE_TEXT("INSERTGUILD ===>GUILDNAME[%s]===>GUILDID[%d]====>masterName[%s]\n"),guild.guildName_.c_str(),guild.guildId_,guild.masterName_.c_str()));
	DB_EXEC_UNGUARD_RETURN
}

void RefreshSql::mergeGuildMember(){
ACE_DEBUG((LM_INFO,ACE_TEXT("Begin Merge GuildMember...\n")));
DB_EXEC_GUARD
	std::string str="SELECT * FROM GuildMember ";

	std::auto_ptr< sql::Statement > stmt(formmysql_->createStatement());
	std::auto_ptr< sql::ResultSet > res(stmt->executeQuery(str.c_str()));
	while (res->next()) 
	{
		COM_GuildMember	guildMember;
		guildMember.guildId_  = res->getInt("GuildId");
		guildMember.roleId_   = res->getInt64("RoleId");
		guildMember.job_ = GuildJob(res->getInt("Job"));
		guildMember.contribution_ = res->getInt("Contribution");
		guildMember.level_	  = res->getInt("Rolelevel");
		guildMember.offlineTime_ = res->getInt64("OfflineTime");
		guildMember.roleName_ = res->getString("RoleName");
		guildMember.joinTime_ = SexType(res->getInt("JoinTime"));
		guildMember.profType_ = res->getInt("Proftype");
		guildMember.profLevel_ = res->getInt("Proflevel");

		{
			std::stringstream sstream;
			sstream << guildMember.roleName_;
			if(checkSamePlayerName(guildMember.roleName_))
				sstream << Env::get<int>(V_FromGameServId);
			guildMember.roleName_ = sstream.str();
		}

		if(checkSamePlayerInstID(guildMember.roleId_)){
			guildMember.roleId_+=catchMaxid_[PLAYERMAXID];
		}

		insertGuildMember(guildMember);
	}
	ACE_DEBUG((LM_INFO,ACE_TEXT("End Merge GuildMember...\n")));
DB_EXEC_UNGUARD_RETURN
}

void RefreshSql::insertGuildMember(COM_GuildMember &member){
DB_EXEC_GUARD
	const char* pMemberCode = "INSERT INTO GuildMember(GuildId,RoleId,Job,Contribution,Rolelevel,JoinTime,Proftype,Proflevel,OfflineTime,RoleName)"
		"values(?,?,?,?,?,?,?,?,?,?) ;";

	std::auto_ptr< sql::PreparedStatement > prep_stmt(tomysql_->prepareStatement(pMemberCode));
	prep_stmt->setInt( 1 , int(member.guildId_) );
	prep_stmt->setInt64( 2 , member.roleId_);
	prep_stmt->setInt( 3 , member.job_);
	prep_stmt->setInt( 4 , member.contribution_);
	prep_stmt->setInt( 5 , member.level_);
	prep_stmt->setInt( 6 , int(member.joinTime_));
	prep_stmt->setInt( 7 , int(member.profType_));
	prep_stmt->setInt( 8 , int(member.profLevel_));
	prep_stmt->setInt( 9 , int(member.offlineTime_));
	prep_stmt->setString( 10 , member.roleName_.c_str());
	prep_stmt->execute();
	ACE_DEBUG((LM_INFO,ACE_TEXT("INSERTGUILDMEMBER ===>MEMBERNAME[%s]====>MEMBERID[%d]\n"),member.roleName_.c_str(),member.roleId_));
DB_EXEC_UNGUARD_RETURN
}

void RefreshSql::updateGuildMember(U32 oldGuildid,U32 newGuildid){
DB_EXEC_GUARD
	std::stringstream strCode;
	strCode<<"UPDATE GuildMember SET GuildId="<<newGuildid<<" WHERE GuildId="<<oldGuildid<<" ";

	std::auto_ptr< sql::PreparedStatement > prep_stmt(formmysql_->prepareStatement(strCode.str().c_str()));
	prep_stmt->execute();
	ACE_DEBUG((LM_INFO,ACE_TEXT("updateGuildMember ===>newGuildid[%d]====>oldGuildid[%d]\n"),newGuildid,oldGuildid));
DB_EXEC_UNGUARD_RETURN
}

//////////////////////////////////////////////////////////////////////////

bool RefreshSql::checkSameMailID(U32 mailid){
DB_EXEC_GUARD
	std::stringstream sstream;
	sstream << "SELECT * FROM Mail WHERE MailGuid = \"" << mailid << "\"";

	std::auto_ptr< sql::Statement > stmt(tomysql_->createStatement());
	std::auto_ptr< sql::ResultSet > res(stmt->executeQuery(sstream.str().c_str()));
	if(res->next()){
		return true;
	}

	return false;
DB_EXEC_UNGUARD_RETURN
}

void RefreshSql::mergeMail(){
	ACE_DEBUG((LM_INFO,ACE_TEXT("Begin Merger Mail...\n")));
	std::stringstream selStr;
	selStr<<"SELECT * FROM Mail";

	std::auto_ptr< sql::Statement > stmt(formmysql_->createStatement());
	std::auto_ptr< sql::ResultSet > res(stmt->executeQuery(selStr.str().c_str()));
DB_EXEC_GUARD
	while(res->next())
	{
		sql::SQLString pCacheBlob= res->getString("BinData");
		ProtocolMemReader mr(pCacheBlob->c_str(),pCacheBlob->length());
		COM_Mail mail;
		mail.deserialize(&mr);
		mail.mailId_ =  res->getInt("MailGuid");
		
		if(checkSameMailID(mail.mailId_))
			mail.mailId_ += catchMaxid_[MAILMAXID];
		{
			std::stringstream sstream;
			sstream << mail.recvPlayerName_;
			if(checkSamePlayerName(mail.recvPlayerName_))
				sstream << Env::get<int>(V_FromGameServId);
			mail.recvPlayerName_ = sstream.str();
		}

		insertMail(mail);
	}
	ACE_DEBUG((LM_INFO,ACE_TEXT("End Merge Mail...\n")));
DB_EXEC_UNGUARD_RETURN
}

void RefreshSql::insertMail(COM_Mail &mail){
	static const char* pCode = "INSERT INTO Mail( RecvName , SendTime , ItemNum , BinData ) VALUES(?,?,?,?);";

	char buffer[204800] = {'\0'};

DB_EXEC_GUARD
	std::auto_ptr< sql::PreparedStatement > prep_stmt(tomysql_->prepareStatement(pCode));
	ProtocolMemWriter mw(buffer,sizeof(buffer));
	mail.serialize(&mw);
	prep_stmt->setString(1,mail.recvPlayerName_.c_str());
	prep_stmt->setInt64(2,mail.timestamp_);
	prep_stmt->setInt(3,mail.items_.size());
	sql::SQLString binString(buffer,mw.length());
	prep_stmt->setString(4,binString);
	prep_stmt->executeUpdate();
	ACE_DEBUG((LM_INFO,ACE_TEXT("INSERTMAIL ===>RECVNAME[%s]\n"),mail.recvPlayerName_.c_str()));
DB_EXEC_UNGUARD_RETURN
}

//////////////////////////////////////////////////////////////////////////

void RefreshSql::mergeMallSell(){
	std::stringstream sstream;
	sstream << "SELECT * FROM MallSelledTable";

DB_EXEC_GUARD
	std::auto_ptr< sql::Statement > stmt(formmysql_->createStatement());
	std::auto_ptr< sql::ResultSet > res(stmt->executeQuery(sstream.str().c_str()));
	while (res->next()) 
	{
		COM_SelledItem item;
		item.guid_ = res->getInt("Guid");
		item.sellPlayerId_ = res->getInt("PlayerId");
		item.price_ = res->getInt("SellPrice");
		item.itemStack_ = res->getInt("Stack");
		item.itemId_ = res->getInt("ItemId");
		item.babyId_ = res->getInt("BabyId");
		item.selledTime_ = res->getInt("SelledTime");
		item.tax_ = res->getInt("Tax");

		if(checkSamePlayerInstID(item.sellPlayerId_)){
			item.sellPlayerId_+=catchMaxid_[PLAYERMAXID];
		}
	}

DB_EXEC_UNGUARD_RETURN
}

void RefreshSql::insertMallSell(COM_SelledItem &item){
	static const char* insertcode = "INSERT INTO MallSelledTable(PlayerId,SellPrice,Stack,BabyId,ItemId,SelledTime,Tax) VALUES(?,?,?,?,?,?,?)";
DB_EXEC_GUARD
	std::auto_ptr< sql::PreparedStatement > prep_stmt(tomysql_->prepareStatement(insertcode));
	prep_stmt->setInt( 1,item.sellPlayerId_);
	prep_stmt->setInt( 2,item.price_);
	prep_stmt->setInt( 3,item.itemStack_);
	prep_stmt->setInt( 4,item.babyId_);
	prep_stmt->setInt( 5,item.itemId_);
	prep_stmt->setInt( 6,item.selledTime_);
	prep_stmt->setInt( 7,item.tax_);
	prep_stmt->execute();
DB_EXEC_UNGUARD_RETURN
}

void RefreshSql::mergeMall(){
	ACE_DEBUG((LM_INFO,ACE_TEXT("Begin Merge Mall...\n")));
	std::stringstream sstream;
	sstream << "SELECT * FROM MallTable";
DB_EXEC_GUARD
	std::auto_ptr< sql::Statement > stmt(formmysql_->createStatement());
	std::auto_ptr< sql::ResultSet > res(stmt->executeQuery(sstream.str().c_str()));
	while (res->next()) 
	{
		sql::SQLString pCacheBlob = res->getString("SellItem");
		ProtocolMemReader mr(pCacheBlob.c_str(),pCacheBlob.length());
		COM_SellItem item;
		item.deserialize(&mr);
		item.guid_ = res->getInt("Guid");
		if(checkSamePlayerInstID(item.sellPlayerId_)){
			item.sellPlayerId_+=catchMaxid_[PLAYERMAXID];
		}

		insertMall(item);
	}
	ACE_DEBUG((LM_INFO,ACE_TEXT("End Merge Mall...\n")));
DB_EXEC_UNGUARD_RETURN
}

void RefreshSql::insertMall(COM_SellItem &item){
	static const char* pCode = "SELECT MAX(Guid) AS GuidMax FROM MallTable ";
	static const char* insertcode = "INSERT INTO MallTable(Guid,PlayerId,Title,ItemSubType,RaceType,ItemId,BabyId,SellPrice,SellItem) VALUES(?,?,?,?,?,?,?,?,?)";
DB_EXEC_GUARD

	std::auto_ptr< sql::Statement > stmt(tomysql_->createStatement());
	std::auto_ptr< sql::ResultSet > res(stmt->executeQuery(pCode));
	if(res->next()){
		item.guid_ =res->getInt("GuidMax") + 1;
	}
	std::auto_ptr< sql::PreparedStatement > prep_stmt(tomysql_->prepareStatement(insertcode));
	prep_stmt->setInt( 1 , int(item.guid_) );
	prep_stmt->setInt( 2 , int(item.sellPlayerId_) );
	prep_stmt->setString( 3 , item.title_.c_str());
	prep_stmt->setInt( 4 , item.ist_);
	prep_stmt->setInt64( 5 , item.rt_);
	prep_stmt->setInt(6,item.item_.itemId_);
	if(item.baby_.properties_.empty())
		prep_stmt->setInt(7,0);
	else
		prep_stmt->setInt(7,item.baby_.properties_[PT_TableId]);
	prep_stmt->setInt( 8 , item.sellPrice);
	char buffer[20480] = {'\0'};
	ProtocolMemWriter mw(buffer,sizeof(buffer));
	item.serialize(&mw);
	sql::SQLString str(buffer,mw.length());
	prep_stmt->setString(9, str);
	prep_stmt->execute();
	ACE_DEBUG((LM_INFO,ACE_TEXT("INSERTMALL ===>SELLPLAYER[%d]\n"),item.sellPlayerId_));
DB_EXEC_UNGUARD_RETURN
}

void RefreshSql::mergeAccount(){
	ACE_DEBUG((LM_DEBUG ,"mergeAccount go\n"));
	std::stringstream sstream;
	sstream << "SELECT * FROM Account";
	
DB_EXEC_GUARD

	std::auto_ptr< sql::Statement > stmt(formmysql_->createStatement());
	std::auto_ptr< sql::ResultSet > res(stmt->executeQuery(sstream.str().c_str()));
	while (res->next()) 
	{
		sql::SQLString pCacheBlob= res->getString("AccountInfo");
		ProtocolMemReader mr(pCacheBlob->c_str(),pCacheBlob->length());
		COM_AccountInfo inst;
		inst.deserialize(&mr);
		bool isSeal = !!res->getInt("Seal");
		insertAccount(inst,isSeal);
	}
	ACE_DEBUG((LM_DEBUG ,"mergeAccount End\n"));
DB_EXEC_UNGUARD_RETURN
}

void RefreshSql::insertAccount(COM_AccountInfo & info,bool isSeal){
	static const char * pCode =
		"INSERT INTO Account(UserName,Password,AccountInfo,Seal,PhoneNumber) "
		"VALUES(?,?,?,?,?);";
DB_EXEC_GUARD

	char buffer[20480] = {'\0'};
	ProtocolMemWriter mw(buffer,sizeof(buffer));
	
	std::auto_ptr< sql::PreparedStatement > prep_stmt(tomysql_->prepareStatement(pCode));
	prep_stmt->setString(1,info.username_.c_str());
	prep_stmt->setString(2,info.password_.c_str());
	prep_stmt->setInt(4,isSeal ? 1 : 0);
	sql::SQLString binString(buffer,mw.length());
	prep_stmt->setString(3,binString);
	prep_stmt->setString(5,info.phoneNumber_);
	prep_stmt->executeUpdate();
	ACE_DEBUG((LM_INFO,ACE_TEXT("INSERTACCOUNT ===>USERNAME[%s]\n"),info.username_.c_str()));
DB_EXEC_UNGUARD_RETURN		
}

void RefreshSql::checkAccount(){

}