package prpc

// enum EquipmentSlot
const (
	ES_None       = 0
	ES_Boot       = 1
	ES_SingleHand = 2
	ES_Ornament_0 = 3
	ES_Head       = 4
	ES_Ornament_1 = 5
	ES_Body       = 6
	ES_DoubleHand = 7
	ES_Crystal    = 8
	ES_Fashion    = 9
	ES_Max        = 10
)

func ToName_EquipmentSlot(id int) string {
	switch id {
	case ES_None:
		return "ES_None"
	case ES_Boot:
		return "ES_Boot"
	case ES_SingleHand:
		return "ES_SingleHand"
	case ES_Ornament_0:
		return "ES_Ornament_0"
	case ES_Head:
		return "ES_Head"
	case ES_Ornament_1:
		return "ES_Ornament_1"
	case ES_Body:
		return "ES_Body"
	case ES_DoubleHand:
		return "ES_DoubleHand"
	case ES_Crystal:
		return "ES_Crystal"
	case ES_Fashion:
		return "ES_Fashion"
	case ES_Max:
		return "ES_Max"
	default:
		return ""
	}
}
func ToId_EquipmentSlot(name string) int {
	switch name {
	case "ES_None":
		return ES_None
	case "ES_Boot":
		return ES_Boot
	case "ES_SingleHand":
		return ES_SingleHand
	case "ES_Ornament_0":
		return ES_Ornament_0
	case "ES_Head":
		return ES_Head
	case "ES_Ornament_1":
		return ES_Ornament_1
	case "ES_Body":
		return ES_Body
	case "ES_DoubleHand":
		return ES_DoubleHand
	case "ES_Crystal":
		return ES_Crystal
	case "ES_Fashion":
		return ES_Fashion
	case "ES_Max":
		return ES_Max
	default:
		return -1
	}
}
