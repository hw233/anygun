package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type COM_ShowbabyInstInfo struct {
	showId_     int32  //0
	playerName_ string //1
	babyId_     uint32 //2
}

func (this *COM_ShowbabyInstInfo) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.showId_ != 0)
	mask.WriteBit(len(this.playerName_) != 0)
	mask.WriteBit(this.babyId_ != 0)
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
	// serialize babyId_
	{
		if this.babyId_ != 0 {
			err := prpc.Write(buffer, this.babyId_)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *COM_ShowbabyInstInfo) Deserialize(buffer *bytes.Buffer) error {
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
	// deserialize babyId_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.babyId_)
		if err != nil {
			return err
		}
	}
	return nil
}
