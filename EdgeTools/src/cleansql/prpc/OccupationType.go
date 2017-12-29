package prpc

// enum OccupationType
const (
	OT_None       = 0
	OT_HeavyArmor = 1
	OT_LightArmor = 2
	OT_Spell      = 3
	OT_Max        = 4
)

func ToName_OccupationType(id int) string {
	switch id {
	case OT_None:
		return "OT_None"
	case OT_HeavyArmor:
		return "OT_HeavyArmor"
	case OT_LightArmor:
		return "OT_LightArmor"
	case OT_Spell:
		return "OT_Spell"
	case OT_Max:
		return "OT_Max"
	default:
		return ""
	}
}
func ToId_OccupationType(name string) int {
	switch name {
	case "OT_None":
		return OT_None
	case "OT_HeavyArmor":
		return OT_HeavyArmor
	case "OT_LightArmor":
		return OT_LightArmor
	case "OT_Spell":
		return OT_Spell
	case "OT_Max":
		return OT_Max
	default:
		return -1
	}
}
