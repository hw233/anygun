package prpc

// enum VipLevel
const (
	VL_None = 0
	VL_1    = 1
	VL_2    = 2
	VL_Max  = 3
)

func ToName_VipLevel(id int) string {
	switch id {
	case VL_None:
		return "VL_None"
	case VL_1:
		return "VL_1"
	case VL_2:
		return "VL_2"
	case VL_Max:
		return "VL_Max"
	default:
		return ""
	}
}
func ToId_VipLevel(name string) int {
	switch name {
	case "VL_None":
		return VL_None
	case "VL_1":
		return VL_1
	case "VL_2":
		return VL_2
	case "VL_Max":
		return VL_Max
	default:
		return -1
	}
}
