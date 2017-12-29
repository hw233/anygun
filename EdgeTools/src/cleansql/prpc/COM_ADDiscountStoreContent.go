package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type COM_ADDiscountStoreContent struct {
	price_    uint32  //0
	itemId_   uint32  //1
	discount_ float32 //2
	buyLimit_ uint32  //3
}

func (this *COM_ADDiscountStoreContent) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.price_ != 0)
	mask.WriteBit(this.itemId_ != 0)
	mask.WriteBit(this.discount_ != 0)
	mask.WriteBit(this.buyLimit_ != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
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
	// serialize itemId_
	{
		if this.itemId_ != 0 {
			err := prpc.Write(buffer, this.itemId_)
			if err != nil {
				return err
			}
		}
	}
	// serialize discount_
	{
		if this.discount_ != 0 {
			err := prpc.Write(buffer, this.discount_)
			if err != nil {
				return err
			}
		}
	}
	// serialize buyLimit_
	{
		if this.buyLimit_ != 0 {
			err := prpc.Write(buffer, this.buyLimit_)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *COM_ADDiscountStoreContent) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize price_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.price_)
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
	// deserialize discount_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.discount_)
		if err != nil {
			return err
		}
	}
	// deserialize buyLimit_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.buyLimit_)
		if err != nil {
			return err
		}
	}
	return nil
}
