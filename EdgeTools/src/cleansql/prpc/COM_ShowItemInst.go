package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type COM_ShowItemInst struct {
	showId_     int32    //0
	playerName_ string   //1
	itemInst_   COM_Item //2
}

func (this *COM_ShowItemInst) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.showId_ != 0)
	mask.WriteBit(len(this.playerName_) != 0)
	mask.WriteBit(true) //itemInst_
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
	// serialize itemInst_
	{
		err := this.itemInst_.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *COM_ShowItemInst) Deserialize(buffer *bytes.Buffer) error {
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
	// deserialize itemInst_
	if mask.ReadBit() {
		err := this.itemInst_.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
