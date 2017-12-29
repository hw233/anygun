package prpc

// enum GatherStateType
const (
	GST_None     = 0
	GST_Vulgar   = 1
	GST_Advanced = 2
	GST_Max      = 3
)

func ToName_GatherStateType(id int) string {
	switch id {
	case GST_None:
		return "GST_None"
	case GST_Vulgar:
		return "GST_Vulgar"
	case GST_Advanced:
		return "GST_Advanced"
	case GST_Max:
		return "GST_Max"
	default:
		return ""
	}
}
func ToId_GatherStateType(name string) int {
	switch name {
	case "GST_None":
		return GST_None
	case "GST_Vulgar":
		return GST_Vulgar
	case "GST_Advanced":
		return GST_Advanced
	case "GST_Max":
		return GST_Max
	default:
		return -1
	}
}
