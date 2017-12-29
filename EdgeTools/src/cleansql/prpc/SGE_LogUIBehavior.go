package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type SGE_LogUIBehavior struct {
	accountGuid_ uint32 //0
	accountName_ string //1
	playerGuid_  uint32 //2
	playerName_  string //3
	clientIp_    string //4
	type_        int    //5
}

func (this *SGE_LogUIBehavior) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.accountGuid_ != 0)
	mask.WriteBit(len(this.accountName_) != 0)
	mask.WriteBit(this.playerGuid_ != 0)
	mask.WriteBit(len(this.playerName_) != 0)
	mask.WriteBit(len(this.clientIp_) != 0)
	mask.WriteBit(this.type_ != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize accountGuid_
	{
		if this.accountGuid_ != 0 {
			err := prpc.Write(buffer, this.accountGuid_)
			if err != nil {
				return err
			}
		}
	}
	// serialize accountName_
	if len(this.accountName_) != 0 {
		err := prpc.Write(buffer, this.accountName_)
		if err != nil {
			return err
		}
	}
	// serialize playerGuid_
	{
		if this.playerGuid_ != 0 {
			err := prpc.Write(buffer, this.playerGuid_)
			if err != nil {
				return err
			}
		}
	}
	// serialize playerName_
	if len(this.playerName_) != 0 {
		err := prpc.Write(buffer, this.playerName_)
		if err != nil {
			return err
		}
	}
	// serialize clientIp_
	if len(this.clientIp_) != 0 {
		err := prpc.Write(buffer, this.clientIp_)
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
	return nil
}
func (this *SGE_LogUIBehavior) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize accountGuid_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.accountGuid_)
		if err != nil {
			return err
		}
	}
	// deserialize accountName_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.accountName_)
		if err != nil {
			return err
		}
	}
	// deserialize playerGuid_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.playerGuid_)
		if err != nil {
			return err
		}
	}
	// deserialize playerName_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.playerName_)
		if err != nil {
			return err
		}
	}
	// deserialize clientIp_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.clientIp_)
		if err != nil {
			return err
		}
	}
	// deserialize type_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.type_)
		if err != nil {
			return err
		}
	}
	return nil
}
