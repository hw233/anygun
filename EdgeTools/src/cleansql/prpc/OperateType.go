package prpc

// enum OperateType
const (
	OT_0  = 0
	OT_P1 = 1
	OT_P2 = 2
	OT_B  = 3
)

func ToName_OperateType(id int) string {
	switch id {
	case OT_0:
		return "OT_0"
	case OT_P1:
		return "OT_P1"
	case OT_P2:
		return "OT_P2"
	case OT_B:
		return "OT_B"
	default:
		return ""
	}
}
func ToId_OperateType(name string) int {
	switch name {
	case "OT_0":
		return OT_0
	case "OT_P1":
		return OT_P1
	case "OT_P2":
		return OT_P2
	case "OT_B":
		return OT_B
	default:
		return -1
	}
}
