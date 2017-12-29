package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type COM_BattleReport struct {
	stateIds_      []COM_ReportState             //0
	actionResults_ []COM_ReportAction            //1
	targets_       []COM_ReportTarget            //2
	waveEntities_  []COM_BattleEntityInformation //3
}

func (this *COM_BattleReport) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.stateIds_) != 0)
	mask.WriteBit(len(this.actionResults_) != 0)
	mask.WriteBit(len(this.targets_) != 0)
	mask.WriteBit(len(this.waveEntities_) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize stateIds_
	if len(this.stateIds_) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.stateIds_)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.stateIds_ {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// serialize actionResults_
	if len(this.actionResults_) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.actionResults_)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.actionResults_ {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// serialize targets_
	if len(this.targets_) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.targets_)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.targets_ {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// serialize waveEntities_
	if len(this.waveEntities_) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.waveEntities_)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.waveEntities_ {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *COM_BattleReport) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize stateIds_
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.stateIds_ = make([]COM_ReportState, size)
		for i, _ := range this.stateIds_ {
			err := this.stateIds_[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// deserialize actionResults_
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.actionResults_ = make([]COM_ReportAction, size)
		for i, _ := range this.actionResults_ {
			err := this.actionResults_[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// deserialize targets_
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.targets_ = make([]COM_ReportTarget, size)
		for i, _ := range this.targets_ {
			err := this.targets_[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// deserialize waveEntities_
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.waveEntities_ = make([]COM_BattleEntityInformation, size)
		for i, _ := range this.waveEntities_ {
			err := this.waveEntities_[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
