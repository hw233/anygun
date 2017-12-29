package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type COM_ScenePlayerInformation struct {
	isLeader_          bool          //0
	isInBattle_        bool          //1
	vip_               int16         //2
	instId_            int32         //3
	assetId_           int32         //4
	weaponItemId_      int32         //5
	fashionId_         int32         //6
	hpMax_             int32         //7
	hpCrt_             int32         //8
	mpMax_             int32         //9
	mpCrt_             int32         //10
	level_             int32         //11
	battlePower_       int32         //12
	jl_                int32         //13
	openSubSystemFlag_ int32         //14
	title_             int32         //15
	instName_          string        //16
	guildeName_        string        //17
	jt_                int           //18
	type_              int           //19
	originPos_         COM_FPosition //20
	showBabyTableId_   int32         //21
	showBabyName_      string        //22
}

func (this *COM_ScenePlayerInformation) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(3)
	mask.WriteBit(this.isLeader_)
	mask.WriteBit(this.isInBattle_)
	mask.WriteBit(this.vip_ != 0)
	mask.WriteBit(this.instId_ != 0)
	mask.WriteBit(this.assetId_ != 0)
	mask.WriteBit(this.weaponItemId_ != 0)
	mask.WriteBit(this.fashionId_ != 0)
	mask.WriteBit(this.hpMax_ != 0)
	mask.WriteBit(this.hpCrt_ != 0)
	mask.WriteBit(this.mpMax_ != 0)
	mask.WriteBit(this.mpCrt_ != 0)
	mask.WriteBit(this.level_ != 0)
	mask.WriteBit(this.battlePower_ != 0)
	mask.WriteBit(this.jl_ != 0)
	mask.WriteBit(this.openSubSystemFlag_ != 0)
	mask.WriteBit(this.title_ != 0)
	mask.WriteBit(len(this.instName_) != 0)
	mask.WriteBit(len(this.guildeName_) != 0)
	mask.WriteBit(this.jt_ != 0)
	mask.WriteBit(this.type_ != 0)
	mask.WriteBit(true) //originPos_
	mask.WriteBit(this.showBabyTableId_ != 0)
	mask.WriteBit(len(this.showBabyName_) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize isLeader_
	{
	}
	// serialize isInBattle_
	{
	}
	// serialize vip_
	{
		if this.vip_ != 0 {
			err := prpc.Write(buffer, this.vip_)
			if err != nil {
				return err
			}
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
	// serialize assetId_
	{
		if this.assetId_ != 0 {
			err := prpc.Write(buffer, this.assetId_)
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
	// serialize battlePower_
	{
		if this.battlePower_ != 0 {
			err := prpc.Write(buffer, this.battlePower_)
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
	// serialize openSubSystemFlag_
	{
		if this.openSubSystemFlag_ != 0 {
			err := prpc.Write(buffer, this.openSubSystemFlag_)
			if err != nil {
				return err
			}
		}
	}
	// serialize title_
	{
		if this.title_ != 0 {
			err := prpc.Write(buffer, this.title_)
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
	// serialize guildeName_
	if len(this.guildeName_) != 0 {
		err := prpc.Write(buffer, this.guildeName_)
		if err != nil {
			return err
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
	// serialize type_
	{
		if this.type_ != 0 {
			err := prpc.Write(buffer, this.type_)
			if err != nil {
				return err
			}
		}
	}
	// serialize originPos_
	{
		err := this.originPos_.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	// serialize showBabyTableId_
	{
		if this.showBabyTableId_ != 0 {
			err := prpc.Write(buffer, this.showBabyTableId_)
			if err != nil {
				return err
			}
		}
	}
	// serialize showBabyName_
	if len(this.showBabyName_) != 0 {
		err := prpc.Write(buffer, this.showBabyName_)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *COM_ScenePlayerInformation) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 3)
	if err != nil {
		return err
	}
	// deserialize isLeader_
	this.isLeader_ = mask.ReadBit()
	// deserialize isInBattle_
	this.isInBattle_ = mask.ReadBit()
	// deserialize vip_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.vip_)
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
	// deserialize assetId_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.assetId_)
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
	// deserialize battlePower_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.battlePower_)
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
	// deserialize openSubSystemFlag_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.openSubSystemFlag_)
		if err != nil {
			return err
		}
	}
	// deserialize title_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.title_)
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
	// deserialize guildeName_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.guildeName_)
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
	// deserialize type_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.type_)
		if err != nil {
			return err
		}
	}
	// deserialize originPos_
	if mask.ReadBit() {
		err := this.originPos_.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	// deserialize showBabyTableId_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.showBabyTableId_)
		if err != nil {
			return err
		}
	}
	// deserialize showBabyName_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.showBabyName_)
		if err != nil {
			return err
		}
	}
	return nil
}
