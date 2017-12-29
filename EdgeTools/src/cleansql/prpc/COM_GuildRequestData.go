package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type COM_GuildRequestData struct {
	roleId_    int64  //0
	level_     int8   //1
	roleName_  string //2
	time_      int32  //3
	prof_      int8   //4
	profLevel_ int8   //5
}

func (this *COM_GuildRequestData) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.roleId_ != 0)
	mask.WriteBit(this.level_ != 0)
	mask.WriteBit(len(this.roleName_) != 0)
	mask.WriteBit(this.time_ != 0)
	mask.WriteBit(this.prof_ != 0)
	mask.WriteBit(this.profLevel_ != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize roleId_
	{
		if this.roleId_ != 0 {
			err := prpc.Write(buffer, this.roleId_)
			if err != nil {
				return err
			}
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
	// serialize roleName_
	if len(this.roleName_) != 0 {
		err := prpc.Write(buffer, this.roleName_)
		if err != nil {
			return err
		}
	}
	// serialize time_
	{
		if this.time_ != 0 {
			err := prpc.Write(buffer, this.time_)
			if err != nil {
				return err
			}
		}
	}
	// serialize prof_
	{
		if this.prof_ != 0 {
			err := prpc.Write(buffer, this.prof_)
			if err != nil {
				return err
			}
		}
	}
	// serialize profLevel_
	{
		if this.profLevel_ != 0 {
			err := prpc.Write(buffer, this.profLevel_)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *COM_GuildRequestData) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize roleId_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.roleId_)
		if err != nil {
			return err
		}
	}
	// deserialize level_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.level_)
		if err != nil {
			return err
		}
	}
	// deserialize roleName_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.roleName_)
		if err != nil {
			return err
		}
	}
	// deserialize time_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.time_)
		if err != nil {
			return err
		}
	}
	// deserialize prof_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.prof_)
		if err != nil {
			return err
		}
	}
	// deserialize profLevel_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.profLevel_)
		if err != nil {
			return err
		}
	}
	return nil
}
