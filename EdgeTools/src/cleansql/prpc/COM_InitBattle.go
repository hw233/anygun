package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type COM_InitBattle struct {
	battleId_    uint32                        //0
	bt_          int                           //1
	roundCount_  uint32                        //2
	opType_      int                           //3
	sneakAttack_ int                           //4
	actors_      []COM_BattleEntityInformation //5
}

func (this *COM_InitBattle) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.battleId_ != 0)
	mask.WriteBit(this.bt_ != 0)
	mask.WriteBit(this.roundCount_ != 0)
	mask.WriteBit(this.opType_ != 0)
	mask.WriteBit(this.sneakAttack_ != 0)
	mask.WriteBit(len(this.actors_) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize battleId_
	{
		if this.battleId_ != 0 {
			err := prpc.Write(buffer, this.battleId_)
			if err != nil {
				return err
			}
		}
	}
	// serialize bt_
	{
		if this.bt_ != 0 {
			err := prpc.Write(buffer, this.bt_)
			if err != nil {
				return err
			}
		}
	}
	// serialize roundCount_
	{
		if this.roundCount_ != 0 {
			err := prpc.Write(buffer, this.roundCount_)
			if err != nil {
				return err
			}
		}
	}
	// serialize opType_
	{
		if this.opType_ != 0 {
			err := prpc.Write(buffer, this.opType_)
			if err != nil {
				return err
			}
		}
	}
	// serialize sneakAttack_
	{
		if this.sneakAttack_ != 0 {
			err := prpc.Write(buffer, this.sneakAttack_)
			if err != nil {
				return err
			}
		}
	}
	// serialize actors_
	if len(this.actors_) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.actors_)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.actors_ {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *COM_InitBattle) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize battleId_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.battleId_)
		if err != nil {
			return err
		}
	}
	// deserialize bt_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.bt_)
		if err != nil {
			return err
		}
	}
	// deserialize roundCount_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.roundCount_)
		if err != nil {
			return err
		}
	}
	// deserialize opType_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.opType_)
		if err != nil {
			return err
		}
	}
	// deserialize sneakAttack_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.sneakAttack_)
		if err != nil {
			return err
		}
	}
	// deserialize actors_
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.actors_ = make([]COM_BattleEntityInformation, size)
		for i, _ := range this.actors_ {
			err := this.actors_[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
