package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type COM_Item struct {
	itemId_        uint32          //0
	instId_        uint32          //1
	stack_         int16           //2
	isBind_        bool            //3
	isLock_        bool            //4
	strLevel_      int8            //5
	slot_          int16           //6
	skillID_       int32           //7
	durability_    int32           //8
	durabilityMax_ int32           //9
	usedTimeout_   int32           //10
	lastSellTime_  int32           //11
	propArr        []COM_PropValue //12
}

func (this *COM_Item) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(2)
	mask.WriteBit(this.itemId_ != 0)
	mask.WriteBit(this.instId_ != 0)
	mask.WriteBit(this.stack_ != 0)
	mask.WriteBit(this.isBind_)
	mask.WriteBit(this.isLock_)
	mask.WriteBit(this.strLevel_ != 0)
	mask.WriteBit(this.slot_ != 0)
	mask.WriteBit(this.skillID_ != 0)
	mask.WriteBit(this.durability_ != 0)
	mask.WriteBit(this.durabilityMax_ != 0)
	mask.WriteBit(this.usedTimeout_ != 0)
	mask.WriteBit(this.lastSellTime_ != 0)
	mask.WriteBit(len(this.propArr) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize itemId_
	{
		if this.itemId_ != 0 {
			err := prpc.Write(buffer, this.itemId_)
			if err != nil {
				return err
			}
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
	// serialize stack_
	{
		if this.stack_ != 0 {
			err := prpc.Write(buffer, this.stack_)
			if err != nil {
				return err
			}
		}
	}
	// serialize isBind_
	{
	}
	// serialize isLock_
	{
	}
	// serialize strLevel_
	{
		if this.strLevel_ != 0 {
			err := prpc.Write(buffer, this.strLevel_)
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
	// serialize skillID_
	{
		if this.skillID_ != 0 {
			err := prpc.Write(buffer, this.skillID_)
			if err != nil {
				return err
			}
		}
	}
	// serialize durability_
	{
		if this.durability_ != 0 {
			err := prpc.Write(buffer, this.durability_)
			if err != nil {
				return err
			}
		}
	}
	// serialize durabilityMax_
	{
		if this.durabilityMax_ != 0 {
			err := prpc.Write(buffer, this.durabilityMax_)
			if err != nil {
				return err
			}
		}
	}
	// serialize usedTimeout_
	{
		if this.usedTimeout_ != 0 {
			err := prpc.Write(buffer, this.usedTimeout_)
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
	// serialize propArr
	if len(this.propArr) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.propArr)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.propArr {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *COM_Item) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 2)
	if err != nil {
		return err
	}
	// deserialize itemId_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.itemId_)
		if err != nil {
			return err
		}
	}
	// deserialize instId_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.instId_)
		if err != nil {
			return err
		}
	}
	// deserialize stack_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.stack_)
		if err != nil {
			return err
		}
	}
	// deserialize isBind_
	this.isBind_ = mask.ReadBit()
	// deserialize isLock_
	this.isLock_ = mask.ReadBit()
	// deserialize strLevel_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.strLevel_)
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
	// deserialize skillID_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.skillID_)
		if err != nil {
			return err
		}
	}
	// deserialize durability_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.durability_)
		if err != nil {
			return err
		}
	}
	// deserialize durabilityMax_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.durabilityMax_)
		if err != nil {
			return err
		}
	}
	// deserialize usedTimeout_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.usedTimeout_)
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
	// deserialize propArr
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.propArr = make([]COM_PropValue, size)
		for i, _ := range this.propArr {
			err := this.propArr[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
