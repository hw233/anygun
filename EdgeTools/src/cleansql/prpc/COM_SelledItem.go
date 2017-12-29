package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type COM_SelledItem struct {
	guid_         int32 //0
	sellPlayerId_ int32 //1
	sellTime_     int32 //2
	selledTime_   int32 //3
	itemId_       int32 //4
	babyId_       int32 //5
	itemStack_    int32 //6
	price_        int32 //7
	tax_          int32 //8
}

func (this *COM_SelledItem) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(2)
	mask.WriteBit(this.guid_ != 0)
	mask.WriteBit(this.sellPlayerId_ != 0)
	mask.WriteBit(this.sellTime_ != 0)
	mask.WriteBit(this.selledTime_ != 0)
	mask.WriteBit(this.itemId_ != 0)
	mask.WriteBit(this.babyId_ != 0)
	mask.WriteBit(this.itemStack_ != 0)
	mask.WriteBit(this.price_ != 0)
	mask.WriteBit(this.tax_ != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize guid_
	{
		if this.guid_ != 0 {
			err := prpc.Write(buffer, this.guid_)
			if err != nil {
				return err
			}
		}
	}
	// serialize sellPlayerId_
	{
		if this.sellPlayerId_ != 0 {
			err := prpc.Write(buffer, this.sellPlayerId_)
			if err != nil {
				return err
			}
		}
	}
	// serialize sellTime_
	{
		if this.sellTime_ != 0 {
			err := prpc.Write(buffer, this.sellTime_)
			if err != nil {
				return err
			}
		}
	}
	// serialize selledTime_
	{
		if this.selledTime_ != 0 {
			err := prpc.Write(buffer, this.selledTime_)
			if err != nil {
				return err
			}
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
	// serialize babyId_
	{
		if this.babyId_ != 0 {
			err := prpc.Write(buffer, this.babyId_)
			if err != nil {
				return err
			}
		}
	}
	// serialize itemStack_
	{
		if this.itemStack_ != 0 {
			err := prpc.Write(buffer, this.itemStack_)
			if err != nil {
				return err
			}
		}
	}
	// serialize price_
	{
		if this.price_ != 0 {
			err := prpc.Write(buffer, this.price_)
			if err != nil {
				return err
			}
		}
	}
	// serialize tax_
	{
		if this.tax_ != 0 {
			err := prpc.Write(buffer, this.tax_)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *COM_SelledItem) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 2)
	if err != nil {
		return err
	}
	// deserialize guid_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.guid_)
		if err != nil {
			return err
		}
	}
	// deserialize sellPlayerId_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.sellPlayerId_)
		if err != nil {
			return err
		}
	}
	// deserialize sellTime_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.sellTime_)
		if err != nil {
			return err
		}
	}
	// deserialize selledTime_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.selledTime_)
		if err != nil {
			return err
		}
	}
	// deserialize itemId_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.itemId_)
		if err != nil {
			return err
		}
	}
	// deserialize babyId_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.babyId_)
		if err != nil {
			return err
		}
	}
	// deserialize itemStack_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.itemStack_)
		if err != nil {
			return err
		}
	}
	// deserialize price_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.price_)
		if err != nil {
			return err
		}
	}
	// deserialize tax_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.tax_)
		if err != nil {
			return err
		}
	}
	return nil
}
