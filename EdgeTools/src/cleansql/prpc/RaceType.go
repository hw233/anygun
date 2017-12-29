package prpc

// enum RaceType
const (
	RT_None   = 0
	RT_Human  = 1
	RT_Insect = 2
	RT_Plant  = 3
	RT_Extra  = 4
	RT_Dragon = 5
	RT_Animal = 6
	RT_Fly    = 7
	RT_Undead = 8
	RT_Metal  = 9
	RT_Max    = 10
)

func ToName_RaceType(id int) string {
	switch id {
	case RT_None:
		return "RT_None"
	case RT_Human:
		return "RT_Human"
	case RT_Insect:
		return "RT_Insect"
	case RT_Plant:
		return "RT_Plant"
	case RT_Extra:
		return "RT_Extra"
	case RT_Dragon:
		return "RT_Dragon"
	case RT_Animal:
		return "RT_Animal"
	case RT_Fly:
		return "RT_Fly"
	case RT_Undead:
		return "RT_Undead"
	case RT_Metal:
		return "RT_Metal"
	case RT_Max:
		return "RT_Max"
	default:
		return ""
	}
}
func ToId_RaceType(name string) int {
	switch name {
	case "RT_None":
		return RT_None
	case "RT_Human":
		return RT_Human
	case "RT_Insect":
		return RT_Insect
	case "RT_Plant":
		return RT_Plant
	case "RT_Extra":
		return RT_Extra
	case "RT_Dragon":
		return RT_Dragon
	case "RT_Animal":
		return RT_Animal
	case "RT_Fly":
		return RT_Fly
	case "RT_Undead":
		return RT_Undead
	case "RT_Metal":
		return RT_Metal
	case "RT_Max":
		return RT_Max
	default:
		return -1
	}
}
