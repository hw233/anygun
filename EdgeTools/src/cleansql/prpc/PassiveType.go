package prpc

// enum PassiveType
const (
	PAT_None       = 0
	PAT_Buff       = 1
	PAT_Deference  = 2
	PAT_Dodge      = 3
	PAT_Counter    = 4
	PAT_Change     = 5
	PAT_Guard      = 6
	PAT_Runaway    = 7
	PAT_BabyInnout = 8
	PAT_SecKill    = 9
)

func ToName_PassiveType(id int) string {
	switch id {
	case PAT_None:
		return "PAT_None"
	case PAT_Buff:
		return "PAT_Buff"
	case PAT_Deference:
		return "PAT_Deference"
	case PAT_Dodge:
		return "PAT_Dodge"
	case PAT_Counter:
		return "PAT_Counter"
	case PAT_Change:
		return "PAT_Change"
	case PAT_Guard:
		return "PAT_Guard"
	case PAT_Runaway:
		return "PAT_Runaway"
	case PAT_BabyInnout:
		return "PAT_BabyInnout"
	case PAT_SecKill:
		return "PAT_SecKill"
	default:
		return ""
	}
}
func ToId_PassiveType(name string) int {
	switch name {
	case "PAT_None":
		return PAT_None
	case "PAT_Buff":
		return PAT_Buff
	case "PAT_Deference":
		return PAT_Deference
	case "PAT_Dodge":
		return PAT_Dodge
	case "PAT_Counter":
		return PAT_Counter
	case "PAT_Change":
		return PAT_Change
	case "PAT_Guard":
		return PAT_Guard
	case "PAT_Runaway":
		return PAT_Runaway
	case "PAT_BabyInnout":
		return PAT_BabyInnout
	case "PAT_SecKill":
		return PAT_SecKill
	default:
		return -1
	}
}
