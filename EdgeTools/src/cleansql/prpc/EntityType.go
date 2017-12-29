package prpc

// enum EntityType
const (
	ET_None    = 0
	ET_Player  = 1
	ET_Baby    = 2
	ET_Monster = 3
	ET_Boss    = 4
	ET_Emplyee = 5
	ET_Max     = 6
)

func ToName_EntityType(id int) string {
	switch id {
	case ET_None:
		return "ET_None"
	case ET_Player:
		return "ET_Player"
	case ET_Baby:
		return "ET_Baby"
	case ET_Monster:
		return "ET_Monster"
	case ET_Boss:
		return "ET_Boss"
	case ET_Emplyee:
		return "ET_Emplyee"
	case ET_Max:
		return "ET_Max"
	default:
		return ""
	}
}
func ToId_EntityType(name string) int {
	switch name {
	case "ET_None":
		return ET_None
	case "ET_Player":
		return ET_Player
	case "ET_Baby":
		return ET_Baby
	case "ET_Monster":
		return ET_Monster
	case "ET_Boss":
		return ET_Boss
	case "ET_Emplyee":
		return ET_Emplyee
	case "ET_Max":
		return ET_Max
	default:
		return -1
	}
}
