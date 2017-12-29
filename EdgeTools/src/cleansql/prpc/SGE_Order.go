package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type SGE_Order struct {
	gamename_  string  //0
	pfid_      string  //1
	pfname_    string  //2
	orderid_   string  //3
	roleId_    uint32  //4
	rolelv_    uint32  //5
	accountid_ string  //6
	payment_   float32 //7
	paytime_   string  //8
}

func (this *SGE_Order) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(2)
	mask.WriteBit(len(this.gamename_) != 0)
	mask.WriteBit(len(this.pfid_) != 0)
	mask.WriteBit(len(this.pfname_) != 0)
	mask.WriteBit(len(this.orderid_) != 0)
	mask.WriteBit(this.roleId_ != 0)
	mask.WriteBit(this.rolelv_ != 0)
	mask.WriteBit(len(this.accountid_) != 0)
	mask.WriteBit(this.payment_ != 0)
	mask.WriteBit(len(this.paytime_) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize gamename_
	if len(this.gamename_) != 0 {
		err := prpc.Write(buffer, this.gamename_)
		if err != nil {
			return err
		}
	}
	// serialize pfid_
	if len(this.pfid_) != 0 {
		err := prpc.Write(buffer, this.pfid_)
		if err != nil {
			return err
		}
	}
	// serialize pfname_
	if len(this.pfname_) != 0 {
		err := prpc.Write(buffer, this.pfname_)
		if err != nil {
			return err
		}
	}
	// serialize orderid_
	if len(this.orderid_) != 0 {
		err := prpc.Write(buffer, this.orderid_)
		if err != nil {
			return err
		}
	}
	// serialize roleId_
	{
		if this.roleId_ != 0 {
			err := prpc.Write(buffer, this.roleId_)
			if err != nil {
				return err
			}
		}
	}
	// serialize rolelv_
	{
		if this.rolelv_ != 0 {
			err := prpc.Write(buffer, this.rolelv_)
			if err != nil {
				return err
			}
		}
	}
	// serialize accountid_
	if len(this.accountid_) != 0 {
		err := prpc.Write(buffer, this.accountid_)
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
	// serialize paytime_
	if len(this.paytime_) != 0 {
		err := prpc.Write(buffer, this.paytime_)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_Order) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 2)
	if err != nil {
		return err
	}
	// deserialize gamename_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.gamename_)
		if err != nil {
			return err
		}
	}
	// deserialize pfid_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.pfid_)
		if err != nil {
			return err
		}
	}
	// deserialize pfname_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.pfname_)
		if err != nil {
			return err
		}
	}
	// deserialize orderid_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.orderid_)
		if err != nil {
			return err
		}
	}
	// deserialize roleId_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.roleId_)
		if err != nil {
			return err
		}
	}
	// deserialize rolelv_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.rolelv_)
		if err != nil {
			return err
		}
	}
	// deserialize accountid_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.accountid_)
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
	// deserialize paytime_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.paytime_)
		if err != nil {
			return err
		}
	}
	return nil
}
