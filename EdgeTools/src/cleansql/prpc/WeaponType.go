package prpc

// enum WeaponType
const (
	WT_None      = 0
	WT_Axe       = 1
	WT_Sword     = 2
	WT_Spear     = 3
	WT_Bow       = 4
	WT_Stick     = 5
	WT_Knife     = 6
	WT_Hoe       = 7
	WT_Pickax    = 8
	WT_Lumberaxe = 9
	WT_Gloves    = 10
	WT_Max       = 11
)

func ToName_WeaponType(id int) string {
	switch id {
	case WT_None:
		return "WT_None"
	case WT_Axe:
		return "WT_Axe"
	case WT_Sword:
		return "WT_Sword"
	case WT_Spear:
		return "WT_Spear"
	case WT_Bow:
		return "WT_Bow"
	case WT_Stick:
		return "WT_Stick"
	case WT_Knife:
		return "WT_Knife"
	case WT_Hoe:
		return "WT_Hoe"
	case WT_Pickax:
		return "WT_Pickax"
	case WT_Lumberaxe:
		return "WT_Lumberaxe"
	case WT_Gloves:
		return "WT_Gloves"
	case WT_Max:
		return "WT_Max"
	default:
		return ""
	}
}
func ToId_WeaponType(name string) int {
	switch name {
	case "WT_None":
		return WT_None
	case "WT_Axe":
		return WT_Axe
	case "WT_Sword":
		return WT_Sword
	case "WT_Spear":
		return WT_Spear
	case "WT_Bow":
		return WT_Bow
	case "WT_Stick":
		return WT_Stick
	case "WT_Knife":
		return WT_Knife
	case "WT_Hoe":
		return WT_Hoe
	case "WT_Pickax":
		return WT_Pickax
	case "WT_Lumberaxe":
		return WT_Lumberaxe
	case "WT_Gloves":
		return WT_Gloves
	case "WT_Max":
		return WT_Max
	default:
		return -1
	}
}
