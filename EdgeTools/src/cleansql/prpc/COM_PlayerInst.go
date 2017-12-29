package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type COM_PlayerInst struct {
	COM_Entity
	isLeavingTeam_        bool                //0
	isBattle_             bool                //1
	autoBattle_           bool                //2
	isTeamLeader_         bool                //3
	sceneId_              int32               //4
	openSubSystemFlag_    int32               //5
	createTime_           int64               //6
	guildName_            string              //7
	scenePos_             COM_FPosition       //8
	pvpInfo_              COM_PlayerVsPlayer  //9
	onlineTimeFlag_       bool                //10
	onlineTime_           float32             //11
	onlineTimeReward_     []uint32            //12
	isFund_               bool                //13
	fundtags_             []uint32            //14
	openDoubleTimeFlag_   bool                //15
	isFirstLogin_         bool                //16
	firstRechargeDiamond_ bool                //17
	isFirstRechargeGift_  bool                //18
	offlineExp_           float32             //19
	rivalTime_            float32             //20
	rivalNum_             int8                //21
	promoteAward_         int8                //22
	guideIdx_             uint64              //23
	noTalkTime_           float32             //24
	wishShareNum_         uint32              //25
	warriortrophyNum_     uint32              //26
	employeelasttime_     uint32              //27
	employeeonecount_     uint32              //28
	employeetencount_     uint32              //29
	greenBoxTimes_        float32             //30
	blueBoxTimes_         float32             //31
	greenBoxFreeNum_      uint32              //32
	hbInfo_               COM_HundredBattle   //33
	openScenes_           []int32             //34
	copyNum_              []uint32            //35
	magicItemLevel_       int32               //36
	magicItemeExp_        int32               //37
	magicItemeJob_        int                 //38
	magicTupoLevel_       int32               //39
	cachedNpcs_           []int32             //40
	gft_                  []string            //41
	babycache_            []uint32            //42
	titles_               []int32             //43
	guildContribution_    int32               //44
	exitGuildTime_        uint32              //45
	guildSkills_          []COM_Skill         //46
	gmActivities_         []int               //47
	festival_             COM_ADLoginTotal    //48
	selfRecharge_         COM_ADChargeTotal   //49
	sysRecharge_          COM_ADChargeTotal   //50
	selfDiscountStore_    COM_ADDiscountStore //51
	sysDiscountStore_     COM_ADDiscountStore //52
	selfOnceRecharge_     COM_ADChargeEvery   //53
	sysOnceRecharge_      COM_ADChargeEvery   //54
	empact_               COM_ADEmployeeTotal //55
	selfCards_            COM_ADCards         //56
	myselfRecharge_       COM_ADChargeTotal   //57
	hotdata_              COM_ADHotRole       //58
	gbdata_               COM_ADGiftBag       //59
	sevenflag_            bool                //60
	signFlag_             bool                //61
	sevendate_            []COM_Sevenday      //62
	viprewardflag_        bool                //63
	phoneNumber_          string              //64
	levelgift_            []uint32            //65
	activity_             COM_ActivityTable   //66
	fuwen_                []COM_Item          //67
}

func (this *COM_PlayerInst) Serialize(buffer *bytes.Buffer) error {
	{
		err := this.COM_Entity.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	//field mask
	mask := prpc.NewMask1(9)
	mask.WriteBit(this.isLeavingTeam_)
	mask.WriteBit(this.isBattle_)
	mask.WriteBit(this.autoBattle_)
	mask.WriteBit(this.isTeamLeader_)
	mask.WriteBit(this.sceneId_ != 0)
	mask.WriteBit(this.openSubSystemFlag_ != 0)
	mask.WriteBit(this.createTime_ != 0)
	mask.WriteBit(len(this.guildName_) != 0)
	mask.WriteBit(true) //scenePos_
	mask.WriteBit(true) //pvpInfo_
	mask.WriteBit(this.onlineTimeFlag_)
	mask.WriteBit(this.onlineTime_ != 0)
	mask.WriteBit(len(this.onlineTimeReward_) != 0)
	mask.WriteBit(this.isFund_)
	mask.WriteBit(len(this.fundtags_) != 0)
	mask.WriteBit(this.openDoubleTimeFlag_)
	mask.WriteBit(this.isFirstLogin_)
	mask.WriteBit(this.firstRechargeDiamond_)
	mask.WriteBit(this.isFirstRechargeGift_)
	mask.WriteBit(this.offlineExp_ != 0)
	mask.WriteBit(this.rivalTime_ != 0)
	mask.WriteBit(this.rivalNum_ != 0)
	mask.WriteBit(this.promoteAward_ != 0)
	mask.WriteBit(this.guideIdx_ != 0)
	mask.WriteBit(this.noTalkTime_ != 0)
	mask.WriteBit(this.wishShareNum_ != 0)
	mask.WriteBit(this.warriortrophyNum_ != 0)
	mask.WriteBit(this.employeelasttime_ != 0)
	mask.WriteBit(this.employeeonecount_ != 0)
	mask.WriteBit(this.employeetencount_ != 0)
	mask.WriteBit(this.greenBoxTimes_ != 0)
	mask.WriteBit(this.blueBoxTimes_ != 0)
	mask.WriteBit(this.greenBoxFreeNum_ != 0)
	mask.WriteBit(true) //hbInfo_
	mask.WriteBit(len(this.openScenes_) != 0)
	mask.WriteBit(len(this.copyNum_) != 0)
	mask.WriteBit(this.magicItemLevel_ != 0)
	mask.WriteBit(this.magicItemeExp_ != 0)
	mask.WriteBit(this.magicItemeJob_ != 0)
	mask.WriteBit(this.magicTupoLevel_ != 0)
	mask.WriteBit(len(this.cachedNpcs_) != 0)
	mask.WriteBit(len(this.gft_) != 0)
	mask.WriteBit(len(this.babycache_) != 0)
	mask.WriteBit(len(this.titles_) != 0)
	mask.WriteBit(this.guildContribution_ != 0)
	mask.WriteBit(this.exitGuildTime_ != 0)
	mask.WriteBit(len(this.guildSkills_) != 0)
	mask.WriteBit(len(this.gmActivities_) != 0)
	mask.WriteBit(true) //festival_
	mask.WriteBit(true) //selfRecharge_
	mask.WriteBit(true) //sysRecharge_
	mask.WriteBit(true) //selfDiscountStore_
	mask.WriteBit(true) //sysDiscountStore_
	mask.WriteBit(true) //selfOnceRecharge_
	mask.WriteBit(true) //sysOnceRecharge_
	mask.WriteBit(true) //empact_
	mask.WriteBit(true) //selfCards_
	mask.WriteBit(true) //myselfRecharge_
	mask.WriteBit(true) //hotdata_
	mask.WriteBit(true) //gbdata_
	mask.WriteBit(this.sevenflag_)
	mask.WriteBit(this.signFlag_)
	mask.WriteBit(len(this.sevendate_) != 0)
	mask.WriteBit(this.viprewardflag_)
	mask.WriteBit(len(this.phoneNumber_) != 0)
	mask.WriteBit(len(this.levelgift_) != 0)
	mask.WriteBit(true) //activity_
	mask.WriteBit(len(this.fuwen_) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize isLeavingTeam_
	{
	}
	// serialize isBattle_
	{
	}
	// serialize autoBattle_
	{
	}
	// serialize isTeamLeader_
	{
	}
	// serialize sceneId_
	{
		if this.sceneId_ != 0 {
			err := prpc.Write(buffer, this.sceneId_)
			if err != nil {
				return err
			}
		}
	}
	// serialize openSubSystemFlag_
	{
		if this.openSubSystemFlag_ != 0 {
			err := prpc.Write(buffer, this.openSubSystemFlag_)
			if err != nil {
				return err
			}
		}
	}
	// serialize createTime_
	{
		if this.createTime_ != 0 {
			err := prpc.Write(buffer, this.createTime_)
			if err != nil {
				return err
			}
		}
	}
	// serialize guildName_
	if len(this.guildName_) != 0 {
		err := prpc.Write(buffer, this.guildName_)
		if err != nil {
			return err
		}
	}
	// serialize scenePos_
	{
		err := this.scenePos_.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	// serialize pvpInfo_
	{
		err := this.pvpInfo_.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	// serialize onlineTimeFlag_
	{
	}
	// serialize onlineTime_
	{
		if this.onlineTime_ != 0 {
			err := prpc.Write(buffer, this.onlineTime_)
			if err != nil {
				return err
			}
		}
	}
	// serialize onlineTimeReward_
	if len(this.onlineTimeReward_) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.onlineTimeReward_)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.onlineTimeReward_ {
			err := prpc.Write(buffer, value)
			if err != nil {
				return err
			}
		}
	}
	// serialize isFund_
	{
	}
	// serialize fundtags_
	if len(this.fundtags_) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.fundtags_)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.fundtags_ {
			err := prpc.Write(buffer, value)
			if err != nil {
				return err
			}
		}
	}
	// serialize openDoubleTimeFlag_
	{
	}
	// serialize isFirstLogin_
	{
	}
	// serialize firstRechargeDiamond_
	{
	}
	// serialize isFirstRechargeGift_
	{
	}
	// serialize offlineExp_
	{
		if this.offlineExp_ != 0 {
			err := prpc.Write(buffer, this.offlineExp_)
			if err != nil {
				return err
			}
		}
	}
	// serialize rivalTime_
	{
		if this.rivalTime_ != 0 {
			err := prpc.Write(buffer, this.rivalTime_)
			if err != nil {
				return err
			}
		}
	}
	// serialize rivalNum_
	{
		if this.rivalNum_ != 0 {
			err := prpc.Write(buffer, this.rivalNum_)
			if err != nil {
				return err
			}
		}
	}
	// serialize promoteAward_
	{
		if this.promoteAward_ != 0 {
			err := prpc.Write(buffer, this.promoteAward_)
			if err != nil {
				return err
			}
		}
	}
	// serialize guideIdx_
	{
		if this.guideIdx_ != 0 {
			err := prpc.Write(buffer, this.guideIdx_)
			if err != nil {
				return err
			}
		}
	}
	// serialize noTalkTime_
	{
		if this.noTalkTime_ != 0 {
			err := prpc.Write(buffer, this.noTalkTime_)
			if err != nil {
				return err
			}
		}
	}
	// serialize wishShareNum_
	{
		if this.wishShareNum_ != 0 {
			err := prpc.Write(buffer, this.wishShareNum_)
			if err != nil {
				return err
			}
		}
	}
	// serialize warriortrophyNum_
	{
		if this.warriortrophyNum_ != 0 {
			err := prpc.Write(buffer, this.warriortrophyNum_)
			if err != nil {
				return err
			}
		}
	}
	// serialize employeelasttime_
	{
		if this.employeelasttime_ != 0 {
			err := prpc.Write(buffer, this.employeelasttime_)
			if err != nil {
				return err
			}
		}
	}
	// serialize employeeonecount_
	{
		if this.employeeonecount_ != 0 {
			err := prpc.Write(buffer, this.employeeonecount_)
			if err != nil {
				return err
			}
		}
	}
	// serialize employeetencount_
	{
		if this.employeetencount_ != 0 {
			err := prpc.Write(buffer, this.employeetencount_)
			if err != nil {
				return err
			}
		}
	}
	// serialize greenBoxTimes_
	{
		if this.greenBoxTimes_ != 0 {
			err := prpc.Write(buffer, this.greenBoxTimes_)
			if err != nil {
				return err
			}
		}
	}
	// serialize blueBoxTimes_
	{
		if this.blueBoxTimes_ != 0 {
			err := prpc.Write(buffer, this.blueBoxTimes_)
			if err != nil {
				return err
			}
		}
	}
	// serialize greenBoxFreeNum_
	{
		if this.greenBoxFreeNum_ != 0 {
			err := prpc.Write(buffer, this.greenBoxFreeNum_)
			if err != nil {
				return err
			}
		}
	}
	// serialize hbInfo_
	{
		err := this.hbInfo_.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	// serialize openScenes_
	if len(this.openScenes_) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.openScenes_)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.openScenes_ {
			err := prpc.Write(buffer, value)
			if err != nil {
				return err
			}
		}
	}
	// serialize copyNum_
	if len(this.copyNum_) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.copyNum_)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.copyNum_ {
			err := prpc.Write(buffer, value)
			if err != nil {
				return err
			}
		}
	}
	// serialize magicItemLevel_
	{
		if this.magicItemLevel_ != 0 {
			err := prpc.Write(buffer, this.magicItemLevel_)
			if err != nil {
				return err
			}
		}
	}
	// serialize magicItemeExp_
	{
		if this.magicItemeExp_ != 0 {
			err := prpc.Write(buffer, this.magicItemeExp_)
			if err != nil {
				return err
			}
		}
	}
	// serialize magicItemeJob_
	{
		if this.magicItemeJob_ != 0 {
			err := prpc.Write(buffer, this.magicItemeJob_)
			if err != nil {
				return err
			}
		}
	}
	// serialize magicTupoLevel_
	{
		if this.magicTupoLevel_ != 0 {
			err := prpc.Write(buffer, this.magicTupoLevel_)
			if err != nil {
				return err
			}
		}
	}
	// serialize cachedNpcs_
	if len(this.cachedNpcs_) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.cachedNpcs_)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.cachedNpcs_ {
			err := prpc.Write(buffer, value)
			if err != nil {
				return err
			}
		}
	}
	// serialize gft_
	if len(this.gft_) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.gft_)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.gft_ {
			err := prpc.Write(buffer, value)
			if err != nil {
				return err
			}
		}
	}
	// serialize babycache_
	if len(this.babycache_) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.babycache_)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.babycache_ {
			err := prpc.Write(buffer, value)
			if err != nil {
				return err
			}
		}
	}
	// serialize titles_
	if len(this.titles_) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.titles_)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.titles_ {
			err := prpc.Write(buffer, value)
			if err != nil {
				return err
			}
		}
	}
	// serialize guildContribution_
	{
		if this.guildContribution_ != 0 {
			err := prpc.Write(buffer, this.guildContribution_)
			if err != nil {
				return err
			}
		}
	}
	// serialize exitGuildTime_
	{
		if this.exitGuildTime_ != 0 {
			err := prpc.Write(buffer, this.exitGuildTime_)
			if err != nil {
				return err
			}
		}
	}
	// serialize guildSkills_
	if len(this.guildSkills_) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.guildSkills_)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.guildSkills_ {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// serialize gmActivities_
	if len(this.gmActivities_) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.gmActivities_)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.gmActivities_ {
			err := prpc.Write(buffer, value)
			if err != nil {
				return err
			}
		}
	}
	// serialize festival_
	{
		err := this.festival_.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	// serialize selfRecharge_
	{
		err := this.selfRecharge_.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	// serialize sysRecharge_
	{
		err := this.sysRecharge_.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	// serialize selfDiscountStore_
	{
		err := this.selfDiscountStore_.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	// serialize sysDiscountStore_
	{
		err := this.sysDiscountStore_.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	// serialize selfOnceRecharge_
	{
		err := this.selfOnceRecharge_.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	// serialize sysOnceRecharge_
	{
		err := this.sysOnceRecharge_.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	// serialize empact_
	{
		err := this.empact_.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	// serialize selfCards_
	{
		err := this.selfCards_.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	// serialize myselfRecharge_
	{
		err := this.myselfRecharge_.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	// serialize hotdata_
	{
		err := this.hotdata_.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	// serialize gbdata_
	{
		err := this.gbdata_.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	// serialize sevenflag_
	{
	}
	// serialize signFlag_
	{
	}
	// serialize sevendate_
	if len(this.sevendate_) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.sevendate_)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.sevendate_ {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// serialize viprewardflag_
	{
	}
	// serialize phoneNumber_
	if len(this.phoneNumber_) != 0 {
		err := prpc.Write(buffer, this.phoneNumber_)
		if err != nil {
			return err
		}
	}
	// serialize levelgift_
	if len(this.levelgift_) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.levelgift_)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.levelgift_ {
			err := prpc.Write(buffer, value)
			if err != nil {
				return err
			}
		}
	}
	// serialize activity_
	{
		err := this.activity_.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	// serialize fuwen_
	if len(this.fuwen_) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.fuwen_)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.fuwen_ {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *COM_PlayerInst) Deserialize(buffer *bytes.Buffer) error {
	{
		this.COM_Entity.Deserialize(buffer)
	}
	//field mask
	mask, err := prpc.NewMask0(buffer, 9)
	if err != nil {
		return err
	}
	// deserialize isLeavingTeam_
	this.isLeavingTeam_ = mask.ReadBit()
	// deserialize isBattle_
	this.isBattle_ = mask.ReadBit()
	// deserialize autoBattle_
	this.autoBattle_ = mask.ReadBit()
	// deserialize isTeamLeader_
	this.isTeamLeader_ = mask.ReadBit()
	// deserialize sceneId_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.sceneId_)
		if err != nil {
			return err
		}
	}
	// deserialize openSubSystemFlag_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.openSubSystemFlag_)
		if err != nil {
			return err
		}
	}
	// deserialize createTime_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.createTime_)
		if err != nil {
			return err
		}
	}
	// deserialize guildName_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.guildName_)
		if err != nil {
			return err
		}
	}
	// deserialize scenePos_
	if mask.ReadBit() {
		err := this.scenePos_.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	// deserialize pvpInfo_
	if mask.ReadBit() {
		err := this.pvpInfo_.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	// deserialize onlineTimeFlag_
	this.onlineTimeFlag_ = mask.ReadBit()
	// deserialize onlineTime_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.onlineTime_)
		if err != nil {
			return err
		}
	}
	// deserialize onlineTimeReward_
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.onlineTimeReward_ = make([]uint32, size)
		for i, _ := range this.onlineTimeReward_ {
			err := prpc.Read(buffer, &this.onlineTimeReward_[i])
			if err != nil {
				return err
			}
		}
	}
	// deserialize isFund_
	this.isFund_ = mask.ReadBit()
	// deserialize fundtags_
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.fundtags_ = make([]uint32, size)
		for i, _ := range this.fundtags_ {
			err := prpc.Read(buffer, &this.fundtags_[i])
			if err != nil {
				return err
			}
		}
	}
	// deserialize openDoubleTimeFlag_
	this.openDoubleTimeFlag_ = mask.ReadBit()
	// deserialize isFirstLogin_
	this.isFirstLogin_ = mask.ReadBit()
	// deserialize firstRechargeDiamond_
	this.firstRechargeDiamond_ = mask.ReadBit()
	// deserialize isFirstRechargeGift_
	this.isFirstRechargeGift_ = mask.ReadBit()
	// deserialize offlineExp_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.offlineExp_)
		if err != nil {
			return err
		}
	}
	// deserialize rivalTime_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.rivalTime_)
		if err != nil {
			return err
		}
	}
	// deserialize rivalNum_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.rivalNum_)
		if err != nil {
			return err
		}
	}
	// deserialize promoteAward_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.promoteAward_)
		if err != nil {
			return err
		}
	}
	// deserialize guideIdx_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.guideIdx_)
		if err != nil {
			return err
		}
	}
	// deserialize noTalkTime_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.noTalkTime_)
		if err != nil {
			return err
		}
	}
	// deserialize wishShareNum_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.wishShareNum_)
		if err != nil {
			return err
		}
	}
	// deserialize warriortrophyNum_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.warriortrophyNum_)
		if err != nil {
			return err
		}
	}
	// deserialize employeelasttime_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.employeelasttime_)
		if err != nil {
			return err
		}
	}
	// deserialize employeeonecount_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.employeeonecount_)
		if err != nil {
			return err
		}
	}
	// deserialize employeetencount_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.employeetencount_)
		if err != nil {
			return err
		}
	}
	// deserialize greenBoxTimes_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.greenBoxTimes_)
		if err != nil {
			return err
		}
	}
	// deserialize blueBoxTimes_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.blueBoxTimes_)
		if err != nil {
			return err
		}
	}
	// deserialize greenBoxFreeNum_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.greenBoxFreeNum_)
		if err != nil {
			return err
		}
	}
	// deserialize hbInfo_
	if mask.ReadBit() {
		err := this.hbInfo_.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	// deserialize openScenes_
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.openScenes_ = make([]int32, size)
		for i, _ := range this.openScenes_ {
			err := prpc.Read(buffer, &this.openScenes_[i])
			if err != nil {
				return err
			}
		}
	}
	// deserialize copyNum_
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.copyNum_ = make([]uint32, size)
		for i, _ := range this.copyNum_ {
			err := prpc.Read(buffer, &this.copyNum_[i])
			if err != nil {
				return err
			}
		}
	}
	// deserialize magicItemLevel_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.magicItemLevel_)
		if err != nil {
			return err
		}
	}
	// deserialize magicItemeExp_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.magicItemeExp_)
		if err != nil {
			return err
		}
	}
	// deserialize magicItemeJob_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.magicItemeJob_)
		if err != nil {
			return err
		}
	}
	// deserialize magicTupoLevel_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.magicTupoLevel_)
		if err != nil {
			return err
		}
	}
	// deserialize cachedNpcs_
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.cachedNpcs_ = make([]int32, size)
		for i, _ := range this.cachedNpcs_ {
			err := prpc.Read(buffer, &this.cachedNpcs_[i])
			if err != nil {
				return err
			}
		}
	}
	// deserialize gft_
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.gft_ = make([]string, size)
		for i, _ := range this.gft_ {
			err := prpc.Read(buffer, &this.gft_[i])
			if err != nil {
				return err
			}
		}
	}
	// deserialize babycache_
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.babycache_ = make([]uint32, size)
		for i, _ := range this.babycache_ {
			err := prpc.Read(buffer, &this.babycache_[i])
			if err != nil {
				return err
			}
		}
	}
	// deserialize titles_
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.titles_ = make([]int32, size)
		for i, _ := range this.titles_ {
			err := prpc.Read(buffer, &this.titles_[i])
			if err != nil {
				return err
			}
		}
	}
	// deserialize guildContribution_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.guildContribution_)
		if err != nil {
			return err
		}
	}
	// deserialize exitGuildTime_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.exitGuildTime_)
		if err != nil {
			return err
		}
	}
	// deserialize guildSkills_
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.guildSkills_ = make([]COM_Skill, size)
		for i, _ := range this.guildSkills_ {
			err := this.guildSkills_[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// deserialize gmActivities_
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.gmActivities_ = make([]int, size)
		for i, _ := range this.gmActivities_ {
			err := prpc.Read(buffer, &this.gmActivities_[i])
			if err != nil {
				return err
			}
		}
	}
	// deserialize festival_
	if mask.ReadBit() {
		err := this.festival_.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	// deserialize selfRecharge_
	if mask.ReadBit() {
		err := this.selfRecharge_.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	// deserialize sysRecharge_
	if mask.ReadBit() {
		err := this.sysRecharge_.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	// deserialize selfDiscountStore_
	if mask.ReadBit() {
		err := this.selfDiscountStore_.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	// deserialize sysDiscountStore_
	if mask.ReadBit() {
		err := this.sysDiscountStore_.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	// deserialize selfOnceRecharge_
	if mask.ReadBit() {
		err := this.selfOnceRecharge_.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	// deserialize sysOnceRecharge_
	if mask.ReadBit() {
		err := this.sysOnceRecharge_.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	// deserialize empact_
	if mask.ReadBit() {
		err := this.empact_.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	// deserialize selfCards_
	if mask.ReadBit() {
		err := this.selfCards_.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	// deserialize myselfRecharge_
	if mask.ReadBit() {
		err := this.myselfRecharge_.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	// deserialize hotdata_
	if mask.ReadBit() {
		err := this.hotdata_.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	// deserialize gbdata_
	if mask.ReadBit() {
		err := this.gbdata_.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	// deserialize sevenflag_
	this.sevenflag_ = mask.ReadBit()
	// deserialize signFlag_
	this.signFlag_ = mask.ReadBit()
	// deserialize sevendate_
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.sevendate_ = make([]COM_Sevenday, size)
		for i, _ := range this.sevendate_ {
			err := this.sevendate_[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// deserialize viprewardflag_
	this.viprewardflag_ = mask.ReadBit()
	// deserialize phoneNumber_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.phoneNumber_)
		if err != nil {
			return err
		}
	}
	// deserialize levelgift_
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.levelgift_ = make([]uint32, size)
		for i, _ := range this.levelgift_ {
			err := prpc.Read(buffer, &this.levelgift_[i])
			if err != nil {
				return err
			}
		}
	}
	// deserialize activity_
	if mask.ReadBit() {
		err := this.activity_.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	// deserialize fuwen_
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.fuwen_ = make([]COM_Item, size)
		for i, _ := range this.fuwen_ {
			err := this.fuwen_[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
