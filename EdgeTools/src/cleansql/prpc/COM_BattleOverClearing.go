package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type COM_BattleOverClearing struct {
	isFly_       bool           //0
	playExp_     int32          //1
	playLevel_   uint32         //2
	playFree_    uint32         //3
	pvpJJCGrade_ int32          //4
	money_       int32          //5
	babyExp_     uint32         //6
	babyLevel_   uint32         //7
	items_       []COM_DropItem //8
	skills_      []COM_Skill    //9
}

func (this *COM_BattleOverClearing) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(2)
	mask.WriteBit(this.isFly_)
	mask.WriteBit(this.playExp_ != 0)
	mask.WriteBit(this.playLevel_ != 0)
	mask.WriteBit(this.playFree_ != 0)
	mask.WriteBit(this.pvpJJCGrade_ != 0)
	mask.WriteBit(this.money_ != 0)
	mask.WriteBit(this.babyExp_ != 0)
	mask.WriteBit(this.babyLevel_ != 0)
	mask.WriteBit(len(this.items_) != 0)
	mask.WriteBit(len(this.skills_) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize isFly_
	{
	}
	// serialize playExp_
	{
		if this.playExp_ != 0 {
			err := prpc.Write(buffer, this.playExp_)
			if err != nil {
				return err
			}
		}
	}
	// serialize playLevel_
	{
		if this.playLevel_ != 0 {
			err := prpc.Write(buffer, this.playLevel_)
			if err != nil {
				return err
			}
		}
	}
	// serialize playFree_
	{
		if this.playFree_ != 0 {
			err := prpc.Write(buffer, this.playFree_)
			if err != nil {
				return err
			}
		}
	}
	// serialize pvpJJCGrade_
	{
		if this.pvpJJCGrade_ != 0 {
			err := prpc.Write(buffer, this.pvpJJCGrade_)
			if err != nil {
				return err
			}
		}
	}
	// serialize money_
	{
		if this.money_ != 0 {
			err := prpc.Write(buffer, this.money_)
			if err != nil {
				return err
			}
		}
	}
	// serialize babyExp_
	{
		if this.babyExp_ != 0 {
			err := prpc.Write(buffer, this.babyExp_)
			if err != nil {
				return err
			}
		}
	}
	// serialize babyLevel_
	{
		if this.babyLevel_ != 0 {
			err := prpc.Write(buffer, this.babyLevel_)
			if err != nil {
				return err
			}
		}
	}
	// serialize items_
	if len(this.items_) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.items_)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.items_ {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// serialize skills_
	if len(this.skills_) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.skills_)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.skills_ {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *COM_BattleOverClearing) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 2)
	if err != nil {
		return err
	}
	// deserialize isFly_
	this.isFly_ = mask.ReadBit()
	// deserialize playExp_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.playExp_)
		if err != nil {
			return err
		}
	}
	// deserialize playLevel_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.playLevel_)
		if err != nil {
			return err
		}
	}
	// deserialize playFree_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.playFree_)
		if err != nil {
			return err
		}
	}
	// deserialize pvpJJCGrade_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.pvpJJCGrade_)
		if err != nil {
			return err
		}
	}
	// deserialize money_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.money_)
		if err != nil {
			return err
		}
	}
	// deserialize babyExp_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.babyExp_)
		if err != nil {
			return err
		}
	}
	// deserialize babyLevel_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.babyLevel_)
		if err != nil {
			return err
		}
	}
	// deserialize items_
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.items_ = make([]COM_DropItem, size)
		for i, _ := range this.items_ {
			err := this.items_[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// deserialize skills_
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.skills_ = make([]COM_Skill, size)
		for i, _ := range this.skills_ {
			err := this.skills_[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
