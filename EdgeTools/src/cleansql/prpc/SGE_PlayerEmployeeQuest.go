package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type SGE_PlayerEmployeeQuest struct {
	playerId_      int32                          //0
	usedEmployees_ []int32                        //1
	quests_        []SGE_PlayerEmployeeQuestArray //2
}

func (this *SGE_PlayerEmployeeQuest) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.playerId_ != 0)
	mask.WriteBit(len(this.usedEmployees_) != 0)
	mask.WriteBit(len(this.quests_) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize playerId_
	{
		if this.playerId_ != 0 {
			err := prpc.Write(buffer, this.playerId_)
			if err != nil {
				return err
			}
		}
	}
	// serialize usedEmployees_
	if len(this.usedEmployees_) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.usedEmployees_)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.usedEmployees_ {
			err := prpc.Write(buffer, value)
			if err != nil {
				return err
			}
		}
	}
	// serialize quests_
	if len(this.quests_) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.quests_)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.quests_ {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *SGE_PlayerEmployeeQuest) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize playerId_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.playerId_)
		if err != nil {
			return err
		}
	}
	// deserialize usedEmployees_
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.usedEmployees_ = make([]int32, size)
		for i, _ := range this.usedEmployees_ {
			err := prpc.Read(buffer, &this.usedEmployees_[i])
			if err != nil {
				return err
			}
		}
	}
	// deserialize quests_
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.quests_ = make([]SGE_PlayerEmployeeQuestArray, size)
		for i, _ := range this.quests_ {
			err := this.quests_[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
