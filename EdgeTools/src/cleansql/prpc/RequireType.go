package prpc

// enum RequireType
const (
	RT_Nil = 0
)

func ToName_RequireType(id int) string {
	switch id {
	case RT_Nil:
		return "RT_Nil"
	default:
		return ""
	}
}
func ToId_RequireType(name string) int {
	switch name {
	case "RT_Nil":
		return RT_Nil
	default:
		return -1
	}
}
