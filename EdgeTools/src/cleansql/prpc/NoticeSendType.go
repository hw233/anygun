package prpc

// enum NoticeSendType
const (
	NST_Immediately = 0
	NST_Timming     = 1
	NST_Loop        = 2
)

func ToName_NoticeSendType(id int) string {
	switch id {
	case NST_Immediately:
		return "NST_Immediately"
	case NST_Timming:
		return "NST_Timming"
	case NST_Loop:
		return "NST_Loop"
	default:
		return ""
	}
}
func ToId_NoticeSendType(name string) int {
	switch name {
	case "NST_Immediately":
		return NST_Immediately
	case "NST_Timming":
		return NST_Timming
	case "NST_Loop":
		return NST_Loop
	default:
		return -1
	}
}
