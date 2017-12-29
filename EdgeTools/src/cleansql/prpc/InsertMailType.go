package prpc

// enum InsertMailType
const (
	IGMT_PlayerId  = 0
	IGMT_AllOnline = 1
	IGMT_AllRegist = 2
)

func ToName_InsertMailType(id int) string {
	switch id {
	case IGMT_PlayerId:
		return "IGMT_PlayerId"
	case IGMT_AllOnline:
		return "IGMT_AllOnline"
	case IGMT_AllRegist:
		return "IGMT_AllRegist"
	default:
		return ""
	}
}
func ToId_InsertMailType(name string) int {
	switch name {
	case "IGMT_PlayerId":
		return IGMT_PlayerId
	case "IGMT_AllOnline":
		return IGMT_AllOnline
	case "IGMT_AllRegist":
		return IGMT_AllRegist
	default:
		return -1
	}
}
