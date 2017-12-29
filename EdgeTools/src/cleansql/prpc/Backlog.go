package prpc

import (
	"bytes"
	"errors"
	"suzuki/prpc"
)

type Backlog_log struct {
	key     string //0
	msg     string //1
	stack   string //2
	version string //3
}
type BacklogStub struct {
	Sender prpc.StubSender
}
type BacklogProxy interface {
	log(key string, msg string, stack string, version string) error // 0
}

func (this *Backlog_log) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.key) != 0)
	mask.WriteBit(len(this.msg) != 0)
	mask.WriteBit(len(this.stack) != 0)
	mask.WriteBit(len(this.version) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize key
	if len(this.key) != 0 {
		err := prpc.Write(buffer, this.key)
		if err != nil {
			return err
		}
	}
	// serialize msg
	if len(this.msg) != 0 {
		err := prpc.Write(buffer, this.msg)
		if err != nil {
			return err
		}
	}
	// serialize stack
	if len(this.stack) != 0 {
		err := prpc.Write(buffer, this.stack)
		if err != nil {
			return err
		}
	}
	// serialize version
	if len(this.version) != 0 {
		err := prpc.Write(buffer, this.version)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *Backlog_log) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize key
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.key)
		if err != nil {
			return err
		}
	}
	// deserialize msg
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.msg)
		if err != nil {
			return err
		}
	}
	// deserialize stack
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.stack)
		if err != nil {
			return err
		}
	}
	// deserialize version
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.version)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *BacklogStub) log(key string, msg string, stack string, version string) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(0))
	if err != nil {
		return err
	}
	_0 := Backlog_log{}
	_0.key = key
	_0.msg = msg
	_0.stack = stack
	_0.version = version
	err = _0.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func Bridging_Backlog_log(buffer *bytes.Buffer, p BacklogProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_0 := Backlog_log{}
	err := _0.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.log(_0.key, _0.msg, _0.stack, _0.version)
}
func BacklogDispatch(buffer *bytes.Buffer, p BacklogProxy) error {
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
		return Bridging_Backlog_log(buffer, p)
	default:
		return errors.New(prpc.NoneDispatchMatchError)
	}
}
