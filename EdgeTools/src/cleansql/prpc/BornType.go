package prpc

// enum BornType
const (
	BOT_None            = 0
	BOT_BornPos         = 1
	BOT_Cached          = 2
	BOT_FromSceneEntry  = 3
	BOT_FromMazeEntry   = 4
	BOT_NormalMazeEntry = 5
	BOT_Max             = 6
)

func ToName_BornType(id int) string {
	switch id {
	case BOT_None:
		return "BOT_None"
	case BOT_BornPos:
		return "BOT_BornPos"
	case BOT_Cached:
		return "BOT_Cached"
	case BOT_FromSceneEntry:
		return "BOT_FromSceneEntry"
	case BOT_FromMazeEntry:
		return "BOT_FromMazeEntry"
	case BOT_NormalMazeEntry:
		return "BOT_NormalMazeEntry"
	case BOT_Max:
		return "BOT_Max"
	default:
		return ""
	}
}
func ToId_BornType(name string) int {
	switch name {
	case "BOT_None":
		return BOT_None
	case "BOT_BornPos":
		return BOT_BornPos
	case "BOT_Cached":
		return BOT_Cached
	case "BOT_FromSceneEntry":
		return BOT_FromSceneEntry
	case "BOT_FromMazeEntry":
		return BOT_FromMazeEntry
	case "BOT_NormalMazeEntry":
		return BOT_NormalMazeEntry
	case "BOT_Max":
		return BOT_Max
	default:
		return -1
	}
}
