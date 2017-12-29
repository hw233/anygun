package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type SGE_EmployeeQuestInst struct {
	COM_EmployeeQuestInst
	doTime_      int64 //0
	refreshTime_ int64 //1
}

func (this *SGE_EmployeeQuestInst) Serialize(buffer *bytes.Buffer) error {
	{
		err := this.COM_EmployeeQuestInst.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.doTime_ != 0)
	mask.WriteBit(this.refreshTime_ != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize doTime_
	{
		if this.doTime_ != 0 {
			err := prpc.Write(buffer, this.doTime_)
			if err != nil {
				return err
			}
		}
	}
	// serialize refreshTime_
	{
		if this.refreshTime_ != 0 {
			err := prpc.Write(buffer, this.refreshTime_)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *SGE_EmployeeQuestInst) Deserialize(buffer *bytes.Buffer) error {
	{
		this.COM_EmployeeQuestInst.Deserialize(buffer)
	}
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize doTime_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.doTime_)
		if err != nil {
			return err
		}
	}
	// deserialize refreshTime_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.refreshTime_)
		if err != nil {
			return err
		}
	}
	return nil
}
