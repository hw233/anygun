package prpc

// enum UIBehaviorType
const (
	UBT_None            = 0
	UBT_Bag             = 1
	UBT_Baby            = 2
	UBT_Employee        = 3
	UBT_SkillView       = 4
	UBT_SkillLearn      = 5
	UBT_Task            = 6
	UBT_Make            = 7
	UBT_Gather          = 8
	UBT_MagicItem       = 9
	UBT_Store           = 10
	UBT_Help            = 11
	UBT_Friend          = 12
	UBT_Email           = 13
	UBT_Auction         = 14
	UBT_Activity        = 15
	UBT_SignUp          = 16
	UBT_EmployessList   = 17
	UBT_EmployessPos    = 18
	UBT_EmployessTavern = 19
	UBT_PlayerProperty  = 20
	UBT_Max             = 21
)

func ToName_UIBehaviorType(id int) string {
	switch id {
	case UBT_None:
		return "UBT_None"
	case UBT_Bag:
		return "UBT_Bag"
	case UBT_Baby:
		return "UBT_Baby"
	case UBT_Employee:
		return "UBT_Employee"
	case UBT_SkillView:
		return "UBT_SkillView"
	case UBT_SkillLearn:
		return "UBT_SkillLearn"
	case UBT_Task:
		return "UBT_Task"
	case UBT_Make:
		return "UBT_Make"
	case UBT_Gather:
		return "UBT_Gather"
	case UBT_MagicItem:
		return "UBT_MagicItem"
	case UBT_Store:
		return "UBT_Store"
	case UBT_Help:
		return "UBT_Help"
	case UBT_Friend:
		return "UBT_Friend"
	case UBT_Email:
		return "UBT_Email"
	case UBT_Auction:
		return "UBT_Auction"
	case UBT_Activity:
		return "UBT_Activity"
	case UBT_SignUp:
		return "UBT_SignUp"
	case UBT_EmployessList:
		return "UBT_EmployessList"
	case UBT_EmployessPos:
		return "UBT_EmployessPos"
	case UBT_EmployessTavern:
		return "UBT_EmployessTavern"
	case UBT_PlayerProperty:
		return "UBT_PlayerProperty"
	case UBT_Max:
		return "UBT_Max"
	default:
		return ""
	}
}
func ToId_UIBehaviorType(name string) int {
	switch name {
	case "UBT_None":
		return UBT_None
	case "UBT_Bag":
		return UBT_Bag
	case "UBT_Baby":
		return UBT_Baby
	case "UBT_Employee":
		return UBT_Employee
	case "UBT_SkillView":
		return UBT_SkillView
	case "UBT_SkillLearn":
		return UBT_SkillLearn
	case "UBT_Task":
		return UBT_Task
	case "UBT_Make":
		return UBT_Make
	case "UBT_Gather":
		return UBT_Gather
	case "UBT_MagicItem":
		return UBT_MagicItem
	case "UBT_Store":
		return UBT_Store
	case "UBT_Help":
		return UBT_Help
	case "UBT_Friend":
		return UBT_Friend
	case "UBT_Email":
		return UBT_Email
	case "UBT_Auction":
		return UBT_Auction
	case "UBT_Activity":
		return UBT_Activity
	case "UBT_SignUp":
		return UBT_SignUp
	case "UBT_EmployessList":
		return UBT_EmployessList
	case "UBT_EmployessPos":
		return UBT_EmployessPos
	case "UBT_EmployessTavern":
		return UBT_EmployessTavern
	case "UBT_PlayerProperty":
		return UBT_PlayerProperty
	case "UBT_Max":
		return UBT_Max
	default:
		return -1
	}
}
