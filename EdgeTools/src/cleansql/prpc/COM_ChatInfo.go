package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type COM_ChatInfo struct {
	COM_Chat
	audioId_    int32  //0
	assetId_    uint16 //1
	playerName_ string //2
	guildName_  string //3
	instId_     uint32 //4
}

func (this *COM_ChatInfo) Serialize(buffer *bytes.Buffer) error {
	{
		err := this.COM_Chat.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.audioId_ != 0)
	mask.WriteBit(this.assetId_ != 0)
	mask.WriteBit(len(this.playerName_) != 0)
	mask.WriteBit(len(this.guildName_) != 0)
	mask.WriteBit(this.instId_ != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize audioId_
	{
		if this.audioId_ != 0 {
			err := prpc.Write(buffer, this.audioId_)
			if err != nil {
				return err
			}
		}
	}
	// serialize assetId_
	{
		if this.assetId_ != 0 {
			err := prpc.Write(buffer, this.assetId_)
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
	// serialize guildName_
	if len(this.guildName_) != 0 {
		err := prpc.Write(buffer, this.guildName_)
		if err != nil {
			return err
		}
	}
	// serialize instId_
	{
		if this.instId_ != 0 {
			err := prpc.Write(buffer, this.instId_)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *COM_ChatInfo) Deserialize(buffer *bytes.Buffer) error {
	{
		this.COM_Chat.Deserialize(buffer)
	}
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize audioId_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.audioId_)
		if err != nil {
			return err
		}
	}
	// deserialize assetId_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.assetId_)
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
	// deserialize guildName_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.guildName_)
		if err != nil {
			return err
		}
	}
	// deserialize instId_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.instId_)
		if err != nil {
			return err
		}
	}
	return nil
}
