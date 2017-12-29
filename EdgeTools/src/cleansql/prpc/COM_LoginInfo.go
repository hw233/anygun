package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type COM_LoginInfo struct {
	username_   string //0
	password_   string //1
	version_    uint32 //2
	sessionkey_ string //3
	mac_        string //4
	idfa_       string //5
	devicetype_ string //6
}

func (this *COM_LoginInfo) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.username_) != 0)
	mask.WriteBit(len(this.password_) != 0)
	mask.WriteBit(this.version_ != 0)
	mask.WriteBit(len(this.sessionkey_) != 0)
	mask.WriteBit(len(this.mac_) != 0)
	mask.WriteBit(len(this.idfa_) != 0)
	mask.WriteBit(len(this.devicetype_) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
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
	// serialize version_
	{
		if this.version_ != 0 {
			err := prpc.Write(buffer, this.version_)
			if err != nil {
				return err
			}
		}
	}
	// serialize sessionkey_
	if len(this.sessionkey_) != 0 {
		err := prpc.Write(buffer, this.sessionkey_)
		if err != nil {
			return err
		}
	}
	// serialize mac_
	if len(this.mac_) != 0 {
		err := prpc.Write(buffer, this.mac_)
		if err != nil {
			return err
		}
	}
	// serialize idfa_
	if len(this.idfa_) != 0 {
		err := prpc.Write(buffer, this.idfa_)
		if err != nil {
			return err
		}
	}
	// serialize devicetype_
	if len(this.devicetype_) != 0 {
		err := prpc.Write(buffer, this.devicetype_)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *COM_LoginInfo) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
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
	// deserialize version_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.version_)
		if err != nil {
			return err
		}
	}
	// deserialize sessionkey_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.sessionkey_)
		if err != nil {
			return err
		}
	}
	// deserialize mac_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.mac_)
		if err != nil {
			return err
		}
	}
	// deserialize idfa_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.idfa_)
		if err != nil {
			return err
		}
	}
	// deserialize devicetype_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.devicetype_)
		if err != nil {
			return err
		}
	}
	return nil
}
