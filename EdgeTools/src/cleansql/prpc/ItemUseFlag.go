package prpc

// enum ItemUseFlag
const (
	IUF_None   = 0
	IUF_Battle = 1
	IUF_Scene  = 2
	IUF_All    = 3
)

func ToName_ItemUseFlag(id int) string {
	switch id {
	case IUF_None:
		return "IUF_None"
	case IUF_Battle:
		return "IUF_Battle"
	case IUF_Scene:
		return "IUF_Scene"
	case IUF_All:
		return "IUF_All"
	default:
		return ""
	}
}
func ToId_ItemUseFlag(name string) int {
	switch name {
	case "IUF_None":
		return IUF_None
	case "IUF_Battle":
		return IUF_Battle
	case "IUF_Scene":
		return IUF_Scene
	case "IUF_All":
		return IUF_All
	default:
		return -1
	}
}
