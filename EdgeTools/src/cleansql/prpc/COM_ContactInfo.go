package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type COM_ContactInfo struct {
	instId_   uint32 //0
	name_     string //1
	level_    uint32 //2
	exp_      uint32 //3
	job_      int    //4
	assetId_  uint32 //5
	jobLevel_ uint32 //6
	vip_      int    //7
	ff_       uint32 //8
	rank_     uint32 //9
	section_  int32  //10
	value_    uint32 //11
	isLine_   bool   //12
}

func (this *COM_ContactInfo) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(2)
	mask.WriteBit(this.instId_ != 0)
	mask.WriteBit(len(this.name_) != 0)
	mask.WriteBit(this.level_ != 0)
	mask.WriteBit(this.exp_ != 0)
	mask.WriteBit(this.job_ != 0)
	mask.WriteBit(this.assetId_ != 0)
	mask.WriteBit(this.jobLevel_ != 0)
	mask.WriteBit(this.vip_ != 0)
	mask.WriteBit(this.ff_ != 0)
	mask.WriteBit(this.rank_ != 0)
	mask.WriteBit(this.section_ != 0)
	mask.WriteBit(this.value_ != 0)
	mask.WriteBit(this.isLine_)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize instId_
	{
		if this.instId_ != 0 {
			err := prpc.Write(buffer, this.instId_)
			if err != nil {
				return err
			}
		}
	}
	// serialize name_
	if len(this.name_) != 0 {
		err := prpc.Write(buffer, this.name_)
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
	// serialize exp_
	{
		if this.exp_ != 0 {
			err := prpc.Write(buffer, this.exp_)
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
	// serialize assetId_
	{
		if this.assetId_ != 0 {
			err := prpc.Write(buffer, this.assetId_)
			if err != nil {
				return err
			}
		}
	}
	// serialize jobLevel_
	{
		if this.jobLevel_ != 0 {
			err := prpc.Write(buffer, this.jobLevel_)
			if err != nil {
				return err
			}
		}
	}
	// serialize vip_
	{
		if this.vip_ != 0 {
			err := prpc.Write(buffer, this.vip_)
			if err != nil {
				return err
			}
		}
	}
	// serialize ff_
	{
		if this.ff_ != 0 {
			err := prpc.Write(buffer, this.ff_)
			if err != nil {
				return err
			}
		}
	}
	// serialize rank_
	{
		if this.rank_ != 0 {
			err := prpc.Write(buffer, this.rank_)
			if err != nil {
				return err
			}
		}
	}
	// serialize section_
	{
		if this.section_ != 0 {
			err := prpc.Write(buffer, this.section_)
			if err != nil {
				return err
			}
		}
	}
	// serialize value_
	{
		if this.value_ != 0 {
			err := prpc.Write(buffer, this.value_)
			if err != nil {
				return err
			}
		}
	}
	// serialize isLine_
	{
	}
	return nil
}
func (this *COM_ContactInfo) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 2)
	if err != nil {
		return err
	}
	// deserialize instId_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.instId_)
		if err != nil {
			return err
		}
	}
	// deserialize name_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.name_)
		if err != nil {
			return err
		}
	}
	// deserialize level_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.level_)
		if err != nil {
			return err
		}
	}
	// deserialize exp_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.exp_)
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
	// deserialize assetId_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.assetId_)
		if err != nil {
			return err
		}
	}
	// deserialize jobLevel_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.jobLevel_)
		if err != nil {
			return err
		}
	}
	// deserialize vip_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.vip_)
		if err != nil {
			return err
		}
	}
	// deserialize ff_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.ff_)
		if err != nil {
			return err
		}
	}
	// deserialize rank_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.rank_)
		if err != nil {
			return err
		}
	}
	// deserialize section_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.section_)
		if err != nil {
			return err
		}
	}
	// deserialize value_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.value_)
		if err != nil {
			return err
		}
	}
	// deserialize isLine_
	this.isLine_ = mask.ReadBit()
	return nil
}
