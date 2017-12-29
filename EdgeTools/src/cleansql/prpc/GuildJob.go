package prpc

// enum GuildJob
const (
	GJ_None          = 0
	GJ_People        = 1
	GJ_Minister      = 2
	GJ_SecretaryHead = 3
	GJ_VicePremier   = 4
	GJ_Premier       = 5
	GJ_Max           = 6
)

func ToName_GuildJob(id int) string {
	switch id {
	case GJ_None:
		return "GJ_None"
	case GJ_People:
		return "GJ_People"
	case GJ_Minister:
		return "GJ_Minister"
	case GJ_SecretaryHead:
		return "GJ_SecretaryHead"
	case GJ_VicePremier:
		return "GJ_VicePremier"
	case GJ_Premier:
		return "GJ_Premier"
	case GJ_Max:
		return "GJ_Max"
	default:
		return ""
	}
}
func ToId_GuildJob(name string) int {
	switch name {
	case "GJ_None":
		return GJ_None
	case "GJ_People":
		return GJ_People
	case "GJ_Minister":
		return GJ_Minister
	case "GJ_SecretaryHead":
		return GJ_SecretaryHead
	case "GJ_VicePremier":
		return GJ_VicePremier
	case "GJ_Premier":
		return GJ_Premier
	case "GJ_Max":
		return GJ_Max
	default:
		return -1
	}
}
