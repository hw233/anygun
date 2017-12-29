package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type COM_CourseGift struct {
	id_      uint32  //0
	timeout_ float32 //1
}

func (this *COM_CourseGift) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.id_ != 0)
	mask.WriteBit(this.timeout_ != 0)
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
	// serialize timeout_
	{
		if this.timeout_ != 0 {
			err := prpc.Write(buffer, this.timeout_)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *COM_CourseGift) Deserialize(buffer *bytes.Buffer) error {
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
	// deserialize timeout_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.timeout_)
		if err != nil {
			return err
		}
	}
	return nil
}
