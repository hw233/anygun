package prpc

// enum GameEventType
const (
	GET_None                    = 0
	GET_Online                  = 1
	GET_Offline                 = 2
	GET_CreatePlayerOK          = 3
	GET_Kill                    = 4
	GET_Die                     = 5
	GET_LevelUp                 = 6
	GET_Flee                    = 7
	GET_LearnSkill              = 8
	GET_SkillLevelUp            = 9
	GET_UseSkill                = 10
	GET_EnterJJC                = 11
	GET_MakeEquip               = 12
	GET_RecruitEmp              = 13
	GET_EmployeeEvolve          = 14
	GET_AddBaby                 = 15
	GET_DelBaby                 = 16
	GET_CatchBaby               = 17
	GET_DepositBaby             = 18
	GET_BabyLevelUp             = 19
	GET_ResetBaby               = 20
	GET_BabyNo                  = 21
	GET_BabyLearnSkill          = 22
	GET_AddSkillExp             = 23
	GET_HalfHourAgo             = 24
	GET_Sign                    = 25
	GET_BattleChangeProp        = 26
	GET_BattleOver              = 27
	GET_TalkNpc                 = 28
	GET_Activity                = 29
	GET_PvR                     = 30
	GET_JJC                     = 31
	GET_Challenge               = 32
	GET_Zhuanpan                = 33
	GET_Richang                 = 34
	GET_Pet                     = 35
	GET_Tongji                  = 36
	GET_Babyintensify           = 37
	GET_CreateGuild             = 38
	GET_JoinGuild               = 39
	GET_GuildBattleWin          = 40
	GET_GuildBattleLose         = 41
	GET_OpenGuildBattle         = 42
	GET_CloseGuildBattle        = 43
	GET_OpenGuildDemonInvaded   = 44
	GET_CloseGuildDemonInvaded  = 45
	GET_OpenGuildLeaderInvaded  = 46
	GET_CloseGuildLeaderInvaded = 47
	GET_Exam                    = 48
	GET_Wish                    = 49
	GET_ChangeMoney             = 50
	GET_ChangeDiamond           = 51
	GET_ChangeMagicCurrency     = 52
	GET_WearEquip               = 53
	GET_AddFleeProp             = 54
	GET_Gather                  = 55
	GET_AddFriend               = 56
	GET_ExtendStorage           = 57
	GET_FinishAch               = 58
	GET_AddTeamMemberCondition  = 59
	GET_Shenqishengji           = 60
	GET_NewServer               = 61
	GET_Recharge                = 62
	GET_PhoneNumber             = 63
	GET_ChangeProfession        = 64
	GET_Max                     = 65
)

func ToName_GameEventType(id int) string {
	switch id {
	case GET_None:
		return "GET_None"
	case GET_Online:
		return "GET_Online"
	case GET_Offline:
		return "GET_Offline"
	case GET_CreatePlayerOK:
		return "GET_CreatePlayerOK"
	case GET_Kill:
		return "GET_Kill"
	case GET_Die:
		return "GET_Die"
	case GET_LevelUp:
		return "GET_LevelUp"
	case GET_Flee:
		return "GET_Flee"
	case GET_LearnSkill:
		return "GET_LearnSkill"
	case GET_SkillLevelUp:
		return "GET_SkillLevelUp"
	case GET_UseSkill:
		return "GET_UseSkill"
	case GET_EnterJJC:
		return "GET_EnterJJC"
	case GET_MakeEquip:
		return "GET_MakeEquip"
	case GET_RecruitEmp:
		return "GET_RecruitEmp"
	case GET_EmployeeEvolve:
		return "GET_EmployeeEvolve"
	case GET_AddBaby:
		return "GET_AddBaby"
	case GET_DelBaby:
		return "GET_DelBaby"
	case GET_CatchBaby:
		return "GET_CatchBaby"
	case GET_DepositBaby:
		return "GET_DepositBaby"
	case GET_BabyLevelUp:
		return "GET_BabyLevelUp"
	case GET_ResetBaby:
		return "GET_ResetBaby"
	case GET_BabyNo:
		return "GET_BabyNo"
	case GET_BabyLearnSkill:
		return "GET_BabyLearnSkill"
	case GET_AddSkillExp:
		return "GET_AddSkillExp"
	case GET_HalfHourAgo:
		return "GET_HalfHourAgo"
	case GET_Sign:
		return "GET_Sign"
	case GET_BattleChangeProp:
		return "GET_BattleChangeProp"
	case GET_BattleOver:
		return "GET_BattleOver"
	case GET_TalkNpc:
		return "GET_TalkNpc"
	case GET_Activity:
		return "GET_Activity"
	case GET_PvR:
		return "GET_PvR"
	case GET_JJC:
		return "GET_JJC"
	case GET_Challenge:
		return "GET_Challenge"
	case GET_Zhuanpan:
		return "GET_Zhuanpan"
	case GET_Richang:
		return "GET_Richang"
	case GET_Pet:
		return "GET_Pet"
	case GET_Tongji:
		return "GET_Tongji"
	case GET_Babyintensify:
		return "GET_Babyintensify"
	case GET_CreateGuild:
		return "GET_CreateGuild"
	case GET_JoinGuild:
		return "GET_JoinGuild"
	case GET_GuildBattleWin:
		return "GET_GuildBattleWin"
	case GET_GuildBattleLose:
		return "GET_GuildBattleLose"
	case GET_OpenGuildBattle:
		return "GET_OpenGuildBattle"
	case GET_CloseGuildBattle:
		return "GET_CloseGuildBattle"
	case GET_OpenGuildDemonInvaded:
		return "GET_OpenGuildDemonInvaded"
	case GET_CloseGuildDemonInvaded:
		return "GET_CloseGuildDemonInvaded"
	case GET_OpenGuildLeaderInvaded:
		return "GET_OpenGuildLeaderInvaded"
	case GET_CloseGuildLeaderInvaded:
		return "GET_CloseGuildLeaderInvaded"
	case GET_Exam:
		return "GET_Exam"
	case GET_Wish:
		return "GET_Wish"
	case GET_ChangeMoney:
		return "GET_ChangeMoney"
	case GET_ChangeDiamond:
		return "GET_ChangeDiamond"
	case GET_ChangeMagicCurrency:
		return "GET_ChangeMagicCurrency"
	case GET_WearEquip:
		return "GET_WearEquip"
	case GET_AddFleeProp:
		return "GET_AddFleeProp"
	case GET_Gather:
		return "GET_Gather"
	case GET_AddFriend:
		return "GET_AddFriend"
	case GET_ExtendStorage:
		return "GET_ExtendStorage"
	case GET_FinishAch:
		return "GET_FinishAch"
	case GET_AddTeamMemberCondition:
		return "GET_AddTeamMemberCondition"
	case GET_Shenqishengji:
		return "GET_Shenqishengji"
	case GET_NewServer:
		return "GET_NewServer"
	case GET_Recharge:
		return "GET_Recharge"
	case GET_PhoneNumber:
		return "GET_PhoneNumber"
	case GET_ChangeProfession:
		return "GET_ChangeProfession"
	case GET_Max:
		return "GET_Max"
	default:
		return ""
	}
}
func ToId_GameEventType(name string) int {
	switch name {
	case "GET_None":
		return GET_None
	case "GET_Online":
		return GET_Online
	case "GET_Offline":
		return GET_Offline
	case "GET_CreatePlayerOK":
		return GET_CreatePlayerOK
	case "GET_Kill":
		return GET_Kill
	case "GET_Die":
		return GET_Die
	case "GET_LevelUp":
		return GET_LevelUp
	case "GET_Flee":
		return GET_Flee
	case "GET_LearnSkill":
		return GET_LearnSkill
	case "GET_SkillLevelUp":
		return GET_SkillLevelUp
	case "GET_UseSkill":
		return GET_UseSkill
	case "GET_EnterJJC":
		return GET_EnterJJC
	case "GET_MakeEquip":
		return GET_MakeEquip
	case "GET_RecruitEmp":
		return GET_RecruitEmp
	case "GET_EmployeeEvolve":
		return GET_EmployeeEvolve
	case "GET_AddBaby":
		return GET_AddBaby
	case "GET_DelBaby":
		return GET_DelBaby
	case "GET_CatchBaby":
		return GET_CatchBaby
	case "GET_DepositBaby":
		return GET_DepositBaby
	case "GET_BabyLevelUp":
		return GET_BabyLevelUp
	case "GET_ResetBaby":
		return GET_ResetBaby
	case "GET_BabyNo":
		return GET_BabyNo
	case "GET_BabyLearnSkill":
		return GET_BabyLearnSkill
	case "GET_AddSkillExp":
		return GET_AddSkillExp
	case "GET_HalfHourAgo":
		return GET_HalfHourAgo
	case "GET_Sign":
		return GET_Sign
	case "GET_BattleChangeProp":
		return GET_BattleChangeProp
	case "GET_BattleOver":
		return GET_BattleOver
	case "GET_TalkNpc":
		return GET_TalkNpc
	case "GET_Activity":
		return GET_Activity
	case "GET_PvR":
		return GET_PvR
	case "GET_JJC":
		return GET_JJC
	case "GET_Challenge":
		return GET_Challenge
	case "GET_Zhuanpan":
		return GET_Zhuanpan
	case "GET_Richang":
		return GET_Richang
	case "GET_Pet":
		return GET_Pet
	case "GET_Tongji":
		return GET_Tongji
	case "GET_Babyintensify":
		return GET_Babyintensify
	case "GET_CreateGuild":
		return GET_CreateGuild
	case "GET_JoinGuild":
		return GET_JoinGuild
	case "GET_GuildBattleWin":
		return GET_GuildBattleWin
	case "GET_GuildBattleLose":
		return GET_GuildBattleLose
	case "GET_OpenGuildBattle":
		return GET_OpenGuildBattle
	case "GET_CloseGuildBattle":
		return GET_CloseGuildBattle
	case "GET_OpenGuildDemonInvaded":
		return GET_OpenGuildDemonInvaded
	case "GET_CloseGuildDemonInvaded":
		return GET_CloseGuildDemonInvaded
	case "GET_OpenGuildLeaderInvaded":
		return GET_OpenGuildLeaderInvaded
	case "GET_CloseGuildLeaderInvaded":
		return GET_CloseGuildLeaderInvaded
	case "GET_Exam":
		return GET_Exam
	case "GET_Wish":
		return GET_Wish
	case "GET_ChangeMoney":
		return GET_ChangeMoney
	case "GET_ChangeDiamond":
		return GET_ChangeDiamond
	case "GET_ChangeMagicCurrency":
		return GET_ChangeMagicCurrency
	case "GET_WearEquip":
		return GET_WearEquip
	case "GET_AddFleeProp":
		return GET_AddFleeProp
	case "GET_Gather":
		return GET_Gather
	case "GET_AddFriend":
		return GET_AddFriend
	case "GET_ExtendStorage":
		return GET_ExtendStorage
	case "GET_FinishAch":
		return GET_FinishAch
	case "GET_AddTeamMemberCondition":
		return GET_AddTeamMemberCondition
	case "GET_Shenqishengji":
		return GET_Shenqishengji
	case "GET_NewServer":
		return GET_NewServer
	case "GET_Recharge":
		return GET_Recharge
	case "GET_PhoneNumber":
		return GET_PhoneNumber
	case "GET_ChangeProfession":
		return GET_ChangeProfession
	case "GET_Max":
		return GET_Max
	default:
		return -1
	}
}
