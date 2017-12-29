package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type COM_SearchContext struct {
	title_  string //0
	ist_    int    //1
	rt_     int    //2
	itemId_ int32  //3
	babyId_ int32  //4
	begin_  int32  //5
	limit_  int32  //6
}

func (this *COM_SearchContext) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.title_) != 0)
	mask.WriteBit(this.ist_ != 0)
	mask.WriteBit(this.rt_ != 0)
	mask.WriteBit(this.itemId_ != 0)
	mask.WriteBit(this.babyId_ != 0)
	mask.WriteBit(this.begin_ != 0)
	mask.WriteBit(this.limit_ != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
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
	// serialize ist_
	{
		if this.ist_ != 0 {
			err := prpc.Write(buffer, this.ist_)
			if err != nil {
				return err
			}
		}
	}
	// serialize rt_
	{
		if this.rt_ != 0 {
			err := prpc.Write(buffer, this.rt_)
			if err != nil {
				return err
			}
		}
	}
	// serialize itemId_
	{
		if this.itemId_ != 0 {
			err := prpc.Write(buffer, this.itemId_)
			if err != nil {
				return err
			}
		}
	}
	// serialize babyId_
	{
		if this.babyId_ != 0 {
			err := prpc.Write(buffer, this.babyId_)
			if err != nil {
				return err
			}
		}
	}
	// serialize begin_
	{
		if this.begin_ != 0 {
			err := prpc.Write(buffer, this.begin_)
			if err != nil {
				return err
			}
		}
	}
	// serialize limit_
	{
		if this.limit_ != 0 {
			err := prpc.Write(buffer, this.limit_)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *COM_SearchContext) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize title_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.title_)
		if err != nil {
			return err
		}
	}
	// deserialize ist_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.ist_)
		if err != nil {
			return err
		}
	}
	// deserialize rt_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.rt_)
		if err != nil {
			return err
		}
	}
	// deserialize itemId_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.itemId_)
		if err != nil {
			return err
		}
	}
	// deserialize babyId_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.babyId_)
		if err != nil {
			return err
		}
	}
	// deserialize begin_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.begin_)
		if err != nil {
			return err
		}
	}
	// deserialize limit_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.limit_)
		if err != nil {
			return err
		}
	}
	return nil
}
