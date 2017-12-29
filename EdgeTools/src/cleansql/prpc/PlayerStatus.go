package prpc

// enum PlayerStatus
const (
	PS_Idle    = 0
	PS_Login   = 1
	PS_Game    = 2
	PS_Logout  = 3
	PS_Illegal = 4
)

func ToName_PlayerStatus(id int) string {
	switch id {
	case PS_Idle:
		return "PS_Idle"
	case PS_Login:
		return "PS_Login"
	case PS_Game:
		return "PS_Game"
	case PS_Logout:
		return "PS_Logout"
	case PS_Illegal:
		return "PS_Illegal"
	default:
		return ""
	}
}
func ToId_PlayerStatus(name string) int {
	switch name {
	case "PS_Idle":
		return PS_Idle
	case "PS_Login":
		return PS_Login
	case "PS_Game":
		return PS_Game
	case "PS_Logout":
		return PS_Logout
	case "PS_Illegal":
		return PS_Illegal
	default:
		return -1
	}
}
