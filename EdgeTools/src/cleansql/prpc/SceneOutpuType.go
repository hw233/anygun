package prpc

// enum SceneOutpuType
const (
	SOT_None = 0
	SOT_PVE  = 1
	SOT_PVP  = 2
	SOT_Max  = 3
)

func ToName_SceneOutpuType(id int) string {
	switch id {
	case SOT_None:
		return "SOT_None"
	case SOT_PVE:
		return "SOT_PVE"
	case SOT_PVP:
		return "SOT_PVP"
	case SOT_Max:
		return "SOT_Max"
	default:
		return ""
	}
}
func ToId_SceneOutpuType(name string) int {
	switch name {
	case "SOT_None":
		return SOT_None
	case "SOT_PVE":
		return SOT_PVE
	case "SOT_PVP":
		return SOT_PVP
	case "SOT_Max":
		return SOT_Max
	default:
		return -1
	}
}
