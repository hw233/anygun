package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type COM_GuildBuilding struct {
	level_     int32 //0
	struction_ int32 //1
}

func (this *COM_GuildBuilding) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.level_ != 0)
	mask.WriteBit(this.struction_ != 0)
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
	// serialize struction_
	{
		if this.struction_ != 0 {
			err := prpc.Write(buffer, this.struction_)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *COM_GuildBuilding) Deserialize(buffer *bytes.Buffer) error {
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
	// deserialize struction_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.struction_)
		if err != nil {
			return err
		}
	}
	return nil
}
