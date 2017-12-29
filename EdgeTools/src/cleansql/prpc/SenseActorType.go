package prpc

// enum SenseActorType
const (
	SAT_Guard       = 0
	SAT_Ambassdor   = 1
	SAT_King        = 2
	SAT_Yingzi      = 3
	SAT_VillageKing = 4
	SAT_Archor      = 5
	SAT_Axe         = 6
	SAT_Sage        = 7
	SAT_Mage        = 8
	SAT_Girl        = 9
	SAT_AllMonster  = 10
)

func ToName_SenseActorType(id int) string {
	switch id {
	case SAT_Guard:
		return "SAT_Guard"
	case SAT_Ambassdor:
		return "SAT_Ambassdor"
	case SAT_King:
		return "SAT_King"
	case SAT_Yingzi:
		return "SAT_Yingzi"
	case SAT_VillageKing:
		return "SAT_VillageKing"
	case SAT_Archor:
		return "SAT_Archor"
	case SAT_Axe:
		return "SAT_Axe"
	case SAT_Sage:
		return "SAT_Sage"
	case SAT_Mage:
		return "SAT_Mage"
	case SAT_Girl:
		return "SAT_Girl"
	case SAT_AllMonster:
		return "SAT_AllMonster"
	default:
		return ""
	}
}
func ToId_SenseActorType(name string) int {
	switch name {
	case "SAT_Guard":
		return SAT_Guard
	case "SAT_Ambassdor":
		return SAT_Ambassdor
	case "SAT_King":
		return SAT_King
	case "SAT_Yingzi":
		return SAT_Yingzi
	case "SAT_VillageKing":
		return SAT_VillageKing
	case "SAT_Archor":
		return SAT_Archor
	case "SAT_Axe":
		return SAT_Axe
	case "SAT_Sage":
		return SAT_Sage
	case "SAT_Mage":
		return SAT_Mage
	case "SAT_Girl":
		return SAT_Girl
	case "SAT_AllMonster":
		return SAT_AllMonster
	default:
		return -1
	}
}
