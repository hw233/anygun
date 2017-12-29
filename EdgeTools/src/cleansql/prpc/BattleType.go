package prpc

// enum BattleType
const (
	BT_None  = 0
	BT_PVE   = 1
	BT_PVR   = 2
	BT_PVP   = 3
	BT_PVH   = 4
	BT_PET   = 5
	BT_PK1   = 6
	BT_PK2   = 7
	BT_Guild = 8
	BT_MAX   = 9
)

func ToName_BattleType(id int) string {
	switch id {
	case BT_None:
		return "BT_None"
	case BT_PVE:
		return "BT_PVE"
	case BT_PVR:
		return "BT_PVR"
	case BT_PVP:
		return "BT_PVP"
	case BT_PVH:
		return "BT_PVH"
	case BT_PET:
		return "BT_PET"
	case BT_PK1:
		return "BT_PK1"
	case BT_PK2:
		return "BT_PK2"
	case BT_Guild:
		return "BT_Guild"
	case BT_MAX:
		return "BT_MAX"
	default:
		return ""
	}
}
func ToId_BattleType(name string) int {
	switch name {
	case "BT_None":
		return BT_None
	case "BT_PVE":
		return BT_PVE
	case "BT_PVR":
		return BT_PVR
	case "BT_PVP":
		return BT_PVP
	case "BT_PVH":
		return BT_PVH
	case "BT_PET":
		return BT_PET
	case "BT_PK1":
		return BT_PK1
	case "BT_PK2":
		return BT_PK2
	case "BT_Guild":
		return BT_Guild
	case "BT_MAX":
		return BT_MAX
	default:
		return -1
	}
}
