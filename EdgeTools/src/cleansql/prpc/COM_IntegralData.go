package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type COM_IntegralData struct {
	sinceStamp_ uint64                //0
	endStamp_   uint64                //1
	integral_   uint32                //2
	isflag_     bool                  //3
	contents_   []COM_IntegralContent //4
}

func (this *COM_IntegralData) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.sinceStamp_ != 0)
	mask.WriteBit(this.endStamp_ != 0)
	mask.WriteBit(this.integral_ != 0)
	mask.WriteBit(this.isflag_)
	mask.WriteBit(len(this.contents_) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize sinceStamp_
	{
		if this.sinceStamp_ != 0 {
			err := prpc.Write(buffer, this.sinceStamp_)
			if err != nil {
				return err
			}
		}
	}
	// serialize endStamp_
	{
		if this.endStamp_ != 0 {
			err := prpc.Write(buffer, this.endStamp_)
			if err != nil {
				return err
			}
		}
	}
	// serialize integral_
	{
		if this.integral_ != 0 {
			err := prpc.Write(buffer, this.integral_)
			if err != nil {
				return err
			}
		}
	}
	// serialize isflag_
	{
	}
	// serialize contents_
	if len(this.contents_) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.contents_)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.contents_ {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *COM_IntegralData) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize sinceStamp_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.sinceStamp_)
		if err != nil {
			return err
		}
	}
	// deserialize endStamp_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.endStamp_)
		if err != nil {
			return err
		}
	}
	// deserialize integral_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.integral_)
		if err != nil {
			return err
		}
	}
	// deserialize isflag_
	this.isflag_ = mask.ReadBit()
	// deserialize contents_
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.contents_ = make([]COM_IntegralContent, size)
		for i, _ := range this.contents_ {
			err := this.contents_[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
