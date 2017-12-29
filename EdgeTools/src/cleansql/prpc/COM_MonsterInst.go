package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type COM_MonsterInst struct {
	COM_Entity
	gear_ []int32 //0
}

func (this *COM_MonsterInst) Serialize(buffer *bytes.Buffer) error {
	{
		err := this.COM_Entity.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.gear_) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize gear_
	if len(this.gear_) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.gear_)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.gear_ {
			err := prpc.Write(buffer, value)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *COM_MonsterInst) Deserialize(buffer *bytes.Buffer) error {
	{
		this.COM_Entity.Deserialize(buffer)
	}
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize gear_
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.gear_ = make([]int32, size)
		for i, _ := range this.gear_ {
			err := prpc.Read(buffer, &this.gear_[i])
			if err != nil {
				return err
			}
		}
	}
	return nil
}
