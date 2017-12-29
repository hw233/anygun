package prpc

// enum SexType
const (
	ST_Unknown = 0
	ST_Male    = 1
	ST_Female  = 2
)

func ToName_SexType(id int) string {
	switch id {
	case ST_Unknown:
		return "ST_Unknown"
	case ST_Male:
		return "ST_Male"
	case ST_Female:
		return "ST_Female"
	default:
		return ""
	}
}
func ToId_SexType(name string) int {
	switch name {
	case "ST_Unknown":
		return ST_Unknown
	case "ST_Male":
		return ST_Male
	case "ST_Female":
		return ST_Female
	default:
		return -1
	}
}
