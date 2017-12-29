package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type SGE_LogProduceTrack struct {
	timestamp_  uint64 //0
	from_       int32  //1
	playerId_   uint32 //2
	playerName_ string //3
	itemId_     uint32 //4
	itemInstId_ uint32 //5
	itemStack_  int32  //6
	diamond_    int32  //7
	money_      int32  //8
	exp_        int32  //9
	magic_      int32  //10
}

func (this *SGE_LogProduceTrack) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(2)
	mask.WriteBit(this.timestamp_ != 0)
	mask.WriteBit(this.from_ != 0)
	mask.WriteBit(this.playerId_ != 0)
	mask.WriteBit(len(this.playerName_) != 0)
	mask.WriteBit(this.itemId_ != 0)
	mask.WriteBit(this.itemInstId_ != 0)
	mask.WriteBit(this.itemStack_ != 0)
	mask.WriteBit(this.diamond_ != 0)
	mask.WriteBit(this.money_ != 0)
	mask.WriteBit(this.exp_ != 0)
	mask.WriteBit(this.magic_ != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize timestamp_
	{
		if this.timestamp_ != 0 {
			err := prpc.Write(buffer, this.timestamp_)
			if err != nil {
				return err
			}
		}
	}
	// serialize from_
	{
		if this.from_ != 0 {
			err := prpc.Write(buffer, this.from_)
			if err != nil {
				return err
			}
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
	// serialize playerName_
	if len(this.playerName_) != 0 {
		err := prpc.Write(buffer, this.playerName_)
		if err != nil {
			return err
		}
	}
	// serialize itemId_
	{
		if this.itemId_ != 0 {
			err := prpc.Write(buffer, this.itemId_)
			if err != nil {
				return err
			}
		}
	}
	// serialize itemInstId_
	{
		if this.itemInstId_ != 0 {
			err := prpc.Write(buffer, this.itemInstId_)
			if err != nil {
				return err
			}
		}
	}
	// serialize itemStack_
	{
		if this.itemStack_ != 0 {
			err := prpc.Write(buffer, this.itemStack_)
			if err != nil {
				return err
			}
		}
	}
	// serialize diamond_
	{
		if this.diamond_ != 0 {
			err := prpc.Write(buffer, this.diamond_)
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
	// serialize magic_
	{
		if this.magic_ != 0 {
			err := prpc.Write(buffer, this.magic_)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *SGE_LogProduceTrack) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 2)
	if err != nil {
		return err
	}
	// deserialize timestamp_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.timestamp_)
		if err != nil {
			return err
		}
	}
	// deserialize from_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.from_)
		if err != nil {
			return err
		}
	}
	// deserialize playerId_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.playerId_)
		if err != nil {
			return err
		}
	}
	// deserialize playerName_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.playerName_)
		if err != nil {
			return err
		}
	}
	// deserialize itemId_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.itemId_)
		if err != nil {
			return err
		}
	}
	// deserialize itemInstId_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.itemInstId_)
		if err != nil {
			return err
		}
	}
	// deserialize itemStack_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.itemStack_)
		if err != nil {
			return err
		}
	}
	// deserialize diamond_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.diamond_)
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
	// deserialize magic_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.magic_)
		if err != nil {
			return err
		}
	}
	return nil
}
