package prpc

// enum QualityColor
const (
	QC_None    = 0
	QC_White   = 1
	QC_Green   = 2
	QC_Blue    = 3
	QC_Blue1   = 4
	QC_Purple  = 5
	QC_Purple1 = 6
	QC_Purple2 = 7
	QC_Golden  = 8
	QC_Golden1 = 9
	QC_Golden2 = 10
	QC_Orange  = 11
	QC_Orange1 = 12
	QC_Orange2 = 13
	QC_Pink    = 14
	QC_Max     = 15
)

func ToName_QualityColor(id int) string {
	switch id {
	case QC_None:
		return "QC_None"
	case QC_White:
		return "QC_White"
	case QC_Green:
		return "QC_Green"
	case QC_Blue:
		return "QC_Blue"
	case QC_Blue1:
		return "QC_Blue1"
	case QC_Purple:
		return "QC_Purple"
	case QC_Purple1:
		return "QC_Purple1"
	case QC_Purple2:
		return "QC_Purple2"
	case QC_Golden:
		return "QC_Golden"
	case QC_Golden1:
		return "QC_Golden1"
	case QC_Golden2:
		return "QC_Golden2"
	case QC_Orange:
		return "QC_Orange"
	case QC_Orange1:
		return "QC_Orange1"
	case QC_Orange2:
		return "QC_Orange2"
	case QC_Pink:
		return "QC_Pink"
	case QC_Max:
		return "QC_Max"
	default:
		return ""
	}
}
func ToId_QualityColor(name string) int {
	switch name {
	case "QC_None":
		return QC_None
	case "QC_White":
		return QC_White
	case "QC_Green":
		return QC_Green
	case "QC_Blue":
		return QC_Blue
	case "QC_Blue1":
		return QC_Blue1
	case "QC_Purple":
		return QC_Purple
	case "QC_Purple1":
		return QC_Purple1
	case "QC_Purple2":
		return QC_Purple2
	case "QC_Golden":
		return QC_Golden
	case "QC_Golden1":
		return QC_Golden1
	case "QC_Golden2":
		return QC_Golden2
	case "QC_Orange":
		return QC_Orange
	case "QC_Orange1":
		return QC_Orange1
	case "QC_Orange2":
		return QC_Orange2
	case "QC_Pink":
		return QC_Pink
	case "QC_Max":
		return QC_Max
	default:
		return -1
	}
}
