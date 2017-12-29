package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type COM_ExpendItem struct {
	itemInstId_ uint32 //0
	num_        uint32 //1
}

func (this *COM_ExpendItem) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.itemInstId_ != 0)
	mask.WriteBit(this.num_ != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize itemInstId_
	{
		if this.itemInstId_ != 0 {
			err := prpc.Write(buffer, this.itemInstId_)
			if err != nil {
				return err
			}
		}
	}
	// serialize num_
	{
		if this.num_ != 0 {
			err := prpc.Write(buffer, this.num_)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *COM_ExpendItem) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize itemInstId_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.itemInstId_)
		if err != nil {
			return err
		}
	}
	// deserialize num_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.num_)
		if err != nil {
			return err
		}
	}
	return nil
}
