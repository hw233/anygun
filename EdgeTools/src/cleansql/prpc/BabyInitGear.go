package prpc

// enum BabyInitGear
const (
	BIG_None     = 0
	BIG_Stama    = 1
	BIG_Strength = 2
	BIG_Power    = 3
	BIG_Speed    = 4
	BIG_Magic    = 5
	BIG_Max      = 6
)

func ToName_BabyInitGear(id int) string {
	switch id {
	case BIG_None:
		return "BIG_None"
	case BIG_Stama:
		return "BIG_Stama"
	case BIG_Strength:
		return "BIG_Strength"
	case BIG_Power:
		return "BIG_Power"
	case BIG_Speed:
		return "BIG_Speed"
	case BIG_Magic:
		return "BIG_Magic"
	case BIG_Max:
		return "BIG_Max"
	default:
		return ""
	}
}
func ToId_BabyInitGear(name string) int {
	switch name {
	case "BIG_None":
		return BIG_None
	case "BIG_Stama":
		return BIG_Stama
	case "BIG_Strength":
		return BIG_Strength
	case "BIG_Power":
		return BIG_Power
	case "BIG_Speed":
		return BIG_Speed
	case "BIG_Magic":
		return BIG_Magic
	case "BIG_Max":
		return BIG_Max
	default:
		return -1
	}
}
