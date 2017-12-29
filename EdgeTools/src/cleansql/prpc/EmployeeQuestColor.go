package prpc

// enum EmployeeQuestColor
const (
	EQC_White  = 0
	EQC_Blue   = 1
	EQC_Purple = 2
	EQC_Max    = 3
)

func ToName_EmployeeQuestColor(id int) string {
	switch id {
	case EQC_White:
		return "EQC_White"
	case EQC_Blue:
		return "EQC_Blue"
	case EQC_Purple:
		return "EQC_Purple"
	case EQC_Max:
		return "EQC_Max"
	default:
		return ""
	}
}
func ToId_EmployeeQuestColor(name string) int {
	switch name {
	case "EQC_White":
		return EQC_White
	case "EQC_Blue":
		return EQC_Blue
	case "EQC_Purple":
		return EQC_Purple
	case "EQC_Max":
		return EQC_Max
	default:
		return -1
	}
}
