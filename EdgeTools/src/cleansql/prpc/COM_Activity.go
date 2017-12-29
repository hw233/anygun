package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type COM_Activity struct {
	actId_   int32 //0
	counter_ int32 //1
}

func (this *COM_Activity) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.actId_ != 0)
	mask.WriteBit(this.counter_ != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize actId_
	{
		if this.actId_ != 0 {
			err := prpc.Write(buffer, this.actId_)
			if err != nil {
				return err
			}
		}
	}
	// serialize counter_
	{
		if this.counter_ != 0 {
			err := prpc.Write(buffer, this.counter_)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *COM_Activity) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize actId_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.actId_)
		if err != nil {
			return err
		}
	}
	// deserialize counter_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.counter_)
		if err != nil {
			return err
		}
	}
	return nil
}
