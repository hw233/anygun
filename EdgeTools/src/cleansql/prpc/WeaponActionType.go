package prpc

// enum WeaponActionType
const (
	WAT_None  = 0
	WAT_Chop  = 1
	WAT_Stab  = 2
	WAT_Bow   = 3
	WAT_Throw = 4
	WAT_Max   = 5
)

func ToName_WeaponActionType(id int) string {
	switch id {
	case WAT_None:
		return "WAT_None"
	case WAT_Chop:
		return "WAT_Chop"
	case WAT_Stab:
		return "WAT_Stab"
	case WAT_Bow:
		return "WAT_Bow"
	case WAT_Throw:
		return "WAT_Throw"
	case WAT_Max:
		return "WAT_Max"
	default:
		return ""
	}
}
func ToId_WeaponActionType(name string) int {
	switch name {
	case "WAT_None":
		return WAT_None
	case "WAT_Chop":
		return WAT_Chop
	case "WAT_Stab":
		return WAT_Stab
	case "WAT_Bow":
		return WAT_Bow
	case "WAT_Throw":
		return WAT_Throw
	case "WAT_Max":
		return WAT_Max
	default:
		return -1
	}
}
