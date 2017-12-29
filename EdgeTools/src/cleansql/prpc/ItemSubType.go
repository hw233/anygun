package prpc

// enum ItemSubType
const (
	IST_None           = 0
	IST_Axe            = 1
	IST_Sword          = 2
	IST_Spear          = 3
	IST_Bow            = 4
	IST_Stick          = 5
	IST_Knife          = 6
	IST_Hat            = 7
	IST_Helmet         = 8
	IST_Clothes        = 9
	IST_Robe           = 10
	IST_Armor          = 11
	IST_Boot           = 12
	IST_Shoes          = 13
	IST_Shield         = 14
	IST_Crystal        = 15
	IST_Charm          = 16
	IST_Earrings       = 17
	IST_Bracelet       = 18
	IST_Ring           = 19
	IST_Necklace       = 20
	IST_Headband       = 21
	IST_Instruments    = 22
	IST_EquipMax       = 23
	IST_Ornament       = 24
	IST_Lottery        = 25
	IST_Coupun         = 26
	IST_OpenGird       = 27
	IST_ConsBegin      = 28
	IST_Consumables    = 29
	IST_Blood          = 30
	IST_Buff           = 31
	IST_Gem            = 32
	IST_Material       = 33
	IST_ItemDebris     = 34
	IST_BabyDebris     = 35
	IST_EmployeeDebris = 36
	IST_BabyExp        = 37
	IST_SkillExp       = 38
	IST_ConsEnd        = 39
	IST_Gloves         = 40
	IST_EmployeeEquip  = 41
	IST_PVP            = 42
	IST_Fashion        = 43
	IST_FuWenAttack    = 44
	IST_FuWenDefense   = 45
	IST_FuWenAssist    = 46
	IST_Max            = 47
)

func ToName_ItemSubType(id int) string {
	switch id {
	case IST_None:
		return "IST_None"
	case IST_Axe:
		return "IST_Axe"
	case IST_Sword:
		return "IST_Sword"
	case IST_Spear:
		return "IST_Spear"
	case IST_Bow:
		return "IST_Bow"
	case IST_Stick:
		return "IST_Stick"
	case IST_Knife:
		return "IST_Knife"
	case IST_Hat:
		return "IST_Hat"
	case IST_Helmet:
		return "IST_Helmet"
	case IST_Clothes:
		return "IST_Clothes"
	case IST_Robe:
		return "IST_Robe"
	case IST_Armor:
		return "IST_Armor"
	case IST_Boot:
		return "IST_Boot"
	case IST_Shoes:
		return "IST_Shoes"
	case IST_Shield:
		return "IST_Shield"
	case IST_Crystal:
		return "IST_Crystal"
	case IST_Charm:
		return "IST_Charm"
	case IST_Earrings:
		return "IST_Earrings"
	case IST_Bracelet:
		return "IST_Bracelet"
	case IST_Ring:
		return "IST_Ring"
	case IST_Necklace:
		return "IST_Necklace"
	case IST_Headband:
		return "IST_Headband"
	case IST_Instruments:
		return "IST_Instruments"
	case IST_EquipMax:
		return "IST_EquipMax"
	case IST_Ornament:
		return "IST_Ornament"
	case IST_Lottery:
		return "IST_Lottery"
	case IST_Coupun:
		return "IST_Coupun"
	case IST_OpenGird:
		return "IST_OpenGird"
	case IST_ConsBegin:
		return "IST_ConsBegin"
	case IST_Consumables:
		return "IST_Consumables"
	case IST_Blood:
		return "IST_Blood"
	case IST_Buff:
		return "IST_Buff"
	case IST_Gem:
		return "IST_Gem"
	case IST_Material:
		return "IST_Material"
	case IST_ItemDebris:
		return "IST_ItemDebris"
	case IST_BabyDebris:
		return "IST_BabyDebris"
	case IST_EmployeeDebris:
		return "IST_EmployeeDebris"
	case IST_BabyExp:
		return "IST_BabyExp"
	case IST_SkillExp:
		return "IST_SkillExp"
	case IST_ConsEnd:
		return "IST_ConsEnd"
	case IST_Gloves:
		return "IST_Gloves"
	case IST_EmployeeEquip:
		return "IST_EmployeeEquip"
	case IST_PVP:
		return "IST_PVP"
	case IST_Fashion:
		return "IST_Fashion"
	case IST_FuWenAttack:
		return "IST_FuWenAttack"
	case IST_FuWenDefense:
		return "IST_FuWenDefense"
	case IST_FuWenAssist:
		return "IST_FuWenAssist"
	case IST_Max:
		return "IST_Max"
	default:
		return ""
	}
}
func ToId_ItemSubType(name string) int {
	switch name {
	case "IST_None":
		return IST_None
	case "IST_Axe":
		return IST_Axe
	case "IST_Sword":
		return IST_Sword
	case "IST_Spear":
		return IST_Spear
	case "IST_Bow":
		return IST_Bow
	case "IST_Stick":
		return IST_Stick
	case "IST_Knife":
		return IST_Knife
	case "IST_Hat":
		return IST_Hat
	case "IST_Helmet":
		return IST_Helmet
	case "IST_Clothes":
		return IST_Clothes
	case "IST_Robe":
		return IST_Robe
	case "IST_Armor":
		return IST_Armor
	case "IST_Boot":
		return IST_Boot
	case "IST_Shoes":
		return IST_Shoes
	case "IST_Shield":
		return IST_Shield
	case "IST_Crystal":
		return IST_Crystal
	case "IST_Charm":
		return IST_Charm
	case "IST_Earrings":
		return IST_Earrings
	case "IST_Bracelet":
		return IST_Bracelet
	case "IST_Ring":
		return IST_Ring
	case "IST_Necklace":
		return IST_Necklace
	case "IST_Headband":
		return IST_Headband
	case "IST_Instruments":
		return IST_Instruments
	case "IST_EquipMax":
		return IST_EquipMax
	case "IST_Ornament":
		return IST_Ornament
	case "IST_Lottery":
		return IST_Lottery
	case "IST_Coupun":
		return IST_Coupun
	case "IST_OpenGird":
		return IST_OpenGird
	case "IST_ConsBegin":
		return IST_ConsBegin
	case "IST_Consumables":
		return IST_Consumables
	case "IST_Blood":
		return IST_Blood
	case "IST_Buff":
		return IST_Buff
	case "IST_Gem":
		return IST_Gem
	case "IST_Material":
		return IST_Material
	case "IST_ItemDebris":
		return IST_ItemDebris
	case "IST_BabyDebris":
		return IST_BabyDebris
	case "IST_EmployeeDebris":
		return IST_EmployeeDebris
	case "IST_BabyExp":
		return IST_BabyExp
	case "IST_SkillExp":
		return IST_SkillExp
	case "IST_ConsEnd":
		return IST_ConsEnd
	case "IST_Gloves":
		return IST_Gloves
	case "IST_EmployeeEquip":
		return IST_EmployeeEquip
	case "IST_PVP":
		return IST_PVP
	case "IST_Fashion":
		return IST_Fashion
	case "IST_FuWenAttack":
		return IST_FuWenAttack
	case "IST_FuWenDefense":
		return IST_FuWenDefense
	case "IST_FuWenAssist":
		return IST_FuWenAssist
	case "IST_Max":
		return IST_Max
	default:
		return -1
	}
}
