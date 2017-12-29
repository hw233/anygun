package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type COM_Mail struct {
	mailId_         int32          //0
	mailType_       int            //1
	timestamp_      int64          //2
	sendPlayerName_ string         //3
	recvPlayerName_ string         //4
	title_          string         //5
	content_        string         //6
	money_          int32          //7
	diamond_        int32          //8
	items_          []COM_MailItem //9
	isRead_         bool           //10
}

func (this *COM_Mail) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(2)
	mask.WriteBit(this.mailId_ != 0)
	mask.WriteBit(this.mailType_ != 0)
	mask.WriteBit(this.timestamp_ != 0)
	mask.WriteBit(len(this.sendPlayerName_) != 0)
	mask.WriteBit(len(this.recvPlayerName_) != 0)
	mask.WriteBit(len(this.title_) != 0)
	mask.WriteBit(len(this.content_) != 0)
	mask.WriteBit(this.money_ != 0)
	mask.WriteBit(this.diamond_ != 0)
	mask.WriteBit(len(this.items_) != 0)
	mask.WriteBit(this.isRead_)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize mailId_
	{
		if this.mailId_ != 0 {
			err := prpc.Write(buffer, this.mailId_)
			if err != nil {
				return err
			}
		}
	}
	// serialize mailType_
	{
		if this.mailType_ != 0 {
			err := prpc.Write(buffer, this.mailType_)
			if err != nil {
				return err
			}
		}
	}
	// serialize timestamp_
	{
		if this.timestamp_ != 0 {
			err := prpc.Write(buffer, this.timestamp_)
			if err != nil {
				return err
			}
		}
	}
	// serialize sendPlayerName_
	if len(this.sendPlayerName_) != 0 {
		err := prpc.Write(buffer, this.sendPlayerName_)
		if err != nil {
			return err
		}
	}
	// serialize recvPlayerName_
	if len(this.recvPlayerName_) != 0 {
		err := prpc.Write(buffer, this.recvPlayerName_)
		if err != nil {
			return err
		}
	}
	// serialize title_
	if len(this.title_) != 0 {
		err := prpc.Write(buffer, this.title_)
		if err != nil {
			return err
		}
	}
	// serialize content_
	if len(this.content_) != 0 {
		err := prpc.Write(buffer, this.content_)
		if err != nil {
			return err
		}
	}
	// serialize money_
	{
		if this.money_ != 0 {
			err := prpc.Write(buffer, this.money_)
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
	// serialize items_
	if len(this.items_) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.items_)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.items_ {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// serialize isRead_
	{
	}
	return nil
}
func (this *COM_Mail) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 2)
	if err != nil {
		return err
	}
	// deserialize mailId_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.mailId_)
		if err != nil {
			return err
		}
	}
	// deserialize mailType_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.mailType_)
		if err != nil {
			return err
		}
	}
	// deserialize timestamp_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.timestamp_)
		if err != nil {
			return err
		}
	}
	// deserialize sendPlayerName_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.sendPlayerName_)
		if err != nil {
			return err
		}
	}
	// deserialize recvPlayerName_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.recvPlayerName_)
		if err != nil {
			return err
		}
	}
	// deserialize title_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.title_)
		if err != nil {
			return err
		}
	}
	// deserialize content_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.content_)
		if err != nil {
			return err
		}
	}
	// deserialize money_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.money_)
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
	// deserialize items_
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.items_ = make([]COM_MailItem, size)
		for i, _ := range this.items_ {
			err := this.items_[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// deserialize isRead_
	this.isRead_ = mask.ReadBit()
	return nil
}
