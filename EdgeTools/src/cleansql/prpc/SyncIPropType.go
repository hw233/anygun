package prpc

// enum SyncIPropType
const (
	SPT_None     = 0
	SPT_Player   = 1
	SPT_Baby     = 2
	SPT_Employee = 3
	SPT_Max      = 4
)

func ToName_SyncIPropType(id int) string {
	switch id {
	case SPT_None:
		return "SPT_None"
	case SPT_Player:
		return "SPT_Player"
	case SPT_Baby:
		return "SPT_Baby"
	case SPT_Employee:
		return "SPT_Employee"
	case SPT_Max:
		return "SPT_Max"
	default:
		return ""
	}
}
func ToId_SyncIPropType(name string) int {
	switch name {
	case "SPT_None":
		return SPT_None
	case "SPT_Player":
		return SPT_Player
	case "SPT_Baby":
		return SPT_Baby
	case "SPT_Employee":
		return SPT_Employee
	case "SPT_Max":
		return SPT_Max
	default:
		return -1
	}
}
