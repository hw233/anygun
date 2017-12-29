package prpc

// enum TogetherStateType
const (
	TST_None  = 0
	TST_Self  = 1
	TST_Enemy = 2
	TST_Max   = 3
)

func ToName_TogetherStateType(id int) string {
	switch id {
	case TST_None:
		return "TST_None"
	case TST_Self:
		return "TST_Self"
	case TST_Enemy:
		return "TST_Enemy"
	case TST_Max:
		return "TST_Max"
	default:
		return ""
	}
}
func ToId_TogetherStateType(name string) int {
	switch name {
	case "TST_None":
		return TST_None
	case "TST_Self":
		return TST_Self
	case "TST_Enemy":
		return TST_Enemy
	case "TST_Max":
		return TST_Max
	default:
		return -1
	}
}
