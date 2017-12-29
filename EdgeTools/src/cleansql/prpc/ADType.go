package prpc

// enum ADType
const (
	ADT_None              = 0
	ADT_7Days             = 1
	ADT_Cards             = 2
	ADT_ChargeTotal       = 3
	ADT_ChargeEvery       = 4
	ADT_DiscountStore     = 5
	ADT_Foundation        = 6
	ADT_LoginTotal        = 7
	ADT_OnlineReward      = 8
	ADT_HotRole           = 9
	ADT_SelfChargeTotal   = 10
	ADT_SelfChargeEvery   = 11
	ADT_SelfDiscountStore = 12
	ADT_BuyEmployee       = 13
	ADT_PhoneNumber       = 14
	ADT_Level             = 15
	ADT_Sign              = 16
	ADT_GiftBag           = 17
	ADT_Zhuanpan          = 18
	ADT_Max               = 19
)

func ToName_ADType(id int) string {
	switch id {
	case ADT_None:
		return "ADT_None"
	case ADT_7Days:
		return "ADT_7Days"
	case ADT_Cards:
		return "ADT_Cards"
	case ADT_ChargeTotal:
		return "ADT_ChargeTotal"
	case ADT_ChargeEvery:
		return "ADT_ChargeEvery"
	case ADT_DiscountStore:
		return "ADT_DiscountStore"
	case ADT_Foundation:
		return "ADT_Foundation"
	case ADT_LoginTotal:
		return "ADT_LoginTotal"
	case ADT_OnlineReward:
		return "ADT_OnlineReward"
	case ADT_HotRole:
		return "ADT_HotRole"
	case ADT_SelfChargeTotal:
		return "ADT_SelfChargeTotal"
	case ADT_SelfChargeEvery:
		return "ADT_SelfChargeEvery"
	case ADT_SelfDiscountStore:
		return "ADT_SelfDiscountStore"
	case ADT_BuyEmployee:
		return "ADT_BuyEmployee"
	case ADT_PhoneNumber:
		return "ADT_PhoneNumber"
	case ADT_Level:
		return "ADT_Level"
	case ADT_Sign:
		return "ADT_Sign"
	case ADT_GiftBag:
		return "ADT_GiftBag"
	case ADT_Zhuanpan:
		return "ADT_Zhuanpan"
	case ADT_Max:
		return "ADT_Max"
	default:
		return ""
	}
}
func ToId_ADType(name string) int {
	switch name {
	case "ADT_None":
		return ADT_None
	case "ADT_7Days":
		return ADT_7Days
	case "ADT_Cards":
		return ADT_Cards
	case "ADT_ChargeTotal":
		return ADT_ChargeTotal
	case "ADT_ChargeEvery":
		return ADT_ChargeEvery
	case "ADT_DiscountStore":
		return ADT_DiscountStore
	case "ADT_Foundation":
		return ADT_Foundation
	case "ADT_LoginTotal":
		return ADT_LoginTotal
	case "ADT_OnlineReward":
		return ADT_OnlineReward
	case "ADT_HotRole":
		return ADT_HotRole
	case "ADT_SelfChargeTotal":
		return ADT_SelfChargeTotal
	case "ADT_SelfChargeEvery":
		return ADT_SelfChargeEvery
	case "ADT_SelfDiscountStore":
		return ADT_SelfDiscountStore
	case "ADT_BuyEmployee":
		return ADT_BuyEmployee
	case "ADT_PhoneNumber":
		return ADT_PhoneNumber
	case "ADT_Level":
		return ADT_Level
	case "ADT_Sign":
		return ADT_Sign
	case "ADT_GiftBag":
		return ADT_GiftBag
	case "ADT_Zhuanpan":
		return ADT_Zhuanpan
	case "ADT_Max":
		return ADT_Max
	default:
		return -1
	}
}
