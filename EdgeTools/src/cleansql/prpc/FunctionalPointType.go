package prpc

// enum FunctionalPointType
const (
	FPT_None   = 0
	FPT_Tongji = 1
	FPT_Mogu   = 2
	FPT_Xiji   = 3
	FPT_Npc    = 4
	FPT_Max    = 5
)

func ToName_FunctionalPointType(id int) string {
	switch id {
	case FPT_None:
		return "FPT_None"
	case FPT_Tongji:
		return "FPT_Tongji"
	case FPT_Mogu:
		return "FPT_Mogu"
	case FPT_Xiji:
		return "FPT_Xiji"
	case FPT_Npc:
		return "FPT_Npc"
	case FPT_Max:
		return "FPT_Max"
	default:
		return ""
	}
}
func ToId_FunctionalPointType(name string) int {
	switch name {
	case "FPT_None":
		return FPT_None
	case "FPT_Tongji":
		return FPT_Tongji
	case "FPT_Mogu":
		return FPT_Mogu
	case "FPT_Xiji":
		return FPT_Xiji
	case "FPT_Npc":
		return FPT_Npc
	case "FPT_Max":
		return FPT_Max
	default:
		return -1
	}
}
