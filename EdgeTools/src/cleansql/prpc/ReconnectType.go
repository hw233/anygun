package prpc

// enum ReconnectType
const (
	RECT_None          = 0
	RECT_LoginOk       = 1
	RECT_EnterGameOk   = 2
	RECT_JoinTeamOk    = 3
	RECT_EnterSceneOk  = 4
	RECT_EnterBattleOk = 5
	RECT_Max           = 6
)

func ToName_ReconnectType(id int) string {
	switch id {
	case RECT_None:
		return "RECT_None"
	case RECT_LoginOk:
		return "RECT_LoginOk"
	case RECT_EnterGameOk:
		return "RECT_EnterGameOk"
	case RECT_JoinTeamOk:
		return "RECT_JoinTeamOk"
	case RECT_EnterSceneOk:
		return "RECT_EnterSceneOk"
	case RECT_EnterBattleOk:
		return "RECT_EnterBattleOk"
	case RECT_Max:
		return "RECT_Max"
	default:
		return ""
	}
}
func ToId_ReconnectType(name string) int {
	switch name {
	case "RECT_None":
		return RECT_None
	case "RECT_LoginOk":
		return RECT_LoginOk
	case "RECT_EnterGameOk":
		return RECT_EnterGameOk
	case "RECT_JoinTeamOk":
		return RECT_JoinTeamOk
	case "RECT_EnterSceneOk":
		return RECT_EnterSceneOk
	case "RECT_EnterBattleOk":
		return RECT_EnterBattleOk
	case "RECT_Max":
		return RECT_Max
	default:
		return -1
	}
}
