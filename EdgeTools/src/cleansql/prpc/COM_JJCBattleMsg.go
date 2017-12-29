package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type COM_JJCBattleMsg struct {
	defier_   string //0
	bydefier_ string //1
	rank_     uint32 //2
	isWin_    bool   //3
	curTime_  int64  //4
}

func (this *COM_JJCBattleMsg) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.defier_) != 0)
	mask.WriteBit(len(this.bydefier_) != 0)
	mask.WriteBit(this.rank_ != 0)
	mask.WriteBit(this.isWin_)
	mask.WriteBit(this.curTime_ != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize defier_
	if len(this.defier_) != 0 {
		err := prpc.Write(buffer, this.defier_)
		if err != nil {
			return err
		}
	}
	// serialize bydefier_
	if len(this.bydefier_) != 0 {
		err := prpc.Write(buffer, this.bydefier_)
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
	// serialize isWin_
	{
	}
	// serialize curTime_
	{
		if this.curTime_ != 0 {
			err := prpc.Write(buffer, this.curTime_)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *COM_JJCBattleMsg) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize defier_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.defier_)
		if err != nil {
			return err
		}
	}
	// deserialize bydefier_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.bydefier_)
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
	// deserialize isWin_
	this.isWin_ = mask.ReadBit()
	// deserialize curTime_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.curTime_)
		if err != nil {
			return err
		}
	}
	return nil
}
