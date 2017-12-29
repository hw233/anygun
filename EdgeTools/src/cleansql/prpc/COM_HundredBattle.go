package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type COM_HundredBattle struct {
	playerId_ uint32 //0
	tier_     uint32 //1
	curTier_  uint32 //2
	surplus_  uint32 //3
	resetNum_ uint32 //4
}

func (this *COM_HundredBattle) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.playerId_ != 0)
	mask.WriteBit(this.tier_ != 0)
	mask.WriteBit(this.curTier_ != 0)
	mask.WriteBit(this.surplus_ != 0)
	mask.WriteBit(this.resetNum_ != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize playerId_
	{
		if this.playerId_ != 0 {
			err := prpc.Write(buffer, this.playerId_)
			if err != nil {
				return err
			}
		}
	}
	// serialize tier_
	{
		if this.tier_ != 0 {
			err := prpc.Write(buffer, this.tier_)
			if err != nil {
				return err
			}
		}
	}
	// serialize curTier_
	{
		if this.curTier_ != 0 {
			err := prpc.Write(buffer, this.curTier_)
			if err != nil {
				return err
			}
		}
	}
	// serialize surplus_
	{
		if this.surplus_ != 0 {
			err := prpc.Write(buffer, this.surplus_)
			if err != nil {
				return err
			}
		}
	}
	// serialize resetNum_
	{
		if this.resetNum_ != 0 {
			err := prpc.Write(buffer, this.resetNum_)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *COM_HundredBattle) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize playerId_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.playerId_)
		if err != nil {
			return err
		}
	}
	// deserialize tier_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.tier_)
		if err != nil {
			return err
		}
	}
	// deserialize curTier_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.curTier_)
		if err != nil {
			return err
		}
	}
	// deserialize surplus_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.surplus_)
		if err != nil {
			return err
		}
	}
	// deserialize resetNum_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.resetNum_)
		if err != nil {
			return err
		}
	}
	return nil
}
