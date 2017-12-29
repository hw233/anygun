package prpc

// enum MailType
const (
	MT_Normal = 0
	MT_System = 1
)

func ToName_MailType(id int) string {
	switch id {
	case MT_Normal:
		return "MT_Normal"
	case MT_System:
		return "MT_System"
	default:
		return ""
	}
}
func ToId_MailType(name string) int {
	switch name {
	case "MT_Normal":
		return MT_Normal
	case "MT_System":
		return MT_System
	default:
		return -1
	}
}
