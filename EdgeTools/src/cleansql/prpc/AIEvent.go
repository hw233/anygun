package prpc

// enum AIEvent
const (
	ME_None     = 0
	ME_Born     = 1
	ME_Deadth   = 2
	ME_AttackGo = 3
	ME_SkillGO  = 4
	ME_Update   = 5
	ME_Max      = 6
)

func ToName_AIEvent(id int) string {
	switch id {
	case ME_None:
		return "ME_None"
	case ME_Born:
		return "ME_Born"
	case ME_Deadth:
		return "ME_Deadth"
	case ME_AttackGo:
		return "ME_AttackGo"
	case ME_SkillGO:
		return "ME_SkillGO"
	case ME_Update:
		return "ME_Update"
	case ME_Max:
		return "ME_Max"
	default:
		return ""
	}
}
func ToId_AIEvent(name string) int {
	switch name {
	case "ME_None":
		return ME_None
	case "ME_Born":
		return ME_Born
	case "ME_Deadth":
		return ME_Deadth
	case "ME_AttackGo":
		return ME_AttackGo
	case "ME_SkillGO":
		return ME_SkillGO
	case "ME_Update":
		return ME_Update
	case "ME_Max":
		return ME_Max
	default:
		return -1
	}
}
