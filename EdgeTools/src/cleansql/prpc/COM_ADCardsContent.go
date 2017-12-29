package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type COM_ADCardsContent struct {
	count_    uint32 //0
	rewardId_ uint32 //1
}

func (this *COM_ADCardsContent) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.count_ != 0)
	mask.WriteBit(this.rewardId_ != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize count_
	{
		if this.count_ != 0 {
			err := prpc.Write(buffer, this.count_)
			if err != nil {
				return err
			}
		}
	}
	// serialize rewardId_
	{
		if this.rewardId_ != 0 {
			err := prpc.Write(buffer, this.rewardId_)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *COM_ADCardsContent) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize count_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.count_)
		if err != nil {
			return err
		}
	}
	// deserialize rewardId_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.rewardId_)
		if err != nil {
			return err
		}
	}
	return nil
}
