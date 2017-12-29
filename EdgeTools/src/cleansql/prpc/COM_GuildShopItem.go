package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type COM_GuildShopItem struct {
	shopId_   int32 //0
	buyLimit_ int32 //1
}

func (this *COM_GuildShopItem) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.shopId_ != 0)
	mask.WriteBit(this.buyLimit_ != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize shopId_
	{
		if this.shopId_ != 0 {
			err := prpc.Write(buffer, this.shopId_)
			if err != nil {
				return err
			}
		}
	}
	// serialize buyLimit_
	{
		if this.buyLimit_ != 0 {
			err := prpc.Write(buffer, this.buyLimit_)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *COM_GuildShopItem) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize shopId_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.shopId_)
		if err != nil {
			return err
		}
	}
	// deserialize buyLimit_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.buyLimit_)
		if err != nil {
			return err
		}
	}
	return nil
}
