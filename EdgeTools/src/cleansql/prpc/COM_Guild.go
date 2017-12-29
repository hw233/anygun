package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type COM_Guild struct {
	guildLevel_          int8                   //0
	createTime_          int32                  //1
	guildId_             uint32                 //2
	guildContribution_   uint32                 //3
	fundz_               uint32                 //4
	presentNum_          int32                  //5
	master_              int64                  //6
	masterName_          string                 //7
	guildName_           string                 //8
	notice_              string                 //9
	requestList_         []COM_GuildRequestData //10
	noFundzDays_         int32                  //11
	buildings_           []COM_GuildBuilding    //12
	progenitus_          []COM_GuildProgen      //13
	progenitusPositions_ []int32                //14
}

func (this *COM_Guild) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(2)
	mask.WriteBit(this.guildLevel_ != 0)
	mask.WriteBit(this.createTime_ != 0)
	mask.WriteBit(this.guildId_ != 0)
	mask.WriteBit(this.guildContribution_ != 0)
	mask.WriteBit(this.fundz_ != 0)
	mask.WriteBit(this.presentNum_ != 0)
	mask.WriteBit(this.master_ != 0)
	mask.WriteBit(len(this.masterName_) != 0)
	mask.WriteBit(len(this.guildName_) != 0)
	mask.WriteBit(len(this.notice_) != 0)
	mask.WriteBit(len(this.requestList_) != 0)
	mask.WriteBit(this.noFundzDays_ != 0)
	mask.WriteBit(len(this.buildings_) != 0)
	mask.WriteBit(len(this.progenitus_) != 0)
	mask.WriteBit(len(this.progenitusPositions_) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize guildLevel_
	{
		if this.guildLevel_ != 0 {
			err := prpc.Write(buffer, this.guildLevel_)
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
	// serialize guildId_
	{
		if this.guildId_ != 0 {
			err := prpc.Write(buffer, this.guildId_)
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
	// serialize fundz_
	{
		if this.fundz_ != 0 {
			err := prpc.Write(buffer, this.fundz_)
			if err != nil {
				return err
			}
		}
	}
	// serialize presentNum_
	{
		if this.presentNum_ != 0 {
			err := prpc.Write(buffer, this.presentNum_)
			if err != nil {
				return err
			}
		}
	}
	// serialize master_
	{
		if this.master_ != 0 {
			err := prpc.Write(buffer, this.master_)
			if err != nil {
				return err
			}
		}
	}
	// serialize masterName_
	if len(this.masterName_) != 0 {
		err := prpc.Write(buffer, this.masterName_)
		if err != nil {
			return err
		}
	}
	// serialize guildName_
	if len(this.guildName_) != 0 {
		err := prpc.Write(buffer, this.guildName_)
		if err != nil {
			return err
		}
	}
	// serialize notice_
	if len(this.notice_) != 0 {
		err := prpc.Write(buffer, this.notice_)
		if err != nil {
			return err
		}
	}
	// serialize requestList_
	if len(this.requestList_) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.requestList_)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.requestList_ {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// serialize noFundzDays_
	{
		if this.noFundzDays_ != 0 {
			err := prpc.Write(buffer, this.noFundzDays_)
			if err != nil {
				return err
			}
		}
	}
	// serialize buildings_
	if len(this.buildings_) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.buildings_)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.buildings_ {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// serialize progenitus_
	if len(this.progenitus_) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.progenitus_)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.progenitus_ {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// serialize progenitusPositions_
	if len(this.progenitusPositions_) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.progenitusPositions_)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.progenitusPositions_ {
			err := prpc.Write(buffer, value)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *COM_Guild) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 2)
	if err != nil {
		return err
	}
	// deserialize guildLevel_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.guildLevel_)
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
	// deserialize guildId_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.guildId_)
		if err != nil {
			return err
		}
	}
	// deserialize guildContribution_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.guildContribution_)
		if err != nil {
			return err
		}
	}
	// deserialize fundz_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.fundz_)
		if err != nil {
			return err
		}
	}
	// deserialize presentNum_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.presentNum_)
		if err != nil {
			return err
		}
	}
	// deserialize master_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.master_)
		if err != nil {
			return err
		}
	}
	// deserialize masterName_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.masterName_)
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
	// deserialize notice_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.notice_)
		if err != nil {
			return err
		}
	}
	// deserialize requestList_
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.requestList_ = make([]COM_GuildRequestData, size)
		for i, _ := range this.requestList_ {
			err := this.requestList_[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// deserialize noFundzDays_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.noFundzDays_)
		if err != nil {
			return err
		}
	}
	// deserialize buildings_
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.buildings_ = make([]COM_GuildBuilding, size)
		for i, _ := range this.buildings_ {
			err := this.buildings_[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// deserialize progenitus_
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.progenitus_ = make([]COM_GuildProgen, size)
		for i, _ := range this.progenitus_ {
			err := this.progenitus_[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// deserialize progenitusPositions_
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.progenitusPositions_ = make([]int32, size)
		for i, _ := range this.progenitusPositions_ {
			err := prpc.Read(buffer, &this.progenitusPositions_[i])
			if err != nil {
				return err
			}
		}
	}
	return nil
}
