package prpc

// enum DiamondConfigType
const (
	DBT_None  = 0
	DBT_Day   = 1
	DBT_Week  = 2
	DBT_Month = 3
	DBT_Max   = 4
)

func ToName_DiamondConfigType(id int) string {
	switch id {
	case DBT_None:
		return "DBT_None"
	case DBT_Day:
		return "DBT_Day"
	case DBT_Week:
		return "DBT_Week"
	case DBT_Month:
		return "DBT_Month"
	case DBT_Max:
		return "DBT_Max"
	default:
		return ""
	}
}
func ToId_DiamondConfigType(name string) int {
	switch name {
	case "DBT_None":
		return DBT_None
	case "DBT_Day":
		return DBT_Day
	case "DBT_Week":
		return DBT_Week
	case "DBT_Month":
		return DBT_Month
	case "DBT_Max":
		return DBT_Max
	default:
		return -1
	}
}
