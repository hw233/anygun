package prpc

// enum GroupType
const (
	GT_None = 0
	GT_Up   = 1
	GT_Down = 2
)

func ToName_GroupType(id int) string {
	switch id {
	case GT_None:
		return "GT_None"
	case GT_Up:
		return "GT_Up"
	case GT_Down:
		return "GT_Down"
	default:
		return ""
	}
}
func ToId_GroupType(name string) int {
	switch name {
	case "GT_None":
		return GT_None
	case "GT_Up":
		return GT_Up
	case "GT_Down":
		return GT_Down
	default:
		return -1
	}
}
