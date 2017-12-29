package prpc

// enum QuestKind
const (
	QK_None       = 0
	QK_Main       = 1
	QK_Daily      = 2
	QK_Profession = 3
	QK_Sub        = 4
	QK_Tongji     = 5
	QK_Copy       = 6
	QK_Wishing    = 7
	QK_Guild      = 8
	QK_Rand       = 9
	QK_Max        = 10
)

func ToName_QuestKind(id int) string {
	switch id {
	case QK_None:
		return "QK_None"
	case QK_Main:
		return "QK_Main"
	case QK_Daily:
		return "QK_Daily"
	case QK_Profession:
		return "QK_Profession"
	case QK_Sub:
		return "QK_Sub"
	case QK_Tongji:
		return "QK_Tongji"
	case QK_Copy:
		return "QK_Copy"
	case QK_Wishing:
		return "QK_Wishing"
	case QK_Guild:
		return "QK_Guild"
	case QK_Rand:
		return "QK_Rand"
	case QK_Max:
		return "QK_Max"
	default:
		return ""
	}
}
func ToId_QuestKind(name string) int {
	switch name {
	case "QK_None":
		return QK_None
	case "QK_Main":
		return QK_Main
	case "QK_Daily":
		return QK_Daily
	case "QK_Profession":
		return QK_Profession
	case "QK_Sub":
		return QK_Sub
	case "QK_Tongji":
		return QK_Tongji
	case "QK_Copy":
		return QK_Copy
	case "QK_Wishing":
		return QK_Wishing
	case "QK_Guild":
		return QK_Guild
	case "QK_Rand":
		return QK_Rand
	case "QK_Max":
		return QK_Max
	default:
		return -1
	}
}
