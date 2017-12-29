package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type COM_SellItem struct {
	guid_         int32        //0
	sellPlayerId_ int32        //1
	sellPrice     int32        //2
	title_        string       //3
	ist_          int          //4
	rt_           int          //5
	item_         COM_Item     //6
	baby_         COM_BabyInst //7
	time_         int32        //8
}

func (this *COM_SellItem) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(2)
	mask.WriteBit(this.guid_ != 0)
	mask.WriteBit(this.sellPlayerId_ != 0)
	mask.WriteBit(this.sellPrice != 0)
	mask.WriteBit(len(this.title_) != 0)
	mask.WriteBit(this.ist_ != 0)
	mask.WriteBit(this.rt_ != 0)
	mask.WriteBit(true) //item_
	mask.WriteBit(true) //baby_
	mask.WriteBit(this.time_ != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
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
	// serialize sellPlayerId_
	{
		if this.sellPlayerId_ != 0 {
			err := prpc.Write(buffer, this.sellPlayerId_)
			if err != nil {
				return err
			}
		}
	}
	// serialize sellPrice
	{
		if this.sellPrice != 0 {
			err := prpc.Write(buffer, this.sellPrice)
			if err != nil {
				return err
			}
		}
	}
	// serialize title_
	if len(this.title_) != 0 {
		err := prpc.Write(buffer, this.title_)
		if err != nil {
			return err
		}
	}
	// serialize ist_
	{
		if this.ist_ != 0 {
			err := prpc.Write(buffer, this.ist_)
			if err != nil {
				return err
			}
		}
	}
	// serialize rt_
	{
		if this.rt_ != 0 {
			err := prpc.Write(buffer, this.rt_)
			if err != nil {
				return err
			}
		}
	}
	// serialize item_
	{
		err := this.item_.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	// serialize baby_
	{
		err := this.baby_.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	// serialize time_
	{
		if this.time_ != 0 {
			err := prpc.Write(buffer, this.time_)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *COM_SellItem) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 2)
	if err != nil {
		return err
	}
	// deserialize guid_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.guid_)
		if err != nil {
			return err
		}
	}
	// deserialize sellPlayerId_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.sellPlayerId_)
		if err != nil {
			return err
		}
	}
	// deserialize sellPrice
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.sellPrice)
		if err != nil {
			return err
		}
	}
	// deserialize title_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.title_)
		if err != nil {
			return err
		}
	}
	// deserialize ist_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.ist_)
		if err != nil {
			return err
		}
	}
	// deserialize rt_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.rt_)
		if err != nil {
			return err
		}
	}
	// deserialize item_
	if mask.ReadBit() {
		err := this.item_.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	// deserialize baby_
	if mask.ReadBit() {
		err := this.baby_.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	// deserialize time_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.time_)
		if err != nil {
			return err
		}
	}
	return nil
}
