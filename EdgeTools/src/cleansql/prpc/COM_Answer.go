package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type COM_Answer struct {
	questionIndex_ uint8  //0
	money_         uint32 //1
	exp_           uint32 //2
	isRigth_       bool   //3
}

func (this *COM_Answer) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.questionIndex_ != 0)
	mask.WriteBit(this.money_ != 0)
	mask.WriteBit(this.exp_ != 0)
	mask.WriteBit(this.isRigth_)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize questionIndex_
	{
		if this.questionIndex_ != 0 {
			err := prpc.Write(buffer, this.questionIndex_)
			if err != nil {
				return err
			}
		}
	}
	// serialize money_
	{
		if this.money_ != 0 {
			err := prpc.Write(buffer, this.money_)
			if err != nil {
				return err
			}
		}
	}
	// serialize exp_
	{
		if this.exp_ != 0 {
			err := prpc.Write(buffer, this.exp_)
			if err != nil {
				return err
			}
		}
	}
	// serialize isRigth_
	{
	}
	return nil
}
func (this *COM_Answer) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize questionIndex_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.questionIndex_)
		if err != nil {
			return err
		}
	}
	// deserialize money_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.money_)
		if err != nil {
			return err
		}
	}
	// deserialize exp_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.exp_)
		if err != nil {
			return err
		}
	}
	// deserialize isRigth_
	this.isRigth_ = mask.ReadBit()
	return nil
}
