package prpc

// enum MajorVersion
const (
	Major_0     = 0
	MajorNumber = 1
)

func ToName_MajorVersion(id int) string {
	switch id {
	case Major_0:
		return "Major_0"
	case MajorNumber:
		return "MajorNumber"
	default:
		return ""
	}
}
func ToId_MajorVersion(name string) int {
	switch name {
	case "Major_0":
		return Major_0
	case "MajorNumber":
		return MajorNumber
	default:
		return -1
	}
}
