package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type COM_BattleEntityInformation struct {
	type_           int    //0
	instName_       string //1
	instId_         int32  //2
	tableId_        int32  //3
	assetId_        int32  //4
	jt_             int    //5
	battlePosition_ int    //6
	weaponItemId_   int32  //7
	fashionId_      int32  //8
	hpMax_          int32  //9
	hpCrt_          int32  //10
	mpMax_          int32  //11
	mpCrt_          int32  //12
	level_          int32  //13
}

func (this *COM_BattleEntityInformation) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(2)
	mask.WriteBit(this.type_ != 0)
	mask.WriteBit(len(this.instName_) != 0)
	mask.WriteBit(this.instId_ != 0)
	mask.WriteBit(this.tableId_ != 0)
	mask.WriteBit(this.assetId_ != 0)
	mask.WriteBit(this.jt_ != 0)
	mask.WriteBit(this.battlePosition_ != 0)
	mask.WriteBit(this.weaponItemId_ != 0)
	mask.WriteBit(this.fashionId_ != 0)
	mask.WriteBit(this.hpMax_ != 0)
	mask.WriteBit(this.hpCrt_ != 0)
	mask.WriteBit(this.mpMax_ != 0)
	mask.WriteBit(this.mpCrt_ != 0)
	mask.WriteBit(this.level_ != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
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
	// serialize instName_
	if len(this.instName_) != 0 {
		err := prpc.Write(buffer, this.instName_)
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
	// serialize tableId_
	{
		if this.tableId_ != 0 {
			err := prpc.Write(buffer, this.tableId_)
			if err != nil {
				return err
			}
		}
	}
	// serialize assetId_
	{
		if this.assetId_ != 0 {
			err := prpc.Write(buffer, this.assetId_)
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
	// serialize battlePosition_
	{
		if this.battlePosition_ != 0 {
			err := prpc.Write(buffer, this.battlePosition_)
			if err != nil {
				return err
			}
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
	// serialize hpMax_
	{
		if this.hpMax_ != 0 {
			err := prpc.Write(buffer, this.hpMax_)
			if err != nil {
				return err
			}
		}
	}
	// serialize hpCrt_
	{
		if this.hpCrt_ != 0 {
			err := prpc.Write(buffer, this.hpCrt_)
			if err != nil {
				return err
			}
		}
	}
	// serialize mpMax_
	{
		if this.mpMax_ != 0 {
			err := prpc.Write(buffer, this.mpMax_)
			if err != nil {
				return err
			}
		}
	}
	// serialize mpCrt_
	{
		if this.mpCrt_ != 0 {
			err := prpc.Write(buffer, this.mpCrt_)
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
	return nil
}
func (this *COM_BattleEntityInformation) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 2)
	if err != nil {
		return err
	}
	// deserialize type_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.type_)
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
	// deserialize instId_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.instId_)
		if err != nil {
			return err
		}
	}
	// deserialize tableId_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.tableId_)
		if err != nil {
			return err
		}
	}
	// deserialize assetId_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.assetId_)
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
	// deserialize battlePosition_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.battlePosition_)
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
	// deserialize hpMax_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.hpMax_)
		if err != nil {
			return err
		}
	}
	// deserialize hpCrt_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.hpCrt_)
		if err != nil {
			return err
		}
	}
	// deserialize mpMax_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.mpMax_)
		if err != nil {
			return err
		}
	}
	// deserialize mpCrt_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.mpCrt_)
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
	return nil
}
