package prpc

// enum ModifyListFlag
const (
	MLF_Add                   = 0
	MLF_Delete                = 1
	MLF_ChangeOnline          = 2
	MLF_ChangeOffline         = 3
	MLF_ChangePosition        = 4
	MLF_ChangeContribution    = 5
	MLF_ChangeLevel           = 6
	MLF_ChangeProfession      = 7
	MLF_ChangeGuuildBattleCon = 8
)

func ToName_ModifyListFlag(id int) string {
	switch id {
	case MLF_Add:
		return "MLF_Add"
	case MLF_Delete:
		return "MLF_Delete"
	case MLF_ChangeOnline:
		return "MLF_ChangeOnline"
	case MLF_ChangeOffline:
		return "MLF_ChangeOffline"
	case MLF_ChangePosition:
		return "MLF_ChangePosition"
	case MLF_ChangeContribution:
		return "MLF_ChangeContribution"
	case MLF_ChangeLevel:
		return "MLF_ChangeLevel"
	case MLF_ChangeProfession:
		return "MLF_ChangeProfession"
	case MLF_ChangeGuuildBattleCon:
		return "MLF_ChangeGuuildBattleCon"
	default:
		return ""
	}
}
func ToId_ModifyListFlag(name string) int {
	switch name {
	case "MLF_Add":
		return MLF_Add
	case "MLF_Delete":
		return MLF_Delete
	case "MLF_ChangeOnline":
		return MLF_ChangeOnline
	case "MLF_ChangeOffline":
		return MLF_ChangeOffline
	case "MLF_ChangePosition":
		return MLF_ChangePosition
	case "MLF_ChangeContribution":
		return MLF_ChangeContribution
	case "MLF_ChangeLevel":
		return MLF_ChangeLevel
	case "MLF_ChangeProfession":
		return MLF_ChangeProfession
	case "MLF_ChangeGuuildBattleCon":
		return MLF_ChangeGuuildBattleCon
	default:
		return -1
	}
}
