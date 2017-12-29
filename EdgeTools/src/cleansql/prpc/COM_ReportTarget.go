package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type COM_ReportTarget struct {
	COM_ReportTargetBase
	prop2_ []COM_ReportTargetBase //0
}

func (this *COM_ReportTarget) Serialize(buffer *bytes.Buffer) error {
	{
		err := this.COM_ReportTargetBase.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.prop2_) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize prop2_
	if len(this.prop2_) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.prop2_)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.prop2_ {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *COM_ReportTarget) Deserialize(buffer *bytes.Buffer) error {
	{
		this.COM_ReportTargetBase.Deserialize(buffer)
	}
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize prop2_
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.prop2_ = make([]COM_ReportTargetBase, size)
		for i, _ := range this.prop2_ {
			err := this.prop2_[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
