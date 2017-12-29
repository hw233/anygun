package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type SGE_LogRole struct {
	gamename_      string //0
	pfid_          string //1
	pfname_        string //2
	roleid_        uint32 //3
	cachetime_     uint64 //4
	accountid_     string //5
	serverid_      int8   //6
	servername_    string //7
	firstserid_    int8   //8
	rolefirsttime_ uint64 //9
	rolelasttime_  uint64 //10
	rolelv_        int8   //11
	gold_          int64  //12
	diamond_       int64  //13
	magicgold_     int64  //14
	vip_           int8   //15
	ce_            int64  //16
}

func (this *SGE_LogRole) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(3)
	mask.WriteBit(len(this.gamename_) != 0)
	mask.WriteBit(len(this.pfid_) != 0)
	mask.WriteBit(len(this.pfname_) != 0)
	mask.WriteBit(this.roleid_ != 0)
	mask.WriteBit(this.cachetime_ != 0)
	mask.WriteBit(len(this.accountid_) != 0)
	mask.WriteBit(this.serverid_ != 0)
	mask.WriteBit(len(this.servername_) != 0)
	mask.WriteBit(this.firstserid_ != 0)
	mask.WriteBit(this.rolefirsttime_ != 0)
	mask.WriteBit(this.rolelasttime_ != 0)
	mask.WriteBit(this.rolelv_ != 0)
	mask.WriteBit(this.gold_ != 0)
	mask.WriteBit(this.diamond_ != 0)
	mask.WriteBit(this.magicgold_ != 0)
	mask.WriteBit(this.vip_ != 0)
	mask.WriteBit(this.ce_ != 0)
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
	// serialize roleid_
	{
		if this.roleid_ != 0 {
			err := prpc.Write(buffer, this.roleid_)
			if err != nil {
				return err
			}
		}
	}
	// serialize cachetime_
	{
		if this.cachetime_ != 0 {
			err := prpc.Write(buffer, this.cachetime_)
			if err != nil {
				return err
			}
		}
	}
	// serialize accountid_
	if len(this.accountid_) != 0 {
		err := prpc.Write(buffer, this.accountid_)
		if err != nil {
			return err
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
	// serialize servername_
	if len(this.servername_) != 0 {
		err := prpc.Write(buffer, this.servername_)
		if err != nil {
			return err
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
	// serialize rolefirsttime_
	{
		if this.rolefirsttime_ != 0 {
			err := prpc.Write(buffer, this.rolefirsttime_)
			if err != nil {
				return err
			}
		}
	}
	// serialize rolelasttime_
	{
		if this.rolelasttime_ != 0 {
			err := prpc.Write(buffer, this.rolelasttime_)
			if err != nil {
				return err
			}
		}
	}
	// serialize rolelv_
	{
		if this.rolelv_ != 0 {
			err := prpc.Write(buffer, this.rolelv_)
			if err != nil {
				return err
			}
		}
	}
	// serialize gold_
	{
		if this.gold_ != 0 {
			err := prpc.Write(buffer, this.gold_)
			if err != nil {
				return err
			}
		}
	}
	// serialize diamond_
	{
		if this.diamond_ != 0 {
			err := prpc.Write(buffer, this.diamond_)
			if err != nil {
				return err
			}
		}
	}
	// serialize magicgold_
	{
		if this.magicgold_ != 0 {
			err := prpc.Write(buffer, this.magicgold_)
			if err != nil {
				return err
			}
		}
	}
	// serialize vip_
	{
		if this.vip_ != 0 {
			err := prpc.Write(buffer, this.vip_)
			if err != nil {
				return err
			}
		}
	}
	// serialize ce_
	{
		if this.ce_ != 0 {
			err := prpc.Write(buffer, this.ce_)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *SGE_LogRole) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 3)
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
	// deserialize roleid_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.roleid_)
		if err != nil {
			return err
		}
	}
	// deserialize cachetime_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.cachetime_)
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
	// deserialize serverid_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.serverid_)
		if err != nil {
			return err
		}
	}
	// deserialize servername_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.servername_)
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
	// deserialize rolefirsttime_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.rolefirsttime_)
		if err != nil {
			return err
		}
	}
	// deserialize rolelasttime_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.rolelasttime_)
		if err != nil {
			return err
		}
	}
	// deserialize rolelv_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.rolelv_)
		if err != nil {
			return err
		}
	}
	// deserialize gold_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.gold_)
		if err != nil {
			return err
		}
	}
	// deserialize diamond_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.diamond_)
		if err != nil {
			return err
		}
	}
	// deserialize magicgold_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.magicgold_)
		if err != nil {
			return err
		}
	}
	// deserialize vip_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.vip_)
		if err != nil {
			return err
		}
	}
	// deserialize ce_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.ce_)
		if err != nil {
			return err
		}
	}
	return nil
}
