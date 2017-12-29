package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type SGE_PlayerEmployeeQuestArray struct {
	value_ []SGE_EmployeeQuestInst //0
}

func (this *SGE_PlayerEmployeeQuestArray) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.value_) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize value_
	if len(this.value_) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.value_)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.value_ {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *SGE_PlayerEmployeeQuestArray) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize value_
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.value_ = make([]SGE_EmployeeQuestInst, size)
		for i, _ := range this.value_ {
			err := this.value_[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
