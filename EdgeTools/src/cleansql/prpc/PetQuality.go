package prpc

// enum PetQuality
const (
	PE_None   = 0
	PE_White  = 1
	PE_Green  = 2
	PE_Blue   = 3
	PE_Purple = 4
	PE_Golden = 5
	PE_Orange = 6
	PE_Pink   = 7
)

func ToName_PetQuality(id int) string {
	switch id {
	case PE_None:
		return "PE_None"
	case PE_White:
		return "PE_White"
	case PE_Green:
		return "PE_Green"
	case PE_Blue:
		return "PE_Blue"
	case PE_Purple:
		return "PE_Purple"
	case PE_Golden:
		return "PE_Golden"
	case PE_Orange:
		return "PE_Orange"
	case PE_Pink:
		return "PE_Pink"
	default:
		return ""
	}
}
func ToId_PetQuality(name string) int {
	switch name {
	case "PE_None":
		return PE_None
	case "PE_White":
		return PE_White
	case "PE_Green":
		return PE_Green
	case "PE_Blue":
		return PE_Blue
	case "PE_Purple":
		return PE_Purple
	case "PE_Golden":
		return PE_Golden
	case "PE_Orange":
		return PE_Orange
	case "PE_Pink":
		return PE_Pink
	default:
		return -1
	}
}
