package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type COM_ADLoginTotalContent struct {
	totalDays_  uint32   //0
	itemIds_    []uint32 //1
	itemStacks_ []uint32 //2
	status_     uint8    //3
}

func (this *COM_ADLoginTotalContent) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.totalDays_ != 0)
	mask.WriteBit(len(this.itemIds_) != 0)
	mask.WriteBit(len(this.itemStacks_) != 0)
	mask.WriteBit(this.status_ != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize totalDays_
	{
		if this.totalDays_ != 0 {
			err := prpc.Write(buffer, this.totalDays_)
			if err != nil {
				return err
			}
		}
	}
	// serialize itemIds_
	if len(this.itemIds_) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.itemIds_)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.itemIds_ {
			err := prpc.Write(buffer, value)
			if err != nil {
				return err
			}
		}
	}
	// serialize itemStacks_
	if len(this.itemStacks_) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.itemStacks_)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.itemStacks_ {
			err := prpc.Write(buffer, value)
			if err != nil {
				return err
			}
		}
	}
	// serialize status_
	{
		if this.status_ != 0 {
			err := prpc.Write(buffer, this.status_)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *COM_ADLoginTotalContent) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize totalDays_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.totalDays_)
		if err != nil {
			return err
		}
	}
	// deserialize itemIds_
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.itemIds_ = make([]uint32, size)
		for i, _ := range this.itemIds_ {
			err := prpc.Read(buffer, &this.itemIds_[i])
			if err != nil {
				return err
			}
		}
	}
	// deserialize itemStacks_
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.itemStacks_ = make([]uint32, size)
		for i, _ := range this.itemStacks_ {
			err := prpc.Read(buffer, &this.itemStacks_[i])
			if err != nil {
				return err
			}
		}
	}
	// deserialize status_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.status_)
		if err != nil {
			return err
		}
	}
	return nil
}
