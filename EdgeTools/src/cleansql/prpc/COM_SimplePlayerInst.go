package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type COM_SimplePlayerInst struct {
	COM_Entity
	isLeavingTeam_     bool               //0
	isBattle_          bool               //1
	autoBattle_        bool               //2
	isTeamLeader_      bool               //3
	sceneId_           int32              //4
	openSubSystemFlag_ int32              //5
	createTime_        int64              //6
	guildName_         string             //7
	scenePos_          COM_FPosition      //8
	pvpInfo_           COM_PlayerVsPlayer //9
	babies1_           []COM_BabyInst     //10
	battleEmps_        []COM_EmployeeInst //11
}

func (this *COM_SimplePlayerInst) Serialize(buffer *bytes.Buffer) error {
	{
		err := this.COM_Entity.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	//field mask
	mask := prpc.NewMask1(2)
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
	mask.WriteBit(len(this.babies1_) != 0)
	mask.WriteBit(len(this.battleEmps_) != 0)
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
	// serialize babies1_
	if len(this.babies1_) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.babies1_)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.babies1_ {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// serialize battleEmps_
	if len(this.battleEmps_) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.battleEmps_)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.battleEmps_ {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *COM_SimplePlayerInst) Deserialize(buffer *bytes.Buffer) error {
	{
		this.COM_Entity.Deserialize(buffer)
	}
	//field mask
	mask, err := prpc.NewMask0(buffer, 2)
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
	// deserialize babies1_
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.babies1_ = make([]COM_BabyInst, size)
		for i, _ := range this.babies1_ {
			err := this.babies1_[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// deserialize battleEmps_
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.battleEmps_ = make([]COM_EmployeeInst, size)
		for i, _ := range this.battleEmps_ {
			err := this.battleEmps_[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
