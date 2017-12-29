package prpc

import (
	"bytes"
	"errors"
	"suzuki/prpc"
)

type SGE_World2Scene_initDynamicNpcs struct {
	type_ int   //0
	count int32 //1
}
type SGE_World2Scene_refreshDynamicNpcs struct {
	type_ int   //0
	count int32 //1
}
type SGE_World2Scene_finiDynamicNpcs struct {
	type_ int //0
}
type SGE_World2Scene_addDynamicNpcs struct {
	sceneId int32   //0
	npcs    []int32 //1
}
type SGE_World2Scene_delDynamicNpc struct {
	npcId int32 //0
}
type SGE_World2Scene_delDynamicNpc2 struct {
	sceneId int32 //0
	npcId   int32 //1
}
type SGE_World2Scene_openSceneCopy struct {
	instId int32 //0
}
type SGE_World2Scene_closeSceneCopy struct {
	instId int32 //0
}
type SGE_World2SceneStub struct {
	Sender prpc.StubSender
}
type SGE_World2SceneProxy interface {
	initDynamicNpcs(type_ int, count int32) error     // 0
	refreshDynamicNpcs(type_ int, count int32) error  // 1
	finiDynamicNpcs(type_ int) error                  // 2
	addDynamicNpcs(sceneId int32, npcs []int32) error // 3
	delDynamicNpc(npcId int32) error                  // 4
	delDynamicNpc2(sceneId int32, npcId int32) error  // 5
	openSceneCopy(instId int32) error                 // 6
	closeSceneCopy(instId int32) error                // 7
}

func (this *SGE_World2Scene_initDynamicNpcs) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.type_ != 0)
	mask.WriteBit(this.count != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize type_
	{
		if this.type_ != 0 {
			err := prpc.Write(buffer, this.type_)
			if err != nil {
				return err
			}
		}
	}
	// serialize count
	{
		if this.count != 0 {
			err := prpc.Write(buffer, this.count)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *SGE_World2Scene_initDynamicNpcs) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize type_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.type_)
		if err != nil {
			return err
		}
	}
	// deserialize count
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.count)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_World2Scene_refreshDynamicNpcs) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.type_ != 0)
	mask.WriteBit(this.count != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize type_
	{
		if this.type_ != 0 {
			err := prpc.Write(buffer, this.type_)
			if err != nil {
				return err
			}
		}
	}
	// serialize count
	{
		if this.count != 0 {
			err := prpc.Write(buffer, this.count)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *SGE_World2Scene_refreshDynamicNpcs) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize type_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.type_)
		if err != nil {
			return err
		}
	}
	// deserialize count
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.count)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_World2Scene_finiDynamicNpcs) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.type_ != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize type_
	{
		if this.type_ != 0 {
			err := prpc.Write(buffer, this.type_)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *SGE_World2Scene_finiDynamicNpcs) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize type_
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.type_)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_World2Scene_addDynamicNpcs) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.sceneId != 0)
	mask.WriteBit(len(this.npcs) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize sceneId
	{
		if this.sceneId != 0 {
			err := prpc.Write(buffer, this.sceneId)
			if err != nil {
				return err
			}
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
func (this *SGE_World2Scene_addDynamicNpcs) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize sceneId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.sceneId)
		if err != nil {
			return err
		}
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
func (this *SGE_World2Scene_delDynamicNpc) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.npcId != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize npcId
	{
		if this.npcId != 0 {
			err := prpc.Write(buffer, this.npcId)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *SGE_World2Scene_delDynamicNpc) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize npcId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.npcId)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_World2Scene_delDynamicNpc2) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.sceneId != 0)
	mask.WriteBit(this.npcId != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize sceneId
	{
		if this.sceneId != 0 {
			err := prpc.Write(buffer, this.sceneId)
			if err != nil {
				return err
			}
		}
	}
	// serialize npcId
	{
		if this.npcId != 0 {
			err := prpc.Write(buffer, this.npcId)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *SGE_World2Scene_delDynamicNpc2) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize sceneId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.sceneId)
		if err != nil {
			return err
		}
	}
	// deserialize npcId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.npcId)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_World2Scene_openSceneCopy) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.instId != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize instId
	{
		if this.instId != 0 {
			err := prpc.Write(buffer, this.instId)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *SGE_World2Scene_openSceneCopy) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize instId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.instId)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_World2Scene_closeSceneCopy) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.instId != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize instId
	{
		if this.instId != 0 {
			err := prpc.Write(buffer, this.instId)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *SGE_World2Scene_closeSceneCopy) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize instId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.instId)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_World2SceneStub) initDynamicNpcs(type_ int, count int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(0))
	if err != nil {
		return err
	}
	_0 := SGE_World2Scene_initDynamicNpcs{}
	_0.type_ = type_
	_0.count = count
	err = _0.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_World2SceneStub) refreshDynamicNpcs(type_ int, count int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(1))
	if err != nil {
		return err
	}
	_1 := SGE_World2Scene_refreshDynamicNpcs{}
	_1.type_ = type_
	_1.count = count
	err = _1.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_World2SceneStub) finiDynamicNpcs(type_ int) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(2))
	if err != nil {
		return err
	}
	_2 := SGE_World2Scene_finiDynamicNpcs{}
	_2.type_ = type_
	err = _2.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_World2SceneStub) addDynamicNpcs(sceneId int32, npcs []int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(3))
	if err != nil {
		return err
	}
	_3 := SGE_World2Scene_addDynamicNpcs{}
	_3.sceneId = sceneId
	_3.npcs = npcs
	err = _3.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_World2SceneStub) delDynamicNpc(npcId int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(4))
	if err != nil {
		return err
	}
	_4 := SGE_World2Scene_delDynamicNpc{}
	_4.npcId = npcId
	err = _4.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_World2SceneStub) delDynamicNpc2(sceneId int32, npcId int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(5))
	if err != nil {
		return err
	}
	_5 := SGE_World2Scene_delDynamicNpc2{}
	_5.sceneId = sceneId
	_5.npcId = npcId
	err = _5.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_World2SceneStub) openSceneCopy(instId int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(6))
	if err != nil {
		return err
	}
	_6 := SGE_World2Scene_openSceneCopy{}
	_6.instId = instId
	err = _6.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_World2SceneStub) closeSceneCopy(instId int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(7))
	if err != nil {
		return err
	}
	_7 := SGE_World2Scene_closeSceneCopy{}
	_7.instId = instId
	err = _7.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func Bridging_SGE_World2Scene_initDynamicNpcs(buffer *bytes.Buffer, p SGE_World2SceneProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_0 := SGE_World2Scene_initDynamicNpcs{}
	err := _0.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.initDynamicNpcs(_0.type_, _0.count)
}
func Bridging_SGE_World2Scene_refreshDynamicNpcs(buffer *bytes.Buffer, p SGE_World2SceneProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_1 := SGE_World2Scene_refreshDynamicNpcs{}
	err := _1.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.refreshDynamicNpcs(_1.type_, _1.count)
}
func Bridging_SGE_World2Scene_finiDynamicNpcs(buffer *bytes.Buffer, p SGE_World2SceneProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_2 := SGE_World2Scene_finiDynamicNpcs{}
	err := _2.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.finiDynamicNpcs(_2.type_)
}
func Bridging_SGE_World2Scene_addDynamicNpcs(buffer *bytes.Buffer, p SGE_World2SceneProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_3 := SGE_World2Scene_addDynamicNpcs{}
	err := _3.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.addDynamicNpcs(_3.sceneId, _3.npcs)
}
func Bridging_SGE_World2Scene_delDynamicNpc(buffer *bytes.Buffer, p SGE_World2SceneProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_4 := SGE_World2Scene_delDynamicNpc{}
	err := _4.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.delDynamicNpc(_4.npcId)
}
func Bridging_SGE_World2Scene_delDynamicNpc2(buffer *bytes.Buffer, p SGE_World2SceneProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_5 := SGE_World2Scene_delDynamicNpc2{}
	err := _5.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.delDynamicNpc2(_5.sceneId, _5.npcId)
}
func Bridging_SGE_World2Scene_openSceneCopy(buffer *bytes.Buffer, p SGE_World2SceneProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_6 := SGE_World2Scene_openSceneCopy{}
	err := _6.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.openSceneCopy(_6.instId)
}
func Bridging_SGE_World2Scene_closeSceneCopy(buffer *bytes.Buffer, p SGE_World2SceneProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_7 := SGE_World2Scene_closeSceneCopy{}
	err := _7.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.closeSceneCopy(_7.instId)
}
func SGE_World2SceneDispatch(buffer *bytes.Buffer, p SGE_World2SceneProxy) error {
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
		return Bridging_SGE_World2Scene_initDynamicNpcs(buffer, p)
	case 1:
		return Bridging_SGE_World2Scene_refreshDynamicNpcs(buffer, p)
	case 2:
		return Bridging_SGE_World2Scene_finiDynamicNpcs(buffer, p)
	case 3:
		return Bridging_SGE_World2Scene_addDynamicNpcs(buffer, p)
	case 4:
		return Bridging_SGE_World2Scene_delDynamicNpc(buffer, p)
	case 5:
		return Bridging_SGE_World2Scene_delDynamicNpc2(buffer, p)
	case 6:
		return Bridging_SGE_World2Scene_openSceneCopy(buffer, p)
	case 7:
		return Bridging_SGE_World2Scene_closeSceneCopy(buffer, p)
	default:
		return errors.New(prpc.NoneDispatchMatchError)
	}
}
