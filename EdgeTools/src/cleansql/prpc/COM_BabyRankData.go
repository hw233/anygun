package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type COM_BabyRankData struct {
	instId_    uint32 //0
	rank_      int32  //1
	name_      string //2
	ownerName_ string //3
	ff_        int32  //4
}

func (this *COM_BabyRankData) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.instId_ != 0)
	mask.WriteBit(this.rank_ != 0)
	mask.WriteBit(len(this.name_) != 0)
	mask.WriteBit(len(this.ownerName_) != 0)
	mask.WriteBit(this.ff_ != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
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
	// serialize rank_
	{
		if this.rank_ != 0 {
			err := prpc.Write(buffer, this.rank_)
			if err != nil {
				return err
			}
		}
	}
	// serialize name_
	if len(this.name_) != 0 {
		err := prpc.Write(buffer, this.name_)
		if err != nil {
			return err
		}
	}
	// serialize ownerName_
	if len(this.ownerName_) != 0 {
		err := prpc.Write(buffer, this.ownerName_)
		if err != nil {
			return err
		}
	}
	// serialize ff_
	{
		if this.ff_ != 0 {
			err := prpc.Write(buffer, this.ff_)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *COM_BabyRankData) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize instId_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.instId_)
		if err != nil {
			return err
		}
	}
	// deserialize rank_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.rank_)
		if err != nil {
			return err
		}
	}
	// deserialize name_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.name_)
		if err != nil {
			return err
		}
	}
	// deserialize ownerName_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.ownerName_)
		if err != nil {
			return err
		}
	}
	// deserialize ff_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.ff_)
		if err != nil {
			return err
		}
	}
	return nil
}
