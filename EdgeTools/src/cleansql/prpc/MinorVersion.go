package prpc

// enum MinorVersion
const (
	Minor_0     = 0
	Minor_1     = 1
	Minor_2     = 2
	Minor_3     = 3
	Minor_4     = 4
	Minor_5     = 5
	Minor_6     = 6
	MinorNumber = 7
)

func ToName_MinorVersion(id int) string {
	switch id {
	case Minor_0:
		return "Minor_0"
	case Minor_1:
		return "Minor_1"
	case Minor_2:
		return "Minor_2"
	case Minor_3:
		return "Minor_3"
	case Minor_4:
		return "Minor_4"
	case Minor_5:
		return "Minor_5"
	case Minor_6:
		return "Minor_6"
	case MinorNumber:
		return "MinorNumber"
	default:
		return ""
	}
}
func ToId_MinorVersion(name string) int {
	switch name {
	case "Minor_0":
		return Minor_0
	case "Minor_1":
		return Minor_1
	case "Minor_2":
		return Minor_2
	case "Minor_3":
		return Minor_3
	case "Minor_4":
		return Minor_4
	case "Minor_5":
		return Minor_5
	case "Minor_6":
		return Minor_6
	case "MinorNumber":
		return MinorNumber
	default:
		return -1
	}
}
