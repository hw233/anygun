package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type COM_PropValue struct {
	type_  int     //0
	value_ float32 //1
}

func (this *COM_PropValue) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.type_ != 0)
	mask.WriteBit(this.value_ != 0)
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
	// serialize value_
	{
		if this.value_ != 0 {
			err := prpc.Write(buffer, this.value_)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *COM_PropValue) Deserialize(buffer *bytes.Buffer) error {
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
	// deserialize value_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.value_)
		if err != nil {
			return err
		}
	}
	return nil
}
