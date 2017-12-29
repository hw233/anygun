package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type COM_ReportTargetBase struct {
	position_ int           //0
	prop_     COM_PropValue //1
	bao_      bool          //2
	fly_      bool          //3
}

func (this *COM_ReportTargetBase) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.position_ != 0)
	mask.WriteBit(true) //prop_
	mask.WriteBit(this.bao_)
	mask.WriteBit(this.fly_)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize position_
	{
		if this.position_ != 0 {
			err := prpc.Write(buffer, this.position_)
			if err != nil {
				return err
			}
		}
	}
	// serialize prop_
	{
		err := this.prop_.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	// serialize bao_
	{
	}
	// serialize fly_
	{
	}
	return nil
}
func (this *COM_ReportTargetBase) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize position_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.position_)
		if err != nil {
			return err
		}
	}
	// deserialize prop_
	if mask.ReadBit() {
		err := this.prop_.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	// deserialize bao_
	this.bao_ = mask.ReadBit()
	// deserialize fly_
	this.fly_ = mask.ReadBit()
	return nil
}
