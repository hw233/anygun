package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type COM_CreateTeamInfo struct {
	type_          int    //0
	maxMemberSize_ uint8  //1
	name_          string //2
	pwd_           string //3
	minLevel_      uint16 //4
	maxLevel_      uint16 //5
}

func (this *COM_CreateTeamInfo) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.type_ != 0)
	mask.WriteBit(this.maxMemberSize_ != 0)
	mask.WriteBit(len(this.name_) != 0)
	mask.WriteBit(len(this.pwd_) != 0)
	mask.WriteBit(this.minLevel_ != 0)
	mask.WriteBit(this.maxLevel_ != 0)
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
	// serialize maxMemberSize_
	{
		if this.maxMemberSize_ != 0 {
			err := prpc.Write(buffer, this.maxMemberSize_)
			if err != nil {
				return err
			}
		}
	}
	// serialize name_
	if len(this.name_) != 0 {
		err := prpc.Write(buffer, this.name_)
		if err != nil {
			return err
		}
	}
	// serialize pwd_
	if len(this.pwd_) != 0 {
		err := prpc.Write(buffer, this.pwd_)
		if err != nil {
			return err
		}
	}
	// serialize minLevel_
	{
		if this.minLevel_ != 0 {
			err := prpc.Write(buffer, this.minLevel_)
			if err != nil {
				return err
			}
		}
	}
	// serialize maxLevel_
	{
		if this.maxLevel_ != 0 {
			err := prpc.Write(buffer, this.maxLevel_)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *COM_CreateTeamInfo) Deserialize(buffer *bytes.Buffer) error {
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
	// deserialize maxMemberSize_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.maxMemberSize_)
		if err != nil {
			return err
		}
	}
	// deserialize name_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.name_)
		if err != nil {
			return err
		}
	}
	// deserialize pwd_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.pwd_)
		if err != nil {
			return err
		}
	}
	// deserialize minLevel_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.minLevel_)
		if err != nil {
			return err
		}
	}
	// deserialize maxLevel_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.maxLevel_)
		if err != nil {
			return err
		}
	}
	return nil
}
