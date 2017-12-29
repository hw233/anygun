package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type COM_ActivityTable struct {
	activities_ []COM_Activity //0
	flag_       []uint32       //1
	reward_     int32          //2
}

func (this *COM_ActivityTable) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.activities_) != 0)
	mask.WriteBit(len(this.flag_) != 0)
	mask.WriteBit(this.reward_ != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize activities_
	if len(this.activities_) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.activities_)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.activities_ {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// serialize flag_
	if len(this.flag_) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.flag_)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.flag_ {
			err := prpc.Write(buffer, value)
			if err != nil {
				return err
			}
		}
	}
	// serialize reward_
	{
		if this.reward_ != 0 {
			err := prpc.Write(buffer, this.reward_)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *COM_ActivityTable) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize activities_
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.activities_ = make([]COM_Activity, size)
		for i, _ := range this.activities_ {
			err := this.activities_[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// deserialize flag_
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.flag_ = make([]uint32, size)
		for i, _ := range this.flag_ {
			err := prpc.Read(buffer, &this.flag_[i])
			if err != nil {
				return err
			}
		}
	}
	// deserialize reward_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.reward_)
		if err != nil {
			return err
		}
	}
	return nil
}
