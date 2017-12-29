package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type COM_Zhuanpan struct {
	playerName_ string //0
	zhuanpanId_ uint32 //1
}

func (this *COM_Zhuanpan) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.playerName_) != 0)
	mask.WriteBit(this.zhuanpanId_ != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
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
	// serialize zhuanpanId_
	{
		if this.zhuanpanId_ != 0 {
			err := prpc.Write(buffer, this.zhuanpanId_)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *COM_Zhuanpan) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize playerName_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.playerName_)
		if err != nil {
			return err
		}
	}
	// deserialize zhuanpanId_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.zhuanpanId_)
		if err != nil {
			return err
		}
	}
	return nil
}
