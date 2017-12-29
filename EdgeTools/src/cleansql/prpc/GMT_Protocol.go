package prpc

// enum GMT_Protocol
const (
	GMT_None            = 0
	GMT_GMCommand       = 1
	GMT_Notice          = 2
	GMT_InsertMail      = 3
	GMT_QueryPlayer     = 4
	GMT_LoginActivity   = 5
	GMT_7Days           = 6
	GMT_Cards           = 7
	GMT_ExtractEmployee = 8
	GMT_ChargeTotal     = 9
	GMT_ChargeEvery     = 10
	GMT_DiscountStore   = 11
	GMT_Foundation      = 12
	GMT_LoginTotal      = 13
	GMT_OnlineReward    = 14
	GMT_HotRole         = 15
	GMT_SelfChargeTotal = 16
	GMT_SelfChargeEvery = 17
	GMT_MinGiftBag      = 18
	GMT_DoScript        = 19
	GNT_MakeOrder       = 20
	GMT_QueryRoleList   = 21
	GMT_QueryRMB        = 22
	GMT_QueryDia        = 23
	GMT_QueryMoney      = 24
	GMT_Zhuanpan        = 25
	GMT_MAX             = 26
)

func ToName_GMT_Protocol(id int) string {
	switch id {
	case GMT_None:
		return "GMT_None"
	case GMT_GMCommand:
		return "GMT_GMCommand"
	case GMT_Notice:
		return "GMT_Notice"
	case GMT_InsertMail:
		return "GMT_InsertMail"
	case GMT_QueryPlayer:
		return "GMT_QueryPlayer"
	case GMT_LoginActivity:
		return "GMT_LoginActivity"
	case GMT_7Days:
		return "GMT_7Days"
	case GMT_Cards:
		return "GMT_Cards"
	case GMT_ExtractEmployee:
		return "GMT_ExtractEmployee"
	case GMT_ChargeTotal:
		return "GMT_ChargeTotal"
	case GMT_ChargeEvery:
		return "GMT_ChargeEvery"
	case GMT_DiscountStore:
		return "GMT_DiscountStore"
	case GMT_Foundation:
		return "GMT_Foundation"
	case GMT_LoginTotal:
		return "GMT_LoginTotal"
	case GMT_OnlineReward:
		return "GMT_OnlineReward"
	case GMT_HotRole:
		return "GMT_HotRole"
	case GMT_SelfChargeTotal:
		return "GMT_SelfChargeTotal"
	case GMT_SelfChargeEvery:
		return "GMT_SelfChargeEvery"
	case GMT_MinGiftBag:
		return "GMT_MinGiftBag"
	case GMT_DoScript:
		return "GMT_DoScript"
	case GNT_MakeOrder:
		return "GNT_MakeOrder"
	case GMT_QueryRoleList:
		return "GMT_QueryRoleList"
	case GMT_QueryRMB:
		return "GMT_QueryRMB"
	case GMT_QueryDia:
		return "GMT_QueryDia"
	case GMT_QueryMoney:
		return "GMT_QueryMoney"
	case GMT_Zhuanpan:
		return "GMT_Zhuanpan"
	case GMT_MAX:
		return "GMT_MAX"
	default:
		return ""
	}
}
func ToId_GMT_Protocol(name string) int {
	switch name {
	case "GMT_None":
		return GMT_None
	case "GMT_GMCommand":
		return GMT_GMCommand
	case "GMT_Notice":
		return GMT_Notice
	case "GMT_InsertMail":
		return GMT_InsertMail
	case "GMT_QueryPlayer":
		return GMT_QueryPlayer
	case "GMT_LoginActivity":
		return GMT_LoginActivity
	case "GMT_7Days":
		return GMT_7Days
	case "GMT_Cards":
		return GMT_Cards
	case "GMT_ExtractEmployee":
		return GMT_ExtractEmployee
	case "GMT_ChargeTotal":
		return GMT_ChargeTotal
	case "GMT_ChargeEvery":
		return GMT_ChargeEvery
	case "GMT_DiscountStore":
		return GMT_DiscountStore
	case "GMT_Foundation":
		return GMT_Foundation
	case "GMT_LoginTotal":
		return GMT_LoginTotal
	case "GMT_OnlineReward":
		return GMT_OnlineReward
	case "GMT_HotRole":
		return GMT_HotRole
	case "GMT_SelfChargeTotal":
		return GMT_SelfChargeTotal
	case "GMT_SelfChargeEvery":
		return GMT_SelfChargeEvery
	case "GMT_MinGiftBag":
		return GMT_MinGiftBag
	case "GMT_DoScript":
		return GMT_DoScript
	case "GNT_MakeOrder":
		return GNT_MakeOrder
	case "GMT_QueryRoleList":
		return GMT_QueryRoleList
	case "GMT_QueryRMB":
		return GMT_QueryRMB
	case "GMT_QueryDia":
		return GMT_QueryDia
	case "GMT_QueryMoney":
		return GMT_QueryMoney
	case "GMT_Zhuanpan":
		return GMT_Zhuanpan
	case "GMT_MAX":
		return GMT_MAX
	default:
		return -1
	}
}
