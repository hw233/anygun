package prpc

// enum OpenSubSystemFlag
const (
	OSSF_None             = 0
	OSSF_Skill            = 1
	OSSF_Baby             = 2
	OSSF_Friend           = 3
	OSSF_EmployeeGet      = 4
	OSSF_EmployeeList     = 5
	OSSF_EmployeePosition = 6
	OSSF_EmployeeEquip    = 7
	OSSF_Bar              = 8
	OSSF_Castle           = 9
	OSSF_JJC              = 10
	OSSF_PVPJJC           = 11
	OSSF_Make             = 12
	OSSF_Hundred          = 13
	OSSF_Activity         = 14
	OSSF_Handbook         = 15
	OSSF_EveryTask        = 16
	OSSF_Achieve          = 17
	OSSF_Rank             = 18
	OSSF_OnKyTalk         = 19
	OSSF_Setting          = 20
	OSSF_Shop             = 21
	OSSF_DoubleExp        = 22
	OSSF_Family           = 23
	OSSF_AuctionHouse     = 24
	OSSF_BabyLeranSkill   = 25
	OSSF_MagicItem        = 26
	OSSF_EmployeePos10    = 27
	OSSF_EmployeePos15    = 28
	OSSF_EmployeePos20    = 29
	OSSF_Guid             = 30
	OSSF_Team             = 31
	OSSF_choujiang        = 32
	OSSF_Max              = 33
)

func ToName_OpenSubSystemFlag(id int) string {
	switch id {
	case OSSF_None:
		return "OSSF_None"
	case OSSF_Skill:
		return "OSSF_Skill"
	case OSSF_Baby:
		return "OSSF_Baby"
	case OSSF_Friend:
		return "OSSF_Friend"
	case OSSF_EmployeeGet:
		return "OSSF_EmployeeGet"
	case OSSF_EmployeeList:
		return "OSSF_EmployeeList"
	case OSSF_EmployeePosition:
		return "OSSF_EmployeePosition"
	case OSSF_EmployeeEquip:
		return "OSSF_EmployeeEquip"
	case OSSF_Bar:
		return "OSSF_Bar"
	case OSSF_Castle:
		return "OSSF_Castle"
	case OSSF_JJC:
		return "OSSF_JJC"
	case OSSF_PVPJJC:
		return "OSSF_PVPJJC"
	case OSSF_Make:
		return "OSSF_Make"
	case OSSF_Hundred:
		return "OSSF_Hundred"
	case OSSF_Activity:
		return "OSSF_Activity"
	case OSSF_Handbook:
		return "OSSF_Handbook"
	case OSSF_EveryTask:
		return "OSSF_EveryTask"
	case OSSF_Achieve:
		return "OSSF_Achieve"
	case OSSF_Rank:
		return "OSSF_Rank"
	case OSSF_OnKyTalk:
		return "OSSF_OnKyTalk"
	case OSSF_Setting:
		return "OSSF_Setting"
	case OSSF_Shop:
		return "OSSF_Shop"
	case OSSF_DoubleExp:
		return "OSSF_DoubleExp"
	case OSSF_Family:
		return "OSSF_Family"
	case OSSF_AuctionHouse:
		return "OSSF_AuctionHouse"
	case OSSF_BabyLeranSkill:
		return "OSSF_BabyLeranSkill"
	case OSSF_MagicItem:
		return "OSSF_MagicItem"
	case OSSF_EmployeePos10:
		return "OSSF_EmployeePos10"
	case OSSF_EmployeePos15:
		return "OSSF_EmployeePos15"
	case OSSF_EmployeePos20:
		return "OSSF_EmployeePos20"
	case OSSF_Guid:
		return "OSSF_Guid"
	case OSSF_Team:
		return "OSSF_Team"
	case OSSF_choujiang:
		return "OSSF_choujiang"
	case OSSF_Max:
		return "OSSF_Max"
	default:
		return ""
	}
}
func ToId_OpenSubSystemFlag(name string) int {
	switch name {
	case "OSSF_None":
		return OSSF_None
	case "OSSF_Skill":
		return OSSF_Skill
	case "OSSF_Baby":
		return OSSF_Baby
	case "OSSF_Friend":
		return OSSF_Friend
	case "OSSF_EmployeeGet":
		return OSSF_EmployeeGet
	case "OSSF_EmployeeList":
		return OSSF_EmployeeList
	case "OSSF_EmployeePosition":
		return OSSF_EmployeePosition
	case "OSSF_EmployeeEquip":
		return OSSF_EmployeeEquip
	case "OSSF_Bar":
		return OSSF_Bar
	case "OSSF_Castle":
		return OSSF_Castle
	case "OSSF_JJC":
		return OSSF_JJC
	case "OSSF_PVPJJC":
		return OSSF_PVPJJC
	case "OSSF_Make":
		return OSSF_Make
	case "OSSF_Hundred":
		return OSSF_Hundred
	case "OSSF_Activity":
		return OSSF_Activity
	case "OSSF_Handbook":
		return OSSF_Handbook
	case "OSSF_EveryTask":
		return OSSF_EveryTask
	case "OSSF_Achieve":
		return OSSF_Achieve
	case "OSSF_Rank":
		return OSSF_Rank
	case "OSSF_OnKyTalk":
		return OSSF_OnKyTalk
	case "OSSF_Setting":
		return OSSF_Setting
	case "OSSF_Shop":
		return OSSF_Shop
	case "OSSF_DoubleExp":
		return OSSF_DoubleExp
	case "OSSF_Family":
		return OSSF_Family
	case "OSSF_AuctionHouse":
		return OSSF_AuctionHouse
	case "OSSF_BabyLeranSkill":
		return OSSF_BabyLeranSkill
	case "OSSF_MagicItem":
		return OSSF_MagicItem
	case "OSSF_EmployeePos10":
		return OSSF_EmployeePos10
	case "OSSF_EmployeePos15":
		return OSSF_EmployeePos15
	case "OSSF_EmployeePos20":
		return OSSF_EmployeePos20
	case "OSSF_Guid":
		return OSSF_Guid
	case "OSSF_Team":
		return OSSF_Team
	case "OSSF_choujiang":
		return OSSF_choujiang
	case "OSSF_Max":
		return OSSF_Max
	default:
		return -1
	}
}
