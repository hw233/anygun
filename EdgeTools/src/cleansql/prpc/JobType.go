package prpc

// enum JobType
const (
	JT_None    = 0
	JT_Newbie  = 1
	JT_Axe     = 2
	JT_Sword   = 3
	JT_Knight  = 4
	JT_Archer  = 5
	JT_Fighter = 6
	JT_Ninja   = 7
	JT_Mage    = 8
	JT_Sage    = 9
	JT_Wizard  = 10
	JT_Word    = 11
)

func ToName_JobType(id int) string {
	switch id {
	case JT_None:
		return "JT_None"
	case JT_Newbie:
		return "JT_Newbie"
	case JT_Axe:
		return "JT_Axe"
	case JT_Sword:
		return "JT_Sword"
	case JT_Knight:
		return "JT_Knight"
	case JT_Archer:
		return "JT_Archer"
	case JT_Fighter:
		return "JT_Fighter"
	case JT_Ninja:
		return "JT_Ninja"
	case JT_Mage:
		return "JT_Mage"
	case JT_Sage:
		return "JT_Sage"
	case JT_Wizard:
		return "JT_Wizard"
	case JT_Word:
		return "JT_Word"
	default:
		return ""
	}
}
func ToId_JobType(name string) int {
	switch name {
	case "JT_None":
		return JT_None
	case "JT_Newbie":
		return JT_Newbie
	case "JT_Axe":
		return JT_Axe
	case "JT_Sword":
		return JT_Sword
	case "JT_Knight":
		return JT_Knight
	case "JT_Archer":
		return JT_Archer
	case "JT_Fighter":
		return JT_Fighter
	case "JT_Ninja":
		return JT_Ninja
	case "JT_Mage":
		return JT_Mage
	case "JT_Sage":
		return JT_Sage
	case "JT_Wizard":
		return JT_Wizard
	case "JT_Word":
		return JT_Word
	default:
		return -1
	}
}
