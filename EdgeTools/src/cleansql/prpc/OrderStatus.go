package prpc

// enum OrderStatus
const (
	OS_None       = 0
	OS_ActiveOk   = 1
	OS_RunawayOk  = 2
	OS_FangBaobao = 3
	OS_ShouBaobao = 4
	OS_Weapon     = 5
	OS_Zhuachong  = 6
	OS_Flee       = 7
	OS_Summon     = 8
)

func ToName_OrderStatus(id int) string {
	switch id {
	case OS_None:
		return "OS_None"
	case OS_ActiveOk:
		return "OS_ActiveOk"
	case OS_RunawayOk:
		return "OS_RunawayOk"
	case OS_FangBaobao:
		return "OS_FangBaobao"
	case OS_ShouBaobao:
		return "OS_ShouBaobao"
	case OS_Weapon:
		return "OS_Weapon"
	case OS_Zhuachong:
		return "OS_Zhuachong"
	case OS_Flee:
		return "OS_Flee"
	case OS_Summon:
		return "OS_Summon"
	default:
		return ""
	}
}
func ToId_OrderStatus(name string) int {
	switch name {
	case "OS_None":
		return OS_None
	case "OS_ActiveOk":
		return OS_ActiveOk
	case "OS_RunawayOk":
		return OS_RunawayOk
	case "OS_FangBaobao":
		return OS_FangBaobao
	case "OS_ShouBaobao":
		return OS_ShouBaobao
	case "OS_Weapon":
		return OS_Weapon
	case "OS_Zhuachong":
		return OS_Zhuachong
	case "OS_Flee":
		return OS_Flee
	case "OS_Summon":
		return OS_Summon
	default:
		return -1
	}
}
