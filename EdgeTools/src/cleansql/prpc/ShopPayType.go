package prpc

// enum ShopPayType
const (
	SPT_Nil           = 0
	SPT_RMB           = 1
	SPT_Diamond       = 2
	SPT_MagicCurrency = 3
	SPT_Gold          = 4
)

func ToName_ShopPayType(id int) string {
	switch id {
	case SPT_Nil:
		return "SPT_Nil"
	case SPT_RMB:
		return "SPT_RMB"
	case SPT_Diamond:
		return "SPT_Diamond"
	case SPT_MagicCurrency:
		return "SPT_MagicCurrency"
	case SPT_Gold:
		return "SPT_Gold"
	default:
		return ""
	}
}
func ToId_ShopPayType(name string) int {
	switch name {
	case "SPT_Nil":
		return SPT_Nil
	case "SPT_RMB":
		return SPT_RMB
	case "SPT_Diamond":
		return SPT_Diamond
	case "SPT_MagicCurrency":
		return SPT_MagicCurrency
	case "SPT_Gold":
		return SPT_Gold
	default:
		return -1
	}
}
