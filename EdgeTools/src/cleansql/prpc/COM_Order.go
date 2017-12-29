package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type COM_Order struct {
	status_       int   //0
	casterId_     int32 //1
	target_       int32 //2
	skill_        int32 //3
	itemId_       int32 //4
	weaponInstId_ int32 //5
	babyId_       int32 //6
	isSec_        int8  //7
	uinteGroup_   int8  //8
	uniteNum_     int32 //9
	isNo_         bool  //10
}

func (this *COM_Order) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(2)
	mask.WriteBit(this.status_ != 0)
	mask.WriteBit(this.casterId_ != 0)
	mask.WriteBit(this.target_ != 0)
	mask.WriteBit(this.skill_ != 0)
	mask.WriteBit(this.itemId_ != 0)
	mask.WriteBit(this.weaponInstId_ != 0)
	mask.WriteBit(this.babyId_ != 0)
	mask.WriteBit(this.isSec_ != 0)
	mask.WriteBit(this.uinteGroup_ != 0)
	mask.WriteBit(this.uniteNum_ != 0)
	mask.WriteBit(this.isNo_)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize status_
	{
		if this.status_ != 0 {
			err := prpc.Write(buffer, this.status_)
			if err != nil {
				return err
			}
		}
	}
	// serialize casterId_
	{
		if this.casterId_ != 0 {
			err := prpc.Write(buffer, this.casterId_)
			if err != nil {
				return err
			}
		}
	}
	// serialize target_
	{
		if this.target_ != 0 {
			err := prpc.Write(buffer, this.target_)
			if err != nil {
				return err
			}
		}
	}
	// serialize skill_
	{
		if this.skill_ != 0 {
			err := prpc.Write(buffer, this.skill_)
			if err != nil {
				return err
			}
		}
	}
	// serialize itemId_
	{
		if this.itemId_ != 0 {
			err := prpc.Write(buffer, this.itemId_)
			if err != nil {
				return err
			}
		}
	}
	// serialize weaponInstId_
	{
		if this.weaponInstId_ != 0 {
			err := prpc.Write(buffer, this.weaponInstId_)
			if err != nil {
				return err
			}
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
	// serialize isSec_
	{
		if this.isSec_ != 0 {
			err := prpc.Write(buffer, this.isSec_)
			if err != nil {
				return err
			}
		}
	}
	// serialize uinteGroup_
	{
		if this.uinteGroup_ != 0 {
			err := prpc.Write(buffer, this.uinteGroup_)
			if err != nil {
				return err
			}
		}
	}
	// serialize uniteNum_
	{
		if this.uniteNum_ != 0 {
			err := prpc.Write(buffer, this.uniteNum_)
			if err != nil {
				return err
			}
		}
	}
	// serialize isNo_
	{
	}
	return nil
}
func (this *COM_Order) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 2)
	if err != nil {
		return err
	}
	// deserialize status_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.status_)
		if err != nil {
			return err
		}
	}
	// deserialize casterId_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.casterId_)
		if err != nil {
			return err
		}
	}
	// deserialize target_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.target_)
		if err != nil {
			return err
		}
	}
	// deserialize skill_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.skill_)
		if err != nil {
			return err
		}
	}
	// deserialize itemId_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.itemId_)
		if err != nil {
			return err
		}
	}
	// deserialize weaponInstId_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.weaponInstId_)
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
	// deserialize isSec_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.isSec_)
		if err != nil {
			return err
		}
	}
	// deserialize uinteGroup_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.uinteGroup_)
		if err != nil {
			return err
		}
	}
	// deserialize uniteNum_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.uniteNum_)
		if err != nil {
			return err
		}
	}
	// deserialize isNo_
	this.isNo_ = mask.ReadBit()
	return nil
}
