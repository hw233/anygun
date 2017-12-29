package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type SGE_Account struct {
	gamename_   string //0
	pfid_       string //1
	pfname_     string //2
	accountid_  string //3
	createtime_ uint64 //4
	mac_        string //5
	idfa_       string //6
	ip_         string //7
	devicetype_ string //8
}

func (this *SGE_Account) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(2)
	mask.WriteBit(len(this.gamename_) != 0)
	mask.WriteBit(len(this.pfid_) != 0)
	mask.WriteBit(len(this.pfname_) != 0)
	mask.WriteBit(len(this.accountid_) != 0)
	mask.WriteBit(this.createtime_ != 0)
	mask.WriteBit(len(this.mac_) != 0)
	mask.WriteBit(len(this.idfa_) != 0)
	mask.WriteBit(len(this.ip_) != 0)
	mask.WriteBit(len(this.devicetype_) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize gamename_
	if len(this.gamename_) != 0 {
		err := prpc.Write(buffer, this.gamename_)
		if err != nil {
			return err
		}
	}
	// serialize pfid_
	if len(this.pfid_) != 0 {
		err := prpc.Write(buffer, this.pfid_)
		if err != nil {
			return err
		}
	}
	// serialize pfname_
	if len(this.pfname_) != 0 {
		err := prpc.Write(buffer, this.pfname_)
		if err != nil {
			return err
		}
	}
	// serialize accountid_
	if len(this.accountid_) != 0 {
		err := prpc.Write(buffer, this.accountid_)
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
	// serialize ip_
	if len(this.ip_) != 0 {
		err := prpc.Write(buffer, this.ip_)
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
func (this *SGE_Account) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 2)
	if err != nil {
		return err
	}
	// deserialize gamename_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.gamename_)
		if err != nil {
			return err
		}
	}
	// deserialize pfid_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.pfid_)
		if err != nil {
			return err
		}
	}
	// deserialize pfname_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.pfname_)
		if err != nil {
			return err
		}
	}
	// deserialize accountid_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.accountid_)
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
	// deserialize ip_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.ip_)
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
