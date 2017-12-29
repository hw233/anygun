package prpc

import (
	"bytes"
	"errors"
	"suzuki/prpc"
)

type SGE_Mall2World_fetchSellOk struct {
	playerid  int32          //0
	items     []COM_SellItem //1
	totalSize int32          //2
}
type SGE_Mall2World_fetchMySellOk struct {
	playerid int32          //0
	items    []COM_SellItem //1
}
type SGE_Mall2World_fetchSelledItemOk struct {
	playerId int32            //0
	items    []COM_SelledItem //1
}
type SGE_Mall2World_sellOk struct {
	playerid int32        //0
	item     COM_SellItem //1
}
type SGE_Mall2World_unSellOk struct {
	playerid int32 //0
	sellid   int32 //1
}
type SGE_Mall2World_buyOk struct {
	playerid int32        //0
	item     COM_SellItem //1
}
type SGE_Mall2World_buyFail struct {
	playerid int32 //0
	errorno  int   //1
}
type SGE_Mall2WorldStub struct {
	Sender prpc.StubSender
}
type SGE_Mall2WorldProxy interface {
	fetchSellOk(playerid int32, items []COM_SellItem, totalSize int32) error // 0
	fetchMySellOk(playerid int32, items []COM_SellItem) error                // 1
	fetchSelledItemOk(playerId int32, items []COM_SelledItem) error          // 2
	sellOk(playerid int32, item COM_SellItem) error                          // 3
	unSellOk(playerid int32, sellid int32) error                             // 4
	buyOk(playerid int32, item COM_SellItem) error                           // 5
	buyFail(playerid int32, errorno int) error                               // 6
}

func (this *SGE_Mall2World_fetchSellOk) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.playerid != 0)
	mask.WriteBit(len(this.items) != 0)
	mask.WriteBit(this.totalSize != 0)
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
	// serialize items
	if len(this.items) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.items)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.items {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// serialize totalSize
	{
		if this.totalSize != 0 {
			err := prpc.Write(buffer, this.totalSize)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *SGE_Mall2World_fetchSellOk) Deserialize(buffer *bytes.Buffer) error {
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
	// deserialize items
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.items = make([]COM_SellItem, size)
		for i, _ := range this.items {
			err := this.items[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// deserialize totalSize
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.totalSize)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_Mall2World_fetchMySellOk) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.playerid != 0)
	mask.WriteBit(len(this.items) != 0)
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
	// serialize items
	if len(this.items) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.items)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.items {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *SGE_Mall2World_fetchMySellOk) Deserialize(buffer *bytes.Buffer) error {
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
	// deserialize items
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.items = make([]COM_SellItem, size)
		for i, _ := range this.items {
			err := this.items[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *SGE_Mall2World_fetchSelledItemOk) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.playerId != 0)
	mask.WriteBit(len(this.items) != 0)
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
	// serialize items
	if len(this.items) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.items)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.items {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *SGE_Mall2World_fetchSelledItemOk) Deserialize(buffer *bytes.Buffer) error {
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
	// deserialize items
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.items = make([]COM_SelledItem, size)
		for i, _ := range this.items {
			err := this.items[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *SGE_Mall2World_sellOk) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.playerid != 0)
	mask.WriteBit(true) //item
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
	// serialize item
	{
		err := this.item.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_Mall2World_sellOk) Deserialize(buffer *bytes.Buffer) error {
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
	// deserialize item
	if mask.ReadBit() {
		err := this.item.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_Mall2World_unSellOk) Serialize(buffer *bytes.Buffer) error {
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
func (this *SGE_Mall2World_unSellOk) Deserialize(buffer *bytes.Buffer) error {
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
func (this *SGE_Mall2World_buyOk) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.playerid != 0)
	mask.WriteBit(true) //item
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
	// serialize item
	{
		err := this.item.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_Mall2World_buyOk) Deserialize(buffer *bytes.Buffer) error {
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
	// deserialize item
	if mask.ReadBit() {
		err := this.item.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_Mall2World_buyFail) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.playerid != 0)
	mask.WriteBit(this.errorno != 0)
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
	// serialize errorno
	{
		if this.errorno != 0 {
			err := prpc.Write(buffer, this.errorno)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *SGE_Mall2World_buyFail) Deserialize(buffer *bytes.Buffer) error {
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
	// deserialize errorno
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.errorno)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_Mall2WorldStub) fetchSellOk(playerid int32, items []COM_SellItem, totalSize int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(0))
	if err != nil {
		return err
	}
	_0 := SGE_Mall2World_fetchSellOk{}
	_0.playerid = playerid
	_0.items = items
	_0.totalSize = totalSize
	err = _0.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_Mall2WorldStub) fetchMySellOk(playerid int32, items []COM_SellItem) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(1))
	if err != nil {
		return err
	}
	_1 := SGE_Mall2World_fetchMySellOk{}
	_1.playerid = playerid
	_1.items = items
	err = _1.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_Mall2WorldStub) fetchSelledItemOk(playerId int32, items []COM_SelledItem) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(2))
	if err != nil {
		return err
	}
	_2 := SGE_Mall2World_fetchSelledItemOk{}
	_2.playerId = playerId
	_2.items = items
	err = _2.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_Mall2WorldStub) sellOk(playerid int32, item COM_SellItem) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(3))
	if err != nil {
		return err
	}
	_3 := SGE_Mall2World_sellOk{}
	_3.playerid = playerid
	_3.item = item
	err = _3.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_Mall2WorldStub) unSellOk(playerid int32, sellid int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(4))
	if err != nil {
		return err
	}
	_4 := SGE_Mall2World_unSellOk{}
	_4.playerid = playerid
	_4.sellid = sellid
	err = _4.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_Mall2WorldStub) buyOk(playerid int32, item COM_SellItem) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(5))
	if err != nil {
		return err
	}
	_5 := SGE_Mall2World_buyOk{}
	_5.playerid = playerid
	_5.item = item
	err = _5.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_Mall2WorldStub) buyFail(playerid int32, errorno int) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(6))
	if err != nil {
		return err
	}
	_6 := SGE_Mall2World_buyFail{}
	_6.playerid = playerid
	_6.errorno = errorno
	err = _6.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func Bridging_SGE_Mall2World_fetchSellOk(buffer *bytes.Buffer, p SGE_Mall2WorldProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_0 := SGE_Mall2World_fetchSellOk{}
	err := _0.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.fetchSellOk(_0.playerid, _0.items, _0.totalSize)
}
func Bridging_SGE_Mall2World_fetchMySellOk(buffer *bytes.Buffer, p SGE_Mall2WorldProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_1 := SGE_Mall2World_fetchMySellOk{}
	err := _1.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.fetchMySellOk(_1.playerid, _1.items)
}
func Bridging_SGE_Mall2World_fetchSelledItemOk(buffer *bytes.Buffer, p SGE_Mall2WorldProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_2 := SGE_Mall2World_fetchSelledItemOk{}
	err := _2.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.fetchSelledItemOk(_2.playerId, _2.items)
}
func Bridging_SGE_Mall2World_sellOk(buffer *bytes.Buffer, p SGE_Mall2WorldProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_3 := SGE_Mall2World_sellOk{}
	err := _3.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.sellOk(_3.playerid, _3.item)
}
func Bridging_SGE_Mall2World_unSellOk(buffer *bytes.Buffer, p SGE_Mall2WorldProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_4 := SGE_Mall2World_unSellOk{}
	err := _4.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.unSellOk(_4.playerid, _4.sellid)
}
func Bridging_SGE_Mall2World_buyOk(buffer *bytes.Buffer, p SGE_Mall2WorldProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_5 := SGE_Mall2World_buyOk{}
	err := _5.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.buyOk(_5.playerid, _5.item)
}
func Bridging_SGE_Mall2World_buyFail(buffer *bytes.Buffer, p SGE_Mall2WorldProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_6 := SGE_Mall2World_buyFail{}
	err := _6.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.buyFail(_6.playerid, _6.errorno)
}
func SGE_Mall2WorldDispatch(buffer *bytes.Buffer, p SGE_Mall2WorldProxy) error {
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
		return Bridging_SGE_Mall2World_fetchSellOk(buffer, p)
	case 1:
		return Bridging_SGE_Mall2World_fetchMySellOk(buffer, p)
	case 2:
		return Bridging_SGE_Mall2World_fetchSelledItemOk(buffer, p)
	case 3:
		return Bridging_SGE_Mall2World_sellOk(buffer, p)
	case 4:
		return Bridging_SGE_Mall2World_unSellOk(buffer, p)
	case 5:
		return Bridging_SGE_Mall2World_buyOk(buffer, p)
	case 6:
		return Bridging_SGE_Mall2World_buyFail(buffer, p)
	default:
		return errors.New(prpc.NoneDispatchMatchError)
	}
}
