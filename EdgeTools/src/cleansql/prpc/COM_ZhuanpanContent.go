package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type COM_ZhuanpanContent struct {
	id_      uint32 //0
	item_    uint32 //1
	itemNum_ uint32 //2
	rate_    uint32 //3
	maxdrop_ uint32 //4
}

func (this *COM_ZhuanpanContent) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.id_ != 0)
	mask.WriteBit(this.item_ != 0)
	mask.WriteBit(this.itemNum_ != 0)
	mask.WriteBit(this.rate_ != 0)
	mask.WriteBit(this.maxdrop_ != 0)
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
	// serialize item_
	{
		if this.item_ != 0 {
			err := prpc.Write(buffer, this.item_)
			if err != nil {
				return err
			}
		}
	}
	// serialize itemNum_
	{
		if this.itemNum_ != 0 {
			err := prpc.Write(buffer, this.itemNum_)
			if err != nil {
				return err
			}
		}
	}
	// serialize rate_
	{
		if this.rate_ != 0 {
			err := prpc.Write(buffer, this.rate_)
			if err != nil {
				return err
			}
		}
	}
	// serialize maxdrop_
	{
		if this.maxdrop_ != 0 {
			err := prpc.Write(buffer, this.maxdrop_)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *COM_ZhuanpanContent) Deserialize(buffer *bytes.Buffer) error {
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
	// deserialize item_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.item_)
		if err != nil {
			return err
		}
	}
	// deserialize itemNum_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.itemNum_)
		if err != nil {
			return err
		}
	}
	// deserialize rate_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.rate_)
		if err != nil {
			return err
		}
	}
	// deserialize maxdrop_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.maxdrop_)
		if err != nil {
			return err
		}
	}
	return nil
}
