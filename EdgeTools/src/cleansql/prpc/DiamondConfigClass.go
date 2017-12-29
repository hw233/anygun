package prpc

// enum DiamondConfigClass
const (
	DBT_Type_None          = 0
	DBT_Type_Mine_Famu     = 1
	DBT_Type_Mine_Caikuang = 2
	DBT_Type_Mine_Zhibu    = 3
	DBT_Type_Max           = 4
)

func ToName_DiamondConfigClass(id int) string {
	switch id {
	case DBT_Type_None:
		return "DBT_Type_None"
	case DBT_Type_Mine_Famu:
		return "DBT_Type_Mine_Famu"
	case DBT_Type_Mine_Caikuang:
		return "DBT_Type_Mine_Caikuang"
	case DBT_Type_Mine_Zhibu:
		return "DBT_Type_Mine_Zhibu"
	case DBT_Type_Max:
		return "DBT_Type_Max"
	default:
		return ""
	}
}
func ToId_DiamondConfigClass(name string) int {
	switch name {
	case "DBT_Type_None":
		return DBT_Type_None
	case "DBT_Type_Mine_Famu":
		return DBT_Type_Mine_Famu
	case "DBT_Type_Mine_Caikuang":
		return DBT_Type_Mine_Caikuang
	case "DBT_Type_Mine_Zhibu":
		return DBT_Type_Mine_Zhibu
	case "DBT_Type_Max":
		return DBT_Type_Max
	default:
		return -1
	}
}
