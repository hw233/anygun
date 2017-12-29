package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type COM_ReportActionCounter struct {
	casterId_       int32             //0
	targetPosition_ uint8             //1
	props_          COM_ReportTarget  //2
	states_         []COM_ReportState //3
}

func (this *COM_ReportActionCounter) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.casterId_ != 0)
	mask.WriteBit(this.targetPosition_ != 0)
	mask.WriteBit(true) //props_
	mask.WriteBit(len(this.states_) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
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
	// serialize targetPosition_
	{
		if this.targetPosition_ != 0 {
			err := prpc.Write(buffer, this.targetPosition_)
			if err != nil {
				return err
			}
		}
	}
	// serialize props_
	{
		err := this.props_.Serialize(buffer)
		if err != nil {
			return err
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
func (this *COM_ReportActionCounter) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize casterId_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.casterId_)
		if err != nil {
			return err
		}
	}
	// deserialize targetPosition_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.targetPosition_)
		if err != nil {
			return err
		}
	}
	// deserialize props_
	if mask.ReadBit() {
		err := this.props_.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	// deserialize states_
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.states_ = make([]COM_ReportState, size)
		for i, _ := range this.states_ {
			err := this.states_[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
