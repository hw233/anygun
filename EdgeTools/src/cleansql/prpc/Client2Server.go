package prpc

import (
	"bytes"
	"errors"
	"suzuki/prpc"
)

type Client2Server_openvip struct {
	vl int //0
}
type Client2Server_sessionlogin struct {
	info COM_LoginInfo //0
}
type Client2Server_login struct {
	info COM_LoginInfo //0
}
type Client2Server_createPlayer struct {
	playername  string //0
	playerTmpId uint8  //1
}
type Client2Server_deletePlayer struct {
	playername string //0
}
type Client2Server_enterGame struct {
	playerInstId uint32 //0
}
type Client2Server_requestStorage struct {
	tp int //0
}
type Client2Server_move struct {
	x float32 //0
	z float32 //1
}
type Client2Server_moveToNpc struct {
	npcid int32 //0
}
type Client2Server_moveToNpc2 struct {
	type_ int //0
}
type Client2Server_moveToZone struct {
	sceneId int32 //0
	zoneId  int32 //1
}
type Client2Server_transforScene struct {
	sceneId int32 //0
}
type Client2Server_querySimplePlayerInst struct {
	instId uint32 //0
}
type Client2Server_changProp struct {
	guid  uint32        //0
	props []COM_Addprop //1
}
type Client2Server_learnSkill struct {
	skid uint32 //0
}
type Client2Server_babyLearnSkill struct {
	instId  uint32 //0
	oldSkId uint32 //1
	newSkId uint32 //2
	newSkLv uint32 //3
}
type Client2Server_forgetSkill struct {
	skid uint32 //0
}
type Client2Server_syncOrder struct {
	order COM_Order //0
}
type Client2Server_sendChat struct {
	content    COM_Chat //0
	targetName string   //1
}
type Client2Server_requestAudio struct {
	audioId int32 //0
}
type Client2Server_publishItemInst struct {
	type_      int    //0
	itemInstId uint32 //1
	chatType   int    //2
	playerName string //3
}
type Client2Server_queryItemInst struct {
	showId int32 //0
}
type Client2Server_publishbabyInst struct {
	type_      int    //0
	babyInstId uint32 //1
	playerName string //2
}
type Client2Server_querybabyInst struct {
	showId int32 //0
}
type Client2Server_useItem struct {
	slot   uint32 //0
	target uint32 //1
	stack  uint32 //2
}
type Client2Server_wearEquipment struct {
	target     uint32 //0
	itemInstId uint32 //1
}
type Client2Server_delEquipment struct {
	target     uint32 //0
	itemInstId uint32 //1
}
type Client2Server_setPlayerFront struct {
	isFront bool //0
}
type Client2Server_setBattlebaby struct {
	babyID   uint32 //0
	isBattle bool   //1
}
type Client2Server_changeBabyName struct {
	babyID uint32 //0
	name   string //1
}
type Client2Server_intensifyBaby struct {
	babyid uint32 //0
}
type Client2Server_setBattleEmp struct {
	empID    uint32 //0
	group    int    //1
	isBattle bool   //2
}
type Client2Server_changeEmpBattleGroup struct {
	group int //0
}
type Client2Server_requestEvolve struct {
	empInstId uint32 //0
}
type Client2Server_requestUpStar struct {
	empInstId uint32 //0
}
type Client2Server_requestDelEmp struct {
	empInstId uint32 //0
}
type Client2Server_delEmployee struct {
	emps []uint32 //0
}
type Client2Server_delEmployeeSoul struct {
	instid  uint32 //0
	soulNum uint32 //1
}
type Client2Server_sellBagItem struct {
	instId uint32 //0
	stack  uint32 //1
}
type Client2Server_depositItemToStorage struct {
	instid uint32 //0
}
type Client2Server_depositBabyToStorage struct {
	instid uint32 //0
}
type Client2Server_storageItemToBag struct {
	instid uint32 //0
}
type Client2Server_storageBabyToPlayer struct {
	instid uint32 //0
}
type Client2Server_sortStorage struct {
	tp int //0
}
type Client2Server_delStorageBaby struct {
	instid uint32 //0
}
type Client2Server_createTeam struct {
	cti COM_CreateTeamInfo //0
}
type Client2Server_changeTeam struct {
	info COM_CreateTeamInfo //0
}
type Client2Server_kickTeamMember struct {
	uuid uint32 //0
}
type Client2Server_changeTeamLeader struct {
	uuid uint32 //0
}
type Client2Server_joinTeam struct {
	teamId uint32 //0
	pwd    string //1
}
type Client2Server_changeTeamPassword struct {
	pwd string //0
}
type Client2Server_inviteTeamMember struct {
	name string //0
}
type Client2Server_isjoinTeam struct {
	teamId uint32 //0
	isFlag bool   //1
}
type Client2Server_teamCallMember struct {
	playerId int32 //0
}
type Client2Server_requestJoinTeam struct {
	targetName string //0
}
type Client2Server_ratifyJoinTeam struct {
	sendName string //0
}
type Client2Server_drawLotteryBox struct {
	type_  int  //0
	isFree bool //1
}
type Client2Server_acceptQuest struct {
	questId int32 //0
}
type Client2Server_submitQuest struct {
	npcId   int32 //0
	questId int32 //1
}
type Client2Server_giveupQuest struct {
	questId int32 //0
}
type Client2Server_requestContactInfoById struct {
	instId uint32 //0
}
type Client2Server_requestContactInfoByName struct {
	instName string //0
}
type Client2Server_addFriend struct {
	instId uint32 //0
}
type Client2Server_delFriend struct {
	instId uint32 //0
}
type Client2Server_addBlacklist struct {
	instId uint32 //0
}
type Client2Server_delBlacklist struct {
	instId uint32 //0
}
type Client2Server_mining struct {
	gatherId int32 //0
	times    int32 //1
}
type Client2Server_compoundItem struct {
	itemId int32 //0
	gemId  int32 //1
}
type Client2Server_bagItemSplit struct {
	instId   int32 //0
	splitNum int32 //1
}
type Client2Server_requestChallenge struct {
	name string //0
}
type Client2Server_requestCheckMsg struct {
	name string //0
}
type Client2Server_queryOnlinePlayerbyName struct {
	name string //0
}
type Client2Server_queryPlayerbyName struct {
	name string //0
}
type Client2Server_queryBaby struct {
	instId uint32 //0
}
type Client2Server_queryEmployee struct {
	instId uint32 //0
}
type Client2Server_guideFinish struct {
	guideIdx uint64 //0
}
type Client2Server_enterBattle struct {
	battleId int32 //0
}
type Client2Server_shopBuyItem struct {
	id  int32 //0
	num int32 //1
}
type Client2Server_requestLevelGift struct {
	level int32 //0
}
type Client2Server_setCurrentTitle struct {
	title int32 //0
}
type Client2Server_requestAchaward struct {
	achId int32 //0
}
type Client2Server_sign struct {
	index int32 //0
}
type Client2Server_requestActivityReward struct {
	index int32 //0
}
type Client2Server_enterHundredScene struct {
	level int32 //0
}
type Client2Server_delBaby struct {
	instId int32 //0
}
type Client2Server_resetBaby struct {
	instId int32 //0
}
type Client2Server_resetBabyProp struct {
	instId int32 //0
}
type Client2Server_remouldBaby struct {
	instid int32 //0
}
type Client2Server_empSkillLevelUp struct {
	empId   uint32 //0
	skillId int32  //1
}
type Client2Server_setOpenDoubleTimeFlag struct {
	isFlag bool //0
}
type Client2Server_talkedNpc struct {
	npcId int32 //0
}
type Client2Server_jjcBattleGo struct {
	id uint32 //0
}
type Client2Server_sendMail struct {
	playername string //0
	title      string //1
	content    string //2
}
type Client2Server_readMail struct {
	mailId int32 //0
}
type Client2Server_delMail struct {
	mailId int32 //0
}
type Client2Server_getMailItem struct {
	mailId int32 //0
}
type Client2Server_createGuild struct {
	guildName string //0
}
type Client2Server_delGuild struct {
	guildId uint32 //0
}
type Client2Server_requestJoinGuild struct {
	guid uint32 //0
}
type Client2Server_kickOut struct {
	guid int32 //0
}
type Client2Server_acceptRequestGuild struct {
	playerId int32 //0
}
type Client2Server_refuseRequestGuild struct {
	playerId int32 //0
}
type Client2Server_changeMemberPosition struct {
	targetId int32 //0
	job      int   //1
}
type Client2Server_transferPremier struct {
	targetId int32 //0
}
type Client2Server_changeGuildNotice struct {
	notice string //0
}
type Client2Server_queryGuildList struct {
	page int16 //0
}
type Client2Server_inviteJoinGuild struct {
	playerName string //0
}
type Client2Server_respondInviteJoinGuild struct {
	sendName string //0
}
type Client2Server_buyGuildItem struct {
	tableId int32 //0
	times   int32 //1
}
type Client2Server_addGuildMoney struct {
	money int32 //0
}
type Client2Server_updateGuildBuiling struct {
	gbt int //0
}
type Client2Server_levelupGuildSkill struct {
	skId int32 //0
}
type Client2Server_presentGuildItem struct {
	num int32 //0
}
type Client2Server_progenitusAddExp struct {
	monsterId int32 //0
	isSuper   bool  //1
}
type Client2Server_setProgenitusPosition struct {
	mId int32 //0
	pos int32 //1
}
type Client2Server_fetchSelling struct {
	context COM_SearchContext //0
}
type Client2Server_fetchSelling2 struct {
	context COM_SearchContext //0
}
type Client2Server_selling struct {
	iteminstid int32 //0
	babyinstid int32 //1
	price      int32 //2
}
type Client2Server_unselling struct {
	sellid int32 //0
}
type Client2Server_buy struct {
	sellid int32 //0
}
type Client2Server_fixItem struct {
	instId int32 //0
	type_  int   //1
}
type Client2Server_fixAllItem struct {
	items []uint32 //0
	type_ int      //1
}
type Client2Server_makeDebirsItem struct {
	instId int32 //0
	num    int32 //1
}
type Client2Server_levelUpMagicItem struct {
	items []uint32 //0
}
type Client2Server_tupoMagicItem struct {
	level int32 //0
}
type Client2Server_changeMagicJob struct {
	job int //0
}
type Client2Server_requestPk struct {
	playerId uint32 //0
}
type Client2Server_uiBehavior struct {
	type_ int //0
}
type Client2Server_zhuanpanGo struct {
	counter uint32 //0
}
type Client2Server_redemptionSpree struct {
	code string //0
}
type Client2Server_sceneFilter struct {
	sfType []int //0
}
type Client2Server_sendExamAnswer struct {
	questionId uint32 //0
	answer     uint8  //1
}
type Client2Server_sendwishing struct {
	wish COM_Wishing //0
}
type Client2Server_requestOnlineReward struct {
	index uint32 //0
}
type Client2Server_requestFundReward struct {
	level uint32 //0
}
type Client2Server_openCard struct {
	index uint16 //0
}
type Client2Server_requestSevenReward struct {
	qid uint32 //0
}
type Client2Server_requestChargeTotalSingleReward struct {
	index uint32 //0
}
type Client2Server_requestChargeTotalReward struct {
	index uint32 //0
}
type Client2Server_requestChargeEverySingleReward struct {
	index uint32 //0
}
type Client2Server_requestChargeEveryReward struct {
	index uint32 //0
}
type Client2Server_requestLoginTotal struct {
	index uint32 //0
}
type Client2Server_buyDiscountStoreSingle struct {
	itemId    int32 //0
	itemStack int32 //1
}
type Client2Server_buyDiscountStore struct {
	itemId    int32 //0
	itemStack int32 //1
}
type Client2Server_requestEmployeeActivityReward struct {
	index uint32 //0
}
type Client2Server_requestmyselfrechargeleReward struct {
	index uint32 //0
}
type Client2Server_buyIntegralItem struct {
	id  uint32 //0
	num uint32 //1
}
type Client2Server_verificationSMS struct {
	phoneNumber string //0
	code        string //1
}
type Client2Server_lockItem struct {
	instId int32 //0
	isLock bool  //1
}
type Client2Server_lockBaby struct {
	instId int32 //0
	isLock bool  //1
}
type Client2Server_showBaby struct {
	instId int32 //0
}
type Client2Server_wearFuwen struct {
	itemInstId int32 //0
}
type Client2Server_takeoffFuwen struct {
	slotId int32 //0
}
type Client2Server_compFuwen struct {
	itemInstId int32 //0
}
type Client2Server_acceptEmployeeQuest struct {
	questId   int32   //0
	employees []int32 //1
}
type Client2Server_submitEmployeeQuest struct {
	questId int32 //0
}
type Client2Server_resetCrystalProp struct {
	locklist []int32 //0
}
type Client2ServerStub struct {
	Sender prpc.StubSender
}
type Client2ServerProxy interface {
	openvip(vl int) error                                                                // 0
	requestPhoto() error                                                                 // 1
	ping() error                                                                         // 2
	sessionlogin(info COM_LoginInfo) error                                               // 3
	login(info COM_LoginInfo) error                                                      // 4
	createPlayer(playername string, playerTmpId uint8) error                             // 5
	deletePlayer(playername string) error                                                // 6
	enterGame(playerInstId uint32) error                                                 // 7
	requestBag() error                                                                   // 8
	requestEmployees() error                                                             // 9
	requestStorage(tp int) error                                                         // 10
	requestAchievement() error                                                           // 11
	initminig() error                                                                    // 12
	requestCompound() error                                                              // 13
	move(x float32, z float32) error                                                     // 14
	moveToNpc(npcid int32) error                                                         // 15
	moveToNpc2(type_ int) error                                                          // 16
	moveToZone(sceneId int32, zoneId int32) error                                        // 17
	autoBattle() error                                                                   // 18
	stopAutoBattle() error                                                               // 19
	stopMove() error                                                                     // 20
	exitCopy() error                                                                     // 21
	transforScene(sceneId int32) error                                                   // 22
	sceneLoaded() error                                                                  // 23
	querySimplePlayerInst(instId uint32) error                                           // 24
	logout() error                                                                       // 25
	changProp(guid uint32, props []COM_Addprop) error                                    // 26
	learnSkill(skid uint32) error                                                        // 27
	babyLearnSkill(instId uint32, oldSkId uint32, newSkId uint32, newSkLv uint32) error  // 28
	forgetSkill(skid uint32) error                                                       // 29
	syncOrder(order COM_Order) error                                                     // 30
	syncOrderTimeout() error                                                             // 31
	sendChat(content COM_Chat, targetName string) error                                  // 32
	requestAudio(audioId int32) error                                                    // 33
	publishItemInst(type_ int, itemInstId uint32, chatType int, playerName string) error // 34
	queryItemInst(showId int32) error                                                    // 35
	publishbabyInst(type_ int, babyInstId uint32, playerName string) error               // 36
	querybabyInst(showId int32) error                                                    // 37
	useItem(slot uint32, target uint32, stack uint32) error                              // 38
	wearEquipment(target uint32, itemInstId uint32) error                                // 39
	delEquipment(target uint32, itemInstId uint32) error                                 // 40
	setPlayerFront(isFront bool) error                                                   // 41
	setBattlebaby(babyID uint32, isBattle bool) error                                    // 42
	changeBabyName(babyID uint32, name string) error                                     // 43
	intensifyBaby(babyid uint32) error                                                   // 44
	setBattleEmp(empID uint32, group int, isBattle bool) error                           // 45
	changeEmpBattleGroup(group int) error                                                // 46
	requestEvolve(empInstId uint32) error                                                // 47
	requestUpStar(empInstId uint32) error                                                // 48
	requestDelEmp(empInstId uint32) error                                                // 49
	delEmployee(emps []uint32) error                                                     // 50
	onekeyDelEmp() error                                                                 // 51
	delEmployeeSoul(instid uint32, soulNum uint32) error                                 // 52
	sortBagItem() error                                                                  // 53
	sellBagItem(instId uint32, stack uint32) error                                       // 54
	depositItemToStorage(instid uint32) error                                            // 55
	depositBabyToStorage(instid uint32) error                                            // 56
	storageItemToBag(instid uint32) error                                                // 57
	storageBabyToPlayer(instid uint32) error                                             // 58
	sortStorage(tp int) error                                                            // 59
	delStorageBaby(instid uint32) error                                                  // 60
	jointLobby() error                                                                   // 61
	exitLobby() error                                                                    // 62
	createTeam(cti COM_CreateTeamInfo) error                                             // 63
	changeTeam(info COM_CreateTeamInfo) error                                            // 64
	kickTeamMember(uuid uint32) error                                                    // 65
	changeTeamLeader(uuid uint32) error                                                  // 66
	joinTeam(teamId uint32, pwd string) error                                            // 67
	exitTeam() error                                                                     // 68
	changeTeamPassword(pwd string) error                                                 // 69
	joinTeamRoom() error                                                                 // 70
	inviteTeamMember(name string) error                                                  // 71
	isjoinTeam(teamId uint32, isFlag bool) error                                         // 72
	leaveTeam() error                                                                    // 73
	backTeam() error                                                                     // 74
	refuseBackTeam() error                                                               // 75
	teamCallMember(playerId int32) error                                                 // 76
	requestJoinTeam(targetName string) error                                             // 77
	ratifyJoinTeam(sendName string) error                                                // 78
	drawLotteryBox(type_ int, isFree bool) error                                         // 79
	acceptQuest(questId int32) error                                                     // 80
	submitQuest(npcId int32, questId int32) error                                        // 81
	giveupQuest(questId int32) error                                                     // 82
	requestContactInfoById(instId uint32) error                                          // 83
	requestContactInfoByName(instName string) error                                      // 84
	requestFriendList() error                                                            // 85
	addFriend(instId uint32) error                                                       // 86
	delFriend(instId uint32) error                                                       // 87
	addBlacklist(instId uint32) error                                                    // 88
	delBlacklist(instId uint32) error                                                    // 89
	requestReferrFriend() error                                                          // 90
	mining(gatherId int32, times int32) error                                            // 91
	compoundItem(itemId int32, gemId int32) error                                        // 92
	bagItemSplit(instId int32, splitNum int32) error                                     // 93
	requestChallenge(name string) error                                                  // 94
	requestRival() error                                                                 // 95
	requestMySelfJJCData() error                                                         // 96
	requestCheckMsg(name string) error                                                   // 97
	requestMyAllbattleMsg() error                                                        // 98
	requestJJCRank() error                                                               // 99
	requestLevelRank() error                                                             // 100
	requestBabyRank() error                                                              // 101
	requestEmpRank() error                                                               // 102
	requestPlayerFFRank() error                                                          // 103
	queryOnlinePlayerbyName(name string) error                                           // 104
	queryPlayerbyName(name string) error                                                 // 105
	queryBaby(instId uint32) error                                                       // 106
	queryEmployee(instId uint32) error                                                   // 107
	guideFinish(guideIdx uint64) error                                                   // 108
	enterBattle(battleId int32) error                                                    // 109
	shopBuyItem(id int32, num int32) error                                               // 110
	getFirstRechargeItem() error                                                         // 111
	requestLevelGift(level int32) error                                                  // 112
	setCurrentTitle(title int32) error                                                   // 113
	openBuyBox() error                                                                   // 114
	requestAchaward(achId int32) error                                                   // 115
	sign(index int32) error                                                              // 116
	requestSignupReward7() error                                                         // 117
	requestSignupReward14() error                                                        // 118
	requestSignupReward28() error                                                        // 119
	requestActivityReward(index int32) error                                             // 120
	resetHundredTier() error                                                             // 121
	enterHundredScene(level int32) error                                                 // 122
	delBaby(instId int32) error                                                          // 123
	resetBaby(instId int32) error                                                        // 124
	resetBabyProp(instId int32) error                                                    // 125
	remouldBaby(instid int32) error                                                      // 126
	empSkillLevelUp(empId uint32, skillId int32) error                                   // 127
	setOpenDoubleTimeFlag(isFlag bool) error                                             // 128
	talkedNpc(npcId int32) error                                                         // 129
	jjcBattleGo(id uint32) error                                                         // 130
	requestMyJJCTeamMsg() error                                                          // 131
	startMatching() error                                                                // 132
	stopMatching() error                                                                 // 133
	exitPvpJJC() error                                                                   // 134
	joinPvpLobby() error                                                                 // 135
	exitPvpLobby() error                                                                 // 136
	requestpvprank() error                                                               // 137
	joinWarriorchoose() error                                                            // 138
	warriorStart() error                                                                 // 139
	warriorStop() error                                                                  // 140
	sendMail(playername string, title string, content string) error                      // 141
	readMail(mailId int32) error                                                         // 142
	delMail(mailId int32) error                                                          // 143
	getMailItem(mailId int32) error                                                      // 144
	requestState() error                                                                 // 145
	createGuild(guildName string) error                                                  // 146
	delGuild(guildId uint32) error                                                       // 147
	requestJoinGuild(guid uint32) error                                                  // 148
	leaveGuild() error                                                                   // 149
	kickOut(guid int32) error                                                            // 150
	acceptRequestGuild(playerId int32) error                                             // 151
	refuseRequestGuild(playerId int32) error                                             // 152
	changeMemberPosition(targetId int32, job int) error                                  // 153
	transferPremier(targetId int32) error                                                // 154
	changeGuildNotice(notice string) error                                               // 155
	queryGuildList(page int16) error                                                     // 156
	inviteJoinGuild(playerName string) error                                             // 157
	respondInviteJoinGuild(sendName string) error                                        // 158
	buyGuildItem(tableId int32, times int32) error                                       // 159
	entryGuildBattle() error                                                             // 160
	transforGuildBattleScene() error                                                     // 161
	addGuildMoney(money int32) error                                                     // 162
	updateGuildBuiling(gbt int) error                                                    // 163
	refreshGuildShop() error                                                             // 164
	levelupGuildSkill(skId int32) error                                                  // 165
	presentGuildItem(num int32) error                                                    // 166
	progenitusAddExp(monsterId int32, isSuper bool) error                                // 167
	setProgenitusPosition(mId int32, pos int32) error                                    // 168
	guildsign() error                                                                    // 169
	fetchSelling(context COM_SearchContext) error                                        // 170
	fetchSelling2(context COM_SearchContext) error                                       // 171
	selling(iteminstid int32, babyinstid int32, price int32) error                       // 172
	unselling(sellid int32) error                                                        // 173
	buy(sellid int32) error                                                              // 174
	fixItem(instId int32, type_ int) error                                               // 175
	fixAllItem(items []uint32, type_ int) error                                          // 176
	makeDebirsItem(instId int32, num int32) error                                        // 177
	levelUpMagicItem(items []uint32) error                                               // 178
	tupoMagicItem(level int32) error                                                     // 179
	changeMagicJob(job int) error                                                        // 180
	requestPk(playerId uint32) error                                                     // 181
	uiBehavior(type_ int) error                                                          // 182
	openZhuanpan() error                                                                 // 183
	zhuanpanGo(counter uint32) error                                                     // 184
	redemptionSpree(code string) error                                                   // 185
	sceneFilter(sfType []int) error                                                      // 186
	sendExamAnswer(questionId uint32, answer uint8) error                                // 187
	sendwishing(wish COM_Wishing) error                                                  // 188
	requestWish() error                                                                  // 189
	leaderCloseDialog() error                                                            // 190
	requestOnlineReward(index uint32) error                                              // 191
	requestFundReward(level uint32) error                                                // 192
	openCard(index uint16) error                                                         // 193
	resetCard() error                                                                    // 194
	hotRoleBuy() error                                                                   // 195
	requestSevenReward(qid uint32) error                                                 // 196
	vipreward() error                                                                    // 197
	requestChargeTotalSingleReward(index uint32) error                                   // 198
	requestChargeTotalReward(index uint32) error                                         // 199
	requestChargeEverySingleReward(index uint32) error                                   // 200
	requestChargeEveryReward(index uint32) error                                         // 201
	requestLoginTotal(index uint32) error                                                // 202
	buyDiscountStoreSingle(itemId int32, itemStack int32) error                          // 203
	buyDiscountStore(itemId int32, itemStack int32) error                                // 204
	requestEmployeeActivityReward(index uint32) error                                    // 205
	requestmyselfrechargeleReward(index uint32) error                                    // 206
	requestEverydayIntegral() error                                                      // 207
	buyIntegralItem(id uint32, num uint32) error                                         // 208
	familyLoseLeader() error                                                             // 209
	verificationSMS(phoneNumber string, code string) error                               // 210
	lockItem(instId int32, isLock bool) error                                            // 211
	lockBaby(instId int32, isLock bool) error                                            // 212
	showBaby(instId int32) error                                                         // 213
	wearFuwen(itemInstId int32) error                                                    // 214
	takeoffFuwen(slotId int32) error                                                     // 215
	compFuwen(itemInstId int32) error                                                    // 216
	requestEmployeeQuest() error                                                         // 217
	acceptEmployeeQuest(questId int32, employees []int32) error                          // 218
	submitEmployeeQuest(questId int32) error                                             // 219
	crystalUpLevel() error                                                               // 220
	resetCrystalProp(locklist []int32) error                                             // 221
	magicItemOneKeyLevel() error                                                         // 222
}

func (this *Client2Server_openvip) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.vl != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize vl
	{
		if this.vl != 0 {
			err := prpc.Write(buffer, this.vl)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Client2Server_openvip) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize vl
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.vl)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Client2Server_sessionlogin) Serialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_sessionlogin) Deserialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_login) Serialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_login) Deserialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_createPlayer) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.playername) != 0)
	mask.WriteBit(this.playerTmpId != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize playername
	if len(this.playername) != 0 {
		err := prpc.Write(buffer, this.playername)
		if err != nil {
			return err
		}
	}
	// serialize playerTmpId
	{
		if this.playerTmpId != 0 {
			err := prpc.Write(buffer, this.playerTmpId)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Client2Server_createPlayer) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize playername
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.playername)
		if err != nil {
			return err
		}
	}
	// deserialize playerTmpId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.playerTmpId)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Client2Server_deletePlayer) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.playername) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize playername
	if len(this.playername) != 0 {
		err := prpc.Write(buffer, this.playername)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Client2Server_deletePlayer) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize playername
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.playername)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Client2Server_enterGame) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.playerInstId != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize playerInstId
	{
		if this.playerInstId != 0 {
			err := prpc.Write(buffer, this.playerInstId)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Client2Server_enterGame) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize playerInstId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.playerInstId)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Client2Server_requestStorage) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.tp != 0)
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
	return nil
}
func (this *Client2Server_requestStorage) Deserialize(buffer *bytes.Buffer) error {
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
	return nil
}
func (this *Client2Server_move) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.x != 0)
	mask.WriteBit(this.z != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize x
	{
		if this.x != 0 {
			err := prpc.Write(buffer, this.x)
			if err != nil {
				return err
			}
		}
	}
	// serialize z
	{
		if this.z != 0 {
			err := prpc.Write(buffer, this.z)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Client2Server_move) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize x
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.x)
		if err != nil {
			return err
		}
	}
	// deserialize z
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.z)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Client2Server_moveToNpc) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.npcid != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize npcid
	{
		if this.npcid != 0 {
			err := prpc.Write(buffer, this.npcid)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Client2Server_moveToNpc) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize npcid
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.npcid)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Client2Server_moveToNpc2) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.type_ != 0)
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
	return nil
}
func (this *Client2Server_moveToNpc2) Deserialize(buffer *bytes.Buffer) error {
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
	return nil
}
func (this *Client2Server_moveToZone) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.sceneId != 0)
	mask.WriteBit(this.zoneId != 0)
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
	// serialize zoneId
	{
		if this.zoneId != 0 {
			err := prpc.Write(buffer, this.zoneId)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Client2Server_moveToZone) Deserialize(buffer *bytes.Buffer) error {
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
	// deserialize zoneId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.zoneId)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Client2Server_transforScene) Serialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_transforScene) Deserialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_querySimplePlayerInst) Serialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_querySimplePlayerInst) Deserialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_changProp) Serialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_changProp) Deserialize(buffer *bytes.Buffer) error {
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
		this.props = make([]COM_Addprop, size)
		for i, _ := range this.props {
			err := this.props[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Client2Server_learnSkill) Serialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_learnSkill) Deserialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_babyLearnSkill) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.instId != 0)
	mask.WriteBit(this.oldSkId != 0)
	mask.WriteBit(this.newSkId != 0)
	mask.WriteBit(this.newSkLv != 0)
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
	// serialize oldSkId
	{
		if this.oldSkId != 0 {
			err := prpc.Write(buffer, this.oldSkId)
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
	// serialize newSkLv
	{
		if this.newSkLv != 0 {
			err := prpc.Write(buffer, this.newSkLv)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Client2Server_babyLearnSkill) Deserialize(buffer *bytes.Buffer) error {
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
	// deserialize oldSkId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.oldSkId)
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
	// deserialize newSkLv
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.newSkLv)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Client2Server_forgetSkill) Serialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_forgetSkill) Deserialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_syncOrder) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //order
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize order
	{
		err := this.order.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Client2Server_syncOrder) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize order
	if mask.ReadBit() {
		err := this.order.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Client2Server_sendChat) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //content
	mask.WriteBit(len(this.targetName) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize content
	{
		err := this.content.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	// serialize targetName
	if len(this.targetName) != 0 {
		err := prpc.Write(buffer, this.targetName)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Client2Server_sendChat) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize content
	if mask.ReadBit() {
		err := this.content.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	// deserialize targetName
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.targetName)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Client2Server_requestAudio) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.audioId != 0)
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
	return nil
}
func (this *Client2Server_requestAudio) Deserialize(buffer *bytes.Buffer) error {
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
	return nil
}
func (this *Client2Server_publishItemInst) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.type_ != 0)
	mask.WriteBit(this.itemInstId != 0)
	mask.WriteBit(this.chatType != 0)
	mask.WriteBit(len(this.playerName) != 0)
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
	// serialize itemInstId
	{
		if this.itemInstId != 0 {
			err := prpc.Write(buffer, this.itemInstId)
			if err != nil {
				return err
			}
		}
	}
	// serialize chatType
	{
		if this.chatType != 0 {
			err := prpc.Write(buffer, this.chatType)
			if err != nil {
				return err
			}
		}
	}
	// serialize playerName
	if len(this.playerName) != 0 {
		err := prpc.Write(buffer, this.playerName)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Client2Server_publishItemInst) Deserialize(buffer *bytes.Buffer) error {
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
	// deserialize itemInstId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.itemInstId)
		if err != nil {
			return err
		}
	}
	// deserialize chatType
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.chatType)
		if err != nil {
			return err
		}
	}
	// deserialize playerName
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.playerName)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Client2Server_queryItemInst) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.showId != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize showId
	{
		if this.showId != 0 {
			err := prpc.Write(buffer, this.showId)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Client2Server_queryItemInst) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize showId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.showId)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Client2Server_publishbabyInst) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.type_ != 0)
	mask.WriteBit(this.babyInstId != 0)
	mask.WriteBit(len(this.playerName) != 0)
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
	// serialize babyInstId
	{
		if this.babyInstId != 0 {
			err := prpc.Write(buffer, this.babyInstId)
			if err != nil {
				return err
			}
		}
	}
	// serialize playerName
	if len(this.playerName) != 0 {
		err := prpc.Write(buffer, this.playerName)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Client2Server_publishbabyInst) Deserialize(buffer *bytes.Buffer) error {
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
	// deserialize babyInstId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.babyInstId)
		if err != nil {
			return err
		}
	}
	// deserialize playerName
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.playerName)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Client2Server_querybabyInst) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.showId != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize showId
	{
		if this.showId != 0 {
			err := prpc.Write(buffer, this.showId)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Client2Server_querybabyInst) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize showId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.showId)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Client2Server_useItem) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.slot != 0)
	mask.WriteBit(this.target != 0)
	mask.WriteBit(this.stack != 0)
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
	// serialize target
	{
		if this.target != 0 {
			err := prpc.Write(buffer, this.target)
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
func (this *Client2Server_useItem) Deserialize(buffer *bytes.Buffer) error {
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
	// deserialize target
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.target)
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
func (this *Client2Server_wearEquipment) Serialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_wearEquipment) Deserialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_delEquipment) Serialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_delEquipment) Deserialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_setPlayerFront) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.isFront)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize isFront
	{
	}
	return nil
}
func (this *Client2Server_setPlayerFront) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize isFront
	this.isFront = mask.ReadBit()
	return nil
}
func (this *Client2Server_setBattlebaby) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.babyID != 0)
	mask.WriteBit(this.isBattle)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize babyID
	{
		if this.babyID != 0 {
			err := prpc.Write(buffer, this.babyID)
			if err != nil {
				return err
			}
		}
	}
	// serialize isBattle
	{
	}
	return nil
}
func (this *Client2Server_setBattlebaby) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize babyID
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.babyID)
		if err != nil {
			return err
		}
	}
	// deserialize isBattle
	this.isBattle = mask.ReadBit()
	return nil
}
func (this *Client2Server_changeBabyName) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.babyID != 0)
	mask.WriteBit(len(this.name) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize babyID
	{
		if this.babyID != 0 {
			err := prpc.Write(buffer, this.babyID)
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
func (this *Client2Server_changeBabyName) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize babyID
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.babyID)
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
func (this *Client2Server_intensifyBaby) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.babyid != 0)
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
	return nil
}
func (this *Client2Server_intensifyBaby) Deserialize(buffer *bytes.Buffer) error {
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
	return nil
}
func (this *Client2Server_setBattleEmp) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.empID != 0)
	mask.WriteBit(this.group != 0)
	mask.WriteBit(this.isBattle)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize empID
	{
		if this.empID != 0 {
			err := prpc.Write(buffer, this.empID)
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
	// serialize isBattle
	{
	}
	return nil
}
func (this *Client2Server_setBattleEmp) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize empID
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.empID)
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
	// deserialize isBattle
	this.isBattle = mask.ReadBit()
	return nil
}
func (this *Client2Server_changeEmpBattleGroup) Serialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_changeEmpBattleGroup) Deserialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_requestEvolve) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.empInstId != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize empInstId
	{
		if this.empInstId != 0 {
			err := prpc.Write(buffer, this.empInstId)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Client2Server_requestEvolve) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize empInstId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.empInstId)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Client2Server_requestUpStar) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.empInstId != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize empInstId
	{
		if this.empInstId != 0 {
			err := prpc.Write(buffer, this.empInstId)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Client2Server_requestUpStar) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize empInstId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.empInstId)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Client2Server_requestDelEmp) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.empInstId != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize empInstId
	{
		if this.empInstId != 0 {
			err := prpc.Write(buffer, this.empInstId)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Client2Server_requestDelEmp) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize empInstId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.empInstId)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Client2Server_delEmployee) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.emps) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize emps
	if len(this.emps) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.emps)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.emps {
			err := prpc.Write(buffer, value)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Client2Server_delEmployee) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize emps
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.emps = make([]uint32, size)
		for i, _ := range this.emps {
			err := prpc.Read(buffer, &this.emps[i])
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Client2Server_delEmployeeSoul) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.instid != 0)
	mask.WriteBit(this.soulNum != 0)
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
func (this *Client2Server_delEmployeeSoul) Deserialize(buffer *bytes.Buffer) error {
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
	// deserialize soulNum
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.soulNum)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Client2Server_sellBagItem) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.instId != 0)
	mask.WriteBit(this.stack != 0)
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
func (this *Client2Server_sellBagItem) Deserialize(buffer *bytes.Buffer) error {
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
	// deserialize stack
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.stack)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Client2Server_depositItemToStorage) Serialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_depositItemToStorage) Deserialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_depositBabyToStorage) Serialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_depositBabyToStorage) Deserialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_storageItemToBag) Serialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_storageItemToBag) Deserialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_storageBabyToPlayer) Serialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_storageBabyToPlayer) Deserialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_sortStorage) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.tp != 0)
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
	return nil
}
func (this *Client2Server_sortStorage) Deserialize(buffer *bytes.Buffer) error {
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
	return nil
}
func (this *Client2Server_delStorageBaby) Serialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_delStorageBaby) Deserialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_createTeam) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //cti
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize cti
	{
		err := this.cti.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Client2Server_createTeam) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize cti
	if mask.ReadBit() {
		err := this.cti.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Client2Server_changeTeam) Serialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_changeTeam) Deserialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_kickTeamMember) Serialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_kickTeamMember) Deserialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_changeTeamLeader) Serialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_changeTeamLeader) Deserialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_joinTeam) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.teamId != 0)
	mask.WriteBit(len(this.pwd) != 0)
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
	// serialize pwd
	if len(this.pwd) != 0 {
		err := prpc.Write(buffer, this.pwd)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Client2Server_joinTeam) Deserialize(buffer *bytes.Buffer) error {
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
	// deserialize pwd
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.pwd)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Client2Server_changeTeamPassword) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.pwd) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize pwd
	if len(this.pwd) != 0 {
		err := prpc.Write(buffer, this.pwd)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Client2Server_changeTeamPassword) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize pwd
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.pwd)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Client2Server_inviteTeamMember) Serialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_inviteTeamMember) Deserialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_isjoinTeam) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.teamId != 0)
	mask.WriteBit(this.isFlag)
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
	// serialize isFlag
	{
	}
	return nil
}
func (this *Client2Server_isjoinTeam) Deserialize(buffer *bytes.Buffer) error {
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
	// deserialize isFlag
	this.isFlag = mask.ReadBit()
	return nil
}
func (this *Client2Server_teamCallMember) Serialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_teamCallMember) Deserialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_requestJoinTeam) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.targetName) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize targetName
	if len(this.targetName) != 0 {
		err := prpc.Write(buffer, this.targetName)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Client2Server_requestJoinTeam) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize targetName
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.targetName)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Client2Server_ratifyJoinTeam) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.sendName) != 0)
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
	return nil
}
func (this *Client2Server_ratifyJoinTeam) Deserialize(buffer *bytes.Buffer) error {
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
	return nil
}
func (this *Client2Server_drawLotteryBox) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.type_ != 0)
	mask.WriteBit(this.isFree)
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
	// serialize isFree
	{
	}
	return nil
}
func (this *Client2Server_drawLotteryBox) Deserialize(buffer *bytes.Buffer) error {
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
	// deserialize isFree
	this.isFree = mask.ReadBit()
	return nil
}
func (this *Client2Server_acceptQuest) Serialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_acceptQuest) Deserialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_submitQuest) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.npcId != 0)
	mask.WriteBit(this.questId != 0)
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
func (this *Client2Server_submitQuest) Deserialize(buffer *bytes.Buffer) error {
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
	// deserialize questId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.questId)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Client2Server_giveupQuest) Serialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_giveupQuest) Deserialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_requestContactInfoById) Serialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_requestContactInfoById) Deserialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_requestContactInfoByName) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.instName) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize instName
	if len(this.instName) != 0 {
		err := prpc.Write(buffer, this.instName)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Client2Server_requestContactInfoByName) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize instName
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.instName)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Client2Server_addFriend) Serialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_addFriend) Deserialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_delFriend) Serialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_delFriend) Deserialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_addBlacklist) Serialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_addBlacklist) Deserialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_delBlacklist) Serialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_delBlacklist) Deserialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_mining) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.gatherId != 0)
	mask.WriteBit(this.times != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize gatherId
	{
		if this.gatherId != 0 {
			err := prpc.Write(buffer, this.gatherId)
			if err != nil {
				return err
			}
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
func (this *Client2Server_mining) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize gatherId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.gatherId)
		if err != nil {
			return err
		}
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
func (this *Client2Server_compoundItem) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.itemId != 0)
	mask.WriteBit(this.gemId != 0)
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
	// serialize gemId
	{
		if this.gemId != 0 {
			err := prpc.Write(buffer, this.gemId)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Client2Server_compoundItem) Deserialize(buffer *bytes.Buffer) error {
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
	// deserialize gemId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.gemId)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Client2Server_bagItemSplit) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.instId != 0)
	mask.WriteBit(this.splitNum != 0)
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
	// serialize splitNum
	{
		if this.splitNum != 0 {
			err := prpc.Write(buffer, this.splitNum)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Client2Server_bagItemSplit) Deserialize(buffer *bytes.Buffer) error {
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
	// deserialize splitNum
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.splitNum)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Client2Server_requestChallenge) Serialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_requestChallenge) Deserialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_requestCheckMsg) Serialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_requestCheckMsg) Deserialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_queryOnlinePlayerbyName) Serialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_queryOnlinePlayerbyName) Deserialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_queryPlayerbyName) Serialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_queryPlayerbyName) Deserialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_queryBaby) Serialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_queryBaby) Deserialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_queryEmployee) Serialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_queryEmployee) Deserialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_guideFinish) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.guideIdx != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize guideIdx
	{
		if this.guideIdx != 0 {
			err := prpc.Write(buffer, this.guideIdx)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Client2Server_guideFinish) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize guideIdx
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.guideIdx)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Client2Server_enterBattle) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.battleId != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize battleId
	{
		if this.battleId != 0 {
			err := prpc.Write(buffer, this.battleId)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Client2Server_enterBattle) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize battleId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.battleId)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Client2Server_shopBuyItem) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.id != 0)
	mask.WriteBit(this.num != 0)
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
func (this *Client2Server_shopBuyItem) Deserialize(buffer *bytes.Buffer) error {
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
	// deserialize num
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.num)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Client2Server_requestLevelGift) Serialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_requestLevelGift) Deserialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_setCurrentTitle) Serialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_setCurrentTitle) Deserialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_requestAchaward) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.achId != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize achId
	{
		if this.achId != 0 {
			err := prpc.Write(buffer, this.achId)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Client2Server_requestAchaward) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize achId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.achId)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Client2Server_sign) Serialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_sign) Deserialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_requestActivityReward) Serialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_requestActivityReward) Deserialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_enterHundredScene) Serialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_enterHundredScene) Deserialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_delBaby) Serialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_delBaby) Deserialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_resetBaby) Serialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_resetBaby) Deserialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_resetBabyProp) Serialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_resetBabyProp) Deserialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_remouldBaby) Serialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_remouldBaby) Deserialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_empSkillLevelUp) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.empId != 0)
	mask.WriteBit(this.skillId != 0)
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
	// serialize skillId
	{
		if this.skillId != 0 {
			err := prpc.Write(buffer, this.skillId)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Client2Server_empSkillLevelUp) Deserialize(buffer *bytes.Buffer) error {
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
	// deserialize skillId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.skillId)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Client2Server_setOpenDoubleTimeFlag) Serialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_setOpenDoubleTimeFlag) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize isFlag
	this.isFlag = mask.ReadBit()
	return nil
}
func (this *Client2Server_talkedNpc) Serialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_talkedNpc) Deserialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_jjcBattleGo) Serialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_jjcBattleGo) Deserialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_sendMail) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.playername) != 0)
	mask.WriteBit(len(this.title) != 0)
	mask.WriteBit(len(this.content) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize playername
	if len(this.playername) != 0 {
		err := prpc.Write(buffer, this.playername)
		if err != nil {
			return err
		}
	}
	// serialize title
	if len(this.title) != 0 {
		err := prpc.Write(buffer, this.title)
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
	return nil
}
func (this *Client2Server_sendMail) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize playername
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.playername)
		if err != nil {
			return err
		}
	}
	// deserialize title
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.title)
		if err != nil {
			return err
		}
	}
	// deserialize content
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.content)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Client2Server_readMail) Serialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_readMail) Deserialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_delMail) Serialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_delMail) Deserialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_getMailItem) Serialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_getMailItem) Deserialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_createGuild) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.guildName) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
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
func (this *Client2Server_createGuild) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
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
func (this *Client2Server_delGuild) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.guildId != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize guildId
	{
		if this.guildId != 0 {
			err := prpc.Write(buffer, this.guildId)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Client2Server_delGuild) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize guildId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.guildId)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Client2Server_requestJoinGuild) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.guid != 0)
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
	return nil
}
func (this *Client2Server_requestJoinGuild) Deserialize(buffer *bytes.Buffer) error {
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
	return nil
}
func (this *Client2Server_kickOut) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.guid != 0)
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
	return nil
}
func (this *Client2Server_kickOut) Deserialize(buffer *bytes.Buffer) error {
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
	return nil
}
func (this *Client2Server_acceptRequestGuild) Serialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_acceptRequestGuild) Deserialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_refuseRequestGuild) Serialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_refuseRequestGuild) Deserialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_changeMemberPosition) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.targetId != 0)
	mask.WriteBit(this.job != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize targetId
	{
		if this.targetId != 0 {
			err := prpc.Write(buffer, this.targetId)
			if err != nil {
				return err
			}
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
func (this *Client2Server_changeMemberPosition) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize targetId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.targetId)
		if err != nil {
			return err
		}
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
func (this *Client2Server_transferPremier) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.targetId != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize targetId
	{
		if this.targetId != 0 {
			err := prpc.Write(buffer, this.targetId)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Client2Server_transferPremier) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize targetId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.targetId)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Client2Server_changeGuildNotice) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.notice) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize notice
	if len(this.notice) != 0 {
		err := prpc.Write(buffer, this.notice)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Client2Server_changeGuildNotice) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize notice
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.notice)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Client2Server_queryGuildList) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.page != 0)
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
	return nil
}
func (this *Client2Server_queryGuildList) Deserialize(buffer *bytes.Buffer) error {
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
	return nil
}
func (this *Client2Server_inviteJoinGuild) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.playerName) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize playerName
	if len(this.playerName) != 0 {
		err := prpc.Write(buffer, this.playerName)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Client2Server_inviteJoinGuild) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize playerName
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.playerName)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Client2Server_respondInviteJoinGuild) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.sendName) != 0)
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
	return nil
}
func (this *Client2Server_respondInviteJoinGuild) Deserialize(buffer *bytes.Buffer) error {
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
	return nil
}
func (this *Client2Server_buyGuildItem) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.tableId != 0)
	mask.WriteBit(this.times != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize tableId
	{
		if this.tableId != 0 {
			err := prpc.Write(buffer, this.tableId)
			if err != nil {
				return err
			}
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
func (this *Client2Server_buyGuildItem) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize tableId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.tableId)
		if err != nil {
			return err
		}
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
func (this *Client2Server_addGuildMoney) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.money != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize money
	{
		if this.money != 0 {
			err := prpc.Write(buffer, this.money)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Client2Server_addGuildMoney) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize money
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.money)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Client2Server_updateGuildBuiling) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.gbt != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize gbt
	{
		if this.gbt != 0 {
			err := prpc.Write(buffer, this.gbt)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Client2Server_updateGuildBuiling) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize gbt
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.gbt)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Client2Server_levelupGuildSkill) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.skId != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize skId
	{
		if this.skId != 0 {
			err := prpc.Write(buffer, this.skId)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Client2Server_levelupGuildSkill) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize skId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.skId)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Client2Server_presentGuildItem) Serialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_presentGuildItem) Deserialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_progenitusAddExp) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.monsterId != 0)
	mask.WriteBit(this.isSuper)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize monsterId
	{
		if this.monsterId != 0 {
			err := prpc.Write(buffer, this.monsterId)
			if err != nil {
				return err
			}
		}
	}
	// serialize isSuper
	{
	}
	return nil
}
func (this *Client2Server_progenitusAddExp) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize monsterId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.monsterId)
		if err != nil {
			return err
		}
	}
	// deserialize isSuper
	this.isSuper = mask.ReadBit()
	return nil
}
func (this *Client2Server_setProgenitusPosition) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.mId != 0)
	mask.WriteBit(this.pos != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize mId
	{
		if this.mId != 0 {
			err := prpc.Write(buffer, this.mId)
			if err != nil {
				return err
			}
		}
	}
	// serialize pos
	{
		if this.pos != 0 {
			err := prpc.Write(buffer, this.pos)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Client2Server_setProgenitusPosition) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize mId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.mId)
		if err != nil {
			return err
		}
	}
	// deserialize pos
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.pos)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Client2Server_fetchSelling) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //context
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize context
	{
		err := this.context.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Client2Server_fetchSelling) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize context
	if mask.ReadBit() {
		err := this.context.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Client2Server_fetchSelling2) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //context
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize context
	{
		err := this.context.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Client2Server_fetchSelling2) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize context
	if mask.ReadBit() {
		err := this.context.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Client2Server_selling) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.iteminstid != 0)
	mask.WriteBit(this.babyinstid != 0)
	mask.WriteBit(this.price != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize iteminstid
	{
		if this.iteminstid != 0 {
			err := prpc.Write(buffer, this.iteminstid)
			if err != nil {
				return err
			}
		}
	}
	// serialize babyinstid
	{
		if this.babyinstid != 0 {
			err := prpc.Write(buffer, this.babyinstid)
			if err != nil {
				return err
			}
		}
	}
	// serialize price
	{
		if this.price != 0 {
			err := prpc.Write(buffer, this.price)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Client2Server_selling) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize iteminstid
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.iteminstid)
		if err != nil {
			return err
		}
	}
	// deserialize babyinstid
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.babyinstid)
		if err != nil {
			return err
		}
	}
	// deserialize price
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.price)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Client2Server_unselling) Serialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_unselling) Deserialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_buy) Serialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_buy) Deserialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_fixItem) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.instId != 0)
	mask.WriteBit(this.type_ != 0)
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
func (this *Client2Server_fixItem) Deserialize(buffer *bytes.Buffer) error {
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
	// deserialize type_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.type_)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Client2Server_fixAllItem) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.items) != 0)
	mask.WriteBit(this.type_ != 0)
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
			err := prpc.Write(buffer, value)
			if err != nil {
				return err
			}
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
func (this *Client2Server_fixAllItem) Deserialize(buffer *bytes.Buffer) error {
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
		this.items = make([]uint32, size)
		for i, _ := range this.items {
			err := prpc.Read(buffer, &this.items[i])
			if err != nil {
				return err
			}
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
func (this *Client2Server_makeDebirsItem) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.instId != 0)
	mask.WriteBit(this.num != 0)
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
func (this *Client2Server_makeDebirsItem) Deserialize(buffer *bytes.Buffer) error {
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
	// deserialize num
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.num)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Client2Server_levelUpMagicItem) Serialize(buffer *bytes.Buffer) error {
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
			err := prpc.Write(buffer, value)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Client2Server_levelUpMagicItem) Deserialize(buffer *bytes.Buffer) error {
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
		this.items = make([]uint32, size)
		for i, _ := range this.items {
			err := prpc.Read(buffer, &this.items[i])
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Client2Server_tupoMagicItem) Serialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_tupoMagicItem) Deserialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_changeMagicJob) Serialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_changeMagicJob) Deserialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_requestPk) Serialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_requestPk) Deserialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_uiBehavior) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.type_ != 0)
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
	return nil
}
func (this *Client2Server_uiBehavior) Deserialize(buffer *bytes.Buffer) error {
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
	return nil
}
func (this *Client2Server_zhuanpanGo) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.counter != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
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
	return nil
}
func (this *Client2Server_zhuanpanGo) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize counter
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.counter)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Client2Server_redemptionSpree) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.code) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize code
	if len(this.code) != 0 {
		err := prpc.Write(buffer, this.code)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Client2Server_redemptionSpree) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize code
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.code)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Client2Server_sceneFilter) Serialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_sceneFilter) Deserialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_sendExamAnswer) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.questionId != 0)
	mask.WriteBit(this.answer != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize questionId
	{
		if this.questionId != 0 {
			err := prpc.Write(buffer, this.questionId)
			if err != nil {
				return err
			}
		}
	}
	// serialize answer
	{
		if this.answer != 0 {
			err := prpc.Write(buffer, this.answer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Client2Server_sendExamAnswer) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize questionId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.questionId)
		if err != nil {
			return err
		}
	}
	// deserialize answer
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.answer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Client2Server_sendwishing) Serialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_sendwishing) Deserialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_requestOnlineReward) Serialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_requestOnlineReward) Deserialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_requestFundReward) Serialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_requestFundReward) Deserialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_openCard) Serialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_openCard) Deserialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_requestSevenReward) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.qid != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize qid
	{
		if this.qid != 0 {
			err := prpc.Write(buffer, this.qid)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Client2Server_requestSevenReward) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize qid
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.qid)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Client2Server_requestChargeTotalSingleReward) Serialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_requestChargeTotalSingleReward) Deserialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_requestChargeTotalReward) Serialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_requestChargeTotalReward) Deserialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_requestChargeEverySingleReward) Serialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_requestChargeEverySingleReward) Deserialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_requestChargeEveryReward) Serialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_requestChargeEveryReward) Deserialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_requestLoginTotal) Serialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_requestLoginTotal) Deserialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_buyDiscountStoreSingle) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.itemId != 0)
	mask.WriteBit(this.itemStack != 0)
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
	// serialize itemStack
	{
		if this.itemStack != 0 {
			err := prpc.Write(buffer, this.itemStack)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Client2Server_buyDiscountStoreSingle) Deserialize(buffer *bytes.Buffer) error {
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
	// deserialize itemStack
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.itemStack)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Client2Server_buyDiscountStore) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.itemId != 0)
	mask.WriteBit(this.itemStack != 0)
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
	// serialize itemStack
	{
		if this.itemStack != 0 {
			err := prpc.Write(buffer, this.itemStack)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Client2Server_buyDiscountStore) Deserialize(buffer *bytes.Buffer) error {
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
	// deserialize itemStack
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.itemStack)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Client2Server_requestEmployeeActivityReward) Serialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_requestEmployeeActivityReward) Deserialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_requestmyselfrechargeleReward) Serialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_requestmyselfrechargeleReward) Deserialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_buyIntegralItem) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.id != 0)
	mask.WriteBit(this.num != 0)
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
func (this *Client2Server_buyIntegralItem) Deserialize(buffer *bytes.Buffer) error {
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
	// deserialize num
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.num)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Client2Server_verificationSMS) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.phoneNumber) != 0)
	mask.WriteBit(len(this.code) != 0)
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
	// serialize code
	if len(this.code) != 0 {
		err := prpc.Write(buffer, this.code)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Client2Server_verificationSMS) Deserialize(buffer *bytes.Buffer) error {
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
	// deserialize code
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.code)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Client2Server_lockItem) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.instId != 0)
	mask.WriteBit(this.isLock)
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
	// serialize isLock
	{
	}
	return nil
}
func (this *Client2Server_lockItem) Deserialize(buffer *bytes.Buffer) error {
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
	// deserialize isLock
	this.isLock = mask.ReadBit()
	return nil
}
func (this *Client2Server_lockBaby) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.instId != 0)
	mask.WriteBit(this.isLock)
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
	// serialize isLock
	{
	}
	return nil
}
func (this *Client2Server_lockBaby) Deserialize(buffer *bytes.Buffer) error {
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
	// deserialize isLock
	this.isLock = mask.ReadBit()
	return nil
}
func (this *Client2Server_showBaby) Serialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_showBaby) Deserialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_wearFuwen) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.itemInstId != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
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
func (this *Client2Server_wearFuwen) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
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
func (this *Client2Server_takeoffFuwen) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.slotId != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize slotId
	{
		if this.slotId != 0 {
			err := prpc.Write(buffer, this.slotId)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Client2Server_takeoffFuwen) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize slotId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.slotId)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Client2Server_compFuwen) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.itemInstId != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
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
func (this *Client2Server_compFuwen) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
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
func (this *Client2Server_acceptEmployeeQuest) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.questId != 0)
	mask.WriteBit(len(this.employees) != 0)
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
	// serialize employees
	if len(this.employees) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.employees)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.employees {
			err := prpc.Write(buffer, value)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Client2Server_acceptEmployeeQuest) Deserialize(buffer *bytes.Buffer) error {
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
	// deserialize employees
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.employees = make([]int32, size)
		for i, _ := range this.employees {
			err := prpc.Read(buffer, &this.employees[i])
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Client2Server_submitEmployeeQuest) Serialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_submitEmployeeQuest) Deserialize(buffer *bytes.Buffer) error {
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
func (this *Client2Server_resetCrystalProp) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.locklist) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize locklist
	if len(this.locklist) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.locklist)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.locklist {
			err := prpc.Write(buffer, value)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Client2Server_resetCrystalProp) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize locklist
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.locklist = make([]int32, size)
		for i, _ := range this.locklist {
			err := prpc.Read(buffer, &this.locklist[i])
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *Client2ServerStub) openvip(vl int) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(0))
	if err != nil {
		return err
	}
	_0 := Client2Server_openvip{}
	_0.vl = vl
	err = _0.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) requestPhoto() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(1))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) ping() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(2))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) sessionlogin(info COM_LoginInfo) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(3))
	if err != nil {
		return err
	}
	_3 := Client2Server_sessionlogin{}
	_3.info = info
	err = _3.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) login(info COM_LoginInfo) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(4))
	if err != nil {
		return err
	}
	_4 := Client2Server_login{}
	_4.info = info
	err = _4.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) createPlayer(playername string, playerTmpId uint8) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(5))
	if err != nil {
		return err
	}
	_5 := Client2Server_createPlayer{}
	_5.playername = playername
	_5.playerTmpId = playerTmpId
	err = _5.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) deletePlayer(playername string) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(6))
	if err != nil {
		return err
	}
	_6 := Client2Server_deletePlayer{}
	_6.playername = playername
	err = _6.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) enterGame(playerInstId uint32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(7))
	if err != nil {
		return err
	}
	_7 := Client2Server_enterGame{}
	_7.playerInstId = playerInstId
	err = _7.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) requestBag() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(8))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) requestEmployees() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(9))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) requestStorage(tp int) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(10))
	if err != nil {
		return err
	}
	_10 := Client2Server_requestStorage{}
	_10.tp = tp
	err = _10.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) requestAchievement() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(11))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) initminig() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(12))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) requestCompound() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(13))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) move(x float32, z float32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(14))
	if err != nil {
		return err
	}
	_14 := Client2Server_move{}
	_14.x = x
	_14.z = z
	err = _14.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) moveToNpc(npcid int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(15))
	if err != nil {
		return err
	}
	_15 := Client2Server_moveToNpc{}
	_15.npcid = npcid
	err = _15.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) moveToNpc2(type_ int) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(16))
	if err != nil {
		return err
	}
	_16 := Client2Server_moveToNpc2{}
	_16.type_ = type_
	err = _16.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) moveToZone(sceneId int32, zoneId int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(17))
	if err != nil {
		return err
	}
	_17 := Client2Server_moveToZone{}
	_17.sceneId = sceneId
	_17.zoneId = zoneId
	err = _17.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) autoBattle() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(18))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) stopAutoBattle() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(19))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) stopMove() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(20))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) exitCopy() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(21))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) transforScene(sceneId int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(22))
	if err != nil {
		return err
	}
	_22 := Client2Server_transforScene{}
	_22.sceneId = sceneId
	err = _22.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) sceneLoaded() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(23))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) querySimplePlayerInst(instId uint32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(24))
	if err != nil {
		return err
	}
	_24 := Client2Server_querySimplePlayerInst{}
	_24.instId = instId
	err = _24.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) logout() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(25))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) changProp(guid uint32, props []COM_Addprop) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(26))
	if err != nil {
		return err
	}
	_26 := Client2Server_changProp{}
	_26.guid = guid
	_26.props = props
	err = _26.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) learnSkill(skid uint32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(27))
	if err != nil {
		return err
	}
	_27 := Client2Server_learnSkill{}
	_27.skid = skid
	err = _27.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) babyLearnSkill(instId uint32, oldSkId uint32, newSkId uint32, newSkLv uint32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(28))
	if err != nil {
		return err
	}
	_28 := Client2Server_babyLearnSkill{}
	_28.instId = instId
	_28.oldSkId = oldSkId
	_28.newSkId = newSkId
	_28.newSkLv = newSkLv
	err = _28.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) forgetSkill(skid uint32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(29))
	if err != nil {
		return err
	}
	_29 := Client2Server_forgetSkill{}
	_29.skid = skid
	err = _29.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) syncOrder(order COM_Order) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(30))
	if err != nil {
		return err
	}
	_30 := Client2Server_syncOrder{}
	_30.order = order
	err = _30.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) syncOrderTimeout() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(31))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) sendChat(content COM_Chat, targetName string) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(32))
	if err != nil {
		return err
	}
	_32 := Client2Server_sendChat{}
	_32.content = content
	_32.targetName = targetName
	err = _32.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) requestAudio(audioId int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(33))
	if err != nil {
		return err
	}
	_33 := Client2Server_requestAudio{}
	_33.audioId = audioId
	err = _33.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) publishItemInst(type_ int, itemInstId uint32, chatType int, playerName string) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(34))
	if err != nil {
		return err
	}
	_34 := Client2Server_publishItemInst{}
	_34.type_ = type_
	_34.itemInstId = itemInstId
	_34.chatType = chatType
	_34.playerName = playerName
	err = _34.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) queryItemInst(showId int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(35))
	if err != nil {
		return err
	}
	_35 := Client2Server_queryItemInst{}
	_35.showId = showId
	err = _35.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) publishbabyInst(type_ int, babyInstId uint32, playerName string) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(36))
	if err != nil {
		return err
	}
	_36 := Client2Server_publishbabyInst{}
	_36.type_ = type_
	_36.babyInstId = babyInstId
	_36.playerName = playerName
	err = _36.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) querybabyInst(showId int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(37))
	if err != nil {
		return err
	}
	_37 := Client2Server_querybabyInst{}
	_37.showId = showId
	err = _37.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) useItem(slot uint32, target uint32, stack uint32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(38))
	if err != nil {
		return err
	}
	_38 := Client2Server_useItem{}
	_38.slot = slot
	_38.target = target
	_38.stack = stack
	err = _38.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) wearEquipment(target uint32, itemInstId uint32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(39))
	if err != nil {
		return err
	}
	_39 := Client2Server_wearEquipment{}
	_39.target = target
	_39.itemInstId = itemInstId
	err = _39.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) delEquipment(target uint32, itemInstId uint32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(40))
	if err != nil {
		return err
	}
	_40 := Client2Server_delEquipment{}
	_40.target = target
	_40.itemInstId = itemInstId
	err = _40.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) setPlayerFront(isFront bool) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(41))
	if err != nil {
		return err
	}
	_41 := Client2Server_setPlayerFront{}
	_41.isFront = isFront
	err = _41.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) setBattlebaby(babyID uint32, isBattle bool) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(42))
	if err != nil {
		return err
	}
	_42 := Client2Server_setBattlebaby{}
	_42.babyID = babyID
	_42.isBattle = isBattle
	err = _42.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) changeBabyName(babyID uint32, name string) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(43))
	if err != nil {
		return err
	}
	_43 := Client2Server_changeBabyName{}
	_43.babyID = babyID
	_43.name = name
	err = _43.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) intensifyBaby(babyid uint32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(44))
	if err != nil {
		return err
	}
	_44 := Client2Server_intensifyBaby{}
	_44.babyid = babyid
	err = _44.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) setBattleEmp(empID uint32, group int, isBattle bool) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(45))
	if err != nil {
		return err
	}
	_45 := Client2Server_setBattleEmp{}
	_45.empID = empID
	_45.group = group
	_45.isBattle = isBattle
	err = _45.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) changeEmpBattleGroup(group int) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(46))
	if err != nil {
		return err
	}
	_46 := Client2Server_changeEmpBattleGroup{}
	_46.group = group
	err = _46.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) requestEvolve(empInstId uint32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(47))
	if err != nil {
		return err
	}
	_47 := Client2Server_requestEvolve{}
	_47.empInstId = empInstId
	err = _47.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) requestUpStar(empInstId uint32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(48))
	if err != nil {
		return err
	}
	_48 := Client2Server_requestUpStar{}
	_48.empInstId = empInstId
	err = _48.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) requestDelEmp(empInstId uint32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(49))
	if err != nil {
		return err
	}
	_49 := Client2Server_requestDelEmp{}
	_49.empInstId = empInstId
	err = _49.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) delEmployee(emps []uint32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(50))
	if err != nil {
		return err
	}
	_50 := Client2Server_delEmployee{}
	_50.emps = emps
	err = _50.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) onekeyDelEmp() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(51))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) delEmployeeSoul(instid uint32, soulNum uint32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(52))
	if err != nil {
		return err
	}
	_52 := Client2Server_delEmployeeSoul{}
	_52.instid = instid
	_52.soulNum = soulNum
	err = _52.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) sortBagItem() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(53))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) sellBagItem(instId uint32, stack uint32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(54))
	if err != nil {
		return err
	}
	_54 := Client2Server_sellBagItem{}
	_54.instId = instId
	_54.stack = stack
	err = _54.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) depositItemToStorage(instid uint32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(55))
	if err != nil {
		return err
	}
	_55 := Client2Server_depositItemToStorage{}
	_55.instid = instid
	err = _55.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) depositBabyToStorage(instid uint32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(56))
	if err != nil {
		return err
	}
	_56 := Client2Server_depositBabyToStorage{}
	_56.instid = instid
	err = _56.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) storageItemToBag(instid uint32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(57))
	if err != nil {
		return err
	}
	_57 := Client2Server_storageItemToBag{}
	_57.instid = instid
	err = _57.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) storageBabyToPlayer(instid uint32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(58))
	if err != nil {
		return err
	}
	_58 := Client2Server_storageBabyToPlayer{}
	_58.instid = instid
	err = _58.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) sortStorage(tp int) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(59))
	if err != nil {
		return err
	}
	_59 := Client2Server_sortStorage{}
	_59.tp = tp
	err = _59.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) delStorageBaby(instid uint32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(60))
	if err != nil {
		return err
	}
	_60 := Client2Server_delStorageBaby{}
	_60.instid = instid
	err = _60.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) jointLobby() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(61))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) exitLobby() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(62))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) createTeam(cti COM_CreateTeamInfo) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(63))
	if err != nil {
		return err
	}
	_63 := Client2Server_createTeam{}
	_63.cti = cti
	err = _63.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) changeTeam(info COM_CreateTeamInfo) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(64))
	if err != nil {
		return err
	}
	_64 := Client2Server_changeTeam{}
	_64.info = info
	err = _64.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) kickTeamMember(uuid uint32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(65))
	if err != nil {
		return err
	}
	_65 := Client2Server_kickTeamMember{}
	_65.uuid = uuid
	err = _65.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) changeTeamLeader(uuid uint32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(66))
	if err != nil {
		return err
	}
	_66 := Client2Server_changeTeamLeader{}
	_66.uuid = uuid
	err = _66.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) joinTeam(teamId uint32, pwd string) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(67))
	if err != nil {
		return err
	}
	_67 := Client2Server_joinTeam{}
	_67.teamId = teamId
	_67.pwd = pwd
	err = _67.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) exitTeam() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(68))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) changeTeamPassword(pwd string) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(69))
	if err != nil {
		return err
	}
	_69 := Client2Server_changeTeamPassword{}
	_69.pwd = pwd
	err = _69.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) joinTeamRoom() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(70))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) inviteTeamMember(name string) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(71))
	if err != nil {
		return err
	}
	_71 := Client2Server_inviteTeamMember{}
	_71.name = name
	err = _71.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) isjoinTeam(teamId uint32, isFlag bool) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(72))
	if err != nil {
		return err
	}
	_72 := Client2Server_isjoinTeam{}
	_72.teamId = teamId
	_72.isFlag = isFlag
	err = _72.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) leaveTeam() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(73))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) backTeam() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(74))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) refuseBackTeam() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(75))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) teamCallMember(playerId int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(76))
	if err != nil {
		return err
	}
	_76 := Client2Server_teamCallMember{}
	_76.playerId = playerId
	err = _76.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) requestJoinTeam(targetName string) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(77))
	if err != nil {
		return err
	}
	_77 := Client2Server_requestJoinTeam{}
	_77.targetName = targetName
	err = _77.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) ratifyJoinTeam(sendName string) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(78))
	if err != nil {
		return err
	}
	_78 := Client2Server_ratifyJoinTeam{}
	_78.sendName = sendName
	err = _78.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) drawLotteryBox(type_ int, isFree bool) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(79))
	if err != nil {
		return err
	}
	_79 := Client2Server_drawLotteryBox{}
	_79.type_ = type_
	_79.isFree = isFree
	err = _79.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) acceptQuest(questId int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(80))
	if err != nil {
		return err
	}
	_80 := Client2Server_acceptQuest{}
	_80.questId = questId
	err = _80.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) submitQuest(npcId int32, questId int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(81))
	if err != nil {
		return err
	}
	_81 := Client2Server_submitQuest{}
	_81.npcId = npcId
	_81.questId = questId
	err = _81.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) giveupQuest(questId int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(82))
	if err != nil {
		return err
	}
	_82 := Client2Server_giveupQuest{}
	_82.questId = questId
	err = _82.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) requestContactInfoById(instId uint32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(83))
	if err != nil {
		return err
	}
	_83 := Client2Server_requestContactInfoById{}
	_83.instId = instId
	err = _83.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) requestContactInfoByName(instName string) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(84))
	if err != nil {
		return err
	}
	_84 := Client2Server_requestContactInfoByName{}
	_84.instName = instName
	err = _84.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) requestFriendList() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(85))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) addFriend(instId uint32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(86))
	if err != nil {
		return err
	}
	_86 := Client2Server_addFriend{}
	_86.instId = instId
	err = _86.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) delFriend(instId uint32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(87))
	if err != nil {
		return err
	}
	_87 := Client2Server_delFriend{}
	_87.instId = instId
	err = _87.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) addBlacklist(instId uint32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(88))
	if err != nil {
		return err
	}
	_88 := Client2Server_addBlacklist{}
	_88.instId = instId
	err = _88.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) delBlacklist(instId uint32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(89))
	if err != nil {
		return err
	}
	_89 := Client2Server_delBlacklist{}
	_89.instId = instId
	err = _89.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) requestReferrFriend() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(90))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) mining(gatherId int32, times int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(91))
	if err != nil {
		return err
	}
	_91 := Client2Server_mining{}
	_91.gatherId = gatherId
	_91.times = times
	err = _91.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) compoundItem(itemId int32, gemId int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(92))
	if err != nil {
		return err
	}
	_92 := Client2Server_compoundItem{}
	_92.itemId = itemId
	_92.gemId = gemId
	err = _92.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) bagItemSplit(instId int32, splitNum int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(93))
	if err != nil {
		return err
	}
	_93 := Client2Server_bagItemSplit{}
	_93.instId = instId
	_93.splitNum = splitNum
	err = _93.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) requestChallenge(name string) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(94))
	if err != nil {
		return err
	}
	_94 := Client2Server_requestChallenge{}
	_94.name = name
	err = _94.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) requestRival() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(95))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) requestMySelfJJCData() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(96))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) requestCheckMsg(name string) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(97))
	if err != nil {
		return err
	}
	_97 := Client2Server_requestCheckMsg{}
	_97.name = name
	err = _97.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) requestMyAllbattleMsg() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(98))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) requestJJCRank() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(99))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) requestLevelRank() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(100))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) requestBabyRank() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(101))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) requestEmpRank() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(102))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) requestPlayerFFRank() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(103))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) queryOnlinePlayerbyName(name string) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(104))
	if err != nil {
		return err
	}
	_104 := Client2Server_queryOnlinePlayerbyName{}
	_104.name = name
	err = _104.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) queryPlayerbyName(name string) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(105))
	if err != nil {
		return err
	}
	_105 := Client2Server_queryPlayerbyName{}
	_105.name = name
	err = _105.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) queryBaby(instId uint32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(106))
	if err != nil {
		return err
	}
	_106 := Client2Server_queryBaby{}
	_106.instId = instId
	err = _106.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) queryEmployee(instId uint32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(107))
	if err != nil {
		return err
	}
	_107 := Client2Server_queryEmployee{}
	_107.instId = instId
	err = _107.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) guideFinish(guideIdx uint64) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(108))
	if err != nil {
		return err
	}
	_108 := Client2Server_guideFinish{}
	_108.guideIdx = guideIdx
	err = _108.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) enterBattle(battleId int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(109))
	if err != nil {
		return err
	}
	_109 := Client2Server_enterBattle{}
	_109.battleId = battleId
	err = _109.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) shopBuyItem(id int32, num int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(110))
	if err != nil {
		return err
	}
	_110 := Client2Server_shopBuyItem{}
	_110.id = id
	_110.num = num
	err = _110.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) getFirstRechargeItem() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(111))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) requestLevelGift(level int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(112))
	if err != nil {
		return err
	}
	_112 := Client2Server_requestLevelGift{}
	_112.level = level
	err = _112.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) setCurrentTitle(title int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(113))
	if err != nil {
		return err
	}
	_113 := Client2Server_setCurrentTitle{}
	_113.title = title
	err = _113.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) openBuyBox() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(114))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) requestAchaward(achId int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(115))
	if err != nil {
		return err
	}
	_115 := Client2Server_requestAchaward{}
	_115.achId = achId
	err = _115.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) sign(index int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(116))
	if err != nil {
		return err
	}
	_116 := Client2Server_sign{}
	_116.index = index
	err = _116.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) requestSignupReward7() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(117))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) requestSignupReward14() error {
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
func (this *Client2ServerStub) requestSignupReward28() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(119))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) requestActivityReward(index int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(120))
	if err != nil {
		return err
	}
	_120 := Client2Server_requestActivityReward{}
	_120.index = index
	err = _120.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) resetHundredTier() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(121))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) enterHundredScene(level int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(122))
	if err != nil {
		return err
	}
	_122 := Client2Server_enterHundredScene{}
	_122.level = level
	err = _122.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) delBaby(instId int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(123))
	if err != nil {
		return err
	}
	_123 := Client2Server_delBaby{}
	_123.instId = instId
	err = _123.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) resetBaby(instId int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(124))
	if err != nil {
		return err
	}
	_124 := Client2Server_resetBaby{}
	_124.instId = instId
	err = _124.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) resetBabyProp(instId int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(125))
	if err != nil {
		return err
	}
	_125 := Client2Server_resetBabyProp{}
	_125.instId = instId
	err = _125.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) remouldBaby(instid int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(126))
	if err != nil {
		return err
	}
	_126 := Client2Server_remouldBaby{}
	_126.instid = instid
	err = _126.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) empSkillLevelUp(empId uint32, skillId int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(127))
	if err != nil {
		return err
	}
	_127 := Client2Server_empSkillLevelUp{}
	_127.empId = empId
	_127.skillId = skillId
	err = _127.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) setOpenDoubleTimeFlag(isFlag bool) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(128))
	if err != nil {
		return err
	}
	_128 := Client2Server_setOpenDoubleTimeFlag{}
	_128.isFlag = isFlag
	err = _128.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) talkedNpc(npcId int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(129))
	if err != nil {
		return err
	}
	_129 := Client2Server_talkedNpc{}
	_129.npcId = npcId
	err = _129.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) jjcBattleGo(id uint32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(130))
	if err != nil {
		return err
	}
	_130 := Client2Server_jjcBattleGo{}
	_130.id = id
	err = _130.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) requestMyJJCTeamMsg() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(131))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) startMatching() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(132))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) stopMatching() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(133))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) exitPvpJJC() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(134))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) joinPvpLobby() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(135))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) exitPvpLobby() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(136))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) requestpvprank() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(137))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) joinWarriorchoose() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(138))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) warriorStart() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(139))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) warriorStop() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(140))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) sendMail(playername string, title string, content string) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(141))
	if err != nil {
		return err
	}
	_141 := Client2Server_sendMail{}
	_141.playername = playername
	_141.title = title
	_141.content = content
	err = _141.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) readMail(mailId int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(142))
	if err != nil {
		return err
	}
	_142 := Client2Server_readMail{}
	_142.mailId = mailId
	err = _142.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) delMail(mailId int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(143))
	if err != nil {
		return err
	}
	_143 := Client2Server_delMail{}
	_143.mailId = mailId
	err = _143.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) getMailItem(mailId int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(144))
	if err != nil {
		return err
	}
	_144 := Client2Server_getMailItem{}
	_144.mailId = mailId
	err = _144.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) requestState() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(145))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) createGuild(guildName string) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(146))
	if err != nil {
		return err
	}
	_146 := Client2Server_createGuild{}
	_146.guildName = guildName
	err = _146.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) delGuild(guildId uint32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(147))
	if err != nil {
		return err
	}
	_147 := Client2Server_delGuild{}
	_147.guildId = guildId
	err = _147.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) requestJoinGuild(guid uint32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(148))
	if err != nil {
		return err
	}
	_148 := Client2Server_requestJoinGuild{}
	_148.guid = guid
	err = _148.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) leaveGuild() error {
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
func (this *Client2ServerStub) kickOut(guid int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(150))
	if err != nil {
		return err
	}
	_150 := Client2Server_kickOut{}
	_150.guid = guid
	err = _150.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) acceptRequestGuild(playerId int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(151))
	if err != nil {
		return err
	}
	_151 := Client2Server_acceptRequestGuild{}
	_151.playerId = playerId
	err = _151.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) refuseRequestGuild(playerId int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(152))
	if err != nil {
		return err
	}
	_152 := Client2Server_refuseRequestGuild{}
	_152.playerId = playerId
	err = _152.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) changeMemberPosition(targetId int32, job int) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(153))
	if err != nil {
		return err
	}
	_153 := Client2Server_changeMemberPosition{}
	_153.targetId = targetId
	_153.job = job
	err = _153.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) transferPremier(targetId int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(154))
	if err != nil {
		return err
	}
	_154 := Client2Server_transferPremier{}
	_154.targetId = targetId
	err = _154.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) changeGuildNotice(notice string) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(155))
	if err != nil {
		return err
	}
	_155 := Client2Server_changeGuildNotice{}
	_155.notice = notice
	err = _155.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) queryGuildList(page int16) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(156))
	if err != nil {
		return err
	}
	_156 := Client2Server_queryGuildList{}
	_156.page = page
	err = _156.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) inviteJoinGuild(playerName string) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(157))
	if err != nil {
		return err
	}
	_157 := Client2Server_inviteJoinGuild{}
	_157.playerName = playerName
	err = _157.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) respondInviteJoinGuild(sendName string) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(158))
	if err != nil {
		return err
	}
	_158 := Client2Server_respondInviteJoinGuild{}
	_158.sendName = sendName
	err = _158.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) buyGuildItem(tableId int32, times int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(159))
	if err != nil {
		return err
	}
	_159 := Client2Server_buyGuildItem{}
	_159.tableId = tableId
	_159.times = times
	err = _159.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) entryGuildBattle() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(160))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) transforGuildBattleScene() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(161))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) addGuildMoney(money int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(162))
	if err != nil {
		return err
	}
	_162 := Client2Server_addGuildMoney{}
	_162.money = money
	err = _162.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) updateGuildBuiling(gbt int) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(163))
	if err != nil {
		return err
	}
	_163 := Client2Server_updateGuildBuiling{}
	_163.gbt = gbt
	err = _163.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) refreshGuildShop() error {
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
func (this *Client2ServerStub) levelupGuildSkill(skId int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(165))
	if err != nil {
		return err
	}
	_165 := Client2Server_levelupGuildSkill{}
	_165.skId = skId
	err = _165.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) presentGuildItem(num int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(166))
	if err != nil {
		return err
	}
	_166 := Client2Server_presentGuildItem{}
	_166.num = num
	err = _166.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) progenitusAddExp(monsterId int32, isSuper bool) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(167))
	if err != nil {
		return err
	}
	_167 := Client2Server_progenitusAddExp{}
	_167.monsterId = monsterId
	_167.isSuper = isSuper
	err = _167.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) setProgenitusPosition(mId int32, pos int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(168))
	if err != nil {
		return err
	}
	_168 := Client2Server_setProgenitusPosition{}
	_168.mId = mId
	_168.pos = pos
	err = _168.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) guildsign() error {
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
func (this *Client2ServerStub) fetchSelling(context COM_SearchContext) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(170))
	if err != nil {
		return err
	}
	_170 := Client2Server_fetchSelling{}
	_170.context = context
	err = _170.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) fetchSelling2(context COM_SearchContext) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(171))
	if err != nil {
		return err
	}
	_171 := Client2Server_fetchSelling2{}
	_171.context = context
	err = _171.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) selling(iteminstid int32, babyinstid int32, price int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(172))
	if err != nil {
		return err
	}
	_172 := Client2Server_selling{}
	_172.iteminstid = iteminstid
	_172.babyinstid = babyinstid
	_172.price = price
	err = _172.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) unselling(sellid int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(173))
	if err != nil {
		return err
	}
	_173 := Client2Server_unselling{}
	_173.sellid = sellid
	err = _173.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) buy(sellid int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(174))
	if err != nil {
		return err
	}
	_174 := Client2Server_buy{}
	_174.sellid = sellid
	err = _174.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) fixItem(instId int32, type_ int) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(175))
	if err != nil {
		return err
	}
	_175 := Client2Server_fixItem{}
	_175.instId = instId
	_175.type_ = type_
	err = _175.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) fixAllItem(items []uint32, type_ int) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(176))
	if err != nil {
		return err
	}
	_176 := Client2Server_fixAllItem{}
	_176.items = items
	_176.type_ = type_
	err = _176.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) makeDebirsItem(instId int32, num int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(177))
	if err != nil {
		return err
	}
	_177 := Client2Server_makeDebirsItem{}
	_177.instId = instId
	_177.num = num
	err = _177.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) levelUpMagicItem(items []uint32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(178))
	if err != nil {
		return err
	}
	_178 := Client2Server_levelUpMagicItem{}
	_178.items = items
	err = _178.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) tupoMagicItem(level int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(179))
	if err != nil {
		return err
	}
	_179 := Client2Server_tupoMagicItem{}
	_179.level = level
	err = _179.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) changeMagicJob(job int) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(180))
	if err != nil {
		return err
	}
	_180 := Client2Server_changeMagicJob{}
	_180.job = job
	err = _180.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) requestPk(playerId uint32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(181))
	if err != nil {
		return err
	}
	_181 := Client2Server_requestPk{}
	_181.playerId = playerId
	err = _181.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) uiBehavior(type_ int) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(182))
	if err != nil {
		return err
	}
	_182 := Client2Server_uiBehavior{}
	_182.type_ = type_
	err = _182.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) openZhuanpan() error {
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
func (this *Client2ServerStub) zhuanpanGo(counter uint32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(184))
	if err != nil {
		return err
	}
	_184 := Client2Server_zhuanpanGo{}
	_184.counter = counter
	err = _184.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) redemptionSpree(code string) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(185))
	if err != nil {
		return err
	}
	_185 := Client2Server_redemptionSpree{}
	_185.code = code
	err = _185.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) sceneFilter(sfType []int) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(186))
	if err != nil {
		return err
	}
	_186 := Client2Server_sceneFilter{}
	_186.sfType = sfType
	err = _186.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) sendExamAnswer(questionId uint32, answer uint8) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(187))
	if err != nil {
		return err
	}
	_187 := Client2Server_sendExamAnswer{}
	_187.questionId = questionId
	_187.answer = answer
	err = _187.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) sendwishing(wish COM_Wishing) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(188))
	if err != nil {
		return err
	}
	_188 := Client2Server_sendwishing{}
	_188.wish = wish
	err = _188.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) requestWish() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(189))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) leaderCloseDialog() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(190))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) requestOnlineReward(index uint32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(191))
	if err != nil {
		return err
	}
	_191 := Client2Server_requestOnlineReward{}
	_191.index = index
	err = _191.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) requestFundReward(level uint32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(192))
	if err != nil {
		return err
	}
	_192 := Client2Server_requestFundReward{}
	_192.level = level
	err = _192.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) openCard(index uint16) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(193))
	if err != nil {
		return err
	}
	_193 := Client2Server_openCard{}
	_193.index = index
	err = _193.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) resetCard() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(194))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) hotRoleBuy() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(195))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) requestSevenReward(qid uint32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(196))
	if err != nil {
		return err
	}
	_196 := Client2Server_requestSevenReward{}
	_196.qid = qid
	err = _196.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) vipreward() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(197))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) requestChargeTotalSingleReward(index uint32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(198))
	if err != nil {
		return err
	}
	_198 := Client2Server_requestChargeTotalSingleReward{}
	_198.index = index
	err = _198.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) requestChargeTotalReward(index uint32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(199))
	if err != nil {
		return err
	}
	_199 := Client2Server_requestChargeTotalReward{}
	_199.index = index
	err = _199.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) requestChargeEverySingleReward(index uint32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(200))
	if err != nil {
		return err
	}
	_200 := Client2Server_requestChargeEverySingleReward{}
	_200.index = index
	err = _200.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) requestChargeEveryReward(index uint32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(201))
	if err != nil {
		return err
	}
	_201 := Client2Server_requestChargeEveryReward{}
	_201.index = index
	err = _201.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) requestLoginTotal(index uint32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(202))
	if err != nil {
		return err
	}
	_202 := Client2Server_requestLoginTotal{}
	_202.index = index
	err = _202.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) buyDiscountStoreSingle(itemId int32, itemStack int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(203))
	if err != nil {
		return err
	}
	_203 := Client2Server_buyDiscountStoreSingle{}
	_203.itemId = itemId
	_203.itemStack = itemStack
	err = _203.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) buyDiscountStore(itemId int32, itemStack int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(204))
	if err != nil {
		return err
	}
	_204 := Client2Server_buyDiscountStore{}
	_204.itemId = itemId
	_204.itemStack = itemStack
	err = _204.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) requestEmployeeActivityReward(index uint32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(205))
	if err != nil {
		return err
	}
	_205 := Client2Server_requestEmployeeActivityReward{}
	_205.index = index
	err = _205.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) requestmyselfrechargeleReward(index uint32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(206))
	if err != nil {
		return err
	}
	_206 := Client2Server_requestmyselfrechargeleReward{}
	_206.index = index
	err = _206.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) requestEverydayIntegral() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(207))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) buyIntegralItem(id uint32, num uint32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(208))
	if err != nil {
		return err
	}
	_208 := Client2Server_buyIntegralItem{}
	_208.id = id
	_208.num = num
	err = _208.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) familyLoseLeader() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(209))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) verificationSMS(phoneNumber string, code string) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(210))
	if err != nil {
		return err
	}
	_210 := Client2Server_verificationSMS{}
	_210.phoneNumber = phoneNumber
	_210.code = code
	err = _210.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) lockItem(instId int32, isLock bool) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(211))
	if err != nil {
		return err
	}
	_211 := Client2Server_lockItem{}
	_211.instId = instId
	_211.isLock = isLock
	err = _211.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) lockBaby(instId int32, isLock bool) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(212))
	if err != nil {
		return err
	}
	_212 := Client2Server_lockBaby{}
	_212.instId = instId
	_212.isLock = isLock
	err = _212.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) showBaby(instId int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(213))
	if err != nil {
		return err
	}
	_213 := Client2Server_showBaby{}
	_213.instId = instId
	err = _213.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) wearFuwen(itemInstId int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(214))
	if err != nil {
		return err
	}
	_214 := Client2Server_wearFuwen{}
	_214.itemInstId = itemInstId
	err = _214.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) takeoffFuwen(slotId int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(215))
	if err != nil {
		return err
	}
	_215 := Client2Server_takeoffFuwen{}
	_215.slotId = slotId
	err = _215.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) compFuwen(itemInstId int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(216))
	if err != nil {
		return err
	}
	_216 := Client2Server_compFuwen{}
	_216.itemInstId = itemInstId
	err = _216.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) requestEmployeeQuest() error {
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
func (this *Client2ServerStub) acceptEmployeeQuest(questId int32, employees []int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(218))
	if err != nil {
		return err
	}
	_218 := Client2Server_acceptEmployeeQuest{}
	_218.questId = questId
	_218.employees = employees
	err = _218.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) submitEmployeeQuest(questId int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(219))
	if err != nil {
		return err
	}
	_219 := Client2Server_submitEmployeeQuest{}
	_219.questId = questId
	err = _219.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) crystalUpLevel() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(220))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) resetCrystalProp(locklist []int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(221))
	if err != nil {
		return err
	}
	_221 := Client2Server_resetCrystalProp{}
	_221.locklist = locklist
	err = _221.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *Client2ServerStub) magicItemOneKeyLevel() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(222))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func Bridging_Client2Server_openvip(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_0 := Client2Server_openvip{}
	err := _0.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.openvip(_0.vl)
}
func Bridging_Client2Server_requestPhoto(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.requestPhoto()
}
func Bridging_Client2Server_ping(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.ping()
}
func Bridging_Client2Server_sessionlogin(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_3 := Client2Server_sessionlogin{}
	err := _3.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.sessionlogin(_3.info)
}
func Bridging_Client2Server_login(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_4 := Client2Server_login{}
	err := _4.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.login(_4.info)
}
func Bridging_Client2Server_createPlayer(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_5 := Client2Server_createPlayer{}
	err := _5.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.createPlayer(_5.playername, _5.playerTmpId)
}
func Bridging_Client2Server_deletePlayer(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_6 := Client2Server_deletePlayer{}
	err := _6.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.deletePlayer(_6.playername)
}
func Bridging_Client2Server_enterGame(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_7 := Client2Server_enterGame{}
	err := _7.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.enterGame(_7.playerInstId)
}
func Bridging_Client2Server_requestBag(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.requestBag()
}
func Bridging_Client2Server_requestEmployees(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.requestEmployees()
}
func Bridging_Client2Server_requestStorage(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_10 := Client2Server_requestStorage{}
	err := _10.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.requestStorage(_10.tp)
}
func Bridging_Client2Server_requestAchievement(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.requestAchievement()
}
func Bridging_Client2Server_initminig(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.initminig()
}
func Bridging_Client2Server_requestCompound(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.requestCompound()
}
func Bridging_Client2Server_move(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_14 := Client2Server_move{}
	err := _14.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.move(_14.x, _14.z)
}
func Bridging_Client2Server_moveToNpc(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_15 := Client2Server_moveToNpc{}
	err := _15.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.moveToNpc(_15.npcid)
}
func Bridging_Client2Server_moveToNpc2(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_16 := Client2Server_moveToNpc2{}
	err := _16.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.moveToNpc2(_16.type_)
}
func Bridging_Client2Server_moveToZone(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_17 := Client2Server_moveToZone{}
	err := _17.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.moveToZone(_17.sceneId, _17.zoneId)
}
func Bridging_Client2Server_autoBattle(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.autoBattle()
}
func Bridging_Client2Server_stopAutoBattle(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.stopAutoBattle()
}
func Bridging_Client2Server_stopMove(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.stopMove()
}
func Bridging_Client2Server_exitCopy(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.exitCopy()
}
func Bridging_Client2Server_transforScene(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_22 := Client2Server_transforScene{}
	err := _22.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.transforScene(_22.sceneId)
}
func Bridging_Client2Server_sceneLoaded(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.sceneLoaded()
}
func Bridging_Client2Server_querySimplePlayerInst(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_24 := Client2Server_querySimplePlayerInst{}
	err := _24.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.querySimplePlayerInst(_24.instId)
}
func Bridging_Client2Server_logout(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.logout()
}
func Bridging_Client2Server_changProp(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_26 := Client2Server_changProp{}
	err := _26.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.changProp(_26.guid, _26.props)
}
func Bridging_Client2Server_learnSkill(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_27 := Client2Server_learnSkill{}
	err := _27.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.learnSkill(_27.skid)
}
func Bridging_Client2Server_babyLearnSkill(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_28 := Client2Server_babyLearnSkill{}
	err := _28.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.babyLearnSkill(_28.instId, _28.oldSkId, _28.newSkId, _28.newSkLv)
}
func Bridging_Client2Server_forgetSkill(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_29 := Client2Server_forgetSkill{}
	err := _29.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.forgetSkill(_29.skid)
}
func Bridging_Client2Server_syncOrder(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_30 := Client2Server_syncOrder{}
	err := _30.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.syncOrder(_30.order)
}
func Bridging_Client2Server_syncOrderTimeout(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.syncOrderTimeout()
}
func Bridging_Client2Server_sendChat(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_32 := Client2Server_sendChat{}
	err := _32.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.sendChat(_32.content, _32.targetName)
}
func Bridging_Client2Server_requestAudio(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_33 := Client2Server_requestAudio{}
	err := _33.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.requestAudio(_33.audioId)
}
func Bridging_Client2Server_publishItemInst(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_34 := Client2Server_publishItemInst{}
	err := _34.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.publishItemInst(_34.type_, _34.itemInstId, _34.chatType, _34.playerName)
}
func Bridging_Client2Server_queryItemInst(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_35 := Client2Server_queryItemInst{}
	err := _35.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.queryItemInst(_35.showId)
}
func Bridging_Client2Server_publishbabyInst(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_36 := Client2Server_publishbabyInst{}
	err := _36.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.publishbabyInst(_36.type_, _36.babyInstId, _36.playerName)
}
func Bridging_Client2Server_querybabyInst(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_37 := Client2Server_querybabyInst{}
	err := _37.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.querybabyInst(_37.showId)
}
func Bridging_Client2Server_useItem(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_38 := Client2Server_useItem{}
	err := _38.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.useItem(_38.slot, _38.target, _38.stack)
}
func Bridging_Client2Server_wearEquipment(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_39 := Client2Server_wearEquipment{}
	err := _39.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.wearEquipment(_39.target, _39.itemInstId)
}
func Bridging_Client2Server_delEquipment(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_40 := Client2Server_delEquipment{}
	err := _40.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.delEquipment(_40.target, _40.itemInstId)
}
func Bridging_Client2Server_setPlayerFront(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_41 := Client2Server_setPlayerFront{}
	err := _41.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.setPlayerFront(_41.isFront)
}
func Bridging_Client2Server_setBattlebaby(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_42 := Client2Server_setBattlebaby{}
	err := _42.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.setBattlebaby(_42.babyID, _42.isBattle)
}
func Bridging_Client2Server_changeBabyName(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_43 := Client2Server_changeBabyName{}
	err := _43.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.changeBabyName(_43.babyID, _43.name)
}
func Bridging_Client2Server_intensifyBaby(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_44 := Client2Server_intensifyBaby{}
	err := _44.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.intensifyBaby(_44.babyid)
}
func Bridging_Client2Server_setBattleEmp(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_45 := Client2Server_setBattleEmp{}
	err := _45.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.setBattleEmp(_45.empID, _45.group, _45.isBattle)
}
func Bridging_Client2Server_changeEmpBattleGroup(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_46 := Client2Server_changeEmpBattleGroup{}
	err := _46.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.changeEmpBattleGroup(_46.group)
}
func Bridging_Client2Server_requestEvolve(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_47 := Client2Server_requestEvolve{}
	err := _47.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.requestEvolve(_47.empInstId)
}
func Bridging_Client2Server_requestUpStar(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_48 := Client2Server_requestUpStar{}
	err := _48.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.requestUpStar(_48.empInstId)
}
func Bridging_Client2Server_requestDelEmp(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_49 := Client2Server_requestDelEmp{}
	err := _49.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.requestDelEmp(_49.empInstId)
}
func Bridging_Client2Server_delEmployee(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_50 := Client2Server_delEmployee{}
	err := _50.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.delEmployee(_50.emps)
}
func Bridging_Client2Server_onekeyDelEmp(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.onekeyDelEmp()
}
func Bridging_Client2Server_delEmployeeSoul(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_52 := Client2Server_delEmployeeSoul{}
	err := _52.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.delEmployeeSoul(_52.instid, _52.soulNum)
}
func Bridging_Client2Server_sortBagItem(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.sortBagItem()
}
func Bridging_Client2Server_sellBagItem(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_54 := Client2Server_sellBagItem{}
	err := _54.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.sellBagItem(_54.instId, _54.stack)
}
func Bridging_Client2Server_depositItemToStorage(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_55 := Client2Server_depositItemToStorage{}
	err := _55.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.depositItemToStorage(_55.instid)
}
func Bridging_Client2Server_depositBabyToStorage(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_56 := Client2Server_depositBabyToStorage{}
	err := _56.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.depositBabyToStorage(_56.instid)
}
func Bridging_Client2Server_storageItemToBag(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_57 := Client2Server_storageItemToBag{}
	err := _57.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.storageItemToBag(_57.instid)
}
func Bridging_Client2Server_storageBabyToPlayer(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_58 := Client2Server_storageBabyToPlayer{}
	err := _58.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.storageBabyToPlayer(_58.instid)
}
func Bridging_Client2Server_sortStorage(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_59 := Client2Server_sortStorage{}
	err := _59.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.sortStorage(_59.tp)
}
func Bridging_Client2Server_delStorageBaby(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_60 := Client2Server_delStorageBaby{}
	err := _60.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.delStorageBaby(_60.instid)
}
func Bridging_Client2Server_jointLobby(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.jointLobby()
}
func Bridging_Client2Server_exitLobby(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.exitLobby()
}
func Bridging_Client2Server_createTeam(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_63 := Client2Server_createTeam{}
	err := _63.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.createTeam(_63.cti)
}
func Bridging_Client2Server_changeTeam(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_64 := Client2Server_changeTeam{}
	err := _64.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.changeTeam(_64.info)
}
func Bridging_Client2Server_kickTeamMember(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_65 := Client2Server_kickTeamMember{}
	err := _65.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.kickTeamMember(_65.uuid)
}
func Bridging_Client2Server_changeTeamLeader(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_66 := Client2Server_changeTeamLeader{}
	err := _66.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.changeTeamLeader(_66.uuid)
}
func Bridging_Client2Server_joinTeam(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_67 := Client2Server_joinTeam{}
	err := _67.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.joinTeam(_67.teamId, _67.pwd)
}
func Bridging_Client2Server_exitTeam(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.exitTeam()
}
func Bridging_Client2Server_changeTeamPassword(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_69 := Client2Server_changeTeamPassword{}
	err := _69.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.changeTeamPassword(_69.pwd)
}
func Bridging_Client2Server_joinTeamRoom(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.joinTeamRoom()
}
func Bridging_Client2Server_inviteTeamMember(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_71 := Client2Server_inviteTeamMember{}
	err := _71.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.inviteTeamMember(_71.name)
}
func Bridging_Client2Server_isjoinTeam(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_72 := Client2Server_isjoinTeam{}
	err := _72.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.isjoinTeam(_72.teamId, _72.isFlag)
}
func Bridging_Client2Server_leaveTeam(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.leaveTeam()
}
func Bridging_Client2Server_backTeam(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.backTeam()
}
func Bridging_Client2Server_refuseBackTeam(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.refuseBackTeam()
}
func Bridging_Client2Server_teamCallMember(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_76 := Client2Server_teamCallMember{}
	err := _76.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.teamCallMember(_76.playerId)
}
func Bridging_Client2Server_requestJoinTeam(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_77 := Client2Server_requestJoinTeam{}
	err := _77.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.requestJoinTeam(_77.targetName)
}
func Bridging_Client2Server_ratifyJoinTeam(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_78 := Client2Server_ratifyJoinTeam{}
	err := _78.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.ratifyJoinTeam(_78.sendName)
}
func Bridging_Client2Server_drawLotteryBox(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_79 := Client2Server_drawLotteryBox{}
	err := _79.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.drawLotteryBox(_79.type_, _79.isFree)
}
func Bridging_Client2Server_acceptQuest(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_80 := Client2Server_acceptQuest{}
	err := _80.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.acceptQuest(_80.questId)
}
func Bridging_Client2Server_submitQuest(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_81 := Client2Server_submitQuest{}
	err := _81.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.submitQuest(_81.npcId, _81.questId)
}
func Bridging_Client2Server_giveupQuest(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_82 := Client2Server_giveupQuest{}
	err := _82.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.giveupQuest(_82.questId)
}
func Bridging_Client2Server_requestContactInfoById(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_83 := Client2Server_requestContactInfoById{}
	err := _83.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.requestContactInfoById(_83.instId)
}
func Bridging_Client2Server_requestContactInfoByName(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_84 := Client2Server_requestContactInfoByName{}
	err := _84.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.requestContactInfoByName(_84.instName)
}
func Bridging_Client2Server_requestFriendList(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.requestFriendList()
}
func Bridging_Client2Server_addFriend(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_86 := Client2Server_addFriend{}
	err := _86.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.addFriend(_86.instId)
}
func Bridging_Client2Server_delFriend(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_87 := Client2Server_delFriend{}
	err := _87.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.delFriend(_87.instId)
}
func Bridging_Client2Server_addBlacklist(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_88 := Client2Server_addBlacklist{}
	err := _88.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.addBlacklist(_88.instId)
}
func Bridging_Client2Server_delBlacklist(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_89 := Client2Server_delBlacklist{}
	err := _89.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.delBlacklist(_89.instId)
}
func Bridging_Client2Server_requestReferrFriend(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.requestReferrFriend()
}
func Bridging_Client2Server_mining(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_91 := Client2Server_mining{}
	err := _91.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.mining(_91.gatherId, _91.times)
}
func Bridging_Client2Server_compoundItem(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_92 := Client2Server_compoundItem{}
	err := _92.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.compoundItem(_92.itemId, _92.gemId)
}
func Bridging_Client2Server_bagItemSplit(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_93 := Client2Server_bagItemSplit{}
	err := _93.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.bagItemSplit(_93.instId, _93.splitNum)
}
func Bridging_Client2Server_requestChallenge(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_94 := Client2Server_requestChallenge{}
	err := _94.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.requestChallenge(_94.name)
}
func Bridging_Client2Server_requestRival(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.requestRival()
}
func Bridging_Client2Server_requestMySelfJJCData(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.requestMySelfJJCData()
}
func Bridging_Client2Server_requestCheckMsg(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_97 := Client2Server_requestCheckMsg{}
	err := _97.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.requestCheckMsg(_97.name)
}
func Bridging_Client2Server_requestMyAllbattleMsg(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.requestMyAllbattleMsg()
}
func Bridging_Client2Server_requestJJCRank(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.requestJJCRank()
}
func Bridging_Client2Server_requestLevelRank(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.requestLevelRank()
}
func Bridging_Client2Server_requestBabyRank(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.requestBabyRank()
}
func Bridging_Client2Server_requestEmpRank(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.requestEmpRank()
}
func Bridging_Client2Server_requestPlayerFFRank(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.requestPlayerFFRank()
}
func Bridging_Client2Server_queryOnlinePlayerbyName(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_104 := Client2Server_queryOnlinePlayerbyName{}
	err := _104.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.queryOnlinePlayerbyName(_104.name)
}
func Bridging_Client2Server_queryPlayerbyName(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_105 := Client2Server_queryPlayerbyName{}
	err := _105.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.queryPlayerbyName(_105.name)
}
func Bridging_Client2Server_queryBaby(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_106 := Client2Server_queryBaby{}
	err := _106.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.queryBaby(_106.instId)
}
func Bridging_Client2Server_queryEmployee(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_107 := Client2Server_queryEmployee{}
	err := _107.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.queryEmployee(_107.instId)
}
func Bridging_Client2Server_guideFinish(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_108 := Client2Server_guideFinish{}
	err := _108.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.guideFinish(_108.guideIdx)
}
func Bridging_Client2Server_enterBattle(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_109 := Client2Server_enterBattle{}
	err := _109.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.enterBattle(_109.battleId)
}
func Bridging_Client2Server_shopBuyItem(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_110 := Client2Server_shopBuyItem{}
	err := _110.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.shopBuyItem(_110.id, _110.num)
}
func Bridging_Client2Server_getFirstRechargeItem(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.getFirstRechargeItem()
}
func Bridging_Client2Server_requestLevelGift(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_112 := Client2Server_requestLevelGift{}
	err := _112.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.requestLevelGift(_112.level)
}
func Bridging_Client2Server_setCurrentTitle(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_113 := Client2Server_setCurrentTitle{}
	err := _113.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.setCurrentTitle(_113.title)
}
func Bridging_Client2Server_openBuyBox(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.openBuyBox()
}
func Bridging_Client2Server_requestAchaward(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_115 := Client2Server_requestAchaward{}
	err := _115.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.requestAchaward(_115.achId)
}
func Bridging_Client2Server_sign(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_116 := Client2Server_sign{}
	err := _116.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.sign(_116.index)
}
func Bridging_Client2Server_requestSignupReward7(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.requestSignupReward7()
}
func Bridging_Client2Server_requestSignupReward14(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.requestSignupReward14()
}
func Bridging_Client2Server_requestSignupReward28(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.requestSignupReward28()
}
func Bridging_Client2Server_requestActivityReward(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_120 := Client2Server_requestActivityReward{}
	err := _120.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.requestActivityReward(_120.index)
}
func Bridging_Client2Server_resetHundredTier(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.resetHundredTier()
}
func Bridging_Client2Server_enterHundredScene(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_122 := Client2Server_enterHundredScene{}
	err := _122.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.enterHundredScene(_122.level)
}
func Bridging_Client2Server_delBaby(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_123 := Client2Server_delBaby{}
	err := _123.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.delBaby(_123.instId)
}
func Bridging_Client2Server_resetBaby(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_124 := Client2Server_resetBaby{}
	err := _124.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.resetBaby(_124.instId)
}
func Bridging_Client2Server_resetBabyProp(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_125 := Client2Server_resetBabyProp{}
	err := _125.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.resetBabyProp(_125.instId)
}
func Bridging_Client2Server_remouldBaby(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_126 := Client2Server_remouldBaby{}
	err := _126.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.remouldBaby(_126.instid)
}
func Bridging_Client2Server_empSkillLevelUp(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_127 := Client2Server_empSkillLevelUp{}
	err := _127.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.empSkillLevelUp(_127.empId, _127.skillId)
}
func Bridging_Client2Server_setOpenDoubleTimeFlag(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_128 := Client2Server_setOpenDoubleTimeFlag{}
	err := _128.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.setOpenDoubleTimeFlag(_128.isFlag)
}
func Bridging_Client2Server_talkedNpc(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_129 := Client2Server_talkedNpc{}
	err := _129.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.talkedNpc(_129.npcId)
}
func Bridging_Client2Server_jjcBattleGo(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_130 := Client2Server_jjcBattleGo{}
	err := _130.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.jjcBattleGo(_130.id)
}
func Bridging_Client2Server_requestMyJJCTeamMsg(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.requestMyJJCTeamMsg()
}
func Bridging_Client2Server_startMatching(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.startMatching()
}
func Bridging_Client2Server_stopMatching(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.stopMatching()
}
func Bridging_Client2Server_exitPvpJJC(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.exitPvpJJC()
}
func Bridging_Client2Server_joinPvpLobby(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.joinPvpLobby()
}
func Bridging_Client2Server_exitPvpLobby(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.exitPvpLobby()
}
func Bridging_Client2Server_requestpvprank(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.requestpvprank()
}
func Bridging_Client2Server_joinWarriorchoose(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.joinWarriorchoose()
}
func Bridging_Client2Server_warriorStart(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.warriorStart()
}
func Bridging_Client2Server_warriorStop(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.warriorStop()
}
func Bridging_Client2Server_sendMail(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_141 := Client2Server_sendMail{}
	err := _141.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.sendMail(_141.playername, _141.title, _141.content)
}
func Bridging_Client2Server_readMail(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_142 := Client2Server_readMail{}
	err := _142.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.readMail(_142.mailId)
}
func Bridging_Client2Server_delMail(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_143 := Client2Server_delMail{}
	err := _143.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.delMail(_143.mailId)
}
func Bridging_Client2Server_getMailItem(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_144 := Client2Server_getMailItem{}
	err := _144.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.getMailItem(_144.mailId)
}
func Bridging_Client2Server_requestState(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.requestState()
}
func Bridging_Client2Server_createGuild(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_146 := Client2Server_createGuild{}
	err := _146.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.createGuild(_146.guildName)
}
func Bridging_Client2Server_delGuild(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_147 := Client2Server_delGuild{}
	err := _147.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.delGuild(_147.guildId)
}
func Bridging_Client2Server_requestJoinGuild(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_148 := Client2Server_requestJoinGuild{}
	err := _148.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.requestJoinGuild(_148.guid)
}
func Bridging_Client2Server_leaveGuild(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.leaveGuild()
}
func Bridging_Client2Server_kickOut(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_150 := Client2Server_kickOut{}
	err := _150.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.kickOut(_150.guid)
}
func Bridging_Client2Server_acceptRequestGuild(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_151 := Client2Server_acceptRequestGuild{}
	err := _151.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.acceptRequestGuild(_151.playerId)
}
func Bridging_Client2Server_refuseRequestGuild(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_152 := Client2Server_refuseRequestGuild{}
	err := _152.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.refuseRequestGuild(_152.playerId)
}
func Bridging_Client2Server_changeMemberPosition(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_153 := Client2Server_changeMemberPosition{}
	err := _153.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.changeMemberPosition(_153.targetId, _153.job)
}
func Bridging_Client2Server_transferPremier(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_154 := Client2Server_transferPremier{}
	err := _154.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.transferPremier(_154.targetId)
}
func Bridging_Client2Server_changeGuildNotice(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_155 := Client2Server_changeGuildNotice{}
	err := _155.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.changeGuildNotice(_155.notice)
}
func Bridging_Client2Server_queryGuildList(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_156 := Client2Server_queryGuildList{}
	err := _156.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.queryGuildList(_156.page)
}
func Bridging_Client2Server_inviteJoinGuild(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_157 := Client2Server_inviteJoinGuild{}
	err := _157.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.inviteJoinGuild(_157.playerName)
}
func Bridging_Client2Server_respondInviteJoinGuild(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_158 := Client2Server_respondInviteJoinGuild{}
	err := _158.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.respondInviteJoinGuild(_158.sendName)
}
func Bridging_Client2Server_buyGuildItem(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_159 := Client2Server_buyGuildItem{}
	err := _159.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.buyGuildItem(_159.tableId, _159.times)
}
func Bridging_Client2Server_entryGuildBattle(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.entryGuildBattle()
}
func Bridging_Client2Server_transforGuildBattleScene(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.transforGuildBattleScene()
}
func Bridging_Client2Server_addGuildMoney(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_162 := Client2Server_addGuildMoney{}
	err := _162.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.addGuildMoney(_162.money)
}
func Bridging_Client2Server_updateGuildBuiling(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_163 := Client2Server_updateGuildBuiling{}
	err := _163.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.updateGuildBuiling(_163.gbt)
}
func Bridging_Client2Server_refreshGuildShop(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.refreshGuildShop()
}
func Bridging_Client2Server_levelupGuildSkill(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_165 := Client2Server_levelupGuildSkill{}
	err := _165.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.levelupGuildSkill(_165.skId)
}
func Bridging_Client2Server_presentGuildItem(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_166 := Client2Server_presentGuildItem{}
	err := _166.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.presentGuildItem(_166.num)
}
func Bridging_Client2Server_progenitusAddExp(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_167 := Client2Server_progenitusAddExp{}
	err := _167.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.progenitusAddExp(_167.monsterId, _167.isSuper)
}
func Bridging_Client2Server_setProgenitusPosition(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_168 := Client2Server_setProgenitusPosition{}
	err := _168.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.setProgenitusPosition(_168.mId, _168.pos)
}
func Bridging_Client2Server_guildsign(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.guildsign()
}
func Bridging_Client2Server_fetchSelling(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_170 := Client2Server_fetchSelling{}
	err := _170.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.fetchSelling(_170.context)
}
func Bridging_Client2Server_fetchSelling2(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_171 := Client2Server_fetchSelling2{}
	err := _171.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.fetchSelling2(_171.context)
}
func Bridging_Client2Server_selling(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_172 := Client2Server_selling{}
	err := _172.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.selling(_172.iteminstid, _172.babyinstid, _172.price)
}
func Bridging_Client2Server_unselling(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_173 := Client2Server_unselling{}
	err := _173.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.unselling(_173.sellid)
}
func Bridging_Client2Server_buy(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_174 := Client2Server_buy{}
	err := _174.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.buy(_174.sellid)
}
func Bridging_Client2Server_fixItem(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_175 := Client2Server_fixItem{}
	err := _175.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.fixItem(_175.instId, _175.type_)
}
func Bridging_Client2Server_fixAllItem(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_176 := Client2Server_fixAllItem{}
	err := _176.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.fixAllItem(_176.items, _176.type_)
}
func Bridging_Client2Server_makeDebirsItem(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_177 := Client2Server_makeDebirsItem{}
	err := _177.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.makeDebirsItem(_177.instId, _177.num)
}
func Bridging_Client2Server_levelUpMagicItem(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_178 := Client2Server_levelUpMagicItem{}
	err := _178.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.levelUpMagicItem(_178.items)
}
func Bridging_Client2Server_tupoMagicItem(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_179 := Client2Server_tupoMagicItem{}
	err := _179.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.tupoMagicItem(_179.level)
}
func Bridging_Client2Server_changeMagicJob(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_180 := Client2Server_changeMagicJob{}
	err := _180.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.changeMagicJob(_180.job)
}
func Bridging_Client2Server_requestPk(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_181 := Client2Server_requestPk{}
	err := _181.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.requestPk(_181.playerId)
}
func Bridging_Client2Server_uiBehavior(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_182 := Client2Server_uiBehavior{}
	err := _182.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.uiBehavior(_182.type_)
}
func Bridging_Client2Server_openZhuanpan(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.openZhuanpan()
}
func Bridging_Client2Server_zhuanpanGo(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_184 := Client2Server_zhuanpanGo{}
	err := _184.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.zhuanpanGo(_184.counter)
}
func Bridging_Client2Server_redemptionSpree(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_185 := Client2Server_redemptionSpree{}
	err := _185.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.redemptionSpree(_185.code)
}
func Bridging_Client2Server_sceneFilter(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_186 := Client2Server_sceneFilter{}
	err := _186.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.sceneFilter(_186.sfType)
}
func Bridging_Client2Server_sendExamAnswer(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_187 := Client2Server_sendExamAnswer{}
	err := _187.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.sendExamAnswer(_187.questionId, _187.answer)
}
func Bridging_Client2Server_sendwishing(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_188 := Client2Server_sendwishing{}
	err := _188.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.sendwishing(_188.wish)
}
func Bridging_Client2Server_requestWish(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.requestWish()
}
func Bridging_Client2Server_leaderCloseDialog(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.leaderCloseDialog()
}
func Bridging_Client2Server_requestOnlineReward(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_191 := Client2Server_requestOnlineReward{}
	err := _191.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.requestOnlineReward(_191.index)
}
func Bridging_Client2Server_requestFundReward(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_192 := Client2Server_requestFundReward{}
	err := _192.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.requestFundReward(_192.level)
}
func Bridging_Client2Server_openCard(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_193 := Client2Server_openCard{}
	err := _193.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.openCard(_193.index)
}
func Bridging_Client2Server_resetCard(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.resetCard()
}
func Bridging_Client2Server_hotRoleBuy(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.hotRoleBuy()
}
func Bridging_Client2Server_requestSevenReward(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_196 := Client2Server_requestSevenReward{}
	err := _196.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.requestSevenReward(_196.qid)
}
func Bridging_Client2Server_vipreward(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.vipreward()
}
func Bridging_Client2Server_requestChargeTotalSingleReward(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_198 := Client2Server_requestChargeTotalSingleReward{}
	err := _198.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.requestChargeTotalSingleReward(_198.index)
}
func Bridging_Client2Server_requestChargeTotalReward(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_199 := Client2Server_requestChargeTotalReward{}
	err := _199.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.requestChargeTotalReward(_199.index)
}
func Bridging_Client2Server_requestChargeEverySingleReward(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_200 := Client2Server_requestChargeEverySingleReward{}
	err := _200.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.requestChargeEverySingleReward(_200.index)
}
func Bridging_Client2Server_requestChargeEveryReward(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_201 := Client2Server_requestChargeEveryReward{}
	err := _201.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.requestChargeEveryReward(_201.index)
}
func Bridging_Client2Server_requestLoginTotal(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_202 := Client2Server_requestLoginTotal{}
	err := _202.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.requestLoginTotal(_202.index)
}
func Bridging_Client2Server_buyDiscountStoreSingle(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_203 := Client2Server_buyDiscountStoreSingle{}
	err := _203.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.buyDiscountStoreSingle(_203.itemId, _203.itemStack)
}
func Bridging_Client2Server_buyDiscountStore(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_204 := Client2Server_buyDiscountStore{}
	err := _204.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.buyDiscountStore(_204.itemId, _204.itemStack)
}
func Bridging_Client2Server_requestEmployeeActivityReward(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_205 := Client2Server_requestEmployeeActivityReward{}
	err := _205.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.requestEmployeeActivityReward(_205.index)
}
func Bridging_Client2Server_requestmyselfrechargeleReward(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_206 := Client2Server_requestmyselfrechargeleReward{}
	err := _206.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.requestmyselfrechargeleReward(_206.index)
}
func Bridging_Client2Server_requestEverydayIntegral(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.requestEverydayIntegral()
}
func Bridging_Client2Server_buyIntegralItem(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_208 := Client2Server_buyIntegralItem{}
	err := _208.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.buyIntegralItem(_208.id, _208.num)
}
func Bridging_Client2Server_familyLoseLeader(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.familyLoseLeader()
}
func Bridging_Client2Server_verificationSMS(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_210 := Client2Server_verificationSMS{}
	err := _210.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.verificationSMS(_210.phoneNumber, _210.code)
}
func Bridging_Client2Server_lockItem(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_211 := Client2Server_lockItem{}
	err := _211.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.lockItem(_211.instId, _211.isLock)
}
func Bridging_Client2Server_lockBaby(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_212 := Client2Server_lockBaby{}
	err := _212.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.lockBaby(_212.instId, _212.isLock)
}
func Bridging_Client2Server_showBaby(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_213 := Client2Server_showBaby{}
	err := _213.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.showBaby(_213.instId)
}
func Bridging_Client2Server_wearFuwen(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_214 := Client2Server_wearFuwen{}
	err := _214.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.wearFuwen(_214.itemInstId)
}
func Bridging_Client2Server_takeoffFuwen(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_215 := Client2Server_takeoffFuwen{}
	err := _215.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.takeoffFuwen(_215.slotId)
}
func Bridging_Client2Server_compFuwen(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_216 := Client2Server_compFuwen{}
	err := _216.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.compFuwen(_216.itemInstId)
}
func Bridging_Client2Server_requestEmployeeQuest(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.requestEmployeeQuest()
}
func Bridging_Client2Server_acceptEmployeeQuest(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_218 := Client2Server_acceptEmployeeQuest{}
	err := _218.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.acceptEmployeeQuest(_218.questId, _218.employees)
}
func Bridging_Client2Server_submitEmployeeQuest(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_219 := Client2Server_submitEmployeeQuest{}
	err := _219.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.submitEmployeeQuest(_219.questId)
}
func Bridging_Client2Server_crystalUpLevel(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.crystalUpLevel()
}
func Bridging_Client2Server_resetCrystalProp(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_221 := Client2Server_resetCrystalProp{}
	err := _221.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.resetCrystalProp(_221.locklist)
}
func Bridging_Client2Server_magicItemOneKeyLevel(buffer *bytes.Buffer, p Client2ServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.magicItemOneKeyLevel()
}
func Client2ServerDispatch(buffer *bytes.Buffer, p Client2ServerProxy) error {
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
		return Bridging_Client2Server_openvip(buffer, p)
	case 1:
		return Bridging_Client2Server_requestPhoto(buffer, p)
	case 2:
		return Bridging_Client2Server_ping(buffer, p)
	case 3:
		return Bridging_Client2Server_sessionlogin(buffer, p)
	case 4:
		return Bridging_Client2Server_login(buffer, p)
	case 5:
		return Bridging_Client2Server_createPlayer(buffer, p)
	case 6:
		return Bridging_Client2Server_deletePlayer(buffer, p)
	case 7:
		return Bridging_Client2Server_enterGame(buffer, p)
	case 8:
		return Bridging_Client2Server_requestBag(buffer, p)
	case 9:
		return Bridging_Client2Server_requestEmployees(buffer, p)
	case 10:
		return Bridging_Client2Server_requestStorage(buffer, p)
	case 11:
		return Bridging_Client2Server_requestAchievement(buffer, p)
	case 12:
		return Bridging_Client2Server_initminig(buffer, p)
	case 13:
		return Bridging_Client2Server_requestCompound(buffer, p)
	case 14:
		return Bridging_Client2Server_move(buffer, p)
	case 15:
		return Bridging_Client2Server_moveToNpc(buffer, p)
	case 16:
		return Bridging_Client2Server_moveToNpc2(buffer, p)
	case 17:
		return Bridging_Client2Server_moveToZone(buffer, p)
	case 18:
		return Bridging_Client2Server_autoBattle(buffer, p)
	case 19:
		return Bridging_Client2Server_stopAutoBattle(buffer, p)
	case 20:
		return Bridging_Client2Server_stopMove(buffer, p)
	case 21:
		return Bridging_Client2Server_exitCopy(buffer, p)
	case 22:
		return Bridging_Client2Server_transforScene(buffer, p)
	case 23:
		return Bridging_Client2Server_sceneLoaded(buffer, p)
	case 24:
		return Bridging_Client2Server_querySimplePlayerInst(buffer, p)
	case 25:
		return Bridging_Client2Server_logout(buffer, p)
	case 26:
		return Bridging_Client2Server_changProp(buffer, p)
	case 27:
		return Bridging_Client2Server_learnSkill(buffer, p)
	case 28:
		return Bridging_Client2Server_babyLearnSkill(buffer, p)
	case 29:
		return Bridging_Client2Server_forgetSkill(buffer, p)
	case 30:
		return Bridging_Client2Server_syncOrder(buffer, p)
	case 31:
		return Bridging_Client2Server_syncOrderTimeout(buffer, p)
	case 32:
		return Bridging_Client2Server_sendChat(buffer, p)
	case 33:
		return Bridging_Client2Server_requestAudio(buffer, p)
	case 34:
		return Bridging_Client2Server_publishItemInst(buffer, p)
	case 35:
		return Bridging_Client2Server_queryItemInst(buffer, p)
	case 36:
		return Bridging_Client2Server_publishbabyInst(buffer, p)
	case 37:
		return Bridging_Client2Server_querybabyInst(buffer, p)
	case 38:
		return Bridging_Client2Server_useItem(buffer, p)
	case 39:
		return Bridging_Client2Server_wearEquipment(buffer, p)
	case 40:
		return Bridging_Client2Server_delEquipment(buffer, p)
	case 41:
		return Bridging_Client2Server_setPlayerFront(buffer, p)
	case 42:
		return Bridging_Client2Server_setBattlebaby(buffer, p)
	case 43:
		return Bridging_Client2Server_changeBabyName(buffer, p)
	case 44:
		return Bridging_Client2Server_intensifyBaby(buffer, p)
	case 45:
		return Bridging_Client2Server_setBattleEmp(buffer, p)
	case 46:
		return Bridging_Client2Server_changeEmpBattleGroup(buffer, p)
	case 47:
		return Bridging_Client2Server_requestEvolve(buffer, p)
	case 48:
		return Bridging_Client2Server_requestUpStar(buffer, p)
	case 49:
		return Bridging_Client2Server_requestDelEmp(buffer, p)
	case 50:
		return Bridging_Client2Server_delEmployee(buffer, p)
	case 51:
		return Bridging_Client2Server_onekeyDelEmp(buffer, p)
	case 52:
		return Bridging_Client2Server_delEmployeeSoul(buffer, p)
	case 53:
		return Bridging_Client2Server_sortBagItem(buffer, p)
	case 54:
		return Bridging_Client2Server_sellBagItem(buffer, p)
	case 55:
		return Bridging_Client2Server_depositItemToStorage(buffer, p)
	case 56:
		return Bridging_Client2Server_depositBabyToStorage(buffer, p)
	case 57:
		return Bridging_Client2Server_storageItemToBag(buffer, p)
	case 58:
		return Bridging_Client2Server_storageBabyToPlayer(buffer, p)
	case 59:
		return Bridging_Client2Server_sortStorage(buffer, p)
	case 60:
		return Bridging_Client2Server_delStorageBaby(buffer, p)
	case 61:
		return Bridging_Client2Server_jointLobby(buffer, p)
	case 62:
		return Bridging_Client2Server_exitLobby(buffer, p)
	case 63:
		return Bridging_Client2Server_createTeam(buffer, p)
	case 64:
		return Bridging_Client2Server_changeTeam(buffer, p)
	case 65:
		return Bridging_Client2Server_kickTeamMember(buffer, p)
	case 66:
		return Bridging_Client2Server_changeTeamLeader(buffer, p)
	case 67:
		return Bridging_Client2Server_joinTeam(buffer, p)
	case 68:
		return Bridging_Client2Server_exitTeam(buffer, p)
	case 69:
		return Bridging_Client2Server_changeTeamPassword(buffer, p)
	case 70:
		return Bridging_Client2Server_joinTeamRoom(buffer, p)
	case 71:
		return Bridging_Client2Server_inviteTeamMember(buffer, p)
	case 72:
		return Bridging_Client2Server_isjoinTeam(buffer, p)
	case 73:
		return Bridging_Client2Server_leaveTeam(buffer, p)
	case 74:
		return Bridging_Client2Server_backTeam(buffer, p)
	case 75:
		return Bridging_Client2Server_refuseBackTeam(buffer, p)
	case 76:
		return Bridging_Client2Server_teamCallMember(buffer, p)
	case 77:
		return Bridging_Client2Server_requestJoinTeam(buffer, p)
	case 78:
		return Bridging_Client2Server_ratifyJoinTeam(buffer, p)
	case 79:
		return Bridging_Client2Server_drawLotteryBox(buffer, p)
	case 80:
		return Bridging_Client2Server_acceptQuest(buffer, p)
	case 81:
		return Bridging_Client2Server_submitQuest(buffer, p)
	case 82:
		return Bridging_Client2Server_giveupQuest(buffer, p)
	case 83:
		return Bridging_Client2Server_requestContactInfoById(buffer, p)
	case 84:
		return Bridging_Client2Server_requestContactInfoByName(buffer, p)
	case 85:
		return Bridging_Client2Server_requestFriendList(buffer, p)
	case 86:
		return Bridging_Client2Server_addFriend(buffer, p)
	case 87:
		return Bridging_Client2Server_delFriend(buffer, p)
	case 88:
		return Bridging_Client2Server_addBlacklist(buffer, p)
	case 89:
		return Bridging_Client2Server_delBlacklist(buffer, p)
	case 90:
		return Bridging_Client2Server_requestReferrFriend(buffer, p)
	case 91:
		return Bridging_Client2Server_mining(buffer, p)
	case 92:
		return Bridging_Client2Server_compoundItem(buffer, p)
	case 93:
		return Bridging_Client2Server_bagItemSplit(buffer, p)
	case 94:
		return Bridging_Client2Server_requestChallenge(buffer, p)
	case 95:
		return Bridging_Client2Server_requestRival(buffer, p)
	case 96:
		return Bridging_Client2Server_requestMySelfJJCData(buffer, p)
	case 97:
		return Bridging_Client2Server_requestCheckMsg(buffer, p)
	case 98:
		return Bridging_Client2Server_requestMyAllbattleMsg(buffer, p)
	case 99:
		return Bridging_Client2Server_requestJJCRank(buffer, p)
	case 100:
		return Bridging_Client2Server_requestLevelRank(buffer, p)
	case 101:
		return Bridging_Client2Server_requestBabyRank(buffer, p)
	case 102:
		return Bridging_Client2Server_requestEmpRank(buffer, p)
	case 103:
		return Bridging_Client2Server_requestPlayerFFRank(buffer, p)
	case 104:
		return Bridging_Client2Server_queryOnlinePlayerbyName(buffer, p)
	case 105:
		return Bridging_Client2Server_queryPlayerbyName(buffer, p)
	case 106:
		return Bridging_Client2Server_queryBaby(buffer, p)
	case 107:
		return Bridging_Client2Server_queryEmployee(buffer, p)
	case 108:
		return Bridging_Client2Server_guideFinish(buffer, p)
	case 109:
		return Bridging_Client2Server_enterBattle(buffer, p)
	case 110:
		return Bridging_Client2Server_shopBuyItem(buffer, p)
	case 111:
		return Bridging_Client2Server_getFirstRechargeItem(buffer, p)
	case 112:
		return Bridging_Client2Server_requestLevelGift(buffer, p)
	case 113:
		return Bridging_Client2Server_setCurrentTitle(buffer, p)
	case 114:
		return Bridging_Client2Server_openBuyBox(buffer, p)
	case 115:
		return Bridging_Client2Server_requestAchaward(buffer, p)
	case 116:
		return Bridging_Client2Server_sign(buffer, p)
	case 117:
		return Bridging_Client2Server_requestSignupReward7(buffer, p)
	case 118:
		return Bridging_Client2Server_requestSignupReward14(buffer, p)
	case 119:
		return Bridging_Client2Server_requestSignupReward28(buffer, p)
	case 120:
		return Bridging_Client2Server_requestActivityReward(buffer, p)
	case 121:
		return Bridging_Client2Server_resetHundredTier(buffer, p)
	case 122:
		return Bridging_Client2Server_enterHundredScene(buffer, p)
	case 123:
		return Bridging_Client2Server_delBaby(buffer, p)
	case 124:
		return Bridging_Client2Server_resetBaby(buffer, p)
	case 125:
		return Bridging_Client2Server_resetBabyProp(buffer, p)
	case 126:
		return Bridging_Client2Server_remouldBaby(buffer, p)
	case 127:
		return Bridging_Client2Server_empSkillLevelUp(buffer, p)
	case 128:
		return Bridging_Client2Server_setOpenDoubleTimeFlag(buffer, p)
	case 129:
		return Bridging_Client2Server_talkedNpc(buffer, p)
	case 130:
		return Bridging_Client2Server_jjcBattleGo(buffer, p)
	case 131:
		return Bridging_Client2Server_requestMyJJCTeamMsg(buffer, p)
	case 132:
		return Bridging_Client2Server_startMatching(buffer, p)
	case 133:
		return Bridging_Client2Server_stopMatching(buffer, p)
	case 134:
		return Bridging_Client2Server_exitPvpJJC(buffer, p)
	case 135:
		return Bridging_Client2Server_joinPvpLobby(buffer, p)
	case 136:
		return Bridging_Client2Server_exitPvpLobby(buffer, p)
	case 137:
		return Bridging_Client2Server_requestpvprank(buffer, p)
	case 138:
		return Bridging_Client2Server_joinWarriorchoose(buffer, p)
	case 139:
		return Bridging_Client2Server_warriorStart(buffer, p)
	case 140:
		return Bridging_Client2Server_warriorStop(buffer, p)
	case 141:
		return Bridging_Client2Server_sendMail(buffer, p)
	case 142:
		return Bridging_Client2Server_readMail(buffer, p)
	case 143:
		return Bridging_Client2Server_delMail(buffer, p)
	case 144:
		return Bridging_Client2Server_getMailItem(buffer, p)
	case 145:
		return Bridging_Client2Server_requestState(buffer, p)
	case 146:
		return Bridging_Client2Server_createGuild(buffer, p)
	case 147:
		return Bridging_Client2Server_delGuild(buffer, p)
	case 148:
		return Bridging_Client2Server_requestJoinGuild(buffer, p)
	case 149:
		return Bridging_Client2Server_leaveGuild(buffer, p)
	case 150:
		return Bridging_Client2Server_kickOut(buffer, p)
	case 151:
		return Bridging_Client2Server_acceptRequestGuild(buffer, p)
	case 152:
		return Bridging_Client2Server_refuseRequestGuild(buffer, p)
	case 153:
		return Bridging_Client2Server_changeMemberPosition(buffer, p)
	case 154:
		return Bridging_Client2Server_transferPremier(buffer, p)
	case 155:
		return Bridging_Client2Server_changeGuildNotice(buffer, p)
	case 156:
		return Bridging_Client2Server_queryGuildList(buffer, p)
	case 157:
		return Bridging_Client2Server_inviteJoinGuild(buffer, p)
	case 158:
		return Bridging_Client2Server_respondInviteJoinGuild(buffer, p)
	case 159:
		return Bridging_Client2Server_buyGuildItem(buffer, p)
	case 160:
		return Bridging_Client2Server_entryGuildBattle(buffer, p)
	case 161:
		return Bridging_Client2Server_transforGuildBattleScene(buffer, p)
	case 162:
		return Bridging_Client2Server_addGuildMoney(buffer, p)
	case 163:
		return Bridging_Client2Server_updateGuildBuiling(buffer, p)
	case 164:
		return Bridging_Client2Server_refreshGuildShop(buffer, p)
	case 165:
		return Bridging_Client2Server_levelupGuildSkill(buffer, p)
	case 166:
		return Bridging_Client2Server_presentGuildItem(buffer, p)
	case 167:
		return Bridging_Client2Server_progenitusAddExp(buffer, p)
	case 168:
		return Bridging_Client2Server_setProgenitusPosition(buffer, p)
	case 169:
		return Bridging_Client2Server_guildsign(buffer, p)
	case 170:
		return Bridging_Client2Server_fetchSelling(buffer, p)
	case 171:
		return Bridging_Client2Server_fetchSelling2(buffer, p)
	case 172:
		return Bridging_Client2Server_selling(buffer, p)
	case 173:
		return Bridging_Client2Server_unselling(buffer, p)
	case 174:
		return Bridging_Client2Server_buy(buffer, p)
	case 175:
		return Bridging_Client2Server_fixItem(buffer, p)
	case 176:
		return Bridging_Client2Server_fixAllItem(buffer, p)
	case 177:
		return Bridging_Client2Server_makeDebirsItem(buffer, p)
	case 178:
		return Bridging_Client2Server_levelUpMagicItem(buffer, p)
	case 179:
		return Bridging_Client2Server_tupoMagicItem(buffer, p)
	case 180:
		return Bridging_Client2Server_changeMagicJob(buffer, p)
	case 181:
		return Bridging_Client2Server_requestPk(buffer, p)
	case 182:
		return Bridging_Client2Server_uiBehavior(buffer, p)
	case 183:
		return Bridging_Client2Server_openZhuanpan(buffer, p)
	case 184:
		return Bridging_Client2Server_zhuanpanGo(buffer, p)
	case 185:
		return Bridging_Client2Server_redemptionSpree(buffer, p)
	case 186:
		return Bridging_Client2Server_sceneFilter(buffer, p)
	case 187:
		return Bridging_Client2Server_sendExamAnswer(buffer, p)
	case 188:
		return Bridging_Client2Server_sendwishing(buffer, p)
	case 189:
		return Bridging_Client2Server_requestWish(buffer, p)
	case 190:
		return Bridging_Client2Server_leaderCloseDialog(buffer, p)
	case 191:
		return Bridging_Client2Server_requestOnlineReward(buffer, p)
	case 192:
		return Bridging_Client2Server_requestFundReward(buffer, p)
	case 193:
		return Bridging_Client2Server_openCard(buffer, p)
	case 194:
		return Bridging_Client2Server_resetCard(buffer, p)
	case 195:
		return Bridging_Client2Server_hotRoleBuy(buffer, p)
	case 196:
		return Bridging_Client2Server_requestSevenReward(buffer, p)
	case 197:
		return Bridging_Client2Server_vipreward(buffer, p)
	case 198:
		return Bridging_Client2Server_requestChargeTotalSingleReward(buffer, p)
	case 199:
		return Bridging_Client2Server_requestChargeTotalReward(buffer, p)
	case 200:
		return Bridging_Client2Server_requestChargeEverySingleReward(buffer, p)
	case 201:
		return Bridging_Client2Server_requestChargeEveryReward(buffer, p)
	case 202:
		return Bridging_Client2Server_requestLoginTotal(buffer, p)
	case 203:
		return Bridging_Client2Server_buyDiscountStoreSingle(buffer, p)
	case 204:
		return Bridging_Client2Server_buyDiscountStore(buffer, p)
	case 205:
		return Bridging_Client2Server_requestEmployeeActivityReward(buffer, p)
	case 206:
		return Bridging_Client2Server_requestmyselfrechargeleReward(buffer, p)
	case 207:
		return Bridging_Client2Server_requestEverydayIntegral(buffer, p)
	case 208:
		return Bridging_Client2Server_buyIntegralItem(buffer, p)
	case 209:
		return Bridging_Client2Server_familyLoseLeader(buffer, p)
	case 210:
		return Bridging_Client2Server_verificationSMS(buffer, p)
	case 211:
		return Bridging_Client2Server_lockItem(buffer, p)
	case 212:
		return Bridging_Client2Server_lockBaby(buffer, p)
	case 213:
		return Bridging_Client2Server_showBaby(buffer, p)
	case 214:
		return Bridging_Client2Server_wearFuwen(buffer, p)
	case 215:
		return Bridging_Client2Server_takeoffFuwen(buffer, p)
	case 216:
		return Bridging_Client2Server_compFuwen(buffer, p)
	case 217:
		return Bridging_Client2Server_requestEmployeeQuest(buffer, p)
	case 218:
		return Bridging_Client2Server_acceptEmployeeQuest(buffer, p)
	case 219:
		return Bridging_Client2Server_submitEmployeeQuest(buffer, p)
	case 220:
		return Bridging_Client2Server_crystalUpLevel(buffer, p)
	case 221:
		return Bridging_Client2Server_resetCrystalProp(buffer, p)
	case 222:
		return Bridging_Client2Server_magicItemOneKeyLevel(buffer, p)
	default:
		return errors.New(prpc.NoneDispatchMatchError)
	}
}
