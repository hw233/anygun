package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type COM_ZhuanpanData struct {
	sinceStamp_ uint64                //0
	endStamp_   uint64                //1
	contents_   []COM_ZhuanpanContent //2
	rarity_     []COM_Zhuanpan        //3
}

func (this *COM_ZhuanpanData) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.sinceStamp_ != 0)
	mask.WriteBit(this.endStamp_ != 0)
	mask.WriteBit(len(this.contents_) != 0)
	mask.WriteBit(len(this.rarity_) != 0)
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
	// serialize contents_
	if len(this.contents_) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.contents_)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.contents_ {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// serialize rarity_
	if len(this.rarity_) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.rarity_)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.rarity_ {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *COM_ZhuanpanData) Deserialize(buffer *bytes.Buffer) error {
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
	// deserialize contents_
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.contents_ = make([]COM_ZhuanpanContent, size)
		for i, _ := range this.contents_ {
			err := this.contents_[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// deserialize rarity_
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.rarity_ = make([]COM_Zhuanpan, size)
		for i, _ := range this.rarity_ {
			err := this.rarity_[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
