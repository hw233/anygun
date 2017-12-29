package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type COM_Entity struct {
	type_           int         //0
	instName_       string      //1
	instId_         uint32      //2
	battlePosition_ int         //3
	properties_     []float32   //4
	skill_          []COM_Skill //5
	equips_         []COM_Item  //6
	states_         []COM_State //7
}

func (this *COM_Entity) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.type_ != 0)
	mask.WriteBit(len(this.instName_) != 0)
	mask.WriteBit(this.instId_ != 0)
	mask.WriteBit(this.battlePosition_ != 0)
	mask.WriteBit(len(this.properties_) != 0)
	mask.WriteBit(len(this.skill_) != 0)
	mask.WriteBit(len(this.equips_) != 0)
	mask.WriteBit(len(this.states_) != 0)
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
	// serialize battlePosition_
	{
		if this.battlePosition_ != 0 {
			err := prpc.Write(buffer, this.battlePosition_)
			if err != nil {
				return err
			}
		}
	}
	// serialize properties_
	if len(this.properties_) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.properties_)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.properties_ {
			err := prpc.Write(buffer, value)
			if err != nil {
				return err
			}
		}
	}
	// serialize skill_
	if len(this.skill_) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.skill_)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.skill_ {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// serialize equips_
	if len(this.equips_) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.equips_)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.equips_ {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// serialize states_
	if len(this.states_) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.states_)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.states_ {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *COM_Entity) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
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
	// deserialize battlePosition_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.battlePosition_)
		if err != nil {
			return err
		}
	}
	// deserialize properties_
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.properties_ = make([]float32, size)
		for i, _ := range this.properties_ {
			err := prpc.Read(buffer, &this.properties_[i])
			if err != nil {
				return err
			}
		}
	}
	// deserialize skill_
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.skill_ = make([]COM_Skill, size)
		for i, _ := range this.skill_ {
			err := this.skill_[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// deserialize equips_
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.equips_ = make([]COM_Item, size)
		for i, _ := range this.equips_ {
			err := this.equips_[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// deserialize states_
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.states_ = make([]COM_State, size)
		for i, _ := range this.states_ {
			err := this.states_[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
