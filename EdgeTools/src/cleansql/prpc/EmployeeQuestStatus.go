package prpc

// enum EmployeeQuestStatus
const (
	EQS_None     = 0
	EQS_Running  = 1
	EQS_Complate = 2
)

func ToName_EmployeeQuestStatus(id int) string {
	switch id {
	case EQS_None:
		return "EQS_None"
	case EQS_Running:
		return "EQS_Running"
	case EQS_Complate:
		return "EQS_Complate"
	default:
		return ""
	}
}
func ToId_EmployeeQuestStatus(name string) int {
	switch name {
	case "EQS_None":
		return EQS_None
	case "EQS_Running":
		return EQS_Running
	case "EQS_Complate":
		return EQS_Complate
	default:
		return -1
	}
}
