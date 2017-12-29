package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type SGE_ScenePlayerInfo struct {
	playerId_        int32   //0
	playerLevel_     int32   //1
	sceneId_         int32   //2
	entryId_         int32   //3
	currentQuestIds_ []int32 //4
	accecptQuestIds_ []int32 //5
	openScenes_      []int32 //6
}

func (this *SGE_ScenePlayerInfo) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.playerId_ != 0)
	mask.WriteBit(this.playerLevel_ != 0)
	mask.WriteBit(this.sceneId_ != 0)
	mask.WriteBit(this.entryId_ != 0)
	mask.WriteBit(len(this.currentQuestIds_) != 0)
	mask.WriteBit(len(this.accecptQuestIds_) != 0)
	mask.WriteBit(len(this.openScenes_) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize playerId_
	{
		if this.playerId_ != 0 {
			err := prpc.Write(buffer, this.playerId_)
			if err != nil {
				return err
			}
		}
	}
	// serialize playerLevel_
	{
		if this.playerLevel_ != 0 {
			err := prpc.Write(buffer, this.playerLevel_)
			if err != nil {
				return err
			}
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
	// serialize entryId_
	{
		if this.entryId_ != 0 {
			err := prpc.Write(buffer, this.entryId_)
			if err != nil {
				return err
			}
		}
	}
	// serialize currentQuestIds_
	if len(this.currentQuestIds_) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.currentQuestIds_)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.currentQuestIds_ {
			err := prpc.Write(buffer, value)
			if err != nil {
				return err
			}
		}
	}
	// serialize accecptQuestIds_
	if len(this.accecptQuestIds_) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.accecptQuestIds_)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.accecptQuestIds_ {
			err := prpc.Write(buffer, value)
			if err != nil {
				return err
			}
		}
	}
	// serialize openScenes_
	if len(this.openScenes_) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.openScenes_)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.openScenes_ {
			err := prpc.Write(buffer, value)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *SGE_ScenePlayerInfo) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize playerId_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.playerId_)
		if err != nil {
			return err
		}
	}
	// deserialize playerLevel_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.playerLevel_)
		if err != nil {
			return err
		}
	}
	// deserialize sceneId_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.sceneId_)
		if err != nil {
			return err
		}
	}
	// deserialize entryId_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.entryId_)
		if err != nil {
			return err
		}
	}
	// deserialize currentQuestIds_
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.currentQuestIds_ = make([]int32, size)
		for i, _ := range this.currentQuestIds_ {
			err := prpc.Read(buffer, &this.currentQuestIds_[i])
			if err != nil {
				return err
			}
		}
	}
	// deserialize accecptQuestIds_
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.accecptQuestIds_ = make([]int32, size)
		for i, _ := range this.accecptQuestIds_ {
			err := prpc.Read(buffer, &this.accecptQuestIds_[i])
			if err != nil {
				return err
			}
		}
	}
	// deserialize openScenes_
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.openScenes_ = make([]int32, size)
		for i, _ := range this.openScenes_ {
			err := prpc.Read(buffer, &this.openScenes_[i])
			if err != nil {
				return err
			}
		}
	}
	return nil
}
