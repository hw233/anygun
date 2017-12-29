package prpc

// enum ClassifyType
const (
	SD_Debris  = 0
	SD_Data    = 1
	SD_Pet     = 2
	SD_Fashion = 3
	SD_Diamond = 4
)

func ToName_ClassifyType(id int) string {
	switch id {
	case SD_Debris:
		return "SD_Debris"
	case SD_Data:
		return "SD_Data"
	case SD_Pet:
		return "SD_Pet"
	case SD_Fashion:
		return "SD_Fashion"
	case SD_Diamond:
		return "SD_Diamond"
	default:
		return ""
	}
}
func ToId_ClassifyType(name string) int {
	switch name {
	case "SD_Debris":
		return SD_Debris
	case "SD_Data":
		return SD_Data
	case "SD_Pet":
		return SD_Pet
	case "SD_Fashion":
		return SD_Fashion
	case "SD_Diamond":
		return SD_Diamond
	default:
		return -1
	}
}
