package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type COM_ReportState struct {
	COM_State
	add_      bool  //0
	ownerId_  int32 //1
	addQueue_ int8  //2
}

func (this *COM_ReportState) Serialize(buffer *bytes.Buffer) error {
	{
		err := this.COM_State.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.add_)
	mask.WriteBit(this.ownerId_ != 0)
	mask.WriteBit(this.addQueue_ != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize add_
	{
	}
	// serialize ownerId_
	{
		if this.ownerId_ != 0 {
			err := prpc.Write(buffer, this.ownerId_)
			if err != nil {
				return err
			}
		}
	}
	// serialize addQueue_
	{
		if this.addQueue_ != 0 {
			err := prpc.Write(buffer, this.addQueue_)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *COM_ReportState) Deserialize(buffer *bytes.Buffer) error {
	{
		this.COM_State.Deserialize(buffer)
	}
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize add_
	this.add_ = mask.ReadBit()
	// deserialize ownerId_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.ownerId_)
		if err != nil {
			return err
		}
	}
	// deserialize addQueue_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.addQueue_)
		if err != nil {
			return err
		}
	}
	return nil
}
