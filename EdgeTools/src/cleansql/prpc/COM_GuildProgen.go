package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type COM_GuildProgen struct {
	mId_ int32 //0
	lev_ int32 //1
	exp_ int32 //2
}

func (this *COM_GuildProgen) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.mId_ != 0)
	mask.WriteBit(this.lev_ != 0)
	mask.WriteBit(this.exp_ != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize mId_
	{
		if this.mId_ != 0 {
			err := prpc.Write(buffer, this.mId_)
			if err != nil {
				return err
			}
		}
	}
	// serialize lev_
	{
		if this.lev_ != 0 {
			err := prpc.Write(buffer, this.lev_)
			if err != nil {
				return err
			}
		}
	}
	// serialize exp_
	{
		if this.exp_ != 0 {
			err := prpc.Write(buffer, this.exp_)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *COM_GuildProgen) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize mId_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.mId_)
		if err != nil {
			return err
		}
	}
	// deserialize lev_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.lev_)
		if err != nil {
			return err
		}
	}
	// deserialize exp_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.exp_)
		if err != nil {
			return err
		}
	}
	return nil
}
