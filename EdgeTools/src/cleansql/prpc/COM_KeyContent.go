package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type COM_KeyContent struct {
	giftname_   string         //0
	pfid_       string         //1
	key_        string         //2
	playerName_ string         //3
	usetime_    int64          //4
	items_      []COM_GiftItem //5
}

func (this *COM_KeyContent) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.giftname_) != 0)
	mask.WriteBit(len(this.pfid_) != 0)
	mask.WriteBit(len(this.key_) != 0)
	mask.WriteBit(len(this.playerName_) != 0)
	mask.WriteBit(this.usetime_ != 0)
	mask.WriteBit(len(this.items_) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize giftname_
	if len(this.giftname_) != 0 {
		err := prpc.Write(buffer, this.giftname_)
		if err != nil {
			return err
		}
	}
	// serialize pfid_
	if len(this.pfid_) != 0 {
		err := prpc.Write(buffer, this.pfid_)
		if err != nil {
			return err
		}
	}
	// serialize key_
	if len(this.key_) != 0 {
		err := prpc.Write(buffer, this.key_)
		if err != nil {
			return err
		}
	}
	// serialize playerName_
	if len(this.playerName_) != 0 {
		err := prpc.Write(buffer, this.playerName_)
		if err != nil {
			return err
		}
	}
	// serialize usetime_
	{
		if this.usetime_ != 0 {
			err := prpc.Write(buffer, this.usetime_)
			if err != nil {
				return err
			}
		}
	}
	// serialize items_
	if len(this.items_) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.items_)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.items_ {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *COM_KeyContent) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize giftname_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.giftname_)
		if err != nil {
			return err
		}
	}
	// deserialize pfid_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.pfid_)
		if err != nil {
			return err
		}
	}
	// deserialize key_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.key_)
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
	// deserialize usetime_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.usetime_)
		if err != nil {
			return err
		}
	}
	// deserialize items_
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.items_ = make([]COM_GiftItem, size)
		for i, _ := range this.items_ {
			err := this.items_[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
