package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type COM_BattleEmp struct {
	empBattleGroup_ int      //0
	employeeGroup1_ []uint32 //1
	employeeGroup2_ []uint32 //2
}

func (this *COM_BattleEmp) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.empBattleGroup_ != 0)
	mask.WriteBit(len(this.employeeGroup1_) != 0)
	mask.WriteBit(len(this.employeeGroup2_) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize empBattleGroup_
	{
		if this.empBattleGroup_ != 0 {
			err := prpc.Write(buffer, this.empBattleGroup_)
			if err != nil {
				return err
			}
		}
	}
	// serialize employeeGroup1_
	if len(this.employeeGroup1_) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.employeeGroup1_)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.employeeGroup1_ {
			err := prpc.Write(buffer, value)
			if err != nil {
				return err
			}
		}
	}
	// serialize employeeGroup2_
	if len(this.employeeGroup2_) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.employeeGroup2_)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.employeeGroup2_ {
			err := prpc.Write(buffer, value)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *COM_BattleEmp) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize empBattleGroup_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.empBattleGroup_)
		if err != nil {
			return err
		}
	}
	// deserialize employeeGroup1_
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.employeeGroup1_ = make([]uint32, size)
		for i, _ := range this.employeeGroup1_ {
			err := prpc.Read(buffer, &this.employeeGroup1_[i])
			if err != nil {
				return err
			}
		}
	}
	// deserialize employeeGroup2_
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.employeeGroup2_ = make([]uint32, size)
		for i, _ := range this.employeeGroup2_ {
			err := prpc.Read(buffer, &this.employeeGroup2_[i])
			if err != nil {
				return err
			}
		}
	}
	return nil
}
