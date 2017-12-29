package prpc

// enum GiftType
const (
	GFT_Bug   = 0
	GFT_UC1   = 1
	GFT_UC2   = 2
	GFT_Group = 3
)

func ToName_GiftType(id int) string {
	switch id {
	case GFT_Bug:
		return "GFT_Bug"
	case GFT_UC1:
		return "GFT_UC1"
	case GFT_UC2:
		return "GFT_UC2"
	case GFT_Group:
		return "GFT_Group"
	default:
		return ""
	}
}
func ToId_GiftType(name string) int {
	switch name {
	case "GFT_Bug":
		return GFT_Bug
	case "GFT_UC1":
		return GFT_UC1
	case "GFT_UC2":
		return GFT_UC2
	case "GFT_Group":
		return GFT_Group
	default:
		return -1
	}
}
