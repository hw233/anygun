package prpc

// enum MineType
const (
	MT_None   = 0
	MT_JinShu = 1
	MT_MuCai  = 2
	MT_BuLiao = 3
	MT_Max    = 4
)

func ToName_MineType(id int) string {
	switch id {
	case MT_None:
		return "MT_None"
	case MT_JinShu:
		return "MT_JinShu"
	case MT_MuCai:
		return "MT_MuCai"
	case MT_BuLiao:
		return "MT_BuLiao"
	case MT_Max:
		return "MT_Max"
	default:
		return ""
	}
}
func ToId_MineType(name string) int {
	switch name {
	case "MT_None":
		return MT_None
	case "MT_JinShu":
		return MT_JinShu
	case "MT_MuCai":
		return MT_MuCai
	case "MT_BuLiao":
		return MT_BuLiao
	case "MT_Max":
		return MT_Max
	default:
		return -1
	}
}
