package prpc

// enum RobotActionType
const (
	RAT_None      = 0
	RAT_Resting   = 1
	RAT_Move      = 2
	RAT_QuestMove = 3
	RAT_TeamMove  = 4
	RAT_Max       = 5
)

func ToName_RobotActionType(id int) string {
	switch id {
	case RAT_None:
		return "RAT_None"
	case RAT_Resting:
		return "RAT_Resting"
	case RAT_Move:
		return "RAT_Move"
	case RAT_QuestMove:
		return "RAT_QuestMove"
	case RAT_TeamMove:
		return "RAT_TeamMove"
	case RAT_Max:
		return "RAT_Max"
	default:
		return ""
	}
}
func ToId_RobotActionType(name string) int {
	switch name {
	case "RAT_None":
		return RAT_None
	case "RAT_Resting":
		return RAT_Resting
	case "RAT_Move":
		return RAT_Move
	case "RAT_QuestMove":
		return RAT_QuestMove
	case "RAT_TeamMove":
		return RAT_TeamMove
	case "RAT_Max":
		return RAT_Max
	default:
		return -1
	}
}
