package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type COM_IntegralContent struct {
	id_     uint32 //0
	itemid_ uint32 //1
	times_  int32  //2
	cost_   int32  //3
}

func (this *COM_IntegralContent) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.id_ != 0)
	mask.WriteBit(this.itemid_ != 0)
	mask.WriteBit(this.times_ != 0)
	mask.WriteBit(this.cost_ != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize id_
	{
		if this.id_ != 0 {
			err := prpc.Write(buffer, this.id_)
			if err != nil {
				return err
			}
		}
	}
	// serialize itemid_
	{
		if this.itemid_ != 0 {
			err := prpc.Write(buffer, this.itemid_)
			if err != nil {
				return err
			}
		}
	}
	// serialize times_
	{
		if this.times_ != 0 {
			err := prpc.Write(buffer, this.times_)
			if err != nil {
				return err
			}
		}
	}
	// serialize cost_
	{
		if this.cost_ != 0 {
			err := prpc.Write(buffer, this.cost_)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *COM_IntegralContent) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize id_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.id_)
		if err != nil {
			return err
		}
	}
	// deserialize itemid_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.itemid_)
		if err != nil {
			return err
		}
	}
	// deserialize times_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.times_)
		if err != nil {
			return err
		}
	}
	// deserialize cost_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.cost_)
		if err != nil {
			return err
		}
	}
	return nil
}
