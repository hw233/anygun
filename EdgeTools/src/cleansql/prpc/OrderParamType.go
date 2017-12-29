package prpc

// enum OrderParamType
const (
	OPT_None   = 0
	OPT_BabyId = 1
	OPT_Unite  = 2
	OPT_Huwei  = 3
	OPT_IsNo   = 4
	OPT_Max    = 5
)

func ToName_OrderParamType(id int) string {
	switch id {
	case OPT_None:
		return "OPT_None"
	case OPT_BabyId:
		return "OPT_BabyId"
	case OPT_Unite:
		return "OPT_Unite"
	case OPT_Huwei:
		return "OPT_Huwei"
	case OPT_IsNo:
		return "OPT_IsNo"
	case OPT_Max:
		return "OPT_Max"
	default:
		return ""
	}
}
func ToId_OrderParamType(name string) int {
	switch name {
	case "OPT_None":
		return OPT_None
	case "OPT_BabyId":
		return OPT_BabyId
	case "OPT_Unite":
		return OPT_Unite
	case "OPT_Huwei":
		return OPT_Huwei
	case "OPT_IsNo":
		return OPT_IsNo
	case "OPT_Max":
		return OPT_Max
	default:
		return -1
	}
}
