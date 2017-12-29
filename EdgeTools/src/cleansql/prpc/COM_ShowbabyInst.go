package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type COM_ShowbabyInst struct {
	showId_     int32        //0
	playerName_ string       //1
	babyInst_   COM_BabyInst //2
}

func (this *COM_ShowbabyInst) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.showId_ != 0)
	mask.WriteBit(len(this.playerName_) != 0)
	mask.WriteBit(true) //babyInst_
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
	// serialize babyInst_
	{
		err := this.babyInst_.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *COM_ShowbabyInst) Deserialize(buffer *bytes.Buffer) error {
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
	// deserialize babyInst_
	if mask.ReadBit() {
		err := this.babyInst_.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
