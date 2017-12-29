package prpc

// enum StorageType
const (
	ST_Item = 0
	ST_Baby = 1
)

func ToName_StorageType(id int) string {
	switch id {
	case ST_Item:
		return "ST_Item"
	case ST_Baby:
		return "ST_Baby"
	default:
		return ""
	}
}
func ToId_StorageType(name string) int {
	switch name {
	case "ST_Item":
		return ST_Item
	case "ST_Baby":
		return ST_Baby
	default:
		return -1
	}
}
