package prpc

// enum CategoryType
const (
	ACH_All           = 0
	ACH_Growup        = 1
	ACH_Battle        = 2
	ACH_Pet           = 3
	ACH_Partner       = 4
	ACH_Illustrations = 5
	ACH_Max           = 6
)

func ToName_CategoryType(id int) string {
	switch id {
	case ACH_All:
		return "ACH_All"
	case ACH_Growup:
		return "ACH_Growup"
	case ACH_Battle:
		return "ACH_Battle"
	case ACH_Pet:
		return "ACH_Pet"
	case ACH_Partner:
		return "ACH_Partner"
	case ACH_Illustrations:
		return "ACH_Illustrations"
	case ACH_Max:
		return "ACH_Max"
	default:
		return ""
	}
}
func ToId_CategoryType(name string) int {
	switch name {
	case "ACH_All":
		return ACH_All
	case "ACH_Growup":
		return ACH_Growup
	case "ACH_Battle":
		return ACH_Battle
	case "ACH_Pet":
		return ACH_Pet
	case "ACH_Partner":
		return ACH_Partner
	case "ACH_Illustrations":
		return ACH_Illustrations
	case "ACH_Max":
		return ACH_Max
	default:
		return -1
	}
}
