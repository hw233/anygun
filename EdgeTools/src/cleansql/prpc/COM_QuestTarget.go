package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type COM_QuestTarget struct {
	targetId_  uint32 //0
	targetNum_ int32  //1
}

func (this *COM_QuestTarget) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.targetId_ != 0)
	mask.WriteBit(this.targetNum_ != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize targetId_
	{
		if this.targetId_ != 0 {
			err := prpc.Write(buffer, this.targetId_)
			if err != nil {
				return err
			}
		}
	}
	// serialize targetNum_
	{
		if this.targetNum_ != 0 {
			err := prpc.Write(buffer, this.targetNum_)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *COM_QuestTarget) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize targetId_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.targetId_)
		if err != nil {
			return err
		}
	}
	// deserialize targetNum_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.targetNum_)
		if err != nil {
			return err
		}
	}
	return nil
}
