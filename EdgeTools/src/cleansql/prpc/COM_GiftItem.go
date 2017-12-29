package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type COM_GiftItem struct {
	itemId_  uint32 //0
	itemNum_ uint32 //1
}

func (this *COM_GiftItem) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.itemId_ != 0)
	mask.WriteBit(this.itemNum_ != 0)
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
	// serialize itemNum_
	{
		if this.itemNum_ != 0 {
			err := prpc.Write(buffer, this.itemNum_)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *COM_GiftItem) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
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
	// deserialize itemNum_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.itemNum_)
		if err != nil {
			return err
		}
	}
	return nil
}
