package prpc

import (
	"bytes"
	"errors"
	"suzuki/prpc"
)

type Server2Client_errorno struct {
	e int //0
}
type Server2Client_teamerrorno struct {
	name string //0
	e    int    //1
}
type Server2Client_reconnection struct {
	recInfo COM_ReconnectInfo //0
}
type Server2Client_loginok struct {
	sessionkey string                  //0
	players    []COM_SimpleInformation //1
}
type Server2Client_createPlayerOk struct {
	player COM_SimpleInformation //0
}
type Server2Client_deletePlayerOk struct {
	name string //0
}
type Server2Client_enterGameOk struct {
	inst COM_PlayerInst //0
}
type Server2Client_initBabies struct {
	insts []COM_BabyInst //0
}
type Server2Client_initEmployees struct {
	insts  []COM_EmployeeInst //0
	isFlag bool               //1
}
type Server2Client_initEmpBattleGroup struct {
	bep COM_BattleEmp //0
}
type Server2Client_initNpc struct {
	npcList []int32 //0
}
type Server2Client_initAchievement struct {
	actlist []COM_Achievement //0
}
type Server2Client_initGather struct {
	allnum  uint32       //0
	gathers []COM_Gather //1
}
type Server2Client_initcompound struct {
	compounds []uint32 //0
}
type Server2Client_addBaby struct {
	inst COM_BabyInst //0
}
type Server2Client_refreshBaby struct {
	inst COM_BabyInst //0
}
type Server2Client_delBabyOK struct {
	babyInstId uint32 //0
}
type Server2Client_changeBabyNameOK struct {
	babyId uint32 //0
	name   string //1
}
type Server2Client_remouldBabyOK struct {
	instid uint32 //0
}
type Server2Client_intensifyBabyOK struct {
	babyid         uint32 //0
	intensifyLevel uint32 //1
}
type Server2Client_learnSkillOk struct {
	inst COM_Skill //0
}
type Server2Client_forgetSkillOk struct {
	skid uint32 //0
}
type Server2Client_addSkillExp struct {
	skid uint32 //0
	uExp uint32 //1
	flag int    //2
}
type Server2Client_babyLearnSkillOK struct {
	instId  uint32 //0
	newSkId uint32 //1
}
type Server2Client_skillLevelUp struct {
	instId uint32    //0
	inst   COM_Skill //1
}
type Server2Client_joinScene struct {
	info COM_SceneInfo //0
}
type Server2Client_joinCopySceneOK struct {
	secneid int32 //0
}
type Server2Client_addToScene struct {
	inst COM_ScenePlayerInformation //0
}
type Server2Client_delFormScene struct {
	instId int32 //0
}
type Server2Client_move2 struct {
	instId int32         //0
	pos    COM_FPosition //1
}
type Server2Client_querySimplePlayerInstOk struct {
	player COM_SimplePlayerInst //0
}
type Server2Client_transfor2 struct {
	instId int32         //0
	pos    COM_FPosition //1
}
type Server2Client_openScene struct {
	sceneId int32 //0
}
type Server2Client_autoBattleResult struct {
	isOk bool //0
}
type Server2Client_talked2Npc struct {
	npcId int32 //0
}
type Server2Client_talked2Player struct {
	playerId int32 //0
}
type Server2Client_useItemOk struct {
	itemId int32 //0
	stack  int32 //1
}
type Server2Client_syncBattleStatus struct {
	playerId int32 //0
	inBattle bool  //1
}
type Server2Client_enterBattleOk struct {
	initBattle COM_InitBattle //0
}
type Server2Client_exitBattleOk struct {
	bjt  int                    //0
	init COM_BattleOverClearing //1
}
type Server2Client_syncOrderOk struct {
	uid uint32 //0
}
type Server2Client_syncOneTurnAction struct {
	reports COM_BattleReport //0
}
type Server2Client_syncProperties struct {
	guid  uint32          //0
	props []COM_PropValue //1
}
type Server2Client_receiveChat struct {
	info   COM_ChatInfo    //0
	myinfo COM_ContactInfo //1
}
type Server2Client_requestAudioOk struct {
	audioId int32   //0
	content []uint8 //1
}
type Server2Client_publishItemInstRes struct {
	info  COM_ShowItemInstInfo //0
	type_ int                  //1
}
type Server2Client_queryItemInstRes struct {
	item COM_ShowItemInst //0
}
type Server2Client_publishBabyInstRes struct {
	info  COM_ShowbabyInstInfo //0
	type_ int                  //1
}
type Server2Client_queryBabyInstRes struct {
	inst COM_ShowbabyInst //0
}
type Server2Client_setNoTalkTime struct {
	t float32 //0
}
type Server2Client_addNpc struct {
	npcList []int32 //0
}
type Server2Client_delNpc struct {
	npcList []int32 //0
}
type Server2Client_setTeamLeader struct {
	playerId int32 //0
	isLeader bool  //1
}
type Server2Client_initBag struct {
	items []COM_Item //0
}
type Server2Client_addBagItem struct {
	item COM_Item //0
}
type Server2Client_delBagItem struct {
	slot uint16 //0
}
type Server2Client_updateBagItem struct {
	item COM_Item //0
}
type Server2Client_depositItemOK struct {
	item COM_Item //0
}
type Server2Client_getoutItemOK struct {
	slot uint16 //0
}
type Server2Client_depositBabyOK struct {
	baby COM_BabyInst //0
}
type Server2Client_getoutBabyOK struct {
	slot uint16 //0
}
type Server2Client_sortItemStorageOK struct {
	items []COM_Item //0
}
type Server2Client_sortBabyStorageOK struct {
	babys []uint32 //0
}
type Server2Client_initItemStorage struct {
	gridNum uint16     //0
	items   []COM_Item //1
}
type Server2Client_initBabyStorage struct {
	gridNum uint16         //0
	babys   []COM_BabyInst //1
}
type Server2Client_openStorageGrid struct {
	tp      int    //0
	gridNum uint16 //1
}
type Server2Client_delStorageBabyOK struct {
	slot uint16 //0
}
type Server2Client_initPlayerEquips struct {
	equips []COM_Item //0
}
type Server2Client_wearEquipmentOk struct {
	target uint32   //0
	equip  COM_Item //1
}
type Server2Client_scenePlayerWearEquipment struct {
	target uint32 //0
	itemId uint32 //1
}
type Server2Client_delEquipmentOk struct {
	target     uint32 //0
	itemInstId uint32 //1
}
type Server2Client_scenePlayerDoffEquipment struct {
	target uint32 //0
	itemId uint32 //1
}
type Server2Client_jointLobbyOk struct {
	infos []COM_SimpleTeamInfo //0
}
type Server2Client_syncDelLobbyTeam struct {
	teamId uint32 //0
}
type Server2Client_syncUpdateLobbyTeam struct {
	info COM_SimpleTeamInfo //0
}
type Server2Client_syncAddLobbyTeam struct {
	team COM_SimpleTeamInfo //0
}
type Server2Client_createTeamOk struct {
	team COM_TeamInfo //0
}
type Server2Client_changeTeamOk struct {
	team COM_TeamInfo //0
}
type Server2Client_joinTeamOk struct {
	team COM_TeamInfo //0
}
type Server2Client_addTeamMember struct {
	info COM_SimplePlayerInst //0
}
type Server2Client_delTeamMember struct {
	instId int32 //0
}
type Server2Client_changeTeamLeaderOk struct {
	uuid int32 //0
}
type Server2Client_exitTeamOk struct {
	iskick bool //0
}
type Server2Client_updateTeam struct {
	team COM_TeamInfo //0
}
type Server2Client_joinTeamRoomOK struct {
	team COM_TeamInfo //0
}
type Server2Client_inviteJoinTeam struct {
	teamId uint32 //0
	name   string //1
}
type Server2Client_syncTeamDirtyProp struct {
	guid  int32           //0
	props []COM_PropValue //1
}
type Server2Client_leaveTeamOk struct {
	playerId int32 //0
}
type Server2Client_backTeamOK struct {
	playerId int32 //0
}
type Server2Client_refuseBackTeamOk struct {
	playerId int32 //0
}
type Server2Client_requestJoinTeamTranspond struct {
	reqName string //0
}
type Server2Client_drawLotteryBoxRep struct {
	items []COM_Item //0
}
type Server2Client_addEmployee struct {
	employee COM_EmployeeInst //0
}
type Server2Client_battleEmployee struct {
	empId     int32 //0
	group     int   //1
	forbattle bool  //2
}
type Server2Client_changeEmpBattleGroupOK struct {
	group int //0
}
type Server2Client_evolveOK struct {
	guid int32 //0
	qc   int   //1
}
type Server2Client_upStarOK struct {
	guid int32     //0
	star int32     //1
	sk   COM_Skill //2
}
type Server2Client_delEmployeeOK struct {
	instids []uint32 //0
}
type Server2Client_sycnEmployeeSoul struct {
	guid    int32  //0
	soulNum uint32 //1
}
type Server2Client_initQuest struct {
	qlist []COM_QuestInst //0
	clist []int32         //1
}
type Server2Client_acceptQuestOk struct {
	inst COM_QuestInst //0
}
type Server2Client_submitQuestOk struct {
	questId int32 //0
}
type Server2Client_giveupQuestOk struct {
	questId int32 //0
}
type Server2Client_updateQuestInst struct {
	inst COM_QuestInst //0
}
type Server2Client_requestContactInfoOk struct {
	contact COM_ContactInfo //0
}
type Server2Client_addFriendOK struct {
	inst COM_ContactInfo //0
}
type Server2Client_delFriendOK struct {
	instId uint32 //0
}
type Server2Client_addBlacklistOK struct {
	inst COM_ContactInfo //0
}
type Server2Client_delBlacklistOK struct {
	instId uint32 //0
}
type Server2Client_referrFriendOK struct {
	insts []COM_ContactInfo //0
}
type Server2Client_requestFriendListOK struct {
	insts []COM_ContactInfo //0
}
type Server2Client_lotteryOk struct {
	lotteryId int32          //0
	dropItem  []COM_DropItem //1
}
type Server2Client_openGatherOK struct {
	gather COM_Gather //0
}
type Server2Client_miningOk struct {
	items     []COM_DropItem //0
	gather    COM_Gather     //1
	gatherNum uint32         //2
}
type Server2Client_openCompound struct {
	compoundId uint32 //0
}
type Server2Client_compoundItemOk struct {
	item COM_Item //0
}
type Server2Client_openBagGridOk struct {
	num int32 //0
}
type Server2Client_requestChallengeOK struct {
	isOK bool //0
}
type Server2Client_requestMySelfJJCDataOK struct {
	info COM_EndlessStair //0
}
type Server2Client_requestRivalOK struct {
	infos []COM_EndlessStair //0
}
type Server2Client_checkMsgOK struct {
	inst COM_SimplePlayerInst //0
}
type Server2Client_requestMyAllbattleMsgOK struct {
	infos []COM_JJCBattleMsg //0
}
type Server2Client_myBattleMsgOK struct {
	info COM_JJCBattleMsg //0
}
type Server2Client_requestJJCRankOK struct {
	myRank uint32             //0
	infos  []COM_EndlessStair //1
}
type Server2Client_requestLevelRankOK struct {
	myRank uint32            //0
	infos  []COM_ContactInfo //1
}
type Server2Client_requestBabyRankOK struct {
	myRank uint32             //0
	infos  []COM_BabyRankData //1
}
type Server2Client_requestEmpRankOK struct {
	myRank uint32                 //0
	infos  []COM_EmployeeRankData //1
}
type Server2Client_requestPlayerFFRankOK struct {
	myRank uint32            //0
	infos  []COM_ContactInfo //1
}
type Server2Client_queryOnlinePlayerOK struct {
	isOnline bool //0
}
type Server2Client_queryPlayerOK struct {
	inst COM_SimplePlayerInst //0
}
type Server2Client_queryBabyOK struct {
	inst COM_BabyInst //0
}
type Server2Client_queryEmployeeOK struct {
	inst COM_EmployeeInst //0
}
type Server2Client_initGuide struct {
	guideMask uint32 //0
}
type Server2Client_buyShopItemOk struct {
	id int32 //0
}
type Server2Client_addPlayerTitle struct {
	title int32 //0
}
type Server2Client_delPlayerTitle struct {
	title int32 //0
}
type Server2Client_requestOpenBuyBox struct {
	greenTime    float32 //0
	blueTime     float32 //1
	greenFreeNum int32   //2
}
type Server2Client_updateAchievementinfo struct {
	achs COM_Achievement //0
}
type Server2Client_syncOpenSystemFlag struct {
	flag uint64 //0
}
type Server2Client_requestActivityRewardOK struct {
	ar uint32 //0
}
type Server2Client_syncActivity struct {
	table COM_ActivityTable //0
}
type Server2Client_updateActivityStatus struct {
	type_ int  //0
	open  bool //1
}
type Server2Client_updateActivityCounter struct {
	type_   int   //0
	counter int32 //1
	reward  int32 //2
}
type Server2Client_syncExam struct {
	exam COM_Exam //0
}
type Server2Client_syncExamAnswer struct {
	answer COM_Answer //0
}
type Server2Client_petActivityNoNum struct {
	name string //0
}
type Server2Client_syncHundredInfo struct {
	hb COM_HundredBattle //0
}
type Server2Client_initSignUp struct {
	info    []int32 //0
	process int32   //1
	sign7   bool    //2
	sign14  bool    //3
	sign28  bool    //4
}
type Server2Client_signUp struct {
	flag bool //0
}
type Server2Client_sycnDoubleExpTime struct {
	isFlag bool    //0
	times  float32 //1
}
type Server2Client_sycnStates struct {
	states []COM_State //0
}
type Server2Client_requestpvprankOK struct {
	infos []COM_ContactInfo //0
}
type Server2Client_stopMatchingOK struct {
	times float32 //0
}
type Server2Client_updatePvpJJCinfo struct {
	info COM_PlayerVsPlayer //0
}
type Server2Client_syncEnemyPvpJJCPlayerInfo struct {
	info COM_SimpleInformation //0
}
type Server2Client_syncEnemyPvpJJCTeamInfo struct {
	infos   []COM_SimpleInformation //0
	teamID_ uint32                  //1
}
type Server2Client_syncWarriorEnemyTeamInfo struct {
	infos   []COM_SimpleInformation //0
	teamID_ uint32                  //1
}
type Server2Client_appendMail struct {
	mails []COM_Mail //0
}
type Server2Client_delMail struct {
	mailId int32 //0
}
type Server2Client_updateMailOk struct {
	mail COM_Mail //0
}
type Server2Client_boardcastNotice struct {
	content string //0
	isGm    bool   //1
}
type Server2Client_leaveGuildOk struct {
	who    string //0
	isKick bool   //1
}
type Server2Client_initGuildData struct {
	guild COM_Guild //0
}
type Server2Client_initGuildMemberList struct {
	member []COM_GuildMember //0
}
type Server2Client_modifyGuildMemberList struct {
	member COM_GuildMember //0
	flag   int             //1
}
type Server2Client_modifyGuildList struct {
	data COM_GuildViewerData //0
	flag int                 //1
}
type Server2Client_queryGuildListResult struct {
	page      int16                 //0
	pageNum   int16                 //1
	guildList []COM_GuildViewerData //2
}
type Server2Client_inviteGuild struct {
	sendName  string //0
	guildName string //1
}
type Server2Client_updateGuildShopItems struct {
	items []COM_GuildShopItem //0
}
type Server2Client_updateGuildBuilding struct {
	type_    int               //0
	building COM_GuildBuilding //1
}
type Server2Client_updateGuildMyMember struct {
	member COM_GuildMember //0
}
type Server2Client_levelupGuildSkillOk struct {
	skInst COM_Skill //0
}
type Server2Client_presentGuildItemOk struct {
	val int32 //0
}
type Server2Client_progenitusAddExpOk struct {
	mInst COM_GuildProgen //0
}
type Server2Client_setProgenitusPositionOk struct {
	positions []int32 //0
}
type Server2Client_updateGuildFundz struct {
	val int32 //0
}
type Server2Client_updateGuildMemberContribution struct {
	val int32 //0
}
type Server2Client_openGuildBattle struct {
	otherName string //0
	playerNum int32  //1
	level     int32  //2
	isLeft    bool   //3
	lstime    int32  //4
}
type Server2Client_startGuildBattle struct {
	otherName string //0
	otherCon  int32  //1
	selfCon   int32  //2
}
type Server2Client_closeGuildBattle struct {
	isWinner bool //0
}
type Server2Client_syncGuildBattleWinCount struct {
	myWin    int32 //0
	otherWin int32 //1
}
type Server2Client_initMySelling struct {
	items []COM_SellItem //0
}
type Server2Client_initMySelled struct {
	items []COM_SelledItem //0
}
type Server2Client_fetchSellingOk struct {
	items []COM_SellItem //0
	total int32          //1
}
type Server2Client_fetchSellingOk2 struct {
	items []COM_SellItem //0
	total int32          //1
}
type Server2Client_sellingOk struct {
	item COM_SellItem //0
}
type Server2Client_selledOk struct {
	item COM_SelledItem //0
}
type Server2Client_unsellingOk struct {
	sellid int32 //0
}
type Server2Client_insertState struct {
	st COM_State //0
}
type Server2Client_updattState struct {
	st COM_State //0
}
type Server2Client_removeState struct {
	stid uint32 //0
}
type Server2Client_requestFixItemOk struct {
	item COM_Item //0
}
type Server2Client_updateMagicItem struct {
	level int32 //0
	exp   int32 //1
}
type Server2Client_changeMagicJobOk struct {
	job int //0
}
type Server2Client_magicItemTupoOk struct {
	level int32 //0
}
type Server2Client_zhuanpanOK struct {
	pond []uint32 //0
}
type Server2Client_updateZhuanpanNotice struct {
	zhuanp COM_Zhuanpan //0
}
type Server2Client_sycnZhuanpanData struct {
	data COM_ZhuanpanData //0
}
type Server2Client_copynonum struct {
	name string //0
}
type Server2Client_sceneFilterOk struct {
	sfType []int //0
}
type Server2Client_shareWishOK struct {
	wish COM_Wish //0
}
type Server2Client_requestOnlineTimeRewardOK struct {
	index uint32 //0
}
type Server2Client_sycnVipflag struct {
	flag bool //0
}
type Server2Client_buyFundOK struct {
	flag bool //0
}
type Server2Client_requestFundRewardOK struct {
	level uint32 //0
}
type Server2Client_firstRechargeOK struct {
	isFlag bool //0
}
type Server2Client_firstRechargeGiftOK struct {
	isFlag bool //0
}
type Server2Client_agencyActivity struct {
	type_  int  //0
	isFlag bool //1
}
type Server2Client_updateFestival struct {
	festival COM_ADLoginTotal //0
}
type Server2Client_updateSelfRecharge struct {
	val COM_ADChargeTotal //0
}
type Server2Client_updateSysRecharge struct {
	val COM_ADChargeTotal //0
}
type Server2Client_updateSelfDiscountStore struct {
	val COM_ADDiscountStore //0
}
type Server2Client_updateSysDiscountStore struct {
	val COM_ADDiscountStore //0
}
type Server2Client_updateSelfOnceRecharge struct {
	val COM_ADChargeEvery //0
}
type Server2Client_updateSysOnceRecharge struct {
	val COM_ADChargeEvery //0
}
type Server2Client_openCardOK struct {
	data COM_ADCardsContent //0
}
type Server2Client_sycnHotRole struct {
	data COM_ADHotRole //0
}
type Server2Client_hotRoleBuyOk struct {
	buyNum uint16 //0
}
type Server2Client_updateSevenday struct {
	data COM_Sevenday //0
}
type Server2Client_updateEmployeeActivity struct {
	data COM_ADEmployeeTotal //0
}
type Server2Client_updateMinGiftActivity struct {
	data COM_ADGiftBag //0
}
type Server2Client_updateIntegralShop struct {
	data COM_IntegralData //0
}
type Server2Client_updateShowBaby struct {
	playerId        uint32 //0
	showBabyTableId uint32 //1
	showBabyName    string //2
}
type Server2Client_updateMySelfRecharge struct {
	val COM_ADChargeTotal //0
}
type Server2Client_verificationSMSOk struct {
	phoneNumber string //0
}
type Server2Client_requestLevelGiftOK struct {
	level int32 //0
}
type Server2Client_sycnConvertExp struct {
	val int32 //0
}
type Server2Client_wearFuwenOk struct {
	inst COM_Item //0
}
type Server2Client_takeoffFuwenOk struct {
	slot int32 //0
}
type Server2Client_requestEmployeeQuestOk struct {
	questList []COM_EmployeeQuestInst //0
}
type Server2Client_acceptEmployeeQuestOk struct {
	inst COM_EmployeeQuestInst //0
}
type Server2Client_submitEmployeeQuestOk struct {
	questId   int32 //0
	isSuccess bool  //1
}
type Server2Client_sycnCrystal struct {
	data COM_CrystalData //0
}
type Server2Client_crystalUpLeveResult struct {
	isOK bool //0
}
type Server2Client_sycnCourseGift struct {
	data []COM_CourseGift //0
}
type Server2Client_orderOk struct {
	orderId string //0
	shopId  int32  //1
}
type Server2Client_updateRandSubmitQuestCount struct {
	submitCount int32 //0
}
type Server2ClientStub struct {
	Sender prpc.StubSender
}
type Server2ClientProxy interface {
	pong() error                                                                                     // 0
	errorno(e int) error                                                                             // 1
	teamerrorno(name string, e int) error                                                            // 2
	reconnection(recInfo COM_ReconnectInfo) error                                                    // 3
	sessionfailed() error                                                                            // 4
	loginok(sessionkey string, players []COM_SimpleInformation) error                                // 5
	logoutOk() error                                                                                 // 6
	createPlayerOk(player COM_SimpleInformation) error                                               // 7
	deletePlayerOk(name string) error                                                                // 8
	enterGameOk(inst COM_PlayerInst) error                                                           // 9
	initBabies(insts []COM_BabyInst) error                                                           // 10
	initEmployees(insts []COM_EmployeeInst, isFlag bool) error                                       // 11
	initEmpBattleGroup(bep COM_BattleEmp) error                                                      // 12
	initNpc(npcList []int32) error                                                                   // 13
	initAchievement(actlist []COM_Achievement) error                                                 // 14
	initGather(allnum uint32, gathers []COM_Gather) error                                            // 15
	initcompound(compounds []uint32) error                                                           // 16
	addBaby(inst COM_BabyInst) error                                                                 // 17
	refreshBaby(inst COM_BabyInst) error                                                             // 18
	delBabyOK(babyInstId uint32) error                                                               // 19
	changeBabyNameOK(babyId uint32, name string) error                                               // 20
	remouldBabyOK(instid uint32) error                                                               // 21
	intensifyBabyOK(babyid uint32, intensifyLevel uint32) error                                      // 22
	learnSkillOk(inst COM_Skill) error                                                               // 23
	forgetSkillOk(skid uint32) error                                                                 // 24
	addSkillExp(skid uint32, uExp uint32, flag int) error                                            // 25
	babyLearnSkillOK(instId uint32, newSkId uint32) error                                            // 26
	skillLevelUp(instId uint32, inst COM_Skill) error                                                // 27
	joinScene(info COM_SceneInfo) error                                                              // 28
	joinCopySceneOK(secneid int32) error                                                             // 29
	initCopyNums() error                                                                             // 30
	addToScene(inst COM_ScenePlayerInformation) error                                                // 31
	delFormScene(instId int32) error                                                                 // 32
	move2(instId int32, pos COM_FPosition) error                                                     // 33
	cantMove() error                                                                                 // 34
	querySimplePlayerInstOk(player COM_SimplePlayerInst) error                                       // 35
	transfor2(instId int32, pos COM_FPosition) error                                                 // 36
	openScene(sceneId int32) error                                                                   // 37
	autoBattleResult(isOk bool) error                                                                // 38
	talked2Npc(npcId int32) error                                                                    // 39
	talked2Player(playerId int32) error                                                              // 40
	useItemOk(itemId int32, stack int32) error                                                       // 41
	syncBattleStatus(playerId int32, inBattle bool) error                                            // 42
	enterBattleOk(initBattle COM_InitBattle) error                                                   // 43
	exitBattleOk(bjt int, init COM_BattleOverClearing) error                                         // 44
	syncOrderOk(uid uint32) error                                                                    // 45
	syncOrderOkEX() error                                                                            // 46
	syncOneTurnAction(reports COM_BattleReport) error                                                // 47
	syncProperties(guid uint32, props []COM_PropValue) error                                         // 48
	receiveChat(info COM_ChatInfo, myinfo COM_ContactInfo) error                                     // 49
	requestAudioOk(audioId int32, content []uint8) error                                             // 50
	publishItemInstRes(info COM_ShowItemInstInfo, type_ int) error                                   // 51
	queryItemInstRes(item COM_ShowItemInst) error                                                    // 52
	publishBabyInstRes(info COM_ShowbabyInstInfo, type_ int) error                                   // 53
	queryBabyInstRes(inst COM_ShowbabyInst) error                                                    // 54
	setNoTalkTime(t float32) error                                                                   // 55
	addNpc(npcList []int32) error                                                                    // 56
	delNpc(npcList []int32) error                                                                    // 57
	setTeamLeader(playerId int32, isLeader bool) error                                               // 58
	initBag(items []COM_Item) error                                                                  // 59
	addBagItem(item COM_Item) error                                                                  // 60
	delBagItem(slot uint16) error                                                                    // 61
	updateBagItem(item COM_Item) error                                                               // 62
	depositItemOK(item COM_Item) error                                                               // 63
	getoutItemOK(slot uint16) error                                                                  // 64
	depositBabyOK(baby COM_BabyInst) error                                                           // 65
	getoutBabyOK(slot uint16) error                                                                  // 66
	sortItemStorageOK(items []COM_Item) error                                                        // 67
	sortBabyStorageOK(babys []uint32) error                                                          // 68
	initItemStorage(gridNum uint16, items []COM_Item) error                                          // 69
	initBabyStorage(gridNum uint16, babys []COM_BabyInst) error                                      // 70
	openStorageGrid(tp int, gridNum uint16) error                                                    // 71
	delStorageBabyOK(slot uint16) error                                                              // 72
	initPlayerEquips(equips []COM_Item) error                                                        // 73
	wearEquipmentOk(target uint32, equip COM_Item) error                                             // 74
	scenePlayerWearEquipment(target uint32, itemId uint32) error                                     // 75
	delEquipmentOk(target uint32, itemInstId uint32) error                                           // 76
	scenePlayerDoffEquipment(target uint32, itemId uint32) error                                     // 77
	sortBagItemOk() error                                                                            // 78
	jointLobbyOk(infos []COM_SimpleTeamInfo) error                                                   // 79
	exitLobbyOk() error                                                                              // 80
	syncDelLobbyTeam(teamId uint32) error                                                            // 81
	syncUpdateLobbyTeam(info COM_SimpleTeamInfo) error                                               // 82
	syncAddLobbyTeam(team COM_SimpleTeamInfo) error                                                  // 83
	createTeamOk(team COM_TeamInfo) error                                                            // 84
	changeTeamOk(team COM_TeamInfo) error                                                            // 85
	joinTeamOk(team COM_TeamInfo) error                                                              // 86
	addTeamMember(info COM_SimplePlayerInst) error                                                   // 87
	delTeamMember(instId int32) error                                                                // 88
	changeTeamLeaderOk(uuid int32) error                                                             // 89
	exitTeamOk(iskick bool) error                                                                    // 90
	updateTeam(team COM_TeamInfo) error                                                              // 91
	joinTeamRoomOK(team COM_TeamInfo) error                                                          // 92
	inviteJoinTeam(teamId uint32, name string) error                                                 // 93
	syncTeamDirtyProp(guid int32, props []COM_PropValue) error                                       // 94
	leaveTeamOk(playerId int32) error                                                                // 95
	backTeamOK(playerId int32) error                                                                 // 96
	teamCallMemberBack() error                                                                       // 97
	refuseBackTeamOk(playerId int32) error                                                           // 98
	requestJoinTeamTranspond(reqName string) error                                                   // 99
	drawLotteryBoxRep(items []COM_Item) error                                                        // 100
	addEmployee(employee COM_EmployeeInst) error                                                     // 101
	battleEmployee(empId int32, group int, forbattle bool) error                                     // 102
	changeEmpBattleGroupOK(group int) error                                                          // 103
	evolveOK(guid int32, qc int) error                                                               // 104
	upStarOK(guid int32, star int32, sk COM_Skill) error                                             // 105
	delEmployeeOK(instids []uint32) error                                                            // 106
	sycnEmployeeSoul(guid int32, soulNum uint32) error                                               // 107
	initQuest(qlist []COM_QuestInst, clist []int32) error                                            // 108
	acceptQuestOk(inst COM_QuestInst) error                                                          // 109
	submitQuestOk(questId int32) error                                                               // 110
	giveupQuestOk(questId int32) error                                                               // 111
	updateQuestInst(inst COM_QuestInst) error                                                        // 112
	requestContactInfoOk(contact COM_ContactInfo) error                                              // 113
	addFriendOK(inst COM_ContactInfo) error                                                          // 114
	delFriendOK(instId uint32) error                                                                 // 115
	addBlacklistOK(inst COM_ContactInfo) error                                                       // 116
	delBlacklistOK(instId uint32) error                                                              // 117
	findFriendFail() error                                                                           // 118
	referrFriendOK(insts []COM_ContactInfo) error                                                    // 119
	requestFriendListOK(insts []COM_ContactInfo) error                                               // 120
	lotteryOk(lotteryId int32, dropItem []COM_DropItem) error                                        // 121
	openGatherOK(gather COM_Gather) error                                                            // 122
	miningOk(items []COM_DropItem, gather COM_Gather, gatherNum uint32) error                        // 123
	openCompound(compoundId uint32) error                                                            // 124
	compoundItemOk(item COM_Item) error                                                              // 125
	openBagGridOk(num int32) error                                                                   // 126
	requestChallengeOK(isOK bool) error                                                              // 127
	requestMySelfJJCDataOK(info COM_EndlessStair) error                                              // 128
	requestRivalOK(infos []COM_EndlessStair) error                                                   // 129
	rivalTimeOK() error                                                                              // 130
	checkMsgOK(inst COM_SimplePlayerInst) error                                                      // 131
	requestMyAllbattleMsgOK(infos []COM_JJCBattleMsg) error                                          // 132
	myBattleMsgOK(info COM_JJCBattleMsg) error                                                       // 133
	requestJJCRankOK(myRank uint32, infos []COM_EndlessStair) error                                  // 134
	requestLevelRankOK(myRank uint32, infos []COM_ContactInfo) error                                 // 135
	requestBabyRankOK(myRank uint32, infos []COM_BabyRankData) error                                 // 136
	requestEmpRankOK(myRank uint32, infos []COM_EmployeeRankData) error                              // 137
	requestPlayerFFRankOK(myRank uint32, infos []COM_ContactInfo) error                              // 138
	queryOnlinePlayerOK(isOnline bool) error                                                         // 139
	queryPlayerOK(inst COM_SimplePlayerInst) error                                                   // 140
	queryBabyOK(inst COM_BabyInst) error                                                             // 141
	queryEmployeeOK(inst COM_EmployeeInst) error                                                     // 142
	initGuide(guideMask uint32) error                                                                // 143
	buyShopItemOk(id int32) error                                                                    // 144
	addPlayerTitle(title int32) error                                                                // 145
	delPlayerTitle(title int32) error                                                                // 146
	requestOpenBuyBox(greenTime float32, blueTime float32, greenFreeNum int32) error                 // 147
	requestGreenBoxTimeOk() error                                                                    // 148
	requestBlueBoxTimeOk() error                                                                     // 149
	updateAchievementinfo(achs COM_Achievement) error                                                // 150
	syncOpenSystemFlag(flag uint64) error                                                            // 151
	requestActivityRewardOK(ar uint32) error                                                         // 152
	syncActivity(table COM_ActivityTable) error                                                      // 153
	updateActivityStatus(type_ int, open bool) error                                                 // 154
	updateActivityCounter(type_ int, counter int32, reward int32) error                              // 155
	syncExam(exam COM_Exam) error                                                                    // 156
	syncExamAnswer(answer COM_Answer) error                                                          // 157
	petActivityNoNum(name string) error                                                              // 158
	syncHundredInfo(hb COM_HundredBattle) error                                                      // 159
	initSignUp(info []int32, process int32, sign7 bool, sign14 bool, sign28 bool) error              // 160
	signUp(flag bool) error                                                                          // 161
	requestSignupRewardOk7() error                                                                   // 162
	requestSignupRewardOk14() error                                                                  // 163
	requestSignupRewardOk28() error                                                                  // 164
	sycnDoubleExpTime(isFlag bool, times float32) error                                              // 165
	sycnStates(states []COM_State) error                                                             // 166
	requestpvprankOK(infos []COM_ContactInfo) error                                                  // 167
	syncMyJJCTeamMember() error                                                                      // 168
	startMatchingOK() error                                                                          // 169
	stopMatchingOK(times float32) error                                                              // 170
	updatePvpJJCinfo(info COM_PlayerVsPlayer) error                                                  // 171
	exitPvpJJCOk() error                                                                             // 172
	syncEnemyPvpJJCPlayerInfo(info COM_SimpleInformation) error                                      // 173
	syncEnemyPvpJJCTeamInfo(infos []COM_SimpleInformation, teamID_ uint32) error                     // 174
	openWarriorchooseUI() error                                                                      // 175
	warriorStartOK() error                                                                           // 176
	warriorStopOK() error                                                                            // 177
	syncWarriorEnemyTeamInfo(infos []COM_SimpleInformation, teamID_ uint32) error                    // 178
	appendMail(mails []COM_Mail) error                                                               // 179
	delMail(mailId int32) error                                                                      // 180
	updateMailOk(mail COM_Mail) error                                                                // 181
	boardcastNotice(content string, isGm bool) error                                                 // 182
	createGuildOK() error                                                                            // 183
	delGuildOK() error                                                                               // 184
	leaveGuildOk(who string, isKick bool) error                                                      // 185
	initGuildData(guild COM_Guild) error                                                             // 186
	initGuildMemberList(member []COM_GuildMember) error                                              // 187
	modifyGuildMemberList(member COM_GuildMember, flag int) error                                    // 188
	modifyGuildList(data COM_GuildViewerData, flag int) error                                        // 189
	queryGuildListResult(page int16, pageNum int16, guildList []COM_GuildViewerData) error           // 190
	inviteGuild(sendName string, guildName string) error                                             // 191
	updateGuildShopItems(items []COM_GuildShopItem) error                                            // 192
	updateGuildBuilding(type_ int, building COM_GuildBuilding) error                                 // 193
	updateGuildMyMember(member COM_GuildMember) error                                                // 194
	levelupGuildSkillOk(skInst COM_Skill) error                                                      // 195
	presentGuildItemOk(val int32) error                                                              // 196
	progenitusAddExpOk(mInst COM_GuildProgen) error                                                  // 197
	setProgenitusPositionOk(positions []int32) error                                                 // 198
	updateGuildFundz(val int32) error                                                                // 199
	updateGuildMemberContribution(val int32) error                                                   // 200
	openGuildBattle(otherName string, playerNum int32, level int32, isLeft bool, lstime int32) error // 201
	startGuildBattle(otherName string, otherCon int32, selfCon int32) error                          // 202
	closeGuildBattle(isWinner bool) error                                                            // 203
	syncGuildBattleWinCount(myWin int32, otherWin int32) error                                       // 204
	initMySelling(items []COM_SellItem) error                                                        // 205
	initMySelled(items []COM_SelledItem) error                                                       // 206
	fetchSellingOk(items []COM_SellItem, total int32) error                                          // 207
	fetchSellingOk2(items []COM_SellItem, total int32) error                                         // 208
	sellingOk(item COM_SellItem) error                                                               // 209
	selledOk(item COM_SelledItem) error                                                              // 210
	unsellingOk(sellid int32) error                                                                  // 211
	redemptionSpreeOk() error                                                                        // 212
	insertState(st COM_State) error                                                                  // 213
	updattState(st COM_State) error                                                                  // 214
	removeState(stid uint32) error                                                                   // 215
	requestFixItemOk(item COM_Item) error                                                            // 216
	makeDebirsItemOK() error                                                                         // 217
	updateMagicItem(level int32, exp int32) error                                                    // 218
	changeMagicJobOk(job int) error                                                                  // 219
	magicItemTupoOk(level int32) error                                                               // 220
	zhuanpanOK(pond []uint32) error                                                                  // 221
	updateZhuanpanNotice(zhuanp COM_Zhuanpan) error                                                  // 222
	sycnZhuanpanData(data COM_ZhuanpanData) error                                                    // 223
	copynonum(name string) error                                                                     // 224
	sceneFilterOk(sfType []int) error                                                                // 225
	wishingOK() error                                                                                // 226
	shareWishOK(wish COM_Wish) error                                                                 // 227
	leaderCloseDialogOk() error                                                                      // 228
	startOnlineTime() error                                                                          // 229
	requestOnlineTimeRewardOK(index uint32) error                                                    // 230
	sycnVipflag(flag bool) error                                                                     // 231
	buyFundOK(flag bool) error                                                                       // 232
	requestFundRewardOK(level uint32) error                                                          // 233
	firstRechargeOK(isFlag bool) error                                                               // 234
	firstRechargeGiftOK(isFlag bool) error                                                           // 235
	agencyActivity(type_ int, isFlag bool) error                                                     // 236
	updateFestival(festival COM_ADLoginTotal) error                                                  // 237
	updateSelfRecharge(val COM_ADChargeTotal) error                                                  // 238
	updateSysRecharge(val COM_ADChargeTotal) error                                                   // 239
	updateSelfDiscountStore(val COM_ADDiscountStore) error                                           // 240
	updateSysDiscountStore(val COM_ADDiscountStore) error                                            // 241
	updateSelfOnceRecharge(val COM_ADChargeEvery) error                                              // 242
	updateSysOnceRecharge(val COM_ADChargeEvery) error                                               // 243
	openCardOK(data COM_ADCardsContent) error                                                        // 244
	resetCardOK() error                                                                              // 245
	sycnHotRole(data COM_ADHotRole) error                                                            // 246
	hotRoleBuyOk(buyNum uint16) error                                                                // 247
	updateSevenday(data COM_Sevenday) error                                                          // 248
	updateEmployeeActivity(data COM_ADEmployeeTotal) error                                           // 249
	updateMinGiftActivity(data COM_ADGiftBag) error                                                  // 250
	updateIntegralShop(data COM_IntegralData) error                                                  // 251
	updateShowBaby(playerId uint32, showBabyTableId uint32, showBabyName string) error               // 252
	updateMySelfRecharge(val COM_ADChargeTotal) error                                                // 253
	verificationSMSOk(phoneNumber string) error                                                      // 254
	requestLevelGiftOK(level int32) error                                                            // 255
	sycnConvertExp(val int32) error                                                                  // 256
	wearFuwenOk(inst COM_Item) error                                                                 // 257
	takeoffFuwenOk(slot int32) error                                                                 // 258
	compFuwenOk() error                                                                              // 259
	requestEmployeeQuestOk(questList []COM_EmployeeQuestInst) error                                  // 260
	acceptEmployeeQuestOk(inst COM_EmployeeQuestInst) error                                          // 261
	submitEmployeeQuestOk(questId int32, isSuccess bool) error                                       // 262
	sycnCrystal(data COM_CrystalData) error                                                          // 263
	crystalUpLeveResult(isOK bool) error                                                             // 264
	resetCrystalPropOK() error                                                                       // 265
	sycnCourseGift(data []COM_CourseGift) error                                                      // 266
	orderOk(orderId string, shopId int32) error                                                      // 267
	updateRandSubmitQuestCount(submitCount int32) error                                              // 268
}

func (this *Server2Client_errorno) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.e != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize e
	{
		if this.e != 0 {
			err := prpc.Write(buffer, this.e)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_errorno) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize e
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.e)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_teamerrorno) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.name) != 0)
	mask.WriteBit(this.e != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize name
	if len(this.name) != 0 {
		err := prpc.Write(buffer, this.name)
		if err != nil {
			return err
		}
	}
	// serialize e
	{
		if this.e != 0 {
			err := prpc.Write(buffer, this.e)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_teamerrorno) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize name
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.name)
		if err != nil {
			return err
		}
	}
	// deserialize e
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.e)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_reconnection) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //recInfo
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize recInfo
	{
		err := this.recInfo.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_reconnection) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize recInfo
	if mask.ReadBit() {
		err := this.recInfo.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_loginok) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.sessionkey) != 0)
	mask.WriteBit(len(this.players) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize sessionkey
	if len(this.sessionkey) != 0 {
		err := prpc.Write(buffer, this.sessionkey)
		if err != nil {
			return err
		}
	}
	// serialize players
	if len(this.players) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.players)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.players {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_loginok) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize sessionkey
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.sessionkey)
		if err != nil {
			return err
		}
	}
	// deserialize players
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.players = make([]COM_SimpleInformation, size)
		for i, _ := range this.players {
			err := this.players[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_createPlayerOk) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //player
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize player
	{
		err := this.player.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_createPlayerOk) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize player
	if mask.ReadBit() {
		err := this.player.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_deletePlayerOk) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.name) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize name
	if len(this.name) != 0 {
		err := prpc.Write(buffer, this.name)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_deletePlayerOk) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize name
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.name)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_enterGameOk) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //inst
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize inst
	{
		err := this.inst.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_enterGameOk) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize inst
	if mask.ReadBit() {
		err := this.inst.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_initBabies) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.insts) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize insts
	if len(this.insts) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.insts)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.insts {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_initBabies) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize insts
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.insts = make([]COM_BabyInst, size)
		for i, _ := range this.insts {
			err := this.insts[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_initEmployees) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.insts) != 0)
	mask.WriteBit(this.isFlag)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize insts
	if len(this.insts) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.insts)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.insts {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// serialize isFlag
	{
	}
	return nil
}
func (this *Server2Client_initEmployees) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize insts
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.insts = make([]COM_EmployeeInst, size)
		for i, _ := range this.insts {
			err := this.insts[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// deserialize isFlag
	this.isFlag = mask.ReadBit()
	return nil
}
func (this *Server2Client_initEmpBattleGroup) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //bep
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize bep
	{
		err := this.bep.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_initEmpBattleGroup) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize bep
	if mask.ReadBit() {
		err := this.bep.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_initNpc) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.npcList) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize npcList
	if len(this.npcList) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.npcList)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.npcList {
			err := prpc.Write(buffer, value)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_initNpc) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize npcList
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.npcList = make([]int32, size)
		for i, _ := range this.npcList {
			err := prpc.Read(buffer, &this.npcList[i])
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_initAchievement) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.actlist) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize actlist
	if len(this.actlist) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.actlist)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.actlist {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_initAchievement) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize actlist
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.actlist = make([]COM_Achievement, size)
		for i, _ := range this.actlist {
			err := this.actlist[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_initGather) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.allnum != 0)
	mask.WriteBit(len(this.gathers) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize allnum
	{
		if this.allnum != 0 {
			err := prpc.Write(buffer, this.allnum)
			if err != nil {
				return err
			}
		}
	}
	// serialize gathers
	if len(this.gathers) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.gathers)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.gathers {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_initGather) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize allnum
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.allnum)
		if err != nil {
			return err
		}
	}
	// deserialize gathers
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.gathers = make([]COM_Gather, size)
		for i, _ := range this.gathers {
			err := this.gathers[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_initcompound) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.compounds) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize compounds
	if len(this.compounds) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.compounds)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.compounds {
			err := prpc.Write(buffer, value)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_initcompound) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize compounds
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.compounds = make([]uint32, size)
		for i, _ := range this.compounds {
			err := prpc.Read(buffer, &this.compounds[i])
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_addBaby) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //inst
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize inst
	{
		err := this.inst.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_addBaby) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize inst
	if mask.ReadBit() {
		err := this.inst.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_refreshBaby) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //inst
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize inst
	{
		err := this.inst.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_refreshBaby) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize inst
	if mask.ReadBit() {
		err := this.inst.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_delBabyOK) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.babyInstId != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize babyInstId
	{
		if this.babyInstId != 0 {
			err := prpc.Write(buffer, this.babyInstId)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_delBabyOK) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize babyInstId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.babyInstId)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_changeBabyNameOK) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.babyId != 0)
	mask.WriteBit(len(this.name) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize babyId
	{
		if this.babyId != 0 {
			err := prpc.Write(buffer, this.babyId)
			if err != nil {
				return err
			}
		}
	}
	// serialize name
	if len(this.name) != 0 {
		err := prpc.Write(buffer, this.name)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_changeBabyNameOK) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize babyId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.babyId)
		if err != nil {
			return err
		}
	}
	// deserialize name
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.name)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_remouldBabyOK) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.instid != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize instid
	{
		if this.instid != 0 {
			err := prpc.Write(buffer, this.instid)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_remouldBabyOK) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize instid
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.instid)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_intensifyBabyOK) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.babyid != 0)
	mask.WriteBit(this.intensifyLevel != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize babyid
	{
		if this.babyid != 0 {
			err := prpc.Write(buffer, this.babyid)
			if err != nil {
				return err
			}
		}
	}
	// serialize intensifyLevel
	{
		if this.intensifyLevel != 0 {
			err := prpc.Write(buffer, this.intensifyLevel)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_intensifyBabyOK) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize babyid
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.babyid)
		if err != nil {
			return err
		}
	}
	// deserialize intensifyLevel
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.intensifyLevel)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_learnSkillOk) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //inst
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize inst
	{
		err := this.inst.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_learnSkillOk) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize inst
	if mask.ReadBit() {
		err := this.inst.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_forgetSkillOk) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.skid != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize skid
	{
		if this.skid != 0 {
			err := prpc.Write(buffer, this.skid)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_forgetSkillOk) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize skid
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.skid)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_addSkillExp) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.skid != 0)
	mask.WriteBit(this.uExp != 0)
	mask.WriteBit(this.flag != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize skid
	{
		if this.skid != 0 {
			err := prpc.Write(buffer, this.skid)
			if err != nil {
				return err
			}
		}
	}
	// serialize uExp
	{
		if this.uExp != 0 {
			err := prpc.Write(buffer, this.uExp)
			if err != nil {
				return err
			}
		}
	}
	// serialize flag
	{
		if this.flag != 0 {
			err := prpc.Write(buffer, this.flag)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_addSkillExp) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize skid
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.skid)
		if err != nil {
			return err
		}
	}
	// deserialize uExp
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.uExp)
		if err != nil {
			return err
		}
	}
	// deserialize flag
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.flag)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_babyLearnSkillOK) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.instId != 0)
	mask.WriteBit(this.newSkId != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize instId
	{
		if this.instId != 0 {
			err := prpc.Write(buffer, this.instId)
			if err != nil {
				return err
			}
		}
	}
	// serialize newSkId
	{
		if this.newSkId != 0 {
			err := prpc.Write(buffer, this.newSkId)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_babyLearnSkillOK) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize instId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.instId)
		if err != nil {
			return err
		}
	}
	// deserialize newSkId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.newSkId)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_skillLevelUp) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.instId != 0)
	mask.WriteBit(true) //inst
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize instId
	{
		if this.instId != 0 {
			err := prpc.Write(buffer, this.instId)
			if err != nil {
				return err
			}
		}
	}
	// serialize inst
	{
		err := this.inst.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_skillLevelUp) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize instId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.instId)
		if err != nil {
			return err
		}
	}
	// deserialize inst
	if mask.ReadBit() {
		err := this.inst.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_joinScene) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //info
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize info
	{
		err := this.info.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_joinScene) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize info
	if mask.ReadBit() {
		err := this.info.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_joinCopySceneOK) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.secneid != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize secneid
	{
		if this.secneid != 0 {
			err := prpc.Write(buffer, this.secneid)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_joinCopySceneOK) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize secneid
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.secneid)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_addToScene) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //inst
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize inst
	{
		err := this.inst.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_addToScene) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize inst
	if mask.ReadBit() {
		err := this.inst.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_delFormScene) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.instId != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize instId
	{
		if this.instId != 0 {
			err := prpc.Write(buffer, this.instId)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_delFormScene) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize instId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.instId)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_move2) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.instId != 0)
	mask.WriteBit(true) //pos
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize instId
	{
		if this.instId != 0 {
			err := prpc.Write(buffer, this.instId)
			if err != nil {
				return err
			}
		}
	}
	// serialize pos
	{
		err := this.pos.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_move2) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize instId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.instId)
		if err != nil {
			return err
		}
	}
	// deserialize pos
	if mask.ReadBit() {
		err := this.pos.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_querySimplePlayerInstOk) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //player
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize player
	{
		err := this.player.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_querySimplePlayerInstOk) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize player
	if mask.ReadBit() {
		err := this.player.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_transfor2) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.instId != 0)
	mask.WriteBit(true) //pos
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize instId
	{
		if this.instId != 0 {
			err := prpc.Write(buffer, this.instId)
			if err != nil {
				return err
			}
		}
	}
	// serialize pos
	{
		err := this.pos.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_transfor2) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize instId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.instId)
		if err != nil {
			return err
		}
	}
	// deserialize pos
	if mask.ReadBit() {
		err := this.pos.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_openScene) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.sceneId != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize sceneId
	{
		if this.sceneId != 0 {
			err := prpc.Write(buffer, this.sceneId)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_openScene) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize sceneId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.sceneId)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_autoBattleResult) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.isOk)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize isOk
	{
	}
	return nil
}
func (this *Server2Client_autoBattleResult) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize isOk
	this.isOk = mask.ReadBit()
	return nil
}
func (this *Server2Client_talked2Npc) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.npcId != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize npcId
	{
		if this.npcId != 0 {
			err := prpc.Write(buffer, this.npcId)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_talked2Npc) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize npcId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.npcId)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_talked2Player) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.playerId != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize playerId
	{
		if this.playerId != 0 {
			err := prpc.Write(buffer, this.playerId)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_talked2Player) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize playerId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.playerId)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_useItemOk) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.itemId != 0)
	mask.WriteBit(this.stack != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize itemId
	{
		if this.itemId != 0 {
			err := prpc.Write(buffer, this.itemId)
			if err != nil {
				return err
			}
		}
	}
	// serialize stack
	{
		if this.stack != 0 {
			err := prpc.Write(buffer, this.stack)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_useItemOk) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize itemId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.itemId)
		if err != nil {
			return err
		}
	}
	// deserialize stack
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.stack)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_syncBattleStatus) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.playerId != 0)
	mask.WriteBit(this.inBattle)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize playerId
	{
		if this.playerId != 0 {
			err := prpc.Write(buffer, this.playerId)
			if err != nil {
				return err
			}
		}
	}
	// serialize inBattle
	{
	}
	return nil
}
func (this *Server2Client_syncBattleStatus) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize playerId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.playerId)
		if err != nil {
			return err
		}
	}
	// deserialize inBattle
	this.inBattle = mask.ReadBit()
	return nil
}
func (this *Server2Client_enterBattleOk) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //initBattle
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize initBattle
	{
		err := this.initBattle.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_enterBattleOk) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize initBattle
	if mask.ReadBit() {
		err := this.initBattle.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_exitBattleOk) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.bjt != 0)
	mask.WriteBit(true) //init
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize bjt
	{
		if this.bjt != 0 {
			err := prpc.Write(buffer, this.bjt)
			if err != nil {
				return err
			}
		}
	}
	// serialize init
	{
		err := this.init.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_exitBattleOk) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize bjt
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.bjt)
		if err != nil {
			return err
		}
	}
	// deserialize init
	if mask.ReadBit() {
		err := this.init.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_syncOrderOk) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.uid != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize uid
	{
		if this.uid != 0 {
			err := prpc.Write(buffer, this.uid)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_syncOrderOk) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize uid
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.uid)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_syncOneTurnAction) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //reports
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize reports
	{
		err := this.reports.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_syncOneTurnAction) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize reports
	if mask.ReadBit() {
		err := this.reports.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_syncProperties) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.guid != 0)
	mask.WriteBit(len(this.props) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize guid
	{
		if this.guid != 0 {
			err := prpc.Write(buffer, this.guid)
			if err != nil {
				return err
			}
		}
	}
	// serialize props
	if len(this.props) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.props)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.props {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_syncProperties) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize guid
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.guid)
		if err != nil {
			return err
		}
	}
	// deserialize props
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.props = make([]COM_PropValue, size)
		for i, _ := range this.props {
			err := this.props[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_receiveChat) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //info
	mask.WriteBit(true) //myinfo
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize info
	{
		err := this.info.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	// serialize myinfo
	{
		err := this.myinfo.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_receiveChat) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize info
	if mask.ReadBit() {
		err := this.info.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	// deserialize myinfo
	if mask.ReadBit() {
		err := this.myinfo.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_requestAudioOk) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.audioId != 0)
	mask.WriteBit(len(this.content) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize audioId
	{
		if this.audioId != 0 {
			err := prpc.Write(buffer, this.audioId)
			if err != nil {
				return err
			}
		}
	}
	// serialize content
	if len(this.content) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.content)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.content {
			err := prpc.Write(buffer, value)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_requestAudioOk) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize audioId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.audioId)
		if err != nil {
			return err
		}
	}
	// deserialize content
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.content = make([]uint8, size)
		for i, _ := range this.content {
			err := prpc.Read(buffer, &this.content[i])
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_publishItemInstRes) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //info
	mask.WriteBit(this.type_ != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize info
	{
		err := this.info.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	// serialize type_
	{
		if this.type_ != 0 {
			err := prpc.Write(buffer, this.type_)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_publishItemInstRes) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize info
	if mask.ReadBit() {
		err := this.info.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	// deserialize type_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.type_)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_queryItemInstRes) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //item
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize item
	{
		err := this.item.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_queryItemInstRes) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize item
	if mask.ReadBit() {
		err := this.item.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_publishBabyInstRes) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //info
	mask.WriteBit(this.type_ != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize info
	{
		err := this.info.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	// serialize type_
	{
		if this.type_ != 0 {
			err := prpc.Write(buffer, this.type_)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_publishBabyInstRes) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize info
	if mask.ReadBit() {
		err := this.info.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	// deserialize type_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.type_)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_queryBabyInstRes) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //inst
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize inst
	{
		err := this.inst.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_queryBabyInstRes) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize inst
	if mask.ReadBit() {
		err := this.inst.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_setNoTalkTime) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.t != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize t
	{
		if this.t != 0 {
			err := prpc.Write(buffer, this.t)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_setNoTalkTime) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize t
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.t)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_addNpc) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.npcList) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize npcList
	if len(this.npcList) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.npcList)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.npcList {
			err := prpc.Write(buffer, value)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_addNpc) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize npcList
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.npcList = make([]int32, size)
		for i, _ := range this.npcList {
			err := prpc.Read(buffer, &this.npcList[i])
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_delNpc) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.npcList) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize npcList
	if len(this.npcList) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.npcList)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.npcList {
			err := prpc.Write(buffer, value)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_delNpc) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize npcList
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.npcList = make([]int32, size)
		for i, _ := range this.npcList {
			err := prpc.Read(buffer, &this.npcList[i])
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_setTeamLeader) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.playerId != 0)
	mask.WriteBit(this.isLeader)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize playerId
	{
		if this.playerId != 0 {
			err := prpc.Write(buffer, this.playerId)
			if err != nil {
				return err
			}
		}
	}
	// serialize isLeader
	{
	}
	return nil
}
func (this *Server2Client_setTeamLeader) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize playerId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.playerId)
		if err != nil {
			return err
		}
	}
	// deserialize isLeader
	this.isLeader = mask.ReadBit()
	return nil
}
func (this *Server2Client_initBag) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.items) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize items
	if len(this.items) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.items)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.items {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_initBag) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize items
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.items = make([]COM_Item, size)
		for i, _ := range this.items {
			err := this.items[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_addBagItem) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //item
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize item
	{
		err := this.item.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_addBagItem) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize item
	if mask.ReadBit() {
		err := this.item.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_delBagItem) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.slot != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize slot
	{
		if this.slot != 0 {
			err := prpc.Write(buffer, this.slot)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_delBagItem) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize slot
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.slot)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_updateBagItem) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //item
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize item
	{
		err := this.item.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_updateBagItem) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize item
	if mask.ReadBit() {
		err := this.item.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_depositItemOK) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //item
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize item
	{
		err := this.item.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_depositItemOK) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize item
	if mask.ReadBit() {
		err := this.item.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_getoutItemOK) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.slot != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize slot
	{
		if this.slot != 0 {
			err := prpc.Write(buffer, this.slot)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_getoutItemOK) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize slot
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.slot)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_depositBabyOK) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //baby
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize baby
	{
		err := this.baby.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_depositBabyOK) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize baby
	if mask.ReadBit() {
		err := this.baby.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_getoutBabyOK) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.slot != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize slot
	{
		if this.slot != 0 {
			err := prpc.Write(buffer, this.slot)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_getoutBabyOK) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize slot
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.slot)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_sortItemStorageOK) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.items) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize items
	if len(this.items) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.items)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.items {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_sortItemStorageOK) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize items
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.items = make([]COM_Item, size)
		for i, _ := range this.items {
			err := this.items[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_sortBabyStorageOK) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.babys) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize babys
	if len(this.babys) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.babys)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.babys {
			err := prpc.Write(buffer, value)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_sortBabyStorageOK) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize babys
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.babys = make([]uint32, size)
		for i, _ := range this.babys {
			err := prpc.Read(buffer, &this.babys[i])
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_initItemStorage) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.gridNum != 0)
	mask.WriteBit(len(this.items) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize gridNum
	{
		if this.gridNum != 0 {
			err := prpc.Write(buffer, this.gridNum)
			if err != nil {
				return err
			}
		}
	}
	// serialize items
	if len(this.items) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.items)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.items {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_initItemStorage) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize gridNum
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.gridNum)
		if err != nil {
			return err
		}
	}
	// deserialize items
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.items = make([]COM_Item, size)
		for i, _ := range this.items {
			err := this.items[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_initBabyStorage) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.gridNum != 0)
	mask.WriteBit(len(this.babys) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize gridNum
	{
		if this.gridNum != 0 {
			err := prpc.Write(buffer, this.gridNum)
			if err != nil {
				return err
			}
		}
	}
	// serialize babys
	if len(this.babys) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.babys)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.babys {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_initBabyStorage) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize gridNum
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.gridNum)
		if err != nil {
			return err
		}
	}
	// deserialize babys
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.babys = make([]COM_BabyInst, size)
		for i, _ := range this.babys {
			err := this.babys[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_openStorageGrid) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.tp != 0)
	mask.WriteBit(this.gridNum != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize tp
	{
		if this.tp != 0 {
			err := prpc.Write(buffer, this.tp)
			if err != nil {
				return err
			}
		}
	}
	// serialize gridNum
	{
		if this.gridNum != 0 {
			err := prpc.Write(buffer, this.gridNum)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_openStorageGrid) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize tp
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.tp)
		if err != nil {
			return err
		}
	}
	// deserialize gridNum
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.gridNum)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_delStorageBabyOK) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.slot != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize slot
	{
		if this.slot != 0 {
			err := prpc.Write(buffer, this.slot)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_delStorageBabyOK) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize slot
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.slot)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_initPlayerEquips) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.equips) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize equips
	if len(this.equips) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.equips)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.equips {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_initPlayerEquips) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize equips
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.equips = make([]COM_Item, size)
		for i, _ := range this.equips {
			err := this.equips[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_wearEquipmentOk) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.target != 0)
	mask.WriteBit(true) //equip
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize target
	{
		if this.target != 0 {
			err := prpc.Write(buffer, this.target)
			if err != nil {
				return err
			}
		}
	}
	// serialize equip
	{
		err := this.equip.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_wearEquipmentOk) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize target
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.target)
		if err != nil {
			return err
		}
	}
	// deserialize equip
	if mask.ReadBit() {
		err := this.equip.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_scenePlayerWearEquipment) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.target != 0)
	mask.WriteBit(this.itemId != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize target
	{
		if this.target != 0 {
			err := prpc.Write(buffer, this.target)
			if err != nil {
				return err
			}
		}
	}
	// serialize itemId
	{
		if this.itemId != 0 {
			err := prpc.Write(buffer, this.itemId)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_scenePlayerWearEquipment) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize target
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.target)
		if err != nil {
			return err
		}
	}
	// deserialize itemId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.itemId)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_delEquipmentOk) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.target != 0)
	mask.WriteBit(this.itemInstId != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize target
	{
		if this.target != 0 {
			err := prpc.Write(buffer, this.target)
			if err != nil {
				return err
			}
		}
	}
	// serialize itemInstId
	{
		if this.itemInstId != 0 {
			err := prpc.Write(buffer, this.itemInstId)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_delEquipmentOk) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize target
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.target)
		if err != nil {
			return err
		}
	}
	// deserialize itemInstId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.itemInstId)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_scenePlayerDoffEquipment) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.target != 0)
	mask.WriteBit(this.itemId != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize target
	{
		if this.target != 0 {
			err := prpc.Write(buffer, this.target)
			if err != nil {
				return err
			}
		}
	}
	// serialize itemId
	{
		if this.itemId != 0 {
			err := prpc.Write(buffer, this.itemId)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_scenePlayerDoffEquipment) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize target
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.target)
		if err != nil {
			return err
		}
	}
	// deserialize itemId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.itemId)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_jointLobbyOk) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.infos) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize infos
	if len(this.infos) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.infos)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.infos {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_jointLobbyOk) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize infos
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.infos = make([]COM_SimpleTeamInfo, size)
		for i, _ := range this.infos {
			err := this.infos[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_syncDelLobbyTeam) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.teamId != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize teamId
	{
		if this.teamId != 0 {
			err := prpc.Write(buffer, this.teamId)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_syncDelLobbyTeam) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize teamId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.teamId)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_syncUpdateLobbyTeam) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //info
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize info
	{
		err := this.info.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_syncUpdateLobbyTeam) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize info
	if mask.ReadBit() {
		err := this.info.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_syncAddLobbyTeam) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //team
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize team
	{
		err := this.team.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_syncAddLobbyTeam) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize team
	if mask.ReadBit() {
		err := this.team.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_createTeamOk) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //team
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize team
	{
		err := this.team.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_createTeamOk) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize team
	if mask.ReadBit() {
		err := this.team.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_changeTeamOk) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //team
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize team
	{
		err := this.team.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_changeTeamOk) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize team
	if mask.ReadBit() {
		err := this.team.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_joinTeamOk) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //team
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize team
	{
		err := this.team.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_joinTeamOk) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize team
	if mask.ReadBit() {
		err := this.team.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_addTeamMember) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //info
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize info
	{
		err := this.info.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_addTeamMember) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize info
	if mask.ReadBit() {
		err := this.info.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_delTeamMember) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.instId != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize instId
	{
		if this.instId != 0 {
			err := prpc.Write(buffer, this.instId)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_delTeamMember) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize instId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.instId)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_changeTeamLeaderOk) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.uuid != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize uuid
	{
		if this.uuid != 0 {
			err := prpc.Write(buffer, this.uuid)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_changeTeamLeaderOk) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize uuid
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.uuid)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_exitTeamOk) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.iskick)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize iskick
	{
	}
	return nil
}
func (this *Server2Client_exitTeamOk) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize iskick
	this.iskick = mask.ReadBit()
	return nil
}
func (this *Server2Client_updateTeam) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //team
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize team
	{
		err := this.team.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_updateTeam) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize team
	if mask.ReadBit() {
		err := this.team.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_joinTeamRoomOK) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //team
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize team
	{
		err := this.team.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_joinTeamRoomOK) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize team
	if mask.ReadBit() {
		err := this.team.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_inviteJoinTeam) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.teamId != 0)
	mask.WriteBit(len(this.name) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize teamId
	{
		if this.teamId != 0 {
			err := prpc.Write(buffer, this.teamId)
			if err != nil {
				return err
			}
		}
	}
	// serialize name
	if len(this.name) != 0 {
		err := prpc.Write(buffer, this.name)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_inviteJoinTeam) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize teamId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.teamId)
		if err != nil {
			return err
		}
	}
	// deserialize name
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.name)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_syncTeamDirtyProp) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.guid != 0)
	mask.WriteBit(len(this.props) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize guid
	{
		if this.guid != 0 {
			err := prpc.Write(buffer, this.guid)
			if err != nil {
				return err
			}
		}
	}
	// serialize props
	if len(this.props) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.props)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.props {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_syncTeamDirtyProp) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize guid
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.guid)
		if err != nil {
			return err
		}
	}
	// deserialize props
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.props = make([]COM_PropValue, size)
		for i, _ := range this.props {
			err := this.props[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_leaveTeamOk) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.playerId != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize playerId
	{
		if this.playerId != 0 {
			err := prpc.Write(buffer, this.playerId)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_leaveTeamOk) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize playerId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.playerId)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_backTeamOK) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.playerId != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize playerId
	{
		if this.playerId != 0 {
			err := prpc.Write(buffer, this.playerId)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_backTeamOK) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize playerId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.playerId)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_refuseBackTeamOk) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.playerId != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize playerId
	{
		if this.playerId != 0 {
			err := prpc.Write(buffer, this.playerId)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_refuseBackTeamOk) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize playerId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.playerId)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_requestJoinTeamTranspond) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.reqName) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize reqName
	if len(this.reqName) != 0 {
		err := prpc.Write(buffer, this.reqName)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_requestJoinTeamTranspond) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize reqName
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.reqName)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_drawLotteryBoxRep) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.items) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize items
	if len(this.items) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.items)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.items {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_drawLotteryBoxRep) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize items
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.items = make([]COM_Item, size)
		for i, _ := range this.items {
			err := this.items[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_addEmployee) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //employee
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize employee
	{
		err := this.employee.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_addEmployee) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize employee
	if mask.ReadBit() {
		err := this.employee.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_battleEmployee) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.empId != 0)
	mask.WriteBit(this.group != 0)
	mask.WriteBit(this.forbattle)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize empId
	{
		if this.empId != 0 {
			err := prpc.Write(buffer, this.empId)
			if err != nil {
				return err
			}
		}
	}
	// serialize group
	{
		if this.group != 0 {
			err := prpc.Write(buffer, this.group)
			if err != nil {
				return err
			}
		}
	}
	// serialize forbattle
	{
	}
	return nil
}
func (this *Server2Client_battleEmployee) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize empId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.empId)
		if err != nil {
			return err
		}
	}
	// deserialize group
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.group)
		if err != nil {
			return err
		}
	}
	// deserialize forbattle
	this.forbattle = mask.ReadBit()
	return nil
}
func (this *Server2Client_changeEmpBattleGroupOK) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.group != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize group
	{
		if this.group != 0 {
			err := prpc.Write(buffer, this.group)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_changeEmpBattleGroupOK) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize group
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.group)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_evolveOK) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.guid != 0)
	mask.WriteBit(this.qc != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize guid
	{
		if this.guid != 0 {
			err := prpc.Write(buffer, this.guid)
			if err != nil {
				return err
			}
		}
	}
	// serialize qc
	{
		if this.qc != 0 {
			err := prpc.Write(buffer, this.qc)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_evolveOK) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize guid
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.guid)
		if err != nil {
			return err
		}
	}
	// deserialize qc
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.qc)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_upStarOK) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.guid != 0)
	mask.WriteBit(this.star != 0)
	mask.WriteBit(true) //sk
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize guid
	{
		if this.guid != 0 {
			err := prpc.Write(buffer, this.guid)
			if err != nil {
				return err
			}
		}
	}
	// serialize star
	{
		if this.star != 0 {
			err := prpc.Write(buffer, this.star)
			if err != nil {
				return err
			}
		}
	}
	// serialize sk
	{
		err := this.sk.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_upStarOK) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize guid
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.guid)
		if err != nil {
			return err
		}
	}
	// deserialize star
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.star)
		if err != nil {
			return err
		}
	}
	// deserialize sk
	if mask.ReadBit() {
		err := this.sk.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_delEmployeeOK) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.instids) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize instids
	if len(this.instids) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.instids)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.instids {
			err := prpc.Write(buffer, value)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_delEmployeeOK) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize instids
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.instids = make([]uint32, size)
		for i, _ := range this.instids {
			err := prpc.Read(buffer, &this.instids[i])
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_sycnEmployeeSoul) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.guid != 0)
	mask.WriteBit(this.soulNum != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize guid
	{
		if this.guid != 0 {
			err := prpc.Write(buffer, this.guid)
			if err != nil {
				return err
			}
		}
	}
	// serialize soulNum
	{
		if this.soulNum != 0 {
			err := prpc.Write(buffer, this.soulNum)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_sycnEmployeeSoul) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize guid
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.guid)
		if err != nil {
			return err
		}
	}
	// deserialize soulNum
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.soulNum)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_initQuest) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.qlist) != 0)
	mask.WriteBit(len(this.clist) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize qlist
	if len(this.qlist) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.qlist)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.qlist {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// serialize clist
	if len(this.clist) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.clist)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.clist {
			err := prpc.Write(buffer, value)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_initQuest) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize qlist
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.qlist = make([]COM_QuestInst, size)
		for i, _ := range this.qlist {
			err := this.qlist[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// deserialize clist
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.clist = make([]int32, size)
		for i, _ := range this.clist {
			err := prpc.Read(buffer, &this.clist[i])
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_acceptQuestOk) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //inst
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize inst
	{
		err := this.inst.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_acceptQuestOk) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize inst
	if mask.ReadBit() {
		err := this.inst.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_submitQuestOk) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.questId != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize questId
	{
		if this.questId != 0 {
			err := prpc.Write(buffer, this.questId)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_submitQuestOk) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize questId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.questId)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_giveupQuestOk) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.questId != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize questId
	{
		if this.questId != 0 {
			err := prpc.Write(buffer, this.questId)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_giveupQuestOk) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize questId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.questId)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_updateQuestInst) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //inst
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize inst
	{
		err := this.inst.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_updateQuestInst) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize inst
	if mask.ReadBit() {
		err := this.inst.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_requestContactInfoOk) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //contact
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize contact
	{
		err := this.contact.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_requestContactInfoOk) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize contact
	if mask.ReadBit() {
		err := this.contact.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_addFriendOK) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //inst
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize inst
	{
		err := this.inst.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_addFriendOK) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize inst
	if mask.ReadBit() {
		err := this.inst.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_delFriendOK) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.instId != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize instId
	{
		if this.instId != 0 {
			err := prpc.Write(buffer, this.instId)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_delFriendOK) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize instId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.instId)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_addBlacklistOK) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //inst
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize inst
	{
		err := this.inst.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_addBlacklistOK) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize inst
	if mask.ReadBit() {
		err := this.inst.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_delBlacklistOK) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.instId != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize instId
	{
		if this.instId != 0 {
			err := prpc.Write(buffer, this.instId)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_delBlacklistOK) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize instId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.instId)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_referrFriendOK) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.insts) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize insts
	if len(this.insts) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.insts)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.insts {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_referrFriendOK) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize insts
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.insts = make([]COM_ContactInfo, size)
		for i, _ := range this.insts {
			err := this.insts[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_requestFriendListOK) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.insts) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize insts
	if len(this.insts) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.insts)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.insts {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_requestFriendListOK) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize insts
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.insts = make([]COM_ContactInfo, size)
		for i, _ := range this.insts {
			err := this.insts[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_lotteryOk) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.lotteryId != 0)
	mask.WriteBit(len(this.dropItem) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize lotteryId
	{
		if this.lotteryId != 0 {
			err := prpc.Write(buffer, this.lotteryId)
			if err != nil {
				return err
			}
		}
	}
	// serialize dropItem
	if len(this.dropItem) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.dropItem)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.dropItem {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_lotteryOk) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize lotteryId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.lotteryId)
		if err != nil {
			return err
		}
	}
	// deserialize dropItem
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.dropItem = make([]COM_DropItem, size)
		for i, _ := range this.dropItem {
			err := this.dropItem[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_openGatherOK) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //gather
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize gather
	{
		err := this.gather.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_openGatherOK) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize gather
	if mask.ReadBit() {
		err := this.gather.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_miningOk) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.items) != 0)
	mask.WriteBit(true) //gather
	mask.WriteBit(this.gatherNum != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize items
	if len(this.items) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.items)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.items {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// serialize gather
	{
		err := this.gather.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	// serialize gatherNum
	{
		if this.gatherNum != 0 {
			err := prpc.Write(buffer, this.gatherNum)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_miningOk) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize items
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.items = make([]COM_DropItem, size)
		for i, _ := range this.items {
			err := this.items[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// deserialize gather
	if mask.ReadBit() {
		err := this.gather.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	// deserialize gatherNum
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.gatherNum)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_openCompound) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.compoundId != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize compoundId
	{
		if this.compoundId != 0 {
			err := prpc.Write(buffer, this.compoundId)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_openCompound) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize compoundId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.compoundId)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_compoundItemOk) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //item
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize item
	{
		err := this.item.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_compoundItemOk) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize item
	if mask.ReadBit() {
		err := this.item.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_openBagGridOk) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.num != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize num
	{
		if this.num != 0 {
			err := prpc.Write(buffer, this.num)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_openBagGridOk) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize num
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.num)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_requestChallengeOK) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.isOK)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize isOK
	{
	}
	return nil
}
func (this *Server2Client_requestChallengeOK) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize isOK
	this.isOK = mask.ReadBit()
	return nil
}
func (this *Server2Client_requestMySelfJJCDataOK) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //info
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize info
	{
		err := this.info.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_requestMySelfJJCDataOK) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize info
	if mask.ReadBit() {
		err := this.info.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_requestRivalOK) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.infos) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize infos
	if len(this.infos) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.infos)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.infos {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_requestRivalOK) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize infos
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.infos = make([]COM_EndlessStair, size)
		for i, _ := range this.infos {
			err := this.infos[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_checkMsgOK) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //inst
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize inst
	{
		err := this.inst.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_checkMsgOK) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize inst
	if mask.ReadBit() {
		err := this.inst.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_requestMyAllbattleMsgOK) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.infos) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize infos
	if len(this.infos) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.infos)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.infos {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_requestMyAllbattleMsgOK) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize infos
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.infos = make([]COM_JJCBattleMsg, size)
		for i, _ := range this.infos {
			err := this.infos[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_myBattleMsgOK) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //info
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize info
	{
		err := this.info.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_myBattleMsgOK) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize info
	if mask.ReadBit() {
		err := this.info.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_requestJJCRankOK) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.myRank != 0)
	mask.WriteBit(len(this.infos) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize myRank
	{
		if this.myRank != 0 {
			err := prpc.Write(buffer, this.myRank)
			if err != nil {
				return err
			}
		}
	}
	// serialize infos
	if len(this.infos) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.infos)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.infos {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_requestJJCRankOK) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize myRank
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.myRank)
		if err != nil {
			return err
		}
	}
	// deserialize infos
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.infos = make([]COM_EndlessStair, size)
		for i, _ := range this.infos {
			err := this.infos[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_requestLevelRankOK) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.myRank != 0)
	mask.WriteBit(len(this.infos) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize myRank
	{
		if this.myRank != 0 {
			err := prpc.Write(buffer, this.myRank)
			if err != nil {
				return err
			}
		}
	}
	// serialize infos
	if len(this.infos) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.infos)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.infos {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_requestLevelRankOK) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize myRank
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.myRank)
		if err != nil {
			return err
		}
	}
	// deserialize infos
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.infos = make([]COM_ContactInfo, size)
		for i, _ := range this.infos {
			err := this.infos[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_requestBabyRankOK) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.myRank != 0)
	mask.WriteBit(len(this.infos) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize myRank
	{
		if this.myRank != 0 {
			err := prpc.Write(buffer, this.myRank)
			if err != nil {
				return err
			}
		}
	}
	// serialize infos
	if len(this.infos) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.infos)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.infos {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_requestBabyRankOK) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize myRank
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.myRank)
		if err != nil {
			return err
		}
	}
	// deserialize infos
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.infos = make([]COM_BabyRankData, size)
		for i, _ := range this.infos {
			err := this.infos[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_requestEmpRankOK) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.myRank != 0)
	mask.WriteBit(len(this.infos) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize myRank
	{
		if this.myRank != 0 {
			err := prpc.Write(buffer, this.myRank)
			if err != nil {
				return err
			}
		}
	}
	// serialize infos
	if len(this.infos) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.infos)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.infos {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_requestEmpRankOK) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize myRank
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.myRank)
		if err != nil {
			return err
		}
	}
	// deserialize infos
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.infos = make([]COM_EmployeeRankData, size)
		for i, _ := range this.infos {
			err := this.infos[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_requestPlayerFFRankOK) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.myRank != 0)
	mask.WriteBit(len(this.infos) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize myRank
	{
		if this.myRank != 0 {
			err := prpc.Write(buffer, this.myRank)
			if err != nil {
				return err
			}
		}
	}
	// serialize infos
	if len(this.infos) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.infos)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.infos {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_requestPlayerFFRankOK) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize myRank
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.myRank)
		if err != nil {
			return err
		}
	}
	// deserialize infos
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.infos = make([]COM_ContactInfo, size)
		for i, _ := range this.infos {
			err := this.infos[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_queryOnlinePlayerOK) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.isOnline)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize isOnline
	{
	}
	return nil
}
func (this *Server2Client_queryOnlinePlayerOK) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize isOnline
	this.isOnline = mask.ReadBit()
	return nil
}
func (this *Server2Client_queryPlayerOK) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //inst
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize inst
	{
		err := this.inst.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_queryPlayerOK) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize inst
	if mask.ReadBit() {
		err := this.inst.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_queryBabyOK) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //inst
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize inst
	{
		err := this.inst.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_queryBabyOK) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize inst
	if mask.ReadBit() {
		err := this.inst.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_queryEmployeeOK) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //inst
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize inst
	{
		err := this.inst.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_queryEmployeeOK) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize inst
	if mask.ReadBit() {
		err := this.inst.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_initGuide) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.guideMask != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize guideMask
	{
		if this.guideMask != 0 {
			err := prpc.Write(buffer, this.guideMask)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_initGuide) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize guideMask
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.guideMask)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_buyShopItemOk) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.id != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize id
	{
		if this.id != 0 {
			err := prpc.Write(buffer, this.id)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_buyShopItemOk) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize id
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.id)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_addPlayerTitle) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.title != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize title
	{
		if this.title != 0 {
			err := prpc.Write(buffer, this.title)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_addPlayerTitle) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize title
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.title)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_delPlayerTitle) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.title != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize title
	{
		if this.title != 0 {
			err := prpc.Write(buffer, this.title)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_delPlayerTitle) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize title
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.title)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_requestOpenBuyBox) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.greenTime != 0)
	mask.WriteBit(this.blueTime != 0)
	mask.WriteBit(this.greenFreeNum != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize greenTime
	{
		if this.greenTime != 0 {
			err := prpc.Write(buffer, this.greenTime)
			if err != nil {
				return err
			}
		}
	}
	// serialize blueTime
	{
		if this.blueTime != 0 {
			err := prpc.Write(buffer, this.blueTime)
			if err != nil {
				return err
			}
		}
	}
	// serialize greenFreeNum
	{
		if this.greenFreeNum != 0 {
			err := prpc.Write(buffer, this.greenFreeNum)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_requestOpenBuyBox) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize greenTime
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.greenTime)
		if err != nil {
			return err
		}
	}
	// deserialize blueTime
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.blueTime)
		if err != nil {
			return err
		}
	}
	// deserialize greenFreeNum
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.greenFreeNum)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_updateAchievementinfo) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //achs
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize achs
	{
		err := this.achs.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_updateAchievementinfo) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize achs
	if mask.ReadBit() {
		err := this.achs.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_syncOpenSystemFlag) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.flag != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize flag
	{
		if this.flag != 0 {
			err := prpc.Write(buffer, this.flag)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_syncOpenSystemFlag) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize flag
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.flag)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_requestActivityRewardOK) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.ar != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize ar
	{
		if this.ar != 0 {
			err := prpc.Write(buffer, this.ar)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_requestActivityRewardOK) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize ar
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.ar)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_syncActivity) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //table
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize table
	{
		err := this.table.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_syncActivity) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize table
	if mask.ReadBit() {
		err := this.table.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_updateActivityStatus) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.type_ != 0)
	mask.WriteBit(this.open)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize type_
	{
		if this.type_ != 0 {
			err := prpc.Write(buffer, this.type_)
			if err != nil {
				return err
			}
		}
	}
	// serialize open
	{
	}
	return nil
}
func (this *Server2Client_updateActivityStatus) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize type_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.type_)
		if err != nil {
			return err
		}
	}
	// deserialize open
	this.open = mask.ReadBit()
	return nil
}
func (this *Server2Client_updateActivityCounter) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.type_ != 0)
	mask.WriteBit(this.counter != 0)
	mask.WriteBit(this.reward != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize type_
	{
		if this.type_ != 0 {
			err := prpc.Write(buffer, this.type_)
			if err != nil {
				return err
			}
		}
	}
	// serialize counter
	{
		if this.counter != 0 {
			err := prpc.Write(buffer, this.counter)
			if err != nil {
				return err
			}
		}
	}
	// serialize reward
	{
		if this.reward != 0 {
			err := prpc.Write(buffer, this.reward)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_updateActivityCounter) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize type_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.type_)
		if err != nil {
			return err
		}
	}
	// deserialize counter
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.counter)
		if err != nil {
			return err
		}
	}
	// deserialize reward
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.reward)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_syncExam) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //exam
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize exam
	{
		err := this.exam.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_syncExam) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize exam
	if mask.ReadBit() {
		err := this.exam.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_syncExamAnswer) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //answer
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize answer
	{
		err := this.answer.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_syncExamAnswer) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize answer
	if mask.ReadBit() {
		err := this.answer.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_petActivityNoNum) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.name) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize name
	if len(this.name) != 0 {
		err := prpc.Write(buffer, this.name)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_petActivityNoNum) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize name
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.name)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_syncHundredInfo) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //hb
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize hb
	{
		err := this.hb.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_syncHundredInfo) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize hb
	if mask.ReadBit() {
		err := this.hb.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_initSignUp) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.info) != 0)
	mask.WriteBit(this.process != 0)
	mask.WriteBit(this.sign7)
	mask.WriteBit(this.sign14)
	mask.WriteBit(this.sign28)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize info
	if len(this.info) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.info)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.info {
			err := prpc.Write(buffer, value)
			if err != nil {
				return err
			}
		}
	}
	// serialize process
	{
		if this.process != 0 {
			err := prpc.Write(buffer, this.process)
			if err != nil {
				return err
			}
		}
	}
	// serialize sign7
	{
	}
	// serialize sign14
	{
	}
	// serialize sign28
	{
	}
	return nil
}
func (this *Server2Client_initSignUp) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize info
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.info = make([]int32, size)
		for i, _ := range this.info {
			err := prpc.Read(buffer, &this.info[i])
			if err != nil {
				return err
			}
		}
	}
	// deserialize process
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.process)
		if err != nil {
			return err
		}
	}
	// deserialize sign7
	this.sign7 = mask.ReadBit()
	// deserialize sign14
	this.sign14 = mask.ReadBit()
	// deserialize sign28
	this.sign28 = mask.ReadBit()
	return nil
}
func (this *Server2Client_signUp) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.flag)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize flag
	{
	}
	return nil
}
func (this *Server2Client_signUp) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize flag
	this.flag = mask.ReadBit()
	return nil
}
func (this *Server2Client_sycnDoubleExpTime) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.isFlag)
	mask.WriteBit(this.times != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize isFlag
	{
	}
	// serialize times
	{
		if this.times != 0 {
			err := prpc.Write(buffer, this.times)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_sycnDoubleExpTime) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize isFlag
	this.isFlag = mask.ReadBit()
	// deserialize times
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.times)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_sycnStates) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.states) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize states
	if len(this.states) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.states)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.states {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_sycnStates) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize states
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.states = make([]COM_State, size)
		for i, _ := range this.states {
			err := this.states[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_requestpvprankOK) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.infos) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize infos
	if len(this.infos) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.infos)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.infos {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_requestpvprankOK) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize infos
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.infos = make([]COM_ContactInfo, size)
		for i, _ := range this.infos {
			err := this.infos[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_stopMatchingOK) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.times != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize times
	{
		if this.times != 0 {
			err := prpc.Write(buffer, this.times)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_stopMatchingOK) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize times
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.times)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_updatePvpJJCinfo) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //info
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize info
	{
		err := this.info.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_updatePvpJJCinfo) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize info
	if mask.ReadBit() {
		err := this.info.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_syncEnemyPvpJJCPlayerInfo) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //info
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize info
	{
		err := this.info.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_syncEnemyPvpJJCPlayerInfo) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize info
	if mask.ReadBit() {
		err := this.info.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_syncEnemyPvpJJCTeamInfo) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.infos) != 0)
	mask.WriteBit(this.teamID_ != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize infos
	if len(this.infos) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.infos)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.infos {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// serialize teamID_
	{
		if this.teamID_ != 0 {
			err := prpc.Write(buffer, this.teamID_)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_syncEnemyPvpJJCTeamInfo) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize infos
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.infos = make([]COM_SimpleInformation, size)
		for i, _ := range this.infos {
			err := this.infos[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// deserialize teamID_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.teamID_)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_syncWarriorEnemyTeamInfo) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.infos) != 0)
	mask.WriteBit(this.teamID_ != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize infos
	if len(this.infos) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.infos)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.infos {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// serialize teamID_
	{
		if this.teamID_ != 0 {
			err := prpc.Write(buffer, this.teamID_)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_syncWarriorEnemyTeamInfo) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize infos
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.infos = make([]COM_SimpleInformation, size)
		for i, _ := range this.infos {
			err := this.infos[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// deserialize teamID_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.teamID_)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_appendMail) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.mails) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize mails
	if len(this.mails) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.mails)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.mails {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_appendMail) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize mails
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.mails = make([]COM_Mail, size)
		for i, _ := range this.mails {
			err := this.mails[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_delMail) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.mailId != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize mailId
	{
		if this.mailId != 0 {
			err := prpc.Write(buffer, this.mailId)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_delMail) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize mailId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.mailId)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_updateMailOk) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //mail
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize mail
	{
		err := this.mail.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_updateMailOk) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize mail
	if mask.ReadBit() {
		err := this.mail.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_boardcastNotice) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.content) != 0)
	mask.WriteBit(this.isGm)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize content
	if len(this.content) != 0 {
		err := prpc.Write(buffer, this.content)
		if err != nil {
			return err
		}
	}
	// serialize isGm
	{
	}
	return nil
}
func (this *Server2Client_boardcastNotice) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize content
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.content)
		if err != nil {
			return err
		}
	}
	// deserialize isGm
	this.isGm = mask.ReadBit()
	return nil
}
func (this *Server2Client_leaveGuildOk) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.who) != 0)
	mask.WriteBit(this.isKick)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize who
	if len(this.who) != 0 {
		err := prpc.Write(buffer, this.who)
		if err != nil {
			return err
		}
	}
	// serialize isKick
	{
	}
	return nil
}
func (this *Server2Client_leaveGuildOk) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize who
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.who)
		if err != nil {
			return err
		}
	}
	// deserialize isKick
	this.isKick = mask.ReadBit()
	return nil
}
func (this *Server2Client_initGuildData) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //guild
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize guild
	{
		err := this.guild.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_initGuildData) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize guild
	if mask.ReadBit() {
		err := this.guild.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_initGuildMemberList) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.member) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize member
	if len(this.member) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.member)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.member {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_initGuildMemberList) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize member
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.member = make([]COM_GuildMember, size)
		for i, _ := range this.member {
			err := this.member[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_modifyGuildMemberList) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //member
	mask.WriteBit(this.flag != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize member
	{
		err := this.member.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	// serialize flag
	{
		if this.flag != 0 {
			err := prpc.Write(buffer, this.flag)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_modifyGuildMemberList) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize member
	if mask.ReadBit() {
		err := this.member.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	// deserialize flag
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.flag)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_modifyGuildList) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //data
	mask.WriteBit(this.flag != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize data
	{
		err := this.data.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	// serialize flag
	{
		if this.flag != 0 {
			err := prpc.Write(buffer, this.flag)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_modifyGuildList) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize data
	if mask.ReadBit() {
		err := this.data.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	// deserialize flag
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.flag)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_queryGuildListResult) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.page != 0)
	mask.WriteBit(this.pageNum != 0)
	mask.WriteBit(len(this.guildList) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize page
	{
		if this.page != 0 {
			err := prpc.Write(buffer, this.page)
			if err != nil {
				return err
			}
		}
	}
	// serialize pageNum
	{
		if this.pageNum != 0 {
			err := prpc.Write(buffer, this.pageNum)
			if err != nil {
				return err
			}
		}
	}
	// serialize guildList
	if len(this.guildList) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.guildList)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.guildList {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_queryGuildListResult) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize page
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.page)
		if err != nil {
			return err
		}
	}
	// deserialize pageNum
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.pageNum)
		if err != nil {
			return err
		}
	}
	// deserialize guildList
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.guildList = make([]COM_GuildViewerData, size)
		for i, _ := range this.guildList {
			err := this.guildList[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_inviteGuild) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.sendName) != 0)
	mask.WriteBit(len(this.guildName) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize sendName
	if len(this.sendName) != 0 {
		err := prpc.Write(buffer, this.sendName)
		if err != nil {
			return err
		}
	}
	// serialize guildName
	if len(this.guildName) != 0 {
		err := prpc.Write(buffer, this.guildName)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_inviteGuild) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize sendName
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.sendName)
		if err != nil {
			return err
		}
	}
	// deserialize guildName
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.guildName)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_updateGuildShopItems) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.items) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize items
	if len(this.items) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.items)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.items {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_updateGuildShopItems) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize items
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.items = make([]COM_GuildShopItem, size)
		for i, _ := range this.items {
			err := this.items[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_updateGuildBuilding) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.type_ != 0)
	mask.WriteBit(true) //building
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize type_
	{
		if this.type_ != 0 {
			err := prpc.Write(buffer, this.type_)
			if err != nil {
				return err
			}
		}
	}
	// serialize building
	{
		err := this.building.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_updateGuildBuilding) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize type_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.type_)
		if err != nil {
			return err
		}
	}
	// deserialize building
	if mask.ReadBit() {
		err := this.building.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_updateGuildMyMember) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //member
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize member
	{
		err := this.member.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_updateGuildMyMember) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize member
	if mask.ReadBit() {
		err := this.member.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_levelupGuildSkillOk) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //skInst
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize skInst
	{
		err := this.skInst.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_levelupGuildSkillOk) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize skInst
	if mask.ReadBit() {
		err := this.skInst.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_presentGuildItemOk) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.val != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize val
	{
		if this.val != 0 {
			err := prpc.Write(buffer, this.val)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_presentGuildItemOk) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize val
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.val)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_progenitusAddExpOk) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //mInst
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize mInst
	{
		err := this.mInst.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_progenitusAddExpOk) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize mInst
	if mask.ReadBit() {
		err := this.mInst.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_setProgenitusPositionOk) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.positions) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize positions
	if len(this.positions) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.positions)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.positions {
			err := prpc.Write(buffer, value)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_setProgenitusPositionOk) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize positions
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.positions = make([]int32, size)
		for i, _ := range this.positions {
			err := prpc.Read(buffer, &this.positions[i])
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_updateGuildFundz) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.val != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize val
	{
		if this.val != 0 {
			err := prpc.Write(buffer, this.val)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_updateGuildFundz) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize val
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.val)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_updateGuildMemberContribution) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.val != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize val
	{
		if this.val != 0 {
			err := prpc.Write(buffer, this.val)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_updateGuildMemberContribution) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize val
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.val)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_openGuildBattle) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.otherName) != 0)
	mask.WriteBit(this.playerNum != 0)
	mask.WriteBit(this.level != 0)
	mask.WriteBit(this.isLeft)
	mask.WriteBit(this.lstime != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize otherName
	if len(this.otherName) != 0 {
		err := prpc.Write(buffer, this.otherName)
		if err != nil {
			return err
		}
	}
	// serialize playerNum
	{
		if this.playerNum != 0 {
			err := prpc.Write(buffer, this.playerNum)
			if err != nil {
				return err
			}
		}
	}
	// serialize level
	{
		if this.level != 0 {
			err := prpc.Write(buffer, this.level)
			if err != nil {
				return err
			}
		}
	}
	// serialize isLeft
	{
	}
	// serialize lstime
	{
		if this.lstime != 0 {
			err := prpc.Write(buffer, this.lstime)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_openGuildBattle) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize otherName
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.otherName)
		if err != nil {
			return err
		}
	}
	// deserialize playerNum
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.playerNum)
		if err != nil {
			return err
		}
	}
	// deserialize level
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.level)
		if err != nil {
			return err
		}
	}
	// deserialize isLeft
	this.isLeft = mask.ReadBit()
	// deserialize lstime
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.lstime)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_startGuildBattle) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.otherName) != 0)
	mask.WriteBit(this.otherCon != 0)
	mask.WriteBit(this.selfCon != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize otherName
	if len(this.otherName) != 0 {
		err := prpc.Write(buffer, this.otherName)
		if err != nil {
			return err
		}
	}
	// serialize otherCon
	{
		if this.otherCon != 0 {
			err := prpc.Write(buffer, this.otherCon)
			if err != nil {
				return err
			}
		}
	}
	// serialize selfCon
	{
		if this.selfCon != 0 {
			err := prpc.Write(buffer, this.selfCon)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_startGuildBattle) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize otherName
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.otherName)
		if err != nil {
			return err
		}
	}
	// deserialize otherCon
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.otherCon)
		if err != nil {
			return err
		}
	}
	// deserialize selfCon
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.selfCon)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_closeGuildBattle) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.isWinner)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize isWinner
	{
	}
	return nil
}
func (this *Server2Client_closeGuildBattle) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize isWinner
	this.isWinner = mask.ReadBit()
	return nil
}
func (this *Server2Client_syncGuildBattleWinCount) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.myWin != 0)
	mask.WriteBit(this.otherWin != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize myWin
	{
		if this.myWin != 0 {
			err := prpc.Write(buffer, this.myWin)
			if err != nil {
				return err
			}
		}
	}
	// serialize otherWin
	{
		if this.otherWin != 0 {
			err := prpc.Write(buffer, this.otherWin)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_syncGuildBattleWinCount) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize myWin
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.myWin)
		if err != nil {
			return err
		}
	}
	// deserialize otherWin
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.otherWin)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_initMySelling) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.items) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize items
	if len(this.items) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.items)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.items {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_initMySelling) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize items
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.items = make([]COM_SellItem, size)
		for i, _ := range this.items {
			err := this.items[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_initMySelled) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.items) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize items
	if len(this.items) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.items)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.items {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_initMySelled) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize items
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.items = make([]COM_SelledItem, size)
		for i, _ := range this.items {
			err := this.items[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_fetchSellingOk) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.items) != 0)
	mask.WriteBit(this.total != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize items
	if len(this.items) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.items)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.items {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// serialize total
	{
		if this.total != 0 {
			err := prpc.Write(buffer, this.total)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_fetchSellingOk) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize items
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.items = make([]COM_SellItem, size)
		for i, _ := range this.items {
			err := this.items[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// deserialize total
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.total)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_fetchSellingOk2) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.items) != 0)
	mask.WriteBit(this.total != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize items
	if len(this.items) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.items)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.items {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// serialize total
	{
		if this.total != 0 {
			err := prpc.Write(buffer, this.total)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_fetchSellingOk2) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize items
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.items = make([]COM_SellItem, size)
		for i, _ := range this.items {
			err := this.items[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// deserialize total
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.total)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_sellingOk) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //item
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize item
	{
		err := this.item.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_sellingOk) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize item
	if mask.ReadBit() {
		err := this.item.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_selledOk) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //item
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize item
	{
		err := this.item.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_selledOk) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize item
	if mask.ReadBit() {
		err := this.item.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_unsellingOk) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.sellid != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize sellid
	{
		if this.sellid != 0 {
			err := prpc.Write(buffer, this.sellid)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_unsellingOk) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize sellid
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.sellid)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_insertState) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //st
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize st
	{
		err := this.st.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_insertState) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize st
	if mask.ReadBit() {
		err := this.st.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_updattState) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //st
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize st
	{
		err := this.st.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_updattState) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize st
	if mask.ReadBit() {
		err := this.st.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_removeState) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.stid != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize stid
	{
		if this.stid != 0 {
			err := prpc.Write(buffer, this.stid)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_removeState) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize stid
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.stid)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_requestFixItemOk) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //item
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize item
	{
		err := this.item.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_requestFixItemOk) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize item
	if mask.ReadBit() {
		err := this.item.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_updateMagicItem) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.level != 0)
	mask.WriteBit(this.exp != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize level
	{
		if this.level != 0 {
			err := prpc.Write(buffer, this.level)
			if err != nil {
				return err
			}
		}
	}
	// serialize exp
	{
		if this.exp != 0 {
			err := prpc.Write(buffer, this.exp)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_updateMagicItem) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize level
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.level)
		if err != nil {
			return err
		}
	}
	// deserialize exp
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.exp)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_changeMagicJobOk) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.job != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize job
	{
		if this.job != 0 {
			err := prpc.Write(buffer, this.job)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_changeMagicJobOk) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize job
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.job)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_magicItemTupoOk) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.level != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize level
	{
		if this.level != 0 {
			err := prpc.Write(buffer, this.level)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_magicItemTupoOk) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize level
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.level)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_zhuanpanOK) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.pond) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize pond
	if len(this.pond) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.pond)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.pond {
			err := prpc.Write(buffer, value)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_zhuanpanOK) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize pond
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.pond = make([]uint32, size)
		for i, _ := range this.pond {
			err := prpc.Read(buffer, &this.pond[i])
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_updateZhuanpanNotice) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //zhuanp
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize zhuanp
	{
		err := this.zhuanp.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_updateZhuanpanNotice) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize zhuanp
	if mask.ReadBit() {
		err := this.zhuanp.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_sycnZhuanpanData) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //data
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize data
	{
		err := this.data.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_sycnZhuanpanData) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize data
	if mask.ReadBit() {
		err := this.data.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_copynonum) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.name) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize name
	if len(this.name) != 0 {
		err := prpc.Write(buffer, this.name)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_copynonum) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize name
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.name)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_sceneFilterOk) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.sfType) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize sfType
	if len(this.sfType) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.sfType)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.sfType {
			err := prpc.Write(buffer, value)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_sceneFilterOk) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize sfType
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.sfType = make([]int, size)
		for i, _ := range this.sfType {
			err := prpc.Read(buffer, &this.sfType[i])
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_shareWishOK) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //wish
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize wish
	{
		err := this.wish.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_shareWishOK) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize wish
	if mask.ReadBit() {
		err := this.wish.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_requestOnlineTimeRewardOK) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.index != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize index
	{
		if this.index != 0 {
			err := prpc.Write(buffer, this.index)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_requestOnlineTimeRewardOK) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize index
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.index)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_sycnVipflag) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.flag)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize flag
	{
	}
	return nil
}
func (this *Server2Client_sycnVipflag) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize flag
	this.flag = mask.ReadBit()
	return nil
}
func (this *Server2Client_buyFundOK) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.flag)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize flag
	{
	}
	return nil
}
func (this *Server2Client_buyFundOK) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize flag
	this.flag = mask.ReadBit()
	return nil
}
func (this *Server2Client_requestFundRewardOK) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.level != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize level
	{
		if this.level != 0 {
			err := prpc.Write(buffer, this.level)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_requestFundRewardOK) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize level
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.level)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_firstRechargeOK) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.isFlag)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize isFlag
	{
	}
	return nil
}
func (this *Server2Client_firstRechargeOK) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize isFlag
	this.isFlag = mask.ReadBit()
	return nil
}
func (this *Server2Client_firstRechargeGiftOK) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.isFlag)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize isFlag
	{
	}
	return nil
}
func (this *Server2Client_firstRechargeGiftOK) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize isFlag
	this.isFlag = mask.ReadBit()
	return nil
}
func (this *Server2Client_agencyActivity) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.type_ != 0)
	mask.WriteBit(this.isFlag)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize type_
	{
		if this.type_ != 0 {
			err := prpc.Write(buffer, this.type_)
			if err != nil {
				return err
			}
		}
	}
	// serialize isFlag
	{
	}
	return nil
}
func (this *Server2Client_agencyActivity) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize type_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.type_)
		if err != nil {
			return err
		}
	}
	// deserialize isFlag
	this.isFlag = mask.ReadBit()
	return nil
}
func (this *Server2Client_updateFestival) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //festival
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize festival
	{
		err := this.festival.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_updateFestival) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize festival
	if mask.ReadBit() {
		err := this.festival.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_updateSelfRecharge) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //val
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize val
	{
		err := this.val.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_updateSelfRecharge) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize val
	if mask.ReadBit() {
		err := this.val.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_updateSysRecharge) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //val
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize val
	{
		err := this.val.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_updateSysRecharge) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize val
	if mask.ReadBit() {
		err := this.val.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_updateSelfDiscountStore) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //val
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize val
	{
		err := this.val.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_updateSelfDiscountStore) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize val
	if mask.ReadBit() {
		err := this.val.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_updateSysDiscountStore) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //val
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize val
	{
		err := this.val.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_updateSysDiscountStore) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize val
	if mask.ReadBit() {
		err := this.val.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_updateSelfOnceRecharge) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //val
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize val
	{
		err := this.val.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_updateSelfOnceRecharge) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize val
	if mask.ReadBit() {
		err := this.val.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_updateSysOnceRecharge) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //val
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize val
	{
		err := this.val.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_updateSysOnceRecharge) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize val
	if mask.ReadBit() {
		err := this.val.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_openCardOK) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //data
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize data
	{
		err := this.data.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_openCardOK) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize data
	if mask.ReadBit() {
		err := this.data.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_sycnHotRole) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //data
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize data
	{
		err := this.data.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_sycnHotRole) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize data
	if mask.ReadBit() {
		err := this.data.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_hotRoleBuyOk) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.buyNum != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize buyNum
	{
		if this.buyNum != 0 {
			err := prpc.Write(buffer, this.buyNum)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_hotRoleBuyOk) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize buyNum
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.buyNum)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_updateSevenday) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //data
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize data
	{
		err := this.data.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_updateSevenday) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize data
	if mask.ReadBit() {
		err := this.data.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_updateEmployeeActivity) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //data
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize data
	{
		err := this.data.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_updateEmployeeActivity) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize data
	if mask.ReadBit() {
		err := this.data.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_updateMinGiftActivity) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //data
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize data
	{
		err := this.data.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_updateMinGiftActivity) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize data
	if mask.ReadBit() {
		err := this.data.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_updateIntegralShop) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //data
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize data
	{
		err := this.data.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_updateIntegralShop) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize data
	if mask.ReadBit() {
		err := this.data.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_updateShowBaby) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.playerId != 0)
	mask.WriteBit(this.showBabyTableId != 0)
	mask.WriteBit(len(this.showBabyName) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize playerId
	{
		if this.playerId != 0 {
			err := prpc.Write(buffer, this.playerId)
			if err != nil {
				return err
			}
		}
	}
	// serialize showBabyTableId
	{
		if this.showBabyTableId != 0 {
			err := prpc.Write(buffer, this.showBabyTableId)
			if err != nil {
				return err
			}
		}
	}
	// serialize showBabyName
	if len(this.showBabyName) != 0 {
		err := prpc.Write(buffer, this.showBabyName)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_updateShowBaby) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize playerId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.playerId)
		if err != nil {
			return err
		}
	}
	// deserialize showBabyTableId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.showBabyTableId)
		if err != nil {
			return err
		}
	}
	// deserialize showBabyName
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.showBabyName)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_updateMySelfRecharge) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //val
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize val
	{
		err := this.val.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_updateMySelfRecharge) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize val
	if mask.ReadBit() {
		err := this.val.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_verificationSMSOk) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.phoneNumber) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize phoneNumber
	if len(this.phoneNumber) != 0 {
		err := prpc.Write(buffer, this.phoneNumber)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_verificationSMSOk) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize phoneNumber
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.phoneNumber)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_requestLevelGiftOK) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.level != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize level
	{
		if this.level != 0 {
			err := prpc.Write(buffer, this.level)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_requestLevelGiftOK) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize level
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.level)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_sycnConvertExp) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.val != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize val
	{
		if this.val != 0 {
			err := prpc.Write(buffer, this.val)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_sycnConvertExp) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize val
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.val)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_wearFuwenOk) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //inst
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize inst
	{
		err := this.inst.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_wearFuwenOk) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize inst
	if mask.ReadBit() {
		err := this.inst.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_takeoffFuwenOk) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.slot != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize slot
	{
		if this.slot != 0 {
			err := prpc.Write(buffer, this.slot)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_takeoffFuwenOk) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize slot
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.slot)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_requestEmployeeQuestOk) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.questList) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize questList
	if len(this.questList) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.questList)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.questList {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_requestEmployeeQuestOk) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize questList
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.questList = make([]COM_EmployeeQuestInst, size)
		for i, _ := range this.questList {
			err := this.questList[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_acceptEmployeeQuestOk) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //inst
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize inst
	{
		err := this.inst.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_acceptEmployeeQuestOk) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize inst
	if mask.ReadBit() {
		err := this.inst.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_submitEmployeeQuestOk) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.questId != 0)
	mask.WriteBit(this.isSuccess)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize questId
	{
		if this.questId != 0 {
			err := prpc.Write(buffer, this.questId)
			if err != nil {
				return err
			}
		}
	}
	// serialize isSuccess
	{
	}
	return nil
}
func (this *Server2Client_submitEmployeeQuestOk) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize questId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.questId)
		if err != nil {
			return err
		}
	}
	// deserialize isSuccess
	this.isSuccess = mask.ReadBit()
	return nil
}
func (this *Server2Client_sycnCrystal) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //data
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize data
	{
		err := this.data.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_sycnCrystal) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize data
	if mask.ReadBit() {
		err := this.data.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_crystalUpLeveResult) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.isOK)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize isOK
	{
	}
	return nil
}
func (this *Server2Client_crystalUpLeveResult) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize isOK
	this.isOK = mask.ReadBit()
	return nil
}
func (this *Server2Client_sycnCourseGift) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.data) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize data
	if len(this.data) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.data)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.data {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_sycnCourseGift) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize data
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.data = make([]COM_CourseGift, size)
		for i, _ := range this.data {
			err := this.data[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_orderOk) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.orderId) != 0)
	mask.WriteBit(this.shopId != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize orderId
	if len(this.orderId) != 0 {
		err := prpc.Write(buffer, this.orderId)
		if err != nil {
			return err
		}
	}
	// serialize shopId
	{
		if this.shopId != 0 {
			err := prpc.Write(buffer, this.shopId)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_orderOk) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize orderId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.orderId)
		if err != nil {
			return err
		}
	}
	// deserialize shopId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.shopId)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2Client_updateRandSubmitQuestCount) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.submitCount != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize submitCount
	{
		if this.submitCount != 0 {
			err := prpc.Write(buffer, this.submitCount)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Server2Client_updateRandSubmitQuestCount) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize submitCount
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.submitCount)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Server2ClientStub) pong() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(0))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) errorno(e int) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(1))
	if err != nil {
		return err
	}
	_1 := Server2Client_errorno{}
	_1.e = e
	err = _1.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) teamerrorno(name string, e int) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(2))
	if err != nil {
		return err
	}
	_2 := Server2Client_teamerrorno{}
	_2.name = name
	_2.e = e
	err = _2.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) reconnection(recInfo COM_ReconnectInfo) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(3))
	if err != nil {
		return err
	}
	_3 := Server2Client_reconnection{}
	_3.recInfo = recInfo
	err = _3.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) sessionfailed() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(4))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) loginok(sessionkey string, players []COM_SimpleInformation) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(5))
	if err != nil {
		return err
	}
	_5 := Server2Client_loginok{}
	_5.sessionkey = sessionkey
	_5.players = players
	err = _5.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) logoutOk() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(6))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) createPlayerOk(player COM_SimpleInformation) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(7))
	if err != nil {
		return err
	}
	_7 := Server2Client_createPlayerOk{}
	_7.player = player
	err = _7.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) deletePlayerOk(name string) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(8))
	if err != nil {
		return err
	}
	_8 := Server2Client_deletePlayerOk{}
	_8.name = name
	err = _8.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) enterGameOk(inst COM_PlayerInst) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(9))
	if err != nil {
		return err
	}
	_9 := Server2Client_enterGameOk{}
	_9.inst = inst
	err = _9.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) initBabies(insts []COM_BabyInst) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(10))
	if err != nil {
		return err
	}
	_10 := Server2Client_initBabies{}
	_10.insts = insts
	err = _10.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) initEmployees(insts []COM_EmployeeInst, isFlag bool) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(11))
	if err != nil {
		return err
	}
	_11 := Server2Client_initEmployees{}
	_11.insts = insts
	_11.isFlag = isFlag
	err = _11.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) initEmpBattleGroup(bep COM_BattleEmp) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(12))
	if err != nil {
		return err
	}
	_12 := Server2Client_initEmpBattleGroup{}
	_12.bep = bep
	err = _12.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) initNpc(npcList []int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(13))
	if err != nil {
		return err
	}
	_13 := Server2Client_initNpc{}
	_13.npcList = npcList
	err = _13.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) initAchievement(actlist []COM_Achievement) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(14))
	if err != nil {
		return err
	}
	_14 := Server2Client_initAchievement{}
	_14.actlist = actlist
	err = _14.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) initGather(allnum uint32, gathers []COM_Gather) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(15))
	if err != nil {
		return err
	}
	_15 := Server2Client_initGather{}
	_15.allnum = allnum
	_15.gathers = gathers
	err = _15.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) initcompound(compounds []uint32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(16))
	if err != nil {
		return err
	}
	_16 := Server2Client_initcompound{}
	_16.compounds = compounds
	err = _16.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) addBaby(inst COM_BabyInst) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(17))
	if err != nil {
		return err
	}
	_17 := Server2Client_addBaby{}
	_17.inst = inst
	err = _17.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) refreshBaby(inst COM_BabyInst) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(18))
	if err != nil {
		return err
	}
	_18 := Server2Client_refreshBaby{}
	_18.inst = inst
	err = _18.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) delBabyOK(babyInstId uint32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(19))
	if err != nil {
		return err
	}
	_19 := Server2Client_delBabyOK{}
	_19.babyInstId = babyInstId
	err = _19.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) changeBabyNameOK(babyId uint32, name string) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(20))
	if err != nil {
		return err
	}
	_20 := Server2Client_changeBabyNameOK{}
	_20.babyId = babyId
	_20.name = name
	err = _20.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) remouldBabyOK(instid uint32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(21))
	if err != nil {
		return err
	}
	_21 := Server2Client_remouldBabyOK{}
	_21.instid = instid
	err = _21.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) intensifyBabyOK(babyid uint32, intensifyLevel uint32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(22))
	if err != nil {
		return err
	}
	_22 := Server2Client_intensifyBabyOK{}
	_22.babyid = babyid
	_22.intensifyLevel = intensifyLevel
	err = _22.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) learnSkillOk(inst COM_Skill) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(23))
	if err != nil {
		return err
	}
	_23 := Server2Client_learnSkillOk{}
	_23.inst = inst
	err = _23.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) forgetSkillOk(skid uint32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(24))
	if err != nil {
		return err
	}
	_24 := Server2Client_forgetSkillOk{}
	_24.skid = skid
	err = _24.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) addSkillExp(skid uint32, uExp uint32, flag int) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(25))
	if err != nil {
		return err
	}
	_25 := Server2Client_addSkillExp{}
	_25.skid = skid
	_25.uExp = uExp
	_25.flag = flag
	err = _25.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) babyLearnSkillOK(instId uint32, newSkId uint32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(26))
	if err != nil {
		return err
	}
	_26 := Server2Client_babyLearnSkillOK{}
	_26.instId = instId
	_26.newSkId = newSkId
	err = _26.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) skillLevelUp(instId uint32, inst COM_Skill) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(27))
	if err != nil {
		return err
	}
	_27 := Server2Client_skillLevelUp{}
	_27.instId = instId
	_27.inst = inst
	err = _27.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) joinScene(info COM_SceneInfo) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(28))
	if err != nil {
		return err
	}
	_28 := Server2Client_joinScene{}
	_28.info = info
	err = _28.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) joinCopySceneOK(secneid int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(29))
	if err != nil {
		return err
	}
	_29 := Server2Client_joinCopySceneOK{}
	_29.secneid = secneid
	err = _29.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) initCopyNums() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(30))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) addToScene(inst COM_ScenePlayerInformation) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(31))
	if err != nil {
		return err
	}
	_31 := Server2Client_addToScene{}
	_31.inst = inst
	err = _31.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) delFormScene(instId int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(32))
	if err != nil {
		return err
	}
	_32 := Server2Client_delFormScene{}
	_32.instId = instId
	err = _32.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) move2(instId int32, pos COM_FPosition) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(33))
	if err != nil {
		return err
	}
	_33 := Server2Client_move2{}
	_33.instId = instId
	_33.pos = pos
	err = _33.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) cantMove() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(34))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) querySimplePlayerInstOk(player COM_SimplePlayerInst) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(35))
	if err != nil {
		return err
	}
	_35 := Server2Client_querySimplePlayerInstOk{}
	_35.player = player
	err = _35.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) transfor2(instId int32, pos COM_FPosition) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(36))
	if err != nil {
		return err
	}
	_36 := Server2Client_transfor2{}
	_36.instId = instId
	_36.pos = pos
	err = _36.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) openScene(sceneId int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(37))
	if err != nil {
		return err
	}
	_37 := Server2Client_openScene{}
	_37.sceneId = sceneId
	err = _37.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) autoBattleResult(isOk bool) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(38))
	if err != nil {
		return err
	}
	_38 := Server2Client_autoBattleResult{}
	_38.isOk = isOk
	err = _38.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) talked2Npc(npcId int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(39))
	if err != nil {
		return err
	}
	_39 := Server2Client_talked2Npc{}
	_39.npcId = npcId
	err = _39.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) talked2Player(playerId int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(40))
	if err != nil {
		return err
	}
	_40 := Server2Client_talked2Player{}
	_40.playerId = playerId
	err = _40.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) useItemOk(itemId int32, stack int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(41))
	if err != nil {
		return err
	}
	_41 := Server2Client_useItemOk{}
	_41.itemId = itemId
	_41.stack = stack
	err = _41.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) syncBattleStatus(playerId int32, inBattle bool) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(42))
	if err != nil {
		return err
	}
	_42 := Server2Client_syncBattleStatus{}
	_42.playerId = playerId
	_42.inBattle = inBattle
	err = _42.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) enterBattleOk(initBattle COM_InitBattle) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(43))
	if err != nil {
		return err
	}
	_43 := Server2Client_enterBattleOk{}
	_43.initBattle = initBattle
	err = _43.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) exitBattleOk(bjt int, init COM_BattleOverClearing) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(44))
	if err != nil {
		return err
	}
	_44 := Server2Client_exitBattleOk{}
	_44.bjt = bjt
	_44.init = init
	err = _44.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) syncOrderOk(uid uint32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(45))
	if err != nil {
		return err
	}
	_45 := Server2Client_syncOrderOk{}
	_45.uid = uid
	err = _45.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) syncOrderOkEX() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(46))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) syncOneTurnAction(reports COM_BattleReport) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(47))
	if err != nil {
		return err
	}
	_47 := Server2Client_syncOneTurnAction{}
	_47.reports = reports
	err = _47.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) syncProperties(guid uint32, props []COM_PropValue) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(48))
	if err != nil {
		return err
	}
	_48 := Server2Client_syncProperties{}
	_48.guid = guid
	_48.props = props
	err = _48.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) receiveChat(info COM_ChatInfo, myinfo COM_ContactInfo) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(49))
	if err != nil {
		return err
	}
	_49 := Server2Client_receiveChat{}
	_49.info = info
	_49.myinfo = myinfo
	err = _49.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) requestAudioOk(audioId int32, content []uint8) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(50))
	if err != nil {
		return err
	}
	_50 := Server2Client_requestAudioOk{}
	_50.audioId = audioId
	_50.content = content
	err = _50.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) publishItemInstRes(info COM_ShowItemInstInfo, type_ int) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(51))
	if err != nil {
		return err
	}
	_51 := Server2Client_publishItemInstRes{}
	_51.info = info
	_51.type_ = type_
	err = _51.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) queryItemInstRes(item COM_ShowItemInst) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(52))
	if err != nil {
		return err
	}
	_52 := Server2Client_queryItemInstRes{}
	_52.item = item
	err = _52.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) publishBabyInstRes(info COM_ShowbabyInstInfo, type_ int) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(53))
	if err != nil {
		return err
	}
	_53 := Server2Client_publishBabyInstRes{}
	_53.info = info
	_53.type_ = type_
	err = _53.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) queryBabyInstRes(inst COM_ShowbabyInst) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(54))
	if err != nil {
		return err
	}
	_54 := Server2Client_queryBabyInstRes{}
	_54.inst = inst
	err = _54.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) setNoTalkTime(t float32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(55))
	if err != nil {
		return err
	}
	_55 := Server2Client_setNoTalkTime{}
	_55.t = t
	err = _55.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) addNpc(npcList []int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(56))
	if err != nil {
		return err
	}
	_56 := Server2Client_addNpc{}
	_56.npcList = npcList
	err = _56.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) delNpc(npcList []int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(57))
	if err != nil {
		return err
	}
	_57 := Server2Client_delNpc{}
	_57.npcList = npcList
	err = _57.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) setTeamLeader(playerId int32, isLeader bool) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(58))
	if err != nil {
		return err
	}
	_58 := Server2Client_setTeamLeader{}
	_58.playerId = playerId
	_58.isLeader = isLeader
	err = _58.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) initBag(items []COM_Item) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(59))
	if err != nil {
		return err
	}
	_59 := Server2Client_initBag{}
	_59.items = items
	err = _59.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) addBagItem(item COM_Item) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(60))
	if err != nil {
		return err
	}
	_60 := Server2Client_addBagItem{}
	_60.item = item
	err = _60.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) delBagItem(slot uint16) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(61))
	if err != nil {
		return err
	}
	_61 := Server2Client_delBagItem{}
	_61.slot = slot
	err = _61.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) updateBagItem(item COM_Item) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(62))
	if err != nil {
		return err
	}
	_62 := Server2Client_updateBagItem{}
	_62.item = item
	err = _62.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) depositItemOK(item COM_Item) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(63))
	if err != nil {
		return err
	}
	_63 := Server2Client_depositItemOK{}
	_63.item = item
	err = _63.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) getoutItemOK(slot uint16) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(64))
	if err != nil {
		return err
	}
	_64 := Server2Client_getoutItemOK{}
	_64.slot = slot
	err = _64.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) depositBabyOK(baby COM_BabyInst) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(65))
	if err != nil {
		return err
	}
	_65 := Server2Client_depositBabyOK{}
	_65.baby = baby
	err = _65.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) getoutBabyOK(slot uint16) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(66))
	if err != nil {
		return err
	}
	_66 := Server2Client_getoutBabyOK{}
	_66.slot = slot
	err = _66.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) sortItemStorageOK(items []COM_Item) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(67))
	if err != nil {
		return err
	}
	_67 := Server2Client_sortItemStorageOK{}
	_67.items = items
	err = _67.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) sortBabyStorageOK(babys []uint32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(68))
	if err != nil {
		return err
	}
	_68 := Server2Client_sortBabyStorageOK{}
	_68.babys = babys
	err = _68.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) initItemStorage(gridNum uint16, items []COM_Item) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(69))
	if err != nil {
		return err
	}
	_69 := Server2Client_initItemStorage{}
	_69.gridNum = gridNum
	_69.items = items
	err = _69.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) initBabyStorage(gridNum uint16, babys []COM_BabyInst) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(70))
	if err != nil {
		return err
	}
	_70 := Server2Client_initBabyStorage{}
	_70.gridNum = gridNum
	_70.babys = babys
	err = _70.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) openStorageGrid(tp int, gridNum uint16) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(71))
	if err != nil {
		return err
	}
	_71 := Server2Client_openStorageGrid{}
	_71.tp = tp
	_71.gridNum = gridNum
	err = _71.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) delStorageBabyOK(slot uint16) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(72))
	if err != nil {
		return err
	}
	_72 := Server2Client_delStorageBabyOK{}
	_72.slot = slot
	err = _72.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) initPlayerEquips(equips []COM_Item) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(73))
	if err != nil {
		return err
	}
	_73 := Server2Client_initPlayerEquips{}
	_73.equips = equips
	err = _73.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) wearEquipmentOk(target uint32, equip COM_Item) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(74))
	if err != nil {
		return err
	}
	_74 := Server2Client_wearEquipmentOk{}
	_74.target = target
	_74.equip = equip
	err = _74.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) scenePlayerWearEquipment(target uint32, itemId uint32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(75))
	if err != nil {
		return err
	}
	_75 := Server2Client_scenePlayerWearEquipment{}
	_75.target = target
	_75.itemId = itemId
	err = _75.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) delEquipmentOk(target uint32, itemInstId uint32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(76))
	if err != nil {
		return err
	}
	_76 := Server2Client_delEquipmentOk{}
	_76.target = target
	_76.itemInstId = itemInstId
	err = _76.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) scenePlayerDoffEquipment(target uint32, itemId uint32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(77))
	if err != nil {
		return err
	}
	_77 := Server2Client_scenePlayerDoffEquipment{}
	_77.target = target
	_77.itemId = itemId
	err = _77.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) sortBagItemOk() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(78))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) jointLobbyOk(infos []COM_SimpleTeamInfo) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(79))
	if err != nil {
		return err
	}
	_79 := Server2Client_jointLobbyOk{}
	_79.infos = infos
	err = _79.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) exitLobbyOk() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(80))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) syncDelLobbyTeam(teamId uint32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(81))
	if err != nil {
		return err
	}
	_81 := Server2Client_syncDelLobbyTeam{}
	_81.teamId = teamId
	err = _81.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) syncUpdateLobbyTeam(info COM_SimpleTeamInfo) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(82))
	if err != nil {
		return err
	}
	_82 := Server2Client_syncUpdateLobbyTeam{}
	_82.info = info
	err = _82.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) syncAddLobbyTeam(team COM_SimpleTeamInfo) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(83))
	if err != nil {
		return err
	}
	_83 := Server2Client_syncAddLobbyTeam{}
	_83.team = team
	err = _83.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) createTeamOk(team COM_TeamInfo) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(84))
	if err != nil {
		return err
	}
	_84 := Server2Client_createTeamOk{}
	_84.team = team
	err = _84.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) changeTeamOk(team COM_TeamInfo) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(85))
	if err != nil {
		return err
	}
	_85 := Server2Client_changeTeamOk{}
	_85.team = team
	err = _85.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) joinTeamOk(team COM_TeamInfo) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(86))
	if err != nil {
		return err
	}
	_86 := Server2Client_joinTeamOk{}
	_86.team = team
	err = _86.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) addTeamMember(info COM_SimplePlayerInst) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(87))
	if err != nil {
		return err
	}
	_87 := Server2Client_addTeamMember{}
	_87.info = info
	err = _87.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) delTeamMember(instId int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(88))
	if err != nil {
		return err
	}
	_88 := Server2Client_delTeamMember{}
	_88.instId = instId
	err = _88.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) changeTeamLeaderOk(uuid int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(89))
	if err != nil {
		return err
	}
	_89 := Server2Client_changeTeamLeaderOk{}
	_89.uuid = uuid
	err = _89.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) exitTeamOk(iskick bool) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(90))
	if err != nil {
		return err
	}
	_90 := Server2Client_exitTeamOk{}
	_90.iskick = iskick
	err = _90.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) updateTeam(team COM_TeamInfo) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(91))
	if err != nil {
		return err
	}
	_91 := Server2Client_updateTeam{}
	_91.team = team
	err = _91.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) joinTeamRoomOK(team COM_TeamInfo) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(92))
	if err != nil {
		return err
	}
	_92 := Server2Client_joinTeamRoomOK{}
	_92.team = team
	err = _92.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) inviteJoinTeam(teamId uint32, name string) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(93))
	if err != nil {
		return err
	}
	_93 := Server2Client_inviteJoinTeam{}
	_93.teamId = teamId
	_93.name = name
	err = _93.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) syncTeamDirtyProp(guid int32, props []COM_PropValue) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(94))
	if err != nil {
		return err
	}
	_94 := Server2Client_syncTeamDirtyProp{}
	_94.guid = guid
	_94.props = props
	err = _94.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) leaveTeamOk(playerId int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(95))
	if err != nil {
		return err
	}
	_95 := Server2Client_leaveTeamOk{}
	_95.playerId = playerId
	err = _95.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) backTeamOK(playerId int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(96))
	if err != nil {
		return err
	}
	_96 := Server2Client_backTeamOK{}
	_96.playerId = playerId
	err = _96.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) teamCallMemberBack() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(97))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) refuseBackTeamOk(playerId int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(98))
	if err != nil {
		return err
	}
	_98 := Server2Client_refuseBackTeamOk{}
	_98.playerId = playerId
	err = _98.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) requestJoinTeamTranspond(reqName string) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(99))
	if err != nil {
		return err
	}
	_99 := Server2Client_requestJoinTeamTranspond{}
	_99.reqName = reqName
	err = _99.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) drawLotteryBoxRep(items []COM_Item) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(100))
	if err != nil {
		return err
	}
	_100 := Server2Client_drawLotteryBoxRep{}
	_100.items = items
	err = _100.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) addEmployee(employee COM_EmployeeInst) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(101))
	if err != nil {
		return err
	}
	_101 := Server2Client_addEmployee{}
	_101.employee = employee
	err = _101.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) battleEmployee(empId int32, group int, forbattle bool) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(102))
	if err != nil {
		return err
	}
	_102 := Server2Client_battleEmployee{}
	_102.empId = empId
	_102.group = group
	_102.forbattle = forbattle
	err = _102.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) changeEmpBattleGroupOK(group int) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(103))
	if err != nil {
		return err
	}
	_103 := Server2Client_changeEmpBattleGroupOK{}
	_103.group = group
	err = _103.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) evolveOK(guid int32, qc int) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(104))
	if err != nil {
		return err
	}
	_104 := Server2Client_evolveOK{}
	_104.guid = guid
	_104.qc = qc
	err = _104.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) upStarOK(guid int32, star int32, sk COM_Skill) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(105))
	if err != nil {
		return err
	}
	_105 := Server2Client_upStarOK{}
	_105.guid = guid
	_105.star = star
	_105.sk = sk
	err = _105.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) delEmployeeOK(instids []uint32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(106))
	if err != nil {
		return err
	}
	_106 := Server2Client_delEmployeeOK{}
	_106.instids = instids
	err = _106.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) sycnEmployeeSoul(guid int32, soulNum uint32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(107))
	if err != nil {
		return err
	}
	_107 := Server2Client_sycnEmployeeSoul{}
	_107.guid = guid
	_107.soulNum = soulNum
	err = _107.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) initQuest(qlist []COM_QuestInst, clist []int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(108))
	if err != nil {
		return err
	}
	_108 := Server2Client_initQuest{}
	_108.qlist = qlist
	_108.clist = clist
	err = _108.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) acceptQuestOk(inst COM_QuestInst) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(109))
	if err != nil {
		return err
	}
	_109 := Server2Client_acceptQuestOk{}
	_109.inst = inst
	err = _109.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) submitQuestOk(questId int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(110))
	if err != nil {
		return err
	}
	_110 := Server2Client_submitQuestOk{}
	_110.questId = questId
	err = _110.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) giveupQuestOk(questId int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(111))
	if err != nil {
		return err
	}
	_111 := Server2Client_giveupQuestOk{}
	_111.questId = questId
	err = _111.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) updateQuestInst(inst COM_QuestInst) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(112))
	if err != nil {
		return err
	}
	_112 := Server2Client_updateQuestInst{}
	_112.inst = inst
	err = _112.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) requestContactInfoOk(contact COM_ContactInfo) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(113))
	if err != nil {
		return err
	}
	_113 := Server2Client_requestContactInfoOk{}
	_113.contact = contact
	err = _113.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) addFriendOK(inst COM_ContactInfo) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(114))
	if err != nil {
		return err
	}
	_114 := Server2Client_addFriendOK{}
	_114.inst = inst
	err = _114.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) delFriendOK(instId uint32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(115))
	if err != nil {
		return err
	}
	_115 := Server2Client_delFriendOK{}
	_115.instId = instId
	err = _115.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) addBlacklistOK(inst COM_ContactInfo) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(116))
	if err != nil {
		return err
	}
	_116 := Server2Client_addBlacklistOK{}
	_116.inst = inst
	err = _116.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) delBlacklistOK(instId uint32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(117))
	if err != nil {
		return err
	}
	_117 := Server2Client_delBlacklistOK{}
	_117.instId = instId
	err = _117.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) findFriendFail() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(118))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) referrFriendOK(insts []COM_ContactInfo) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(119))
	if err != nil {
		return err
	}
	_119 := Server2Client_referrFriendOK{}
	_119.insts = insts
	err = _119.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) requestFriendListOK(insts []COM_ContactInfo) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(120))
	if err != nil {
		return err
	}
	_120 := Server2Client_requestFriendListOK{}
	_120.insts = insts
	err = _120.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) lotteryOk(lotteryId int32, dropItem []COM_DropItem) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(121))
	if err != nil {
		return err
	}
	_121 := Server2Client_lotteryOk{}
	_121.lotteryId = lotteryId
	_121.dropItem = dropItem
	err = _121.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) openGatherOK(gather COM_Gather) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(122))
	if err != nil {
		return err
	}
	_122 := Server2Client_openGatherOK{}
	_122.gather = gather
	err = _122.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) miningOk(items []COM_DropItem, gather COM_Gather, gatherNum uint32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(123))
	if err != nil {
		return err
	}
	_123 := Server2Client_miningOk{}
	_123.items = items
	_123.gather = gather
	_123.gatherNum = gatherNum
	err = _123.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) openCompound(compoundId uint32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(124))
	if err != nil {
		return err
	}
	_124 := Server2Client_openCompound{}
	_124.compoundId = compoundId
	err = _124.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) compoundItemOk(item COM_Item) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(125))
	if err != nil {
		return err
	}
	_125 := Server2Client_compoundItemOk{}
	_125.item = item
	err = _125.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) openBagGridOk(num int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(126))
	if err != nil {
		return err
	}
	_126 := Server2Client_openBagGridOk{}
	_126.num = num
	err = _126.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) requestChallengeOK(isOK bool) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(127))
	if err != nil {
		return err
	}
	_127 := Server2Client_requestChallengeOK{}
	_127.isOK = isOK
	err = _127.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) requestMySelfJJCDataOK(info COM_EndlessStair) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(128))
	if err != nil {
		return err
	}
	_128 := Server2Client_requestMySelfJJCDataOK{}
	_128.info = info
	err = _128.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) requestRivalOK(infos []COM_EndlessStair) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(129))
	if err != nil {
		return err
	}
	_129 := Server2Client_requestRivalOK{}
	_129.infos = infos
	err = _129.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) rivalTimeOK() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(130))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) checkMsgOK(inst COM_SimplePlayerInst) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(131))
	if err != nil {
		return err
	}
	_131 := Server2Client_checkMsgOK{}
	_131.inst = inst
	err = _131.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) requestMyAllbattleMsgOK(infos []COM_JJCBattleMsg) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(132))
	if err != nil {
		return err
	}
	_132 := Server2Client_requestMyAllbattleMsgOK{}
	_132.infos = infos
	err = _132.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) myBattleMsgOK(info COM_JJCBattleMsg) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(133))
	if err != nil {
		return err
	}
	_133 := Server2Client_myBattleMsgOK{}
	_133.info = info
	err = _133.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) requestJJCRankOK(myRank uint32, infos []COM_EndlessStair) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(134))
	if err != nil {
		return err
	}
	_134 := Server2Client_requestJJCRankOK{}
	_134.myRank = myRank
	_134.infos = infos
	err = _134.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) requestLevelRankOK(myRank uint32, infos []COM_ContactInfo) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(135))
	if err != nil {
		return err
	}
	_135 := Server2Client_requestLevelRankOK{}
	_135.myRank = myRank
	_135.infos = infos
	err = _135.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) requestBabyRankOK(myRank uint32, infos []COM_BabyRankData) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(136))
	if err != nil {
		return err
	}
	_136 := Server2Client_requestBabyRankOK{}
	_136.myRank = myRank
	_136.infos = infos
	err = _136.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) requestEmpRankOK(myRank uint32, infos []COM_EmployeeRankData) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(137))
	if err != nil {
		return err
	}
	_137 := Server2Client_requestEmpRankOK{}
	_137.myRank = myRank
	_137.infos = infos
	err = _137.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) requestPlayerFFRankOK(myRank uint32, infos []COM_ContactInfo) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(138))
	if err != nil {
		return err
	}
	_138 := Server2Client_requestPlayerFFRankOK{}
	_138.myRank = myRank
	_138.infos = infos
	err = _138.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) queryOnlinePlayerOK(isOnline bool) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(139))
	if err != nil {
		return err
	}
	_139 := Server2Client_queryOnlinePlayerOK{}
	_139.isOnline = isOnline
	err = _139.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) queryPlayerOK(inst COM_SimplePlayerInst) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(140))
	if err != nil {
		return err
	}
	_140 := Server2Client_queryPlayerOK{}
	_140.inst = inst
	err = _140.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) queryBabyOK(inst COM_BabyInst) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(141))
	if err != nil {
		return err
	}
	_141 := Server2Client_queryBabyOK{}
	_141.inst = inst
	err = _141.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) queryEmployeeOK(inst COM_EmployeeInst) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(142))
	if err != nil {
		return err
	}
	_142 := Server2Client_queryEmployeeOK{}
	_142.inst = inst
	err = _142.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) initGuide(guideMask uint32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(143))
	if err != nil {
		return err
	}
	_143 := Server2Client_initGuide{}
	_143.guideMask = guideMask
	err = _143.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) buyShopItemOk(id int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(144))
	if err != nil {
		return err
	}
	_144 := Server2Client_buyShopItemOk{}
	_144.id = id
	err = _144.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) addPlayerTitle(title int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(145))
	if err != nil {
		return err
	}
	_145 := Server2Client_addPlayerTitle{}
	_145.title = title
	err = _145.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) delPlayerTitle(title int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(146))
	if err != nil {
		return err
	}
	_146 := Server2Client_delPlayerTitle{}
	_146.title = title
	err = _146.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) requestOpenBuyBox(greenTime float32, blueTime float32, greenFreeNum int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(147))
	if err != nil {
		return err
	}
	_147 := Server2Client_requestOpenBuyBox{}
	_147.greenTime = greenTime
	_147.blueTime = blueTime
	_147.greenFreeNum = greenFreeNum
	err = _147.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) requestGreenBoxTimeOk() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(148))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) requestBlueBoxTimeOk() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(149))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) updateAchievementinfo(achs COM_Achievement) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(150))
	if err != nil {
		return err
	}
	_150 := Server2Client_updateAchievementinfo{}
	_150.achs = achs
	err = _150.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) syncOpenSystemFlag(flag uint64) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(151))
	if err != nil {
		return err
	}
	_151 := Server2Client_syncOpenSystemFlag{}
	_151.flag = flag
	err = _151.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) requestActivityRewardOK(ar uint32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(152))
	if err != nil {
		return err
	}
	_152 := Server2Client_requestActivityRewardOK{}
	_152.ar = ar
	err = _152.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) syncActivity(table COM_ActivityTable) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(153))
	if err != nil {
		return err
	}
	_153 := Server2Client_syncActivity{}
	_153.table = table
	err = _153.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) updateActivityStatus(type_ int, open bool) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(154))
	if err != nil {
		return err
	}
	_154 := Server2Client_updateActivityStatus{}
	_154.type_ = type_
	_154.open = open
	err = _154.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) updateActivityCounter(type_ int, counter int32, reward int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(155))
	if err != nil {
		return err
	}
	_155 := Server2Client_updateActivityCounter{}
	_155.type_ = type_
	_155.counter = counter
	_155.reward = reward
	err = _155.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) syncExam(exam COM_Exam) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(156))
	if err != nil {
		return err
	}
	_156 := Server2Client_syncExam{}
	_156.exam = exam
	err = _156.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) syncExamAnswer(answer COM_Answer) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(157))
	if err != nil {
		return err
	}
	_157 := Server2Client_syncExamAnswer{}
	_157.answer = answer
	err = _157.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) petActivityNoNum(name string) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(158))
	if err != nil {
		return err
	}
	_158 := Server2Client_petActivityNoNum{}
	_158.name = name
	err = _158.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) syncHundredInfo(hb COM_HundredBattle) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(159))
	if err != nil {
		return err
	}
	_159 := Server2Client_syncHundredInfo{}
	_159.hb = hb
	err = _159.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) initSignUp(info []int32, process int32, sign7 bool, sign14 bool, sign28 bool) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(160))
	if err != nil {
		return err
	}
	_160 := Server2Client_initSignUp{}
	_160.info = info
	_160.process = process
	_160.sign7 = sign7
	_160.sign14 = sign14
	_160.sign28 = sign28
	err = _160.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) signUp(flag bool) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(161))
	if err != nil {
		return err
	}
	_161 := Server2Client_signUp{}
	_161.flag = flag
	err = _161.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) requestSignupRewardOk7() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(162))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) requestSignupRewardOk14() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(163))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) requestSignupRewardOk28() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(164))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) sycnDoubleExpTime(isFlag bool, times float32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(165))
	if err != nil {
		return err
	}
	_165 := Server2Client_sycnDoubleExpTime{}
	_165.isFlag = isFlag
	_165.times = times
	err = _165.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) sycnStates(states []COM_State) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(166))
	if err != nil {
		return err
	}
	_166 := Server2Client_sycnStates{}
	_166.states = states
	err = _166.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) requestpvprankOK(infos []COM_ContactInfo) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(167))
	if err != nil {
		return err
	}
	_167 := Server2Client_requestpvprankOK{}
	_167.infos = infos
	err = _167.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) syncMyJJCTeamMember() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(168))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) startMatchingOK() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(169))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) stopMatchingOK(times float32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(170))
	if err != nil {
		return err
	}
	_170 := Server2Client_stopMatchingOK{}
	_170.times = times
	err = _170.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) updatePvpJJCinfo(info COM_PlayerVsPlayer) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(171))
	if err != nil {
		return err
	}
	_171 := Server2Client_updatePvpJJCinfo{}
	_171.info = info
	err = _171.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) exitPvpJJCOk() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(172))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) syncEnemyPvpJJCPlayerInfo(info COM_SimpleInformation) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(173))
	if err != nil {
		return err
	}
	_173 := Server2Client_syncEnemyPvpJJCPlayerInfo{}
	_173.info = info
	err = _173.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) syncEnemyPvpJJCTeamInfo(infos []COM_SimpleInformation, teamID_ uint32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(174))
	if err != nil {
		return err
	}
	_174 := Server2Client_syncEnemyPvpJJCTeamInfo{}
	_174.infos = infos
	_174.teamID_ = teamID_
	err = _174.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) openWarriorchooseUI() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(175))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) warriorStartOK() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(176))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) warriorStopOK() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(177))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) syncWarriorEnemyTeamInfo(infos []COM_SimpleInformation, teamID_ uint32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(178))
	if err != nil {
		return err
	}
	_178 := Server2Client_syncWarriorEnemyTeamInfo{}
	_178.infos = infos
	_178.teamID_ = teamID_
	err = _178.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) appendMail(mails []COM_Mail) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(179))
	if err != nil {
		return err
	}
	_179 := Server2Client_appendMail{}
	_179.mails = mails
	err = _179.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) delMail(mailId int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(180))
	if err != nil {
		return err
	}
	_180 := Server2Client_delMail{}
	_180.mailId = mailId
	err = _180.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) updateMailOk(mail COM_Mail) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(181))
	if err != nil {
		return err
	}
	_181 := Server2Client_updateMailOk{}
	_181.mail = mail
	err = _181.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) boardcastNotice(content string, isGm bool) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(182))
	if err != nil {
		return err
	}
	_182 := Server2Client_boardcastNotice{}
	_182.content = content
	_182.isGm = isGm
	err = _182.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) createGuildOK() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(183))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) delGuildOK() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(184))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) leaveGuildOk(who string, isKick bool) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(185))
	if err != nil {
		return err
	}
	_185 := Server2Client_leaveGuildOk{}
	_185.who = who
	_185.isKick = isKick
	err = _185.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) initGuildData(guild COM_Guild) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(186))
	if err != nil {
		return err
	}
	_186 := Server2Client_initGuildData{}
	_186.guild = guild
	err = _186.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) initGuildMemberList(member []COM_GuildMember) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(187))
	if err != nil {
		return err
	}
	_187 := Server2Client_initGuildMemberList{}
	_187.member = member
	err = _187.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) modifyGuildMemberList(member COM_GuildMember, flag int) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(188))
	if err != nil {
		return err
	}
	_188 := Server2Client_modifyGuildMemberList{}
	_188.member = member
	_188.flag = flag
	err = _188.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) modifyGuildList(data COM_GuildViewerData, flag int) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(189))
	if err != nil {
		return err
	}
	_189 := Server2Client_modifyGuildList{}
	_189.data = data
	_189.flag = flag
	err = _189.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) queryGuildListResult(page int16, pageNum int16, guildList []COM_GuildViewerData) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(190))
	if err != nil {
		return err
	}
	_190 := Server2Client_queryGuildListResult{}
	_190.page = page
	_190.pageNum = pageNum
	_190.guildList = guildList
	err = _190.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) inviteGuild(sendName string, guildName string) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(191))
	if err != nil {
		return err
	}
	_191 := Server2Client_inviteGuild{}
	_191.sendName = sendName
	_191.guildName = guildName
	err = _191.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) updateGuildShopItems(items []COM_GuildShopItem) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(192))
	if err != nil {
		return err
	}
	_192 := Server2Client_updateGuildShopItems{}
	_192.items = items
	err = _192.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) updateGuildBuilding(type_ int, building COM_GuildBuilding) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(193))
	if err != nil {
		return err
	}
	_193 := Server2Client_updateGuildBuilding{}
	_193.type_ = type_
	_193.building = building
	err = _193.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) updateGuildMyMember(member COM_GuildMember) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(194))
	if err != nil {
		return err
	}
	_194 := Server2Client_updateGuildMyMember{}
	_194.member = member
	err = _194.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) levelupGuildSkillOk(skInst COM_Skill) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(195))
	if err != nil {
		return err
	}
	_195 := Server2Client_levelupGuildSkillOk{}
	_195.skInst = skInst
	err = _195.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) presentGuildItemOk(val int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(196))
	if err != nil {
		return err
	}
	_196 := Server2Client_presentGuildItemOk{}
	_196.val = val
	err = _196.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) progenitusAddExpOk(mInst COM_GuildProgen) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(197))
	if err != nil {
		return err
	}
	_197 := Server2Client_progenitusAddExpOk{}
	_197.mInst = mInst
	err = _197.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) setProgenitusPositionOk(positions []int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(198))
	if err != nil {
		return err
	}
	_198 := Server2Client_setProgenitusPositionOk{}
	_198.positions = positions
	err = _198.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) updateGuildFundz(val int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(199))
	if err != nil {
		return err
	}
	_199 := Server2Client_updateGuildFundz{}
	_199.val = val
	err = _199.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) updateGuildMemberContribution(val int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(200))
	if err != nil {
		return err
	}
	_200 := Server2Client_updateGuildMemberContribution{}
	_200.val = val
	err = _200.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) openGuildBattle(otherName string, playerNum int32, level int32, isLeft bool, lstime int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(201))
	if err != nil {
		return err
	}
	_201 := Server2Client_openGuildBattle{}
	_201.otherName = otherName
	_201.playerNum = playerNum
	_201.level = level
	_201.isLeft = isLeft
	_201.lstime = lstime
	err = _201.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) startGuildBattle(otherName string, otherCon int32, selfCon int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(202))
	if err != nil {
		return err
	}
	_202 := Server2Client_startGuildBattle{}
	_202.otherName = otherName
	_202.otherCon = otherCon
	_202.selfCon = selfCon
	err = _202.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) closeGuildBattle(isWinner bool) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(203))
	if err != nil {
		return err
	}
	_203 := Server2Client_closeGuildBattle{}
	_203.isWinner = isWinner
	err = _203.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) syncGuildBattleWinCount(myWin int32, otherWin int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(204))
	if err != nil {
		return err
	}
	_204 := Server2Client_syncGuildBattleWinCount{}
	_204.myWin = myWin
	_204.otherWin = otherWin
	err = _204.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) initMySelling(items []COM_SellItem) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(205))
	if err != nil {
		return err
	}
	_205 := Server2Client_initMySelling{}
	_205.items = items
	err = _205.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) initMySelled(items []COM_SelledItem) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(206))
	if err != nil {
		return err
	}
	_206 := Server2Client_initMySelled{}
	_206.items = items
	err = _206.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) fetchSellingOk(items []COM_SellItem, total int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(207))
	if err != nil {
		return err
	}
	_207 := Server2Client_fetchSellingOk{}
	_207.items = items
	_207.total = total
	err = _207.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) fetchSellingOk2(items []COM_SellItem, total int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(208))
	if err != nil {
		return err
	}
	_208 := Server2Client_fetchSellingOk2{}
	_208.items = items
	_208.total = total
	err = _208.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) sellingOk(item COM_SellItem) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(209))
	if err != nil {
		return err
	}
	_209 := Server2Client_sellingOk{}
	_209.item = item
	err = _209.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) selledOk(item COM_SelledItem) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(210))
	if err != nil {
		return err
	}
	_210 := Server2Client_selledOk{}
	_210.item = item
	err = _210.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) unsellingOk(sellid int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(211))
	if err != nil {
		return err
	}
	_211 := Server2Client_unsellingOk{}
	_211.sellid = sellid
	err = _211.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) redemptionSpreeOk() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(212))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) insertState(st COM_State) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(213))
	if err != nil {
		return err
	}
	_213 := Server2Client_insertState{}
	_213.st = st
	err = _213.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) updattState(st COM_State) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(214))
	if err != nil {
		return err
	}
	_214 := Server2Client_updattState{}
	_214.st = st
	err = _214.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) removeState(stid uint32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(215))
	if err != nil {
		return err
	}
	_215 := Server2Client_removeState{}
	_215.stid = stid
	err = _215.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) requestFixItemOk(item COM_Item) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(216))
	if err != nil {
		return err
	}
	_216 := Server2Client_requestFixItemOk{}
	_216.item = item
	err = _216.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) makeDebirsItemOK() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(217))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) updateMagicItem(level int32, exp int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(218))
	if err != nil {
		return err
	}
	_218 := Server2Client_updateMagicItem{}
	_218.level = level
	_218.exp = exp
	err = _218.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) changeMagicJobOk(job int) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(219))
	if err != nil {
		return err
	}
	_219 := Server2Client_changeMagicJobOk{}
	_219.job = job
	err = _219.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) magicItemTupoOk(level int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(220))
	if err != nil {
		return err
	}
	_220 := Server2Client_magicItemTupoOk{}
	_220.level = level
	err = _220.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) zhuanpanOK(pond []uint32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(221))
	if err != nil {
		return err
	}
	_221 := Server2Client_zhuanpanOK{}
	_221.pond = pond
	err = _221.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) updateZhuanpanNotice(zhuanp COM_Zhuanpan) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(222))
	if err != nil {
		return err
	}
	_222 := Server2Client_updateZhuanpanNotice{}
	_222.zhuanp = zhuanp
	err = _222.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) sycnZhuanpanData(data COM_ZhuanpanData) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(223))
	if err != nil {
		return err
	}
	_223 := Server2Client_sycnZhuanpanData{}
	_223.data = data
	err = _223.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) copynonum(name string) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(224))
	if err != nil {
		return err
	}
	_224 := Server2Client_copynonum{}
	_224.name = name
	err = _224.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) sceneFilterOk(sfType []int) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(225))
	if err != nil {
		return err
	}
	_225 := Server2Client_sceneFilterOk{}
	_225.sfType = sfType
	err = _225.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) wishingOK() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(226))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) shareWishOK(wish COM_Wish) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(227))
	if err != nil {
		return err
	}
	_227 := Server2Client_shareWishOK{}
	_227.wish = wish
	err = _227.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) leaderCloseDialogOk() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(228))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) startOnlineTime() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(229))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) requestOnlineTimeRewardOK(index uint32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(230))
	if err != nil {
		return err
	}
	_230 := Server2Client_requestOnlineTimeRewardOK{}
	_230.index = index
	err = _230.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) sycnVipflag(flag bool) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(231))
	if err != nil {
		return err
	}
	_231 := Server2Client_sycnVipflag{}
	_231.flag = flag
	err = _231.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) buyFundOK(flag bool) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(232))
	if err != nil {
		return err
	}
	_232 := Server2Client_buyFundOK{}
	_232.flag = flag
	err = _232.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) requestFundRewardOK(level uint32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(233))
	if err != nil {
		return err
	}
	_233 := Server2Client_requestFundRewardOK{}
	_233.level = level
	err = _233.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) firstRechargeOK(isFlag bool) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(234))
	if err != nil {
		return err
	}
	_234 := Server2Client_firstRechargeOK{}
	_234.isFlag = isFlag
	err = _234.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) firstRechargeGiftOK(isFlag bool) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(235))
	if err != nil {
		return err
	}
	_235 := Server2Client_firstRechargeGiftOK{}
	_235.isFlag = isFlag
	err = _235.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) agencyActivity(type_ int, isFlag bool) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(236))
	if err != nil {
		return err
	}
	_236 := Server2Client_agencyActivity{}
	_236.type_ = type_
	_236.isFlag = isFlag
	err = _236.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) updateFestival(festival COM_ADLoginTotal) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(237))
	if err != nil {
		return err
	}
	_237 := Server2Client_updateFestival{}
	_237.festival = festival
	err = _237.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) updateSelfRecharge(val COM_ADChargeTotal) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(238))
	if err != nil {
		return err
	}
	_238 := Server2Client_updateSelfRecharge{}
	_238.val = val
	err = _238.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) updateSysRecharge(val COM_ADChargeTotal) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(239))
	if err != nil {
		return err
	}
	_239 := Server2Client_updateSysRecharge{}
	_239.val = val
	err = _239.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) updateSelfDiscountStore(val COM_ADDiscountStore) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(240))
	if err != nil {
		return err
	}
	_240 := Server2Client_updateSelfDiscountStore{}
	_240.val = val
	err = _240.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) updateSysDiscountStore(val COM_ADDiscountStore) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(241))
	if err != nil {
		return err
	}
	_241 := Server2Client_updateSysDiscountStore{}
	_241.val = val
	err = _241.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) updateSelfOnceRecharge(val COM_ADChargeEvery) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(242))
	if err != nil {
		return err
	}
	_242 := Server2Client_updateSelfOnceRecharge{}
	_242.val = val
	err = _242.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) updateSysOnceRecharge(val COM_ADChargeEvery) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(243))
	if err != nil {
		return err
	}
	_243 := Server2Client_updateSysOnceRecharge{}
	_243.val = val
	err = _243.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) openCardOK(data COM_ADCardsContent) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(244))
	if err != nil {
		return err
	}
	_244 := Server2Client_openCardOK{}
	_244.data = data
	err = _244.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) resetCardOK() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(245))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) sycnHotRole(data COM_ADHotRole) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(246))
	if err != nil {
		return err
	}
	_246 := Server2Client_sycnHotRole{}
	_246.data = data
	err = _246.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) hotRoleBuyOk(buyNum uint16) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(247))
	if err != nil {
		return err
	}
	_247 := Server2Client_hotRoleBuyOk{}
	_247.buyNum = buyNum
	err = _247.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) updateSevenday(data COM_Sevenday) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(248))
	if err != nil {
		return err
	}
	_248 := Server2Client_updateSevenday{}
	_248.data = data
	err = _248.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) updateEmployeeActivity(data COM_ADEmployeeTotal) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(249))
	if err != nil {
		return err
	}
	_249 := Server2Client_updateEmployeeActivity{}
	_249.data = data
	err = _249.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) updateMinGiftActivity(data COM_ADGiftBag) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(250))
	if err != nil {
		return err
	}
	_250 := Server2Client_updateMinGiftActivity{}
	_250.data = data
	err = _250.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) updateIntegralShop(data COM_IntegralData) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(251))
	if err != nil {
		return err
	}
	_251 := Server2Client_updateIntegralShop{}
	_251.data = data
	err = _251.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) updateShowBaby(playerId uint32, showBabyTableId uint32, showBabyName string) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(252))
	if err != nil {
		return err
	}
	_252 := Server2Client_updateShowBaby{}
	_252.playerId = playerId
	_252.showBabyTableId = showBabyTableId
	_252.showBabyName = showBabyName
	err = _252.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) updateMySelfRecharge(val COM_ADChargeTotal) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(253))
	if err != nil {
		return err
	}
	_253 := Server2Client_updateMySelfRecharge{}
	_253.val = val
	err = _253.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) verificationSMSOk(phoneNumber string) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(254))
	if err != nil {
		return err
	}
	_254 := Server2Client_verificationSMSOk{}
	_254.phoneNumber = phoneNumber
	err = _254.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) requestLevelGiftOK(level int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(255))
	if err != nil {
		return err
	}
	_255 := Server2Client_requestLevelGiftOK{}
	_255.level = level
	err = _255.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) sycnConvertExp(val int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(256))
	if err != nil {
		return err
	}
	_256 := Server2Client_sycnConvertExp{}
	_256.val = val
	err = _256.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) wearFuwenOk(inst COM_Item) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(257))
	if err != nil {
		return err
	}
	_257 := Server2Client_wearFuwenOk{}
	_257.inst = inst
	err = _257.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) takeoffFuwenOk(slot int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(258))
	if err != nil {
		return err
	}
	_258 := Server2Client_takeoffFuwenOk{}
	_258.slot = slot
	err = _258.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) compFuwenOk() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(259))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) requestEmployeeQuestOk(questList []COM_EmployeeQuestInst) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(260))
	if err != nil {
		return err
	}
	_260 := Server2Client_requestEmployeeQuestOk{}
	_260.questList = questList
	err = _260.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) acceptEmployeeQuestOk(inst COM_EmployeeQuestInst) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(261))
	if err != nil {
		return err
	}
	_261 := Server2Client_acceptEmployeeQuestOk{}
	_261.inst = inst
	err = _261.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) submitEmployeeQuestOk(questId int32, isSuccess bool) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(262))
	if err != nil {
		return err
	}
	_262 := Server2Client_submitEmployeeQuestOk{}
	_262.questId = questId
	_262.isSuccess = isSuccess
	err = _262.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) sycnCrystal(data COM_CrystalData) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(263))
	if err != nil {
		return err
	}
	_263 := Server2Client_sycnCrystal{}
	_263.data = data
	err = _263.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) crystalUpLeveResult(isOK bool) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(264))
	if err != nil {
		return err
	}
	_264 := Server2Client_crystalUpLeveResult{}
	_264.isOK = isOK
	err = _264.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) resetCrystalPropOK() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(265))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) sycnCourseGift(data []COM_CourseGift) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(266))
	if err != nil {
		return err
	}
	_266 := Server2Client_sycnCourseGift{}
	_266.data = data
	err = _266.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) orderOk(orderId string, shopId int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(267))
	if err != nil {
		return err
	}
	_267 := Server2Client_orderOk{}
	_267.orderId = orderId
	_267.shopId = shopId
	err = _267.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Server2ClientStub) updateRandSubmitQuestCount(submitCount int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(268))
	if err != nil {
		return err
	}
	_268 := Server2Client_updateRandSubmitQuestCount{}
	_268.submitCount = submitCount
	err = _268.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func Bridging_Server2Client_pong(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.pong()
}
func Bridging_Server2Client_errorno(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_1 := Server2Client_errorno{}
	err := _1.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.errorno(_1.e)
}
func Bridging_Server2Client_teamerrorno(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_2 := Server2Client_teamerrorno{}
	err := _2.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.teamerrorno(_2.name, _2.e)
}
func Bridging_Server2Client_reconnection(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_3 := Server2Client_reconnection{}
	err := _3.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.reconnection(_3.recInfo)
}
func Bridging_Server2Client_sessionfailed(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.sessionfailed()
}
func Bridging_Server2Client_loginok(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_5 := Server2Client_loginok{}
	err := _5.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.loginok(_5.sessionkey, _5.players)
}
func Bridging_Server2Client_logoutOk(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.logoutOk()
}
func Bridging_Server2Client_createPlayerOk(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_7 := Server2Client_createPlayerOk{}
	err := _7.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.createPlayerOk(_7.player)
}
func Bridging_Server2Client_deletePlayerOk(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_8 := Server2Client_deletePlayerOk{}
	err := _8.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.deletePlayerOk(_8.name)
}
func Bridging_Server2Client_enterGameOk(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_9 := Server2Client_enterGameOk{}
	err := _9.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.enterGameOk(_9.inst)
}
func Bridging_Server2Client_initBabies(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_10 := Server2Client_initBabies{}
	err := _10.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.initBabies(_10.insts)
}
func Bridging_Server2Client_initEmployees(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_11 := Server2Client_initEmployees{}
	err := _11.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.initEmployees(_11.insts, _11.isFlag)
}
func Bridging_Server2Client_initEmpBattleGroup(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_12 := Server2Client_initEmpBattleGroup{}
	err := _12.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.initEmpBattleGroup(_12.bep)
}
func Bridging_Server2Client_initNpc(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_13 := Server2Client_initNpc{}
	err := _13.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.initNpc(_13.npcList)
}
func Bridging_Server2Client_initAchievement(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_14 := Server2Client_initAchievement{}
	err := _14.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.initAchievement(_14.actlist)
}
func Bridging_Server2Client_initGather(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_15 := Server2Client_initGather{}
	err := _15.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.initGather(_15.allnum, _15.gathers)
}
func Bridging_Server2Client_initcompound(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_16 := Server2Client_initcompound{}
	err := _16.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.initcompound(_16.compounds)
}
func Bridging_Server2Client_addBaby(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_17 := Server2Client_addBaby{}
	err := _17.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.addBaby(_17.inst)
}
func Bridging_Server2Client_refreshBaby(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_18 := Server2Client_refreshBaby{}
	err := _18.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.refreshBaby(_18.inst)
}
func Bridging_Server2Client_delBabyOK(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_19 := Server2Client_delBabyOK{}
	err := _19.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.delBabyOK(_19.babyInstId)
}
func Bridging_Server2Client_changeBabyNameOK(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_20 := Server2Client_changeBabyNameOK{}
	err := _20.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.changeBabyNameOK(_20.babyId, _20.name)
}
func Bridging_Server2Client_remouldBabyOK(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_21 := Server2Client_remouldBabyOK{}
	err := _21.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.remouldBabyOK(_21.instid)
}
func Bridging_Server2Client_intensifyBabyOK(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_22 := Server2Client_intensifyBabyOK{}
	err := _22.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.intensifyBabyOK(_22.babyid, _22.intensifyLevel)
}
func Bridging_Server2Client_learnSkillOk(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_23 := Server2Client_learnSkillOk{}
	err := _23.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.learnSkillOk(_23.inst)
}
func Bridging_Server2Client_forgetSkillOk(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_24 := Server2Client_forgetSkillOk{}
	err := _24.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.forgetSkillOk(_24.skid)
}
func Bridging_Server2Client_addSkillExp(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_25 := Server2Client_addSkillExp{}
	err := _25.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.addSkillExp(_25.skid, _25.uExp, _25.flag)
}
func Bridging_Server2Client_babyLearnSkillOK(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_26 := Server2Client_babyLearnSkillOK{}
	err := _26.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.babyLearnSkillOK(_26.instId, _26.newSkId)
}
func Bridging_Server2Client_skillLevelUp(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_27 := Server2Client_skillLevelUp{}
	err := _27.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.skillLevelUp(_27.instId, _27.inst)
}
func Bridging_Server2Client_joinScene(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_28 := Server2Client_joinScene{}
	err := _28.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.joinScene(_28.info)
}
func Bridging_Server2Client_joinCopySceneOK(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_29 := Server2Client_joinCopySceneOK{}
	err := _29.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.joinCopySceneOK(_29.secneid)
}
func Bridging_Server2Client_initCopyNums(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.initCopyNums()
}
func Bridging_Server2Client_addToScene(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_31 := Server2Client_addToScene{}
	err := _31.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.addToScene(_31.inst)
}
func Bridging_Server2Client_delFormScene(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_32 := Server2Client_delFormScene{}
	err := _32.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.delFormScene(_32.instId)
}
func Bridging_Server2Client_move2(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_33 := Server2Client_move2{}
	err := _33.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.move2(_33.instId, _33.pos)
}
func Bridging_Server2Client_cantMove(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.cantMove()
}
func Bridging_Server2Client_querySimplePlayerInstOk(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_35 := Server2Client_querySimplePlayerInstOk{}
	err := _35.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.querySimplePlayerInstOk(_35.player)
}
func Bridging_Server2Client_transfor2(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_36 := Server2Client_transfor2{}
	err := _36.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.transfor2(_36.instId, _36.pos)
}
func Bridging_Server2Client_openScene(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_37 := Server2Client_openScene{}
	err := _37.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.openScene(_37.sceneId)
}
func Bridging_Server2Client_autoBattleResult(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_38 := Server2Client_autoBattleResult{}
	err := _38.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.autoBattleResult(_38.isOk)
}
func Bridging_Server2Client_talked2Npc(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_39 := Server2Client_talked2Npc{}
	err := _39.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.talked2Npc(_39.npcId)
}
func Bridging_Server2Client_talked2Player(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_40 := Server2Client_talked2Player{}
	err := _40.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.talked2Player(_40.playerId)
}
func Bridging_Server2Client_useItemOk(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_41 := Server2Client_useItemOk{}
	err := _41.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.useItemOk(_41.itemId, _41.stack)
}
func Bridging_Server2Client_syncBattleStatus(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_42 := Server2Client_syncBattleStatus{}
	err := _42.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.syncBattleStatus(_42.playerId, _42.inBattle)
}
func Bridging_Server2Client_enterBattleOk(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_43 := Server2Client_enterBattleOk{}
	err := _43.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.enterBattleOk(_43.initBattle)
}
func Bridging_Server2Client_exitBattleOk(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_44 := Server2Client_exitBattleOk{}
	err := _44.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.exitBattleOk(_44.bjt, _44.init)
}
func Bridging_Server2Client_syncOrderOk(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_45 := Server2Client_syncOrderOk{}
	err := _45.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.syncOrderOk(_45.uid)
}
func Bridging_Server2Client_syncOrderOkEX(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.syncOrderOkEX()
}
func Bridging_Server2Client_syncOneTurnAction(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_47 := Server2Client_syncOneTurnAction{}
	err := _47.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.syncOneTurnAction(_47.reports)
}
func Bridging_Server2Client_syncProperties(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_48 := Server2Client_syncProperties{}
	err := _48.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.syncProperties(_48.guid, _48.props)
}
func Bridging_Server2Client_receiveChat(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_49 := Server2Client_receiveChat{}
	err := _49.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.receiveChat(_49.info, _49.myinfo)
}
func Bridging_Server2Client_requestAudioOk(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_50 := Server2Client_requestAudioOk{}
	err := _50.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.requestAudioOk(_50.audioId, _50.content)
}
func Bridging_Server2Client_publishItemInstRes(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_51 := Server2Client_publishItemInstRes{}
	err := _51.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.publishItemInstRes(_51.info, _51.type_)
}
func Bridging_Server2Client_queryItemInstRes(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_52 := Server2Client_queryItemInstRes{}
	err := _52.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.queryItemInstRes(_52.item)
}
func Bridging_Server2Client_publishBabyInstRes(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_53 := Server2Client_publishBabyInstRes{}
	err := _53.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.publishBabyInstRes(_53.info, _53.type_)
}
func Bridging_Server2Client_queryBabyInstRes(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_54 := Server2Client_queryBabyInstRes{}
	err := _54.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.queryBabyInstRes(_54.inst)
}
func Bridging_Server2Client_setNoTalkTime(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_55 := Server2Client_setNoTalkTime{}
	err := _55.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.setNoTalkTime(_55.t)
}
func Bridging_Server2Client_addNpc(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_56 := Server2Client_addNpc{}
	err := _56.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.addNpc(_56.npcList)
}
func Bridging_Server2Client_delNpc(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_57 := Server2Client_delNpc{}
	err := _57.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.delNpc(_57.npcList)
}
func Bridging_Server2Client_setTeamLeader(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_58 := Server2Client_setTeamLeader{}
	err := _58.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.setTeamLeader(_58.playerId, _58.isLeader)
}
func Bridging_Server2Client_initBag(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_59 := Server2Client_initBag{}
	err := _59.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.initBag(_59.items)
}
func Bridging_Server2Client_addBagItem(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_60 := Server2Client_addBagItem{}
	err := _60.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.addBagItem(_60.item)
}
func Bridging_Server2Client_delBagItem(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_61 := Server2Client_delBagItem{}
	err := _61.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.delBagItem(_61.slot)
}
func Bridging_Server2Client_updateBagItem(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_62 := Server2Client_updateBagItem{}
	err := _62.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.updateBagItem(_62.item)
}
func Bridging_Server2Client_depositItemOK(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_63 := Server2Client_depositItemOK{}
	err := _63.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.depositItemOK(_63.item)
}
func Bridging_Server2Client_getoutItemOK(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_64 := Server2Client_getoutItemOK{}
	err := _64.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.getoutItemOK(_64.slot)
}
func Bridging_Server2Client_depositBabyOK(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_65 := Server2Client_depositBabyOK{}
	err := _65.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.depositBabyOK(_65.baby)
}
func Bridging_Server2Client_getoutBabyOK(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_66 := Server2Client_getoutBabyOK{}
	err := _66.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.getoutBabyOK(_66.slot)
}
func Bridging_Server2Client_sortItemStorageOK(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_67 := Server2Client_sortItemStorageOK{}
	err := _67.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.sortItemStorageOK(_67.items)
}
func Bridging_Server2Client_sortBabyStorageOK(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_68 := Server2Client_sortBabyStorageOK{}
	err := _68.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.sortBabyStorageOK(_68.babys)
}
func Bridging_Server2Client_initItemStorage(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_69 := Server2Client_initItemStorage{}
	err := _69.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.initItemStorage(_69.gridNum, _69.items)
}
func Bridging_Server2Client_initBabyStorage(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_70 := Server2Client_initBabyStorage{}
	err := _70.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.initBabyStorage(_70.gridNum, _70.babys)
}
func Bridging_Server2Client_openStorageGrid(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_71 := Server2Client_openStorageGrid{}
	err := _71.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.openStorageGrid(_71.tp, _71.gridNum)
}
func Bridging_Server2Client_delStorageBabyOK(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_72 := Server2Client_delStorageBabyOK{}
	err := _72.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.delStorageBabyOK(_72.slot)
}
func Bridging_Server2Client_initPlayerEquips(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_73 := Server2Client_initPlayerEquips{}
	err := _73.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.initPlayerEquips(_73.equips)
}
func Bridging_Server2Client_wearEquipmentOk(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_74 := Server2Client_wearEquipmentOk{}
	err := _74.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.wearEquipmentOk(_74.target, _74.equip)
}
func Bridging_Server2Client_scenePlayerWearEquipment(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_75 := Server2Client_scenePlayerWearEquipment{}
	err := _75.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.scenePlayerWearEquipment(_75.target, _75.itemId)
}
func Bridging_Server2Client_delEquipmentOk(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_76 := Server2Client_delEquipmentOk{}
	err := _76.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.delEquipmentOk(_76.target, _76.itemInstId)
}
func Bridging_Server2Client_scenePlayerDoffEquipment(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_77 := Server2Client_scenePlayerDoffEquipment{}
	err := _77.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.scenePlayerDoffEquipment(_77.target, _77.itemId)
}
func Bridging_Server2Client_sortBagItemOk(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.sortBagItemOk()
}
func Bridging_Server2Client_jointLobbyOk(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_79 := Server2Client_jointLobbyOk{}
	err := _79.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.jointLobbyOk(_79.infos)
}
func Bridging_Server2Client_exitLobbyOk(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.exitLobbyOk()
}
func Bridging_Server2Client_syncDelLobbyTeam(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_81 := Server2Client_syncDelLobbyTeam{}
	err := _81.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.syncDelLobbyTeam(_81.teamId)
}
func Bridging_Server2Client_syncUpdateLobbyTeam(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_82 := Server2Client_syncUpdateLobbyTeam{}
	err := _82.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.syncUpdateLobbyTeam(_82.info)
}
func Bridging_Server2Client_syncAddLobbyTeam(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_83 := Server2Client_syncAddLobbyTeam{}
	err := _83.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.syncAddLobbyTeam(_83.team)
}
func Bridging_Server2Client_createTeamOk(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_84 := Server2Client_createTeamOk{}
	err := _84.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.createTeamOk(_84.team)
}
func Bridging_Server2Client_changeTeamOk(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_85 := Server2Client_changeTeamOk{}
	err := _85.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.changeTeamOk(_85.team)
}
func Bridging_Server2Client_joinTeamOk(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_86 := Server2Client_joinTeamOk{}
	err := _86.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.joinTeamOk(_86.team)
}
func Bridging_Server2Client_addTeamMember(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_87 := Server2Client_addTeamMember{}
	err := _87.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.addTeamMember(_87.info)
}
func Bridging_Server2Client_delTeamMember(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_88 := Server2Client_delTeamMember{}
	err := _88.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.delTeamMember(_88.instId)
}
func Bridging_Server2Client_changeTeamLeaderOk(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_89 := Server2Client_changeTeamLeaderOk{}
	err := _89.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.changeTeamLeaderOk(_89.uuid)
}
func Bridging_Server2Client_exitTeamOk(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_90 := Server2Client_exitTeamOk{}
	err := _90.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.exitTeamOk(_90.iskick)
}
func Bridging_Server2Client_updateTeam(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_91 := Server2Client_updateTeam{}
	err := _91.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.updateTeam(_91.team)
}
func Bridging_Server2Client_joinTeamRoomOK(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_92 := Server2Client_joinTeamRoomOK{}
	err := _92.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.joinTeamRoomOK(_92.team)
}
func Bridging_Server2Client_inviteJoinTeam(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_93 := Server2Client_inviteJoinTeam{}
	err := _93.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.inviteJoinTeam(_93.teamId, _93.name)
}
func Bridging_Server2Client_syncTeamDirtyProp(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_94 := Server2Client_syncTeamDirtyProp{}
	err := _94.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.syncTeamDirtyProp(_94.guid, _94.props)
}
func Bridging_Server2Client_leaveTeamOk(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_95 := Server2Client_leaveTeamOk{}
	err := _95.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.leaveTeamOk(_95.playerId)
}
func Bridging_Server2Client_backTeamOK(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_96 := Server2Client_backTeamOK{}
	err := _96.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.backTeamOK(_96.playerId)
}
func Bridging_Server2Client_teamCallMemberBack(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.teamCallMemberBack()
}
func Bridging_Server2Client_refuseBackTeamOk(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_98 := Server2Client_refuseBackTeamOk{}
	err := _98.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.refuseBackTeamOk(_98.playerId)
}
func Bridging_Server2Client_requestJoinTeamTranspond(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_99 := Server2Client_requestJoinTeamTranspond{}
	err := _99.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.requestJoinTeamTranspond(_99.reqName)
}
func Bridging_Server2Client_drawLotteryBoxRep(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_100 := Server2Client_drawLotteryBoxRep{}
	err := _100.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.drawLotteryBoxRep(_100.items)
}
func Bridging_Server2Client_addEmployee(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_101 := Server2Client_addEmployee{}
	err := _101.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.addEmployee(_101.employee)
}
func Bridging_Server2Client_battleEmployee(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_102 := Server2Client_battleEmployee{}
	err := _102.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.battleEmployee(_102.empId, _102.group, _102.forbattle)
}
func Bridging_Server2Client_changeEmpBattleGroupOK(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_103 := Server2Client_changeEmpBattleGroupOK{}
	err := _103.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.changeEmpBattleGroupOK(_103.group)
}
func Bridging_Server2Client_evolveOK(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_104 := Server2Client_evolveOK{}
	err := _104.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.evolveOK(_104.guid, _104.qc)
}
func Bridging_Server2Client_upStarOK(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_105 := Server2Client_upStarOK{}
	err := _105.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.upStarOK(_105.guid, _105.star, _105.sk)
}
func Bridging_Server2Client_delEmployeeOK(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_106 := Server2Client_delEmployeeOK{}
	err := _106.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.delEmployeeOK(_106.instids)
}
func Bridging_Server2Client_sycnEmployeeSoul(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_107 := Server2Client_sycnEmployeeSoul{}
	err := _107.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.sycnEmployeeSoul(_107.guid, _107.soulNum)
}
func Bridging_Server2Client_initQuest(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_108 := Server2Client_initQuest{}
	err := _108.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.initQuest(_108.qlist, _108.clist)
}
func Bridging_Server2Client_acceptQuestOk(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_109 := Server2Client_acceptQuestOk{}
	err := _109.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.acceptQuestOk(_109.inst)
}
func Bridging_Server2Client_submitQuestOk(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_110 := Server2Client_submitQuestOk{}
	err := _110.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.submitQuestOk(_110.questId)
}
func Bridging_Server2Client_giveupQuestOk(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_111 := Server2Client_giveupQuestOk{}
	err := _111.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.giveupQuestOk(_111.questId)
}
func Bridging_Server2Client_updateQuestInst(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_112 := Server2Client_updateQuestInst{}
	err := _112.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.updateQuestInst(_112.inst)
}
func Bridging_Server2Client_requestContactInfoOk(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_113 := Server2Client_requestContactInfoOk{}
	err := _113.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.requestContactInfoOk(_113.contact)
}
func Bridging_Server2Client_addFriendOK(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_114 := Server2Client_addFriendOK{}
	err := _114.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.addFriendOK(_114.inst)
}
func Bridging_Server2Client_delFriendOK(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_115 := Server2Client_delFriendOK{}
	err := _115.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.delFriendOK(_115.instId)
}
func Bridging_Server2Client_addBlacklistOK(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_116 := Server2Client_addBlacklistOK{}
	err := _116.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.addBlacklistOK(_116.inst)
}
func Bridging_Server2Client_delBlacklistOK(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_117 := Server2Client_delBlacklistOK{}
	err := _117.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.delBlacklistOK(_117.instId)
}
func Bridging_Server2Client_findFriendFail(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.findFriendFail()
}
func Bridging_Server2Client_referrFriendOK(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_119 := Server2Client_referrFriendOK{}
	err := _119.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.referrFriendOK(_119.insts)
}
func Bridging_Server2Client_requestFriendListOK(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_120 := Server2Client_requestFriendListOK{}
	err := _120.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.requestFriendListOK(_120.insts)
}
func Bridging_Server2Client_lotteryOk(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_121 := Server2Client_lotteryOk{}
	err := _121.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.lotteryOk(_121.lotteryId, _121.dropItem)
}
func Bridging_Server2Client_openGatherOK(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_122 := Server2Client_openGatherOK{}
	err := _122.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.openGatherOK(_122.gather)
}
func Bridging_Server2Client_miningOk(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_123 := Server2Client_miningOk{}
	err := _123.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.miningOk(_123.items, _123.gather, _123.gatherNum)
}
func Bridging_Server2Client_openCompound(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_124 := Server2Client_openCompound{}
	err := _124.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.openCompound(_124.compoundId)
}
func Bridging_Server2Client_compoundItemOk(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_125 := Server2Client_compoundItemOk{}
	err := _125.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.compoundItemOk(_125.item)
}
func Bridging_Server2Client_openBagGridOk(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_126 := Server2Client_openBagGridOk{}
	err := _126.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.openBagGridOk(_126.num)
}
func Bridging_Server2Client_requestChallengeOK(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_127 := Server2Client_requestChallengeOK{}
	err := _127.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.requestChallengeOK(_127.isOK)
}
func Bridging_Server2Client_requestMySelfJJCDataOK(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_128 := Server2Client_requestMySelfJJCDataOK{}
	err := _128.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.requestMySelfJJCDataOK(_128.info)
}
func Bridging_Server2Client_requestRivalOK(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_129 := Server2Client_requestRivalOK{}
	err := _129.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.requestRivalOK(_129.infos)
}
func Bridging_Server2Client_rivalTimeOK(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.rivalTimeOK()
}
func Bridging_Server2Client_checkMsgOK(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_131 := Server2Client_checkMsgOK{}
	err := _131.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.checkMsgOK(_131.inst)
}
func Bridging_Server2Client_requestMyAllbattleMsgOK(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_132 := Server2Client_requestMyAllbattleMsgOK{}
	err := _132.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.requestMyAllbattleMsgOK(_132.infos)
}
func Bridging_Server2Client_myBattleMsgOK(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_133 := Server2Client_myBattleMsgOK{}
	err := _133.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.myBattleMsgOK(_133.info)
}
func Bridging_Server2Client_requestJJCRankOK(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_134 := Server2Client_requestJJCRankOK{}
	err := _134.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.requestJJCRankOK(_134.myRank, _134.infos)
}
func Bridging_Server2Client_requestLevelRankOK(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_135 := Server2Client_requestLevelRankOK{}
	err := _135.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.requestLevelRankOK(_135.myRank, _135.infos)
}
func Bridging_Server2Client_requestBabyRankOK(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_136 := Server2Client_requestBabyRankOK{}
	err := _136.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.requestBabyRankOK(_136.myRank, _136.infos)
}
func Bridging_Server2Client_requestEmpRankOK(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_137 := Server2Client_requestEmpRankOK{}
	err := _137.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.requestEmpRankOK(_137.myRank, _137.infos)
}
func Bridging_Server2Client_requestPlayerFFRankOK(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_138 := Server2Client_requestPlayerFFRankOK{}
	err := _138.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.requestPlayerFFRankOK(_138.myRank, _138.infos)
}
func Bridging_Server2Client_queryOnlinePlayerOK(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_139 := Server2Client_queryOnlinePlayerOK{}
	err := _139.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.queryOnlinePlayerOK(_139.isOnline)
}
func Bridging_Server2Client_queryPlayerOK(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_140 := Server2Client_queryPlayerOK{}
	err := _140.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.queryPlayerOK(_140.inst)
}
func Bridging_Server2Client_queryBabyOK(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_141 := Server2Client_queryBabyOK{}
	err := _141.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.queryBabyOK(_141.inst)
}
func Bridging_Server2Client_queryEmployeeOK(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_142 := Server2Client_queryEmployeeOK{}
	err := _142.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.queryEmployeeOK(_142.inst)
}
func Bridging_Server2Client_initGuide(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_143 := Server2Client_initGuide{}
	err := _143.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.initGuide(_143.guideMask)
}
func Bridging_Server2Client_buyShopItemOk(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_144 := Server2Client_buyShopItemOk{}
	err := _144.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.buyShopItemOk(_144.id)
}
func Bridging_Server2Client_addPlayerTitle(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_145 := Server2Client_addPlayerTitle{}
	err := _145.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.addPlayerTitle(_145.title)
}
func Bridging_Server2Client_delPlayerTitle(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_146 := Server2Client_delPlayerTitle{}
	err := _146.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.delPlayerTitle(_146.title)
}
func Bridging_Server2Client_requestOpenBuyBox(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_147 := Server2Client_requestOpenBuyBox{}
	err := _147.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.requestOpenBuyBox(_147.greenTime, _147.blueTime, _147.greenFreeNum)
}
func Bridging_Server2Client_requestGreenBoxTimeOk(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.requestGreenBoxTimeOk()
}
func Bridging_Server2Client_requestBlueBoxTimeOk(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.requestBlueBoxTimeOk()
}
func Bridging_Server2Client_updateAchievementinfo(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_150 := Server2Client_updateAchievementinfo{}
	err := _150.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.updateAchievementinfo(_150.achs)
}
func Bridging_Server2Client_syncOpenSystemFlag(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_151 := Server2Client_syncOpenSystemFlag{}
	err := _151.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.syncOpenSystemFlag(_151.flag)
}
func Bridging_Server2Client_requestActivityRewardOK(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_152 := Server2Client_requestActivityRewardOK{}
	err := _152.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.requestActivityRewardOK(_152.ar)
}
func Bridging_Server2Client_syncActivity(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_153 := Server2Client_syncActivity{}
	err := _153.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.syncActivity(_153.table)
}
func Bridging_Server2Client_updateActivityStatus(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_154 := Server2Client_updateActivityStatus{}
	err := _154.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.updateActivityStatus(_154.type_, _154.open)
}
func Bridging_Server2Client_updateActivityCounter(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_155 := Server2Client_updateActivityCounter{}
	err := _155.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.updateActivityCounter(_155.type_, _155.counter, _155.reward)
}
func Bridging_Server2Client_syncExam(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_156 := Server2Client_syncExam{}
	err := _156.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.syncExam(_156.exam)
}
func Bridging_Server2Client_syncExamAnswer(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_157 := Server2Client_syncExamAnswer{}
	err := _157.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.syncExamAnswer(_157.answer)
}
func Bridging_Server2Client_petActivityNoNum(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_158 := Server2Client_petActivityNoNum{}
	err := _158.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.petActivityNoNum(_158.name)
}
func Bridging_Server2Client_syncHundredInfo(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_159 := Server2Client_syncHundredInfo{}
	err := _159.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.syncHundredInfo(_159.hb)
}
func Bridging_Server2Client_initSignUp(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_160 := Server2Client_initSignUp{}
	err := _160.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.initSignUp(_160.info, _160.process, _160.sign7, _160.sign14, _160.sign28)
}
func Bridging_Server2Client_signUp(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_161 := Server2Client_signUp{}
	err := _161.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.signUp(_161.flag)
}
func Bridging_Server2Client_requestSignupRewardOk7(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.requestSignupRewardOk7()
}
func Bridging_Server2Client_requestSignupRewardOk14(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.requestSignupRewardOk14()
}
func Bridging_Server2Client_requestSignupRewardOk28(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.requestSignupRewardOk28()
}
func Bridging_Server2Client_sycnDoubleExpTime(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_165 := Server2Client_sycnDoubleExpTime{}
	err := _165.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.sycnDoubleExpTime(_165.isFlag, _165.times)
}
func Bridging_Server2Client_sycnStates(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_166 := Server2Client_sycnStates{}
	err := _166.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.sycnStates(_166.states)
}
func Bridging_Server2Client_requestpvprankOK(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_167 := Server2Client_requestpvprankOK{}
	err := _167.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.requestpvprankOK(_167.infos)
}
func Bridging_Server2Client_syncMyJJCTeamMember(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.syncMyJJCTeamMember()
}
func Bridging_Server2Client_startMatchingOK(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.startMatchingOK()
}
func Bridging_Server2Client_stopMatchingOK(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_170 := Server2Client_stopMatchingOK{}
	err := _170.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.stopMatchingOK(_170.times)
}
func Bridging_Server2Client_updatePvpJJCinfo(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_171 := Server2Client_updatePvpJJCinfo{}
	err := _171.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.updatePvpJJCinfo(_171.info)
}
func Bridging_Server2Client_exitPvpJJCOk(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.exitPvpJJCOk()
}
func Bridging_Server2Client_syncEnemyPvpJJCPlayerInfo(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_173 := Server2Client_syncEnemyPvpJJCPlayerInfo{}
	err := _173.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.syncEnemyPvpJJCPlayerInfo(_173.info)
}
func Bridging_Server2Client_syncEnemyPvpJJCTeamInfo(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_174 := Server2Client_syncEnemyPvpJJCTeamInfo{}
	err := _174.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.syncEnemyPvpJJCTeamInfo(_174.infos, _174.teamID_)
}
func Bridging_Server2Client_openWarriorchooseUI(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.openWarriorchooseUI()
}
func Bridging_Server2Client_warriorStartOK(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.warriorStartOK()
}
func Bridging_Server2Client_warriorStopOK(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.warriorStopOK()
}
func Bridging_Server2Client_syncWarriorEnemyTeamInfo(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_178 := Server2Client_syncWarriorEnemyTeamInfo{}
	err := _178.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.syncWarriorEnemyTeamInfo(_178.infos, _178.teamID_)
}
func Bridging_Server2Client_appendMail(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_179 := Server2Client_appendMail{}
	err := _179.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.appendMail(_179.mails)
}
func Bridging_Server2Client_delMail(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_180 := Server2Client_delMail{}
	err := _180.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.delMail(_180.mailId)
}
func Bridging_Server2Client_updateMailOk(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_181 := Server2Client_updateMailOk{}
	err := _181.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.updateMailOk(_181.mail)
}
func Bridging_Server2Client_boardcastNotice(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_182 := Server2Client_boardcastNotice{}
	err := _182.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.boardcastNotice(_182.content, _182.isGm)
}
func Bridging_Server2Client_createGuildOK(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.createGuildOK()
}
func Bridging_Server2Client_delGuildOK(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.delGuildOK()
}
func Bridging_Server2Client_leaveGuildOk(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_185 := Server2Client_leaveGuildOk{}
	err := _185.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.leaveGuildOk(_185.who, _185.isKick)
}
func Bridging_Server2Client_initGuildData(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_186 := Server2Client_initGuildData{}
	err := _186.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.initGuildData(_186.guild)
}
func Bridging_Server2Client_initGuildMemberList(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_187 := Server2Client_initGuildMemberList{}
	err := _187.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.initGuildMemberList(_187.member)
}
func Bridging_Server2Client_modifyGuildMemberList(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_188 := Server2Client_modifyGuildMemberList{}
	err := _188.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.modifyGuildMemberList(_188.member, _188.flag)
}
func Bridging_Server2Client_modifyGuildList(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_189 := Server2Client_modifyGuildList{}
	err := _189.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.modifyGuildList(_189.data, _189.flag)
}
func Bridging_Server2Client_queryGuildListResult(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_190 := Server2Client_queryGuildListResult{}
	err := _190.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.queryGuildListResult(_190.page, _190.pageNum, _190.guildList)
}
func Bridging_Server2Client_inviteGuild(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_191 := Server2Client_inviteGuild{}
	err := _191.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.inviteGuild(_191.sendName, _191.guildName)
}
func Bridging_Server2Client_updateGuildShopItems(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_192 := Server2Client_updateGuildShopItems{}
	err := _192.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.updateGuildShopItems(_192.items)
}
func Bridging_Server2Client_updateGuildBuilding(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_193 := Server2Client_updateGuildBuilding{}
	err := _193.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.updateGuildBuilding(_193.type_, _193.building)
}
func Bridging_Server2Client_updateGuildMyMember(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_194 := Server2Client_updateGuildMyMember{}
	err := _194.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.updateGuildMyMember(_194.member)
}
func Bridging_Server2Client_levelupGuildSkillOk(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_195 := Server2Client_levelupGuildSkillOk{}
	err := _195.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.levelupGuildSkillOk(_195.skInst)
}
func Bridging_Server2Client_presentGuildItemOk(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_196 := Server2Client_presentGuildItemOk{}
	err := _196.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.presentGuildItemOk(_196.val)
}
func Bridging_Server2Client_progenitusAddExpOk(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_197 := Server2Client_progenitusAddExpOk{}
	err := _197.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.progenitusAddExpOk(_197.mInst)
}
func Bridging_Server2Client_setProgenitusPositionOk(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_198 := Server2Client_setProgenitusPositionOk{}
	err := _198.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.setProgenitusPositionOk(_198.positions)
}
func Bridging_Server2Client_updateGuildFundz(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_199 := Server2Client_updateGuildFundz{}
	err := _199.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.updateGuildFundz(_199.val)
}
func Bridging_Server2Client_updateGuildMemberContribution(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_200 := Server2Client_updateGuildMemberContribution{}
	err := _200.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.updateGuildMemberContribution(_200.val)
}
func Bridging_Server2Client_openGuildBattle(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_201 := Server2Client_openGuildBattle{}
	err := _201.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.openGuildBattle(_201.otherName, _201.playerNum, _201.level, _201.isLeft, _201.lstime)
}
func Bridging_Server2Client_startGuildBattle(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_202 := Server2Client_startGuildBattle{}
	err := _202.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.startGuildBattle(_202.otherName, _202.otherCon, _202.selfCon)
}
func Bridging_Server2Client_closeGuildBattle(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_203 := Server2Client_closeGuildBattle{}
	err := _203.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.closeGuildBattle(_203.isWinner)
}
func Bridging_Server2Client_syncGuildBattleWinCount(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_204 := Server2Client_syncGuildBattleWinCount{}
	err := _204.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.syncGuildBattleWinCount(_204.myWin, _204.otherWin)
}
func Bridging_Server2Client_initMySelling(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_205 := Server2Client_initMySelling{}
	err := _205.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.initMySelling(_205.items)
}
func Bridging_Server2Client_initMySelled(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_206 := Server2Client_initMySelled{}
	err := _206.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.initMySelled(_206.items)
}
func Bridging_Server2Client_fetchSellingOk(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_207 := Server2Client_fetchSellingOk{}
	err := _207.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.fetchSellingOk(_207.items, _207.total)
}
func Bridging_Server2Client_fetchSellingOk2(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_208 := Server2Client_fetchSellingOk2{}
	err := _208.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.fetchSellingOk2(_208.items, _208.total)
}
func Bridging_Server2Client_sellingOk(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_209 := Server2Client_sellingOk{}
	err := _209.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.sellingOk(_209.item)
}
func Bridging_Server2Client_selledOk(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_210 := Server2Client_selledOk{}
	err := _210.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.selledOk(_210.item)
}
func Bridging_Server2Client_unsellingOk(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_211 := Server2Client_unsellingOk{}
	err := _211.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.unsellingOk(_211.sellid)
}
func Bridging_Server2Client_redemptionSpreeOk(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.redemptionSpreeOk()
}
func Bridging_Server2Client_insertState(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_213 := Server2Client_insertState{}
	err := _213.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.insertState(_213.st)
}
func Bridging_Server2Client_updattState(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_214 := Server2Client_updattState{}
	err := _214.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.updattState(_214.st)
}
func Bridging_Server2Client_removeState(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_215 := Server2Client_removeState{}
	err := _215.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.removeState(_215.stid)
}
func Bridging_Server2Client_requestFixItemOk(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_216 := Server2Client_requestFixItemOk{}
	err := _216.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.requestFixItemOk(_216.item)
}
func Bridging_Server2Client_makeDebirsItemOK(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.makeDebirsItemOK()
}
func Bridging_Server2Client_updateMagicItem(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_218 := Server2Client_updateMagicItem{}
	err := _218.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.updateMagicItem(_218.level, _218.exp)
}
func Bridging_Server2Client_changeMagicJobOk(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_219 := Server2Client_changeMagicJobOk{}
	err := _219.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.changeMagicJobOk(_219.job)
}
func Bridging_Server2Client_magicItemTupoOk(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_220 := Server2Client_magicItemTupoOk{}
	err := _220.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.magicItemTupoOk(_220.level)
}
func Bridging_Server2Client_zhuanpanOK(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_221 := Server2Client_zhuanpanOK{}
	err := _221.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.zhuanpanOK(_221.pond)
}
func Bridging_Server2Client_updateZhuanpanNotice(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_222 := Server2Client_updateZhuanpanNotice{}
	err := _222.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.updateZhuanpanNotice(_222.zhuanp)
}
func Bridging_Server2Client_sycnZhuanpanData(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_223 := Server2Client_sycnZhuanpanData{}
	err := _223.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.sycnZhuanpanData(_223.data)
}
func Bridging_Server2Client_copynonum(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_224 := Server2Client_copynonum{}
	err := _224.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.copynonum(_224.name)
}
func Bridging_Server2Client_sceneFilterOk(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_225 := Server2Client_sceneFilterOk{}
	err := _225.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.sceneFilterOk(_225.sfType)
}
func Bridging_Server2Client_wishingOK(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.wishingOK()
}
func Bridging_Server2Client_shareWishOK(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_227 := Server2Client_shareWishOK{}
	err := _227.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.shareWishOK(_227.wish)
}
func Bridging_Server2Client_leaderCloseDialogOk(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.leaderCloseDialogOk()
}
func Bridging_Server2Client_startOnlineTime(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.startOnlineTime()
}
func Bridging_Server2Client_requestOnlineTimeRewardOK(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_230 := Server2Client_requestOnlineTimeRewardOK{}
	err := _230.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.requestOnlineTimeRewardOK(_230.index)
}
func Bridging_Server2Client_sycnVipflag(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_231 := Server2Client_sycnVipflag{}
	err := _231.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.sycnVipflag(_231.flag)
}
func Bridging_Server2Client_buyFundOK(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_232 := Server2Client_buyFundOK{}
	err := _232.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.buyFundOK(_232.flag)
}
func Bridging_Server2Client_requestFundRewardOK(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_233 := Server2Client_requestFundRewardOK{}
	err := _233.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.requestFundRewardOK(_233.level)
}
func Bridging_Server2Client_firstRechargeOK(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_234 := Server2Client_firstRechargeOK{}
	err := _234.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.firstRechargeOK(_234.isFlag)
}
func Bridging_Server2Client_firstRechargeGiftOK(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_235 := Server2Client_firstRechargeGiftOK{}
	err := _235.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.firstRechargeGiftOK(_235.isFlag)
}
func Bridging_Server2Client_agencyActivity(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_236 := Server2Client_agencyActivity{}
	err := _236.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.agencyActivity(_236.type_, _236.isFlag)
}
func Bridging_Server2Client_updateFestival(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_237 := Server2Client_updateFestival{}
	err := _237.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.updateFestival(_237.festival)
}
func Bridging_Server2Client_updateSelfRecharge(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_238 := Server2Client_updateSelfRecharge{}
	err := _238.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.updateSelfRecharge(_238.val)
}
func Bridging_Server2Client_updateSysRecharge(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_239 := Server2Client_updateSysRecharge{}
	err := _239.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.updateSysRecharge(_239.val)
}
func Bridging_Server2Client_updateSelfDiscountStore(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_240 := Server2Client_updateSelfDiscountStore{}
	err := _240.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.updateSelfDiscountStore(_240.val)
}
func Bridging_Server2Client_updateSysDiscountStore(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_241 := Server2Client_updateSysDiscountStore{}
	err := _241.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.updateSysDiscountStore(_241.val)
}
func Bridging_Server2Client_updateSelfOnceRecharge(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_242 := Server2Client_updateSelfOnceRecharge{}
	err := _242.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.updateSelfOnceRecharge(_242.val)
}
func Bridging_Server2Client_updateSysOnceRecharge(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_243 := Server2Client_updateSysOnceRecharge{}
	err := _243.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.updateSysOnceRecharge(_243.val)
}
func Bridging_Server2Client_openCardOK(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_244 := Server2Client_openCardOK{}
	err := _244.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.openCardOK(_244.data)
}
func Bridging_Server2Client_resetCardOK(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.resetCardOK()
}
func Bridging_Server2Client_sycnHotRole(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_246 := Server2Client_sycnHotRole{}
	err := _246.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.sycnHotRole(_246.data)
}
func Bridging_Server2Client_hotRoleBuyOk(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_247 := Server2Client_hotRoleBuyOk{}
	err := _247.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.hotRoleBuyOk(_247.buyNum)
}
func Bridging_Server2Client_updateSevenday(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_248 := Server2Client_updateSevenday{}
	err := _248.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.updateSevenday(_248.data)
}
func Bridging_Server2Client_updateEmployeeActivity(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_249 := Server2Client_updateEmployeeActivity{}
	err := _249.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.updateEmployeeActivity(_249.data)
}
func Bridging_Server2Client_updateMinGiftActivity(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_250 := Server2Client_updateMinGiftActivity{}
	err := _250.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.updateMinGiftActivity(_250.data)
}
func Bridging_Server2Client_updateIntegralShop(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_251 := Server2Client_updateIntegralShop{}
	err := _251.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.updateIntegralShop(_251.data)
}
func Bridging_Server2Client_updateShowBaby(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_252 := Server2Client_updateShowBaby{}
	err := _252.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.updateShowBaby(_252.playerId, _252.showBabyTableId, _252.showBabyName)
}
func Bridging_Server2Client_updateMySelfRecharge(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_253 := Server2Client_updateMySelfRecharge{}
	err := _253.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.updateMySelfRecharge(_253.val)
}
func Bridging_Server2Client_verificationSMSOk(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_254 := Server2Client_verificationSMSOk{}
	err := _254.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.verificationSMSOk(_254.phoneNumber)
}
func Bridging_Server2Client_requestLevelGiftOK(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_255 := Server2Client_requestLevelGiftOK{}
	err := _255.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.requestLevelGiftOK(_255.level)
}
func Bridging_Server2Client_sycnConvertExp(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_256 := Server2Client_sycnConvertExp{}
	err := _256.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.sycnConvertExp(_256.val)
}
func Bridging_Server2Client_wearFuwenOk(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_257 := Server2Client_wearFuwenOk{}
	err := _257.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.wearFuwenOk(_257.inst)
}
func Bridging_Server2Client_takeoffFuwenOk(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_258 := Server2Client_takeoffFuwenOk{}
	err := _258.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.takeoffFuwenOk(_258.slot)
}
func Bridging_Server2Client_compFuwenOk(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.compFuwenOk()
}
func Bridging_Server2Client_requestEmployeeQuestOk(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_260 := Server2Client_requestEmployeeQuestOk{}
	err := _260.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.requestEmployeeQuestOk(_260.questList)
}
func Bridging_Server2Client_acceptEmployeeQuestOk(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_261 := Server2Client_acceptEmployeeQuestOk{}
	err := _261.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.acceptEmployeeQuestOk(_261.inst)
}
func Bridging_Server2Client_submitEmployeeQuestOk(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_262 := Server2Client_submitEmployeeQuestOk{}
	err := _262.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.submitEmployeeQuestOk(_262.questId, _262.isSuccess)
}
func Bridging_Server2Client_sycnCrystal(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_263 := Server2Client_sycnCrystal{}
	err := _263.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.sycnCrystal(_263.data)
}
func Bridging_Server2Client_crystalUpLeveResult(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_264 := Server2Client_crystalUpLeveResult{}
	err := _264.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.crystalUpLeveResult(_264.isOK)
}
func Bridging_Server2Client_resetCrystalPropOK(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.resetCrystalPropOK()
}
func Bridging_Server2Client_sycnCourseGift(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_266 := Server2Client_sycnCourseGift{}
	err := _266.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.sycnCourseGift(_266.data)
}
func Bridging_Server2Client_orderOk(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_267 := Server2Client_orderOk{}
	err := _267.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.orderOk(_267.orderId, _267.shopId)
}
func Bridging_Server2Client_updateRandSubmitQuestCount(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_268 := Server2Client_updateRandSubmitQuestCount{}
	err := _268.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.updateRandSubmitQuestCount(_268.submitCount)
}
func Server2ClientDispatch(buffer *bytes.Buffer, p Server2ClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	pid := uint16(0XFFFF)
	err := prpc.Read(buffer, &pid)
	if err != nil {
		return err
	}
	switch pid {
	case 0:
		return Bridging_Server2Client_pong(buffer, p)
	case 1:
		return Bridging_Server2Client_errorno(buffer, p)
	case 2:
		return Bridging_Server2Client_teamerrorno(buffer, p)
	case 3:
		return Bridging_Server2Client_reconnection(buffer, p)
	case 4:
		return Bridging_Server2Client_sessionfailed(buffer, p)
	case 5:
		return Bridging_Server2Client_loginok(buffer, p)
	case 6:
		return Bridging_Server2Client_logoutOk(buffer, p)
	case 7:
		return Bridging_Server2Client_createPlayerOk(buffer, p)
	case 8:
		return Bridging_Server2Client_deletePlayerOk(buffer, p)
	case 9:
		return Bridging_Server2Client_enterGameOk(buffer, p)
	case 10:
		return Bridging_Server2Client_initBabies(buffer, p)
	case 11:
		return Bridging_Server2Client_initEmployees(buffer, p)
	case 12:
		return Bridging_Server2Client_initEmpBattleGroup(buffer, p)
	case 13:
		return Bridging_Server2Client_initNpc(buffer, p)
	case 14:
		return Bridging_Server2Client_initAchievement(buffer, p)
	case 15:
		return Bridging_Server2Client_initGather(buffer, p)
	case 16:
		return Bridging_Server2Client_initcompound(buffer, p)
	case 17:
		return Bridging_Server2Client_addBaby(buffer, p)
	case 18:
		return Bridging_Server2Client_refreshBaby(buffer, p)
	case 19:
		return Bridging_Server2Client_delBabyOK(buffer, p)
	case 20:
		return Bridging_Server2Client_changeBabyNameOK(buffer, p)
	case 21:
		return Bridging_Server2Client_remouldBabyOK(buffer, p)
	case 22:
		return Bridging_Server2Client_intensifyBabyOK(buffer, p)
	case 23:
		return Bridging_Server2Client_learnSkillOk(buffer, p)
	case 24:
		return Bridging_Server2Client_forgetSkillOk(buffer, p)
	case 25:
		return Bridging_Server2Client_addSkillExp(buffer, p)
	case 26:
		return Bridging_Server2Client_babyLearnSkillOK(buffer, p)
	case 27:
		return Bridging_Server2Client_skillLevelUp(buffer, p)
	case 28:
		return Bridging_Server2Client_joinScene(buffer, p)
	case 29:
		return Bridging_Server2Client_joinCopySceneOK(buffer, p)
	case 30:
		return Bridging_Server2Client_initCopyNums(buffer, p)
	case 31:
		return Bridging_Server2Client_addToScene(buffer, p)
	case 32:
		return Bridging_Server2Client_delFormScene(buffer, p)
	case 33:
		return Bridging_Server2Client_move2(buffer, p)
	case 34:
		return Bridging_Server2Client_cantMove(buffer, p)
	case 35:
		return Bridging_Server2Client_querySimplePlayerInstOk(buffer, p)
	case 36:
		return Bridging_Server2Client_transfor2(buffer, p)
	case 37:
		return Bridging_Server2Client_openScene(buffer, p)
	case 38:
		return Bridging_Server2Client_autoBattleResult(buffer, p)
	case 39:
		return Bridging_Server2Client_talked2Npc(buffer, p)
	case 40:
		return Bridging_Server2Client_talked2Player(buffer, p)
	case 41:
		return Bridging_Server2Client_useItemOk(buffer, p)
	case 42:
		return Bridging_Server2Client_syncBattleStatus(buffer, p)
	case 43:
		return Bridging_Server2Client_enterBattleOk(buffer, p)
	case 44:
		return Bridging_Server2Client_exitBattleOk(buffer, p)
	case 45:
		return Bridging_Server2Client_syncOrderOk(buffer, p)
	case 46:
		return Bridging_Server2Client_syncOrderOkEX(buffer, p)
	case 47:
		return Bridging_Server2Client_syncOneTurnAction(buffer, p)
	case 48:
		return Bridging_Server2Client_syncProperties(buffer, p)
	case 49:
		return Bridging_Server2Client_receiveChat(buffer, p)
	case 50:
		return Bridging_Server2Client_requestAudioOk(buffer, p)
	case 51:
		return Bridging_Server2Client_publishItemInstRes(buffer, p)
	case 52:
		return Bridging_Server2Client_queryItemInstRes(buffer, p)
	case 53:
		return Bridging_Server2Client_publishBabyInstRes(buffer, p)
	case 54:
		return Bridging_Server2Client_queryBabyInstRes(buffer, p)
	case 55:
		return Bridging_Server2Client_setNoTalkTime(buffer, p)
	case 56:
		return Bridging_Server2Client_addNpc(buffer, p)
	case 57:
		return Bridging_Server2Client_delNpc(buffer, p)
	case 58:
		return Bridging_Server2Client_setTeamLeader(buffer, p)
	case 59:
		return Bridging_Server2Client_initBag(buffer, p)
	case 60:
		return Bridging_Server2Client_addBagItem(buffer, p)
	case 61:
		return Bridging_Server2Client_delBagItem(buffer, p)
	case 62:
		return Bridging_Server2Client_updateBagItem(buffer, p)
	case 63:
		return Bridging_Server2Client_depositItemOK(buffer, p)
	case 64:
		return Bridging_Server2Client_getoutItemOK(buffer, p)
	case 65:
		return Bridging_Server2Client_depositBabyOK(buffer, p)
	case 66:
		return Bridging_Server2Client_getoutBabyOK(buffer, p)
	case 67:
		return Bridging_Server2Client_sortItemStorageOK(buffer, p)
	case 68:
		return Bridging_Server2Client_sortBabyStorageOK(buffer, p)
	case 69:
		return Bridging_Server2Client_initItemStorage(buffer, p)
	case 70:
		return Bridging_Server2Client_initBabyStorage(buffer, p)
	case 71:
		return Bridging_Server2Client_openStorageGrid(buffer, p)
	case 72:
		return Bridging_Server2Client_delStorageBabyOK(buffer, p)
	case 73:
		return Bridging_Server2Client_initPlayerEquips(buffer, p)
	case 74:
		return Bridging_Server2Client_wearEquipmentOk(buffer, p)
	case 75:
		return Bridging_Server2Client_scenePlayerWearEquipment(buffer, p)
	case 76:
		return Bridging_Server2Client_delEquipmentOk(buffer, p)
	case 77:
		return Bridging_Server2Client_scenePlayerDoffEquipment(buffer, p)
	case 78:
		return Bridging_Server2Client_sortBagItemOk(buffer, p)
	case 79:
		return Bridging_Server2Client_jointLobbyOk(buffer, p)
	case 80:
		return Bridging_Server2Client_exitLobbyOk(buffer, p)
	case 81:
		return Bridging_Server2Client_syncDelLobbyTeam(buffer, p)
	case 82:
		return Bridging_Server2Client_syncUpdateLobbyTeam(buffer, p)
	case 83:
		return Bridging_Server2Client_syncAddLobbyTeam(buffer, p)
	case 84:
		return Bridging_Server2Client_createTeamOk(buffer, p)
	case 85:
		return Bridging_Server2Client_changeTeamOk(buffer, p)
	case 86:
		return Bridging_Server2Client_joinTeamOk(buffer, p)
	case 87:
		return Bridging_Server2Client_addTeamMember(buffer, p)
	case 88:
		return Bridging_Server2Client_delTeamMember(buffer, p)
	case 89:
		return Bridging_Server2Client_changeTeamLeaderOk(buffer, p)
	case 90:
		return Bridging_Server2Client_exitTeamOk(buffer, p)
	case 91:
		return Bridging_Server2Client_updateTeam(buffer, p)
	case 92:
		return Bridging_Server2Client_joinTeamRoomOK(buffer, p)
	case 93:
		return Bridging_Server2Client_inviteJoinTeam(buffer, p)
	case 94:
		return Bridging_Server2Client_syncTeamDirtyProp(buffer, p)
	case 95:
		return Bridging_Server2Client_leaveTeamOk(buffer, p)
	case 96:
		return Bridging_Server2Client_backTeamOK(buffer, p)
	case 97:
		return Bridging_Server2Client_teamCallMemberBack(buffer, p)
	case 98:
		return Bridging_Server2Client_refuseBackTeamOk(buffer, p)
	case 99:
		return Bridging_Server2Client_requestJoinTeamTranspond(buffer, p)
	case 100:
		return Bridging_Server2Client_drawLotteryBoxRep(buffer, p)
	case 101:
		return Bridging_Server2Client_addEmployee(buffer, p)
	case 102:
		return Bridging_Server2Client_battleEmployee(buffer, p)
	case 103:
		return Bridging_Server2Client_changeEmpBattleGroupOK(buffer, p)
	case 104:
		return Bridging_Server2Client_evolveOK(buffer, p)
	case 105:
		return Bridging_Server2Client_upStarOK(buffer, p)
	case 106:
		return Bridging_Server2Client_delEmployeeOK(buffer, p)
	case 107:
		return Bridging_Server2Client_sycnEmployeeSoul(buffer, p)
	case 108:
		return Bridging_Server2Client_initQuest(buffer, p)
	case 109:
		return Bridging_Server2Client_acceptQuestOk(buffer, p)
	case 110:
		return Bridging_Server2Client_submitQuestOk(buffer, p)
	case 111:
		return Bridging_Server2Client_giveupQuestOk(buffer, p)
	case 112:
		return Bridging_Server2Client_updateQuestInst(buffer, p)
	case 113:
		return Bridging_Server2Client_requestContactInfoOk(buffer, p)
	case 114:
		return Bridging_Server2Client_addFriendOK(buffer, p)
	case 115:
		return Bridging_Server2Client_delFriendOK(buffer, p)
	case 116:
		return Bridging_Server2Client_addBlacklistOK(buffer, p)
	case 117:
		return Bridging_Server2Client_delBlacklistOK(buffer, p)
	case 118:
		return Bridging_Server2Client_findFriendFail(buffer, p)
	case 119:
		return Bridging_Server2Client_referrFriendOK(buffer, p)
	case 120:
		return Bridging_Server2Client_requestFriendListOK(buffer, p)
	case 121:
		return Bridging_Server2Client_lotteryOk(buffer, p)
	case 122:
		return Bridging_Server2Client_openGatherOK(buffer, p)
	case 123:
		return Bridging_Server2Client_miningOk(buffer, p)
	case 124:
		return Bridging_Server2Client_openCompound(buffer, p)
	case 125:
		return Bridging_Server2Client_compoundItemOk(buffer, p)
	case 126:
		return Bridging_Server2Client_openBagGridOk(buffer, p)
	case 127:
		return Bridging_Server2Client_requestChallengeOK(buffer, p)
	case 128:
		return Bridging_Server2Client_requestMySelfJJCDataOK(buffer, p)
	case 129:
		return Bridging_Server2Client_requestRivalOK(buffer, p)
	case 130:
		return Bridging_Server2Client_rivalTimeOK(buffer, p)
	case 131:
		return Bridging_Server2Client_checkMsgOK(buffer, p)
	case 132:
		return Bridging_Server2Client_requestMyAllbattleMsgOK(buffer, p)
	case 133:
		return Bridging_Server2Client_myBattleMsgOK(buffer, p)
	case 134:
		return Bridging_Server2Client_requestJJCRankOK(buffer, p)
	case 135:
		return Bridging_Server2Client_requestLevelRankOK(buffer, p)
	case 136:
		return Bridging_Server2Client_requestBabyRankOK(buffer, p)
	case 137:
		return Bridging_Server2Client_requestEmpRankOK(buffer, p)
	case 138:
		return Bridging_Server2Client_requestPlayerFFRankOK(buffer, p)
	case 139:
		return Bridging_Server2Client_queryOnlinePlayerOK(buffer, p)
	case 140:
		return Bridging_Server2Client_queryPlayerOK(buffer, p)
	case 141:
		return Bridging_Server2Client_queryBabyOK(buffer, p)
	case 142:
		return Bridging_Server2Client_queryEmployeeOK(buffer, p)
	case 143:
		return Bridging_Server2Client_initGuide(buffer, p)
	case 144:
		return Bridging_Server2Client_buyShopItemOk(buffer, p)
	case 145:
		return Bridging_Server2Client_addPlayerTitle(buffer, p)
	case 146:
		return Bridging_Server2Client_delPlayerTitle(buffer, p)
	case 147:
		return Bridging_Server2Client_requestOpenBuyBox(buffer, p)
	case 148:
		return Bridging_Server2Client_requestGreenBoxTimeOk(buffer, p)
	case 149:
		return Bridging_Server2Client_requestBlueBoxTimeOk(buffer, p)
	case 150:
		return Bridging_Server2Client_updateAchievementinfo(buffer, p)
	case 151:
		return Bridging_Server2Client_syncOpenSystemFlag(buffer, p)
	case 152:
		return Bridging_Server2Client_requestActivityRewardOK(buffer, p)
	case 153:
		return Bridging_Server2Client_syncActivity(buffer, p)
	case 154:
		return Bridging_Server2Client_updateActivityStatus(buffer, p)
	case 155:
		return Bridging_Server2Client_updateActivityCounter(buffer, p)
	case 156:
		return Bridging_Server2Client_syncExam(buffer, p)
	case 157:
		return Bridging_Server2Client_syncExamAnswer(buffer, p)
	case 158:
		return Bridging_Server2Client_petActivityNoNum(buffer, p)
	case 159:
		return Bridging_Server2Client_syncHundredInfo(buffer, p)
	case 160:
		return Bridging_Server2Client_initSignUp(buffer, p)
	case 161:
		return Bridging_Server2Client_signUp(buffer, p)
	case 162:
		return Bridging_Server2Client_requestSignupRewardOk7(buffer, p)
	case 163:
		return Bridging_Server2Client_requestSignupRewardOk14(buffer, p)
	case 164:
		return Bridging_Server2Client_requestSignupRewardOk28(buffer, p)
	case 165:
		return Bridging_Server2Client_sycnDoubleExpTime(buffer, p)
	case 166:
		return Bridging_Server2Client_sycnStates(buffer, p)
	case 167:
		return Bridging_Server2Client_requestpvprankOK(buffer, p)
	case 168:
		return Bridging_Server2Client_syncMyJJCTeamMember(buffer, p)
	case 169:
		return Bridging_Server2Client_startMatchingOK(buffer, p)
	case 170:
		return Bridging_Server2Client_stopMatchingOK(buffer, p)
	case 171:
		return Bridging_Server2Client_updatePvpJJCinfo(buffer, p)
	case 172:
		return Bridging_Server2Client_exitPvpJJCOk(buffer, p)
	case 173:
		return Bridging_Server2Client_syncEnemyPvpJJCPlayerInfo(buffer, p)
	case 174:
		return Bridging_Server2Client_syncEnemyPvpJJCTeamInfo(buffer, p)
	case 175:
		return Bridging_Server2Client_openWarriorchooseUI(buffer, p)
	case 176:
		return Bridging_Server2Client_warriorStartOK(buffer, p)
	case 177:
		return Bridging_Server2Client_warriorStopOK(buffer, p)
	case 178:
		return Bridging_Server2Client_syncWarriorEnemyTeamInfo(buffer, p)
	case 179:
		return Bridging_Server2Client_appendMail(buffer, p)
	case 180:
		return Bridging_Server2Client_delMail(buffer, p)
	case 181:
		return Bridging_Server2Client_updateMailOk(buffer, p)
	case 182:
		return Bridging_Server2Client_boardcastNotice(buffer, p)
	case 183:
		return Bridging_Server2Client_createGuildOK(buffer, p)
	case 184:
		return Bridging_Server2Client_delGuildOK(buffer, p)
	case 185:
		return Bridging_Server2Client_leaveGuildOk(buffer, p)
	case 186:
		return Bridging_Server2Client_initGuildData(buffer, p)
	case 187:
		return Bridging_Server2Client_initGuildMemberList(buffer, p)
	case 188:
		return Bridging_Server2Client_modifyGuildMemberList(buffer, p)
	case 189:
		return Bridging_Server2Client_modifyGuildList(buffer, p)
	case 190:
		return Bridging_Server2Client_queryGuildListResult(buffer, p)
	case 191:
		return Bridging_Server2Client_inviteGuild(buffer, p)
	case 192:
		return Bridging_Server2Client_updateGuildShopItems(buffer, p)
	case 193:
		return Bridging_Server2Client_updateGuildBuilding(buffer, p)
	case 194:
		return Bridging_Server2Client_updateGuildMyMember(buffer, p)
	case 195:
		return Bridging_Server2Client_levelupGuildSkillOk(buffer, p)
	case 196:
		return Bridging_Server2Client_presentGuildItemOk(buffer, p)
	case 197:
		return Bridging_Server2Client_progenitusAddExpOk(buffer, p)
	case 198:
		return Bridging_Server2Client_setProgenitusPositionOk(buffer, p)
	case 199:
		return Bridging_Server2Client_updateGuildFundz(buffer, p)
	case 200:
		return Bridging_Server2Client_updateGuildMemberContribution(buffer, p)
	case 201:
		return Bridging_Server2Client_openGuildBattle(buffer, p)
	case 202:
		return Bridging_Server2Client_startGuildBattle(buffer, p)
	case 203:
		return Bridging_Server2Client_closeGuildBattle(buffer, p)
	case 204:
		return Bridging_Server2Client_syncGuildBattleWinCount(buffer, p)
	case 205:
		return Bridging_Server2Client_initMySelling(buffer, p)
	case 206:
		return Bridging_Server2Client_initMySelled(buffer, p)
	case 207:
		return Bridging_Server2Client_fetchSellingOk(buffer, p)
	case 208:
		return Bridging_Server2Client_fetchSellingOk2(buffer, p)
	case 209:
		return Bridging_Server2Client_sellingOk(buffer, p)
	case 210:
		return Bridging_Server2Client_selledOk(buffer, p)
	case 211:
		return Bridging_Server2Client_unsellingOk(buffer, p)
	case 212:
		return Bridging_Server2Client_redemptionSpreeOk(buffer, p)
	case 213:
		return Bridging_Server2Client_insertState(buffer, p)
	case 214:
		return Bridging_Server2Client_updattState(buffer, p)
	case 215:
		return Bridging_Server2Client_removeState(buffer, p)
	case 216:
		return Bridging_Server2Client_requestFixItemOk(buffer, p)
	case 217:
		return Bridging_Server2Client_makeDebirsItemOK(buffer, p)
	case 218:
		return Bridging_Server2Client_updateMagicItem(buffer, p)
	case 219:
		return Bridging_Server2Client_changeMagicJobOk(buffer, p)
	case 220:
		return Bridging_Server2Client_magicItemTupoOk(buffer, p)
	case 221:
		return Bridging_Server2Client_zhuanpanOK(buffer, p)
	case 222:
		return Bridging_Server2Client_updateZhuanpanNotice(buffer, p)
	case 223:
		return Bridging_Server2Client_sycnZhuanpanData(buffer, p)
	case 224:
		return Bridging_Server2Client_copynonum(buffer, p)
	case 225:
		return Bridging_Server2Client_sceneFilterOk(buffer, p)
	case 226:
		return Bridging_Server2Client_wishingOK(buffer, p)
	case 227:
		return Bridging_Server2Client_shareWishOK(buffer, p)
	case 228:
		return Bridging_Server2Client_leaderCloseDialogOk(buffer, p)
	case 229:
		return Bridging_Server2Client_startOnlineTime(buffer, p)
	case 230:
		return Bridging_Server2Client_requestOnlineTimeRewardOK(buffer, p)
	case 231:
		return Bridging_Server2Client_sycnVipflag(buffer, p)
	case 232:
		return Bridging_Server2Client_buyFundOK(buffer, p)
	case 233:
		return Bridging_Server2Client_requestFundRewardOK(buffer, p)
	case 234:
		return Bridging_Server2Client_firstRechargeOK(buffer, p)
	case 235:
		return Bridging_Server2Client_firstRechargeGiftOK(buffer, p)
	case 236:
		return Bridging_Server2Client_agencyActivity(buffer, p)
	case 237:
		return Bridging_Server2Client_updateFestival(buffer, p)
	case 238:
		return Bridging_Server2Client_updateSelfRecharge(buffer, p)
	case 239:
		return Bridging_Server2Client_updateSysRecharge(buffer, p)
	case 240:
		return Bridging_Server2Client_updateSelfDiscountStore(buffer, p)
	case 241:
		return Bridging_Server2Client_updateSysDiscountStore(buffer, p)
	case 242:
		return Bridging_Server2Client_updateSelfOnceRecharge(buffer, p)
	case 243:
		return Bridging_Server2Client_updateSysOnceRecharge(buffer, p)
	case 244:
		return Bridging_Server2Client_openCardOK(buffer, p)
	case 245:
		return Bridging_Server2Client_resetCardOK(buffer, p)
	case 246:
		return Bridging_Server2Client_sycnHotRole(buffer, p)
	case 247:
		return Bridging_Server2Client_hotRoleBuyOk(buffer, p)
	case 248:
		return Bridging_Server2Client_updateSevenday(buffer, p)
	case 249:
		return Bridging_Server2Client_updateEmployeeActivity(buffer, p)
	case 250:
		return Bridging_Server2Client_updateMinGiftActivity(buffer, p)
	case 251:
		return Bridging_Server2Client_updateIntegralShop(buffer, p)
	case 252:
		return Bridging_Server2Client_updateShowBaby(buffer, p)
	case 253:
		return Bridging_Server2Client_updateMySelfRecharge(buffer, p)
	case 254:
		return Bridging_Server2Client_verificationSMSOk(buffer, p)
	case 255:
		return Bridging_Server2Client_requestLevelGiftOK(buffer, p)
	case 256:
		return Bridging_Server2Client_sycnConvertExp(buffer, p)
	case 257:
		return Bridging_Server2Client_wearFuwenOk(buffer, p)
	case 258:
		return Bridging_Server2Client_takeoffFuwenOk(buffer, p)
	case 259:
		return Bridging_Server2Client_compFuwenOk(buffer, p)
	case 260:
		return Bridging_Server2Client_requestEmployeeQuestOk(buffer, p)
	case 261:
		return Bridging_Server2Client_acceptEmployeeQuestOk(buffer, p)
	case 262:
		return Bridging_Server2Client_submitEmployeeQuestOk(buffer, p)
	case 263:
		return Bridging_Server2Client_sycnCrystal(buffer, p)
	case 264:
		return Bridging_Server2Client_crystalUpLeveResult(buffer, p)
	case 265:
		return Bridging_Server2Client_resetCrystalPropOK(buffer, p)
	case 266:
		return Bridging_Server2Client_sycnCourseGift(buffer, p)
	case 267:
		return Bridging_Server2Client_orderOk(buffer, p)
	case 268:
		return Bridging_Server2Client_updateRandSubmitQuestCount(buffer, p)
	default:
		return errors.New(prpc.NoneDispatchMatchError)
	}
}
