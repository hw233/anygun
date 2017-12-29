package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type COM_SimpleInformation struct {
	instId_       int32  //0
	level_        int32  //1
	asset_id_     int32  //2
	instName_     string //3
	weaponItemId_ int32  //4
	fashionId_    int32  //5
	section_      uint32 //6
	jt_           int    //7
	jl_           int32  //8
}

func (this *COM_SimpleInformation) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(2)
	mask.WriteBit(this.instId_ != 0)
	mask.WriteBit(this.level_ != 0)
	mask.WriteBit(this.asset_id_ != 0)
	mask.WriteBit(len(this.instName_) != 0)
	mask.WriteBit(this.weaponItemId_ != 0)
	mask.WriteBit(this.fashionId_ != 0)
	mask.WriteBit(this.section_ != 0)
	mask.WriteBit(this.jt_ != 0)
	mask.WriteBit(this.jl_ != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize instId_
	{
		if this.instId_ != 0 {
			err := prpc.Write(buffer, this.instId_)
			if err != nil {
				return err
			}
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
	// serialize asset_id_
	{
		if this.asset_id_ != 0 {
			err := prpc.Write(buffer, this.asset_id_)
			if err != nil {
				return err
			}
		}
	}
	// serialize instName_
	if len(this.instName_) != 0 {
		err := prpc.Write(buffer, this.instName_)
		if err != nil {
			return err
		}
	}
	// serialize weaponItemId_
	{
		if this.weaponItemId_ != 0 {
			err := prpc.Write(buffer, this.weaponItemId_)
			if err != nil {
				return err
			}
		}
	}
	// serialize fashionId_
	{
		if this.fashionId_ != 0 {
			err := prpc.Write(buffer, this.fashionId_)
			if err != nil {
				return err
			}
		}
	}
	// serialize section_
	{
		if this.section_ != 0 {
			err := prpc.Write(buffer, this.section_)
			if err != nil {
				return err
			}
		}
	}
	// serialize jt_
	{
		if this.jt_ != 0 {
			err := prpc.Write(buffer, this.jt_)
			if err != nil {
				return err
			}
		}
	}
	// serialize jl_
	{
		if this.jl_ != 0 {
			err := prpc.Write(buffer, this.jl_)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *COM_SimpleInformation) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 2)
	if err != nil {
		return err
	}
	// deserialize instId_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.instId_)
		if err != nil {
			return err
		}
	}
	// deserialize level_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.level_)
		if err != nil {
			return err
		}
	}
	// deserialize asset_id_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.asset_id_)
		if err != nil {
			return err
		}
	}
	// deserialize instName_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.instName_)
		if err != nil {
			return err
		}
	}
	// deserialize weaponItemId_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.weaponItemId_)
		if err != nil {
			return err
		}
	}
	// deserialize fashionId_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.fashionId_)
		if err != nil {
			return err
		}
	}
	// deserialize section_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.section_)
		if err != nil {
			return err
		}
	}
	// deserialize jt_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.jt_)
		if err != nil {
			return err
		}
	}
	// deserialize jl_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.jl_)
		if err != nil {
			return err
		}
	}
	return nil
}
