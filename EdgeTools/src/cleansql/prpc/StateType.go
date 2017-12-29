package prpc

// enum StateType
const (
	ST_None          = 0
	ST_Normal        = 1
	ST_Defense       = 2
	ST_Dodge         = 3
	ST_ActionAbsorb  = 4
	ST_MagicAbsorb   = 5
	ST_Shield        = 6
	ST_ActionBounce  = 7
	ST_MagicBounce   = 8
	ST_ActionInvalid = 9
	ST_MagicInvalid  = 10
	ST_Defend        = 11
	ST_NoDodge       = 12
	ST_Poison        = 13
	ST_Basilisk      = 14
	ST_Sleep         = 15
	ST_Chaos         = 16
	ST_Drunk         = 17
	ST_Forget        = 18
	ST_BeatBack      = 19
	ST_Recover       = 20
	ST_StrongRecover = 21
	ST_GroupRecover  = 22
	ST_MagicDef      = 23
	ST_Max           = 24
)

func ToName_StateType(id int) string {
	switch id {
	case ST_None:
		return "ST_None"
	case ST_Normal:
		return "ST_Normal"
	case ST_Defense:
		return "ST_Defense"
	case ST_Dodge:
		return "ST_Dodge"
	case ST_ActionAbsorb:
		return "ST_ActionAbsorb"
	case ST_MagicAbsorb:
		return "ST_MagicAbsorb"
	case ST_Shield:
		return "ST_Shield"
	case ST_ActionBounce:
		return "ST_ActionBounce"
	case ST_MagicBounce:
		return "ST_MagicBounce"
	case ST_ActionInvalid:
		return "ST_ActionInvalid"
	case ST_MagicInvalid:
		return "ST_MagicInvalid"
	case ST_Defend:
		return "ST_Defend"
	case ST_NoDodge:
		return "ST_NoDodge"
	case ST_Poison:
		return "ST_Poison"
	case ST_Basilisk:
		return "ST_Basilisk"
	case ST_Sleep:
		return "ST_Sleep"
	case ST_Chaos:
		return "ST_Chaos"
	case ST_Drunk:
		return "ST_Drunk"
	case ST_Forget:
		return "ST_Forget"
	case ST_BeatBack:
		return "ST_BeatBack"
	case ST_Recover:
		return "ST_Recover"
	case ST_StrongRecover:
		return "ST_StrongRecover"
	case ST_GroupRecover:
		return "ST_GroupRecover"
	case ST_MagicDef:
		return "ST_MagicDef"
	case ST_Max:
		return "ST_Max"
	default:
		return ""
	}
}
func ToId_StateType(name string) int {
	switch name {
	case "ST_None":
		return ST_None
	case "ST_Normal":
		return ST_Normal
	case "ST_Defense":
		return ST_Defense
	case "ST_Dodge":
		return ST_Dodge
	case "ST_ActionAbsorb":
		return ST_ActionAbsorb
	case "ST_MagicAbsorb":
		return ST_MagicAbsorb
	case "ST_Shield":
		return ST_Shield
	case "ST_ActionBounce":
		return ST_ActionBounce
	case "ST_MagicBounce":
		return ST_MagicBounce
	case "ST_ActionInvalid":
		return ST_ActionInvalid
	case "ST_MagicInvalid":
		return ST_MagicInvalid
	case "ST_Defend":
		return ST_Defend
	case "ST_NoDodge":
		return ST_NoDodge
	case "ST_Poison":
		return ST_Poison
	case "ST_Basilisk":
		return ST_Basilisk
	case "ST_Sleep":
		return ST_Sleep
	case "ST_Chaos":
		return ST_Chaos
	case "ST_Drunk":
		return ST_Drunk
	case "ST_Forget":
		return ST_Forget
	case "ST_BeatBack":
		return ST_BeatBack
	case "ST_Recover":
		return ST_Recover
	case "ST_StrongRecover":
		return ST_StrongRecover
	case "ST_GroupRecover":
		return ST_GroupRecover
	case "ST_MagicDef":
		return ST_MagicDef
	case "ST_Max":
		return ST_Max
	default:
		return -1
	}
}
