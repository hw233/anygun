package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type COM_State struct {
	stateId_ uint32 //0
	turn_    int8   //1
	tick_    int8   //2
	type_    int8   //3
	value0_  int32  //4
	value1_  int32  //5
}

func (this *COM_State) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.stateId_ != 0)
	mask.WriteBit(this.turn_ != 0)
	mask.WriteBit(this.tick_ != 0)
	mask.WriteBit(this.type_ != 0)
	mask.WriteBit(this.value0_ != 0)
	mask.WriteBit(this.value1_ != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize stateId_
	{
		if this.stateId_ != 0 {
			err := prpc.Write(buffer, this.stateId_)
			if err != nil {
				return err
			}
		}
	}
	// serialize turn_
	{
		if this.turn_ != 0 {
			err := prpc.Write(buffer, this.turn_)
			if err != nil {
				return err
			}
		}
	}
	// serialize tick_
	{
		if this.tick_ != 0 {
			err := prpc.Write(buffer, this.tick_)
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
	// serialize value0_
	{
		if this.value0_ != 0 {
			err := prpc.Write(buffer, this.value0_)
			if err != nil {
				return err
			}
		}
	}
	// serialize value1_
	{
		if this.value1_ != 0 {
			err := prpc.Write(buffer, this.value1_)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *COM_State) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize stateId_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.stateId_)
		if err != nil {
			return err
		}
	}
	// deserialize turn_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.turn_)
		if err != nil {
			return err
		}
	}
	// deserialize tick_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.tick_)
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
	// deserialize value0_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.value0_)
		if err != nil {
			return err
		}
	}
	// deserialize value1_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.value1_)
		if err != nil {
			return err
		}
	}
	return nil
}
