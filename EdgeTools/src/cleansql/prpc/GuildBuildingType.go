package prpc

// enum GuildBuildingType
const (
	GBT_MIN        = 0
	GBT_Main       = 1
	GBT_Bank       = 2
	GBT_Shop       = 3
	GBT_Collection = 4
	GBT_Goddess    = 5
	GBT_Progenitus = 6
	GBT_MAX        = 7
)

func ToName_GuildBuildingType(id int) string {
	switch id {
	case GBT_MIN:
		return "GBT_MIN"
	case GBT_Main:
		return "GBT_Main"
	case GBT_Bank:
		return "GBT_Bank"
	case GBT_Shop:
		return "GBT_Shop"
	case GBT_Collection:
		return "GBT_Collection"
	case GBT_Goddess:
		return "GBT_Goddess"
	case GBT_Progenitus:
		return "GBT_Progenitus"
	case GBT_MAX:
		return "GBT_MAX"
	default:
		return ""
	}
}
func ToId_GuildBuildingType(name string) int {
	switch name {
	case "GBT_MIN":
		return GBT_MIN
	case "GBT_Main":
		return GBT_Main
	case "GBT_Bank":
		return GBT_Bank
	case "GBT_Shop":
		return GBT_Shop
	case "GBT_Collection":
		return GBT_Collection
	case "GBT_Goddess":
		return GBT_Goddess
	case "GBT_Progenitus":
		return GBT_Progenitus
	case "GBT_MAX":
		return GBT_MAX
	default:
		return -1
	}
}
