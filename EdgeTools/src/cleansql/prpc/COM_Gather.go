package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type COM_Gather struct {
	gatherId_ uint32 //0
	flag_     int    //1
	num_      uint32 //2
}

func (this *COM_Gather) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.gatherId_ != 0)
	mask.WriteBit(this.flag_ != 0)
	mask.WriteBit(this.num_ != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize gatherId_
	{
		if this.gatherId_ != 0 {
			err := prpc.Write(buffer, this.gatherId_)
			if err != nil {
				return err
			}
		}
	}
	// serialize flag_
	{
		if this.flag_ != 0 {
			err := prpc.Write(buffer, this.flag_)
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
func (this *COM_Gather) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize gatherId_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.gatherId_)
		if err != nil {
			return err
		}
	}
	// deserialize flag_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.flag_)
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
