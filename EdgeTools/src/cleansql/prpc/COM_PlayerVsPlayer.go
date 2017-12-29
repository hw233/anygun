package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type COM_PlayerVsPlayer struct {
	playerInst_ int32   //0
	section_    int32   //1
	value_      int32   //2
	winNum_     int32   //3
	battleNum_  int32   //4
	winValue_   float32 //5
	contWin_    int32   //6
	isCont_     bool    //7
}

func (this *COM_PlayerVsPlayer) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.playerInst_ != 0)
	mask.WriteBit(this.section_ != 0)
	mask.WriteBit(this.value_ != 0)
	mask.WriteBit(this.winNum_ != 0)
	mask.WriteBit(this.battleNum_ != 0)
	mask.WriteBit(this.winValue_ != 0)
	mask.WriteBit(this.contWin_ != 0)
	mask.WriteBit(this.isCont_)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize playerInst_
	{
		if this.playerInst_ != 0 {
			err := prpc.Write(buffer, this.playerInst_)
			if err != nil {
				return err
			}
		}
	}
	// serialize section_
	{
		if this.section_ != 0 {
			err := prpc.Write(buffer, this.section_)
			if err != nil {
				return err
			}
		}
	}
	// serialize value_
	{
		if this.value_ != 0 {
			err := prpc.Write(buffer, this.value_)
			if err != nil {
				return err
			}
		}
	}
	// serialize winNum_
	{
		if this.winNum_ != 0 {
			err := prpc.Write(buffer, this.winNum_)
			if err != nil {
				return err
			}
		}
	}
	// serialize battleNum_
	{
		if this.battleNum_ != 0 {
			err := prpc.Write(buffer, this.battleNum_)
			if err != nil {
				return err
			}
		}
	}
	// serialize winValue_
	{
		if this.winValue_ != 0 {
			err := prpc.Write(buffer, this.winValue_)
			if err != nil {
				return err
			}
		}
	}
	// serialize contWin_
	{
		if this.contWin_ != 0 {
			err := prpc.Write(buffer, this.contWin_)
			if err != nil {
				return err
			}
		}
	}
	// serialize isCont_
	{
	}
	return nil
}
func (this *COM_PlayerVsPlayer) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize playerInst_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.playerInst_)
		if err != nil {
			return err
		}
	}
	// deserialize section_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.section_)
		if err != nil {
			return err
		}
	}
	// deserialize value_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.value_)
		if err != nil {
			return err
		}
	}
	// deserialize winNum_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.winNum_)
		if err != nil {
			return err
		}
	}
	// deserialize battleNum_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.battleNum_)
		if err != nil {
			return err
		}
	}
	// deserialize winValue_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.winValue_)
		if err != nil {
			return err
		}
	}
	// deserialize contWin_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.contWin_)
		if err != nil {
			return err
		}
	}
	// deserialize isCont_
	this.isCont_ = mask.ReadBit()
	return nil
}
