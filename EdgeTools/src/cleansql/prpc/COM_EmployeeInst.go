package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type COM_EmployeeInst struct {
	COM_Entity
	ownerName_ string //0
	isBattle_  bool   //1
	weaponId_  uint32 //2
	quality_   int    //3
	star_      uint32 //4
	soul_      uint32 //5
}

func (this *COM_EmployeeInst) Serialize(buffer *bytes.Buffer) error {
	{
		err := this.COM_Entity.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.ownerName_) != 0)
	mask.WriteBit(this.isBattle_)
	mask.WriteBit(this.weaponId_ != 0)
	mask.WriteBit(this.quality_ != 0)
	mask.WriteBit(this.star_ != 0)
	mask.WriteBit(this.soul_ != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize ownerName_
	if len(this.ownerName_) != 0 {
		err := prpc.Write(buffer, this.ownerName_)
		if err != nil {
			return err
		}
	}
	// serialize isBattle_
	{
	}
	// serialize weaponId_
	{
		if this.weaponId_ != 0 {
			err := prpc.Write(buffer, this.weaponId_)
			if err != nil {
				return err
			}
		}
	}
	// serialize quality_
	{
		if this.quality_ != 0 {
			err := prpc.Write(buffer, this.quality_)
			if err != nil {
				return err
			}
		}
	}
	// serialize star_
	{
		if this.star_ != 0 {
			err := prpc.Write(buffer, this.star_)
			if err != nil {
				return err
			}
		}
	}
	// serialize soul_
	{
		if this.soul_ != 0 {
			err := prpc.Write(buffer, this.soul_)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *COM_EmployeeInst) Deserialize(buffer *bytes.Buffer) error {
	{
		this.COM_Entity.Deserialize(buffer)
	}
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize ownerName_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.ownerName_)
		if err != nil {
			return err
		}
	}
	// deserialize isBattle_
	this.isBattle_ = mask.ReadBit()
	// deserialize weaponId_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.weaponId_)
		if err != nil {
			return err
		}
	}
	// deserialize quality_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.quality_)
		if err != nil {
			return err
		}
	}
	// deserialize star_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.star_)
		if err != nil {
			return err
		}
	}
	// deserialize soul_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.soul_)
		if err != nil {
			return err
		}
	}
	return nil
}
