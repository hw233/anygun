package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type COM_AccountInfo struct {
	guid_        uint32 //0
	username_    string //1
	password_    string //2
	createtime_  uint64 //3
	phoneNumber_ string //4
}

func (this *COM_AccountInfo) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.guid_ != 0)
	mask.WriteBit(len(this.username_) != 0)
	mask.WriteBit(len(this.password_) != 0)
	mask.WriteBit(this.createtime_ != 0)
	mask.WriteBit(len(this.phoneNumber_) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize guid_
	{
		if this.guid_ != 0 {
			err := prpc.Write(buffer, this.guid_)
			if err != nil {
				return err
			}
		}
	}
	// serialize username_
	if len(this.username_) != 0 {
		err := prpc.Write(buffer, this.username_)
		if err != nil {
			return err
		}
	}
	// serialize password_
	if len(this.password_) != 0 {
		err := prpc.Write(buffer, this.password_)
		if err != nil {
			return err
		}
	}
	// serialize createtime_
	{
		if this.createtime_ != 0 {
			err := prpc.Write(buffer, this.createtime_)
			if err != nil {
				return err
			}
		}
	}
	// serialize phoneNumber_
	if len(this.phoneNumber_) != 0 {
		err := prpc.Write(buffer, this.phoneNumber_)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *COM_AccountInfo) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize guid_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.guid_)
		if err != nil {
			return err
		}
	}
	// deserialize username_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.username_)
		if err != nil {
			return err
		}
	}
	// deserialize password_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.password_)
		if err != nil {
			return err
		}
	}
	// deserialize createtime_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.createtime_)
		if err != nil {
			return err
		}
	}
	// deserialize phoneNumber_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.phoneNumber_)
		if err != nil {
			return err
		}
	}
	return nil
}
