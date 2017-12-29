package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type COM_GuildMember struct {
	level_            int8                //0
	shopRefreshTimes_ int8                //1
	guildId_          uint32              //2
	profType_         int8                //3
	profLevel_        int8                //4
	contribution_     int32               //5
	job_              int8                //6
	roleId_           int32               //7
	offlineTime_      uint32              //8
	roleName_         string              //9
	joinTime_         uint32              //10
	signflag_         bool                //11
	shopItems_        []COM_GuildShopItem //12
}

func (this *COM_GuildMember) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(2)
	mask.WriteBit(this.level_ != 0)
	mask.WriteBit(this.shopRefreshTimes_ != 0)
	mask.WriteBit(this.guildId_ != 0)
	mask.WriteBit(this.profType_ != 0)
	mask.WriteBit(this.profLevel_ != 0)
	mask.WriteBit(this.contribution_ != 0)
	mask.WriteBit(this.job_ != 0)
	mask.WriteBit(this.roleId_ != 0)
	mask.WriteBit(this.offlineTime_ != 0)
	mask.WriteBit(len(this.roleName_) != 0)
	mask.WriteBit(this.joinTime_ != 0)
	mask.WriteBit(this.signflag_)
	mask.WriteBit(len(this.shopItems_) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize level_
	{
		if this.level_ != 0 {
			err := prpc.Write(buffer, this.level_)
			if err != nil {
				return err
			}
		}
	}
	// serialize shopRefreshTimes_
	{
		if this.shopRefreshTimes_ != 0 {
			err := prpc.Write(buffer, this.shopRefreshTimes_)
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
	// serialize profType_
	{
		if this.profType_ != 0 {
			err := prpc.Write(buffer, this.profType_)
			if err != nil {
				return err
			}
		}
	}
	// serialize profLevel_
	{
		if this.profLevel_ != 0 {
			err := prpc.Write(buffer, this.profLevel_)
			if err != nil {
				return err
			}
		}
	}
	// serialize contribution_
	{
		if this.contribution_ != 0 {
			err := prpc.Write(buffer, this.contribution_)
			if err != nil {
				return err
			}
		}
	}
	// serialize job_
	{
		if this.job_ != 0 {
			err := prpc.Write(buffer, this.job_)
			if err != nil {
				return err
			}
		}
	}
	// serialize roleId_
	{
		if this.roleId_ != 0 {
			err := prpc.Write(buffer, this.roleId_)
			if err != nil {
				return err
			}
		}
	}
	// serialize offlineTime_
	{
		if this.offlineTime_ != 0 {
			err := prpc.Write(buffer, this.offlineTime_)
			if err != nil {
				return err
			}
		}
	}
	// serialize roleName_
	if len(this.roleName_) != 0 {
		err := prpc.Write(buffer, this.roleName_)
		if err != nil {
			return err
		}
	}
	// serialize joinTime_
	{
		if this.joinTime_ != 0 {
			err := prpc.Write(buffer, this.joinTime_)
			if err != nil {
				return err
			}
		}
	}
	// serialize signflag_
	{
	}
	// serialize shopItems_
	if len(this.shopItems_) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.shopItems_)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.shopItems_ {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *COM_GuildMember) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 2)
	if err != nil {
		return err
	}
	// deserialize level_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.level_)
		if err != nil {
			return err
		}
	}
	// deserialize shopRefreshTimes_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.shopRefreshTimes_)
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
	// deserialize profType_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.profType_)
		if err != nil {
			return err
		}
	}
	// deserialize profLevel_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.profLevel_)
		if err != nil {
			return err
		}
	}
	// deserialize contribution_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.contribution_)
		if err != nil {
			return err
		}
	}
	// deserialize job_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.job_)
		if err != nil {
			return err
		}
	}
	// deserialize roleId_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.roleId_)
		if err != nil {
			return err
		}
	}
	// deserialize offlineTime_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.offlineTime_)
		if err != nil {
			return err
		}
	}
	// deserialize roleName_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.roleName_)
		if err != nil {
			return err
		}
	}
	// deserialize joinTime_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.joinTime_)
		if err != nil {
			return err
		}
	}
	// deserialize signflag_
	this.signflag_ = mask.ReadBit()
	// deserialize shopItems_
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.shopItems_ = make([]COM_GuildShopItem, size)
		for i, _ := range this.shopItems_ {
			err := this.shopItems_[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
