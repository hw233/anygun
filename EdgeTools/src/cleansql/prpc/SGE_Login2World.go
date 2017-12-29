package prpc

import (
	"bytes"
	"errors"
	"suzuki/prpc"
)

type SGE_Login2World_queryAccountOk struct {
	info   COM_AccountInfo //0
	isNew  bool            //1
	isSeal bool            //2
}
type SGE_Login2World_setAccountSealOk struct {
	accountname string //0
}
type SGE_Login2WorldStub struct {
	Sender prpc.StubSender
}
type SGE_Login2WorldProxy interface {
	queryAccountOk(info COM_AccountInfo, isNew bool, isSeal bool) error // 0
	setAccountSealOk(accountname string) error                          // 1
}

func (this *SGE_Login2World_queryAccountOk) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //info
	mask.WriteBit(this.isNew)
	mask.WriteBit(this.isSeal)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize info
	{
		err := this.info.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	// serialize isNew
	{
	}
	// serialize isSeal
	{
	}
	return nil
}
func (this *SGE_Login2World_queryAccountOk) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize info
	if mask.ReadBit() {
		err := this.info.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	// deserialize isNew
	this.isNew = mask.ReadBit()
	// deserialize isSeal
	this.isSeal = mask.ReadBit()
	return nil
}
func (this *SGE_Login2World_setAccountSealOk) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.accountname) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize accountname
	if len(this.accountname) != 0 {
		err := prpc.Write(buffer, this.accountname)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_Login2World_setAccountSealOk) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize accountname
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.accountname)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_Login2WorldStub) queryAccountOk(info COM_AccountInfo, isNew bool, isSeal bool) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(0))
	if err != nil {
		return err
	}
	_0 := SGE_Login2World_queryAccountOk{}
	_0.info = info
	_0.isNew = isNew
	_0.isSeal = isSeal
	err = _0.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_Login2WorldStub) setAccountSealOk(accountname string) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(1))
	if err != nil {
		return err
	}
	_1 := SGE_Login2World_setAccountSealOk{}
	_1.accountname = accountname
	err = _1.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func Bridging_SGE_Login2World_queryAccountOk(buffer *bytes.Buffer, p SGE_Login2WorldProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_0 := SGE_Login2World_queryAccountOk{}
	err := _0.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.queryAccountOk(_0.info, _0.isNew, _0.isSeal)
}
func Bridging_SGE_Login2World_setAccountSealOk(buffer *bytes.Buffer, p SGE_Login2WorldProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_1 := SGE_Login2World_setAccountSealOk{}
	err := _1.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.setAccountSealOk(_1.accountname)
}
func SGE_Login2WorldDispatch(buffer *bytes.Buffer, p SGE_Login2WorldProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	pid := uint16(0XFFFF)
	err := prpc.Read(buffer, &pid)
	if err != nil {
		return err
	}
	switch pid {
	case 0:
		return Bridging_SGE_Login2World_queryAccountOk(buffer, p)
	case 1:
		return Bridging_SGE_Login2World_setAccountSealOk(buffer, p)
	default:
		return errors.New(prpc.NoneDispatchMatchError)
	}
}
