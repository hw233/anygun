package prpc

// enum EmployeeSkillType
const (
	EKT_GroupDamage     = 0
	EKT_DeadlyDamage    = 1
	EKT_BattleTimeLimit = 2
	EKT_Thump           = 3
	EKT_SiegeDamage     = 4
	EKT_SuperMagic      = 5
	EKT_Debuff          = 6
	EKT_PhysicalDefense = 7
	EKT_Max             = 8
)

func ToName_EmployeeSkillType(id int) string {
	switch id {
	case EKT_GroupDamage:
		return "EKT_GroupDamage"
	case EKT_DeadlyDamage:
		return "EKT_DeadlyDamage"
	case EKT_BattleTimeLimit:
		return "EKT_BattleTimeLimit"
	case EKT_Thump:
		return "EKT_Thump"
	case EKT_SiegeDamage:
		return "EKT_SiegeDamage"
	case EKT_SuperMagic:
		return "EKT_SuperMagic"
	case EKT_Debuff:
		return "EKT_Debuff"
	case EKT_PhysicalDefense:
		return "EKT_PhysicalDefense"
	case EKT_Max:
		return "EKT_Max"
	default:
		return ""
	}
}
func ToId_EmployeeSkillType(name string) int {
	switch name {
	case "EKT_GroupDamage":
		return EKT_GroupDamage
	case "EKT_DeadlyDamage":
		return EKT_DeadlyDamage
	case "EKT_BattleTimeLimit":
		return EKT_BattleTimeLimit
	case "EKT_Thump":
		return EKT_Thump
	case "EKT_SiegeDamage":
		return EKT_SiegeDamage
	case "EKT_SuperMagic":
		return EKT_SuperMagic
	case "EKT_Debuff":
		return EKT_Debuff
	case "EKT_PhysicalDefense":
		return EKT_PhysicalDefense
	case "EKT_Max":
		return EKT_Max
	default:
		return -1
	}
}
