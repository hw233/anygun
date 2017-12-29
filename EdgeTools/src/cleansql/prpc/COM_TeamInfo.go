package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type COM_TeamInfo struct {
	COM_SimpleTeamInfo
	members_ []COM_SimplePlayerInst //0
}

func (this *COM_TeamInfo) Serialize(buffer *bytes.Buffer) error {
	{
		err := this.COM_SimpleTeamInfo.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.members_) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize members_
	if len(this.members_) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.members_)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.members_ {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *COM_TeamInfo) Deserialize(buffer *bytes.Buffer) error {
	{
		this.COM_SimpleTeamInfo.Deserialize(buffer)
	}
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize members_
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.members_ = make([]COM_SimplePlayerInst, size)
		for i, _ := range this.members_ {
			err := this.members_[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
