package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type COM_QuestInst struct {
	questId_ int32             //0
	targets_ []COM_QuestTarget //1
}

func (this *COM_QuestInst) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.questId_ != 0)
	mask.WriteBit(len(this.targets_) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
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
	// serialize targets_
	if len(this.targets_) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.targets_)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.targets_ {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *COM_QuestInst) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize questId_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.questId_)
		if err != nil {
			return err
		}
	}
	// deserialize targets_
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.targets_ = make([]COM_QuestTarget, size)
		for i, _ := range this.targets_ {
			err := this.targets_[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
