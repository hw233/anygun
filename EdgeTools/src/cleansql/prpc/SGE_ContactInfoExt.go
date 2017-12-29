package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type SGE_ContactInfoExt struct {
	COM_ContactInfo
	rolefirst_       uint64 //0
	rolelast_        uint64 //1
	gold_            uint32 //2
	diamond_         uint32 //3
	magicgold_       uint32 //4
	guildContribute_ int32  //5
	accName_         string //6
	userid_          string //7
	pfid_            string //8
	pfname_          string //9
	serverid_        uint32 //10
}

func (this *SGE_ContactInfoExt) Serialize(buffer *bytes.Buffer) error {
	{
		err := this.COM_ContactInfo.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	//field mask
	mask := prpc.NewMask1(2)
	mask.WriteBit(this.rolefirst_ != 0)
	mask.WriteBit(this.rolelast_ != 0)
	mask.WriteBit(this.gold_ != 0)
	mask.WriteBit(this.diamond_ != 0)
	mask.WriteBit(this.magicgold_ != 0)
	mask.WriteBit(this.guildContribute_ != 0)
	mask.WriteBit(len(this.accName_) != 0)
	mask.WriteBit(len(this.userid_) != 0)
	mask.WriteBit(len(this.pfid_) != 0)
	mask.WriteBit(len(this.pfname_) != 0)
	mask.WriteBit(this.serverid_ != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize rolefirst_
	{
		if this.rolefirst_ != 0 {
			err := prpc.Write(buffer, this.rolefirst_)
			if err != nil {
				return err
			}
		}
	}
	// serialize rolelast_
	{
		if this.rolelast_ != 0 {
			err := prpc.Write(buffer, this.rolelast_)
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
	// serialize guildContribute_
	{
		if this.guildContribute_ != 0 {
			err := prpc.Write(buffer, this.guildContribute_)
			if err != nil {
				return err
			}
		}
	}
	// serialize accName_
	if len(this.accName_) != 0 {
		err := prpc.Write(buffer, this.accName_)
		if err != nil {
			return err
		}
	}
	// serialize userid_
	if len(this.userid_) != 0 {
		err := prpc.Write(buffer, this.userid_)
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
	// serialize serverid_
	{
		if this.serverid_ != 0 {
			err := prpc.Write(buffer, this.serverid_)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *SGE_ContactInfoExt) Deserialize(buffer *bytes.Buffer) error {
	{
		this.COM_ContactInfo.Deserialize(buffer)
	}
	//field mask
	mask, err := prpc.NewMask0(buffer, 2)
	if err != nil {
		return err
	}
	// deserialize rolefirst_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.rolefirst_)
		if err != nil {
			return err
		}
	}
	// deserialize rolelast_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.rolelast_)
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
	// deserialize guildContribute_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.guildContribute_)
		if err != nil {
			return err
		}
	}
	// deserialize accName_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.accName_)
		if err != nil {
			return err
		}
	}
	// deserialize userid_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.userid_)
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
	// deserialize serverid_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.serverid_)
		if err != nil {
			return err
		}
	}
	return nil
}
