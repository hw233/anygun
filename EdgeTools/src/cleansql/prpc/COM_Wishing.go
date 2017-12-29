package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type COM_Wishing struct {
	wt_   int    //0
	wish_ string //1
}

func (this *COM_Wishing) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.wt_ != 0)
	mask.WriteBit(len(this.wish_) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize wt_
	{
		if this.wt_ != 0 {
			err := prpc.Write(buffer, this.wt_)
			if err != nil {
				return err
			}
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
func (this *COM_Wishing) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize wt_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.wt_)
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
