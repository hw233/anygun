package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type COM_Notice struct {
	SendType_   int     //0
	note_       string  //1
	accumulate_ float32 //2
	startTime_  float32 //3
	stopTime_   float32 //4
	interval_   float32 //5
	validate_   bool    //6
}

func (this *COM_Notice) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.SendType_ != 0)
	mask.WriteBit(len(this.note_) != 0)
	mask.WriteBit(this.accumulate_ != 0)
	mask.WriteBit(this.startTime_ != 0)
	mask.WriteBit(this.stopTime_ != 0)
	mask.WriteBit(this.interval_ != 0)
	mask.WriteBit(this.validate_)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize SendType_
	{
		if this.SendType_ != 0 {
			err := prpc.Write(buffer, this.SendType_)
			if err != nil {
				return err
			}
		}
	}
	// serialize note_
	if len(this.note_) != 0 {
		err := prpc.Write(buffer, this.note_)
		if err != nil {
			return err
		}
	}
	// serialize accumulate_
	{
		if this.accumulate_ != 0 {
			err := prpc.Write(buffer, this.accumulate_)
			if err != nil {
				return err
			}
		}
	}
	// serialize startTime_
	{
		if this.startTime_ != 0 {
			err := prpc.Write(buffer, this.startTime_)
			if err != nil {
				return err
			}
		}
	}
	// serialize stopTime_
	{
		if this.stopTime_ != 0 {
			err := prpc.Write(buffer, this.stopTime_)
			if err != nil {
				return err
			}
		}
	}
	// serialize interval_
	{
		if this.interval_ != 0 {
			err := prpc.Write(buffer, this.interval_)
			if err != nil {
				return err
			}
		}
	}
	// serialize validate_
	{
	}
	return nil
}
func (this *COM_Notice) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize SendType_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.SendType_)
		if err != nil {
			return err
		}
	}
	// deserialize note_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.note_)
		if err != nil {
			return err
		}
	}
	// deserialize accumulate_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.accumulate_)
		if err != nil {
			return err
		}
	}
	// deserialize startTime_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.startTime_)
		if err != nil {
			return err
		}
	}
	// deserialize stopTime_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.stopTime_)
		if err != nil {
			return err
		}
	}
	// deserialize interval_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.interval_)
		if err != nil {
			return err
		}
	}
	// deserialize validate_
	this.validate_ = mask.ReadBit()
	return nil
}
