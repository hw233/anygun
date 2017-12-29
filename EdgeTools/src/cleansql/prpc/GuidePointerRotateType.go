package prpc

// enum GuidePointerRotateType
const (
	GPRT_None = 0
	GPRT_R45  = 1
	GPRT_R90  = 2
	GPRT_R135 = 3
	GPRT_R180 = 4
	GPRT_R225 = 5
	GPRT_R270 = 6
	GPRT_R315 = 7
	GPRT_Max  = 8
)

func ToName_GuidePointerRotateType(id int) string {
	switch id {
	case GPRT_None:
		return "GPRT_None"
	case GPRT_R45:
		return "GPRT_R45"
	case GPRT_R90:
		return "GPRT_R90"
	case GPRT_R135:
		return "GPRT_R135"
	case GPRT_R180:
		return "GPRT_R180"
	case GPRT_R225:
		return "GPRT_R225"
	case GPRT_R270:
		return "GPRT_R270"
	case GPRT_R315:
		return "GPRT_R315"
	case GPRT_Max:
		return "GPRT_Max"
	default:
		return ""
	}
}
func ToId_GuidePointerRotateType(name string) int {
	switch name {
	case "GPRT_None":
		return GPRT_None
	case "GPRT_R45":
		return GPRT_R45
	case "GPRT_R90":
		return GPRT_R90
	case "GPRT_R135":
		return GPRT_R135
	case "GPRT_R180":
		return GPRT_R180
	case "GPRT_R225":
		return GPRT_R225
	case "GPRT_R270":
		return GPRT_R270
	case "GPRT_R315":
		return GPRT_R315
	case "GPRT_Max":
		return GPRT_Max
	default:
		return -1
	}
}
