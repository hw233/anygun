package prpc

// enum NpcType
const (
	NT_None     = 0
	NT_Normal   = 1
	NT_Tongji   = 2
	NT_SinglePK = 3
	NT_TeamPK   = 4
	NT_Mogu     = 5
	NT_Xiji     = 6
	NT_Caiji1   = 7
	NT_Caiji2   = 8
	NT_Caiji3   = 9
	NT_Max      = 10
)

func ToName_NpcType(id int) string {
	switch id {
	case NT_None:
		return "NT_None"
	case NT_Normal:
		return "NT_Normal"
	case NT_Tongji:
		return "NT_Tongji"
	case NT_SinglePK:
		return "NT_SinglePK"
	case NT_TeamPK:
		return "NT_TeamPK"
	case NT_Mogu:
		return "NT_Mogu"
	case NT_Xiji:
		return "NT_Xiji"
	case NT_Caiji1:
		return "NT_Caiji1"
	case NT_Caiji2:
		return "NT_Caiji2"
	case NT_Caiji3:
		return "NT_Caiji3"
	case NT_Max:
		return "NT_Max"
	default:
		return ""
	}
}
func ToId_NpcType(name string) int {
	switch name {
	case "NT_None":
		return NT_None
	case "NT_Normal":
		return NT_Normal
	case "NT_Tongji":
		return NT_Tongji
	case "NT_SinglePK":
		return NT_SinglePK
	case "NT_TeamPK":
		return NT_TeamPK
	case "NT_Mogu":
		return NT_Mogu
	case "NT_Xiji":
		return NT_Xiji
	case "NT_Caiji1":
		return NT_Caiji1
	case "NT_Caiji2":
		return NT_Caiji2
	case "NT_Caiji3":
		return NT_Caiji3
	case "NT_Max":
		return NT_Max
	default:
		return -1
	}
}
