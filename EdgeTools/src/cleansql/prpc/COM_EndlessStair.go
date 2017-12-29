package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type COM_EndlessStair struct {
	rank_      uint32  //0
	name_      string  //1
	job_       int     //2
	joblevel_  uint32  //3
	assetId_   int32   //4
	level_     uint32  //5
	rivalTime_ float32 //6
	rivalNum_  int32   //7
}

func (this *COM_EndlessStair) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.rank_ != 0)
	mask.WriteBit(len(this.name_) != 0)
	mask.WriteBit(this.job_ != 0)
	mask.WriteBit(this.joblevel_ != 0)
	mask.WriteBit(this.assetId_ != 0)
	mask.WriteBit(this.level_ != 0)
	mask.WriteBit(this.rivalTime_ != 0)
	mask.WriteBit(this.rivalNum_ != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
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
	// serialize job_
	{
		if this.job_ != 0 {
			err := prpc.Write(buffer, this.job_)
			if err != nil {
				return err
			}
		}
	}
	// serialize joblevel_
	{
		if this.joblevel_ != 0 {
			err := prpc.Write(buffer, this.joblevel_)
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
	// serialize level_
	{
		if this.level_ != 0 {
			err := prpc.Write(buffer, this.level_)
			if err != nil {
				return err
			}
		}
	}
	// serialize rivalTime_
	{
		if this.rivalTime_ != 0 {
			err := prpc.Write(buffer, this.rivalTime_)
			if err != nil {
				return err
			}
		}
	}
	// serialize rivalNum_
	{
		if this.rivalNum_ != 0 {
			err := prpc.Write(buffer, this.rivalNum_)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *COM_EndlessStair) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
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
	// deserialize job_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.job_)
		if err != nil {
			return err
		}
	}
	// deserialize joblevel_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.joblevel_)
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
	// deserialize level_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.level_)
		if err != nil {
			return err
		}
	}
	// deserialize rivalTime_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.rivalTime_)
		if err != nil {
			return err
		}
	}
	// deserialize rivalNum_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.rivalNum_)
		if err != nil {
			return err
		}
	}
	return nil
}
