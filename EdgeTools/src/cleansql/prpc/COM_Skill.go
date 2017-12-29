package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type COM_Skill struct {
	skillID_    uint32 //0
	skillExp_   uint32 //1
	skillLevel_ uint32 //2
}

func (this *COM_Skill) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.skillID_ != 0)
	mask.WriteBit(this.skillExp_ != 0)
	mask.WriteBit(this.skillLevel_ != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize skillID_
	{
		if this.skillID_ != 0 {
			err := prpc.Write(buffer, this.skillID_)
			if err != nil {
				return err
			}
		}
	}
	// serialize skillExp_
	{
		if this.skillExp_ != 0 {
			err := prpc.Write(buffer, this.skillExp_)
			if err != nil {
				return err
			}
		}
	}
	// serialize skillLevel_
	{
		if this.skillLevel_ != 0 {
			err := prpc.Write(buffer, this.skillLevel_)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *COM_Skill) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize skillID_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.skillID_)
		if err != nil {
			return err
		}
	}
	// deserialize skillExp_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.skillExp_)
		if err != nil {
			return err
		}
	}
	// deserialize skillLevel_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.skillLevel_)
		if err != nil {
			return err
		}
	}
	return nil
}
