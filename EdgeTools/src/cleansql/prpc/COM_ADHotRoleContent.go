package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type COM_ADHotRoleContent struct {
	type_   int    //0
	buyNum_ uint32 //1
	roleId_ uint32 //2
	price_  uint32 //3
}

func (this *COM_ADHotRoleContent) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.type_ != 0)
	mask.WriteBit(this.buyNum_ != 0)
	mask.WriteBit(this.roleId_ != 0)
	mask.WriteBit(this.price_ != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize type_
	{
		if this.type_ != 0 {
			err := prpc.Write(buffer, this.type_)
			if err != nil {
				return err
			}
		}
	}
	// serialize buyNum_
	{
		if this.buyNum_ != 0 {
			err := prpc.Write(buffer, this.buyNum_)
			if err != nil {
				return err
			}
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
	// serialize price_
	{
		if this.price_ != 0 {
			err := prpc.Write(buffer, this.price_)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *COM_ADHotRoleContent) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize type_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.type_)
		if err != nil {
			return err
		}
	}
	// deserialize buyNum_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.buyNum_)
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
	// deserialize price_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.price_)
		if err != nil {
			return err
		}
	}
	return nil
}
