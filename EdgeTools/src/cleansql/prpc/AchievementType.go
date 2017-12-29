package prpc

// enum AchievementType
const (
	AT_None                   = 0
	AT_EarnConis              = 1
	AT_SpendMoney             = 2
	AT_SpendDiamond           = 3
	AT_Recharge               = 4
	AT_RoleLevel              = 5
	AT_PetLevel               = 6
	AT_AttackLevel            = 7
	AT_DefenseLevel           = 8
	AT_AgileLevel             = 9
	AT_WearCrystal            = 10
	AT_WearAccessories        = 11
	AT_TotalDamage            = 12
	AT_TotalTreatment         = 13
	AT_HasSkillNum            = 14
	AT_BabySkill              = 15
	AT_CatchPet               = 16
	AT_RecruitPartner         = 17
	AT_PartnerCard            = 18
	AT_PartnersUpgradeGreen   = 19
	AT_PartnersUpgradeBlue    = 20
	AT_PartnersUpgradePurple  = 21
	AT_PartnersUpgradeGold    = 22
	AT_PartnersUpgradeOrangle = 23
	AT_PartnersUpgradePink    = 24
	AT_QualifyingAdvanced     = 25
	AT_ArenaWin               = 26
	AT_KillMonster            = 27
	AT_KillBoss               = 28
	AT_KillPlayer             = 29
	AT_MakeEquipment          = 30
	AT_Reward50               = 31
	AT_EverydayActivities     = 32
	AT_Sign                   = 33
	AT_Wanted                 = 34
	AT_Copy30                 = 35
	AT_Copy40                 = 36
	AT_Blood                  = 37
	AT_Magic                  = 38
	AT_Bag                    = 39
	AT_PetBag                 = 40
	AT_GoodMake               = 41
	AT_PetIntensive           = 42
	AT_PetHuman               = 43
	AT_PetInsect              = 44
	AT_PetPlant               = 45
	AT_PetExtra               = 46
	AT_PetDragon              = 47
	AT_PetAnimal              = 48
	AT_PetFly                 = 49
	AT_PetUndead              = 50
	AT_PetMetal               = 51
	AT_Home                   = 52
	AT_CollectMaterial        = 53
	AT_Friend                 = 54
	AT_Billboard              = 55
	AT_OwnConis               = 56
	AT_MagicEquip             = 57
	AT_Max                    = 58
)

func ToName_AchievementType(id int) string {
	switch id {
	case AT_None:
		return "AT_None"
	case AT_EarnConis:
		return "AT_EarnConis"
	case AT_SpendMoney:
		return "AT_SpendMoney"
	case AT_SpendDiamond:
		return "AT_SpendDiamond"
	case AT_Recharge:
		return "AT_Recharge"
	case AT_RoleLevel:
		return "AT_RoleLevel"
	case AT_PetLevel:
		return "AT_PetLevel"
	case AT_AttackLevel:
		return "AT_AttackLevel"
	case AT_DefenseLevel:
		return "AT_DefenseLevel"
	case AT_AgileLevel:
		return "AT_AgileLevel"
	case AT_WearCrystal:
		return "AT_WearCrystal"
	case AT_WearAccessories:
		return "AT_WearAccessories"
	case AT_TotalDamage:
		return "AT_TotalDamage"
	case AT_TotalTreatment:
		return "AT_TotalTreatment"
	case AT_HasSkillNum:
		return "AT_HasSkillNum"
	case AT_BabySkill:
		return "AT_BabySkill"
	case AT_CatchPet:
		return "AT_CatchPet"
	case AT_RecruitPartner:
		return "AT_RecruitPartner"
	case AT_PartnerCard:
		return "AT_PartnerCard"
	case AT_PartnersUpgradeGreen:
		return "AT_PartnersUpgradeGreen"
	case AT_PartnersUpgradeBlue:
		return "AT_PartnersUpgradeBlue"
	case AT_PartnersUpgradePurple:
		return "AT_PartnersUpgradePurple"
	case AT_PartnersUpgradeGold:
		return "AT_PartnersUpgradeGold"
	case AT_PartnersUpgradeOrangle:
		return "AT_PartnersUpgradeOrangle"
	case AT_PartnersUpgradePink:
		return "AT_PartnersUpgradePink"
	case AT_QualifyingAdvanced:
		return "AT_QualifyingAdvanced"
	case AT_ArenaWin:
		return "AT_ArenaWin"
	case AT_KillMonster:
		return "AT_KillMonster"
	case AT_KillBoss:
		return "AT_KillBoss"
	case AT_KillPlayer:
		return "AT_KillPlayer"
	case AT_MakeEquipment:
		return "AT_MakeEquipment"
	case AT_Reward50:
		return "AT_Reward50"
	case AT_EverydayActivities:
		return "AT_EverydayActivities"
	case AT_Sign:
		return "AT_Sign"
	case AT_Wanted:
		return "AT_Wanted"
	case AT_Copy30:
		return "AT_Copy30"
	case AT_Copy40:
		return "AT_Copy40"
	case AT_Blood:
		return "AT_Blood"
	case AT_Magic:
		return "AT_Magic"
	case AT_Bag:
		return "AT_Bag"
	case AT_PetBag:
		return "AT_PetBag"
	case AT_GoodMake:
		return "AT_GoodMake"
	case AT_PetIntensive:
		return "AT_PetIntensive"
	case AT_PetHuman:
		return "AT_PetHuman"
	case AT_PetInsect:
		return "AT_PetInsect"
	case AT_PetPlant:
		return "AT_PetPlant"
	case AT_PetExtra:
		return "AT_PetExtra"
	case AT_PetDragon:
		return "AT_PetDragon"
	case AT_PetAnimal:
		return "AT_PetAnimal"
	case AT_PetFly:
		return "AT_PetFly"
	case AT_PetUndead:
		return "AT_PetUndead"
	case AT_PetMetal:
		return "AT_PetMetal"
	case AT_Home:
		return "AT_Home"
	case AT_CollectMaterial:
		return "AT_CollectMaterial"
	case AT_Friend:
		return "AT_Friend"
	case AT_Billboard:
		return "AT_Billboard"
	case AT_OwnConis:
		return "AT_OwnConis"
	case AT_MagicEquip:
		return "AT_MagicEquip"
	case AT_Max:
		return "AT_Max"
	default:
		return ""
	}
}
func ToId_AchievementType(name string) int {
	switch name {
	case "AT_None":
		return AT_None
	case "AT_EarnConis":
		return AT_EarnConis
	case "AT_SpendMoney":
		return AT_SpendMoney
	case "AT_SpendDiamond":
		return AT_SpendDiamond
	case "AT_Recharge":
		return AT_Recharge
	case "AT_RoleLevel":
		return AT_RoleLevel
	case "AT_PetLevel":
		return AT_PetLevel
	case "AT_AttackLevel":
		return AT_AttackLevel
	case "AT_DefenseLevel":
		return AT_DefenseLevel
	case "AT_AgileLevel":
		return AT_AgileLevel
	case "AT_WearCrystal":
		return AT_WearCrystal
	case "AT_WearAccessories":
		return AT_WearAccessories
	case "AT_TotalDamage":
		return AT_TotalDamage
	case "AT_TotalTreatment":
		return AT_TotalTreatment
	case "AT_HasSkillNum":
		return AT_HasSkillNum
	case "AT_BabySkill":
		return AT_BabySkill
	case "AT_CatchPet":
		return AT_CatchPet
	case "AT_RecruitPartner":
		return AT_RecruitPartner
	case "AT_PartnerCard":
		return AT_PartnerCard
	case "AT_PartnersUpgradeGreen":
		return AT_PartnersUpgradeGreen
	case "AT_PartnersUpgradeBlue":
		return AT_PartnersUpgradeBlue
	case "AT_PartnersUpgradePurple":
		return AT_PartnersUpgradePurple
	case "AT_PartnersUpgradeGold":
		return AT_PartnersUpgradeGold
	case "AT_PartnersUpgradeOrangle":
		return AT_PartnersUpgradeOrangle
	case "AT_PartnersUpgradePink":
		return AT_PartnersUpgradePink
	case "AT_QualifyingAdvanced":
		return AT_QualifyingAdvanced
	case "AT_ArenaWin":
		return AT_ArenaWin
	case "AT_KillMonster":
		return AT_KillMonster
	case "AT_KillBoss":
		return AT_KillBoss
	case "AT_KillPlayer":
		return AT_KillPlayer
	case "AT_MakeEquipment":
		return AT_MakeEquipment
	case "AT_Reward50":
		return AT_Reward50
	case "AT_EverydayActivities":
		return AT_EverydayActivities
	case "AT_Sign":
		return AT_Sign
	case "AT_Wanted":
		return AT_Wanted
	case "AT_Copy30":
		return AT_Copy30
	case "AT_Copy40":
		return AT_Copy40
	case "AT_Blood":
		return AT_Blood
	case "AT_Magic":
		return AT_Magic
	case "AT_Bag":
		return AT_Bag
	case "AT_PetBag":
		return AT_PetBag
	case "AT_GoodMake":
		return AT_GoodMake
	case "AT_PetIntensive":
		return AT_PetIntensive
	case "AT_PetHuman":
		return AT_PetHuman
	case "AT_PetInsect":
		return AT_PetInsect
	case "AT_PetPlant":
		return AT_PetPlant
	case "AT_PetExtra":
		return AT_PetExtra
	case "AT_PetDragon":
		return AT_PetDragon
	case "AT_PetAnimal":
		return AT_PetAnimal
	case "AT_PetFly":
		return AT_PetFly
	case "AT_PetUndead":
		return AT_PetUndead
	case "AT_PetMetal":
		return AT_PetMetal
	case "AT_Home":
		return AT_Home
	case "AT_CollectMaterial":
		return AT_CollectMaterial
	case "AT_Friend":
		return AT_Friend
	case "AT_Billboard":
		return AT_Billboard
	case "AT_OwnConis":
		return AT_OwnConis
	case "AT_MagicEquip":
		return AT_MagicEquip
	case "AT_Max":
		return AT_Max
	default:
		return -1
	}
}
