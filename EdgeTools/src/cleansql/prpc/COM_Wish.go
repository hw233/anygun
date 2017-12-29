package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type COM_Wish struct {
	playerInstId_ uint32 //0
	playerName_   string //1
	wish_         string //2
}

func (this *COM_Wish) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.playerInstId_ != 0)
	mask.WriteBit(len(this.playerName_) != 0)
	mask.WriteBit(len(this.wish_) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize playerInstId_
	{
		if this.playerInstId_ != 0 {
			err := prpc.Write(buffer, this.playerInstId_)
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
	// serialize wish_
	if len(this.wish_) != 0 {
		err := prpc.Write(buffer, this.wish_)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *COM_Wish) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize playerInstId_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.playerInstId_)
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
	// deserialize wish_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.wish_)
		if err != nil {
			return err
		}
	}
	return nil
}
