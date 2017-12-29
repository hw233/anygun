package prpc

// enum ItemMainType
const (
	IMT_None          = 0
	IMT_Quest         = 1
	IMT_Consumables   = 2
	IMT_Equip         = 3
	IMT_Employee      = 4
	IMT_EmployeeEquip = 5
	IMT_Debris        = 6
	IMT_FuWen         = 7
	IMT_Max           = 8
)

func ToName_ItemMainType(id int) string {
	switch id {
	case IMT_None:
		return "IMT_None"
	case IMT_Quest:
		return "IMT_Quest"
	case IMT_Consumables:
		return "IMT_Consumables"
	case IMT_Equip:
		return "IMT_Equip"
	case IMT_Employee:
		return "IMT_Employee"
	case IMT_EmployeeEquip:
		return "IMT_EmployeeEquip"
	case IMT_Debris:
		return "IMT_Debris"
	case IMT_FuWen:
		return "IMT_FuWen"
	case IMT_Max:
		return "IMT_Max"
	default:
		return ""
	}
}
func ToId_ItemMainType(name string) int {
	switch name {
	case "IMT_None":
		return IMT_None
	case "IMT_Quest":
		return IMT_Quest
	case "IMT_Consumables":
		return IMT_Consumables
	case "IMT_Equip":
		return IMT_Equip
	case "IMT_Employee":
		return IMT_Employee
	case "IMT_EmployeeEquip":
		return IMT_EmployeeEquip
	case "IMT_Debris":
		return IMT_Debris
	case "IMT_FuWen":
		return IMT_FuWen
	case "IMT_Max":
		return IMT_Max
	default:
		return -1
	}
}
