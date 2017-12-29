package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type SGE_DBPlayerData struct {
	COM_PlayerInst
	freeze_                 bool               //0
	seal_                   bool               //1
	signs_                  uint32             //2
	sellIdMax_              int32              //3
	acceptRandQuestCounter_ int32              //4
	submitRandQuestCounter_ int32              //5
	pfid_                   string             //6
	orders_                 []SGE_OrderInfo    //7
	loginTime_              uint64             //8
	logoutTime_             uint64             //9
	genItemMaxGuid_         uint32             //10
	gaterMaxNum_            uint32             //11
	firstRollEmployeeCon_   bool               //12
	firstRollEmployeeDia_   bool               //13
	employees_              []COM_EmployeeInst //14
	itemStorage_            []COM_Item         //15
	babyStorage_            []COM_BabyInst     //16
	babies_                 []COM_BabyInst     //17
	bagItems_               []COM_Item         //18
	quests_                 []COM_QuestInst    //19
	completeQuests_         []int32            //20
	mineReward_             []COM_Item         //21
	jjcBattleMsg_           []COM_JJCBattleMsg //22
	friend_                 []COM_ContactInfo  //23
	blacklist_              []COM_ContactInfo  //24
	achValues_              []int32            //25
	achievement_            []COM_Achievement  //26
	empBattleGroup_         int                //27
	employeeGroup1_         []uint32           //28
	employeeGroup2_         []uint32           //29
	gatherDate_             []COM_Gather       //30
	compoundList_           []uint32           //31
}

func (this *SGE_DBPlayerData) Serialize(buffer *bytes.Buffer) error {
	{
		err := this.COM_PlayerInst.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	//field mask
	mask := prpc.NewMask1(4)
	mask.WriteBit(this.freeze_)
	mask.WriteBit(this.seal_)
	mask.WriteBit(this.signs_ != 0)
	mask.WriteBit(this.sellIdMax_ != 0)
	mask.WriteBit(this.acceptRandQuestCounter_ != 0)
	mask.WriteBit(this.submitRandQuestCounter_ != 0)
	mask.WriteBit(len(this.pfid_) != 0)
	mask.WriteBit(len(this.orders_) != 0)
	mask.WriteBit(this.loginTime_ != 0)
	mask.WriteBit(this.logoutTime_ != 0)
	mask.WriteBit(this.genItemMaxGuid_ != 0)
	mask.WriteBit(this.gaterMaxNum_ != 0)
	mask.WriteBit(this.firstRollEmployeeCon_)
	mask.WriteBit(this.firstRollEmployeeDia_)
	mask.WriteBit(len(this.employees_) != 0)
	mask.WriteBit(len(this.itemStorage_) != 0)
	mask.WriteBit(len(this.babyStorage_) != 0)
	mask.WriteBit(len(this.babies_) != 0)
	mask.WriteBit(len(this.bagItems_) != 0)
	mask.WriteBit(len(this.quests_) != 0)
	mask.WriteBit(len(this.completeQuests_) != 0)
	mask.WriteBit(len(this.mineReward_) != 0)
	mask.WriteBit(len(this.jjcBattleMsg_) != 0)
	mask.WriteBit(len(this.friend_) != 0)
	mask.WriteBit(len(this.blacklist_) != 0)
	mask.WriteBit(len(this.achValues_) != 0)
	mask.WriteBit(len(this.achievement_) != 0)
	mask.WriteBit(this.empBattleGroup_ != 0)
	mask.WriteBit(len(this.employeeGroup1_) != 0)
	mask.WriteBit(len(this.employeeGroup2_) != 0)
	mask.WriteBit(len(this.gatherDate_) != 0)
	mask.WriteBit(len(this.compoundList_) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize freeze_
	{
	}
	// serialize seal_
	{
	}
	// serialize signs_
	{
		if this.signs_ != 0 {
			err := prpc.Write(buffer, this.signs_)
			if err != nil {
				return err
			}
		}
	}
	// serialize sellIdMax_
	{
		if this.sellIdMax_ != 0 {
			err := prpc.Write(buffer, this.sellIdMax_)
			if err != nil {
				return err
			}
		}
	}
	// serialize acceptRandQuestCounter_
	{
		if this.acceptRandQuestCounter_ != 0 {
			err := prpc.Write(buffer, this.acceptRandQuestCounter_)
			if err != nil {
				return err
			}
		}
	}
	// serialize submitRandQuestCounter_
	{
		if this.submitRandQuestCounter_ != 0 {
			err := prpc.Write(buffer, this.submitRandQuestCounter_)
			if err != nil {
				return err
			}
		}
	}
	// serialize pfid_
	if len(this.pfid_) != 0 {
		err := prpc.Write(buffer, this.pfid_)
		if err != nil {
			return err
		}
	}
	// serialize orders_
	if len(this.orders_) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.orders_)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.orders_ {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// serialize loginTime_
	{
		if this.loginTime_ != 0 {
			err := prpc.Write(buffer, this.loginTime_)
			if err != nil {
				return err
			}
		}
	}
	// serialize logoutTime_
	{
		if this.logoutTime_ != 0 {
			err := prpc.Write(buffer, this.logoutTime_)
			if err != nil {
				return err
			}
		}
	}
	// serialize genItemMaxGuid_
	{
		if this.genItemMaxGuid_ != 0 {
			err := prpc.Write(buffer, this.genItemMaxGuid_)
			if err != nil {
				return err
			}
		}
	}
	// serialize gaterMaxNum_
	{
		if this.gaterMaxNum_ != 0 {
			err := prpc.Write(buffer, this.gaterMaxNum_)
			if err != nil {
				return err
			}
		}
	}
	// serialize firstRollEmployeeCon_
	{
	}
	// serialize firstRollEmployeeDia_
	{
	}
	// serialize employees_
	if len(this.employees_) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.employees_)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.employees_ {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// serialize itemStorage_
	if len(this.itemStorage_) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.itemStorage_)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.itemStorage_ {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// serialize babyStorage_
	if len(this.babyStorage_) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.babyStorage_)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.babyStorage_ {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// serialize babies_
	if len(this.babies_) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.babies_)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.babies_ {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// serialize bagItems_
	if len(this.bagItems_) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.bagItems_)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.bagItems_ {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// serialize quests_
	if len(this.quests_) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.quests_)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.quests_ {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// serialize completeQuests_
	if len(this.completeQuests_) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.completeQuests_)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.completeQuests_ {
			err := prpc.Write(buffer, value)
			if err != nil {
				return err
			}
		}
	}
	// serialize mineReward_
	if len(this.mineReward_) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.mineReward_)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.mineReward_ {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// serialize jjcBattleMsg_
	if len(this.jjcBattleMsg_) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.jjcBattleMsg_)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.jjcBattleMsg_ {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// serialize friend_
	if len(this.friend_) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.friend_)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.friend_ {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// serialize blacklist_
	if len(this.blacklist_) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.blacklist_)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.blacklist_ {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// serialize achValues_
	if len(this.achValues_) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.achValues_)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.achValues_ {
			err := prpc.Write(buffer, value)
			if err != nil {
				return err
			}
		}
	}
	// serialize achievement_
	if len(this.achievement_) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.achievement_)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.achievement_ {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// serialize empBattleGroup_
	{
		if this.empBattleGroup_ != 0 {
			err := prpc.Write(buffer, this.empBattleGroup_)
			if err != nil {
				return err
			}
		}
	}
	// serialize employeeGroup1_
	if len(this.employeeGroup1_) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.employeeGroup1_)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.employeeGroup1_ {
			err := prpc.Write(buffer, value)
			if err != nil {
				return err
			}
		}
	}
	// serialize employeeGroup2_
	if len(this.employeeGroup2_) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.employeeGroup2_)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.employeeGroup2_ {
			err := prpc.Write(buffer, value)
			if err != nil {
				return err
			}
		}
	}
	// serialize gatherDate_
	if len(this.gatherDate_) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.gatherDate_)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.gatherDate_ {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// serialize compoundList_
	if len(this.compoundList_) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.compoundList_)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.compoundList_ {
			err := prpc.Write(buffer, value)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *SGE_DBPlayerData) Deserialize(buffer *bytes.Buffer) error {
	{
		this.COM_PlayerInst.Deserialize(buffer)
	}
	//field mask
	mask, err := prpc.NewMask0(buffer, 4)
	if err != nil {
		return err
	}
	// deserialize freeze_
	this.freeze_ = mask.ReadBit()
	// deserialize seal_
	this.seal_ = mask.ReadBit()
	// deserialize signs_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.signs_)
		if err != nil {
			return err
		}
	}
	// deserialize sellIdMax_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.sellIdMax_)
		if err != nil {
			return err
		}
	}
	// deserialize acceptRandQuestCounter_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.acceptRandQuestCounter_)
		if err != nil {
			return err
		}
	}
	// deserialize submitRandQuestCounter_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.submitRandQuestCounter_)
		if err != nil {
			return err
		}
	}
	// deserialize pfid_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.pfid_)
		if err != nil {
			return err
		}
	}
	// deserialize orders_
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.orders_ = make([]SGE_OrderInfo, size)
		for i, _ := range this.orders_ {
			err := this.orders_[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// deserialize loginTime_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.loginTime_)
		if err != nil {
			return err
		}
	}
	// deserialize logoutTime_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.logoutTime_)
		if err != nil {
			return err
		}
	}
	// deserialize genItemMaxGuid_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.genItemMaxGuid_)
		if err != nil {
			return err
		}
	}
	// deserialize gaterMaxNum_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.gaterMaxNum_)
		if err != nil {
			return err
		}
	}
	// deserialize firstRollEmployeeCon_
	this.firstRollEmployeeCon_ = mask.ReadBit()
	// deserialize firstRollEmployeeDia_
	this.firstRollEmployeeDia_ = mask.ReadBit()
	// deserialize employees_
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.employees_ = make([]COM_EmployeeInst, size)
		for i, _ := range this.employees_ {
			err := this.employees_[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// deserialize itemStorage_
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.itemStorage_ = make([]COM_Item, size)
		for i, _ := range this.itemStorage_ {
			err := this.itemStorage_[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// deserialize babyStorage_
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.babyStorage_ = make([]COM_BabyInst, size)
		for i, _ := range this.babyStorage_ {
			err := this.babyStorage_[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// deserialize babies_
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.babies_ = make([]COM_BabyInst, size)
		for i, _ := range this.babies_ {
			err := this.babies_[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// deserialize bagItems_
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.bagItems_ = make([]COM_Item, size)
		for i, _ := range this.bagItems_ {
			err := this.bagItems_[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// deserialize quests_
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.quests_ = make([]COM_QuestInst, size)
		for i, _ := range this.quests_ {
			err := this.quests_[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// deserialize completeQuests_
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.completeQuests_ = make([]int32, size)
		for i, _ := range this.completeQuests_ {
			err := prpc.Read(buffer, &this.completeQuests_[i])
			if err != nil {
				return err
			}
		}
	}
	// deserialize mineReward_
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.mineReward_ = make([]COM_Item, size)
		for i, _ := range this.mineReward_ {
			err := this.mineReward_[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// deserialize jjcBattleMsg_
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.jjcBattleMsg_ = make([]COM_JJCBattleMsg, size)
		for i, _ := range this.jjcBattleMsg_ {
			err := this.jjcBattleMsg_[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// deserialize friend_
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.friend_ = make([]COM_ContactInfo, size)
		for i, _ := range this.friend_ {
			err := this.friend_[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// deserialize blacklist_
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.blacklist_ = make([]COM_ContactInfo, size)
		for i, _ := range this.blacklist_ {
			err := this.blacklist_[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// deserialize achValues_
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.achValues_ = make([]int32, size)
		for i, _ := range this.achValues_ {
			err := prpc.Read(buffer, &this.achValues_[i])
			if err != nil {
				return err
			}
		}
	}
	// deserialize achievement_
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.achievement_ = make([]COM_Achievement, size)
		for i, _ := range this.achievement_ {
			err := this.achievement_[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// deserialize empBattleGroup_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.empBattleGroup_)
		if err != nil {
			return err
		}
	}
	// deserialize employeeGroup1_
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.employeeGroup1_ = make([]uint32, size)
		for i, _ := range this.employeeGroup1_ {
			err := prpc.Read(buffer, &this.employeeGroup1_[i])
			if err != nil {
				return err
			}
		}
	}
	// deserialize employeeGroup2_
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.employeeGroup2_ = make([]uint32, size)
		for i, _ := range this.employeeGroup2_ {
			err := prpc.Read(buffer, &this.employeeGroup2_[i])
			if err != nil {
				return err
			}
		}
	}
	// deserialize gatherDate_
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.gatherDate_ = make([]COM_Gather, size)
		for i, _ := range this.gatherDate_ {
			err := this.gatherDate_[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// deserialize compoundList_
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.compoundList_ = make([]uint32, size)
		for i, _ := range this.compoundList_ {
			err := prpc.Read(buffer, &this.compoundList_[i])
			if err != nil {
				return err
			}
		}
	}
	return nil
}
