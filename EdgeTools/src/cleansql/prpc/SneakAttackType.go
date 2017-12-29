package prpc

// enum SneakAttackType
const (
	SAT_None           = 0
	SAT_SneakAttack    = 1
	SAT_SurpriseAttack = 2
	SAT_Max            = 3
)

func ToName_SneakAttackType(id int) string {
	switch id {
	case SAT_None:
		return "SAT_None"
	case SAT_SneakAttack:
		return "SAT_SneakAttack"
	case SAT_SurpriseAttack:
		return "SAT_SurpriseAttack"
	case SAT_Max:
		return "SAT_Max"
	default:
		return ""
	}
}
func ToId_SneakAttackType(name string) int {
	switch name {
	case "SAT_None":
		return SAT_None
	case "SAT_SneakAttack":
		return SAT_SneakAttack
	case "SAT_SurpriseAttack":
		return SAT_SurpriseAttack
	case "SAT_Max":
		return SAT_Max
	default:
		return -1
	}
}
