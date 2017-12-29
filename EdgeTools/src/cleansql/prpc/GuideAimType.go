package prpc

// enum GuideAimType
const (
	GAT_None                             = 0
	GAT_FirstAchievement                 = 1
	GAT_FirstSkill                       = 2
	GAT_FirstLevelSkill                  = 3
	GAT_FirstQuest                       = 4
	GAT_DialogUI                         = 5
	GAT_MainTeamBtn                      = 6
	GAT_MainTaskBtn                      = 7
	GAT_QuestMiniFirst                   = 8
	GAT_QuestMiniSecond                  = 9
	GAT_QuestMiniThird                   = 10
	GAT_MainCrystal                      = 11
	GAT_MainCastle                       = 12
	GAT_MainJJC                          = 13
	GAT_OfflineJJC                       = 14
	GAT_OfflineJJC4                      = 15
	GAT_WorldMapER                       = 16
	GAT_WorldMapFL                       = 17
	GAT_WorldMapWorldBtn                 = 18
	GAT_MiniMap                          = 19
	GAT_TeamCreateBtn                    = 20
	GAT_TeamWorldMapBtn                  = 21
	GAT_FirstPartner_PosUI               = 22
	GAT_FreeLootPartner                  = 23
	GAT_FriendAddBtn                     = 24
	GAT_PartnerShowCancel                = 25
	GAT_PartnerPositionTab               = 26
	GAT_PartnerDetailBaseTab             = 27
	GAT_PartnerDetailBodySlot            = 28
	GAT_PartnerDetailEquipBtn            = 29
	GAT_PartnerDetailBaseFirstSkill      = 30
	GAT_PartnerDetailBaseSkillLvUpBtn    = 31
	GAT_LearnSkillAttackSkillTab         = 32
	GAT_LearnSkillAttackMagicTab         = 33
	GAT_LearnSkillStatusMagicTab         = 34
	GAT_LearnSkillAidSkillTab            = 35
	GAT_LearnSkillBtn                    = 36
	GAT_FirstLearningBabySkill           = 37
	GAT_FirstLearningBabySkill_BabyList  = 38
	GAT_ThirdLearningBabySkill_SkillSlot = 39
	GAT_BabySkillLearningBtn             = 40
	GAT_MessageBoxOkBtn                  = 41
	GAT_MainReturn                       = 42
	GAT_MainBag                          = 43
	GAT_MainBagTipUseItem                = 44
	GAT_MainbagTipEquip                  = 45
	GAT_MainSkill                        = 46
	GAT_MainMake                         = 47
	GAT_MainMakeSword                    = 48
	GAT_MainMakeAxe                      = 49
	GAT_MainMakeStick                    = 50
	GAT_MainMakeBow                      = 51
	GAT_MainMakeCompoundBtn              = 52
	GAT_MainMakeLevel10                  = 53
	GAT_MainMakeSubFirst                 = 54
	GAT_MainMakeSubSecond                = 55
	GAT_MainMakeSubThird                 = 56
	GAT_MainMakeGemBtn                   = 57
	GAT_MainMakeGemClose                 = 58
	GAT_MainMakeGemFirst                 = 59
	GAT_MainBaby                         = 60
	GAT_MainBabyStatusBtn                = 61
	GAT_MainBabyPropertyBtn              = 62
	GAT_MainBabyPropertyContainer        = 63
	GAT_MainBabyPropertyConfirm          = 64
	GAT_MainBabyClose                    = 65
	GAT_MainMagic                        = 66
	GAT_MainMagicLevelFirst              = 67
	GAT_MainMagicLevelBtn                = 68
	GAT_MainFriend                       = 69
	GAT_MainPartner                      = 70
	GAT_MainSetting                      = 71
	GAT_MainRecharge                     = 72
	GAT_MainActivity                     = 73
	GAT_MainPlayerInfo                   = 74
	GAT_MainPlayerInfoStatusBtn          = 75
	GAT_MainPlayerInfoPropertyBtn        = 76
	GAT_MainPlayerInfoPropertyContainer  = 77
	GAT_MainPlayerInfoPropertyConfirm    = 78
	GAT_MainPlayerInfoClose              = 79
	GAT_MainJiubaHouse                   = 80
	GAT_MainStick                        = 81
	GAT_MainAchievement                  = 82
	GAT_MainRaise                        = 83
	GAT_MainFamily                       = 84
	GAT_BattleAttack                     = 85
	GAT_BattleSkill                      = 86
	GAT_BattleBabySkill                  = 87
	GAT_BattleDeference                  = 88
	GAT_BattleChangePosition             = 89
	GAT_BattleAuto                       = 90
	GAT_BattleBag                        = 91
	GAT_BattleCatch                      = 92
	GAT_BattleBaby                       = 93
	GAT_BattleRunaway                    = 94
	GAT_BattlePlayerInfo                 = 95
	GAT_BattleRewardClose                = 96
	GAT_Max                              = 97
)

func ToName_GuideAimType(id int) string {
	switch id {
	case GAT_None:
		return "GAT_None"
	case GAT_FirstAchievement:
		return "GAT_FirstAchievement"
	case GAT_FirstSkill:
		return "GAT_FirstSkill"
	case GAT_FirstLevelSkill:
		return "GAT_FirstLevelSkill"
	case GAT_FirstQuest:
		return "GAT_FirstQuest"
	case GAT_DialogUI:
		return "GAT_DialogUI"
	case GAT_MainTeamBtn:
		return "GAT_MainTeamBtn"
	case GAT_MainTaskBtn:
		return "GAT_MainTaskBtn"
	case GAT_QuestMiniFirst:
		return "GAT_QuestMiniFirst"
	case GAT_QuestMiniSecond:
		return "GAT_QuestMiniSecond"
	case GAT_QuestMiniThird:
		return "GAT_QuestMiniThird"
	case GAT_MainCrystal:
		return "GAT_MainCrystal"
	case GAT_MainCastle:
		return "GAT_MainCastle"
	case GAT_MainJJC:
		return "GAT_MainJJC"
	case GAT_OfflineJJC:
		return "GAT_OfflineJJC"
	case GAT_OfflineJJC4:
		return "GAT_OfflineJJC4"
	case GAT_WorldMapER:
		return "GAT_WorldMapER"
	case GAT_WorldMapFL:
		return "GAT_WorldMapFL"
	case GAT_WorldMapWorldBtn:
		return "GAT_WorldMapWorldBtn"
	case GAT_MiniMap:
		return "GAT_MiniMap"
	case GAT_TeamCreateBtn:
		return "GAT_TeamCreateBtn"
	case GAT_TeamWorldMapBtn:
		return "GAT_TeamWorldMapBtn"
	case GAT_FirstPartner_PosUI:
		return "GAT_FirstPartner_PosUI"
	case GAT_FreeLootPartner:
		return "GAT_FreeLootPartner"
	case GAT_FriendAddBtn:
		return "GAT_FriendAddBtn"
	case GAT_PartnerShowCancel:
		return "GAT_PartnerShowCancel"
	case GAT_PartnerPositionTab:
		return "GAT_PartnerPositionTab"
	case GAT_PartnerDetailBaseTab:
		return "GAT_PartnerDetailBaseTab"
	case GAT_PartnerDetailBodySlot:
		return "GAT_PartnerDetailBodySlot"
	case GAT_PartnerDetailEquipBtn:
		return "GAT_PartnerDetailEquipBtn"
	case GAT_PartnerDetailBaseFirstSkill:
		return "GAT_PartnerDetailBaseFirstSkill"
	case GAT_PartnerDetailBaseSkillLvUpBtn:
		return "GAT_PartnerDetailBaseSkillLvUpBtn"
	case GAT_LearnSkillAttackSkillTab:
		return "GAT_LearnSkillAttackSkillTab"
	case GAT_LearnSkillAttackMagicTab:
		return "GAT_LearnSkillAttackMagicTab"
	case GAT_LearnSkillStatusMagicTab:
		return "GAT_LearnSkillStatusMagicTab"
	case GAT_LearnSkillAidSkillTab:
		return "GAT_LearnSkillAidSkillTab"
	case GAT_LearnSkillBtn:
		return "GAT_LearnSkillBtn"
	case GAT_FirstLearningBabySkill:
		return "GAT_FirstLearningBabySkill"
	case GAT_FirstLearningBabySkill_BabyList:
		return "GAT_FirstLearningBabySkill_BabyList"
	case GAT_ThirdLearningBabySkill_SkillSlot:
		return "GAT_ThirdLearningBabySkill_SkillSlot"
	case GAT_BabySkillLearningBtn:
		return "GAT_BabySkillLearningBtn"
	case GAT_MessageBoxOkBtn:
		return "GAT_MessageBoxOkBtn"
	case GAT_MainReturn:
		return "GAT_MainReturn"
	case GAT_MainBag:
		return "GAT_MainBag"
	case GAT_MainBagTipUseItem:
		return "GAT_MainBagTipUseItem"
	case GAT_MainbagTipEquip:
		return "GAT_MainbagTipEquip"
	case GAT_MainSkill:
		return "GAT_MainSkill"
	case GAT_MainMake:
		return "GAT_MainMake"
	case GAT_MainMakeSword:
		return "GAT_MainMakeSword"
	case GAT_MainMakeAxe:
		return "GAT_MainMakeAxe"
	case GAT_MainMakeStick:
		return "GAT_MainMakeStick"
	case GAT_MainMakeBow:
		return "GAT_MainMakeBow"
	case GAT_MainMakeCompoundBtn:
		return "GAT_MainMakeCompoundBtn"
	case GAT_MainMakeLevel10:
		return "GAT_MainMakeLevel10"
	case GAT_MainMakeSubFirst:
		return "GAT_MainMakeSubFirst"
	case GAT_MainMakeSubSecond:
		return "GAT_MainMakeSubSecond"
	case GAT_MainMakeSubThird:
		return "GAT_MainMakeSubThird"
	case GAT_MainMakeGemBtn:
		return "GAT_MainMakeGemBtn"
	case GAT_MainMakeGemClose:
		return "GAT_MainMakeGemClose"
	case GAT_MainMakeGemFirst:
		return "GAT_MainMakeGemFirst"
	case GAT_MainBaby:
		return "GAT_MainBaby"
	case GAT_MainBabyStatusBtn:
		return "GAT_MainBabyStatusBtn"
	case GAT_MainBabyPropertyBtn:
		return "GAT_MainBabyPropertyBtn"
	case GAT_MainBabyPropertyContainer:
		return "GAT_MainBabyPropertyContainer"
	case GAT_MainBabyPropertyConfirm:
		return "GAT_MainBabyPropertyConfirm"
	case GAT_MainBabyClose:
		return "GAT_MainBabyClose"
	case GAT_MainMagic:
		return "GAT_MainMagic"
	case GAT_MainMagicLevelFirst:
		return "GAT_MainMagicLevelFirst"
	case GAT_MainMagicLevelBtn:
		return "GAT_MainMagicLevelBtn"
	case GAT_MainFriend:
		return "GAT_MainFriend"
	case GAT_MainPartner:
		return "GAT_MainPartner"
	case GAT_MainSetting:
		return "GAT_MainSetting"
	case GAT_MainRecharge:
		return "GAT_MainRecharge"
	case GAT_MainActivity:
		return "GAT_MainActivity"
	case GAT_MainPlayerInfo:
		return "GAT_MainPlayerInfo"
	case GAT_MainPlayerInfoStatusBtn:
		return "GAT_MainPlayerInfoStatusBtn"
	case GAT_MainPlayerInfoPropertyBtn:
		return "GAT_MainPlayerInfoPropertyBtn"
	case GAT_MainPlayerInfoPropertyContainer:
		return "GAT_MainPlayerInfoPropertyContainer"
	case GAT_MainPlayerInfoPropertyConfirm:
		return "GAT_MainPlayerInfoPropertyConfirm"
	case GAT_MainPlayerInfoClose:
		return "GAT_MainPlayerInfoClose"
	case GAT_MainJiubaHouse:
		return "GAT_MainJiubaHouse"
	case GAT_MainStick:
		return "GAT_MainStick"
	case GAT_MainAchievement:
		return "GAT_MainAchievement"
	case GAT_MainRaise:
		return "GAT_MainRaise"
	case GAT_MainFamily:
		return "GAT_MainFamily"
	case GAT_BattleAttack:
		return "GAT_BattleAttack"
	case GAT_BattleSkill:
		return "GAT_BattleSkill"
	case GAT_BattleBabySkill:
		return "GAT_BattleBabySkill"
	case GAT_BattleDeference:
		return "GAT_BattleDeference"
	case GAT_BattleChangePosition:
		return "GAT_BattleChangePosition"
	case GAT_BattleAuto:
		return "GAT_BattleAuto"
	case GAT_BattleBag:
		return "GAT_BattleBag"
	case GAT_BattleCatch:
		return "GAT_BattleCatch"
	case GAT_BattleBaby:
		return "GAT_BattleBaby"
	case GAT_BattleRunaway:
		return "GAT_BattleRunaway"
	case GAT_BattlePlayerInfo:
		return "GAT_BattlePlayerInfo"
	case GAT_BattleRewardClose:
		return "GAT_BattleRewardClose"
	case GAT_Max:
		return "GAT_Max"
	default:
		return ""
	}
}
func ToId_GuideAimType(name string) int {
	switch name {
	case "GAT_None":
		return GAT_None
	case "GAT_FirstAchievement":
		return GAT_FirstAchievement
	case "GAT_FirstSkill":
		return GAT_FirstSkill
	case "GAT_FirstLevelSkill":
		return GAT_FirstLevelSkill
	case "GAT_FirstQuest":
		return GAT_FirstQuest
	case "GAT_DialogUI":
		return GAT_DialogUI
	case "GAT_MainTeamBtn":
		return GAT_MainTeamBtn
	case "GAT_MainTaskBtn":
		return GAT_MainTaskBtn
	case "GAT_QuestMiniFirst":
		return GAT_QuestMiniFirst
	case "GAT_QuestMiniSecond":
		return GAT_QuestMiniSecond
	case "GAT_QuestMiniThird":
		return GAT_QuestMiniThird
	case "GAT_MainCrystal":
		return GAT_MainCrystal
	case "GAT_MainCastle":
		return GAT_MainCastle
	case "GAT_MainJJC":
		return GAT_MainJJC
	case "GAT_OfflineJJC":
		return GAT_OfflineJJC
	case "GAT_OfflineJJC4":
		return GAT_OfflineJJC4
	case "GAT_WorldMapER":
		return GAT_WorldMapER
	case "GAT_WorldMapFL":
		return GAT_WorldMapFL
	case "GAT_WorldMapWorldBtn":
		return GAT_WorldMapWorldBtn
	case "GAT_MiniMap":
		return GAT_MiniMap
	case "GAT_TeamCreateBtn":
		return GAT_TeamCreateBtn
	case "GAT_TeamWorldMapBtn":
		return GAT_TeamWorldMapBtn
	case "GAT_FirstPartner_PosUI":
		return GAT_FirstPartner_PosUI
	case "GAT_FreeLootPartner":
		return GAT_FreeLootPartner
	case "GAT_FriendAddBtn":
		return GAT_FriendAddBtn
	case "GAT_PartnerShowCancel":
		return GAT_PartnerShowCancel
	case "GAT_PartnerPositionTab":
		return GAT_PartnerPositionTab
	case "GAT_PartnerDetailBaseTab":
		return GAT_PartnerDetailBaseTab
	case "GAT_PartnerDetailBodySlot":
		return GAT_PartnerDetailBodySlot
	case "GAT_PartnerDetailEquipBtn":
		return GAT_PartnerDetailEquipBtn
	case "GAT_PartnerDetailBaseFirstSkill":
		return GAT_PartnerDetailBaseFirstSkill
	case "GAT_PartnerDetailBaseSkillLvUpBtn":
		return GAT_PartnerDetailBaseSkillLvUpBtn
	case "GAT_LearnSkillAttackSkillTab":
		return GAT_LearnSkillAttackSkillTab
	case "GAT_LearnSkillAttackMagicTab":
		return GAT_LearnSkillAttackMagicTab
	case "GAT_LearnSkillStatusMagicTab":
		return GAT_LearnSkillStatusMagicTab
	case "GAT_LearnSkillAidSkillTab":
		return GAT_LearnSkillAidSkillTab
	case "GAT_LearnSkillBtn":
		return GAT_LearnSkillBtn
	case "GAT_FirstLearningBabySkill":
		return GAT_FirstLearningBabySkill
	case "GAT_FirstLearningBabySkill_BabyList":
		return GAT_FirstLearningBabySkill_BabyList
	case "GAT_ThirdLearningBabySkill_SkillSlot":
		return GAT_ThirdLearningBabySkill_SkillSlot
	case "GAT_BabySkillLearningBtn":
		return GAT_BabySkillLearningBtn
	case "GAT_MessageBoxOkBtn":
		return GAT_MessageBoxOkBtn
	case "GAT_MainReturn":
		return GAT_MainReturn
	case "GAT_MainBag":
		return GAT_MainBag
	case "GAT_MainBagTipUseItem":
		return GAT_MainBagTipUseItem
	case "GAT_MainbagTipEquip":
		return GAT_MainbagTipEquip
	case "GAT_MainSkill":
		return GAT_MainSkill
	case "GAT_MainMake":
		return GAT_MainMake
	case "GAT_MainMakeSword":
		return GAT_MainMakeSword
	case "GAT_MainMakeAxe":
		return GAT_MainMakeAxe
	case "GAT_MainMakeStick":
		return GAT_MainMakeStick
	case "GAT_MainMakeBow":
		return GAT_MainMakeBow
	case "GAT_MainMakeCompoundBtn":
		return GAT_MainMakeCompoundBtn
	case "GAT_MainMakeLevel10":
		return GAT_MainMakeLevel10
	case "GAT_MainMakeSubFirst":
		return GAT_MainMakeSubFirst
	case "GAT_MainMakeSubSecond":
		return GAT_MainMakeSubSecond
	case "GAT_MainMakeSubThird":
		return GAT_MainMakeSubThird
	case "GAT_MainMakeGemBtn":
		return GAT_MainMakeGemBtn
	case "GAT_MainMakeGemClose":
		return GAT_MainMakeGemClose
	case "GAT_MainMakeGemFirst":
		return GAT_MainMakeGemFirst
	case "GAT_MainBaby":
		return GAT_MainBaby
	case "GAT_MainBabyStatusBtn":
		return GAT_MainBabyStatusBtn
	case "GAT_MainBabyPropertyBtn":
		return GAT_MainBabyPropertyBtn
	case "GAT_MainBabyPropertyContainer":
		return GAT_MainBabyPropertyContainer
	case "GAT_MainBabyPropertyConfirm":
		return GAT_MainBabyPropertyConfirm
	case "GAT_MainBabyClose":
		return GAT_MainBabyClose
	case "GAT_MainMagic":
		return GAT_MainMagic
	case "GAT_MainMagicLevelFirst":
		return GAT_MainMagicLevelFirst
	case "GAT_MainMagicLevelBtn":
		return GAT_MainMagicLevelBtn
	case "GAT_MainFriend":
		return GAT_MainFriend
	case "GAT_MainPartner":
		return GAT_MainPartner
	case "GAT_MainSetting":
		return GAT_MainSetting
	case "GAT_MainRecharge":
		return GAT_MainRecharge
	case "GAT_MainActivity":
		return GAT_MainActivity
	case "GAT_MainPlayerInfo":
		return GAT_MainPlayerInfo
	case "GAT_MainPlayerInfoStatusBtn":
		return GAT_MainPlayerInfoStatusBtn
	case "GAT_MainPlayerInfoPropertyBtn":
		return GAT_MainPlayerInfoPropertyBtn
	case "GAT_MainPlayerInfoPropertyContainer":
		return GAT_MainPlayerInfoPropertyContainer
	case "GAT_MainPlayerInfoPropertyConfirm":
		return GAT_MainPlayerInfoPropertyConfirm
	case "GAT_MainPlayerInfoClose":
		return GAT_MainPlayerInfoClose
	case "GAT_MainJiubaHouse":
		return GAT_MainJiubaHouse
	case "GAT_MainStick":
		return GAT_MainStick
	case "GAT_MainAchievement":
		return GAT_MainAchievement
	case "GAT_MainRaise":
		return GAT_MainRaise
	case "GAT_MainFamily":
		return GAT_MainFamily
	case "GAT_BattleAttack":
		return GAT_BattleAttack
	case "GAT_BattleSkill":
		return GAT_BattleSkill
	case "GAT_BattleBabySkill":
		return GAT_BattleBabySkill
	case "GAT_BattleDeference":
		return GAT_BattleDeference
	case "GAT_BattleChangePosition":
		return GAT_BattleChangePosition
	case "GAT_BattleAuto":
		return GAT_BattleAuto
	case "GAT_BattleBag":
		return GAT_BattleBag
	case "GAT_BattleCatch":
		return GAT_BattleCatch
	case "GAT_BattleBaby":
		return GAT_BattleBaby
	case "GAT_BattleRunaway":
		return GAT_BattleRunaway
	case "GAT_BattlePlayerInfo":
		return GAT_BattlePlayerInfo
	case "GAT_BattleRewardClose":
		return GAT_BattleRewardClose
	case "GAT_Max":
		return GAT_Max
	default:
		return -1
	}
}
