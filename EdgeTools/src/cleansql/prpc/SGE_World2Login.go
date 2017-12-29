package prpc

import (
	"bytes"
	"errors"
	"suzuki/prpc"
)

type SGE_World2Login_queryAccount struct {
	info COM_LoginInfo //0
}
type SGE_World2Login_setAccountSeal struct {
	accountname string //0
	val         bool   //1
}
type SGE_World2Login_setPhoneNumber struct {
	accountname string //0
	phoneNumber string //1
}
type SGE_World2LoginStub struct {
	Sender prpc.StubSender
}
type SGE_World2LoginProxy interface {
	queryAccount(info COM_LoginInfo) error                       // 0
	setAccountSeal(accountname string, val bool) error           // 1
	setPhoneNumber(accountname string, phoneNumber string) error // 2
}

func (this *SGE_World2Login_queryAccount) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //info
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
	return nil
}
func (this *SGE_World2Login_queryAccount) Deserialize(buffer *bytes.Buffer) error {
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
	return nil
}
func (this *SGE_World2Login_setAccountSeal) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.accountname) != 0)
	mask.WriteBit(this.val)
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
	// serialize val
	{
	}
	return nil
}
func (this *SGE_World2Login_setAccountSeal) Deserialize(buffer *bytes.Buffer) error {
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
	// deserialize val
	this.val = mask.ReadBit()
	return nil
}
func (this *SGE_World2Login_setPhoneNumber) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.accountname) != 0)
	mask.WriteBit(len(this.phoneNumber) != 0)
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
	// serialize phoneNumber
	if len(this.phoneNumber) != 0 {
		err := prpc.Write(buffer, this.phoneNumber)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_World2Login_setPhoneNumber) Deserialize(buffer *bytes.Buffer) error {
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
	// deserialize phoneNumber
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.phoneNumber)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_World2LoginStub) queryAccount(info COM_LoginInfo) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(0))
	if err != nil {
		return err
	}
	_0 := SGE_World2Login_queryAccount{}
	_0.info = info
	err = _0.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_World2LoginStub) setAccountSeal(accountname string, val bool) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(1))
	if err != nil {
		return err
	}
	_1 := SGE_World2Login_setAccountSeal{}
	_1.accountname = accountname
	_1.val = val
	err = _1.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_World2LoginStub) setPhoneNumber(accountname string, phoneNumber string) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(2))
	if err != nil {
		return err
	}
	_2 := SGE_World2Login_setPhoneNumber{}
	_2.accountname = accountname
	_2.phoneNumber = phoneNumber
	err = _2.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func Bridging_SGE_World2Login_queryAccount(buffer *bytes.Buffer, p SGE_World2LoginProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_0 := SGE_World2Login_queryAccount{}
	err := _0.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.queryAccount(_0.info)
}
func Bridging_SGE_World2Login_setAccountSeal(buffer *bytes.Buffer, p SGE_World2LoginProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_1 := SGE_World2Login_setAccountSeal{}
	err := _1.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.setAccountSeal(_1.accountname, _1.val)
}
func Bridging_SGE_World2Login_setPhoneNumber(buffer *bytes.Buffer, p SGE_World2LoginProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_2 := SGE_World2Login_setPhoneNumber{}
	err := _2.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.setPhoneNumber(_2.accountname, _2.phoneNumber)
}
func SGE_World2LoginDispatch(buffer *bytes.Buffer, p SGE_World2LoginProxy) error {
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
		return Bridging_SGE_World2Login_queryAccount(buffer, p)
	case 1:
		return Bridging_SGE_World2Login_setAccountSeal(buffer, p)
	case 2:
		return Bridging_SGE_World2Login_setPhoneNumber(buffer, p)
	default:
		return errors.New(prpc.NoneDispatchMatchError)
	}
}
