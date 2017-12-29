package prpc

// enum SceneType
const (
	SCT_None             = 0
	SCT_New              = 1
	SCT_Home             = 2
	SCT_Scene            = 3
	SCT_City             = 4
	SCT_Bairen           = 5
	SCT_Instance         = 6
	SCT_AlonePK          = 7
	SCT_TeamPK           = 8
	SCT_GuildHome        = 9
	SCT_GuildBattleScene = 10
	SCT_Max              = 11
)

func ToName_SceneType(id int) string {
	switch id {
	case SCT_None:
		return "SCT_None"
	case SCT_New:
		return "SCT_New"
	case SCT_Home:
		return "SCT_Home"
	case SCT_Scene:
		return "SCT_Scene"
	case SCT_City:
		return "SCT_City"
	case SCT_Bairen:
		return "SCT_Bairen"
	case SCT_Instance:
		return "SCT_Instance"
	case SCT_AlonePK:
		return "SCT_AlonePK"
	case SCT_TeamPK:
		return "SCT_TeamPK"
	case SCT_GuildHome:
		return "SCT_GuildHome"
	case SCT_GuildBattleScene:
		return "SCT_GuildBattleScene"
	case SCT_Max:
		return "SCT_Max"
	default:
		return ""
	}
}
func ToId_SceneType(name string) int {
	switch name {
	case "SCT_None":
		return SCT_None
	case "SCT_New":
		return SCT_New
	case "SCT_Home":
		return SCT_Home
	case "SCT_Scene":
		return SCT_Scene
	case "SCT_City":
		return SCT_City
	case "SCT_Bairen":
		return SCT_Bairen
	case "SCT_Instance":
		return SCT_Instance
	case "SCT_AlonePK":
		return SCT_AlonePK
	case "SCT_TeamPK":
		return SCT_TeamPK
	case "SCT_GuildHome":
		return SCT_GuildHome
	case "SCT_GuildBattleScene":
		return SCT_GuildBattleScene
	case "SCT_Max":
		return SCT_Max
	default:
		return -1
	}
}
