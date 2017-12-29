package prpc

// enum BoxType
const (
	BX_None   = 0
	BX_Normal = 1
	BX_Blue   = 2
	BX_Glod   = 3
)

func ToName_BoxType(id int) string {
	switch id {
	case BX_None:
		return "BX_None"
	case BX_Normal:
		return "BX_Normal"
	case BX_Blue:
		return "BX_Blue"
	case BX_Glod:
		return "BX_Glod"
	default:
		return ""
	}
}
func ToId_BoxType(name string) int {
	switch name {
	case "BX_None":
		return BX_None
	case "BX_Normal":
		return BX_Normal
	case "BX_Blue":
		return BX_Blue
	case "BX_Glod":
		return BX_Glod
	default:
		return -1
	}
}
