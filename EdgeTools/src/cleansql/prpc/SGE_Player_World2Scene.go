package prpc

import (
	"bytes"
	"errors"
	"suzuki/prpc"
)

type SGE_Player_World2Scene_initScenePlayer struct {
	info SGE_ScenePlayerInfo //0
}
type SGE_Player_World2Scene_scenePlayerUpLevel struct {
	level int32 //0
}
type SGE_Player_World2Scene_scenePlayerAddCurrentQuest struct {
	questId int32 //0
}
type SGE_Player_World2Scene_scenePlayerDelCurrentQuest struct {
	questId int32 //0
}
type SGE_Player_World2Scene_scenePlayerAddAcceptableQuest struct {
	questId int32 //0
}
type SGE_Player_World2Scene_scenePlayerDelAcceptableQuest struct {
	questId int32 //0
}
type SGE_Player_World2Scene_openScene struct {
	sceneId int32 //0
}
type SGE_Player_World2Scene_transforScene struct {
	sceneId int32 //0
}
type SGE_Player_World2Scene_transforSceneByEntry struct {
	sceneId int32 //0
	entryId int32 //1
}
type SGE_Player_World2Scene_move struct {
	pos COM_FPosition //0
}
type SGE_Player_World2Scene_moveToNpc struct {
	npcid int32 //0
}
type SGE_Player_World2Scene_moveToNpc2 struct {
	type_ int //0
}
type SGE_Player_World2Scene_moveToZone struct {
	sceneId int32 //0
	zoneId  int32 //1
}
type SGE_Player_World2Scene_addFollow struct {
	scenePlayerId int32 //0
}
type SGE_Player_World2Scene_delFollow struct {
	scenePlayerId int32 //0
}
type SGE_Player_World2Scene_addFollows struct {
	scenePlayers []int32 //0
}
type SGE_Player_World2Scene_setEntryFlag struct {
	scenePlayerId int32 //0
	isFlag        bool  //1
}
type SGE_Player_World2Scene_addNpc struct {
	npcid int32 //0
}
type SGE_Player_World2Scene_delNpc struct {
	npcid int32 //0
}
type SGE_Player_World2Scene_findDynamicNpc struct {
	npcId int32 //0
}
type SGE_Player_World2SceneStub struct {
	Sender prpc.StubSender
}
type SGE_Player_World2SceneProxy interface {
	initScenePlayer(info SGE_ScenePlayerInfo) error          // 0
	scenePlayerUpLevel(level int32) error                    // 1
	scenePlayerAddCurrentQuest(questId int32) error          // 2
	scenePlayerDelCurrentQuest(questId int32) error          // 3
	scenePlayerAddAcceptableQuest(questId int32) error       // 4
	scenePlayerDelAcceptableQuest(questId int32) error       // 5
	openScene(sceneId int32) error                           // 6
	joinBattle() error                                       // 7
	finishBattle() error                                     // 8
	transforScene(sceneId int32) error                       // 9
	transforSceneByEntry(sceneId int32, entryId int32) error // 10
	backHomeScene() error                                    // 11
	sceneLoaded() error                                      // 12
	move(pos COM_FPosition) error                            // 13
	moveToNpc(npcid int32) error                             // 14
	moveToNpc2(type_ int) error                              // 15
	moveToZone(sceneId int32, zoneId int32) error            // 16
	autoBattle() error                                       // 17
	stopMove() error                                         // 18
	addFollow(scenePlayerId int32) error                     // 19
	delFollow(scenePlayerId int32) error                     // 20
	delFollows() error                                       // 21
	addFollows(scenePlayers []int32) error                   // 22
	setEntryFlag(scenePlayerId int32, isFlag bool) error     // 23
	addNpc(npcid int32) error                                // 24
	delNpc(npcid int32) error                                // 25
	findDynamicNpc(npcId int32) error                        // 26
}

func (this *SGE_Player_World2Scene_initScenePlayer) Serialize(buffer *bytes.Buffer) error {
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
func (this *SGE_Player_World2Scene_initScenePlayer) Deserialize(buffer *bytes.Buffer) error {
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
func (this *SGE_Player_World2Scene_scenePlayerUpLevel) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.level != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize level
	{
		if this.level != 0 {
			err := prpc.Write(buffer, this.level)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *SGE_Player_World2Scene_scenePlayerUpLevel) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize level
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.level)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_Player_World2Scene_scenePlayerAddCurrentQuest) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.questId != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize questId
	{
		if this.questId != 0 {
			err := prpc.Write(buffer, this.questId)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *SGE_Player_World2Scene_scenePlayerAddCurrentQuest) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize questId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.questId)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_Player_World2Scene_scenePlayerDelCurrentQuest) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.questId != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize questId
	{
		if this.questId != 0 {
			err := prpc.Write(buffer, this.questId)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *SGE_Player_World2Scene_scenePlayerDelCurrentQuest) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize questId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.questId)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_Player_World2Scene_scenePlayerAddAcceptableQuest) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.questId != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize questId
	{
		if this.questId != 0 {
			err := prpc.Write(buffer, this.questId)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *SGE_Player_World2Scene_scenePlayerAddAcceptableQuest) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize questId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.questId)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_Player_World2Scene_scenePlayerDelAcceptableQuest) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.questId != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize questId
	{
		if this.questId != 0 {
			err := prpc.Write(buffer, this.questId)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *SGE_Player_World2Scene_scenePlayerDelAcceptableQuest) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize questId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.questId)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_Player_World2Scene_openScene) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.sceneId != 0)
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
	return nil
}
func (this *SGE_Player_World2Scene_openScene) Deserialize(buffer *bytes.Buffer) error {
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
	return nil
}
func (this *SGE_Player_World2Scene_transforScene) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.sceneId != 0)
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
	return nil
}
func (this *SGE_Player_World2Scene_transforScene) Deserialize(buffer *bytes.Buffer) error {
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
	return nil
}
func (this *SGE_Player_World2Scene_transforSceneByEntry) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.sceneId != 0)
	mask.WriteBit(this.entryId != 0)
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
	// serialize entryId
	{
		if this.entryId != 0 {
			err := prpc.Write(buffer, this.entryId)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *SGE_Player_World2Scene_transforSceneByEntry) Deserialize(buffer *bytes.Buffer) error {
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
	// deserialize entryId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.entryId)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_Player_World2Scene_move) Serialize(buffer *bytes.Buffer) error {
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
func (this *SGE_Player_World2Scene_move) Deserialize(buffer *bytes.Buffer) error {
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
func (this *SGE_Player_World2Scene_moveToNpc) Serialize(buffer *bytes.Buffer) error {
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
func (this *SGE_Player_World2Scene_moveToNpc) Deserialize(buffer *bytes.Buffer) error {
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
func (this *SGE_Player_World2Scene_moveToNpc2) Serialize(buffer *bytes.Buffer) error {
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
func (this *SGE_Player_World2Scene_moveToNpc2) Deserialize(buffer *bytes.Buffer) error {
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
func (this *SGE_Player_World2Scene_moveToZone) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.sceneId != 0)
	mask.WriteBit(this.zoneId != 0)
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
func (this *SGE_Player_World2Scene_moveToZone) Deserialize(buffer *bytes.Buffer) error {
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
	// deserialize zoneId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.zoneId)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_Player_World2Scene_addFollow) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.scenePlayerId != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize scenePlayerId
	{
		if this.scenePlayerId != 0 {
			err := prpc.Write(buffer, this.scenePlayerId)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *SGE_Player_World2Scene_addFollow) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize scenePlayerId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.scenePlayerId)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_Player_World2Scene_delFollow) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.scenePlayerId != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize scenePlayerId
	{
		if this.scenePlayerId != 0 {
			err := prpc.Write(buffer, this.scenePlayerId)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *SGE_Player_World2Scene_delFollow) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize scenePlayerId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.scenePlayerId)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_Player_World2Scene_addFollows) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.scenePlayers) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize scenePlayers
	if len(this.scenePlayers) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.scenePlayers)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.scenePlayers {
			err := prpc.Write(buffer, value)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *SGE_Player_World2Scene_addFollows) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize scenePlayers
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.scenePlayers = make([]int32, size)
		for i, _ := range this.scenePlayers {
			err := prpc.Read(buffer, &this.scenePlayers[i])
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *SGE_Player_World2Scene_setEntryFlag) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.scenePlayerId != 0)
	mask.WriteBit(this.isFlag)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize scenePlayerId
	{
		if this.scenePlayerId != 0 {
			err := prpc.Write(buffer, this.scenePlayerId)
			if err != nil {
				return err
			}
		}
	}
	// serialize isFlag
	{
	}
	return nil
}
func (this *SGE_Player_World2Scene_setEntryFlag) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize scenePlayerId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.scenePlayerId)
		if err != nil {
			return err
		}
	}
	// deserialize isFlag
	this.isFlag = mask.ReadBit()
	return nil
}
func (this *SGE_Player_World2Scene_addNpc) Serialize(buffer *bytes.Buffer) error {
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
func (this *SGE_Player_World2Scene_addNpc) Deserialize(buffer *bytes.Buffer) error {
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
func (this *SGE_Player_World2Scene_delNpc) Serialize(buffer *bytes.Buffer) error {
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
func (this *SGE_Player_World2Scene_delNpc) Deserialize(buffer *bytes.Buffer) error {
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
func (this *SGE_Player_World2Scene_findDynamicNpc) Serialize(buffer *bytes.Buffer) error {
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
func (this *SGE_Player_World2Scene_findDynamicNpc) Deserialize(buffer *bytes.Buffer) error {
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
func (this *SGE_Player_World2SceneStub) initScenePlayer(info SGE_ScenePlayerInfo) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(0))
	if err != nil {
		return err
	}
	_0 := SGE_Player_World2Scene_initScenePlayer{}
	_0.info = info
	err = _0.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_Player_World2SceneStub) scenePlayerUpLevel(level int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(1))
	if err != nil {
		return err
	}
	_1 := SGE_Player_World2Scene_scenePlayerUpLevel{}
	_1.level = level
	err = _1.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_Player_World2SceneStub) scenePlayerAddCurrentQuest(questId int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(2))
	if err != nil {
		return err
	}
	_2 := SGE_Player_World2Scene_scenePlayerAddCurrentQuest{}
	_2.questId = questId
	err = _2.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_Player_World2SceneStub) scenePlayerDelCurrentQuest(questId int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(3))
	if err != nil {
		return err
	}
	_3 := SGE_Player_World2Scene_scenePlayerDelCurrentQuest{}
	_3.questId = questId
	err = _3.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_Player_World2SceneStub) scenePlayerAddAcceptableQuest(questId int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(4))
	if err != nil {
		return err
	}
	_4 := SGE_Player_World2Scene_scenePlayerAddAcceptableQuest{}
	_4.questId = questId
	err = _4.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_Player_World2SceneStub) scenePlayerDelAcceptableQuest(questId int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(5))
	if err != nil {
		return err
	}
	_5 := SGE_Player_World2Scene_scenePlayerDelAcceptableQuest{}
	_5.questId = questId
	err = _5.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_Player_World2SceneStub) openScene(sceneId int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(6))
	if err != nil {
		return err
	}
	_6 := SGE_Player_World2Scene_openScene{}
	_6.sceneId = sceneId
	err = _6.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_Player_World2SceneStub) joinBattle() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(7))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_Player_World2SceneStub) finishBattle() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(8))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_Player_World2SceneStub) transforScene(sceneId int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(9))
	if err != nil {
		return err
	}
	_9 := SGE_Player_World2Scene_transforScene{}
	_9.sceneId = sceneId
	err = _9.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_Player_World2SceneStub) transforSceneByEntry(sceneId int32, entryId int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(10))
	if err != nil {
		return err
	}
	_10 := SGE_Player_World2Scene_transforSceneByEntry{}
	_10.sceneId = sceneId
	_10.entryId = entryId
	err = _10.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_Player_World2SceneStub) backHomeScene() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(11))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_Player_World2SceneStub) sceneLoaded() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(12))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_Player_World2SceneStub) move(pos COM_FPosition) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(13))
	if err != nil {
		return err
	}
	_13 := SGE_Player_World2Scene_move{}
	_13.pos = pos
	err = _13.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_Player_World2SceneStub) moveToNpc(npcid int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(14))
	if err != nil {
		return err
	}
	_14 := SGE_Player_World2Scene_moveToNpc{}
	_14.npcid = npcid
	err = _14.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_Player_World2SceneStub) moveToNpc2(type_ int) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(15))
	if err != nil {
		return err
	}
	_15 := SGE_Player_World2Scene_moveToNpc2{}
	_15.type_ = type_
	err = _15.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_Player_World2SceneStub) moveToZone(sceneId int32, zoneId int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(16))
	if err != nil {
		return err
	}
	_16 := SGE_Player_World2Scene_moveToZone{}
	_16.sceneId = sceneId
	_16.zoneId = zoneId
	err = _16.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_Player_World2SceneStub) autoBattle() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(17))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_Player_World2SceneStub) stopMove() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(18))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_Player_World2SceneStub) addFollow(scenePlayerId int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(19))
	if err != nil {
		return err
	}
	_19 := SGE_Player_World2Scene_addFollow{}
	_19.scenePlayerId = scenePlayerId
	err = _19.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_Player_World2SceneStub) delFollow(scenePlayerId int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(20))
	if err != nil {
		return err
	}
	_20 := SGE_Player_World2Scene_delFollow{}
	_20.scenePlayerId = scenePlayerId
	err = _20.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_Player_World2SceneStub) delFollows() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(21))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_Player_World2SceneStub) addFollows(scenePlayers []int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(22))
	if err != nil {
		return err
	}
	_22 := SGE_Player_World2Scene_addFollows{}
	_22.scenePlayers = scenePlayers
	err = _22.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_Player_World2SceneStub) setEntryFlag(scenePlayerId int32, isFlag bool) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(23))
	if err != nil {
		return err
	}
	_23 := SGE_Player_World2Scene_setEntryFlag{}
	_23.scenePlayerId = scenePlayerId
	_23.isFlag = isFlag
	err = _23.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_Player_World2SceneStub) addNpc(npcid int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(24))
	if err != nil {
		return err
	}
	_24 := SGE_Player_World2Scene_addNpc{}
	_24.npcid = npcid
	err = _24.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_Player_World2SceneStub) delNpc(npcid int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(25))
	if err != nil {
		return err
	}
	_25 := SGE_Player_World2Scene_delNpc{}
	_25.npcid = npcid
	err = _25.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_Player_World2SceneStub) findDynamicNpc(npcId int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(26))
	if err != nil {
		return err
	}
	_26 := SGE_Player_World2Scene_findDynamicNpc{}
	_26.npcId = npcId
	err = _26.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func Bridging_SGE_Player_World2Scene_initScenePlayer(buffer *bytes.Buffer, p SGE_Player_World2SceneProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_0 := SGE_Player_World2Scene_initScenePlayer{}
	err := _0.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.initScenePlayer(_0.info)
}
func Bridging_SGE_Player_World2Scene_scenePlayerUpLevel(buffer *bytes.Buffer, p SGE_Player_World2SceneProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_1 := SGE_Player_World2Scene_scenePlayerUpLevel{}
	err := _1.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.scenePlayerUpLevel(_1.level)
}
func Bridging_SGE_Player_World2Scene_scenePlayerAddCurrentQuest(buffer *bytes.Buffer, p SGE_Player_World2SceneProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_2 := SGE_Player_World2Scene_scenePlayerAddCurrentQuest{}
	err := _2.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.scenePlayerAddCurrentQuest(_2.questId)
}
func Bridging_SGE_Player_World2Scene_scenePlayerDelCurrentQuest(buffer *bytes.Buffer, p SGE_Player_World2SceneProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_3 := SGE_Player_World2Scene_scenePlayerDelCurrentQuest{}
	err := _3.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.scenePlayerDelCurrentQuest(_3.questId)
}
func Bridging_SGE_Player_World2Scene_scenePlayerAddAcceptableQuest(buffer *bytes.Buffer, p SGE_Player_World2SceneProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_4 := SGE_Player_World2Scene_scenePlayerAddAcceptableQuest{}
	err := _4.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.scenePlayerAddAcceptableQuest(_4.questId)
}
func Bridging_SGE_Player_World2Scene_scenePlayerDelAcceptableQuest(buffer *bytes.Buffer, p SGE_Player_World2SceneProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_5 := SGE_Player_World2Scene_scenePlayerDelAcceptableQuest{}
	err := _5.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.scenePlayerDelAcceptableQuest(_5.questId)
}
func Bridging_SGE_Player_World2Scene_openScene(buffer *bytes.Buffer, p SGE_Player_World2SceneProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_6 := SGE_Player_World2Scene_openScene{}
	err := _6.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.openScene(_6.sceneId)
}
func Bridging_SGE_Player_World2Scene_joinBattle(buffer *bytes.Buffer, p SGE_Player_World2SceneProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.joinBattle()
}
func Bridging_SGE_Player_World2Scene_finishBattle(buffer *bytes.Buffer, p SGE_Player_World2SceneProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.finishBattle()
}
func Bridging_SGE_Player_World2Scene_transforScene(buffer *bytes.Buffer, p SGE_Player_World2SceneProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_9 := SGE_Player_World2Scene_transforScene{}
	err := _9.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.transforScene(_9.sceneId)
}
func Bridging_SGE_Player_World2Scene_transforSceneByEntry(buffer *bytes.Buffer, p SGE_Player_World2SceneProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_10 := SGE_Player_World2Scene_transforSceneByEntry{}
	err := _10.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.transforSceneByEntry(_10.sceneId, _10.entryId)
}
func Bridging_SGE_Player_World2Scene_backHomeScene(buffer *bytes.Buffer, p SGE_Player_World2SceneProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.backHomeScene()
}
func Bridging_SGE_Player_World2Scene_sceneLoaded(buffer *bytes.Buffer, p SGE_Player_World2SceneProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.sceneLoaded()
}
func Bridging_SGE_Player_World2Scene_move(buffer *bytes.Buffer, p SGE_Player_World2SceneProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_13 := SGE_Player_World2Scene_move{}
	err := _13.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.move(_13.pos)
}
func Bridging_SGE_Player_World2Scene_moveToNpc(buffer *bytes.Buffer, p SGE_Player_World2SceneProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_14 := SGE_Player_World2Scene_moveToNpc{}
	err := _14.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.moveToNpc(_14.npcid)
}
func Bridging_SGE_Player_World2Scene_moveToNpc2(buffer *bytes.Buffer, p SGE_Player_World2SceneProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_15 := SGE_Player_World2Scene_moveToNpc2{}
	err := _15.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.moveToNpc2(_15.type_)
}
func Bridging_SGE_Player_World2Scene_moveToZone(buffer *bytes.Buffer, p SGE_Player_World2SceneProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_16 := SGE_Player_World2Scene_moveToZone{}
	err := _16.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.moveToZone(_16.sceneId, _16.zoneId)
}
func Bridging_SGE_Player_World2Scene_autoBattle(buffer *bytes.Buffer, p SGE_Player_World2SceneProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.autoBattle()
}
func Bridging_SGE_Player_World2Scene_stopMove(buffer *bytes.Buffer, p SGE_Player_World2SceneProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.stopMove()
}
func Bridging_SGE_Player_World2Scene_addFollow(buffer *bytes.Buffer, p SGE_Player_World2SceneProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_19 := SGE_Player_World2Scene_addFollow{}
	err := _19.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.addFollow(_19.scenePlayerId)
}
func Bridging_SGE_Player_World2Scene_delFollow(buffer *bytes.Buffer, p SGE_Player_World2SceneProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_20 := SGE_Player_World2Scene_delFollow{}
	err := _20.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.delFollow(_20.scenePlayerId)
}
func Bridging_SGE_Player_World2Scene_delFollows(buffer *bytes.Buffer, p SGE_Player_World2SceneProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.delFollows()
}
func Bridging_SGE_Player_World2Scene_addFollows(buffer *bytes.Buffer, p SGE_Player_World2SceneProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_22 := SGE_Player_World2Scene_addFollows{}
	err := _22.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.addFollows(_22.scenePlayers)
}
func Bridging_SGE_Player_World2Scene_setEntryFlag(buffer *bytes.Buffer, p SGE_Player_World2SceneProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_23 := SGE_Player_World2Scene_setEntryFlag{}
	err := _23.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.setEntryFlag(_23.scenePlayerId, _23.isFlag)
}
func Bridging_SGE_Player_World2Scene_addNpc(buffer *bytes.Buffer, p SGE_Player_World2SceneProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_24 := SGE_Player_World2Scene_addNpc{}
	err := _24.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.addNpc(_24.npcid)
}
func Bridging_SGE_Player_World2Scene_delNpc(buffer *bytes.Buffer, p SGE_Player_World2SceneProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_25 := SGE_Player_World2Scene_delNpc{}
	err := _25.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.delNpc(_25.npcid)
}
func Bridging_SGE_Player_World2Scene_findDynamicNpc(buffer *bytes.Buffer, p SGE_Player_World2SceneProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_26 := SGE_Player_World2Scene_findDynamicNpc{}
	err := _26.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.findDynamicNpc(_26.npcId)
}
func SGE_Player_World2SceneDispatch(buffer *bytes.Buffer, p SGE_Player_World2SceneProxy) error {
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
		return Bridging_SGE_Player_World2Scene_initScenePlayer(buffer, p)
	case 1:
		return Bridging_SGE_Player_World2Scene_scenePlayerUpLevel(buffer, p)
	case 2:
		return Bridging_SGE_Player_World2Scene_scenePlayerAddCurrentQuest(buffer, p)
	case 3:
		return Bridging_SGE_Player_World2Scene_scenePlayerDelCurrentQuest(buffer, p)
	case 4:
		return Bridging_SGE_Player_World2Scene_scenePlayerAddAcceptableQuest(buffer, p)
	case 5:
		return Bridging_SGE_Player_World2Scene_scenePlayerDelAcceptableQuest(buffer, p)
	case 6:
		return Bridging_SGE_Player_World2Scene_openScene(buffer, p)
	case 7:
		return Bridging_SGE_Player_World2Scene_joinBattle(buffer, p)
	case 8:
		return Bridging_SGE_Player_World2Scene_finishBattle(buffer, p)
	case 9:
		return Bridging_SGE_Player_World2Scene_transforScene(buffer, p)
	case 10:
		return Bridging_SGE_Player_World2Scene_transforSceneByEntry(buffer, p)
	case 11:
		return Bridging_SGE_Player_World2Scene_backHomeScene(buffer, p)
	case 12:
		return Bridging_SGE_Player_World2Scene_sceneLoaded(buffer, p)
	case 13:
		return Bridging_SGE_Player_World2Scene_move(buffer, p)
	case 14:
		return Bridging_SGE_Player_World2Scene_moveToNpc(buffer, p)
	case 15:
		return Bridging_SGE_Player_World2Scene_moveToNpc2(buffer, p)
	case 16:
		return Bridging_SGE_Player_World2Scene_moveToZone(buffer, p)
	case 17:
		return Bridging_SGE_Player_World2Scene_autoBattle(buffer, p)
	case 18:
		return Bridging_SGE_Player_World2Scene_stopMove(buffer, p)
	case 19:
		return Bridging_SGE_Player_World2Scene_addFollow(buffer, p)
	case 20:
		return Bridging_SGE_Player_World2Scene_delFollow(buffer, p)
	case 21:
		return Bridging_SGE_Player_World2Scene_delFollows(buffer, p)
	case 22:
		return Bridging_SGE_Player_World2Scene_addFollows(buffer, p)
	case 23:
		return Bridging_SGE_Player_World2Scene_setEntryFlag(buffer, p)
	case 24:
		return Bridging_SGE_Player_World2Scene_addNpc(buffer, p)
	case 25:
		return Bridging_SGE_Player_World2Scene_delNpc(buffer, p)
	case 26:
		return Bridging_SGE_Player_World2Scene_findDynamicNpc(buffer, p)
	default:
		return errors.New(prpc.NoneDispatchMatchError)
	}
}
