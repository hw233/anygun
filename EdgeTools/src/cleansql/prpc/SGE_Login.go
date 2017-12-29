package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type SGE_Login struct {
	gamename_      string //0
	pfid_          string //1
	pfname_        string //2
	accountid_     string //3
	roleId_        uint32 //4
	logintime_     uint64 //5
	logouttime_    uint64 //6
	firsttime_     uint64 //7
	rolefirsttime_ uint64 //8
	firstserid_    uint32 //9
	serverid_      uint32 //10
	mac_           string //11
	idfa_          string //12
	ip_            string //13
	devicetype_    string //14
}

func (this *SGE_Login) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(2)
	mask.WriteBit(len(this.gamename_) != 0)
	mask.WriteBit(len(this.pfid_) != 0)
	mask.WriteBit(len(this.pfname_) != 0)
	mask.WriteBit(len(this.accountid_) != 0)
	mask.WriteBit(this.roleId_ != 0)
	mask.WriteBit(this.logintime_ != 0)
	mask.WriteBit(this.logouttime_ != 0)
	mask.WriteBit(this.firsttime_ != 0)
	mask.WriteBit(this.rolefirsttime_ != 0)
	mask.WriteBit(this.firstserid_ != 0)
	mask.WriteBit(this.serverid_ != 0)
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
	// serialize roleId_
	{
		if this.roleId_ != 0 {
			err := prpc.Write(buffer, this.roleId_)
			if err != nil {
				return err
			}
		}
	}
	// serialize logintime_
	{
		if this.logintime_ != 0 {
			err := prpc.Write(buffer, this.logintime_)
			if err != nil {
				return err
			}
		}
	}
	// serialize logouttime_
	{
		if this.logouttime_ != 0 {
			err := prpc.Write(buffer, this.logouttime_)
			if err != nil {
				return err
			}
		}
	}
	// serialize firsttime_
	{
		if this.firsttime_ != 0 {
			err := prpc.Write(buffer, this.firsttime_)
			if err != nil {
				return err
			}
		}
	}
	// serialize rolefirsttime_
	{
		if this.rolefirsttime_ != 0 {
			err := prpc.Write(buffer, this.rolefirsttime_)
			if err != nil {
				return err
			}
		}
	}
	// serialize firstserid_
	{
		if this.firstserid_ != 0 {
			err := prpc.Write(buffer, this.firstserid_)
			if err != nil {
				return err
			}
		}
	}
	// serialize serverid_
	{
		if this.serverid_ != 0 {
			err := prpc.Write(buffer, this.serverid_)
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
func (this *SGE_Login) Deserialize(buffer *bytes.Buffer) error {
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
	// deserialize roleId_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.roleId_)
		if err != nil {
			return err
		}
	}
	// deserialize logintime_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.logintime_)
		if err != nil {
			return err
		}
	}
	// deserialize logouttime_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.logouttime_)
		if err != nil {
			return err
		}
	}
	// deserialize firsttime_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.firsttime_)
		if err != nil {
			return err
		}
	}
	// deserialize rolefirsttime_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.rolefirsttime_)
		if err != nil {
			return err
		}
	}
	// deserialize firstserid_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.firstserid_)
		if err != nil {
			return err
		}
	}
	// deserialize serverid_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.serverid_)
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
