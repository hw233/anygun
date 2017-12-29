package prpc

// enum SellItemType
const (
	SelIT_None = 0
	SelIT_Max  = 1
)

func ToName_SellItemType(id int) string {
	switch id {
	case SelIT_None:
		return "SelIT_None"
	case SelIT_Max:
		return "SelIT_Max"
	default:
		return ""
	}
}
func ToId_SellItemType(name string) int {
	switch name {
	case "SelIT_None":
		return SelIT_None
	case "SelIT_Max":
		return SelIT_Max
	default:
		return -1
	}
}
