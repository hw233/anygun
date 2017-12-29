package prpc

// enum GMCommandType
const (
	GMCT_NoTalk     = 0
	GMCT_Freeze     = 1
	GMCT_Seal       = 2
	GMCT_Kick       = 3
	GMCT_OpenTalk   = 4
	GMCT_Unseal     = 5
	GMCT_SkipGuide  = 6
	GMCT_AddMoney   = 7
	GMCT_AddDiamond = 8
	GMCT_AddExp     = 9
	GMCT_OpenGM     = 10
	GMCT_CloseGM    = 11
	GMCT_DoScript   = 12
	GMCT_Max        = 13
)

func ToName_GMCommandType(id int) string {
	switch id {
	case GMCT_NoTalk:
		return "GMCT_NoTalk"
	case GMCT_Freeze:
		return "GMCT_Freeze"
	case GMCT_Seal:
		return "GMCT_Seal"
	case GMCT_Kick:
		return "GMCT_Kick"
	case GMCT_OpenTalk:
		return "GMCT_OpenTalk"
	case GMCT_Unseal:
		return "GMCT_Unseal"
	case GMCT_SkipGuide:
		return "GMCT_SkipGuide"
	case GMCT_AddMoney:
		return "GMCT_AddMoney"
	case GMCT_AddDiamond:
		return "GMCT_AddDiamond"
	case GMCT_AddExp:
		return "GMCT_AddExp"
	case GMCT_OpenGM:
		return "GMCT_OpenGM"
	case GMCT_CloseGM:
		return "GMCT_CloseGM"
	case GMCT_DoScript:
		return "GMCT_DoScript"
	case GMCT_Max:
		return "GMCT_Max"
	default:
		return ""
	}
}
func ToId_GMCommandType(name string) int {
	switch name {
	case "GMCT_NoTalk":
		return GMCT_NoTalk
	case "GMCT_Freeze":
		return GMCT_Freeze
	case "GMCT_Seal":
		return GMCT_Seal
	case "GMCT_Kick":
		return GMCT_Kick
	case "GMCT_OpenTalk":
		return GMCT_OpenTalk
	case "GMCT_Unseal":
		return GMCT_Unseal
	case "GMCT_SkipGuide":
		return GMCT_SkipGuide
	case "GMCT_AddMoney":
		return GMCT_AddMoney
	case "GMCT_AddDiamond":
		return GMCT_AddDiamond
	case "GMCT_AddExp":
		return GMCT_AddExp
	case "GMCT_OpenGM":
		return GMCT_OpenGM
	case "GMCT_CloseGM":
		return GMCT_CloseGM
	case "GMCT_DoScript":
		return GMCT_DoScript
	case "GMCT_Max":
		return GMCT_Max
	default:
		return -1
	}
}
