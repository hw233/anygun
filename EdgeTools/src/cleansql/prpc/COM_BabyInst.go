package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type COM_BabyInst struct {
	COM_Entity
	ownerName_      string    //0
	isShow_         bool      //1
	isBattle_       bool      //2
	isBind_         bool      //3
	isLock_         bool      //4
	tableId_        int32     //5
	slot_           int32     //6
	intensifyLevel_ uint32    //7
	intensifynum_   uint32    //8
	lastSellTime_   int32     //9
	gear_           []int32   //10
	addprop_        []float32 //11
}

func (this *COM_BabyInst) Serialize(buffer *bytes.Buffer) error {
	{
		err := this.COM_Entity.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	//field mask
	mask := prpc.NewMask1(2)
	mask.WriteBit(len(this.ownerName_) != 0)
	mask.WriteBit(this.isShow_)
	mask.WriteBit(this.isBattle_)
	mask.WriteBit(this.isBind_)
	mask.WriteBit(this.isLock_)
	mask.WriteBit(this.tableId_ != 0)
	mask.WriteBit(this.slot_ != 0)
	mask.WriteBit(this.intensifyLevel_ != 0)
	mask.WriteBit(this.intensifynum_ != 0)
	mask.WriteBit(this.lastSellTime_ != 0)
	mask.WriteBit(len(this.gear_) != 0)
	mask.WriteBit(len(this.addprop_) != 0)
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
	// serialize isShow_
	{
	}
	// serialize isBattle_
	{
	}
	// serialize isBind_
	{
	}
	// serialize isLock_
	{
	}
	// serialize tableId_
	{
		if this.tableId_ != 0 {
			err := prpc.Write(buffer, this.tableId_)
			if err != nil {
				return err
			}
		}
	}
	// serialize slot_
	{
		if this.slot_ != 0 {
			err := prpc.Write(buffer, this.slot_)
			if err != nil {
				return err
			}
		}
	}
	// serialize intensifyLevel_
	{
		if this.intensifyLevel_ != 0 {
			err := prpc.Write(buffer, this.intensifyLevel_)
			if err != nil {
				return err
			}
		}
	}
	// serialize intensifynum_
	{
		if this.intensifynum_ != 0 {
			err := prpc.Write(buffer, this.intensifynum_)
			if err != nil {
				return err
			}
		}
	}
	// serialize lastSellTime_
	{
		if this.lastSellTime_ != 0 {
			err := prpc.Write(buffer, this.lastSellTime_)
			if err != nil {
				return err
			}
		}
	}
	// serialize gear_
	if len(this.gear_) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.gear_)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.gear_ {
			err := prpc.Write(buffer, value)
			if err != nil {
				return err
			}
		}
	}
	// serialize addprop_
	if len(this.addprop_) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.addprop_)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.addprop_ {
			err := prpc.Write(buffer, value)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *COM_BabyInst) Deserialize(buffer *bytes.Buffer) error {
	{
		this.COM_Entity.Deserialize(buffer)
	}
	//field mask
	mask, err := prpc.NewMask0(buffer, 2)
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
	// deserialize isShow_
	this.isShow_ = mask.ReadBit()
	// deserialize isBattle_
	this.isBattle_ = mask.ReadBit()
	// deserialize isBind_
	this.isBind_ = mask.ReadBit()
	// deserialize isLock_
	this.isLock_ = mask.ReadBit()
	// deserialize tableId_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.tableId_)
		if err != nil {
			return err
		}
	}
	// deserialize slot_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.slot_)
		if err != nil {
			return err
		}
	}
	// deserialize intensifyLevel_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.intensifyLevel_)
		if err != nil {
			return err
		}
	}
	// deserialize intensifynum_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.intensifynum_)
		if err != nil {
			return err
		}
	}
	// deserialize lastSellTime_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.lastSellTime_)
		if err != nil {
			return err
		}
	}
	// deserialize gear_
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.gear_ = make([]int32, size)
		for i, _ := range this.gear_ {
			err := prpc.Read(buffer, &this.gear_[i])
			if err != nil {
				return err
			}
		}
	}
	// deserialize addprop_
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.addprop_ = make([]float32, size)
		for i, _ := range this.addprop_ {
			err := prpc.Read(buffer, &this.addprop_[i])
			if err != nil {
				return err
			}
		}
	}
	return nil
}
