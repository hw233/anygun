package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type SGE_SysActivity struct {
	loginData_ COM_ADLoginTotal    //0
	chData_    COM_ADChargeTotal   //1
	stData_    COM_ADDiscountStore //2
	ceData_    COM_ADChargeEvery   //3
	acData_    COM_ADCards         //4
	hrData_    COM_ADHotRole       //5
	etdata_    COM_ADEmployeeTotal //6
	gbdata_    COM_ADGiftBag       //7
	zpdata_    COM_ZhuanpanData    //8
}

func (this *SGE_SysActivity) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(2)
	mask.WriteBit(true) //loginData_
	mask.WriteBit(true) //chData_
	mask.WriteBit(true) //stData_
	mask.WriteBit(true) //ceData_
	mask.WriteBit(true) //acData_
	mask.WriteBit(true) //hrData_
	mask.WriteBit(true) //etdata_
	mask.WriteBit(true) //gbdata_
	mask.WriteBit(true) //zpdata_
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize loginData_
	{
		err := this.loginData_.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	// serialize chData_
	{
		err := this.chData_.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	// serialize stData_
	{
		err := this.stData_.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	// serialize ceData_
	{
		err := this.ceData_.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	// serialize acData_
	{
		err := this.acData_.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	// serialize hrData_
	{
		err := this.hrData_.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	// serialize etdata_
	{
		err := this.etdata_.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	// serialize gbdata_
	{
		err := this.gbdata_.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	// serialize zpdata_
	{
		err := this.zpdata_.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_SysActivity) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 2)
	if err != nil {
		return err
	}
	// deserialize loginData_
	if mask.ReadBit() {
		err := this.loginData_.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	// deserialize chData_
	if mask.ReadBit() {
		err := this.chData_.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	// deserialize stData_
	if mask.ReadBit() {
		err := this.stData_.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	// deserialize ceData_
	if mask.ReadBit() {
		err := this.ceData_.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	// deserialize acData_
	if mask.ReadBit() {
		err := this.acData_.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	// deserialize hrData_
	if mask.ReadBit() {
		err := this.hrData_.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	// deserialize etdata_
	if mask.ReadBit() {
		err := this.etdata_.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	// deserialize gbdata_
	if mask.ReadBit() {
		err := this.gbdata_.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	// deserialize zpdata_
	if mask.ReadBit() {
		err := this.zpdata_.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
