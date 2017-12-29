package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type COM_Exam struct {
	questionIndex_ uint8    //0
	rightNum_      uint8    //1
	errorNum_      uint8    //2
	money_         uint32   //3
	exp_           uint32   //4
	questions_     []uint32 //5
}

func (this *COM_Exam) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.questionIndex_ != 0)
	mask.WriteBit(this.rightNum_ != 0)
	mask.WriteBit(this.errorNum_ != 0)
	mask.WriteBit(this.money_ != 0)
	mask.WriteBit(this.exp_ != 0)
	mask.WriteBit(len(this.questions_) != 0)
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
	// serialize rightNum_
	{
		if this.rightNum_ != 0 {
			err := prpc.Write(buffer, this.rightNum_)
			if err != nil {
				return err
			}
		}
	}
	// serialize errorNum_
	{
		if this.errorNum_ != 0 {
			err := prpc.Write(buffer, this.errorNum_)
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
	// serialize questions_
	if len(this.questions_) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.questions_)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.questions_ {
			err := prpc.Write(buffer, value)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *COM_Exam) Deserialize(buffer *bytes.Buffer) error {
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
	// deserialize rightNum_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.rightNum_)
		if err != nil {
			return err
		}
	}
	// deserialize errorNum_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.errorNum_)
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
	// deserialize questions_
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.questions_ = make([]uint32, size)
		for i, _ := range this.questions_ {
			err := prpc.Read(buffer, &this.questions_[i])
			if err != nil {
				return err
			}
		}
	}
	return nil
}
