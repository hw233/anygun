package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type COM_ReportAction struct {
	COM_Order
	zhuachongOk_     bool                          //0
	skillLevel_      uint8                         //1
	huweiPosition_   uint8                         //2
	bp0_             int                           //3
	bp1_             int                           //4
	baby_            COM_BattleEntityInformation   //5
	eraseEntities_   []int32                       //6
	targets_         []COM_ReportTarget            //7
	stateIds_        []COM_ReportState             //8
	counters_        []COM_ReportActionCounter     //9
	dynamicEntities_ []COM_BattleEntityInformation //10
}

func (this *COM_ReportAction) Serialize(buffer *bytes.Buffer) error {
	{
		err := this.COM_Order.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	//field mask
	mask := prpc.NewMask1(2)
	mask.WriteBit(this.zhuachongOk_)
	mask.WriteBit(this.skillLevel_ != 0)
	mask.WriteBit(this.huweiPosition_ != 0)
	mask.WriteBit(this.bp0_ != 0)
	mask.WriteBit(this.bp1_ != 0)
	mask.WriteBit(true) //baby_
	mask.WriteBit(len(this.eraseEntities_) != 0)
	mask.WriteBit(len(this.targets_) != 0)
	mask.WriteBit(len(this.stateIds_) != 0)
	mask.WriteBit(len(this.counters_) != 0)
	mask.WriteBit(len(this.dynamicEntities_) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize zhuachongOk_
	{
	}
	// serialize skillLevel_
	{
		if this.skillLevel_ != 0 {
			err := prpc.Write(buffer, this.skillLevel_)
			if err != nil {
				return err
			}
		}
	}
	// serialize huweiPosition_
	{
		if this.huweiPosition_ != 0 {
			err := prpc.Write(buffer, this.huweiPosition_)
			if err != nil {
				return err
			}
		}
	}
	// serialize bp0_
	{
		if this.bp0_ != 0 {
			err := prpc.Write(buffer, this.bp0_)
			if err != nil {
				return err
			}
		}
	}
	// serialize bp1_
	{
		if this.bp1_ != 0 {
			err := prpc.Write(buffer, this.bp1_)
			if err != nil {
				return err
			}
		}
	}
	// serialize baby_
	{
		err := this.baby_.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	// serialize eraseEntities_
	if len(this.eraseEntities_) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.eraseEntities_)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.eraseEntities_ {
			err := prpc.Write(buffer, value)
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
	// serialize counters_
	if len(this.counters_) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.counters_)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.counters_ {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// serialize dynamicEntities_
	if len(this.dynamicEntities_) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.dynamicEntities_)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.dynamicEntities_ {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *COM_ReportAction) Deserialize(buffer *bytes.Buffer) error {
	{
		this.COM_Order.Deserialize(buffer)
	}
	//field mask
	mask, err := prpc.NewMask0(buffer, 2)
	if err != nil {
		return err
	}
	// deserialize zhuachongOk_
	this.zhuachongOk_ = mask.ReadBit()
	// deserialize skillLevel_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.skillLevel_)
		if err != nil {
			return err
		}
	}
	// deserialize huweiPosition_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.huweiPosition_)
		if err != nil {
			return err
		}
	}
	// deserialize bp0_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.bp0_)
		if err != nil {
			return err
		}
	}
	// deserialize bp1_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.bp1_)
		if err != nil {
			return err
		}
	}
	// deserialize baby_
	if mask.ReadBit() {
		err := this.baby_.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	// deserialize eraseEntities_
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.eraseEntities_ = make([]int32, size)
		for i, _ := range this.eraseEntities_ {
			err := prpc.Read(buffer, &this.eraseEntities_[i])
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
	// deserialize counters_
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.counters_ = make([]COM_ReportActionCounter, size)
		for i, _ := range this.counters_ {
			err := this.counters_[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// deserialize dynamicEntities_
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.dynamicEntities_ = make([]COM_BattleEntityInformation, size)
		for i, _ := range this.dynamicEntities_ {
			err := this.dynamicEntities_[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
