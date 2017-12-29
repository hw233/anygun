package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type COM_Chat struct {
	ck_           int     //0
	isAudio_      bool    //1
	content_      string  //2
	audioTime_    int32   //3
	audio_        []uint8 //4
	isMe          bool    //5
	teamId_       int32   //6
	teamType_     int     //7
	teamMinLevel_ int16   //8
	teamMaxLevel_ int16   //9
	teamName_     string  //10
}

func (this *COM_Chat) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(2)
	mask.WriteBit(this.ck_ != 0)
	mask.WriteBit(this.isAudio_)
	mask.WriteBit(len(this.content_) != 0)
	mask.WriteBit(this.audioTime_ != 0)
	mask.WriteBit(len(this.audio_) != 0)
	mask.WriteBit(this.isMe)
	mask.WriteBit(this.teamId_ != 0)
	mask.WriteBit(this.teamType_ != 0)
	mask.WriteBit(this.teamMinLevel_ != 0)
	mask.WriteBit(this.teamMaxLevel_ != 0)
	mask.WriteBit(len(this.teamName_) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize ck_
	{
		if this.ck_ != 0 {
			err := prpc.Write(buffer, this.ck_)
			if err != nil {
				return err
			}
		}
	}
	// serialize isAudio_
	{
	}
	// serialize content_
	if len(this.content_) != 0 {
		err := prpc.Write(buffer, this.content_)
		if err != nil {
			return err
		}
	}
	// serialize audioTime_
	{
		if this.audioTime_ != 0 {
			err := prpc.Write(buffer, this.audioTime_)
			if err != nil {
				return err
			}
		}
	}
	// serialize audio_
	if len(this.audio_) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.audio_)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.audio_ {
			err := prpc.Write(buffer, value)
			if err != nil {
				return err
			}
		}
	}
	// serialize isMe
	{
	}
	// serialize teamId_
	{
		if this.teamId_ != 0 {
			err := prpc.Write(buffer, this.teamId_)
			if err != nil {
				return err
			}
		}
	}
	// serialize teamType_
	{
		if this.teamType_ != 0 {
			err := prpc.Write(buffer, this.teamType_)
			if err != nil {
				return err
			}
		}
	}
	// serialize teamMinLevel_
	{
		if this.teamMinLevel_ != 0 {
			err := prpc.Write(buffer, this.teamMinLevel_)
			if err != nil {
				return err
			}
		}
	}
	// serialize teamMaxLevel_
	{
		if this.teamMaxLevel_ != 0 {
			err := prpc.Write(buffer, this.teamMaxLevel_)
			if err != nil {
				return err
			}
		}
	}
	// serialize teamName_
	if len(this.teamName_) != 0 {
		err := prpc.Write(buffer, this.teamName_)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *COM_Chat) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 2)
	if err != nil {
		return err
	}
	// deserialize ck_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.ck_)
		if err != nil {
			return err
		}
	}
	// deserialize isAudio_
	this.isAudio_ = mask.ReadBit()
	// deserialize content_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.content_)
		if err != nil {
			return err
		}
	}
	// deserialize audioTime_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.audioTime_)
		if err != nil {
			return err
		}
	}
	// deserialize audio_
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.audio_ = make([]uint8, size)
		for i, _ := range this.audio_ {
			err := prpc.Read(buffer, &this.audio_[i])
			if err != nil {
				return err
			}
		}
	}
	// deserialize isMe
	this.isMe = mask.ReadBit()
	// deserialize teamId_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.teamId_)
		if err != nil {
			return err
		}
	}
	// deserialize teamType_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.teamType_)
		if err != nil {
			return err
		}
	}
	// deserialize teamMinLevel_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.teamMinLevel_)
		if err != nil {
			return err
		}
	}
	// deserialize teamMaxLevel_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.teamMaxLevel_)
		if err != nil {
			return err
		}
	}
	// deserialize teamName_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.teamName_)
		if err != nil {
			return err
		}
	}
	return nil
}
