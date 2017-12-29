package prpc

// enum SkillType
const (
	SKT_None              = 0
	SKT_DefaultSecActive  = 1
	SKT_DefaultSecPassive = 2
	SKT_DefaultActive     = 3
	SKT_DefaultPassive    = 4
	SKT_Active            = 5
	SKT_Passive           = 6
	SKT_Gather            = 7
	SKT_Make              = 8
	SKT_CannotUse         = 9
	SKT_GuildPlayerSkill  = 10
	SKT_GuildBabySkill    = 11
	SKT_Max               = 12
)

func ToName_SkillType(id int) string {
	switch id {
	case SKT_None:
		return "SKT_None"
	case SKT_DefaultSecActive:
		return "SKT_DefaultSecActive"
	case SKT_DefaultSecPassive:
		return "SKT_DefaultSecPassive"
	case SKT_DefaultActive:
		return "SKT_DefaultActive"
	case SKT_DefaultPassive:
		return "SKT_DefaultPassive"
	case SKT_Active:
		return "SKT_Active"
	case SKT_Passive:
		return "SKT_Passive"
	case SKT_Gather:
		return "SKT_Gather"
	case SKT_Make:
		return "SKT_Make"
	case SKT_CannotUse:
		return "SKT_CannotUse"
	case SKT_GuildPlayerSkill:
		return "SKT_GuildPlayerSkill"
	case SKT_GuildBabySkill:
		return "SKT_GuildBabySkill"
	case SKT_Max:
		return "SKT_Max"
	default:
		return ""
	}
}
func ToId_SkillType(name string) int {
	switch name {
	case "SKT_None":
		return SKT_None
	case "SKT_DefaultSecActive":
		return SKT_DefaultSecActive
	case "SKT_DefaultSecPassive":
		return SKT_DefaultSecPassive
	case "SKT_DefaultActive":
		return SKT_DefaultActive
	case "SKT_DefaultPassive":
		return SKT_DefaultPassive
	case "SKT_Active":
		return SKT_Active
	case "SKT_Passive":
		return SKT_Passive
	case "SKT_Gather":
		return SKT_Gather
	case "SKT_Make":
		return SKT_Make
	case "SKT_CannotUse":
		return SKT_CannotUse
	case "SKT_GuildPlayerSkill":
		return SKT_GuildPlayerSkill
	case "SKT_GuildBabySkill":
		return SKT_GuildBabySkill
	case "SKT_Max":
		return SKT_Max
	default:
		return -1
	}
}
