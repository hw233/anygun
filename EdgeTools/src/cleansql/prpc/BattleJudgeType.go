package prpc

// enum BattleJudgeType
const (
	BJT_None     = 0
	BJT_Continue = 1
	BJT_Win      = 2
	BJT_Lose     = 3
)

func ToName_BattleJudgeType(id int) string {
	switch id {
	case BJT_None:
		return "BJT_None"
	case BJT_Continue:
		return "BJT_Continue"
	case BJT_Win:
		return "BJT_Win"
	case BJT_Lose:
		return "BJT_Lose"
	default:
		return ""
	}
}
func ToId_BattleJudgeType(name string) int {
	switch name {
	case "BJT_None":
		return BJT_None
	case "BJT_Continue":
		return BJT_Continue
	case "BJT_Win":
		return BJT_Win
	case "BJT_Lose":
		return BJT_Lose
	default:
		return -1
	}
}
