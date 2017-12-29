package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type COM_SimpleTeamInfo struct {
	COM_CreateTeamInfo
	teamId_        uint32 //0
	curMemberSize_ uint8  //1
	leaderName_    string //2
	job_           int    //3
	joblevel_      uint32 //4
	needPassword_  bool   //5
	isRunning_     bool   //6
	isWelcome_     bool   //7
}

func (this *COM_SimpleTeamInfo) Serialize(buffer *bytes.Buffer) error {
	{
		err := this.COM_CreateTeamInfo.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.teamId_ != 0)
	mask.WriteBit(this.curMemberSize_ != 0)
	mask.WriteBit(len(this.leaderName_) != 0)
	mask.WriteBit(this.job_ != 0)
	mask.WriteBit(this.joblevel_ != 0)
	mask.WriteBit(this.needPassword_)
	mask.WriteBit(this.isRunning_)
	mask.WriteBit(this.isWelcome_)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize teamId_
	{
		if this.teamId_ != 0 {
			err := prpc.Write(buffer, this.teamId_)
			if err != nil {
				return err
			}
		}
	}
	// serialize curMemberSize_
	{
		if this.curMemberSize_ != 0 {
			err := prpc.Write(buffer, this.curMemberSize_)
			if err != nil {
				return err
			}
		}
	}
	// serialize leaderName_
	if len(this.leaderName_) != 0 {
		err := prpc.Write(buffer, this.leaderName_)
		if err != nil {
			return err
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
	// serialize joblevel_
	{
		if this.joblevel_ != 0 {
			err := prpc.Write(buffer, this.joblevel_)
			if err != nil {
				return err
			}
		}
	}
	// serialize needPassword_
	{
	}
	// serialize isRunning_
	{
	}
	// serialize isWelcome_
	{
	}
	return nil
}
func (this *COM_SimpleTeamInfo) Deserialize(buffer *bytes.Buffer) error {
	{
		this.COM_CreateTeamInfo.Deserialize(buffer)
	}
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize teamId_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.teamId_)
		if err != nil {
			return err
		}
	}
	// deserialize curMemberSize_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.curMemberSize_)
		if err != nil {
			return err
		}
	}
	// deserialize leaderName_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.leaderName_)
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
	// deserialize joblevel_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.joblevel_)
		if err != nil {
			return err
		}
	}
	// deserialize needPassword_
	this.needPassword_ = mask.ReadBit()
	// deserialize isRunning_
	this.isRunning_ = mask.ReadBit()
	// deserialize isWelcome_
	this.isWelcome_ = mask.ReadBit()
	return nil
}
