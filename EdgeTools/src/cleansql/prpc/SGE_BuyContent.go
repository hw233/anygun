package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type SGE_BuyContent struct {
	playerId_   int32 //0
	guid_       int32 //1
	diamond_    int32 //2
	magic_      int32 //3
	isBabyFull_ bool  //4
	isBagFull_  bool  //5
}

func (this *SGE_BuyContent) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.playerId_ != 0)
	mask.WriteBit(this.guid_ != 0)
	mask.WriteBit(this.diamond_ != 0)
	mask.WriteBit(this.magic_ != 0)
	mask.WriteBit(this.isBabyFull_)
	mask.WriteBit(this.isBagFull_)
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
	// serialize guid_
	{
		if this.guid_ != 0 {
			err := prpc.Write(buffer, this.guid_)
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
	// serialize magic_
	{
		if this.magic_ != 0 {
			err := prpc.Write(buffer, this.magic_)
			if err != nil {
				return err
			}
		}
	}
	// serialize isBabyFull_
	{
	}
	// serialize isBagFull_
	{
	}
	return nil
}
func (this *SGE_BuyContent) Deserialize(buffer *bytes.Buffer) error {
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
	// deserialize guid_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.guid_)
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
	// deserialize magic_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.magic_)
		if err != nil {
			return err
		}
	}
	// deserialize isBabyFull_
	this.isBabyFull_ = mask.ReadBit()
	// deserialize isBagFull_
	this.isBagFull_ = mask.ReadBit()
	return nil
}
