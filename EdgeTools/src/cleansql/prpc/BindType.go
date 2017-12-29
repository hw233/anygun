package prpc

// enum BindType
const (
	BIT_None = 0
	BIT_Bag  = 1
	BIT_Use  = 2
	BIT_Max  = 3
)

func ToName_BindType(id int) string {
	switch id {
	case BIT_None:
		return "BIT_None"
	case BIT_Bag:
		return "BIT_Bag"
	case BIT_Use:
		return "BIT_Use"
	case BIT_Max:
		return "BIT_Max"
	default:
		return ""
	}
}
func ToId_BindType(name string) int {
	switch name {
	case "BIT_None":
		return BIT_None
	case "BIT_Bag":
		return BIT_Bag
	case "BIT_Use":
		return BIT_Use
	case "BIT_Max":
		return BIT_Max
	default:
		return -1
	}
}
