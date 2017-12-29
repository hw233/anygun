package prpc

// enum QuestType
const (
	QT_None       = 0
	QT_Dialog     = 1
	QT_Battle     = 2
	QT_Kill       = 3
	QT_KillAI     = 4
	QT_Item       = 5
	QT_Profession = 6
	QT_Other      = 7
	QT_Max        = 8
)

func ToName_QuestType(id int) string {
	switch id {
	case QT_None:
		return "QT_None"
	case QT_Dialog:
		return "QT_Dialog"
	case QT_Battle:
		return "QT_Battle"
	case QT_Kill:
		return "QT_Kill"
	case QT_KillAI:
		return "QT_KillAI"
	case QT_Item:
		return "QT_Item"
	case QT_Profession:
		return "QT_Profession"
	case QT_Other:
		return "QT_Other"
	case QT_Max:
		return "QT_Max"
	default:
		return ""
	}
}
func ToId_QuestType(name string) int {
	switch name {
	case "QT_None":
		return QT_None
	case "QT_Dialog":
		return QT_Dialog
	case "QT_Battle":
		return QT_Battle
	case "QT_Kill":
		return QT_Kill
	case "QT_KillAI":
		return QT_KillAI
	case "QT_Item":
		return QT_Item
	case "QT_Profession":
		return QT_Profession
	case "QT_Other":
		return QT_Other
	case "QT_Max":
		return QT_Max
	default:
		return -1
	}
}
