package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type SGE_OrderInfo struct {
	productId_    int32   //0
	productCount_ int32   //1
	amount_       float32 //2
	orderId_      string  //3
	payTime_      string  //4
}

func (this *SGE_OrderInfo) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.productId_ != 0)
	mask.WriteBit(this.productCount_ != 0)
	mask.WriteBit(this.amount_ != 0)
	mask.WriteBit(len(this.orderId_) != 0)
	mask.WriteBit(len(this.payTime_) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize productId_
	{
		if this.productId_ != 0 {
			err := prpc.Write(buffer, this.productId_)
			if err != nil {
				return err
			}
		}
	}
	// serialize productCount_
	{
		if this.productCount_ != 0 {
			err := prpc.Write(buffer, this.productCount_)
			if err != nil {
				return err
			}
		}
	}
	// serialize amount_
	{
		if this.amount_ != 0 {
			err := prpc.Write(buffer, this.amount_)
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
	// serialize payTime_
	if len(this.payTime_) != 0 {
		err := prpc.Write(buffer, this.payTime_)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_OrderInfo) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize productId_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.productId_)
		if err != nil {
			return err
		}
	}
	// deserialize productCount_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.productCount_)
		if err != nil {
			return err
		}
	}
	// deserialize amount_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.amount_)
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
	// deserialize payTime_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.payTime_)
		if err != nil {
			return err
		}
	}
	return nil
}
