package prpc

import (
	"bytes"
	"errors"
	"suzuki/prpc"
)

type SGE_Player_Scene2World_joinScene struct {
	info COM_SceneInfo //0
}
type SGE_Player_Scene2World_move2 struct {
	pos COM_FPosition //0
}
type SGE_Player_Scene2World_transfor2 struct {
	pos COM_FPosition //0
}
type SGE_Player_Scene2World_autoBattleResult struct {
	isOk bool //0
}
type SGE_Player_Scene2World_zoneJoinBattle struct {
	zoneId int32 //0
}
type SGE_Player_Scene2World_playerAddNpc struct {
	npcs []int32 //0
}
type SGE_Player_Scene2World_playerDelNpc struct {
	npcs []int32 //0
}
type SGE_Player_Scene2World_talkedNpc struct {
	npcid int32 //0
}
type SGE_Player_Scene2World_findDynamicNpcOK struct {
	npcid  int32 //0
	hasnpc bool  //1
}
type SGE_Player_Scene2WorldStub struct {
	Sender prpc.StubSender
}
type SGE_Player_Scene2WorldProxy interface {
	joinScene(info COM_SceneInfo) error              // 0
	move2(pos COM_FPosition) error                   // 1
	cantMove() error                                 // 2
	transfor2(pos COM_FPosition) error               // 3
	autoBattleResult(isOk bool) error                // 4
	zoneJoinBattle(zoneId int32) error               // 5
	playerAddNpc(npcs []int32) error                 // 6
	playerDelNpc(npcs []int32) error                 // 7
	talkedNpc(npcid int32) error                     // 8
	findDynamicNpcOK(npcid int32, hasnpc bool) error // 9
}

func (this *SGE_Player_Scene2World_joinScene) Serialize(buffer *bytes.Buffer) error {
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
func (this *SGE_Player_Scene2World_joinScene) Deserialize(buffer *bytes.Buffer) error {
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
func (this *SGE_Player_Scene2World_move2) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //pos
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize pos
	{
		err := this.pos.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_Player_Scene2World_move2) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize pos
	if mask.ReadBit() {
		err := this.pos.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_Player_Scene2World_transfor2) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //pos
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize pos
	{
		err := this.pos.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_Player_Scene2World_transfor2) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize pos
	if mask.ReadBit() {
		err := this.pos.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_Player_Scene2World_autoBattleResult) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.isOk)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize isOk
	{
	}
	return nil
}
func (this *SGE_Player_Scene2World_autoBattleResult) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize isOk
	this.isOk = mask.ReadBit()
	return nil
}
func (this *SGE_Player_Scene2World_zoneJoinBattle) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.zoneId != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize zoneId
	{
		if this.zoneId != 0 {
			err := prpc.Write(buffer, this.zoneId)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *SGE_Player_Scene2World_zoneJoinBattle) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize zoneId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.zoneId)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_Player_Scene2World_playerAddNpc) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.npcs) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize npcs
	if len(this.npcs) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.npcs)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.npcs {
			err := prpc.Write(buffer, value)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *SGE_Player_Scene2World_playerAddNpc) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize npcs
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.npcs = make([]int32, size)
		for i, _ := range this.npcs {
			err := prpc.Read(buffer, &this.npcs[i])
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *SGE_Player_Scene2World_playerDelNpc) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.npcs) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize npcs
	if len(this.npcs) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.npcs)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.npcs {
			err := prpc.Write(buffer, value)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *SGE_Player_Scene2World_playerDelNpc) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize npcs
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.npcs = make([]int32, size)
		for i, _ := range this.npcs {
			err := prpc.Read(buffer, &this.npcs[i])
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *SGE_Player_Scene2World_talkedNpc) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.npcid != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize npcid
	{
		if this.npcid != 0 {
			err := prpc.Write(buffer, this.npcid)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *SGE_Player_Scene2World_talkedNpc) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize npcid
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.npcid)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_Player_Scene2World_findDynamicNpcOK) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.npcid != 0)
	mask.WriteBit(this.hasnpc)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize npcid
	{
		if this.npcid != 0 {
			err := prpc.Write(buffer, this.npcid)
			if err != nil {
				return err
			}
		}
	}
	// serialize hasnpc
	{
	}
	return nil
}
func (this *SGE_Player_Scene2World_findDynamicNpcOK) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize npcid
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.npcid)
		if err != nil {
			return err
		}
	}
	// deserialize hasnpc
	this.hasnpc = mask.ReadBit()
	return nil
}
func (this *SGE_Player_Scene2WorldStub) joinScene(info COM_SceneInfo) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(0))
	if err != nil {
		return err
	}
	_0 := SGE_Player_Scene2World_joinScene{}
	_0.info = info
	err = _0.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_Player_Scene2WorldStub) move2(pos COM_FPosition) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(1))
	if err != nil {
		return err
	}
	_1 := SGE_Player_Scene2World_move2{}
	_1.pos = pos
	err = _1.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_Player_Scene2WorldStub) cantMove() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(2))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_Player_Scene2WorldStub) transfor2(pos COM_FPosition) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(3))
	if err != nil {
		return err
	}
	_3 := SGE_Player_Scene2World_transfor2{}
	_3.pos = pos
	err = _3.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_Player_Scene2WorldStub) autoBattleResult(isOk bool) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(4))
	if err != nil {
		return err
	}
	_4 := SGE_Player_Scene2World_autoBattleResult{}
	_4.isOk = isOk
	err = _4.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_Player_Scene2WorldStub) zoneJoinBattle(zoneId int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(5))
	if err != nil {
		return err
	}
	_5 := SGE_Player_Scene2World_zoneJoinBattle{}
	_5.zoneId = zoneId
	err = _5.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_Player_Scene2WorldStub) playerAddNpc(npcs []int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(6))
	if err != nil {
		return err
	}
	_6 := SGE_Player_Scene2World_playerAddNpc{}
	_6.npcs = npcs
	err = _6.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_Player_Scene2WorldStub) playerDelNpc(npcs []int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(7))
	if err != nil {
		return err
	}
	_7 := SGE_Player_Scene2World_playerDelNpc{}
	_7.npcs = npcs
	err = _7.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_Player_Scene2WorldStub) talkedNpc(npcid int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(8))
	if err != nil {
		return err
	}
	_8 := SGE_Player_Scene2World_talkedNpc{}
	_8.npcid = npcid
	err = _8.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_Player_Scene2WorldStub) findDynamicNpcOK(npcid int32, hasnpc bool) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(9))
	if err != nil {
		return err
	}
	_9 := SGE_Player_Scene2World_findDynamicNpcOK{}
	_9.npcid = npcid
	_9.hasnpc = hasnpc
	err = _9.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func Bridging_SGE_Player_Scene2World_joinScene(buffer *bytes.Buffer, p SGE_Player_Scene2WorldProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_0 := SGE_Player_Scene2World_joinScene{}
	err := _0.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.joinScene(_0.info)
}
func Bridging_SGE_Player_Scene2World_move2(buffer *bytes.Buffer, p SGE_Player_Scene2WorldProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_1 := SGE_Player_Scene2World_move2{}
	err := _1.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.move2(_1.pos)
}
func Bridging_SGE_Player_Scene2World_cantMove(buffer *bytes.Buffer, p SGE_Player_Scene2WorldProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.cantMove()
}
func Bridging_SGE_Player_Scene2World_transfor2(buffer *bytes.Buffer, p SGE_Player_Scene2WorldProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_3 := SGE_Player_Scene2World_transfor2{}
	err := _3.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.transfor2(_3.pos)
}
func Bridging_SGE_Player_Scene2World_autoBattleResult(buffer *bytes.Buffer, p SGE_Player_Scene2WorldProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_4 := SGE_Player_Scene2World_autoBattleResult{}
	err := _4.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.autoBattleResult(_4.isOk)
}
func Bridging_SGE_Player_Scene2World_zoneJoinBattle(buffer *bytes.Buffer, p SGE_Player_Scene2WorldProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_5 := SGE_Player_Scene2World_zoneJoinBattle{}
	err := _5.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.zoneJoinBattle(_5.zoneId)
}
func Bridging_SGE_Player_Scene2World_playerAddNpc(buffer *bytes.Buffer, p SGE_Player_Scene2WorldProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_6 := SGE_Player_Scene2World_playerAddNpc{}
	err := _6.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.playerAddNpc(_6.npcs)
}
func Bridging_SGE_Player_Scene2World_playerDelNpc(buffer *bytes.Buffer, p SGE_Player_Scene2WorldProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_7 := SGE_Player_Scene2World_playerDelNpc{}
	err := _7.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.playerDelNpc(_7.npcs)
}
func Bridging_SGE_Player_Scene2World_talkedNpc(buffer *bytes.Buffer, p SGE_Player_Scene2WorldProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_8 := SGE_Player_Scene2World_talkedNpc{}
	err := _8.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.talkedNpc(_8.npcid)
}
func Bridging_SGE_Player_Scene2World_findDynamicNpcOK(buffer *bytes.Buffer, p SGE_Player_Scene2WorldProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_9 := SGE_Player_Scene2World_findDynamicNpcOK{}
	err := _9.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.findDynamicNpcOK(_9.npcid, _9.hasnpc)
}
func SGE_Player_Scene2WorldDispatch(buffer *bytes.Buffer, p SGE_Player_Scene2WorldProxy) error {
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
		return Bridging_SGE_Player_Scene2World_joinScene(buffer, p)
	case 1:
		return Bridging_SGE_Player_Scene2World_move2(buffer, p)
	case 2:
		return Bridging_SGE_Player_Scene2World_cantMove(buffer, p)
	case 3:
		return Bridging_SGE_Player_Scene2World_transfor2(buffer, p)
	case 4:
		return Bridging_SGE_Player_Scene2World_autoBattleResult(buffer, p)
	case 5:
		return Bridging_SGE_Player_Scene2World_zoneJoinBattle(buffer, p)
	case 6:
		return Bridging_SGE_Player_Scene2World_playerAddNpc(buffer, p)
	case 7:
		return Bridging_SGE_Player_Scene2World_playerDelNpc(buffer, p)
	case 8:
		return Bridging_SGE_Player_Scene2World_talkedNpc(buffer, p)
	case 9:
		return Bridging_SGE_Player_Scene2World_findDynamicNpcOK(buffer, p)
	default:
		return errors.New(prpc.NoneDispatchMatchError)
	}
}
