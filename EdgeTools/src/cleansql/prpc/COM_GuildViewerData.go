package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type COM_GuildViewerData struct {
	guid_        uint32 //0
	guildName_   string //1
	playerName_  string //2
	notice_      string //3
	level_       int8   //4
	memberNum_   int16  //5
	memberTotal_ int16  //6
	guildRank_   int8   //7
}

func (this *COM_GuildViewerData) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.guid_ != 0)
	mask.WriteBit(len(this.guildName_) != 0)
	mask.WriteBit(len(this.playerName_) != 0)
	mask.WriteBit(len(this.notice_) != 0)
	mask.WriteBit(this.level_ != 0)
	mask.WriteBit(this.memberNum_ != 0)
	mask.WriteBit(this.memberTotal_ != 0)
	mask.WriteBit(this.guildRank_ != 0)
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
	// serialize guildName_
	if len(this.guildName_) != 0 {
		err := prpc.Write(buffer, this.guildName_)
		if err != nil {
			return err
		}
	}
	// serialize playerName_
	if len(this.playerName_) != 0 {
		err := prpc.Write(buffer, this.playerName_)
		if err != nil {
			return err
		}
	}
	// serialize notice_
	if len(this.notice_) != 0 {
		err := prpc.Write(buffer, this.notice_)
		if err != nil {
			return err
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
	// serialize memberNum_
	{
		if this.memberNum_ != 0 {
			err := prpc.Write(buffer, this.memberNum_)
			if err != nil {
				return err
			}
		}
	}
	// serialize memberTotal_
	{
		if this.memberTotal_ != 0 {
			err := prpc.Write(buffer, this.memberTotal_)
			if err != nil {
				return err
			}
		}
	}
	// serialize guildRank_
	{
		if this.guildRank_ != 0 {
			err := prpc.Write(buffer, this.guildRank_)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *COM_GuildViewerData) Deserialize(buffer *bytes.Buffer) error {
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
	// deserialize guildName_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.guildName_)
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
	// deserialize notice_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.notice_)
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
	// deserialize memberNum_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.memberNum_)
		if err != nil {
			return err
		}
	}
	// deserialize memberTotal_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.memberTotal_)
		if err != nil {
			return err
		}
	}
	// deserialize guildRank_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.guildRank_)
		if err != nil {
			return err
		}
	}
	return nil
}
