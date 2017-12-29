package prpc

// enum HelpRaiseType
const (
	HRT_None            = 0
	HRT_PlayerAddProp   = 1
	HRT_PlayerResetProp = 2
	HRT_PlayerAddEvolve = 3
	HRT_BabyAddProp     = 4
	HRT_BabyReset       = 5
	HRT_BabyStr         = 6
	HRT_BabySkill       = 7
	HRT_BabyChange      = 8
	HRT_SkillAuto       = 9
	HRT_SkillItem       = 10
	HRT_EquipCompound   = 11
	HRT_EquipGem        = 12
	HRT_EmployeeBuy     = 13
	HRT_EmployeePos     = 14
	HRT_EmployeeSkill   = 15
	HRT_EmployeeEquip   = 16
	HRT_EmployeeEvolve  = 17
	HRT_MagicLevelUp    = 18
	HRT_MagicEvolve     = 19
	HRT_Max             = 20
)

func ToName_HelpRaiseType(id int) string {
	switch id {
	case HRT_None:
		return "HRT_None"
	case HRT_PlayerAddProp:
		return "HRT_PlayerAddProp"
	case HRT_PlayerResetProp:
		return "HRT_PlayerResetProp"
	case HRT_PlayerAddEvolve:
		return "HRT_PlayerAddEvolve"
	case HRT_BabyAddProp:
		return "HRT_BabyAddProp"
	case HRT_BabyReset:
		return "HRT_BabyReset"
	case HRT_BabyStr:
		return "HRT_BabyStr"
	case HRT_BabySkill:
		return "HRT_BabySkill"
	case HRT_BabyChange:
		return "HRT_BabyChange"
	case HRT_SkillAuto:
		return "HRT_SkillAuto"
	case HRT_SkillItem:
		return "HRT_SkillItem"
	case HRT_EquipCompound:
		return "HRT_EquipCompound"
	case HRT_EquipGem:
		return "HRT_EquipGem"
	case HRT_EmployeeBuy:
		return "HRT_EmployeeBuy"
	case HRT_EmployeePos:
		return "HRT_EmployeePos"
	case HRT_EmployeeSkill:
		return "HRT_EmployeeSkill"
	case HRT_EmployeeEquip:
		return "HRT_EmployeeEquip"
	case HRT_EmployeeEvolve:
		return "HRT_EmployeeEvolve"
	case HRT_MagicLevelUp:
		return "HRT_MagicLevelUp"
	case HRT_MagicEvolve:
		return "HRT_MagicEvolve"
	case HRT_Max:
		return "HRT_Max"
	default:
		return ""
	}
}
func ToId_HelpRaiseType(name string) int {
	switch name {
	case "HRT_None":
		return HRT_None
	case "HRT_PlayerAddProp":
		return HRT_PlayerAddProp
	case "HRT_PlayerResetProp":
		return HRT_PlayerResetProp
	case "HRT_PlayerAddEvolve":
		return HRT_PlayerAddEvolve
	case "HRT_BabyAddProp":
		return HRT_BabyAddProp
	case "HRT_BabyReset":
		return HRT_BabyReset
	case "HRT_BabyStr":
		return HRT_BabyStr
	case "HRT_BabySkill":
		return HRT_BabySkill
	case "HRT_BabyChange":
		return HRT_BabyChange
	case "HRT_SkillAuto":
		return HRT_SkillAuto
	case "HRT_SkillItem":
		return HRT_SkillItem
	case "HRT_EquipCompound":
		return HRT_EquipCompound
	case "HRT_EquipGem":
		return HRT_EquipGem
	case "HRT_EmployeeBuy":
		return HRT_EmployeeBuy
	case "HRT_EmployeePos":
		return HRT_EmployeePos
	case "HRT_EmployeeSkill":
		return HRT_EmployeeSkill
	case "HRT_EmployeeEquip":
		return HRT_EmployeeEquip
	case "HRT_EmployeeEvolve":
		return HRT_EmployeeEvolve
	case "HRT_MagicLevelUp":
		return HRT_MagicLevelUp
	case "HRT_MagicEvolve":
		return HRT_MagicEvolve
	case "HRT_Max":
		return HRT_Max
	default:
		return -1
	}
}
