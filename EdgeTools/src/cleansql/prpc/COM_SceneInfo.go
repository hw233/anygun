package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type COM_SceneInfo struct {
	sceneId_  int32                        //0
	position_ COM_FPosition                //1
	players_  []COM_ScenePlayerInformation //2
	npcs_     []int32                      //3
}

func (this *COM_SceneInfo) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.sceneId_ != 0)
	mask.WriteBit(true) //position_
	mask.WriteBit(len(this.players_) != 0)
	mask.WriteBit(len(this.npcs_) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize sceneId_
	{
		if this.sceneId_ != 0 {
			err := prpc.Write(buffer, this.sceneId_)
			if err != nil {
				return err
			}
		}
	}
	// serialize position_
	{
		err := this.position_.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	// serialize players_
	if len(this.players_) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.players_)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.players_ {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// serialize npcs_
	if len(this.npcs_) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.npcs_)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.npcs_ {
			err := prpc.Write(buffer, value)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *COM_SceneInfo) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize sceneId_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.sceneId_)
		if err != nil {
			return err
		}
	}
	// deserialize position_
	if mask.ReadBit() {
		err := this.position_.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	// deserialize players_
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.players_ = make([]COM_ScenePlayerInformation, size)
		for i, _ := range this.players_ {
			err := this.players_[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// deserialize npcs_
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.npcs_ = make([]int32, size)
		for i, _ := range this.npcs_ {
			err := prpc.Read(buffer, &this.npcs_[i])
			if err != nil {
				return err
			}
		}
	}
	return nil
}
