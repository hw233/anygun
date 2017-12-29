package prpc

// enum ChatKind
const (
	CK_None   = 0
	CK_World  = 1
	CK_Team   = 2
	CK_System = 3
	CK_Friend = 4
	CK_GM     = 5
	CK_Guild  = 6
	CK_Max    = 7
)

func ToName_ChatKind(id int) string {
	switch id {
	case CK_None:
		return "CK_None"
	case CK_World:
		return "CK_World"
	case CK_Team:
		return "CK_Team"
	case CK_System:
		return "CK_System"
	case CK_Friend:
		return "CK_Friend"
	case CK_GM:
		return "CK_GM"
	case CK_Guild:
		return "CK_Guild"
	case CK_Max:
		return "CK_Max"
	default:
		return ""
	}
}
func ToId_ChatKind(name string) int {
	switch name {
	case "CK_None":
		return CK_None
	case "CK_World":
		return CK_World
	case "CK_Team":
		return CK_Team
	case "CK_System":
		return CK_System
	case "CK_Friend":
		return CK_Friend
	case "CK_GM":
		return CK_GM
	case "CK_Guild":
		return CK_Guild
	case "CK_Max":
		return CK_Max
	default:
		return -1
	}
}
