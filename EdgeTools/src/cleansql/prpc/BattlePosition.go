package prpc

// enum BattlePosition
const (
	BP_None  = 0
	BP_Down0 = 1
	BP_Down1 = 2
	BP_Down2 = 3
	BP_Down3 = 4
	BP_Down4 = 5
	BP_Down5 = 6
	BP_Down6 = 7
	BP_Down7 = 8
	BP_Down8 = 9
	BP_Down9 = 10
	BP_Up0   = 11
	BP_Up1   = 12
	BP_Up2   = 13
	BP_Up3   = 14
	BP_Up4   = 15
	BP_Up5   = 16
	BP_Up6   = 17
	BP_Up7   = 18
	BP_Up8   = 19
	BP_Up9   = 20
	BP_Max   = 21
)

func ToName_BattlePosition(id int) string {
	switch id {
	case BP_None:
		return "BP_None"
	case BP_Down0:
		return "BP_Down0"
	case BP_Down1:
		return "BP_Down1"
	case BP_Down2:
		return "BP_Down2"
	case BP_Down3:
		return "BP_Down3"
	case BP_Down4:
		return "BP_Down4"
	case BP_Down5:
		return "BP_Down5"
	case BP_Down6:
		return "BP_Down6"
	case BP_Down7:
		return "BP_Down7"
	case BP_Down8:
		return "BP_Down8"
	case BP_Down9:
		return "BP_Down9"
	case BP_Up0:
		return "BP_Up0"
	case BP_Up1:
		return "BP_Up1"
	case BP_Up2:
		return "BP_Up2"
	case BP_Up3:
		return "BP_Up3"
	case BP_Up4:
		return "BP_Up4"
	case BP_Up5:
		return "BP_Up5"
	case BP_Up6:
		return "BP_Up6"
	case BP_Up7:
		return "BP_Up7"
	case BP_Up8:
		return "BP_Up8"
	case BP_Up9:
		return "BP_Up9"
	case BP_Max:
		return "BP_Max"
	default:
		return ""
	}
}
func ToId_BattlePosition(name string) int {
	switch name {
	case "BP_None":
		return BP_None
	case "BP_Down0":
		return BP_Down0
	case "BP_Down1":
		return BP_Down1
	case "BP_Down2":
		return BP_Down2
	case "BP_Down3":
		return BP_Down3
	case "BP_Down4":
		return BP_Down4
	case "BP_Down5":
		return BP_Down5
	case "BP_Down6":
		return BP_Down6
	case "BP_Down7":
		return BP_Down7
	case "BP_Down8":
		return BP_Down8
	case "BP_Down9":
		return BP_Down9
	case "BP_Up0":
		return BP_Up0
	case "BP_Up1":
		return BP_Up1
	case "BP_Up2":
		return BP_Up2
	case "BP_Up3":
		return BP_Up3
	case "BP_Up4":
		return BP_Up4
	case "BP_Up5":
		return BP_Up5
	case "BP_Up6":
		return BP_Up6
	case "BP_Up7":
		return BP_Up7
	case "BP_Up8":
		return BP_Up8
	case "BP_Up9":
		return BP_Up9
	case "BP_Max":
		return BP_Max
	default:
		return -1
	}
}
