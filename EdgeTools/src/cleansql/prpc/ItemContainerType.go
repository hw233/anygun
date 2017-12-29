package prpc

// enum ItemContainerType
const (
	ICT_EquipContainer = 0
	ICT_BagContainer   = 1
)

func ToName_ItemContainerType(id int) string {
	switch id {
	case ICT_EquipContainer:
		return "ICT_EquipContainer"
	case ICT_BagContainer:
		return "ICT_BagContainer"
	default:
		return ""
	}
}
func ToId_ItemContainerType(name string) int {
	switch name {
	case "ICT_EquipContainer":
		return ICT_EquipContainer
	case "ICT_BagContainer":
		return ICT_BagContainer
	default:
		return -1
	}
}
