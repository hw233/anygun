package prpc

import (
	"bytes"
	"errors"
	"suzuki/prpc"
)

type SGE_Log_account struct {
	data SGE_Account //0
}
type SGE_Log_login struct {
	data SGE_Login //0
}
type SGE_Log_order struct {
	data SGE_Order //0
}
type SGE_Log_role struct {
	data SGE_LogRole //0
}
type SGE_Log_playersay struct {
	senderId   uint32   //0
	senderName string   //1
	chat       COM_Chat //2
}
type SGE_Log_playerTrack struct {
	data SGE_LogProduceTrack //0
}
type SGE_Log_playerUIBehavior struct {
	core SGE_LogUIBehavior //0
}
type SGE_LogStub struct {
	Sender prpc.StubSender
}
type SGE_LogProxy interface {
	account(data SGE_Account) error                                    // 0
	login(data SGE_Login) error                                        // 1
	order(data SGE_Order) error                                        // 2
	role(data SGE_LogRole) error                                       // 3
	playersay(senderId uint32, senderName string, chat COM_Chat) error // 4
	playerTrack(data SGE_LogProduceTrack) error                        // 5
	playerUIBehavior(core SGE_LogUIBehavior) error                     // 6
}

func (this *SGE_Log_account) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //data
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize data
	{
		err := this.data.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_Log_account) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize data
	if mask.ReadBit() {
		err := this.data.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_Log_login) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //data
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize data
	{
		err := this.data.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_Log_login) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize data
	if mask.ReadBit() {
		err := this.data.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_Log_order) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //data
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize data
	{
		err := this.data.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_Log_order) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize data
	if mask.ReadBit() {
		err := this.data.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_Log_role) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //data
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize data
	{
		err := this.data.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_Log_role) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize data
	if mask.ReadBit() {
		err := this.data.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_Log_playersay) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.senderId != 0)
	mask.WriteBit(len(this.senderName) != 0)
	mask.WriteBit(true) //chat
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize senderId
	{
		if this.senderId != 0 {
			err := prpc.Write(buffer, this.senderId)
			if err != nil {
				return err
			}
		}
	}
	// serialize senderName
	if len(this.senderName) != 0 {
		err := prpc.Write(buffer, this.senderName)
		if err != nil {
			return err
		}
	}
	// serialize chat
	{
		err := this.chat.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_Log_playersay) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize senderId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.senderId)
		if err != nil {
			return err
		}
	}
	// deserialize senderName
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.senderName)
		if err != nil {
			return err
		}
	}
	// deserialize chat
	if mask.ReadBit() {
		err := this.chat.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_Log_playerTrack) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //data
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize data
	{
		err := this.data.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_Log_playerTrack) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize data
	if mask.ReadBit() {
		err := this.data.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_Log_playerUIBehavior) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //core
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize core
	{
		err := this.core.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_Log_playerUIBehavior) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize core
	if mask.ReadBit() {
		err := this.core.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_LogStub) account(data SGE_Account) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(0))
	if err != nil {
		return err
	}
	_0 := SGE_Log_account{}
	_0.data = data
	err = _0.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_LogStub) login(data SGE_Login) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(1))
	if err != nil {
		return err
	}
	_1 := SGE_Log_login{}
	_1.data = data
	err = _1.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_LogStub) order(data SGE_Order) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(2))
	if err != nil {
		return err
	}
	_2 := SGE_Log_order{}
	_2.data = data
	err = _2.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_LogStub) role(data SGE_LogRole) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(3))
	if err != nil {
		return err
	}
	_3 := SGE_Log_role{}
	_3.data = data
	err = _3.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_LogStub) playersay(senderId uint32, senderName string, chat COM_Chat) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(4))
	if err != nil {
		return err
	}
	_4 := SGE_Log_playersay{}
	_4.senderId = senderId
	_4.senderName = senderName
	_4.chat = chat
	err = _4.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_LogStub) playerTrack(data SGE_LogProduceTrack) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(5))
	if err != nil {
		return err
	}
	_5 := SGE_Log_playerTrack{}
	_5.data = data
	err = _5.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_LogStub) playerUIBehavior(core SGE_LogUIBehavior) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(6))
	if err != nil {
		return err
	}
	_6 := SGE_Log_playerUIBehavior{}
	_6.core = core
	err = _6.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func Bridging_SGE_Log_account(buffer *bytes.Buffer, p SGE_LogProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_0 := SGE_Log_account{}
	err := _0.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.account(_0.data)
}
func Bridging_SGE_Log_login(buffer *bytes.Buffer, p SGE_LogProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_1 := SGE_Log_login{}
	err := _1.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.login(_1.data)
}
func Bridging_SGE_Log_order(buffer *bytes.Buffer, p SGE_LogProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_2 := SGE_Log_order{}
	err := _2.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.order(_2.data)
}
func Bridging_SGE_Log_role(buffer *bytes.Buffer, p SGE_LogProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_3 := SGE_Log_role{}
	err := _3.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.role(_3.data)
}
func Bridging_SGE_Log_playersay(buffer *bytes.Buffer, p SGE_LogProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_4 := SGE_Log_playersay{}
	err := _4.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.playersay(_4.senderId, _4.senderName, _4.chat)
}
func Bridging_SGE_Log_playerTrack(buffer *bytes.Buffer, p SGE_LogProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_5 := SGE_Log_playerTrack{}
	err := _5.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.playerTrack(_5.data)
}
func Bridging_SGE_Log_playerUIBehavior(buffer *bytes.Buffer, p SGE_LogProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_6 := SGE_Log_playerUIBehavior{}
	err := _6.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.playerUIBehavior(_6.core)
}
func SGE_LogDispatch(buffer *bytes.Buffer, p SGE_LogProxy) error {
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
		return Bridging_SGE_Log_account(buffer, p)
	case 1:
		return Bridging_SGE_Log_login(buffer, p)
	case 2:
		return Bridging_SGE_Log_order(buffer, p)
	case 3:
		return Bridging_SGE_Log_role(buffer, p)
	case 4:
		return Bridging_SGE_Log_playersay(buffer, p)
	case 5:
		return Bridging_SGE_Log_playerTrack(buffer, p)
	case 6:
		return Bridging_SGE_Log_playerUIBehavior(buffer, p)
	default:
		return errors.New(prpc.NoneDispatchMatchError)
	}
}
