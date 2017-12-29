package prpc

// enum WishType
const (
	WIT_None     = 0
	WIT_Exp      = 1
	WIT_Money    = 2
	WIT_Baby     = 3
	WIT_Employee = 4
	WIT_Max      = 5
)

func ToName_WishType(id int) string {
	switch id {
	case WIT_None:
		return "WIT_None"
	case WIT_Exp:
		return "WIT_Exp"
	case WIT_Money:
		return "WIT_Money"
	case WIT_Baby:
		return "WIT_Baby"
	case WIT_Employee:
		return "WIT_Employee"
	case WIT_Max:
		return "WIT_Max"
	default:
		return ""
	}
}
func ToId_WishType(name string) int {
	switch name {
	case "WIT_None":
		return WIT_None
	case "WIT_Exp":
		return WIT_Exp
	case "WIT_Money":
		return WIT_Money
	case "WIT_Baby":
		return WIT_Baby
	case "WIT_Employee":
		return WIT_Employee
	case "WIT_Max":
		return WIT_Max
	default:
		return -1
	}
}
