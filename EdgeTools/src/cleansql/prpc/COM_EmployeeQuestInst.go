package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type COM_EmployeeQuestInst struct {
	status_        int     //0
	questId_       int32   //1
	timeout_       int32   //2
	usedEmployees_ []int32 //3
}

func (this *COM_EmployeeQuestInst) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.status_ != 0)
	mask.WriteBit(this.questId_ != 0)
	mask.WriteBit(this.timeout_ != 0)
	mask.WriteBit(len(this.usedEmployees_) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize status_
	{
		if this.status_ != 0 {
			err := prpc.Write(buffer, this.status_)
			if err != nil {
				return err
			}
		}
	}
	// serialize questId_
	{
		if this.questId_ != 0 {
			err := prpc.Write(buffer, this.questId_)
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
	// serialize usedEmployees_
	if len(this.usedEmployees_) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.usedEmployees_)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.usedEmployees_ {
			err := prpc.Write(buffer, value)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *COM_EmployeeQuestInst) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize status_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.status_)
		if err != nil {
			return err
		}
	}
	// deserialize questId_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.questId_)
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
	// deserialize usedEmployees_
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.usedEmployees_ = make([]int32, size)
		for i, _ := range this.usedEmployees_ {
			err := prpc.Read(buffer, &this.usedEmployees_[i])
			if err != nil {
				return err
			}
		}
	}
	return nil
}
