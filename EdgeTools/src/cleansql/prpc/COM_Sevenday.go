package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type COM_Sevenday struct {
	quest_    uint32 //0
	stype_    int    //1
	qvalue_   uint32 //2
	isfinish_ bool   //3
	isreward_ bool   //4
}

func (this *COM_Sevenday) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.quest_ != 0)
	mask.WriteBit(this.stype_ != 0)
	mask.WriteBit(this.qvalue_ != 0)
	mask.WriteBit(this.isfinish_)
	mask.WriteBit(this.isreward_)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize quest_
	{
		if this.quest_ != 0 {
			err := prpc.Write(buffer, this.quest_)
			if err != nil {
				return err
			}
		}
	}
	// serialize stype_
	{
		if this.stype_ != 0 {
			err := prpc.Write(buffer, this.stype_)
			if err != nil {
				return err
			}
		}
	}
	// serialize qvalue_
	{
		if this.qvalue_ != 0 {
			err := prpc.Write(buffer, this.qvalue_)
			if err != nil {
				return err
			}
		}
	}
	// serialize isfinish_
	{
	}
	// serialize isreward_
	{
	}
	return nil
}
func (this *COM_Sevenday) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize quest_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.quest_)
		if err != nil {
			return err
		}
	}
	// deserialize stype_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.stype_)
		if err != nil {
			return err
		}
	}
	// deserialize qvalue_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.qvalue_)
		if err != nil {
			return err
		}
	}
	// deserialize isfinish_
	this.isfinish_ = mask.ReadBit()
	// deserialize isreward_
	this.isreward_ = mask.ReadBit()
	return nil
}
