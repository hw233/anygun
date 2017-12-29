package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type COM_ReconnectInfo struct {
	reconnectProcess_ int                     //0
	sessionKey_       string                  //1
	roles_            []COM_SimpleInformation //2
	playerInst_       COM_PlayerInst          //3
	sceneInfo_        COM_SceneInfo           //4
	team_             COM_TeamInfo            //5
	initBattle_       COM_InitBattle          //6
}

func (this *COM_ReconnectInfo) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.reconnectProcess_ != 0)
	mask.WriteBit(len(this.sessionKey_) != 0)
	mask.WriteBit(len(this.roles_) != 0)
	mask.WriteBit(true) //playerInst_
	mask.WriteBit(true) //sceneInfo_
	mask.WriteBit(true) //team_
	mask.WriteBit(true) //initBattle_
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize reconnectProcess_
	{
		if this.reconnectProcess_ != 0 {
			err := prpc.Write(buffer, this.reconnectProcess_)
			if err != nil {
				return err
			}
		}
	}
	// serialize sessionKey_
	if len(this.sessionKey_) != 0 {
		err := prpc.Write(buffer, this.sessionKey_)
		if err != nil {
			return err
		}
	}
	// serialize roles_
	if len(this.roles_) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.roles_)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.roles_ {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// serialize playerInst_
	{
		err := this.playerInst_.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	// serialize sceneInfo_
	{
		err := this.sceneInfo_.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	// serialize team_
	{
		err := this.team_.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	// serialize initBattle_
	{
		err := this.initBattle_.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *COM_ReconnectInfo) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize reconnectProcess_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.reconnectProcess_)
		if err != nil {
			return err
		}
	}
	// deserialize sessionKey_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.sessionKey_)
		if err != nil {
			return err
		}
	}
	// deserialize roles_
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.roles_ = make([]COM_SimpleInformation, size)
		for i, _ := range this.roles_ {
			err := this.roles_[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// deserialize playerInst_
	if mask.ReadBit() {
		err := this.playerInst_.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	// deserialize sceneInfo_
	if mask.ReadBit() {
		err := this.sceneInfo_.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	// deserialize team_
	if mask.ReadBit() {
		err := this.team_.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	// deserialize initBattle_
	if mask.ReadBit() {
		err := this.initBattle_.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
