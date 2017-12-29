package prpc

// enum TeamType
const (
	TT_None        = 0
	TT_MainQuest   = 1
	TT_TongjiQuest = 2
	TT_Daochang    = 3
	TT_CaoMoGu     = 4
	TT_Zhanchang   = 5
	TT_Hero        = 6
	TT_Pet         = 7
	TT_JJC         = 8
	TT_ShuaGuai    = 9
	TT_Copy        = 10
	TT_Max         = 11
)

func ToName_TeamType(id int) string {
	switch id {
	case TT_None:
		return "TT_None"
	case TT_MainQuest:
		return "TT_MainQuest"
	case TT_TongjiQuest:
		return "TT_TongjiQuest"
	case TT_Daochang:
		return "TT_Daochang"
	case TT_CaoMoGu:
		return "TT_CaoMoGu"
	case TT_Zhanchang:
		return "TT_Zhanchang"
	case TT_Hero:
		return "TT_Hero"
	case TT_Pet:
		return "TT_Pet"
	case TT_JJC:
		return "TT_JJC"
	case TT_ShuaGuai:
		return "TT_ShuaGuai"
	case TT_Copy:
		return "TT_Copy"
	case TT_Max:
		return "TT_Max"
	default:
		return ""
	}
}
func ToId_TeamType(name string) int {
	switch name {
	case "TT_None":
		return TT_None
	case "TT_MainQuest":
		return TT_MainQuest
	case "TT_TongjiQuest":
		return TT_TongjiQuest
	case "TT_Daochang":
		return TT_Daochang
	case "TT_CaoMoGu":
		return TT_CaoMoGu
	case "TT_Zhanchang":
		return TT_Zhanchang
	case "TT_Hero":
		return TT_Hero
	case "TT_Pet":
		return TT_Pet
	case "TT_JJC":
		return TT_JJC
	case "TT_ShuaGuai":
		return TT_ShuaGuai
	case "TT_Copy":
		return TT_Copy
	case "TT_Max":
		return TT_Max
	default:
		return -1
	}
}
