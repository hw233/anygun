package prpc

// enum PatchVersion
const (
	Patch_0     = 0
	Patch_1     = 1
	Patch_2     = 2
	Patch_3     = 3
	Patch_4     = 4
	Patch_5     = 5
	PatchNumber = 6
)

func ToName_PatchVersion(id int) string {
	switch id {
	case Patch_0:
		return "Patch_0"
	case Patch_1:
		return "Patch_1"
	case Patch_2:
		return "Patch_2"
	case Patch_3:
		return "Patch_3"
	case Patch_4:
		return "Patch_4"
	case Patch_5:
		return "Patch_5"
	case PatchNumber:
		return "PatchNumber"
	default:
		return ""
	}
}
func ToId_PatchVersion(name string) int {
	switch name {
	case "Patch_0":
		return Patch_0
	case "Patch_1":
		return Patch_1
	case "Patch_2":
		return Patch_2
	case "Patch_3":
		return Patch_3
	case "Patch_4":
		return Patch_4
	case "Patch_5":
		return Patch_5
	case "PatchNumber":
		return PatchNumber
	default:
		return -1
	}
}
