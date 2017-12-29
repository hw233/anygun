package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type COM_FPosition struct {
	x_      float32 //0
	z_      float32 //1
	isLast_ bool    //2
}

func (this *COM_FPosition) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.x_ != 0)
	mask.WriteBit(this.z_ != 0)
	mask.WriteBit(this.isLast_)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize x_
	{
		if this.x_ != 0 {
			err := prpc.Write(buffer, this.x_)
			if err != nil {
				return err
			}
		}
	}
	// serialize z_
	{
		if this.z_ != 0 {
			err := prpc.Write(buffer, this.z_)
			if err != nil {
				return err
			}
		}
	}
	// serialize isLast_
	{
	}
	return nil
}
func (this *COM_FPosition) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize x_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.x_)
		if err != nil {
			return err
		}
	}
	// deserialize z_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.z_)
		if err != nil {
			return err
		}
	}
	// deserialize isLast_
	this.isLast_ = mask.ReadBit()
	return nil
}
