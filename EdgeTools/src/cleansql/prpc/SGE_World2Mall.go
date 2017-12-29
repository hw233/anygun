package prpc

import (
	"bytes"
	"errors"
	"suzuki/prpc"
)

type SGE_World2Mall_fetchSell struct {
	playerid int32             //0
	context  COM_SearchContext //1
}
type SGE_World2Mall_fetchMySell struct {
	playerid int32 //0
}
type SGE_World2Mall_fetchSelledItem struct {
	playerId int32 //0
}
type SGE_World2Mall_sell struct {
	item COM_SellItem //0
}
type SGE_World2Mall_unSell struct {
	playerid int32 //0
	sellid   int32 //1
}
type SGE_World2Mall_buy struct {
	content SGE_BuyContent //0
}
type SGE_World2Mall_insertSelledItem struct {
	item COM_SelledItem //0
}
type SGE_World2MallStub struct {
	Sender prpc.StubSender
}
type SGE_World2MallProxy interface {
	fetchSell(playerid int32, context COM_SearchContext) error // 0
	fetchMySell(playerid int32) error                          // 1
	fetchSelledItem(playerId int32) error                      // 2
	sell(item COM_SellItem) error                              // 3
	unSell(playerid int32, sellid int32) error                 // 4
	buy(content SGE_BuyContent) error                          // 5
	insertSelledItem(item COM_SelledItem) error                // 6
}

func (this *SGE_World2Mall_fetchSell) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.playerid != 0)
	mask.WriteBit(true) //context
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize playerid
	{
		if this.playerid != 0 {
			err := prpc.Write(buffer, this.playerid)
			if err != nil {
				return err
			}
		}
	}
	// serialize context
	{
		err := this.context.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_World2Mall_fetchSell) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize playerid
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.playerid)
		if err != nil {
			return err
		}
	}
	// deserialize context
	if mask.ReadBit() {
		err := this.context.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_World2Mall_fetchMySell) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.playerid != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize playerid
	{
		if this.playerid != 0 {
			err := prpc.Write(buffer, this.playerid)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *SGE_World2Mall_fetchMySell) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize playerid
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.playerid)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_World2Mall_fetchSelledItem) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.playerId != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize playerId
	{
		if this.playerId != 0 {
			err := prpc.Write(buffer, this.playerId)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *SGE_World2Mall_fetchSelledItem) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize playerId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.playerId)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_World2Mall_sell) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //item
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize item
	{
		err := this.item.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_World2Mall_sell) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize item
	if mask.ReadBit() {
		err := this.item.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_World2Mall_unSell) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.playerid != 0)
	mask.WriteBit(this.sellid != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize playerid
	{
		if this.playerid != 0 {
			err := prpc.Write(buffer, this.playerid)
			if err != nil {
				return err
			}
		}
	}
	// serialize sellid
	{
		if this.sellid != 0 {
			err := prpc.Write(buffer, this.sellid)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *SGE_World2Mall_unSell) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize playerid
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.playerid)
		if err != nil {
			return err
		}
	}
	// deserialize sellid
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.sellid)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_World2Mall_buy) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //content
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize content
	{
		err := this.content.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_World2Mall_buy) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize content
	if mask.ReadBit() {
		err := this.content.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_World2Mall_insertSelledItem) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //item
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize item
	{
		err := this.item.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_World2Mall_insertSelledItem) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize item
	if mask.ReadBit() {
		err := this.item.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_World2MallStub) fetchSell(playerid int32, context COM_SearchContext) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(0))
	if err != nil {
		return err
	}
	_0 := SGE_World2Mall_fetchSell{}
	_0.playerid = playerid
	_0.context = context
	err = _0.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_World2MallStub) fetchMySell(playerid int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(1))
	if err != nil {
		return err
	}
	_1 := SGE_World2Mall_fetchMySell{}
	_1.playerid = playerid
	err = _1.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_World2MallStub) fetchSelledItem(playerId int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(2))
	if err != nil {
		return err
	}
	_2 := SGE_World2Mall_fetchSelledItem{}
	_2.playerId = playerId
	err = _2.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_World2MallStub) sell(item COM_SellItem) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(3))
	if err != nil {
		return err
	}
	_3 := SGE_World2Mall_sell{}
	_3.item = item
	err = _3.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_World2MallStub) unSell(playerid int32, sellid int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(4))
	if err != nil {
		return err
	}
	_4 := SGE_World2Mall_unSell{}
	_4.playerid = playerid
	_4.sellid = sellid
	err = _4.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_World2MallStub) buy(content SGE_BuyContent) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(5))
	if err != nil {
		return err
	}
	_5 := SGE_World2Mall_buy{}
	_5.content = content
	err = _5.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_World2MallStub) insertSelledItem(item COM_SelledItem) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(6))
	if err != nil {
		return err
	}
	_6 := SGE_World2Mall_insertSelledItem{}
	_6.item = item
	err = _6.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func Bridging_SGE_World2Mall_fetchSell(buffer *bytes.Buffer, p SGE_World2MallProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_0 := SGE_World2Mall_fetchSell{}
	err := _0.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.fetchSell(_0.playerid, _0.context)
}
func Bridging_SGE_World2Mall_fetchMySell(buffer *bytes.Buffer, p SGE_World2MallProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_1 := SGE_World2Mall_fetchMySell{}
	err := _1.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.fetchMySell(_1.playerid)
}
func Bridging_SGE_World2Mall_fetchSelledItem(buffer *bytes.Buffer, p SGE_World2MallProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_2 := SGE_World2Mall_fetchSelledItem{}
	err := _2.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.fetchSelledItem(_2.playerId)
}
func Bridging_SGE_World2Mall_sell(buffer *bytes.Buffer, p SGE_World2MallProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_3 := SGE_World2Mall_sell{}
	err := _3.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.sell(_3.item)
}
func Bridging_SGE_World2Mall_unSell(buffer *bytes.Buffer, p SGE_World2MallProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_4 := SGE_World2Mall_unSell{}
	err := _4.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.unSell(_4.playerid, _4.sellid)
}
func Bridging_SGE_World2Mall_buy(buffer *bytes.Buffer, p SGE_World2MallProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_5 := SGE_World2Mall_buy{}
	err := _5.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.buy(_5.content)
}
func Bridging_SGE_World2Mall_insertSelledItem(buffer *bytes.Buffer, p SGE_World2MallProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_6 := SGE_World2Mall_insertSelledItem{}
	err := _6.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.insertSelledItem(_6.item)
}
func SGE_World2MallDispatch(buffer *bytes.Buffer, p SGE_World2MallProxy) error {
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
		return Bridging_SGE_World2Mall_fetchSell(buffer, p)
	case 1:
		return Bridging_SGE_World2Mall_fetchMySell(buffer, p)
	case 2:
		return Bridging_SGE_World2Mall_fetchSelledItem(buffer, p)
	case 3:
		return Bridging_SGE_World2Mall_sell(buffer, p)
	case 4:
		return Bridging_SGE_World2Mall_unSell(buffer, p)
	case 5:
		return Bridging_SGE_World2Mall_buy(buffer, p)
	case 6:
		return Bridging_SGE_World2Mall_insertSelledItem(buffer, p)
	default:
		return errors.New(prpc.NoneDispatchMatchError)
	}
}
