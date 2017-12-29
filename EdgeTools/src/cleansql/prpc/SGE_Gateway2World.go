package prpc

import (
	"bytes"
	"errors"
	"suzuki/prpc"
)

type SGE_Gateway2World_syncConnectInfo struct {
	indoor int32  //0
	ip     string //1
}
type SGE_Gateway2WorldStub struct {
	Sender prpc.StubSender
}
type SGE_Gateway2WorldProxy interface {
	syncConnectInfo(indoor int32, ip string) error // 0
}

func (this *SGE_Gateway2World_syncConnectInfo) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.indoor != 0)
	mask.WriteBit(len(this.ip) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize indoor
	{
		if this.indoor != 0 {
			err := prpc.Write(buffer, this.indoor)
			if err != nil {
				return err
			}
		}
	}
	// serialize ip
	if len(this.ip) != 0 {
		err := prpc.Write(buffer, this.ip)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_Gateway2World_syncConnectInfo) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize indoor
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.indoor)
		if err != nil {
			return err
		}
	}
	// deserialize ip
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.ip)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_Gateway2WorldStub) syncConnectInfo(indoor int32, ip string) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(0))
	if err != nil {
		return err
	}
	_0 := SGE_Gateway2World_syncConnectInfo{}
	_0.indoor = indoor
	_0.ip = ip
	err = _0.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func Bridging_SGE_Gateway2World_syncConnectInfo(buffer *bytes.Buffer, p SGE_Gateway2WorldProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_0 := SGE_Gateway2World_syncConnectInfo{}
	err := _0.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.syncConnectInfo(_0.indoor, _0.ip)
}
func SGE_Gateway2WorldDispatch(buffer *bytes.Buffer, p SGE_Gateway2WorldProxy) error {
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
		return Bridging_SGE_Gateway2World_syncConnectInfo(buffer, p)
	default:
		return errors.New(prpc.NoneDispatchMatchError)
	}
}
