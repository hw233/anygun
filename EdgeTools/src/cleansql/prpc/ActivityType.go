package prpc

// enum ActivityType
const (
	ACT_None      = 0
	ACT_Tongji    = 1
	ACT_Mogu      = 2
	ACT_Richang   = 3
	ACT_Pet       = 4
	ACT_AlonePK   = 5
	ACT_TeamPK    = 6
	ACT_JJC       = 7
	ACT_Challenge = 8
	ACT_Exam      = 9
	ACT_Copy      = 10
	ACT_Xuyuan    = 11
	ACT_Family_0  = 12
	ACT_Family_1  = 13
	ACT_Family_2  = 14
	ACT_Family_3  = 15
	ACT_Family_4  = 16
	ACT_Warrior   = 17
	ACT_Money     = 18
	ACT_Rand      = 19
	ACT_Max       = 20
)

func ToName_ActivityType(id int) string {
	switch id {
	case ACT_None:
		return "ACT_None"
	case ACT_Tongji:
		return "ACT_Tongji"
	case ACT_Mogu:
		return "ACT_Mogu"
	case ACT_Richang:
		return "ACT_Richang"
	case ACT_Pet:
		return "ACT_Pet"
	case ACT_AlonePK:
		return "ACT_AlonePK"
	case ACT_TeamPK:
		return "ACT_TeamPK"
	case ACT_JJC:
		return "ACT_JJC"
	case ACT_Challenge:
		return "ACT_Challenge"
	case ACT_Exam:
		return "ACT_Exam"
	case ACT_Copy:
		return "ACT_Copy"
	case ACT_Xuyuan:
		return "ACT_Xuyuan"
	case ACT_Family_0:
		return "ACT_Family_0"
	case ACT_Family_1:
		return "ACT_Family_1"
	case ACT_Family_2:
		return "ACT_Family_2"
	case ACT_Family_3:
		return "ACT_Family_3"
	case ACT_Family_4:
		return "ACT_Family_4"
	case ACT_Warrior:
		return "ACT_Warrior"
	case ACT_Money:
		return "ACT_Money"
	case ACT_Rand:
		return "ACT_Rand"
	case ACT_Max:
		return "ACT_Max"
	default:
		return ""
	}
}
func ToId_ActivityType(name string) int {
	switch name {
	case "ACT_None":
		return ACT_None
	case "ACT_Tongji":
		return ACT_Tongji
	case "ACT_Mogu":
		return ACT_Mogu
	case "ACT_Richang":
		return ACT_Richang
	case "ACT_Pet":
		return ACT_Pet
	case "ACT_AlonePK":
		return ACT_AlonePK
	case "ACT_TeamPK":
		return ACT_TeamPK
	case "ACT_JJC":
		return ACT_JJC
	case "ACT_Challenge":
		return ACT_Challenge
	case "ACT_Exam":
		return ACT_Exam
	case "ACT_Copy":
		return ACT_Copy
	case "ACT_Xuyuan":
		return ACT_Xuyuan
	case "ACT_Family_0":
		return ACT_Family_0
	case "ACT_Family_1":
		return ACT_Family_1
	case "ACT_Family_2":
		return ACT_Family_2
	case "ACT_Family_3":
		return ACT_Family_3
	case "ACT_Family_4":
		return ACT_Family_4
	case "ACT_Warrior":
		return ACT_Warrior
	case "ACT_Money":
		return ACT_Money
	case "ACT_Rand":
		return ACT_Rand
	case "ACT_Max":
		return ACT_Max
	default:
		return -1
	}
}
