package prpc

// enum FixType
const (
	FT_None    = 0
	FT_Money   = 1
	FT_Diamond = 2
	FT_Max     = 3
)

func ToName_FixType(id int) string {
	switch id {
	case FT_None:
		return "FT_None"
	case FT_Money:
		return "FT_Money"
	case FT_Diamond:
		return "FT_Diamond"
	case FT_Max:
		return "FT_Max"
	default:
		return ""
	}
}
func ToId_FixType(name string) int {
	switch name {
	case "FT_None":
		return FT_None
	case "FT_Money":
		return FT_Money
	case "FT_Diamond":
		return FT_Diamond
	case "FT_Max":
		return FT_Max
	default:
		return -1
	}
}
