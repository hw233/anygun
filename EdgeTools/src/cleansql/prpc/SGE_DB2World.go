package prpc

import (
	"bytes"
	"errors"
	"suzuki/prpc"
)

type SGE_DB2World_syncGlobalGuid struct {
	id uint32 //0
}
type SGE_DB2World_syncContactInfo struct {
	info []SGE_ContactInfoExt //0
}
type SGE_DB2World_queryPlayerOk struct {
	username string             //0
	insts    []SGE_DBPlayerData //1
	inDoorId int32              //2
}
type SGE_DB2World_createPlayerOk struct {
	username string           //0
	inst     SGE_DBPlayerData //1
	inDoorId int32            //2
}
type SGE_DB2World_createPlayerSameName struct {
	username string //0
}
type SGE_DB2World_queryPlayerByIdOK struct {
	playername string           //0
	inst       SGE_DBPlayerData //1
	isFlag     bool             //2
}
type SGE_DB2World_queryEndlessStairAllDateOK struct {
	name []string //0
}
type SGE_DB2World_queryPlayerByLevelOK struct {
	info []COM_ContactInfo //0
}
type SGE_DB2World_queryPlayerByFFOK struct {
	info []COM_ContactInfo //0
}
type SGE_DB2World_createBabyOK struct {
	playername  string       //0
	inst        COM_BabyInst //1
	isToStorage bool         //2
}
type SGE_DB2World_deleteBabyOK struct {
	playername string //0
	babyInstId int32  //1
}
type SGE_DB2World_queryBabyByFFOK struct {
	infos []COM_BabyRankData //0
}
type SGE_DB2World_queryBabyByIdOK struct {
	name string       //0
	inst COM_BabyInst //1
}
type SGE_DB2World_UpdateBabysOK struct {
	playername string //0
}
type SGE_DB2World_createEmployeeOK struct {
	playername string           //0
	inst       COM_EmployeeInst //1
}
type SGE_DB2World_deleteEmployeeOK struct {
	playername string   //0
	instIds    []uint32 //1
}
type SGE_DB2World_queryEmployeeByFFOK struct {
	infos []COM_EmployeeRankData //0
}
type SGE_DB2World_queryEmployeeByIdOK struct {
	name string           //0
	inst COM_EmployeeInst //1
}
type SGE_DB2World_appendMail struct {
	mails []COM_Mail //0
}
type SGE_DB2World_insertGuildOK struct {
	guild       COM_Guild       //0
	guildMember COM_GuildMember //1
}
type SGE_DB2World_updateMemberJobOk struct {
	roleId int32 //0
	job    int   //1
}
type SGE_DB2World_syncGuild struct {
	guilds []COM_Guild //0
}
type SGE_DB2World_syncGuildMember struct {
	guildMember []COM_GuildMember //0
}
type SGE_DB2World_queryIdgenOK struct {
	playerName string         //0
	content    COM_KeyContent //1
	isHas      bool           //2
}
type SGE_DB2World_fatchActivity struct {
	date SGE_SysActivity //0
}
type SGE_DB2World_syncEmployeeQuest struct {
	info []SGE_PlayerEmployeeQuest //0
}
type SGE_DB2WorldStub struct {
	Sender prpc.StubSender
}
type SGE_DB2WorldProxy interface {
	syncGlobalGuid(id uint32) error                                                // 0
	syncContactInfo(info []SGE_ContactInfoExt) error                               // 1
	queryPlayerOk(username string, insts []SGE_DBPlayerData, inDoorId int32) error // 2
	createPlayerOk(username string, inst SGE_DBPlayerData, inDoorId int32) error   // 3
	createPlayerSameName(username string) error                                    // 4
	queryPlayerByIdOK(playername string, inst SGE_DBPlayerData, isFlag bool) error // 5
	queryEndlessStairAllDateOK(name []string) error                                // 6
	queryPlayerByLevelOK(info []COM_ContactInfo) error                             // 7
	queryPlayerByFFOK(info []COM_ContactInfo) error                                // 8
	createBabyOK(playername string, inst COM_BabyInst, isToStorage bool) error     // 9
	deleteBabyOK(playername string, babyInstId int32) error                        // 10
	queryBabyByFFOK(infos []COM_BabyRankData) error                                // 11
	queryBabyByIdOK(name string, inst COM_BabyInst) error                          // 12
	UpdateBabysOK(playername string) error                                         // 13
	createEmployeeOK(playername string, inst COM_EmployeeInst) error               // 14
	deleteEmployeeOK(playername string, instIds []uint32) error                    // 15
	queryEmployeeByFFOK(infos []COM_EmployeeRankData) error                        // 16
	queryEmployeeByIdOK(name string, inst COM_EmployeeInst) error                  // 17
	appendMail(mails []COM_Mail) error                                             // 18
	insertGuildOK(guild COM_Guild, guildMember COM_GuildMember) error              // 19
	updateMemberJobOk(roleId int32, job int) error                                 // 20
	syncGuild(guilds []COM_Guild) error                                            // 21
	syncGuildMember(guildMember []COM_GuildMember) error                           // 22
	queryIdgenOK(playerName string, content COM_KeyContent, isHas bool) error      // 23
	fatchActivity(date SGE_SysActivity) error                                      // 24
	syncEmployeeQuest(info []SGE_PlayerEmployeeQuest) error                        // 25
}

func (this *SGE_DB2World_syncGlobalGuid) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.id != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize id
	{
		if this.id != 0 {
			err := prpc.Write(buffer, this.id)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *SGE_DB2World_syncGlobalGuid) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize id
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.id)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_DB2World_syncContactInfo) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.info) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize info
	if len(this.info) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.info)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.info {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *SGE_DB2World_syncContactInfo) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize info
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.info = make([]SGE_ContactInfoExt, size)
		for i, _ := range this.info {
			err := this.info[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *SGE_DB2World_queryPlayerOk) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.username) != 0)
	mask.WriteBit(len(this.insts) != 0)
	mask.WriteBit(this.inDoorId != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize username
	if len(this.username) != 0 {
		err := prpc.Write(buffer, this.username)
		if err != nil {
			return err
		}
	}
	// serialize insts
	if len(this.insts) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.insts)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.insts {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// serialize inDoorId
	{
		if this.inDoorId != 0 {
			err := prpc.Write(buffer, this.inDoorId)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *SGE_DB2World_queryPlayerOk) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize username
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.username)
		if err != nil {
			return err
		}
	}
	// deserialize insts
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.insts = make([]SGE_DBPlayerData, size)
		for i, _ := range this.insts {
			err := this.insts[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// deserialize inDoorId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.inDoorId)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_DB2World_createPlayerOk) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.username) != 0)
	mask.WriteBit(true) //inst
	mask.WriteBit(this.inDoorId != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize username
	if len(this.username) != 0 {
		err := prpc.Write(buffer, this.username)
		if err != nil {
			return err
		}
	}
	// serialize inst
	{
		err := this.inst.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	// serialize inDoorId
	{
		if this.inDoorId != 0 {
			err := prpc.Write(buffer, this.inDoorId)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *SGE_DB2World_createPlayerOk) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize username
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.username)
		if err != nil {
			return err
		}
	}
	// deserialize inst
	if mask.ReadBit() {
		err := this.inst.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	// deserialize inDoorId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.inDoorId)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_DB2World_createPlayerSameName) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.username) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize username
	if len(this.username) != 0 {
		err := prpc.Write(buffer, this.username)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_DB2World_createPlayerSameName) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize username
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.username)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_DB2World_queryPlayerByIdOK) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.playername) != 0)
	mask.WriteBit(true) //inst
	mask.WriteBit(this.isFlag)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize playername
	if len(this.playername) != 0 {
		err := prpc.Write(buffer, this.playername)
		if err != nil {
			return err
		}
	}
	// serialize inst
	{
		err := this.inst.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	// serialize isFlag
	{
	}
	return nil
}
func (this *SGE_DB2World_queryPlayerByIdOK) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize playername
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.playername)
		if err != nil {
			return err
		}
	}
	// deserialize inst
	if mask.ReadBit() {
		err := this.inst.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	// deserialize isFlag
	this.isFlag = mask.ReadBit()
	return nil
}
func (this *SGE_DB2World_queryEndlessStairAllDateOK) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.name) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize name
	if len(this.name) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.name)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.name {
			err := prpc.Write(buffer, value)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *SGE_DB2World_queryEndlessStairAllDateOK) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize name
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.name = make([]string, size)
		for i, _ := range this.name {
			err := prpc.Read(buffer, &this.name[i])
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *SGE_DB2World_queryPlayerByLevelOK) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.info) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize info
	if len(this.info) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.info)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.info {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *SGE_DB2World_queryPlayerByLevelOK) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize info
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.info = make([]COM_ContactInfo, size)
		for i, _ := range this.info {
			err := this.info[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *SGE_DB2World_queryPlayerByFFOK) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.info) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize info
	if len(this.info) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.info)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.info {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *SGE_DB2World_queryPlayerByFFOK) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize info
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.info = make([]COM_ContactInfo, size)
		for i, _ := range this.info {
			err := this.info[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *SGE_DB2World_createBabyOK) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.playername) != 0)
	mask.WriteBit(true) //inst
	mask.WriteBit(this.isToStorage)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize playername
	if len(this.playername) != 0 {
		err := prpc.Write(buffer, this.playername)
		if err != nil {
			return err
		}
	}
	// serialize inst
	{
		err := this.inst.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	// serialize isToStorage
	{
	}
	return nil
}
func (this *SGE_DB2World_createBabyOK) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize playername
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.playername)
		if err != nil {
			return err
		}
	}
	// deserialize inst
	if mask.ReadBit() {
		err := this.inst.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	// deserialize isToStorage
	this.isToStorage = mask.ReadBit()
	return nil
}
func (this *SGE_DB2World_deleteBabyOK) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.playername) != 0)
	mask.WriteBit(this.babyInstId != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize playername
	if len(this.playername) != 0 {
		err := prpc.Write(buffer, this.playername)
		if err != nil {
			return err
		}
	}
	// serialize babyInstId
	{
		if this.babyInstId != 0 {
			err := prpc.Write(buffer, this.babyInstId)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *SGE_DB2World_deleteBabyOK) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize playername
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.playername)
		if err != nil {
			return err
		}
	}
	// deserialize babyInstId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.babyInstId)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_DB2World_queryBabyByFFOK) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.infos) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize infos
	if len(this.infos) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.infos)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.infos {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *SGE_DB2World_queryBabyByFFOK) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize infos
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.infos = make([]COM_BabyRankData, size)
		for i, _ := range this.infos {
			err := this.infos[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *SGE_DB2World_queryBabyByIdOK) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.name) != 0)
	mask.WriteBit(true) //inst
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize name
	if len(this.name) != 0 {
		err := prpc.Write(buffer, this.name)
		if err != nil {
			return err
		}
	}
	// serialize inst
	{
		err := this.inst.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_DB2World_queryBabyByIdOK) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize name
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.name)
		if err != nil {
			return err
		}
	}
	// deserialize inst
	if mask.ReadBit() {
		err := this.inst.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_DB2World_UpdateBabysOK) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.playername) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize playername
	if len(this.playername) != 0 {
		err := prpc.Write(buffer, this.playername)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_DB2World_UpdateBabysOK) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize playername
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.playername)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_DB2World_createEmployeeOK) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.playername) != 0)
	mask.WriteBit(true) //inst
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize playername
	if len(this.playername) != 0 {
		err := prpc.Write(buffer, this.playername)
		if err != nil {
			return err
		}
	}
	// serialize inst
	{
		err := this.inst.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_DB2World_createEmployeeOK) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize playername
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.playername)
		if err != nil {
			return err
		}
	}
	// deserialize inst
	if mask.ReadBit() {
		err := this.inst.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_DB2World_deleteEmployeeOK) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.playername) != 0)
	mask.WriteBit(len(this.instIds) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize playername
	if len(this.playername) != 0 {
		err := prpc.Write(buffer, this.playername)
		if err != nil {
			return err
		}
	}
	// serialize instIds
	if len(this.instIds) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.instIds)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.instIds {
			err := prpc.Write(buffer, value)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *SGE_DB2World_deleteEmployeeOK) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize playername
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.playername)
		if err != nil {
			return err
		}
	}
	// deserialize instIds
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.instIds = make([]uint32, size)
		for i, _ := range this.instIds {
			err := prpc.Read(buffer, &this.instIds[i])
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *SGE_DB2World_queryEmployeeByFFOK) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.infos) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize infos
	if len(this.infos) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.infos)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.infos {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *SGE_DB2World_queryEmployeeByFFOK) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize infos
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.infos = make([]COM_EmployeeRankData, size)
		for i, _ := range this.infos {
			err := this.infos[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *SGE_DB2World_queryEmployeeByIdOK) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.name) != 0)
	mask.WriteBit(true) //inst
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize name
	if len(this.name) != 0 {
		err := prpc.Write(buffer, this.name)
		if err != nil {
			return err
		}
	}
	// serialize inst
	{
		err := this.inst.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_DB2World_queryEmployeeByIdOK) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize name
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.name)
		if err != nil {
			return err
		}
	}
	// deserialize inst
	if mask.ReadBit() {
		err := this.inst.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_DB2World_appendMail) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.mails) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize mails
	if len(this.mails) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.mails)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.mails {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *SGE_DB2World_appendMail) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize mails
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.mails = make([]COM_Mail, size)
		for i, _ := range this.mails {
			err := this.mails[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *SGE_DB2World_insertGuildOK) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //guild
	mask.WriteBit(true) //guildMember
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize guild
	{
		err := this.guild.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	// serialize guildMember
	{
		err := this.guildMember.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_DB2World_insertGuildOK) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize guild
	if mask.ReadBit() {
		err := this.guild.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	// deserialize guildMember
	if mask.ReadBit() {
		err := this.guildMember.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_DB2World_updateMemberJobOk) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.roleId != 0)
	mask.WriteBit(this.job != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize roleId
	{
		if this.roleId != 0 {
			err := prpc.Write(buffer, this.roleId)
			if err != nil {
				return err
			}
		}
	}
	// serialize job
	{
		if this.job != 0 {
			err := prpc.Write(buffer, this.job)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *SGE_DB2World_updateMemberJobOk) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize roleId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.roleId)
		if err != nil {
			return err
		}
	}
	// deserialize job
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.job)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_DB2World_syncGuild) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.guilds) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize guilds
	if len(this.guilds) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.guilds)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.guilds {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *SGE_DB2World_syncGuild) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize guilds
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.guilds = make([]COM_Guild, size)
		for i, _ := range this.guilds {
			err := this.guilds[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *SGE_DB2World_syncGuildMember) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.guildMember) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize guildMember
	if len(this.guildMember) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.guildMember)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.guildMember {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *SGE_DB2World_syncGuildMember) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize guildMember
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.guildMember = make([]COM_GuildMember, size)
		for i, _ := range this.guildMember {
			err := this.guildMember[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *SGE_DB2World_queryIdgenOK) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.playerName) != 0)
	mask.WriteBit(true) //content
	mask.WriteBit(this.isHas)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize playerName
	if len(this.playerName) != 0 {
		err := prpc.Write(buffer, this.playerName)
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
	// serialize isHas
	{
	}
	return nil
}
func (this *SGE_DB2World_queryIdgenOK) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize playerName
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.playerName)
		if err != nil {
			return err
		}
	}
	// deserialize content
	if mask.ReadBit() {
		err := this.content.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	// deserialize isHas
	this.isHas = mask.ReadBit()
	return nil
}
func (this *SGE_DB2World_fatchActivity) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //date
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize date
	{
		err := this.date.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_DB2World_fatchActivity) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize date
	if mask.ReadBit() {
		err := this.date.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_DB2World_syncEmployeeQuest) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.info) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize info
	if len(this.info) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.info)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.info {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *SGE_DB2World_syncEmployeeQuest) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize info
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.info = make([]SGE_PlayerEmployeeQuest, size)
		for i, _ := range this.info {
			err := this.info[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *SGE_DB2WorldStub) syncGlobalGuid(id uint32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(0))
	if err != nil {
		return err
	}
	_0 := SGE_DB2World_syncGlobalGuid{}
	_0.id = id
	err = _0.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_DB2WorldStub) syncContactInfo(info []SGE_ContactInfoExt) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(1))
	if err != nil {
		return err
	}
	_1 := SGE_DB2World_syncContactInfo{}
	_1.info = info
	err = _1.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_DB2WorldStub) queryPlayerOk(username string, insts []SGE_DBPlayerData, inDoorId int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(2))
	if err != nil {
		return err
	}
	_2 := SGE_DB2World_queryPlayerOk{}
	_2.username = username
	_2.insts = insts
	_2.inDoorId = inDoorId
	err = _2.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_DB2WorldStub) createPlayerOk(username string, inst SGE_DBPlayerData, inDoorId int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(3))
	if err != nil {
		return err
	}
	_3 := SGE_DB2World_createPlayerOk{}
	_3.username = username
	_3.inst = inst
	_3.inDoorId = inDoorId
	err = _3.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_DB2WorldStub) createPlayerSameName(username string) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(4))
	if err != nil {
		return err
	}
	_4 := SGE_DB2World_createPlayerSameName{}
	_4.username = username
	err = _4.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_DB2WorldStub) queryPlayerByIdOK(playername string, inst SGE_DBPlayerData, isFlag bool) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(5))
	if err != nil {
		return err
	}
	_5 := SGE_DB2World_queryPlayerByIdOK{}
	_5.playername = playername
	_5.inst = inst
	_5.isFlag = isFlag
	err = _5.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_DB2WorldStub) queryEndlessStairAllDateOK(name []string) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(6))
	if err != nil {
		return err
	}
	_6 := SGE_DB2World_queryEndlessStairAllDateOK{}
	_6.name = name
	err = _6.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_DB2WorldStub) queryPlayerByLevelOK(info []COM_ContactInfo) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(7))
	if err != nil {
		return err
	}
	_7 := SGE_DB2World_queryPlayerByLevelOK{}
	_7.info = info
	err = _7.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_DB2WorldStub) queryPlayerByFFOK(info []COM_ContactInfo) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(8))
	if err != nil {
		return err
	}
	_8 := SGE_DB2World_queryPlayerByFFOK{}
	_8.info = info
	err = _8.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_DB2WorldStub) createBabyOK(playername string, inst COM_BabyInst, isToStorage bool) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(9))
	if err != nil {
		return err
	}
	_9 := SGE_DB2World_createBabyOK{}
	_9.playername = playername
	_9.inst = inst
	_9.isToStorage = isToStorage
	err = _9.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_DB2WorldStub) deleteBabyOK(playername string, babyInstId int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(10))
	if err != nil {
		return err
	}
	_10 := SGE_DB2World_deleteBabyOK{}
	_10.playername = playername
	_10.babyInstId = babyInstId
	err = _10.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_DB2WorldStub) queryBabyByFFOK(infos []COM_BabyRankData) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(11))
	if err != nil {
		return err
	}
	_11 := SGE_DB2World_queryBabyByFFOK{}
	_11.infos = infos
	err = _11.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_DB2WorldStub) queryBabyByIdOK(name string, inst COM_BabyInst) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(12))
	if err != nil {
		return err
	}
	_12 := SGE_DB2World_queryBabyByIdOK{}
	_12.name = name
	_12.inst = inst
	err = _12.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_DB2WorldStub) UpdateBabysOK(playername string) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(13))
	if err != nil {
		return err
	}
	_13 := SGE_DB2World_UpdateBabysOK{}
	_13.playername = playername
	err = _13.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_DB2WorldStub) createEmployeeOK(playername string, inst COM_EmployeeInst) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(14))
	if err != nil {
		return err
	}
	_14 := SGE_DB2World_createEmployeeOK{}
	_14.playername = playername
	_14.inst = inst
	err = _14.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_DB2WorldStub) deleteEmployeeOK(playername string, instIds []uint32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(15))
	if err != nil {
		return err
	}
	_15 := SGE_DB2World_deleteEmployeeOK{}
	_15.playername = playername
	_15.instIds = instIds
	err = _15.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_DB2WorldStub) queryEmployeeByFFOK(infos []COM_EmployeeRankData) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(16))
	if err != nil {
		return err
	}
	_16 := SGE_DB2World_queryEmployeeByFFOK{}
	_16.infos = infos
	err = _16.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_DB2WorldStub) queryEmployeeByIdOK(name string, inst COM_EmployeeInst) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(17))
	if err != nil {
		return err
	}
	_17 := SGE_DB2World_queryEmployeeByIdOK{}
	_17.name = name
	_17.inst = inst
	err = _17.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_DB2WorldStub) appendMail(mails []COM_Mail) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(18))
	if err != nil {
		return err
	}
	_18 := SGE_DB2World_appendMail{}
	_18.mails = mails
	err = _18.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_DB2WorldStub) insertGuildOK(guild COM_Guild, guildMember COM_GuildMember) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(19))
	if err != nil {
		return err
	}
	_19 := SGE_DB2World_insertGuildOK{}
	_19.guild = guild
	_19.guildMember = guildMember
	err = _19.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_DB2WorldStub) updateMemberJobOk(roleId int32, job int) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(20))
	if err != nil {
		return err
	}
	_20 := SGE_DB2World_updateMemberJobOk{}
	_20.roleId = roleId
	_20.job = job
	err = _20.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_DB2WorldStub) syncGuild(guilds []COM_Guild) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(21))
	if err != nil {
		return err
	}
	_21 := SGE_DB2World_syncGuild{}
	_21.guilds = guilds
	err = _21.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_DB2WorldStub) syncGuildMember(guildMember []COM_GuildMember) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(22))
	if err != nil {
		return err
	}
	_22 := SGE_DB2World_syncGuildMember{}
	_22.guildMember = guildMember
	err = _22.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_DB2WorldStub) queryIdgenOK(playerName string, content COM_KeyContent, isHas bool) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(23))
	if err != nil {
		return err
	}
	_23 := SGE_DB2World_queryIdgenOK{}
	_23.playerName = playerName
	_23.content = content
	_23.isHas = isHas
	err = _23.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_DB2WorldStub) fatchActivity(date SGE_SysActivity) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(24))
	if err != nil {
		return err
	}
	_24 := SGE_DB2World_fatchActivity{}
	_24.date = date
	err = _24.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_DB2WorldStub) syncEmployeeQuest(info []SGE_PlayerEmployeeQuest) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(25))
	if err != nil {
		return err
	}
	_25 := SGE_DB2World_syncEmployeeQuest{}
	_25.info = info
	err = _25.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func Bridging_SGE_DB2World_syncGlobalGuid(buffer *bytes.Buffer, p SGE_DB2WorldProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_0 := SGE_DB2World_syncGlobalGuid{}
	err := _0.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.syncGlobalGuid(_0.id)
}
func Bridging_SGE_DB2World_syncContactInfo(buffer *bytes.Buffer, p SGE_DB2WorldProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_1 := SGE_DB2World_syncContactInfo{}
	err := _1.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.syncContactInfo(_1.info)
}
func Bridging_SGE_DB2World_queryPlayerOk(buffer *bytes.Buffer, p SGE_DB2WorldProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_2 := SGE_DB2World_queryPlayerOk{}
	err := _2.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.queryPlayerOk(_2.username, _2.insts, _2.inDoorId)
}
func Bridging_SGE_DB2World_createPlayerOk(buffer *bytes.Buffer, p SGE_DB2WorldProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_3 := SGE_DB2World_createPlayerOk{}
	err := _3.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.createPlayerOk(_3.username, _3.inst, _3.inDoorId)
}
func Bridging_SGE_DB2World_createPlayerSameName(buffer *bytes.Buffer, p SGE_DB2WorldProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_4 := SGE_DB2World_createPlayerSameName{}
	err := _4.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.createPlayerSameName(_4.username)
}
func Bridging_SGE_DB2World_queryPlayerByIdOK(buffer *bytes.Buffer, p SGE_DB2WorldProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_5 := SGE_DB2World_queryPlayerByIdOK{}
	err := _5.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.queryPlayerByIdOK(_5.playername, _5.inst, _5.isFlag)
}
func Bridging_SGE_DB2World_queryEndlessStairAllDateOK(buffer *bytes.Buffer, p SGE_DB2WorldProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_6 := SGE_DB2World_queryEndlessStairAllDateOK{}
	err := _6.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.queryEndlessStairAllDateOK(_6.name)
}
func Bridging_SGE_DB2World_queryPlayerByLevelOK(buffer *bytes.Buffer, p SGE_DB2WorldProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_7 := SGE_DB2World_queryPlayerByLevelOK{}
	err := _7.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.queryPlayerByLevelOK(_7.info)
}
func Bridging_SGE_DB2World_queryPlayerByFFOK(buffer *bytes.Buffer, p SGE_DB2WorldProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_8 := SGE_DB2World_queryPlayerByFFOK{}
	err := _8.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.queryPlayerByFFOK(_8.info)
}
func Bridging_SGE_DB2World_createBabyOK(buffer *bytes.Buffer, p SGE_DB2WorldProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_9 := SGE_DB2World_createBabyOK{}
	err := _9.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.createBabyOK(_9.playername, _9.inst, _9.isToStorage)
}
func Bridging_SGE_DB2World_deleteBabyOK(buffer *bytes.Buffer, p SGE_DB2WorldProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_10 := SGE_DB2World_deleteBabyOK{}
	err := _10.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.deleteBabyOK(_10.playername, _10.babyInstId)
}
func Bridging_SGE_DB2World_queryBabyByFFOK(buffer *bytes.Buffer, p SGE_DB2WorldProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_11 := SGE_DB2World_queryBabyByFFOK{}
	err := _11.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.queryBabyByFFOK(_11.infos)
}
func Bridging_SGE_DB2World_queryBabyByIdOK(buffer *bytes.Buffer, p SGE_DB2WorldProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_12 := SGE_DB2World_queryBabyByIdOK{}
	err := _12.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.queryBabyByIdOK(_12.name, _12.inst)
}
func Bridging_SGE_DB2World_UpdateBabysOK(buffer *bytes.Buffer, p SGE_DB2WorldProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_13 := SGE_DB2World_UpdateBabysOK{}
	err := _13.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.UpdateBabysOK(_13.playername)
}
func Bridging_SGE_DB2World_createEmployeeOK(buffer *bytes.Buffer, p SGE_DB2WorldProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_14 := SGE_DB2World_createEmployeeOK{}
	err := _14.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.createEmployeeOK(_14.playername, _14.inst)
}
func Bridging_SGE_DB2World_deleteEmployeeOK(buffer *bytes.Buffer, p SGE_DB2WorldProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_15 := SGE_DB2World_deleteEmployeeOK{}
	err := _15.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.deleteEmployeeOK(_15.playername, _15.instIds)
}
func Bridging_SGE_DB2World_queryEmployeeByFFOK(buffer *bytes.Buffer, p SGE_DB2WorldProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_16 := SGE_DB2World_queryEmployeeByFFOK{}
	err := _16.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.queryEmployeeByFFOK(_16.infos)
}
func Bridging_SGE_DB2World_queryEmployeeByIdOK(buffer *bytes.Buffer, p SGE_DB2WorldProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_17 := SGE_DB2World_queryEmployeeByIdOK{}
	err := _17.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.queryEmployeeByIdOK(_17.name, _17.inst)
}
func Bridging_SGE_DB2World_appendMail(buffer *bytes.Buffer, p SGE_DB2WorldProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_18 := SGE_DB2World_appendMail{}
	err := _18.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.appendMail(_18.mails)
}
func Bridging_SGE_DB2World_insertGuildOK(buffer *bytes.Buffer, p SGE_DB2WorldProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_19 := SGE_DB2World_insertGuildOK{}
	err := _19.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.insertGuildOK(_19.guild, _19.guildMember)
}
func Bridging_SGE_DB2World_updateMemberJobOk(buffer *bytes.Buffer, p SGE_DB2WorldProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_20 := SGE_DB2World_updateMemberJobOk{}
	err := _20.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.updateMemberJobOk(_20.roleId, _20.job)
}
func Bridging_SGE_DB2World_syncGuild(buffer *bytes.Buffer, p SGE_DB2WorldProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_21 := SGE_DB2World_syncGuild{}
	err := _21.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.syncGuild(_21.guilds)
}
func Bridging_SGE_DB2World_syncGuildMember(buffer *bytes.Buffer, p SGE_DB2WorldProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_22 := SGE_DB2World_syncGuildMember{}
	err := _22.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.syncGuildMember(_22.guildMember)
}
func Bridging_SGE_DB2World_queryIdgenOK(buffer *bytes.Buffer, p SGE_DB2WorldProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_23 := SGE_DB2World_queryIdgenOK{}
	err := _23.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.queryIdgenOK(_23.playerName, _23.content, _23.isHas)
}
func Bridging_SGE_DB2World_fatchActivity(buffer *bytes.Buffer, p SGE_DB2WorldProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_24 := SGE_DB2World_fatchActivity{}
	err := _24.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.fatchActivity(_24.date)
}
func Bridging_SGE_DB2World_syncEmployeeQuest(buffer *bytes.Buffer, p SGE_DB2WorldProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_25 := SGE_DB2World_syncEmployeeQuest{}
	err := _25.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.syncEmployeeQuest(_25.info)
}
func SGE_DB2WorldDispatch(buffer *bytes.Buffer, p SGE_DB2WorldProxy) error {
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
		return Bridging_SGE_DB2World_syncGlobalGuid(buffer, p)
	case 1:
		return Bridging_SGE_DB2World_syncContactInfo(buffer, p)
	case 2:
		return Bridging_SGE_DB2World_queryPlayerOk(buffer, p)
	case 3:
		return Bridging_SGE_DB2World_createPlayerOk(buffer, p)
	case 4:
		return Bridging_SGE_DB2World_createPlayerSameName(buffer, p)
	case 5:
		return Bridging_SGE_DB2World_queryPlayerByIdOK(buffer, p)
	case 6:
		return Bridging_SGE_DB2World_queryEndlessStairAllDateOK(buffer, p)
	case 7:
		return Bridging_SGE_DB2World_queryPlayerByLevelOK(buffer, p)
	case 8:
		return Bridging_SGE_DB2World_queryPlayerByFFOK(buffer, p)
	case 9:
		return Bridging_SGE_DB2World_createBabyOK(buffer, p)
	case 10:
		return Bridging_SGE_DB2World_deleteBabyOK(buffer, p)
	case 11:
		return Bridging_SGE_DB2World_queryBabyByFFOK(buffer, p)
	case 12:
		return Bridging_SGE_DB2World_queryBabyByIdOK(buffer, p)
	case 13:
		return Bridging_SGE_DB2World_UpdateBabysOK(buffer, p)
	case 14:
		return Bridging_SGE_DB2World_createEmployeeOK(buffer, p)
	case 15:
		return Bridging_SGE_DB2World_deleteEmployeeOK(buffer, p)
	case 16:
		return Bridging_SGE_DB2World_queryEmployeeByFFOK(buffer, p)
	case 17:
		return Bridging_SGE_DB2World_queryEmployeeByIdOK(buffer, p)
	case 18:
		return Bridging_SGE_DB2World_appendMail(buffer, p)
	case 19:
		return Bridging_SGE_DB2World_insertGuildOK(buffer, p)
	case 20:
		return Bridging_SGE_DB2World_updateMemberJobOk(buffer, p)
	case 21:
		return Bridging_SGE_DB2World_syncGuild(buffer, p)
	case 22:
		return Bridging_SGE_DB2World_syncGuildMember(buffer, p)
	case 23:
		return Bridging_SGE_DB2World_queryIdgenOK(buffer, p)
	case 24:
		return Bridging_SGE_DB2World_fatchActivity(buffer, p)
	case 25:
		return Bridging_SGE_DB2World_syncEmployeeQuest(buffer, p)
	default:
		return errors.New(prpc.NoneDispatchMatchError)
	}
}
