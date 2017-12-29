package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type COM_ADGiftBag struct {
	sinceStamp_ uint64         //0
	endStamp_   uint64         //1
	isflag_     bool           //2
	price_      uint8          //3
	oldprice_   uint8          //4
	itemdata_   []COM_GiftItem //5
}

func (this *COM_ADGiftBag) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.sinceStamp_ != 0)
	mask.WriteBit(this.endStamp_ != 0)
	mask.WriteBit(this.isflag_)
	mask.WriteBit(this.price_ != 0)
	mask.WriteBit(this.oldprice_ != 0)
	mask.WriteBit(len(this.itemdata_) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize sinceStamp_
	{
		if this.sinceStamp_ != 0 {
			err := prpc.Write(buffer, this.sinceStamp_)
			if err != nil {
				return err
			}
		}
	}
	// serialize endStamp_
	{
		if this.endStamp_ != 0 {
			err := prpc.Write(buffer, this.endStamp_)
			if err != nil {
				return err
			}
		}
	}
	// serialize isflag_
	{
	}
	// serialize price_
	{
		if this.price_ != 0 {
			err := prpc.Write(buffer, this.price_)
			if err != nil {
				return err
			}
		}
	}
	// serialize oldprice_
	{
		if this.oldprice_ != 0 {
			err := prpc.Write(buffer, this.oldprice_)
			if err != nil {
				return err
			}
		}
	}
	// serialize itemdata_
	if len(this.itemdata_) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.itemdata_)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.itemdata_ {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *COM_ADGiftBag) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize sinceStamp_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.sinceStamp_)
		if err != nil {
			return err
		}
	}
	// deserialize endStamp_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.endStamp_)
		if err != nil {
			return err
		}
	}
	// deserialize isflag_
	this.isflag_ = mask.ReadBit()
	// deserialize price_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.price_)
		if err != nil {
			return err
		}
	}
	// deserialize oldprice_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.oldprice_)
		if err != nil {
			return err
		}
	}
	// deserialize itemdata_
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.itemdata_ = make([]COM_GiftItem, size)
		for i, _ := range this.itemdata_ {
			err := this.itemdata_[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
