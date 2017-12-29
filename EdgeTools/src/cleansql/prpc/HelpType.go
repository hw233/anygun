package prpc

// enum HelpType
const (
	HT_None     = 0
	HT_Money    = 1
	HT_Diamond  = 2
	HT_Role     = 3
	HT_Baby     = 4
	HT_Employee = 5
	HT_Skill    = 6
	HT_Exp      = 7
	HT_Magic    = 8
	HT_Equip    = 9
)

func ToName_HelpType(id int) string {
	switch id {
	case HT_None:
		return "HT_None"
	case HT_Money:
		return "HT_Money"
	case HT_Diamond:
		return "HT_Diamond"
	case HT_Role:
		return "HT_Role"
	case HT_Baby:
		return "HT_Baby"
	case HT_Employee:
		return "HT_Employee"
	case HT_Skill:
		return "HT_Skill"
	case HT_Exp:
		return "HT_Exp"
	case HT_Magic:
		return "HT_Magic"
	case HT_Equip:
		return "HT_Equip"
	default:
		return ""
	}
}
func ToId_HelpType(name string) int {
	switch name {
	case "HT_None":
		return HT_None
	case "HT_Money":
		return HT_Money
	case "HT_Diamond":
		return HT_Diamond
	case "HT_Role":
		return HT_Role
	case "HT_Baby":
		return HT_Baby
	case "HT_Employee":
		return HT_Employee
	case "HT_Skill":
		return HT_Skill
	case "HT_Exp":
		return HT_Exp
	case "HT_Magic":
		return HT_Magic
	case "HT_Equip":
		return HT_Equip
	default:
		return -1
	}
}
