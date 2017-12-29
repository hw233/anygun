package prpc

// enum EmployeesBattleGroup
const (
	EBG_None     = 0
	EBG_Free     = 1
	EBG_GroupOne = 2
	EBG_GroupTwo = 3
	EBG_Max      = 4
)

func ToName_EmployeesBattleGroup(id int) string {
	switch id {
	case EBG_None:
		return "EBG_None"
	case EBG_Free:
		return "EBG_Free"
	case EBG_GroupOne:
		return "EBG_GroupOne"
	case EBG_GroupTwo:
		return "EBG_GroupTwo"
	case EBG_Max:
		return "EBG_Max"
	default:
		return ""
	}
}
func ToId_EmployeesBattleGroup(name string) int {
	switch name {
	case "EBG_None":
		return EBG_None
	case "EBG_Free":
		return EBG_Free
	case "EBG_GroupOne":
		return EBG_GroupOne
	case "EBG_GroupTwo":
		return EBG_GroupTwo
	case "EBG_Max":
		return EBG_Max
	default:
		return -1
	}
}
