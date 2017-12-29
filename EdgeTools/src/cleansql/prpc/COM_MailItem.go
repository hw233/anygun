package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type COM_MailItem struct {
	itemId_    int32 //0
	itemStack_ int32 //1
}

func (this *COM_MailItem) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.itemId_ != 0)
	mask.WriteBit(this.itemStack_ != 0)
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
	// serialize itemStack_
	{
		if this.itemStack_ != 0 {
			err := prpc.Write(buffer, this.itemStack_)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *COM_MailItem) Deserialize(buffer *bytes.Buffer) error {
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
	// deserialize itemStack_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.itemStack_)
		if err != nil {
			return err
		}
	}
	return nil
}
