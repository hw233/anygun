package prpc

// enum SceneFilterType
const (
	SFT_None   = 0
	SFT_Team   = 1
	SFT_Friend = 2
	SFT_Guild  = 3
	SFT_All    = 4
	SFT_Max    = 5
)

func ToName_SceneFilterType(id int) string {
	switch id {
	case SFT_None:
		return "SFT_None"
	case SFT_Team:
		return "SFT_Team"
	case SFT_Friend:
		return "SFT_Friend"
	case SFT_Guild:
		return "SFT_Guild"
	case SFT_All:
		return "SFT_All"
	case SFT_Max:
		return "SFT_Max"
	default:
		return ""
	}
}
func ToId_SceneFilterType(name string) int {
	switch name {
	case "SFT_None":
		return SFT_None
	case "SFT_Team":
		return SFT_Team
	case "SFT_Friend":
		return SFT_Friend
	case "SFT_Guild":
		return SFT_Guild
	case "SFT_All":
		return SFT_All
	case "SFT_Max":
		return SFT_Max
	default:
		return -1
	}
}
