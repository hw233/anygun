package prpc

// enum ScriptGameEvent
const (
	SGE_None                           = 0
	SGE_MainPanelOpen                  = 1
	SGE_FirstEnterMainScene            = 2
	SGE_EnterScene                     = 3
	SGE_EnterMainScene                 = 4
	SGE_Talk_FirstEnterMainScene       = 5
	SGE_Talk_BattleReady               = 6
	SGE_Talk_ActorReady                = 7
	SGE_Talk_BattleOver                = 8
	SGE_WorldMapOpen                   = 9
	SGE_WorldMapToWorld                = 10
	SGE_MiniMapOpen                    = 11
	SGE_TeamUIOpen                     = 12
	SGE_AchievementUIOpen              = 13
	SGE_AchievementReceived            = 14
	SGE_TeamUISelectMapOpen            = 15
	SGE_PartnerForBattle               = 16
	SGE_PartnerPositionTabShow         = 17
	SGE_PartnerListTabShow             = 18
	SGE_PartnerDetailUIOpen            = 19
	SGE_PartnerDetailBaseOpen          = 20
	SGE_PartnerDetailEquipClick        = 21
	SGE_PartnerDetailEquipSucc         = 22
	SGE_PartnerDetailBaseSkillOpen     = 23
	SGE_ParnterDetailBaseSkillLvUpSucc = 24
	SGE_MainMakeSub                    = 25
	SGE_MainMakeSubDetail              = 26
	SGE_MainMakeItemOk                 = 27
	SGE_MainMakeGemUI                  = 28
	SGE_MainMakeGemOk                  = 29
	SGE_MainMakeGemUIClose             = 30
	SGE_MainTeamUI                     = 31
	SGE_MainTaskUI                     = 32
	SGE_MainTaskFlushOk                = 33
	SGE_JJCEntryUI                     = 34
	SGE_OfflineJJCUI                   = 35
	SGE_BagItemDoubleClick             = 36
	SGE_NpcDialogBegin                 = 37
	SGE_NpcRenwuUIOpen                 = 38
	SGE_NpcRenwuPreAccept              = 39
	SGE_NpcRenwuAccept                 = 40
	SGE_NpcRenwuSubmit                 = 41
	SGE_EnterNPCBattle                 = 42
	SGE_EnterBattle                    = 43
	SGE_MessageBoxOpen                 = 44
	SGE_BeforeEnterBattle              = 45
	SGE_PlayerLevelUp                  = 46
	SGE_PlayerUIOpen                   = 47
	SGE_PlayerUIStatusSwitch           = 48
	SGE_PlayerUIPropertySwitch         = 49
	SGE_PlayerUIAddPoint               = 50
	SGE_PlayerUIPropertyConfirmClick   = 51
	SGE_PlayerUIClose                  = 52
	SGE_BabyLevelUp                    = 53
	SGE_BabyUIOpen                     = 54
	SGE_BabyUIStatusSwitch             = 55
	SGE_BabyUIPropertySwitch           = 56
	SGE_BabyUIAddPoint                 = 57
	SGE_BabyUIPropertyConfirmClick     = 58
	SGE_BabyUIClose                    = 59
	SGE_BattleTurn                     = 60
	SGE_ExitBattle                     = 61
	SGE_SelectTarget                   = 62
	SGE_SelectTargetOk                 = 63
	SGE_BabySelectSkill                = 64
	SGE_SelectSkill                    = 65
	SGE_SelectSkillLevel               = 66
	SGE_StickDisplay                   = 67
	SGE_StickTouchDown                 = 68
	SGE_StickTouchMove                 = 69
	SGE_StickTouchUp                   = 70
	SGE_BattleOverRewardOpen           = 71
	SGE_BattleOverRewardClose          = 72
	SGE_CheckState                     = 73
	SGE_TogetherState                  = 74
	SGE_BagTipOpen                     = 75
	SGE_UseItem                        = 76
	SGE_GainItem                       = 77
	SGE_EquipItem                      = 78
	SGE_MainLearningUI                 = 79
	SGE_MainLearningClickTab           = 80
	SGE_MainLearningOneSkillClick      = 81
	SGE_MainLearningSkillOk            = 82
	SGE_MainMakeUIOpen                 = 83
	SGE_MainMagicTipClose              = 84
	SGE_MainMagicUIOpen                = 85
	SGE_MainMagicLevelUp               = 86
	SGE_MainMagicFirstClick            = 87
	SGE_PartnerShowUI                  = 88
	SGE_PartnerHideUI                  = 89
	SGE_BabyLearningSkillUI            = 90
	SGE_BabyLearningSkill_BabyListUI   = 91
	SGE_BabyLearningSkill_BabySkillUI  = 92
	SGE_BabyLearningSkillOk            = 93
	SGE_ClickBabyLearningSkill         = 94
	SGE_ClickMiniQuest                 = 95
	SGE_ClickMainBag                   = 96
	SGE_ClickMainSkill                 = 97
	SGE_ClickMainBaby                  = 98
	SGE_ClickMainFriend                = 99
	SGE_ClickMainPartner               = 100
	SGE_ClickMainSetting               = 101
	SGE_ClickMainRecharge              = 102
	SGE_ClickMainActivity              = 103
	SGE_ClickMainInfo                  = 104
	SGE_ClickBattleAttack              = 105
	SGE_ClickBattleSkill               = 106
	SGE_ClickBattleBabySkill           = 107
	SGE_ClickBattleDeference           = 108
	SGE_ClickBattleChangePosition      = 109
	SGE_ClickBattleAuto                = 110
	SGE_ClickBattleBag                 = 111
	SGE_ClickBattleBaby                = 112
	SGE_ClickBattleRunaway             = 113
	SGE_ClickBattleInfo                = 114
	SGE_ClickAddFriendBtn              = 115
	SGE_ClickMainFamily                = 116
	SGE_ClickRaiseUpBtn                = 117
	SGE_UseItemOk                      = 118
	SGE_CheckBuff                      = 119
	SGE_PlayerPropUpdate               = 120
	SGE_NpcTalked                      = 121
	SGE_EnterCopy                      = 122
	SGE_SenseEnter2                    = 123
	SGE_WaitTalk                       = 124
	SGE_SenseTalked                    = 125
	SGE_ExitSense                      = 126
	SGE_Max                            = 127
)

func ToName_ScriptGameEvent(id int) string {
	switch id {
	case SGE_None:
		return "SGE_None"
	case SGE_MainPanelOpen:
		return "SGE_MainPanelOpen"
	case SGE_FirstEnterMainScene:
		return "SGE_FirstEnterMainScene"
	case SGE_EnterScene:
		return "SGE_EnterScene"
	case SGE_EnterMainScene:
		return "SGE_EnterMainScene"
	case SGE_Talk_FirstEnterMainScene:
		return "SGE_Talk_FirstEnterMainScene"
	case SGE_Talk_BattleReady:
		return "SGE_Talk_BattleReady"
	case SGE_Talk_ActorReady:
		return "SGE_Talk_ActorReady"
	case SGE_Talk_BattleOver:
		return "SGE_Talk_BattleOver"
	case SGE_WorldMapOpen:
		return "SGE_WorldMapOpen"
	case SGE_WorldMapToWorld:
		return "SGE_WorldMapToWorld"
	case SGE_MiniMapOpen:
		return "SGE_MiniMapOpen"
	case SGE_TeamUIOpen:
		return "SGE_TeamUIOpen"
	case SGE_AchievementUIOpen:
		return "SGE_AchievementUIOpen"
	case SGE_AchievementReceived:
		return "SGE_AchievementReceived"
	case SGE_TeamUISelectMapOpen:
		return "SGE_TeamUISelectMapOpen"
	case SGE_PartnerForBattle:
		return "SGE_PartnerForBattle"
	case SGE_PartnerPositionTabShow:
		return "SGE_PartnerPositionTabShow"
	case SGE_PartnerListTabShow:
		return "SGE_PartnerListTabShow"
	case SGE_PartnerDetailUIOpen:
		return "SGE_PartnerDetailUIOpen"
	case SGE_PartnerDetailBaseOpen:
		return "SGE_PartnerDetailBaseOpen"
	case SGE_PartnerDetailEquipClick:
		return "SGE_PartnerDetailEquipClick"
	case SGE_PartnerDetailEquipSucc:
		return "SGE_PartnerDetailEquipSucc"
	case SGE_PartnerDetailBaseSkillOpen:
		return "SGE_PartnerDetailBaseSkillOpen"
	case SGE_ParnterDetailBaseSkillLvUpSucc:
		return "SGE_ParnterDetailBaseSkillLvUpSucc"
	case SGE_MainMakeSub:
		return "SGE_MainMakeSub"
	case SGE_MainMakeSubDetail:
		return "SGE_MainMakeSubDetail"
	case SGE_MainMakeItemOk:
		return "SGE_MainMakeItemOk"
	case SGE_MainMakeGemUI:
		return "SGE_MainMakeGemUI"
	case SGE_MainMakeGemOk:
		return "SGE_MainMakeGemOk"
	case SGE_MainMakeGemUIClose:
		return "SGE_MainMakeGemUIClose"
	case SGE_MainTeamUI:
		return "SGE_MainTeamUI"
	case SGE_MainTaskUI:
		return "SGE_MainTaskUI"
	case SGE_MainTaskFlushOk:
		return "SGE_MainTaskFlushOk"
	case SGE_JJCEntryUI:
		return "SGE_JJCEntryUI"
	case SGE_OfflineJJCUI:
		return "SGE_OfflineJJCUI"
	case SGE_BagItemDoubleClick:
		return "SGE_BagItemDoubleClick"
	case SGE_NpcDialogBegin:
		return "SGE_NpcDialogBegin"
	case SGE_NpcRenwuUIOpen:
		return "SGE_NpcRenwuUIOpen"
	case SGE_NpcRenwuPreAccept:
		return "SGE_NpcRenwuPreAccept"
	case SGE_NpcRenwuAccept:
		return "SGE_NpcRenwuAccept"
	case SGE_NpcRenwuSubmit:
		return "SGE_NpcRenwuSubmit"
	case SGE_EnterNPCBattle:
		return "SGE_EnterNPCBattle"
	case SGE_EnterBattle:
		return "SGE_EnterBattle"
	case SGE_MessageBoxOpen:
		return "SGE_MessageBoxOpen"
	case SGE_BeforeEnterBattle:
		return "SGE_BeforeEnterBattle"
	case SGE_PlayerLevelUp:
		return "SGE_PlayerLevelUp"
	case SGE_PlayerUIOpen:
		return "SGE_PlayerUIOpen"
	case SGE_PlayerUIStatusSwitch:
		return "SGE_PlayerUIStatusSwitch"
	case SGE_PlayerUIPropertySwitch:
		return "SGE_PlayerUIPropertySwitch"
	case SGE_PlayerUIAddPoint:
		return "SGE_PlayerUIAddPoint"
	case SGE_PlayerUIPropertyConfirmClick:
		return "SGE_PlayerUIPropertyConfirmClick"
	case SGE_PlayerUIClose:
		return "SGE_PlayerUIClose"
	case SGE_BabyLevelUp:
		return "SGE_BabyLevelUp"
	case SGE_BabyUIOpen:
		return "SGE_BabyUIOpen"
	case SGE_BabyUIStatusSwitch:
		return "SGE_BabyUIStatusSwitch"
	case SGE_BabyUIPropertySwitch:
		return "SGE_BabyUIPropertySwitch"
	case SGE_BabyUIAddPoint:
		return "SGE_BabyUIAddPoint"
	case SGE_BabyUIPropertyConfirmClick:
		return "SGE_BabyUIPropertyConfirmClick"
	case SGE_BabyUIClose:
		return "SGE_BabyUIClose"
	case SGE_BattleTurn:
		return "SGE_BattleTurn"
	case SGE_ExitBattle:
		return "SGE_ExitBattle"
	case SGE_SelectTarget:
		return "SGE_SelectTarget"
	case SGE_SelectTargetOk:
		return "SGE_SelectTargetOk"
	case SGE_BabySelectSkill:
		return "SGE_BabySelectSkill"
	case SGE_SelectSkill:
		return "SGE_SelectSkill"
	case SGE_SelectSkillLevel:
		return "SGE_SelectSkillLevel"
	case SGE_StickDisplay:
		return "SGE_StickDisplay"
	case SGE_StickTouchDown:
		return "SGE_StickTouchDown"
	case SGE_StickTouchMove:
		return "SGE_StickTouchMove"
	case SGE_StickTouchUp:
		return "SGE_StickTouchUp"
	case SGE_BattleOverRewardOpen:
		return "SGE_BattleOverRewardOpen"
	case SGE_BattleOverRewardClose:
		return "SGE_BattleOverRewardClose"
	case SGE_CheckState:
		return "SGE_CheckState"
	case SGE_TogetherState:
		return "SGE_TogetherState"
	case SGE_BagTipOpen:
		return "SGE_BagTipOpen"
	case SGE_UseItem:
		return "SGE_UseItem"
	case SGE_GainItem:
		return "SGE_GainItem"
	case SGE_EquipItem:
		return "SGE_EquipItem"
	case SGE_MainLearningUI:
		return "SGE_MainLearningUI"
	case SGE_MainLearningClickTab:
		return "SGE_MainLearningClickTab"
	case SGE_MainLearningOneSkillClick:
		return "SGE_MainLearningOneSkillClick"
	case SGE_MainLearningSkillOk:
		return "SGE_MainLearningSkillOk"
	case SGE_MainMakeUIOpen:
		return "SGE_MainMakeUIOpen"
	case SGE_MainMagicTipClose:
		return "SGE_MainMagicTipClose"
	case SGE_MainMagicUIOpen:
		return "SGE_MainMagicUIOpen"
	case SGE_MainMagicLevelUp:
		return "SGE_MainMagicLevelUp"
	case SGE_MainMagicFirstClick:
		return "SGE_MainMagicFirstClick"
	case SGE_PartnerShowUI:
		return "SGE_PartnerShowUI"
	case SGE_PartnerHideUI:
		return "SGE_PartnerHideUI"
	case SGE_BabyLearningSkillUI:
		return "SGE_BabyLearningSkillUI"
	case SGE_BabyLearningSkill_BabyListUI:
		return "SGE_BabyLearningSkill_BabyListUI"
	case SGE_BabyLearningSkill_BabySkillUI:
		return "SGE_BabyLearningSkill_BabySkillUI"
	case SGE_BabyLearningSkillOk:
		return "SGE_BabyLearningSkillOk"
	case SGE_ClickBabyLearningSkill:
		return "SGE_ClickBabyLearningSkill"
	case SGE_ClickMiniQuest:
		return "SGE_ClickMiniQuest"
	case SGE_ClickMainBag:
		return "SGE_ClickMainBag"
	case SGE_ClickMainSkill:
		return "SGE_ClickMainSkill"
	case SGE_ClickMainBaby:
		return "SGE_ClickMainBaby"
	case SGE_ClickMainFriend:
		return "SGE_ClickMainFriend"
	case SGE_ClickMainPartner:
		return "SGE_ClickMainPartner"
	case SGE_ClickMainSetting:
		return "SGE_ClickMainSetting"
	case SGE_ClickMainRecharge:
		return "SGE_ClickMainRecharge"
	case SGE_ClickMainActivity:
		return "SGE_ClickMainActivity"
	case SGE_ClickMainInfo:
		return "SGE_ClickMainInfo"
	case SGE_ClickBattleAttack:
		return "SGE_ClickBattleAttack"
	case SGE_ClickBattleSkill:
		return "SGE_ClickBattleSkill"
	case SGE_ClickBattleBabySkill:
		return "SGE_ClickBattleBabySkill"
	case SGE_ClickBattleDeference:
		return "SGE_ClickBattleDeference"
	case SGE_ClickBattleChangePosition:
		return "SGE_ClickBattleChangePosition"
	case SGE_ClickBattleAuto:
		return "SGE_ClickBattleAuto"
	case SGE_ClickBattleBag:
		return "SGE_ClickBattleBag"
	case SGE_ClickBattleBaby:
		return "SGE_ClickBattleBaby"
	case SGE_ClickBattleRunaway:
		return "SGE_ClickBattleRunaway"
	case SGE_ClickBattleInfo:
		return "SGE_ClickBattleInfo"
	case SGE_ClickAddFriendBtn:
		return "SGE_ClickAddFriendBtn"
	case SGE_ClickMainFamily:
		return "SGE_ClickMainFamily"
	case SGE_ClickRaiseUpBtn:
		return "SGE_ClickRaiseUpBtn"
	case SGE_UseItemOk:
		return "SGE_UseItemOk"
	case SGE_CheckBuff:
		return "SGE_CheckBuff"
	case SGE_PlayerPropUpdate:
		return "SGE_PlayerPropUpdate"
	case SGE_NpcTalked:
		return "SGE_NpcTalked"
	case SGE_EnterCopy:
		return "SGE_EnterCopy"
	case SGE_SenseEnter2:
		return "SGE_SenseEnter2"
	case SGE_WaitTalk:
		return "SGE_WaitTalk"
	case SGE_SenseTalked:
		return "SGE_SenseTalked"
	case SGE_ExitSense:
		return "SGE_ExitSense"
	case SGE_Max:
		return "SGE_Max"
	default:
		return ""
	}
}
func ToId_ScriptGameEvent(name string) int {
	switch name {
	case "SGE_None":
		return SGE_None
	case "SGE_MainPanelOpen":
		return SGE_MainPanelOpen
	case "SGE_FirstEnterMainScene":
		return SGE_FirstEnterMainScene
	case "SGE_EnterScene":
		return SGE_EnterScene
	case "SGE_EnterMainScene":
		return SGE_EnterMainScene
	case "SGE_Talk_FirstEnterMainScene":
		return SGE_Talk_FirstEnterMainScene
	case "SGE_Talk_BattleReady":
		return SGE_Talk_BattleReady
	case "SGE_Talk_ActorReady":
		return SGE_Talk_ActorReady
	case "SGE_Talk_BattleOver":
		return SGE_Talk_BattleOver
	case "SGE_WorldMapOpen":
		return SGE_WorldMapOpen
	case "SGE_WorldMapToWorld":
		return SGE_WorldMapToWorld
	case "SGE_MiniMapOpen":
		return SGE_MiniMapOpen
	case "SGE_TeamUIOpen":
		return SGE_TeamUIOpen
	case "SGE_AchievementUIOpen":
		return SGE_AchievementUIOpen
	case "SGE_AchievementReceived":
		return SGE_AchievementReceived
	case "SGE_TeamUISelectMapOpen":
		return SGE_TeamUISelectMapOpen
	case "SGE_PartnerForBattle":
		return SGE_PartnerForBattle
	case "SGE_PartnerPositionTabShow":
		return SGE_PartnerPositionTabShow
	case "SGE_PartnerListTabShow":
		return SGE_PartnerListTabShow
	case "SGE_PartnerDetailUIOpen":
		return SGE_PartnerDetailUIOpen
	case "SGE_PartnerDetailBaseOpen":
		return SGE_PartnerDetailBaseOpen
	case "SGE_PartnerDetailEquipClick":
		return SGE_PartnerDetailEquipClick
	case "SGE_PartnerDetailEquipSucc":
		return SGE_PartnerDetailEquipSucc
	case "SGE_PartnerDetailBaseSkillOpen":
		return SGE_PartnerDetailBaseSkillOpen
	case "SGE_ParnterDetailBaseSkillLvUpSucc":
		return SGE_ParnterDetailBaseSkillLvUpSucc
	case "SGE_MainMakeSub":
		return SGE_MainMakeSub
	case "SGE_MainMakeSubDetail":
		return SGE_MainMakeSubDetail
	case "SGE_MainMakeItemOk":
		return SGE_MainMakeItemOk
	case "SGE_MainMakeGemUI":
		return SGE_MainMakeGemUI
	case "SGE_MainMakeGemOk":
		return SGE_MainMakeGemOk
	case "SGE_MainMakeGemUIClose":
		return SGE_MainMakeGemUIClose
	case "SGE_MainTeamUI":
		return SGE_MainTeamUI
	case "SGE_MainTaskUI":
		return SGE_MainTaskUI
	case "SGE_MainTaskFlushOk":
		return SGE_MainTaskFlushOk
	case "SGE_JJCEntryUI":
		return SGE_JJCEntryUI
	case "SGE_OfflineJJCUI":
		return SGE_OfflineJJCUI
	case "SGE_BagItemDoubleClick":
		return SGE_BagItemDoubleClick
	case "SGE_NpcDialogBegin":
		return SGE_NpcDialogBegin
	case "SGE_NpcRenwuUIOpen":
		return SGE_NpcRenwuUIOpen
	case "SGE_NpcRenwuPreAccept":
		return SGE_NpcRenwuPreAccept
	case "SGE_NpcRenwuAccept":
		return SGE_NpcRenwuAccept
	case "SGE_NpcRenwuSubmit":
		return SGE_NpcRenwuSubmit
	case "SGE_EnterNPCBattle":
		return SGE_EnterNPCBattle
	case "SGE_EnterBattle":
		return SGE_EnterBattle
	case "SGE_MessageBoxOpen":
		return SGE_MessageBoxOpen
	case "SGE_BeforeEnterBattle":
		return SGE_BeforeEnterBattle
	case "SGE_PlayerLevelUp":
		return SGE_PlayerLevelUp
	case "SGE_PlayerUIOpen":
		return SGE_PlayerUIOpen
	case "SGE_PlayerUIStatusSwitch":
		return SGE_PlayerUIStatusSwitch
	case "SGE_PlayerUIPropertySwitch":
		return SGE_PlayerUIPropertySwitch
	case "SGE_PlayerUIAddPoint":
		return SGE_PlayerUIAddPoint
	case "SGE_PlayerUIPropertyConfirmClick":
		return SGE_PlayerUIPropertyConfirmClick
	case "SGE_PlayerUIClose":
		return SGE_PlayerUIClose
	case "SGE_BabyLevelUp":
		return SGE_BabyLevelUp
	case "SGE_BabyUIOpen":
		return SGE_BabyUIOpen
	case "SGE_BabyUIStatusSwitch":
		return SGE_BabyUIStatusSwitch
	case "SGE_BabyUIPropertySwitch":
		return SGE_BabyUIPropertySwitch
	case "SGE_BabyUIAddPoint":
		return SGE_BabyUIAddPoint
	case "SGE_BabyUIPropertyConfirmClick":
		return SGE_BabyUIPropertyConfirmClick
	case "SGE_BabyUIClose":
		return SGE_BabyUIClose
	case "SGE_BattleTurn":
		return SGE_BattleTurn
	case "SGE_ExitBattle":
		return SGE_ExitBattle
	case "SGE_SelectTarget":
		return SGE_SelectTarget
	case "SGE_SelectTargetOk":
		return SGE_SelectTargetOk
	case "SGE_BabySelectSkill":
		return SGE_BabySelectSkill
	case "SGE_SelectSkill":
		return SGE_SelectSkill
	case "SGE_SelectSkillLevel":
		return SGE_SelectSkillLevel
	case "SGE_StickDisplay":
		return SGE_StickDisplay
	case "SGE_StickTouchDown":
		return SGE_StickTouchDown
	case "SGE_StickTouchMove":
		return SGE_StickTouchMove
	case "SGE_StickTouchUp":
		return SGE_StickTouchUp
	case "SGE_BattleOverRewardOpen":
		return SGE_BattleOverRewardOpen
	case "SGE_BattleOverRewardClose":
		return SGE_BattleOverRewardClose
	case "SGE_CheckState":
		return SGE_CheckState
	case "SGE_TogetherState":
		return SGE_TogetherState
	case "SGE_BagTipOpen":
		return SGE_BagTipOpen
	case "SGE_UseItem":
		return SGE_UseItem
	case "SGE_GainItem":
		return SGE_GainItem
	case "SGE_EquipItem":
		return SGE_EquipItem
	case "SGE_MainLearningUI":
		return SGE_MainLearningUI
	case "SGE_MainLearningClickTab":
		return SGE_MainLearningClickTab
	case "SGE_MainLearningOneSkillClick":
		return SGE_MainLearningOneSkillClick
	case "SGE_MainLearningSkillOk":
		return SGE_MainLearningSkillOk
	case "SGE_MainMakeUIOpen":
		return SGE_MainMakeUIOpen
	case "SGE_MainMagicTipClose":
		return SGE_MainMagicTipClose
	case "SGE_MainMagicUIOpen":
		return SGE_MainMagicUIOpen
	case "SGE_MainMagicLevelUp":
		return SGE_MainMagicLevelUp
	case "SGE_MainMagicFirstClick":
		return SGE_MainMagicFirstClick
	case "SGE_PartnerShowUI":
		return SGE_PartnerShowUI
	case "SGE_PartnerHideUI":
		return SGE_PartnerHideUI
	case "SGE_BabyLearningSkillUI":
		return SGE_BabyLearningSkillUI
	case "SGE_BabyLearningSkill_BabyListUI":
		return SGE_BabyLearningSkill_BabyListUI
	case "SGE_BabyLearningSkill_BabySkillUI":
		return SGE_BabyLearningSkill_BabySkillUI
	case "SGE_BabyLearningSkillOk":
		return SGE_BabyLearningSkillOk
	case "SGE_ClickBabyLearningSkill":
		return SGE_ClickBabyLearningSkill
	case "SGE_ClickMiniQuest":
		return SGE_ClickMiniQuest
	case "SGE_ClickMainBag":
		return SGE_ClickMainBag
	case "SGE_ClickMainSkill":
		return SGE_ClickMainSkill
	case "SGE_ClickMainBaby":
		return SGE_ClickMainBaby
	case "SGE_ClickMainFriend":
		return SGE_ClickMainFriend
	case "SGE_ClickMainPartner":
		return SGE_ClickMainPartner
	case "SGE_ClickMainSetting":
		return SGE_ClickMainSetting
	case "SGE_ClickMainRecharge":
		return SGE_ClickMainRecharge
	case "SGE_ClickMainActivity":
		return SGE_ClickMainActivity
	case "SGE_ClickMainInfo":
		return SGE_ClickMainInfo
	case "SGE_ClickBattleAttack":
		return SGE_ClickBattleAttack
	case "SGE_ClickBattleSkill":
		return SGE_ClickBattleSkill
	case "SGE_ClickBattleBabySkill":
		return SGE_ClickBattleBabySkill
	case "SGE_ClickBattleDeference":
		return SGE_ClickBattleDeference
	case "SGE_ClickBattleChangePosition":
		return SGE_ClickBattleChangePosition
	case "SGE_ClickBattleAuto":
		return SGE_ClickBattleAuto
	case "SGE_ClickBattleBag":
		return SGE_ClickBattleBag
	case "SGE_ClickBattleBaby":
		return SGE_ClickBattleBaby
	case "SGE_ClickBattleRunaway":
		return SGE_ClickBattleRunaway
	case "SGE_ClickBattleInfo":
		return SGE_ClickBattleInfo
	case "SGE_ClickAddFriendBtn":
		return SGE_ClickAddFriendBtn
	case "SGE_ClickMainFamily":
		return SGE_ClickMainFamily
	case "SGE_ClickRaiseUpBtn":
		return SGE_ClickRaiseUpBtn
	case "SGE_UseItemOk":
		return SGE_UseItemOk
	case "SGE_CheckBuff":
		return SGE_CheckBuff
	case "SGE_PlayerPropUpdate":
		return SGE_PlayerPropUpdate
	case "SGE_NpcTalked":
		return SGE_NpcTalked
	case "SGE_EnterCopy":
		return SGE_EnterCopy
	case "SGE_SenseEnter2":
		return SGE_SenseEnter2
	case "SGE_WaitTalk":
		return SGE_WaitTalk
	case "SGE_SenseTalked":
		return SGE_SenseTalked
	case "SGE_ExitSense":
		return SGE_ExitSense
	case "SGE_Max":
		return SGE_Max
	default:
		return -1
	}
}
