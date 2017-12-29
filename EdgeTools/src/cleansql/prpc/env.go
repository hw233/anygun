package prpc

import (
	"bytes"
	"database/sql"
	"fmt"
	"time"
)

//=============================================================

func OpenDB() *sql.DB {
	dsn := ""
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil
	}
	return db
}

func (this *COM_Item) GetId() int {
	return int(this.itemId_)
}

func (this *COM_Item) GetStack() int {
	return int(this.stack_)
}

func (this *COM_EmployeeInst) GetColor() int {
	return this.quality_
}

func (this *SGE_DBPlayerData) GetLogoutTime() uint64 {
	return this.logoutTime_
}

func (this *SGE_DBPlayerData) GetPlayerName() string {
	return this.instName_
}

func (this *SGE_DBPlayerData) GetPlayerLevel() int {
	return int(this.properties_[PT_Level])
}

func (this *SGE_DBPlayerData) GetM() int {
	return int(this.properties_[PT_MagicCurrency])
}

func (this *SGE_DBPlayerData) GetD() int {
	return int(this.properties_[PT_Diamond])
}

func (this *SGE_DBPlayerData) GetPlayerInstId() int {
	return int(this.instId_)
}

func (this *SGE_DBPlayerData) GetEmployeeItem() int {
	stack := 0
	for _, v := range this.bagItems_ {
		if v.GetId() == 4511 {
			stack += v.GetStack()
		}
	}
	return stack
}

func (this *SGE_DBPlayerData) GetGoldenEmployeeCount() int {
	count := 0
	for _, v := range this.employees_ {
		if v.GetColor() >= QC_Golden {
			count++
		}
	}
	return count
}

func TestASA() {
	b := bytes.NewBuffer(nil)
	inst := COM_PlayerInst{}
	inst.blueBoxTimes_ = float32(time.Now().Unix())
	inst.instName_ = "123123123124ssadhana2u4wjkhqw3jkhrjkfh23反对减肥咖啡卡拉发开发哈数据库发哈数据来看返回啊时间浪费和啊数据库返回金克拉速度放缓金克拉是否会尽快拉萨的符号啊数据库路哈及刻录"
	err := inst.Serialize(b)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(b.Len())

	b2 := bytes.NewBuffer(b.Bytes())
	fmt.Println(b2.Len())

	inst2 := COM_PlayerInst{}

	err = inst2.Deserialize(b2)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(inst2.blueBoxTimes_)
	fmt.Println(inst2.instName_)
}

func TestDB() {

	db := OpenDB()

	if db == nil {
		return
	}

	defer db.Close()

	records, err := db.Query("SELECT `BinData` FROM `Player`")
	if err != nil {
		fmt.Println(err)
		return
	}
	count := 0
	for records.Next() {
		var bs []byte
		err := records.Scan(&bs)
		if err != nil {
			fmt.Println(err)
			continue
		}
		b := bytes.NewBuffer(bs)
		inst := SGE_DBPlayerData{}
		err = inst.Deserialize(b)
		if err != nil {
			fmt.Println(err)
			continue
		}

		count++
		tm := time.Unix(int64(inst.GetLogoutTime()), 0)
		//fmt.Println(tm, inst.GetPlayerLevel()," ",inst.GetEmployeeItem(),inst.GetD(), inst.GetM(), inst.GetGoldenEmployeeCount())
		if inst.GetEmployeeItem() > 0 || inst.GetM() > 0 || inst.GetD() > 10000 || inst.properties_[PT_Money] > 100000 {
			fmt.Println(tm, inst.GetPlayerName(), inst.GetPlayerLevel(), inst.GetM(), inst.GetD(), inst.properties_[PT_Money], inst.GetEmployeeItem())
		}
		//fmt.Println(inst)
	}

}

type ClientSender struct {
	Buffered *bytes.Buffer
}

func (this *ClientSender) MethodBegin() *bytes.Buffer {
	return this.Buffered
}

func (this *ClientSender) MethodEnd() error {
	return nil
}

type ClientS struct {
	Client2ServerStub
}

type Client struct {
}

func (this *Client) openvip(vl int) error {
	return nil
} // 0
func (this *Client) requestPhoto() error {
	return nil
} // 1
func (this *Client) ping() error {
	return nil
} // 2
func (this *Client) sessionlogin(info COM_LoginInfo) error {
	return nil
} // 3
func (this *Client) login(info COM_LoginInfo) error {
	return nil
} // 4
func (this *Client) createPlayer(playername string, playerTmpId uint8) error {
	return nil
} // 5
func (this *Client) deletePlayer(playername string) error {
	return nil
} // 6
func (this *Client) enterGame(playerInstId uint32) error {
	return nil
} // 7
func (this *Client) requestBag() error {
	return nil
} // 8
func (this *Client) requestEmployees() error {
	return nil
} // 9
func (this *Client) requestStorage(tp int) error {
	return nil
} // 10
func (this *Client) requestAchievement() error {
	return nil
} // 11
func (this *Client) initminig() error {
	return nil
} // 12
func (this *Client) requestCompound() error {
	return nil
} // 13
func (this *Client) move(x float32, z float32) error {
	return nil
} // 14
func (this *Client) moveToNpc(npcid int32) error {
	return nil
} // 15
func (this *Client) moveToNpc2(type_ int) error {
	return nil
} // 16
func (this *Client) moveToZone(sceneId int32, zoneId int32) error {
	return nil
} // 17
func (this *Client) autoBattle() error {
	return nil
} // 18
func (this *Client) stopAutoBattle() error {
	return nil
} // 19
func (this *Client) stopMove() error {
	return nil
} // 20
func (this *Client) exitCopy() error {
	return nil
} // 21
func (this *Client) transforScene(sceneId int32) error {
	return nil
} // 22
func (this *Client) sceneLoaded() error {
	return nil
} // 23
func (this *Client) querySimplePlayerInst(instId uint32) error {
	return nil
} // 24
func (this *Client) logout() error {
	return nil
} // 25
func (this *Client) changProp(guid uint32, props []COM_Addprop) error {
	return nil
} // 26
func (this *Client) learnSkill(skid uint32) error {
	return nil
} // 27
func (this *Client) babyLearnSkill(instId uint32, oldSkId uint32, newSkId uint32, newSkLv uint32) error {
	return nil
} // 28
func (this *Client) forgetSkill(skid uint32) error {
	return nil
} // 29
func (this *Client) syncOrder(order COM_Order) error {
	return nil
} // 30
func (this *Client) syncOrderTimeout() error {
	return nil
} // 31
func (this *Client) sendChat(content COM_Chat, targetName string) error {
	return nil
} // 32
func (this *Client) requestAudio(audioId int32) error {
	return nil
} // 33
func (this *Client) publishItemInst(type_ int, itemInstId uint32, chatType int, playerName string) error {
	return nil
} // 34
func (this *Client) queryItemInst(showId int32) error {
	return nil
} // 35
func (this *Client) publishbabyInst(type_ int, babyInstId uint32, playerName string) error {
	return nil
} // 36
func (this *Client) querybabyInst(showId int32) error {
	return nil
} // 37
func (this *Client) useItem(slot uint32, target uint32, stack uint32) error {
	return nil
} // 38
func (this *Client) wearEquipment(target uint32, itemInstId uint32) error {
	return nil
} // 39
func (this *Client) delEquipment(target uint32, itemInstId uint32) error {
	return nil
} // 40
func (this *Client) setPlayerFront(isFront bool) error {
	return nil
} // 41
func (this *Client) setBattlebaby(babyID uint32, isBattle bool) error {
	return nil
} // 42
func (this *Client) changeBabyName(babyID uint32, name string) error {
	return nil
} // 43
func (this *Client) intensifyBaby(babyid uint32) error {
	return nil
} // 44
func (this *Client) setBattleEmp(empID uint32, group int, isBattle bool) error {
	return nil
} // 45
func (this *Client) changeEmpBattleGroup(group int) error {
	return nil
} // 46
func (this *Client) requestEvolve(empInstId uint32) error {
	return nil
} // 47
func (this *Client) requestUpStar(empInstId uint32) error {
	return nil
} // 48
func (this *Client) requestDelEmp(empInstId uint32) error {
	return nil
} // 49
func (this *Client) delEmployee(emps []uint32) error {
	return nil
} // 50
func (this *Client) onekeyDelEmp() error {
	return nil
} // 51
func (this *Client) delEmployeeSoul(instid uint32, soulNum uint32) error {
	return nil
} // 52
func (this *Client) sortBagItem() error {
	return nil
} // 53
func (this *Client) sellBagItem(instId uint32, stack uint32) error {
	return nil
} // 54
func (this *Client) depositItemToStorage(instid uint32) error {
	return nil
} // 55
func (this *Client) depositBabyToStorage(instid uint32) error {
	return nil
} // 56
func (this *Client) storageItemToBag(instid uint32) error {
	return nil
} // 57
func (this *Client) storageBabyToPlayer(instid uint32) error {
	return nil
} // 58
func (this *Client) sortStorage(tp int) error {
	return nil
} // 59
func (this *Client) delStorageBaby(instid uint32) error {
	return nil
} // 60
func (this *Client) jointLobby() error {
	return nil
} // 61
func (this *Client) exitLobby() error {
	return nil
} // 62
func (this *Client) createTeam(cti COM_CreateTeamInfo) error {
	return nil
} // 63
func (this *Client) changeTeam(info COM_CreateTeamInfo) error {
	return nil
} // 64
func (this *Client) kickTeamMember(uuid uint32) error {
	return nil
} // 65
func (this *Client) changeTeamLeader(uuid uint32) error {
	return nil
} // 66
func (this *Client) joinTeam(teamId uint32, pwd string) error {
	return nil
} // 67
func (this *Client) exitTeam() error {
	return nil
} // 68
func (this *Client) changeTeamPassword(pwd string) error {
	return nil
} // 69
func (this *Client) joinTeamRoom() error {
	return nil
} // 70
func (this *Client) inviteTeamMember(name string) error {
	return nil
} // 71
func (this *Client) isjoinTeam(teamId uint32, isFlag bool) error {
	return nil
} // 72
func (this *Client) leaveTeam() error {
	return nil
} // 73
func (this *Client) backTeam() error {
	return nil
} // 74
func (this *Client) refuseBackTeam() error {
	return nil
} // 75
func (this *Client) teamCallMember(playerId int32) error {
	return nil
} // 76
func (this *Client) requestJoinTeam(targetName string) error {
	return nil
} // 77
func (this *Client) ratifyJoinTeam(sendName string) error {
	return nil
} // 78
func (this *Client) drawLotteryBox(type_ int, isFree bool) error {
	return nil
} // 79
func (this *Client) acceptQuest(questId int32) error {
	return nil
} // 80
func (this *Client) submitQuest(npcId int32, questId int32) error {
	return nil
} // 81
func (this *Client) giveupQuest(questId int32) error {
	return nil
} // 82
func (this *Client) requestContactInfoById(instId uint32) error {
	return nil
} // 83
func (this *Client) requestContactInfoByName(instName string) error {
	return nil
} // 84
func (this *Client) requestFriendList() error {
	return nil
} // 85
func (this *Client) addFriend(instId uint32) error {
	return nil
} // 86
func (this *Client) delFriend(instId uint32) error {
	return nil
} // 87
func (this *Client) addBlacklist(instId uint32) error {
	return nil
} // 88
func (this *Client) delBlacklist(instId uint32) error {
	return nil
} // 89
func (this *Client) requestReferrFriend() error {
	return nil
} // 90
func (this *Client) mining(gatherId int32, times int32) error {
	return nil
} // 91
func (this *Client) compoundItem(itemId int32, gemId int32) error {
	return nil
} // 92
func (this *Client) bagItemSplit(instId int32, splitNum int32) error {
	return nil
} // 93
func (this *Client) requestChallenge(name string) error {
	return nil
} // 94
func (this *Client) requestRival() error {
	return nil
} // 95
func (this *Client) requestMySelfJJCData() error {
	return nil
} // 96
func (this *Client) requestCheckMsg(name string) error {
	return nil
} // 97
func (this *Client) requestMyAllbattleMsg() error {
	return nil
} // 98
func (this *Client) requestJJCRank() error {
	return nil
} // 99
func (this *Client) requestLevelRank() error {
	return nil
} // 100
func (this *Client) requestBabyRank() error {
	return nil
} // 101
func (this *Client) requestEmpRank() error {
	return nil
} // 102
func (this *Client) requestPlayerFFRank() error {
	return nil
} // 103
func (this *Client) queryOnlinePlayerbyName(name string) error {
	return nil
} // 104
func (this *Client) queryPlayerbyName(name string) error {
	return nil
} // 105
func (this *Client) queryBaby(instId uint32) error {
	return nil
} // 106
func (this *Client) queryEmployee(instId uint32) error {
	return nil
} // 107
func (this *Client) guideFinish(guideIdx uint64) error {
	return nil
} // 108
func (this *Client) enterBattle(battleId int32) error {
	return nil
} // 109
func (this *Client) shopBuyItem(id int32, num int32) error {
	return nil
} // 110
func (this *Client) getFirstRechargeItem() error {
	return nil
} // 111
func (this *Client) requestLevelGift(level int32) error {
	return nil
} // 112
func (this *Client) setCurrentTitle(title int32) error {
	return nil
} // 113
func (this *Client) openBuyBox() error {
	return nil
} // 114
func (this *Client) requestAchaward(achId int32) error {
	return nil
} // 115
func (this *Client) sign(index int32) error {
	fmt.Println(index)
	return nil
} // 116
func (this *Client) requestSignupReward7() error {
	return nil
} // 117
func (this *Client) requestSignupReward14() error {
	return nil
} // 118
func (this *Client) requestSignupReward28() error {
	return nil
} // 119
func (this *Client) requestActivityReward(index int32) error {
	return nil
} // 120
func (this *Client) resetHundredTier() error {
	return nil
} // 121
func (this *Client) enterHundredScene(level int32) error {
	return nil
} // 122
func (this *Client) delBaby(instId int32) error {
	return nil
} // 123
func (this *Client) resetBaby(instId int32) error {
	return nil
} // 124
func (this *Client) resetBabyProp(instId int32) error {
	return nil
} // 125
func (this *Client) remouldBaby(instid int32) error {
	return nil
} // 126
func (this *Client) empSkillLevelUp(empId uint32, skillId int32) error {
	return nil
} // 127
func (this *Client) setOpenDoubleTimeFlag(isFlag bool) error {
	return nil
} // 128
func (this *Client) talkedNpc(npcId int32) error {
	return nil
} // 129
func (this *Client) jjcBattleGo(id uint32) error {
	return nil
} // 130
func (this *Client) requestMyJJCTeamMsg() error {
	return nil
} // 131
func (this *Client) startMatching() error {
	return nil
} // 132
func (this *Client) stopMatching() error {
	return nil
} // 133
func (this *Client) exitPvpJJC() error {
	return nil
} // 134
func (this *Client) joinPvpLobby() error {
	return nil
} // 135
func (this *Client) exitPvpLobby() error {
	return nil
} // 136
func (this *Client) requestpvprank() error {
	return nil
} // 137
func (this *Client) joinWarriorchoose() error {
	return nil
} // 138
func (this *Client) warriorStart() error {
	return nil
} // 139
func (this *Client) warriorStop() error {
	return nil
} // 140
func (this *Client) sendMail(playername string, title string, content string) error {
	return nil
} // 141
func (this *Client) readMail(mailId int32) error {
	return nil
} // 142
func (this *Client) delMail(mailId int32) error {
	return nil
} // 143
func (this *Client) getMailItem(mailId int32) error {
	return nil
} // 144
func (this *Client) requestState() error {
	return nil
} // 145
func (this *Client) createGuild(guildName string) error {
	return nil
} // 146
func (this *Client) delGuild(guildId uint32) error {
	return nil
} // 147
func (this *Client) requestJoinGuild(guid uint32) error {
	return nil
} // 148
func (this *Client) leaveGuild() error {
	return nil
} // 149
func (this *Client) kickOut(guid int32) error {
	return nil
} // 150
func (this *Client) acceptRequestGuild(playerId int32) error {
	return nil
} // 151
func (this *Client) refuseRequestGuild(playerId int32) error {
	return nil
} // 152
func (this *Client) changeMemberPosition(targetId int32, job int) error {
	return nil
} // 153
func (this *Client) transferPremier(targetId int32) error {
	return nil
} // 154
func (this *Client) changeGuildNotice(notice string) error {
	return nil
} // 155
func (this *Client) queryGuildList(page int16) error {
	return nil
} // 156
func (this *Client) inviteJoinGuild(playerName string) error {
	return nil
} // 157
func (this *Client) respondInviteJoinGuild(sendName string) error {
	return nil
} // 158
func (this *Client) buyGuildItem(tableId int32, times int32) error {
	return nil
} // 159
func (this *Client) entryGuildBattle() error {
	return nil
} // 160
func (this *Client) transforGuildBattleScene() error {
	return nil
} // 161
func (this *Client) addGuildMoney(money int32) error {
	return nil
} // 162
func (this *Client) updateGuildBuiling(gbt int) error {
	return nil
} // 163
func (this *Client) refreshGuildShop() error {
	return nil
} // 164
func (this *Client) levelupGuildSkill(skId int32) error {
	return nil
} // 165
func (this *Client) presentGuildItem(num int32) error {
	return nil
} // 166
func (this *Client) progenitusAddExp(monsterId int32, isSuper bool) error {
	return nil
} // 167
func (this *Client) setProgenitusPosition(mId int32, pos int32) error {
	return nil
} // 168
func (this *Client) guildsign() error {
	return nil
} // 169
func (this *Client) fetchSelling(context COM_SearchContext) error {
	return nil
} // 170
func (this *Client) fetchSelling2(context COM_SearchContext) error {
	return nil
} // 171
func (this *Client) selling(iteminstid int32, babyinstid int32, price int32) error {
	return nil
} // 172
func (this *Client) unselling(sellid int32) error {
	return nil
} // 173
func (this *Client) buy(sellid int32) error {
	return nil
} // 174
func (this *Client) fixItem(instId int32, type_ int) error {
	return nil
} // 175
func (this *Client) fixAllItem(items []uint32, type_ int) error {
	return nil
} // 176
func (this *Client) makeDebirsItem(instId int32, num int32) error {
	return nil
} // 177
func (this *Client) levelUpMagicItem(items []uint32) error {
	return nil
} // 178
func (this *Client) tupoMagicItem(level int32) error {
	return nil
} // 179
func (this *Client) changeMagicJob(job int) error {
	return nil
} // 180
func (this *Client) requestPk(playerId uint32) error {
	return nil
} // 181
func (this *Client) uiBehavior(type_ int) error {
	return nil
} // 182
func (this *Client) openZhuanpan() error {
	return nil
} // 183
func (this *Client) zhuanpanGo(counter uint32) error {
	return nil
} // 184
func (this *Client) redemptionSpree(code string) error {
	return nil
} // 185
func (this *Client) sceneFilter(sfType []int) error {
	return nil
} // 186
func (this *Client) sendExamAnswer(questionId uint32, answer uint8) error {
	return nil
} // 187
func (this *Client) sendwishing(wish COM_Wishing) error {
	return nil
} // 188
func (this *Client) requestWish() error {
	return nil
} // 189
func (this *Client) leaderCloseDialog() error {
	return nil
} // 190
func (this *Client) requestOnlineReward(index uint32) error {
	return nil
} // 191
func (this *Client) requestFundReward(level uint32) error {
	return nil
} // 192
func (this *Client) openCard(index uint16) error {
	return nil
} // 193
func (this *Client) resetCard() error {
	return nil
} // 194
func (this *Client) hotRoleBuy() error {
	return nil
} // 195
func (this *Client) requestSevenReward(qid uint32) error {
	return nil
} // 196
func (this *Client) vipreward() error {
	return nil
} // 197
func (this *Client) requestChargeTotalSingleReward(index uint32) error {
	return nil
} // 198
func (this *Client) requestChargeTotalReward(index uint32) error {
	return nil
} // 199
func (this *Client) requestChargeEverySingleReward(index uint32) error {
	return nil
} // 200
func (this *Client) requestChargeEveryReward(index uint32) error {
	return nil
} // 201
func (this *Client) requestLoginTotal(index uint32) error {
	return nil
} // 202
func (this *Client) buyDiscountStoreSingle(itemId int32, itemStack int32) error {
	return nil
} // 203
func (this *Client) buyDiscountStore(itemId int32, itemStack int32) error {
	return nil
} // 204
func (this *Client) requestEmployeeActivityReward(index uint32) error {
	return nil
} // 205
func (this *Client) requestmyselfrechargeleReward(index uint32) error {
	return nil
} // 206
func (this *Client) requestEverydayIntegral() error {
	return nil
} // 207
func (this *Client) buyIntegralItem(id uint32, num uint32) error {
	return nil
} // 208
func (this *Client) familyLoseLeader() error {
	return nil
} // 209
func (this *Client) verificationSMS(phoneNumber string, code string) error {
	return nil
} // 210
func (this *Client) lockItem(instId int32, isLock bool) error {
	return nil
} // 211
func (this *Client) lockBaby(instId int32, isLock bool) error {
	return nil
} // 212
func (this *Client) showBaby(instId int32) error {
	return nil
} // 213
func (this *Client) wearFuwen(itemInstId int32) error {
	return nil
} // 214
func (this *Client) takeoffFuwen(slotId int32) error {
	return nil
} // 215
func (this *Client) compFuwen(itemInstId int32) error {
	return nil
} // 216
func (this *Client) requestEmployeeQuest() error {
	return nil
} // 217
func (this *Client) acceptEmployeeQuest(questId int32, employees []int32) error {
	fmt.Println(questId, employees)
	return nil
} // 218
func (this *Client) submitEmployeeQuest(questId int32) error {
	return nil
} // 219
func (this *Client) crystalUpLevel() error {
	return nil
} // 220
func (this *Client) resetCrystalProp(locklist []int32) error {
	return nil
} // 221
func (this *Client) magicItemOneKeyLevel() error {
	return nil
} // 222

func TestProxyABIB() {
	buffered := bytes.NewBuffer(nil)

	sender := ClientSender{buffered}

	cliS := &ClientS{}
	cliS.Sender = &sender

	err := cliS.sign(91)
	if err != nil {
		fmt.Println(err)
	}
	cliS.acceptEmployeeQuest(1, []int32{9, 2, 1})

	fmt.Println(buffered.Len())

	cli := Client{}
	//buffered.Truncate(0)
	fmt.Println(buffered.Len())
	Client2ServerDispatch(buffered, &cli)
	Client2ServerDispatch(buffered, &cli)
}
