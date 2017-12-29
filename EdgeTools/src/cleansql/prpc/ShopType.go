package prpc

// enum ShopType
const (
	SIT_None          = 0
	SIT_FirstRecharge = 1
	SIT_Recharge      = 2
	SIT_Shop          = 3
	SIT_EmployeeShop  = 4
	SIT_VIP           = 5
	SIT_Fund          = 6
	SIT_Giftbag       = 7
	SIT_Max           = 8
)

func ToName_ShopType(id int) string {
	switch id {
	case SIT_None:
		return "SIT_None"
	case SIT_FirstRecharge:
		return "SIT_FirstRecharge"
	case SIT_Recharge:
		return "SIT_Recharge"
	case SIT_Shop:
		return "SIT_Shop"
	case SIT_EmployeeShop:
		return "SIT_EmployeeShop"
	case SIT_VIP:
		return "SIT_VIP"
	case SIT_Fund:
		return "SIT_Fund"
	case SIT_Giftbag:
		return "SIT_Giftbag"
	case SIT_Max:
		return "SIT_Max"
	default:
		return ""
	}
}
func ToId_ShopType(name string) int {
	switch name {
	case "SIT_None":
		return SIT_None
	case "SIT_FirstRecharge":
		return SIT_FirstRecharge
	case "SIT_Recharge":
		return SIT_Recharge
	case "SIT_Shop":
		return SIT_Shop
	case "SIT_EmployeeShop":
		return SIT_EmployeeShop
	case "SIT_VIP":
		return SIT_VIP
	case "SIT_Fund":
		return SIT_Fund
	case "SIT_Giftbag":
		return SIT_Giftbag
	case "SIT_Max":
		return SIT_Max
	default:
		return -1
	}
}
