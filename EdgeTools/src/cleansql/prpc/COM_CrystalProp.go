package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type COM_CrystalProp struct {
	level_ uint32 //0
	type_  int    //1
	val_   uint32 //2
}

func (this *COM_CrystalProp) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.level_ != 0)
	mask.WriteBit(this.type_ != 0)
	mask.WriteBit(this.val_ != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize level_
	{
		if this.level_ != 0 {
			err := prpc.Write(buffer, this.level_)
			if err != nil {
				return err
			}
		}
	}
	// serialize type_
	{
		if this.type_ != 0 {
			err := prpc.Write(buffer, this.type_)
			if err != nil {
				return err
			}
		}
	}
	// serialize val_
	{
		if this.val_ != 0 {
			err := prpc.Write(buffer, this.val_)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *COM_CrystalProp) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize level_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.level_)
		if err != nil {
			return err
		}
	}
	// deserialize type_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.type_)
		if err != nil {
			return err
		}
	}
	// deserialize val_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.val_)
		if err != nil {
			return err
		}
	}
	return nil
}
