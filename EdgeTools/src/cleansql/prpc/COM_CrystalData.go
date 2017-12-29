package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type COM_CrystalData struct {
	level_ uint32            //0
	props_ []COM_CrystalProp //1
}

func (this *COM_CrystalData) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.level_ != 0)
	mask.WriteBit(len(this.props_) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize level_
	{
		if this.level_ != 0 {
			err := prpc.Write(buffer, this.level_)
			if err != nil {
				return err
			}
		}
	}
	// serialize props_
	if len(this.props_) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.props_)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.props_ {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *COM_CrystalData) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize level_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.level_)
		if err != nil {
			return err
		}
	}
	// deserialize props_
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.props_ = make([]COM_CrystalProp, size)
		for i, _ := range this.props_ {
			err := this.props_[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
