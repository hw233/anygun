package prpc

// enum PeriodType
const (
	PT_Daily    = 0
	PT_Weekly   = 1
	PT_Customly = 2
)

func ToName_PeriodType(id int) string {
	switch id {
	case PT_Daily:
		return "PT_Daily"
	case PT_Weekly:
		return "PT_Weekly"
	case PT_Customly:
		return "PT_Customly"
	default:
		return ""
	}
}
func ToId_PeriodType(name string) int {
	switch name {
	case "PT_Daily":
		return PT_Daily
	case "PT_Weekly":
		return PT_Weekly
	case "PT_Customly":
		return PT_Customly
	default:
		return -1
	}
}
