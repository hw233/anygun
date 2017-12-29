package prpc

// enum ErrorNo
const (
	EN_None                                                       = 0
	EN_VersionNotMatch                                            = 1
	EN_AccountNameSame                                            = 2
	EN_PlayerNameSame                                             = 3
	EN_FilterWord                                                 = 4
	EN_CannotfindPlayer                                           = 5
	EN_AcceptQuestNotFound                                        = 6
	EN_AcceptQuestNoItem                                          = 7
	EN_AcceptSecendDaily                                          = 8
	EN_DailyNoNum                                                 = 9
	EN_AcceptSecendProfession                                     = 10
	EN_Battle                                                     = 11
	EN_MoneyLess                                                  = 12
	EN_DiamondLess                                                = 13
	EN_NoSubSyste                                                 = 14
	EN_InTeam                                                     = 15
	EN_NoTeamLeader                                               = 16
	EN_TeamPassword                                               = 17
	EN_TeamIsFull                                                 = 18
	EN_NoTeam                                                     = 19
	EN_TeamIsRunning                                              = 20
	EN_TeamMemberLeaving                                          = 21
	EN_NoBackTeam                                                 = 22
	EN_InTeamBlackList                                            = 23
	EN_EmployeeIsFull                                             = 24
	EN_NoUpSkill                                                  = 25
	EN_PropisNull                                                 = 26
	EN_DoubleExpTimeFull                                          = 27
	EN_DoubleExpTimeNULL                                          = 28
	EN_NoTeamNoTongji                                             = 29
	EN_TongjiTimesMax                                             = 30
	EN_TongjiTeamMemberTimesMax                                   = 31
	EN_NoTeamLeaderNoTongji                                       = 32
	EN_TeamSizeTongjiError                                        = 33
	EN_GetPoisonMushroom                                          = 34
	EN_GetMushroom                                                = 35
	EN_TongjiTeamLevelTooLow                                      = 36
	EN_PlayerIsInTeam                                             = 37
	EN_AcceptQuestBagMax                                          = 38
	EN_SubmitQuestBagMax                                          = 39
	EN_GuildNameSame                                              = 40
	EN_PlayerGoldLess                                             = 41
	EN_PlayerHasGuild                                             = 42
	EN_InRequestErr                                               = 43
	EN_RequestListFull                                            = 44
	EN_joinGuildRequestOk                                         = 45
	EN_JoinOtherGuild                                             = 46
	EN_PremierQuitError                                           = 47
	EN_CommandPositionLess                                        = 48
	EN_PositionUpMax                                              = 49
	EN_MallBuyOk                                                  = 50
	EN_MallBuyFailBagFull                                         = 51
	EN_MallBuyFailBabyFull                                        = 52
	EN_MallBuyFailDiamondLess                                     = 53
	EN_MallBuyFailSelled                                          = 54
	EN_OpenBaoXiangBagFull                                        = 55
	EN_NoBaby                                                     = 56
	EN_BagFull                                                    = 57
	EN_BagSizeMax                                                 = 58
	EN_BabyStorageFull                                            = 59
	EN_BabyFullToStorage                                          = 60
	EN_NewItemError                                               = 61
	EN_BabyFull                                                   = 62
	EN_RemouldBabyLevel                                           = 63
	EN_SkillSoltFull                                              = 64
	EN_WorldChatPayError                                          = 65
	EN_DontTalk                                                   = 66
	EN_BadMushroom                                                = 67
	EN_ItemMushroom                                               = 68
	EN_GetMailItemBagFull                                         = 69
	EN_Materialless                                               = 70
	EN_OpenGatherlose                                             = 71
	EN_OpenGatherRepetition                                       = 72
	EN_GatherLevelLess                                            = 73
	EN_GatherTimesLess                                            = 74
	EN_OpenBaoXiangLevel                                          = 75
	EN_NoBattleBaby                                               = 76
	EN_NoThisPoint                                                = 77
	EN_BabyLevelHigh                                              = 78
	EN_AddMoney1W                                                 = 79
	EN_AddDionmand100                                             = 80
	EN_AddMoney2W                                                 = 81
	EN_AddDionmand200                                             = 82
	EN_AddMoney3W                                                 = 83
	EN_AddDionmand300                                             = 84
	EN_AddMoney4W                                                 = 85
	EN_AddDionmand400                                             = 86
	EN_AddMoney5W                                                 = 87
	EN_AddDionmand500                                             = 88
	EN_AddMoney6W                                                 = 89
	EN_AddDionmand600                                             = 90
	EN_DelBaby1000                                                = 91
	EN_DelBaby30                                                  = 92
	EN_DelBaby54                                                  = 93
	EN_DelBaby10030                                               = 94
	EN_DelBaby10015                                               = 95
	EN_DelBaby10157                                               = 96
	EN_DelDefaultSkill                                            = 97
	EN_GreenBoxFreeNum                                            = 98
	EM_NotNormalVip                                               = 99
	EN_NotSuperVip                                                = 100
	EN_FirendNotOpen                                              = 101
	EN_BlackCannotFriend                                          = 102
	EN_HasFriend                                                  = 103
	EN_HunderdNoNum                                               = 104
	EN_HunderdTier                                                = 105
	EN_HunderdLevel                                               = 106
	EN_SkillExperr                                                = 107
	EN_TimeMushroom                                               = 108
	EN_TimeTongji                                                 = 109
	EN_TimeXiji                                                   = 110
	EN_Storefull                                                  = 111
	EN_HasBattleTime                                              = 112
	EN_HourLess24                                                 = 113
	EN_HourLess24_Join                                            = 114
	EN_TeamMemberHourLess24                                       = 115
	EN_GuildBattleJoinSceneMoveValue                              = 116
	EN_IdgenNull                                                  = 117
	EN_Idgenhas                                                   = 118
	EN_Gifthas                                                    = 119
	EN_UseMakeRepeat                                              = 120
	EN_WishNull                                                   = 121
	EN_NoGuild                                                    = 122
	EN_GuildMemberMax                                             = 123
	EN_LevelupGuildBuilding                                       = 124
	EN_LevelupGuildBuildingLevelMax                               = 125
	EN_LevelupGuildBuildingMoneyLess                              = 126
	EN_LevelupGuildBuildingHallBuildLevelLess                     = 127
	EN_AddGuildMoneyMax                                           = 128
	EN_PresentGuildOk                                             = 129
	EN_RefreshShopTimeLess                                        = 130
	EN_TeamMemberNoGuild                                          = 131
	EN_MagicCurrencyLess                                          = 132
	EN_DisShopLimitLess                                           = 133
	EN_FamilyPremierCanntDelete                                   = 134
	EN_MyFamilyMonster                                            = 135
	EN_NoBattleTime                                               = 136
	EN_OtherNoBattleTime                                          = 137
	EN_GuildBattleTimeout2                                        = 138
	EN_GuildBattleHasTeam                                         = 139
	EN_AccountIsSeal                                              = 140
	EN_PlayerNoOnline                                             = 141
	EN_ActivityNoTime                                             = 142
	EN_NoTeamExist                                                = 143
	EN_GuildNoMember                                              = 144
	EN_MallSellTimeLess                                           = 145
	EN_GuildMemberLess24                                          = 146
	EN_InviteeLeaveGuildLess24                                    = 147
	EN_PboneNumberSuccess                                         = 148
	EN_PhoneNumberError                                           = 149
	EN_OtherPlayerInBattle                                        = 150
	EN_GuildBattleNotMatch                                        = 151
	EN_GuildBattleIsStart                                         = 152
	EN_GuileBattleIsClose                                         = 153
	EN_GuildBattleTeamNoSameGuild                                 = 154
	EN_BackTeamCommandLeaderInGuildHomeSceneAndYouAreNotSameGuild = 155
	EN_AccecptRandQuestSizeLimitError                             = 156
	EN_Max                                                        = 157
)

func ToName_ErrorNo(id int) string {
	switch id {
	case EN_None:
		return "EN_None"
	case EN_VersionNotMatch:
		return "EN_VersionNotMatch"
	case EN_AccountNameSame:
		return "EN_AccountNameSame"
	case EN_PlayerNameSame:
		return "EN_PlayerNameSame"
	case EN_FilterWord:
		return "EN_FilterWord"
	case EN_CannotfindPlayer:
		return "EN_CannotfindPlayer"
	case EN_AcceptQuestNotFound:
		return "EN_AcceptQuestNotFound"
	case EN_AcceptQuestNoItem:
		return "EN_AcceptQuestNoItem"
	case EN_AcceptSecendDaily:
		return "EN_AcceptSecendDaily"
	case EN_DailyNoNum:
		return "EN_DailyNoNum"
	case EN_AcceptSecendProfession:
		return "EN_AcceptSecendProfession"
	case EN_Battle:
		return "EN_Battle"
	case EN_MoneyLess:
		return "EN_MoneyLess"
	case EN_DiamondLess:
		return "EN_DiamondLess"
	case EN_NoSubSyste:
		return "EN_NoSubSyste"
	case EN_InTeam:
		return "EN_InTeam"
	case EN_NoTeamLeader:
		return "EN_NoTeamLeader"
	case EN_TeamPassword:
		return "EN_TeamPassword"
	case EN_TeamIsFull:
		return "EN_TeamIsFull"
	case EN_NoTeam:
		return "EN_NoTeam"
	case EN_TeamIsRunning:
		return "EN_TeamIsRunning"
	case EN_TeamMemberLeaving:
		return "EN_TeamMemberLeaving"
	case EN_NoBackTeam:
		return "EN_NoBackTeam"
	case EN_InTeamBlackList:
		return "EN_InTeamBlackList"
	case EN_EmployeeIsFull:
		return "EN_EmployeeIsFull"
	case EN_NoUpSkill:
		return "EN_NoUpSkill"
	case EN_PropisNull:
		return "EN_PropisNull"
	case EN_DoubleExpTimeFull:
		return "EN_DoubleExpTimeFull"
	case EN_DoubleExpTimeNULL:
		return "EN_DoubleExpTimeNULL"
	case EN_NoTeamNoTongji:
		return "EN_NoTeamNoTongji"
	case EN_TongjiTimesMax:
		return "EN_TongjiTimesMax"
	case EN_TongjiTeamMemberTimesMax:
		return "EN_TongjiTeamMemberTimesMax"
	case EN_NoTeamLeaderNoTongji:
		return "EN_NoTeamLeaderNoTongji"
	case EN_TeamSizeTongjiError:
		return "EN_TeamSizeTongjiError"
	case EN_GetPoisonMushroom:
		return "EN_GetPoisonMushroom"
	case EN_GetMushroom:
		return "EN_GetMushroom"
	case EN_TongjiTeamLevelTooLow:
		return "EN_TongjiTeamLevelTooLow"
	case EN_PlayerIsInTeam:
		return "EN_PlayerIsInTeam"
	case EN_AcceptQuestBagMax:
		return "EN_AcceptQuestBagMax"
	case EN_SubmitQuestBagMax:
		return "EN_SubmitQuestBagMax"
	case EN_GuildNameSame:
		return "EN_GuildNameSame"
	case EN_PlayerGoldLess:
		return "EN_PlayerGoldLess"
	case EN_PlayerHasGuild:
		return "EN_PlayerHasGuild"
	case EN_InRequestErr:
		return "EN_InRequestErr"
	case EN_RequestListFull:
		return "EN_RequestListFull"
	case EN_joinGuildRequestOk:
		return "EN_joinGuildRequestOk"
	case EN_JoinOtherGuild:
		return "EN_JoinOtherGuild"
	case EN_PremierQuitError:
		return "EN_PremierQuitError"
	case EN_CommandPositionLess:
		return "EN_CommandPositionLess"
	case EN_PositionUpMax:
		return "EN_PositionUpMax"
	case EN_MallBuyOk:
		return "EN_MallBuyOk"
	case EN_MallBuyFailBagFull:
		return "EN_MallBuyFailBagFull"
	case EN_MallBuyFailBabyFull:
		return "EN_MallBuyFailBabyFull"
	case EN_MallBuyFailDiamondLess:
		return "EN_MallBuyFailDiamondLess"
	case EN_MallBuyFailSelled:
		return "EN_MallBuyFailSelled"
	case EN_OpenBaoXiangBagFull:
		return "EN_OpenBaoXiangBagFull"
	case EN_NoBaby:
		return "EN_NoBaby"
	case EN_BagFull:
		return "EN_BagFull"
	case EN_BagSizeMax:
		return "EN_BagSizeMax"
	case EN_BabyStorageFull:
		return "EN_BabyStorageFull"
	case EN_BabyFullToStorage:
		return "EN_BabyFullToStorage"
	case EN_NewItemError:
		return "EN_NewItemError"
	case EN_BabyFull:
		return "EN_BabyFull"
	case EN_RemouldBabyLevel:
		return "EN_RemouldBabyLevel"
	case EN_SkillSoltFull:
		return "EN_SkillSoltFull"
	case EN_WorldChatPayError:
		return "EN_WorldChatPayError"
	case EN_DontTalk:
		return "EN_DontTalk"
	case EN_BadMushroom:
		return "EN_BadMushroom"
	case EN_ItemMushroom:
		return "EN_ItemMushroom"
	case EN_GetMailItemBagFull:
		return "EN_GetMailItemBagFull"
	case EN_Materialless:
		return "EN_Materialless"
	case EN_OpenGatherlose:
		return "EN_OpenGatherlose"
	case EN_OpenGatherRepetition:
		return "EN_OpenGatherRepetition"
	case EN_GatherLevelLess:
		return "EN_GatherLevelLess"
	case EN_GatherTimesLess:
		return "EN_GatherTimesLess"
	case EN_OpenBaoXiangLevel:
		return "EN_OpenBaoXiangLevel"
	case EN_NoBattleBaby:
		return "EN_NoBattleBaby"
	case EN_NoThisPoint:
		return "EN_NoThisPoint"
	case EN_BabyLevelHigh:
		return "EN_BabyLevelHigh"
	case EN_AddMoney1W:
		return "EN_AddMoney1W"
	case EN_AddDionmand100:
		return "EN_AddDionmand100"
	case EN_AddMoney2W:
		return "EN_AddMoney2W"
	case EN_AddDionmand200:
		return "EN_AddDionmand200"
	case EN_AddMoney3W:
		return "EN_AddMoney3W"
	case EN_AddDionmand300:
		return "EN_AddDionmand300"
	case EN_AddMoney4W:
		return "EN_AddMoney4W"
	case EN_AddDionmand400:
		return "EN_AddDionmand400"
	case EN_AddMoney5W:
		return "EN_AddMoney5W"
	case EN_AddDionmand500:
		return "EN_AddDionmand500"
	case EN_AddMoney6W:
		return "EN_AddMoney6W"
	case EN_AddDionmand600:
		return "EN_AddDionmand600"
	case EN_DelBaby1000:
		return "EN_DelBaby1000"
	case EN_DelBaby30:
		return "EN_DelBaby30"
	case EN_DelBaby54:
		return "EN_DelBaby54"
	case EN_DelBaby10030:
		return "EN_DelBaby10030"
	case EN_DelBaby10015:
		return "EN_DelBaby10015"
	case EN_DelBaby10157:
		return "EN_DelBaby10157"
	case EN_DelDefaultSkill:
		return "EN_DelDefaultSkill"
	case EN_GreenBoxFreeNum:
		return "EN_GreenBoxFreeNum"
	case EM_NotNormalVip:
		return "EM_NotNormalVip"
	case EN_NotSuperVip:
		return "EN_NotSuperVip"
	case EN_FirendNotOpen:
		return "EN_FirendNotOpen"
	case EN_BlackCannotFriend:
		return "EN_BlackCannotFriend"
	case EN_HasFriend:
		return "EN_HasFriend"
	case EN_HunderdNoNum:
		return "EN_HunderdNoNum"
	case EN_HunderdTier:
		return "EN_HunderdTier"
	case EN_HunderdLevel:
		return "EN_HunderdLevel"
	case EN_SkillExperr:
		return "EN_SkillExperr"
	case EN_TimeMushroom:
		return "EN_TimeMushroom"
	case EN_TimeTongji:
		return "EN_TimeTongji"
	case EN_TimeXiji:
		return "EN_TimeXiji"
	case EN_Storefull:
		return "EN_Storefull"
	case EN_HasBattleTime:
		return "EN_HasBattleTime"
	case EN_HourLess24:
		return "EN_HourLess24"
	case EN_HourLess24_Join:
		return "EN_HourLess24_Join"
	case EN_TeamMemberHourLess24:
		return "EN_TeamMemberHourLess24"
	case EN_GuildBattleJoinSceneMoveValue:
		return "EN_GuildBattleJoinSceneMoveValue"
	case EN_IdgenNull:
		return "EN_IdgenNull"
	case EN_Idgenhas:
		return "EN_Idgenhas"
	case EN_Gifthas:
		return "EN_Gifthas"
	case EN_UseMakeRepeat:
		return "EN_UseMakeRepeat"
	case EN_WishNull:
		return "EN_WishNull"
	case EN_NoGuild:
		return "EN_NoGuild"
	case EN_GuildMemberMax:
		return "EN_GuildMemberMax"
	case EN_LevelupGuildBuilding:
		return "EN_LevelupGuildBuilding"
	case EN_LevelupGuildBuildingLevelMax:
		return "EN_LevelupGuildBuildingLevelMax"
	case EN_LevelupGuildBuildingMoneyLess:
		return "EN_LevelupGuildBuildingMoneyLess"
	case EN_LevelupGuildBuildingHallBuildLevelLess:
		return "EN_LevelupGuildBuildingHallBuildLevelLess"
	case EN_AddGuildMoneyMax:
		return "EN_AddGuildMoneyMax"
	case EN_PresentGuildOk:
		return "EN_PresentGuildOk"
	case EN_RefreshShopTimeLess:
		return "EN_RefreshShopTimeLess"
	case EN_TeamMemberNoGuild:
		return "EN_TeamMemberNoGuild"
	case EN_MagicCurrencyLess:
		return "EN_MagicCurrencyLess"
	case EN_DisShopLimitLess:
		return "EN_DisShopLimitLess"
	case EN_FamilyPremierCanntDelete:
		return "EN_FamilyPremierCanntDelete"
	case EN_MyFamilyMonster:
		return "EN_MyFamilyMonster"
	case EN_NoBattleTime:
		return "EN_NoBattleTime"
	case EN_OtherNoBattleTime:
		return "EN_OtherNoBattleTime"
	case EN_GuildBattleTimeout2:
		return "EN_GuildBattleTimeout2"
	case EN_GuildBattleHasTeam:
		return "EN_GuildBattleHasTeam"
	case EN_AccountIsSeal:
		return "EN_AccountIsSeal"
	case EN_PlayerNoOnline:
		return "EN_PlayerNoOnline"
	case EN_ActivityNoTime:
		return "EN_ActivityNoTime"
	case EN_NoTeamExist:
		return "EN_NoTeamExist"
	case EN_GuildNoMember:
		return "EN_GuildNoMember"
	case EN_MallSellTimeLess:
		return "EN_MallSellTimeLess"
	case EN_GuildMemberLess24:
		return "EN_GuildMemberLess24"
	case EN_InviteeLeaveGuildLess24:
		return "EN_InviteeLeaveGuildLess24"
	case EN_PboneNumberSuccess:
		return "EN_PboneNumberSuccess"
	case EN_PhoneNumberError:
		return "EN_PhoneNumberError"
	case EN_OtherPlayerInBattle:
		return "EN_OtherPlayerInBattle"
	case EN_GuildBattleNotMatch:
		return "EN_GuildBattleNotMatch"
	case EN_GuildBattleIsStart:
		return "EN_GuildBattleIsStart"
	case EN_GuileBattleIsClose:
		return "EN_GuileBattleIsClose"
	case EN_GuildBattleTeamNoSameGuild:
		return "EN_GuildBattleTeamNoSameGuild"
	case EN_BackTeamCommandLeaderInGuildHomeSceneAndYouAreNotSameGuild:
		return "EN_BackTeamCommandLeaderInGuildHomeSceneAndYouAreNotSameGuild"
	case EN_AccecptRandQuestSizeLimitError:
		return "EN_AccecptRandQuestSizeLimitError"
	case EN_Max:
		return "EN_Max"
	default:
		return ""
	}
}
func ToId_ErrorNo(name string) int {
	switch name {
	case "EN_None":
		return EN_None
	case "EN_VersionNotMatch":
		return EN_VersionNotMatch
	case "EN_AccountNameSame":
		return EN_AccountNameSame
	case "EN_PlayerNameSame":
		return EN_PlayerNameSame
	case "EN_FilterWord":
		return EN_FilterWord
	case "EN_CannotfindPlayer":
		return EN_CannotfindPlayer
	case "EN_AcceptQuestNotFound":
		return EN_AcceptQuestNotFound
	case "EN_AcceptQuestNoItem":
		return EN_AcceptQuestNoItem
	case "EN_AcceptSecendDaily":
		return EN_AcceptSecendDaily
	case "EN_DailyNoNum":
		return EN_DailyNoNum
	case "EN_AcceptSecendProfession":
		return EN_AcceptSecendProfession
	case "EN_Battle":
		return EN_Battle
	case "EN_MoneyLess":
		return EN_MoneyLess
	case "EN_DiamondLess":
		return EN_DiamondLess
	case "EN_NoSubSyste":
		return EN_NoSubSyste
	case "EN_InTeam":
		return EN_InTeam
	case "EN_NoTeamLeader":
		return EN_NoTeamLeader
	case "EN_TeamPassword":
		return EN_TeamPassword
	case "EN_TeamIsFull":
		return EN_TeamIsFull
	case "EN_NoTeam":
		return EN_NoTeam
	case "EN_TeamIsRunning":
		return EN_TeamIsRunning
	case "EN_TeamMemberLeaving":
		return EN_TeamMemberLeaving
	case "EN_NoBackTeam":
		return EN_NoBackTeam
	case "EN_InTeamBlackList":
		return EN_InTeamBlackList
	case "EN_EmployeeIsFull":
		return EN_EmployeeIsFull
	case "EN_NoUpSkill":
		return EN_NoUpSkill
	case "EN_PropisNull":
		return EN_PropisNull
	case "EN_DoubleExpTimeFull":
		return EN_DoubleExpTimeFull
	case "EN_DoubleExpTimeNULL":
		return EN_DoubleExpTimeNULL
	case "EN_NoTeamNoTongji":
		return EN_NoTeamNoTongji
	case "EN_TongjiTimesMax":
		return EN_TongjiTimesMax
	case "EN_TongjiTeamMemberTimesMax":
		return EN_TongjiTeamMemberTimesMax
	case "EN_NoTeamLeaderNoTongji":
		return EN_NoTeamLeaderNoTongji
	case "EN_TeamSizeTongjiError":
		return EN_TeamSizeTongjiError
	case "EN_GetPoisonMushroom":
		return EN_GetPoisonMushroom
	case "EN_GetMushroom":
		return EN_GetMushroom
	case "EN_TongjiTeamLevelTooLow":
		return EN_TongjiTeamLevelTooLow
	case "EN_PlayerIsInTeam":
		return EN_PlayerIsInTeam
	case "EN_AcceptQuestBagMax":
		return EN_AcceptQuestBagMax
	case "EN_SubmitQuestBagMax":
		return EN_SubmitQuestBagMax
	case "EN_GuildNameSame":
		return EN_GuildNameSame
	case "EN_PlayerGoldLess":
		return EN_PlayerGoldLess
	case "EN_PlayerHasGuild":
		return EN_PlayerHasGuild
	case "EN_InRequestErr":
		return EN_InRequestErr
	case "EN_RequestListFull":
		return EN_RequestListFull
	case "EN_joinGuildRequestOk":
		return EN_joinGuildRequestOk
	case "EN_JoinOtherGuild":
		return EN_JoinOtherGuild
	case "EN_PremierQuitError":
		return EN_PremierQuitError
	case "EN_CommandPositionLess":
		return EN_CommandPositionLess
	case "EN_PositionUpMax":
		return EN_PositionUpMax
	case "EN_MallBuyOk":
		return EN_MallBuyOk
	case "EN_MallBuyFailBagFull":
		return EN_MallBuyFailBagFull
	case "EN_MallBuyFailBabyFull":
		return EN_MallBuyFailBabyFull
	case "EN_MallBuyFailDiamondLess":
		return EN_MallBuyFailDiamondLess
	case "EN_MallBuyFailSelled":
		return EN_MallBuyFailSelled
	case "EN_OpenBaoXiangBagFull":
		return EN_OpenBaoXiangBagFull
	case "EN_NoBaby":
		return EN_NoBaby
	case "EN_BagFull":
		return EN_BagFull
	case "EN_BagSizeMax":
		return EN_BagSizeMax
	case "EN_BabyStorageFull":
		return EN_BabyStorageFull
	case "EN_BabyFullToStorage":
		return EN_BabyFullToStorage
	case "EN_NewItemError":
		return EN_NewItemError
	case "EN_BabyFull":
		return EN_BabyFull
	case "EN_RemouldBabyLevel":
		return EN_RemouldBabyLevel
	case "EN_SkillSoltFull":
		return EN_SkillSoltFull
	case "EN_WorldChatPayError":
		return EN_WorldChatPayError
	case "EN_DontTalk":
		return EN_DontTalk
	case "EN_BadMushroom":
		return EN_BadMushroom
	case "EN_ItemMushroom":
		return EN_ItemMushroom
	case "EN_GetMailItemBagFull":
		return EN_GetMailItemBagFull
	case "EN_Materialless":
		return EN_Materialless
	case "EN_OpenGatherlose":
		return EN_OpenGatherlose
	case "EN_OpenGatherRepetition":
		return EN_OpenGatherRepetition
	case "EN_GatherLevelLess":
		return EN_GatherLevelLess
	case "EN_GatherTimesLess":
		return EN_GatherTimesLess
	case "EN_OpenBaoXiangLevel":
		return EN_OpenBaoXiangLevel
	case "EN_NoBattleBaby":
		return EN_NoBattleBaby
	case "EN_NoThisPoint":
		return EN_NoThisPoint
	case "EN_BabyLevelHigh":
		return EN_BabyLevelHigh
	case "EN_AddMoney1W":
		return EN_AddMoney1W
	case "EN_AddDionmand100":
		return EN_AddDionmand100
	case "EN_AddMoney2W":
		return EN_AddMoney2W
	case "EN_AddDionmand200":
		return EN_AddDionmand200
	case "EN_AddMoney3W":
		return EN_AddMoney3W
	case "EN_AddDionmand300":
		return EN_AddDionmand300
	case "EN_AddMoney4W":
		return EN_AddMoney4W
	case "EN_AddDionmand400":
		return EN_AddDionmand400
	case "EN_AddMoney5W":
		return EN_AddMoney5W
	case "EN_AddDionmand500":
		return EN_AddDionmand500
	case "EN_AddMoney6W":
		return EN_AddMoney6W
	case "EN_AddDionmand600":
		return EN_AddDionmand600
	case "EN_DelBaby1000":
		return EN_DelBaby1000
	case "EN_DelBaby30":
		return EN_DelBaby30
	case "EN_DelBaby54":
		return EN_DelBaby54
	case "EN_DelBaby10030":
		return EN_DelBaby10030
	case "EN_DelBaby10015":
		return EN_DelBaby10015
	case "EN_DelBaby10157":
		return EN_DelBaby10157
	case "EN_DelDefaultSkill":
		return EN_DelDefaultSkill
	case "EN_GreenBoxFreeNum":
		return EN_GreenBoxFreeNum
	case "EM_NotNormalVip":
		return EM_NotNormalVip
	case "EN_NotSuperVip":
		return EN_NotSuperVip
	case "EN_FirendNotOpen":
		return EN_FirendNotOpen
	case "EN_BlackCannotFriend":
		return EN_BlackCannotFriend
	case "EN_HasFriend":
		return EN_HasFriend
	case "EN_HunderdNoNum":
		return EN_HunderdNoNum
	case "EN_HunderdTier":
		return EN_HunderdTier
	case "EN_HunderdLevel":
		return EN_HunderdLevel
	case "EN_SkillExperr":
		return EN_SkillExperr
	case "EN_TimeMushroom":
		return EN_TimeMushroom
	case "EN_TimeTongji":
		return EN_TimeTongji
	case "EN_TimeXiji":
		return EN_TimeXiji
	case "EN_Storefull":
		return EN_Storefull
	case "EN_HasBattleTime":
		return EN_HasBattleTime
	case "EN_HourLess24":
		return EN_HourLess24
	case "EN_HourLess24_Join":
		return EN_HourLess24_Join
	case "EN_TeamMemberHourLess24":
		return EN_TeamMemberHourLess24
	case "EN_GuildBattleJoinSceneMoveValue":
		return EN_GuildBattleJoinSceneMoveValue
	case "EN_IdgenNull":
		return EN_IdgenNull
	case "EN_Idgenhas":
		return EN_Idgenhas
	case "EN_Gifthas":
		return EN_Gifthas
	case "EN_UseMakeRepeat":
		return EN_UseMakeRepeat
	case "EN_WishNull":
		return EN_WishNull
	case "EN_NoGuild":
		return EN_NoGuild
	case "EN_GuildMemberMax":
		return EN_GuildMemberMax
	case "EN_LevelupGuildBuilding":
		return EN_LevelupGuildBuilding
	case "EN_LevelupGuildBuildingLevelMax":
		return EN_LevelupGuildBuildingLevelMax
	case "EN_LevelupGuildBuildingMoneyLess":
		return EN_LevelupGuildBuildingMoneyLess
	case "EN_LevelupGuildBuildingHallBuildLevelLess":
		return EN_LevelupGuildBuildingHallBuildLevelLess
	case "EN_AddGuildMoneyMax":
		return EN_AddGuildMoneyMax
	case "EN_PresentGuildOk":
		return EN_PresentGuildOk
	case "EN_RefreshShopTimeLess":
		return EN_RefreshShopTimeLess
	case "EN_TeamMemberNoGuild":
		return EN_TeamMemberNoGuild
	case "EN_MagicCurrencyLess":
		return EN_MagicCurrencyLess
	case "EN_DisShopLimitLess":
		return EN_DisShopLimitLess
	case "EN_FamilyPremierCanntDelete":
		return EN_FamilyPremierCanntDelete
	case "EN_MyFamilyMonster":
		return EN_MyFamilyMonster
	case "EN_NoBattleTime":
		return EN_NoBattleTime
	case "EN_OtherNoBattleTime":
		return EN_OtherNoBattleTime
	case "EN_GuildBattleTimeout2":
		return EN_GuildBattleTimeout2
	case "EN_GuildBattleHasTeam":
		return EN_GuildBattleHasTeam
	case "EN_AccountIsSeal":
		return EN_AccountIsSeal
	case "EN_PlayerNoOnline":
		return EN_PlayerNoOnline
	case "EN_ActivityNoTime":
		return EN_ActivityNoTime
	case "EN_NoTeamExist":
		return EN_NoTeamExist
	case "EN_GuildNoMember":
		return EN_GuildNoMember
	case "EN_MallSellTimeLess":
		return EN_MallSellTimeLess
	case "EN_GuildMemberLess24":
		return EN_GuildMemberLess24
	case "EN_InviteeLeaveGuildLess24":
		return EN_InviteeLeaveGuildLess24
	case "EN_PboneNumberSuccess":
		return EN_PboneNumberSuccess
	case "EN_PhoneNumberError":
		return EN_PhoneNumberError
	case "EN_OtherPlayerInBattle":
		return EN_OtherPlayerInBattle
	case "EN_GuildBattleNotMatch":
		return EN_GuildBattleNotMatch
	case "EN_GuildBattleIsStart":
		return EN_GuildBattleIsStart
	case "EN_GuileBattleIsClose":
		return EN_GuileBattleIsClose
	case "EN_GuildBattleTeamNoSameGuild":
		return EN_GuildBattleTeamNoSameGuild
	case "EN_BackTeamCommandLeaderInGuildHomeSceneAndYouAreNotSameGuild":
		return EN_BackTeamCommandLeaderInGuildHomeSceneAndYouAreNotSameGuild
	case "EN_AccecptRandQuestSizeLimitError":
		return EN_AccecptRandQuestSizeLimitError
	case "EN_Max":
		return EN_Max
	default:
		return -1
	}
}
