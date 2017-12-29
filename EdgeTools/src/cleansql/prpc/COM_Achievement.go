package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type COM_Achievement struct {
	achId_    uint32 //0
	achType_  int    //1
	achValue_ uint32 //2
	isAch_    bool   //3
	isAward_  bool   //4
}

func (this *COM_Achievement) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.achId_ != 0)
	mask.WriteBit(this.achType_ != 0)
	mask.WriteBit(this.achValue_ != 0)
	mask.WriteBit(this.isAch_)
	mask.WriteBit(this.isAward_)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize achId_
	{
		if this.achId_ != 0 {
			err := prpc.Write(buffer, this.achId_)
			if err != nil {
				return err
			}
		}
	}
	// serialize achType_
	{
		if this.achType_ != 0 {
			err := prpc.Write(buffer, this.achType_)
			if err != nil {
				return err
			}
		}
	}
	// serialize achValue_
	{
		if this.achValue_ != 0 {
			err := prpc.Write(buffer, this.achValue_)
			if err != nil {
				return err
			}
		}
	}
	// serialize isAch_
	{
	}
	// serialize isAward_
	{
	}
	return nil
}
func (this *COM_Achievement) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize achId_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.achId_)
		if err != nil {
			return err
		}
	}
	// deserialize achType_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.achType_)
		if err != nil {
			return err
		}
	}
	// deserialize achValue_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.achValue_)
		if err != nil {
			return err
		}
	}
	// deserialize isAch_
	this.isAch_ = mask.ReadBit()
	// deserialize isAward_
	this.isAward_ = mask.ReadBit()
	return nil
}
