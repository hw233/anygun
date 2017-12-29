package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type COM_ShowItemInstInfo struct {
	showId_     int32  //0
	playerName_ string //1
	itemId_     uint32 //2
}

func (this *COM_ShowItemInstInfo) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.showId_ != 0)
	mask.WriteBit(len(this.playerName_) != 0)
	mask.WriteBit(this.itemId_ != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize showId_
	{
		if this.showId_ != 0 {
			err := prpc.Write(buffer, this.showId_)
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
	return nil
}
func (this *COM_ShowItemInstInfo) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize showId_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.showId_)
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
	return nil
}
