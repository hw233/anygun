package prpc

// enum SkillTargetType
const (
	STT_None       = 0
	STT_Self       = 1
	STT_Team       = 2
	STT_TeamDead   = 3
	STT_TeamNoSelf = 4
	STT_All        = 5
	STT_AllNoSelf  = 6
	STT_Max        = 7
)

func ToName_SkillTargetType(id int) string {
	switch id {
	case STT_None:
		return "STT_None"
	case STT_Self:
		return "STT_Self"
	case STT_Team:
		return "STT_Team"
	case STT_TeamDead:
		return "STT_TeamDead"
	case STT_TeamNoSelf:
		return "STT_TeamNoSelf"
	case STT_All:
		return "STT_All"
	case STT_AllNoSelf:
		return "STT_AllNoSelf"
	case STT_Max:
		return "STT_Max"
	default:
		return ""
	}
}
func ToId_SkillTargetType(name string) int {
	switch name {
	case "STT_None":
		return STT_None
	case "STT_Self":
		return STT_Self
	case "STT_Team":
		return STT_Team
	case "STT_TeamDead":
		return STT_TeamDead
	case "STT_TeamNoSelf":
		return STT_TeamNoSelf
	case "STT_All":
		return STT_All
	case "STT_AllNoSelf":
		return STT_AllNoSelf
	case "STT_Max":
		return STT_Max
	default:
		return -1
	}
}
