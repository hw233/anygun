package prpc

// enum PropertyType
const (
	PT_None                = 0
	PT_NoSleep             = 1
	PT_NoPetrifaction      = 2
	PT_NoDrunk             = 3
	PT_NoChaos             = 4
	PT_NoForget            = 5
	PT_NoPoison            = 6
	PT_Assassinate         = 7
	PT_Money               = 8
	PT_Diamond             = 9
	PT_MagicCurrency       = 10
	PT_EmployeeCurrency    = 11
	PT_Gather              = 12
	PT_Level               = 13
	PT_Exp                 = 14
	PT_ConvertExp          = 15
	PT_OneDayReputation    = 16
	PT_Reputation          = 17
	PT_TableId             = 18
	PT_AssetId             = 19
	PT_Sex                 = 20
	PT_BagNum              = 21
	PT_Race                = 22
	PT_Profession          = 23
	PT_ProfessionLevel     = 24
	PT_Stama               = 25
	PT_Strength            = 26
	PT_Power               = 27
	PT_Speed               = 28
	PT_Magic               = 29
	PT_Durability          = 30
	PT_HpCurr              = 31
	PT_MpCurr              = 32
	PT_HpMax               = 33
	PT_MpMax               = 34
	PT_Attack              = 35
	PT_Defense             = 36
	PT_Agile               = 37
	PT_Spirit              = 38
	PT_Reply               = 39
	PT_Magicattack         = 40
	PT_Magicdefense        = 41
	PT_Damage              = 42
	PT_SneakAttack         = 43
	PT_Hit                 = 44
	PT_Dodge               = 45
	PT_Crit                = 46
	PT_counterpunch        = 47
	PT_Front               = 48
	PT_Wind                = 49
	PT_Land                = 50
	PT_Water               = 51
	PT_Fire                = 52
	PT_Free                = 53
	PT_Title               = 54
	PT_GuildID             = 55
	PT_Glamour             = 56
	PT_DoubleExp           = 57
	PT_TongjiComplateTimes = 58
	PT_VipLevel            = 59
	PT_VipTime             = 60
	PT_FightingForce       = 61
	PT_Max                 = 62
)

func ToName_PropertyType(id int) string {
	switch id {
	case PT_None:
		return "PT_None"
	case PT_NoSleep:
		return "PT_NoSleep"
	case PT_NoPetrifaction:
		return "PT_NoPetrifaction"
	case PT_NoDrunk:
		return "PT_NoDrunk"
	case PT_NoChaos:
		return "PT_NoChaos"
	case PT_NoForget:
		return "PT_NoForget"
	case PT_NoPoison:
		return "PT_NoPoison"
	case PT_Assassinate:
		return "PT_Assassinate"
	case PT_Money:
		return "PT_Money"
	case PT_Diamond:
		return "PT_Diamond"
	case PT_MagicCurrency:
		return "PT_MagicCurrency"
	case PT_EmployeeCurrency:
		return "PT_EmployeeCurrency"
	case PT_Gather:
		return "PT_Gather"
	case PT_Level:
		return "PT_Level"
	case PT_Exp:
		return "PT_Exp"
	case PT_ConvertExp:
		return "PT_ConvertExp"
	case PT_OneDayReputation:
		return "PT_OneDayReputation"
	case PT_Reputation:
		return "PT_Reputation"
	case PT_TableId:
		return "PT_TableId"
	case PT_AssetId:
		return "PT_AssetId"
	case PT_Sex:
		return "PT_Sex"
	case PT_BagNum:
		return "PT_BagNum"
	case PT_Race:
		return "PT_Race"
	case PT_Profession:
		return "PT_Profession"
	case PT_ProfessionLevel:
		return "PT_ProfessionLevel"
	case PT_Stama:
		return "PT_Stama"
	case PT_Strength:
		return "PT_Strength"
	case PT_Power:
		return "PT_Power"
	case PT_Speed:
		return "PT_Speed"
	case PT_Magic:
		return "PT_Magic"
	case PT_Durability:
		return "PT_Durability"
	case PT_HpCurr:
		return "PT_HpCurr"
	case PT_MpCurr:
		return "PT_MpCurr"
	case PT_HpMax:
		return "PT_HpMax"
	case PT_MpMax:
		return "PT_MpMax"
	case PT_Attack:
		return "PT_Attack"
	case PT_Defense:
		return "PT_Defense"
	case PT_Agile:
		return "PT_Agile"
	case PT_Spirit:
		return "PT_Spirit"
	case PT_Reply:
		return "PT_Reply"
	case PT_Magicattack:
		return "PT_Magicattack"
	case PT_Magicdefense:
		return "PT_Magicdefense"
	case PT_Damage:
		return "PT_Damage"
	case PT_SneakAttack:
		return "PT_SneakAttack"
	case PT_Hit:
		return "PT_Hit"
	case PT_Dodge:
		return "PT_Dodge"
	case PT_Crit:
		return "PT_Crit"
	case PT_counterpunch:
		return "PT_counterpunch"
	case PT_Front:
		return "PT_Front"
	case PT_Wind:
		return "PT_Wind"
	case PT_Land:
		return "PT_Land"
	case PT_Water:
		return "PT_Water"
	case PT_Fire:
		return "PT_Fire"
	case PT_Free:
		return "PT_Free"
	case PT_Title:
		return "PT_Title"
	case PT_GuildID:
		return "PT_GuildID"
	case PT_Glamour:
		return "PT_Glamour"
	case PT_DoubleExp:
		return "PT_DoubleExp"
	case PT_TongjiComplateTimes:
		return "PT_TongjiComplateTimes"
	case PT_VipLevel:
		return "PT_VipLevel"
	case PT_VipTime:
		return "PT_VipTime"
	case PT_FightingForce:
		return "PT_FightingForce"
	case PT_Max:
		return "PT_Max"
	default:
		return ""
	}
}
func ToId_PropertyType(name string) int {
	switch name {
	case "PT_None":
		return PT_None
	case "PT_NoSleep":
		return PT_NoSleep
	case "PT_NoPetrifaction":
		return PT_NoPetrifaction
	case "PT_NoDrunk":
		return PT_NoDrunk
	case "PT_NoChaos":
		return PT_NoChaos
	case "PT_NoForget":
		return PT_NoForget
	case "PT_NoPoison":
		return PT_NoPoison
	case "PT_Assassinate":
		return PT_Assassinate
	case "PT_Money":
		return PT_Money
	case "PT_Diamond":
		return PT_Diamond
	case "PT_MagicCurrency":
		return PT_MagicCurrency
	case "PT_EmployeeCurrency":
		return PT_EmployeeCurrency
	case "PT_Gather":
		return PT_Gather
	case "PT_Level":
		return PT_Level
	case "PT_Exp":
		return PT_Exp
	case "PT_ConvertExp":
		return PT_ConvertExp
	case "PT_OneDayReputation":
		return PT_OneDayReputation
	case "PT_Reputation":
		return PT_Reputation
	case "PT_TableId":
		return PT_TableId
	case "PT_AssetId":
		return PT_AssetId
	case "PT_Sex":
		return PT_Sex
	case "PT_BagNum":
		return PT_BagNum
	case "PT_Race":
		return PT_Race
	case "PT_Profession":
		return PT_Profession
	case "PT_ProfessionLevel":
		return PT_ProfessionLevel
	case "PT_Stama":
		return PT_Stama
	case "PT_Strength":
		return PT_Strength
	case "PT_Power":
		return PT_Power
	case "PT_Speed":
		return PT_Speed
	case "PT_Magic":
		return PT_Magic
	case "PT_Durability":
		return PT_Durability
	case "PT_HpCurr":
		return PT_HpCurr
	case "PT_MpCurr":
		return PT_MpCurr
	case "PT_HpMax":
		return PT_HpMax
	case "PT_MpMax":
		return PT_MpMax
	case "PT_Attack":
		return PT_Attack
	case "PT_Defense":
		return PT_Defense
	case "PT_Agile":
		return PT_Agile
	case "PT_Spirit":
		return PT_Spirit
	case "PT_Reply":
		return PT_Reply
	case "PT_Magicattack":
		return PT_Magicattack
	case "PT_Magicdefense":
		return PT_Magicdefense
	case "PT_Damage":
		return PT_Damage
	case "PT_SneakAttack":
		return PT_SneakAttack
	case "PT_Hit":
		return PT_Hit
	case "PT_Dodge":
		return PT_Dodge
	case "PT_Crit":
		return PT_Crit
	case "PT_counterpunch":
		return PT_counterpunch
	case "PT_Front":
		return PT_Front
	case "PT_Wind":
		return PT_Wind
	case "PT_Land":
		return PT_Land
	case "PT_Water":
		return PT_Water
	case "PT_Fire":
		return PT_Fire
	case "PT_Free":
		return PT_Free
	case "PT_Title":
		return PT_Title
	case "PT_GuildID":
		return PT_GuildID
	case "PT_Glamour":
		return PT_Glamour
	case "PT_DoubleExp":
		return PT_DoubleExp
	case "PT_TongjiComplateTimes":
		return PT_TongjiComplateTimes
	case "PT_VipLevel":
		return PT_VipLevel
	case "PT_VipTime":
		return PT_VipTime
	case "PT_FightingForce":
		return PT_FightingForce
	case "PT_Max":
		return PT_Max
	default:
		return -1
	}
}
