package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type SGE_GmtOrder struct {
	shopId_  int32   //0
	orderId_ string  //1
	payment_ float32 //2
}

func (this *SGE_GmtOrder) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.shopId_ != 0)
	mask.WriteBit(len(this.orderId_) != 0)
	mask.WriteBit(this.payment_ != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize shopId_
	{
		if this.shopId_ != 0 {
			err := prpc.Write(buffer, this.shopId_)
			if err != nil {
				return err
			}
		}
	}
	// serialize orderId_
	if len(this.orderId_) != 0 {
		err := prpc.Write(buffer, this.orderId_)
		if err != nil {
			return err
		}
	}
	// serialize payment_
	{
		if this.payment_ != 0 {
			err := prpc.Write(buffer, this.payment_)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *SGE_GmtOrder) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize shopId_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.shopId_)
		if err != nil {
			return err
		}
	}
	// deserialize orderId_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.orderId_)
		if err != nil {
			return err
		}
	}
	// deserialize payment_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.payment_)
		if err != nil {
			return err
		}
	}
	return nil
}
