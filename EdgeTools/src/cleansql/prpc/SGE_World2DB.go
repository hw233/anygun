package prpc

import (
	"bytes"
	"errors"
	"suzuki/prpc"
)

type SGE_World2DB_queryPlayer struct {
	username string //0
	serverId int32  //1
}
type SGE_World2DB_createPlayer struct {
	username string           //0
	inst     SGE_DBPlayerData //1
	serverId int32            //2
}
type SGE_World2DB_updatePlayer struct {
	username string           //0
	inst     SGE_DBPlayerData //1
}
type SGE_World2DB_deletePlayer struct {
	playername string //0
}
type SGE_World2DB_queryPlayerById struct {
	name   string //0
	instId int32  //1
	isFlag bool   //2
}
type SGE_World2DB_insertEndlessStair struct {
	rank int32  //0
	name string //1
}
type SGE_World2DB_updateEndlessStair struct {
	rank int32  //0
	name string //1
}
type SGE_World2DB_createBaby struct {
	playername  string       //0
	inst        COM_BabyInst //1
	isToStorage bool         //2
}
type SGE_World2DB_deleteBaby struct {
	playername string //0
	babyInstId int32  //1
}
type SGE_World2DB_resetBabyOwner struct {
	playername string //0
	babyInstId int32  //1
}
type SGE_World2DB_updateBaby struct {
	inst COM_BabyInst //0
}
type SGE_World2DB_updateBabys struct {
	playername string         //0
	babys      []COM_BabyInst //1
}
type SGE_World2DB_queryBabyById struct {
	name   string //0
	instid uint32 //1
}
type SGE_World2DB_createEmployee struct {
	playername string           //0
	inst       COM_EmployeeInst //1
}
type SGE_World2DB_deleteEmployee struct {
	playername string   //0
	instIds    []uint32 //1
}
type SGE_World2DB_updateEmployee struct {
	inst COM_EmployeeInst //0
}
type SGE_World2DB_queryEmployeeById struct {
	name   string //0
	instid uint32 //1
}
type SGE_World2DB_insertMail struct {
	mail COM_Mail //0
}
type SGE_World2DB_insertMailAll struct {
	mail COM_Mail //0
}
type SGE_World2DB_insertMailByRecvs struct {
	mail  COM_Mail //0
	recvs []string //1
}
type SGE_World2DB_fatchMail struct {
	recvName string //0
	mailId   int32  //1
}
type SGE_World2DB_delMail struct {
	recvName string //0
	mailId   int32  //1
}
type SGE_World2DB_updateMail struct {
	mail COM_Mail //0
}
type SGE_World2DB_insertGuild struct {
	guild       COM_Guild       //0
	guildMember COM_GuildMember //1
}
type SGE_World2DB_updateGuildRequestList struct {
	guildId uint32                 //0
	data    []COM_GuildRequestData //1
}
type SGE_World2DB_createGuildMember struct {
	guildMember COM_GuildMember //0
}
type SGE_World2DB_delGuild struct {
	guildId int32 //0
}
type SGE_World2DB_updateGuildNotice struct {
	guildId uint32 //0
	notice  string //1
}
type SGE_World2DB_updateGuild struct {
	guild COM_Guild //0
}
type SGE_World2DB_updateMemberPosition struct {
	roleId int32 //0
	job    int   //1
}
type SGE_World2DB_updateMemberContribution struct {
	roleId       int32 //0
	contribution int32 //1
}
type SGE_World2DB_updateGuildStruction struct {
	guildId   uint32 //0
	level     int8   //1
	struction int32  //2
}
type SGE_World2DB_deleteGuildMember struct {
	playerId int32 //0
}
type SGE_World2DB_queryIdgen struct {
	playerName string //0
	idgen      string //1
}
type SGE_World2DB_updateKeyGift struct {
	giftdata COM_KeyContent //0
}
type SGE_World2DB_insertActivity struct {
	adt  int             //0
	date SGE_SysActivity //1
}
type SGE_World2DB_insertLoseCharge struct {
	playerId int32         //0
	order    SGE_OrderInfo //1
}
type SGE_World2DB_insertEmployeeQuest struct {
	playerId uint32                  //0
	data     SGE_PlayerEmployeeQuest //1
}
type SGE_World2DB_delEmployeeQuest struct {
	playerId uint32 //0
}
type SGE_World2DBStub struct {
	Sender prpc.StubSender
}
type SGE_World2DBProxy interface {
	queryPlayer(username string, serverId int32) error                         // 0
	createPlayer(username string, inst SGE_DBPlayerData, serverId int32) error // 1
	updatePlayer(username string, inst SGE_DBPlayerData) error                 // 2
	deletePlayer(playername string) error                                      // 3
	queryPlayerById(name string, instId int32, isFlag bool) error              // 4
	insertEndlessStair(rank int32, name string) error                          // 5
	updateEndlessStair(rank int32, name string) error                          // 6
	createBaby(playername string, inst COM_BabyInst, isToStorage bool) error   // 7
	deleteBaby(playername string, babyInstId int32) error                      // 8
	resetBabyOwner(playername string, babyInstId int32) error                  // 9
	updateBaby(inst COM_BabyInst) error                                        // 10
	updateBabys(playername string, babys []COM_BabyInst) error                 // 11
	queryBabyById(name string, instid uint32) error                            // 12
	createEmployee(playername string, inst COM_EmployeeInst) error             // 13
	deleteEmployee(playername string, instIds []uint32) error                  // 14
	updateEmployee(inst COM_EmployeeInst) error                                // 15
	queryEmployeeById(name string, instid uint32) error                        // 16
	insertMail(mail COM_Mail) error                                            // 17
	insertMailAll(mail COM_Mail) error                                         // 18
	insertMailByRecvs(mail COM_Mail, recvs []string) error                     // 19
	fatchMail(recvName string, mailId int32) error                             // 20
	delMail(recvName string, mailId int32) error                               // 21
	updateMail(mail COM_Mail) error                                            // 22
	insertGuild(guild COM_Guild, guildMember COM_GuildMember) error            // 23
	updateGuildRequestList(guildId uint32, data []COM_GuildRequestData) error  // 24
	createGuildMember(guildMember COM_GuildMember) error                       // 25
	delGuild(guildId int32) error                                              // 26
	updateGuildNotice(guildId uint32, notice string) error                     // 27
	updateGuild(guild COM_Guild) error                                         // 28
	updateMemberPosition(roleId int32, job int) error                          // 29
	updateMemberContribution(roleId int32, contribution int32) error           // 30
	updateGuildStruction(guildId uint32, level int8, struction int32) error    // 31
	deleteGuildMember(playerId int32) error                                    // 32
	queryIdgen(playerName string, idgen string) error                          // 33
	updateKeyGift(giftdata COM_KeyContent) error                               // 34
	insertActivity(adt int, date SGE_SysActivity) error                        // 35
	insertLoseCharge(playerId int32, order SGE_OrderInfo) error                // 36
	insertEmployeeQuest(playerId uint32, data SGE_PlayerEmployeeQuest) error   // 37
	delEmployeeQuest(playerId uint32) error                                    // 38
}

func (this *SGE_World2DB_queryPlayer) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.username) != 0)
	mask.WriteBit(this.serverId != 0)
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
	// serialize serverId
	{
		if this.serverId != 0 {
			err := prpc.Write(buffer, this.serverId)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *SGE_World2DB_queryPlayer) Deserialize(buffer *bytes.Buffer) error {
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
	// deserialize serverId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.serverId)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_World2DB_createPlayer) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.username) != 0)
	mask.WriteBit(true) //inst
	mask.WriteBit(this.serverId != 0)
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
	// serialize serverId
	{
		if this.serverId != 0 {
			err := prpc.Write(buffer, this.serverId)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *SGE_World2DB_createPlayer) Deserialize(buffer *bytes.Buffer) error {
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
	// deserialize serverId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.serverId)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_World2DB_updatePlayer) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.username) != 0)
	mask.WriteBit(true) //inst
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
	return nil
}
func (this *SGE_World2DB_updatePlayer) Deserialize(buffer *bytes.Buffer) error {
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
	return nil
}
func (this *SGE_World2DB_deletePlayer) Serialize(buffer *bytes.Buffer) error {
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
func (this *SGE_World2DB_deletePlayer) Deserialize(buffer *bytes.Buffer) error {
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
func (this *SGE_World2DB_queryPlayerById) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.name) != 0)
	mask.WriteBit(this.instId != 0)
	mask.WriteBit(this.isFlag)
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
	// serialize instId
	{
		if this.instId != 0 {
			err := prpc.Write(buffer, this.instId)
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
func (this *SGE_World2DB_queryPlayerById) Deserialize(buffer *bytes.Buffer) error {
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
	// deserialize instId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.instId)
		if err != nil {
			return err
		}
	}
	// deserialize isFlag
	this.isFlag = mask.ReadBit()
	return nil
}
func (this *SGE_World2DB_insertEndlessStair) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.rank != 0)
	mask.WriteBit(len(this.name) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize rank
	{
		if this.rank != 0 {
			err := prpc.Write(buffer, this.rank)
			if err != nil {
				return err
			}
		}
	}
	// serialize name
	if len(this.name) != 0 {
		err := prpc.Write(buffer, this.name)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_World2DB_insertEndlessStair) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize rank
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.rank)
		if err != nil {
			return err
		}
	}
	// deserialize name
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.name)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_World2DB_updateEndlessStair) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.rank != 0)
	mask.WriteBit(len(this.name) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize rank
	{
		if this.rank != 0 {
			err := prpc.Write(buffer, this.rank)
			if err != nil {
				return err
			}
		}
	}
	// serialize name
	if len(this.name) != 0 {
		err := prpc.Write(buffer, this.name)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_World2DB_updateEndlessStair) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize rank
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.rank)
		if err != nil {
			return err
		}
	}
	// deserialize name
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.name)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_World2DB_createBaby) Serialize(buffer *bytes.Buffer) error {
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
func (this *SGE_World2DB_createBaby) Deserialize(buffer *bytes.Buffer) error {
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
func (this *SGE_World2DB_deleteBaby) Serialize(buffer *bytes.Buffer) error {
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
func (this *SGE_World2DB_deleteBaby) Deserialize(buffer *bytes.Buffer) error {
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
func (this *SGE_World2DB_resetBabyOwner) Serialize(buffer *bytes.Buffer) error {
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
func (this *SGE_World2DB_resetBabyOwner) Deserialize(buffer *bytes.Buffer) error {
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
func (this *SGE_World2DB_updateBaby) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //inst
	{
		err := prpc.Write(buffer, mask.Bytes())
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
func (this *SGE_World2DB_updateBaby) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
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
func (this *SGE_World2DB_updateBabys) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.playername) != 0)
	mask.WriteBit(len(this.babys) != 0)
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
	// serialize babys
	if len(this.babys) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.babys)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.babys {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *SGE_World2DB_updateBabys) Deserialize(buffer *bytes.Buffer) error {
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
	// deserialize babys
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.babys = make([]COM_BabyInst, size)
		for i, _ := range this.babys {
			err := this.babys[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *SGE_World2DB_queryBabyById) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.name) != 0)
	mask.WriteBit(this.instid != 0)
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
	// serialize instid
	{
		if this.instid != 0 {
			err := prpc.Write(buffer, this.instid)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *SGE_World2DB_queryBabyById) Deserialize(buffer *bytes.Buffer) error {
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
	// deserialize instid
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.instid)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_World2DB_createEmployee) Serialize(buffer *bytes.Buffer) error {
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
func (this *SGE_World2DB_createEmployee) Deserialize(buffer *bytes.Buffer) error {
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
func (this *SGE_World2DB_deleteEmployee) Serialize(buffer *bytes.Buffer) error {
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
func (this *SGE_World2DB_deleteEmployee) Deserialize(buffer *bytes.Buffer) error {
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
func (this *SGE_World2DB_updateEmployee) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //inst
	{
		err := prpc.Write(buffer, mask.Bytes())
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
func (this *SGE_World2DB_updateEmployee) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
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
func (this *SGE_World2DB_queryEmployeeById) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.name) != 0)
	mask.WriteBit(this.instid != 0)
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
	// serialize instid
	{
		if this.instid != 0 {
			err := prpc.Write(buffer, this.instid)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *SGE_World2DB_queryEmployeeById) Deserialize(buffer *bytes.Buffer) error {
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
	// deserialize instid
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.instid)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_World2DB_insertMail) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //mail
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize mail
	{
		err := this.mail.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_World2DB_insertMail) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize mail
	if mask.ReadBit() {
		err := this.mail.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_World2DB_insertMailAll) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //mail
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize mail
	{
		err := this.mail.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_World2DB_insertMailAll) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize mail
	if mask.ReadBit() {
		err := this.mail.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_World2DB_insertMailByRecvs) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //mail
	mask.WriteBit(len(this.recvs) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize mail
	{
		err := this.mail.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	// serialize recvs
	if len(this.recvs) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.recvs)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.recvs {
			err := prpc.Write(buffer, value)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *SGE_World2DB_insertMailByRecvs) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize mail
	if mask.ReadBit() {
		err := this.mail.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	// deserialize recvs
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.recvs = make([]string, size)
		for i, _ := range this.recvs {
			err := prpc.Read(buffer, &this.recvs[i])
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *SGE_World2DB_fatchMail) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.recvName) != 0)
	mask.WriteBit(this.mailId != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize recvName
	if len(this.recvName) != 0 {
		err := prpc.Write(buffer, this.recvName)
		if err != nil {
			return err
		}
	}
	// serialize mailId
	{
		if this.mailId != 0 {
			err := prpc.Write(buffer, this.mailId)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *SGE_World2DB_fatchMail) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize recvName
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.recvName)
		if err != nil {
			return err
		}
	}
	// deserialize mailId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.mailId)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_World2DB_delMail) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.recvName) != 0)
	mask.WriteBit(this.mailId != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize recvName
	if len(this.recvName) != 0 {
		err := prpc.Write(buffer, this.recvName)
		if err != nil {
			return err
		}
	}
	// serialize mailId
	{
		if this.mailId != 0 {
			err := prpc.Write(buffer, this.mailId)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *SGE_World2DB_delMail) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize recvName
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.recvName)
		if err != nil {
			return err
		}
	}
	// deserialize mailId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.mailId)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_World2DB_updateMail) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //mail
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize mail
	{
		err := this.mail.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_World2DB_updateMail) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize mail
	if mask.ReadBit() {
		err := this.mail.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_World2DB_insertGuild) Serialize(buffer *bytes.Buffer) error {
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
func (this *SGE_World2DB_insertGuild) Deserialize(buffer *bytes.Buffer) error {
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
func (this *SGE_World2DB_updateGuildRequestList) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.guildId != 0)
	mask.WriteBit(len(this.data) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize guildId
	{
		if this.guildId != 0 {
			err := prpc.Write(buffer, this.guildId)
			if err != nil {
				return err
			}
		}
	}
	// serialize data
	if len(this.data) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.data)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.data {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *SGE_World2DB_updateGuildRequestList) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize guildId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.guildId)
		if err != nil {
			return err
		}
	}
	// deserialize data
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.data = make([]COM_GuildRequestData, size)
		for i, _ := range this.data {
			err := this.data[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *SGE_World2DB_createGuildMember) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //guildMember
	{
		err := prpc.Write(buffer, mask.Bytes())
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
func (this *SGE_World2DB_createGuildMember) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
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
func (this *SGE_World2DB_delGuild) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.guildId != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize guildId
	{
		if this.guildId != 0 {
			err := prpc.Write(buffer, this.guildId)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *SGE_World2DB_delGuild) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize guildId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.guildId)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_World2DB_updateGuildNotice) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.guildId != 0)
	mask.WriteBit(len(this.notice) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize guildId
	{
		if this.guildId != 0 {
			err := prpc.Write(buffer, this.guildId)
			if err != nil {
				return err
			}
		}
	}
	// serialize notice
	if len(this.notice) != 0 {
		err := prpc.Write(buffer, this.notice)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_World2DB_updateGuildNotice) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize guildId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.guildId)
		if err != nil {
			return err
		}
	}
	// deserialize notice
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.notice)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_World2DB_updateGuild) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //guild
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
	return nil
}
func (this *SGE_World2DB_updateGuild) Deserialize(buffer *bytes.Buffer) error {
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
	return nil
}
func (this *SGE_World2DB_updateMemberPosition) Serialize(buffer *bytes.Buffer) error {
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
func (this *SGE_World2DB_updateMemberPosition) Deserialize(buffer *bytes.Buffer) error {
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
func (this *SGE_World2DB_updateMemberContribution) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.roleId != 0)
	mask.WriteBit(this.contribution != 0)
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
	// serialize contribution
	{
		if this.contribution != 0 {
			err := prpc.Write(buffer, this.contribution)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *SGE_World2DB_updateMemberContribution) Deserialize(buffer *bytes.Buffer) error {
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
	// deserialize contribution
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.contribution)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_World2DB_updateGuildStruction) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.guildId != 0)
	mask.WriteBit(this.level != 0)
	mask.WriteBit(this.struction != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize guildId
	{
		if this.guildId != 0 {
			err := prpc.Write(buffer, this.guildId)
			if err != nil {
				return err
			}
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
	// serialize struction
	{
		if this.struction != 0 {
			err := prpc.Write(buffer, this.struction)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *SGE_World2DB_updateGuildStruction) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize guildId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.guildId)
		if err != nil {
			return err
		}
	}
	// deserialize level
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.level)
		if err != nil {
			return err
		}
	}
	// deserialize struction
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.struction)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_World2DB_deleteGuildMember) Serialize(buffer *bytes.Buffer) error {
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
func (this *SGE_World2DB_deleteGuildMember) Deserialize(buffer *bytes.Buffer) error {
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
func (this *SGE_World2DB_queryIdgen) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.playerName) != 0)
	mask.WriteBit(len(this.idgen) != 0)
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
	// serialize idgen
	if len(this.idgen) != 0 {
		err := prpc.Write(buffer, this.idgen)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_World2DB_queryIdgen) Deserialize(buffer *bytes.Buffer) error {
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
	// deserialize idgen
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.idgen)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_World2DB_updateKeyGift) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //giftdata
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize giftdata
	{
		err := this.giftdata.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_World2DB_updateKeyGift) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize giftdata
	if mask.ReadBit() {
		err := this.giftdata.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_World2DB_insertActivity) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.adt != 0)
	mask.WriteBit(true) //date
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize adt
	{
		if this.adt != 0 {
			err := prpc.Write(buffer, this.adt)
			if err != nil {
				return err
			}
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
func (this *SGE_World2DB_insertActivity) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize adt
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.adt)
		if err != nil {
			return err
		}
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
func (this *SGE_World2DB_insertLoseCharge) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.playerId != 0)
	mask.WriteBit(true) //order
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
	// serialize order
	{
		err := this.order.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_World2DB_insertLoseCharge) Deserialize(buffer *bytes.Buffer) error {
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
	// deserialize order
	if mask.ReadBit() {
		err := this.order.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_World2DB_insertEmployeeQuest) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.playerId != 0)
	mask.WriteBit(true) //data
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
	// serialize data
	{
		err := this.data.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_World2DB_insertEmployeeQuest) Deserialize(buffer *bytes.Buffer) error {
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
	// deserialize data
	if mask.ReadBit() {
		err := this.data.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_World2DB_delEmployeeQuest) Serialize(buffer *bytes.Buffer) error {
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
func (this *SGE_World2DB_delEmployeeQuest) Deserialize(buffer *bytes.Buffer) error {
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
func (this *SGE_World2DBStub) queryPlayer(username string, serverId int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(0))
	if err != nil {
		return err
	}
	_0 := SGE_World2DB_queryPlayer{}
	_0.username = username
	_0.serverId = serverId
	err = _0.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_World2DBStub) createPlayer(username string, inst SGE_DBPlayerData, serverId int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(1))
	if err != nil {
		return err
	}
	_1 := SGE_World2DB_createPlayer{}
	_1.username = username
	_1.inst = inst
	_1.serverId = serverId
	err = _1.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_World2DBStub) updatePlayer(username string, inst SGE_DBPlayerData) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(2))
	if err != nil {
		return err
	}
	_2 := SGE_World2DB_updatePlayer{}
	_2.username = username
	_2.inst = inst
	err = _2.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_World2DBStub) deletePlayer(playername string) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(3))
	if err != nil {
		return err
	}
	_3 := SGE_World2DB_deletePlayer{}
	_3.playername = playername
	err = _3.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_World2DBStub) queryPlayerById(name string, instId int32, isFlag bool) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(4))
	if err != nil {
		return err
	}
	_4 := SGE_World2DB_queryPlayerById{}
	_4.name = name
	_4.instId = instId
	_4.isFlag = isFlag
	err = _4.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_World2DBStub) insertEndlessStair(rank int32, name string) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(5))
	if err != nil {
		return err
	}
	_5 := SGE_World2DB_insertEndlessStair{}
	_5.rank = rank
	_5.name = name
	err = _5.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_World2DBStub) updateEndlessStair(rank int32, name string) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(6))
	if err != nil {
		return err
	}
	_6 := SGE_World2DB_updateEndlessStair{}
	_6.rank = rank
	_6.name = name
	err = _6.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_World2DBStub) createBaby(playername string, inst COM_BabyInst, isToStorage bool) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(7))
	if err != nil {
		return err
	}
	_7 := SGE_World2DB_createBaby{}
	_7.playername = playername
	_7.inst = inst
	_7.isToStorage = isToStorage
	err = _7.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_World2DBStub) deleteBaby(playername string, babyInstId int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(8))
	if err != nil {
		return err
	}
	_8 := SGE_World2DB_deleteBaby{}
	_8.playername = playername
	_8.babyInstId = babyInstId
	err = _8.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_World2DBStub) resetBabyOwner(playername string, babyInstId int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(9))
	if err != nil {
		return err
	}
	_9 := SGE_World2DB_resetBabyOwner{}
	_9.playername = playername
	_9.babyInstId = babyInstId
	err = _9.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_World2DBStub) updateBaby(inst COM_BabyInst) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(10))
	if err != nil {
		return err
	}
	_10 := SGE_World2DB_updateBaby{}
	_10.inst = inst
	err = _10.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_World2DBStub) updateBabys(playername string, babys []COM_BabyInst) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(11))
	if err != nil {
		return err
	}
	_11 := SGE_World2DB_updateBabys{}
	_11.playername = playername
	_11.babys = babys
	err = _11.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_World2DBStub) queryBabyById(name string, instid uint32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(12))
	if err != nil {
		return err
	}
	_12 := SGE_World2DB_queryBabyById{}
	_12.name = name
	_12.instid = instid
	err = _12.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_World2DBStub) createEmployee(playername string, inst COM_EmployeeInst) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(13))
	if err != nil {
		return err
	}
	_13 := SGE_World2DB_createEmployee{}
	_13.playername = playername
	_13.inst = inst
	err = _13.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_World2DBStub) deleteEmployee(playername string, instIds []uint32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(14))
	if err != nil {
		return err
	}
	_14 := SGE_World2DB_deleteEmployee{}
	_14.playername = playername
	_14.instIds = instIds
	err = _14.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_World2DBStub) updateEmployee(inst COM_EmployeeInst) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(15))
	if err != nil {
		return err
	}
	_15 := SGE_World2DB_updateEmployee{}
	_15.inst = inst
	err = _15.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_World2DBStub) queryEmployeeById(name string, instid uint32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(16))
	if err != nil {
		return err
	}
	_16 := SGE_World2DB_queryEmployeeById{}
	_16.name = name
	_16.instid = instid
	err = _16.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_World2DBStub) insertMail(mail COM_Mail) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(17))
	if err != nil {
		return err
	}
	_17 := SGE_World2DB_insertMail{}
	_17.mail = mail
	err = _17.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_World2DBStub) insertMailAll(mail COM_Mail) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(18))
	if err != nil {
		return err
	}
	_18 := SGE_World2DB_insertMailAll{}
	_18.mail = mail
	err = _18.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_World2DBStub) insertMailByRecvs(mail COM_Mail, recvs []string) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(19))
	if err != nil {
		return err
	}
	_19 := SGE_World2DB_insertMailByRecvs{}
	_19.mail = mail
	_19.recvs = recvs
	err = _19.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_World2DBStub) fatchMail(recvName string, mailId int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(20))
	if err != nil {
		return err
	}
	_20 := SGE_World2DB_fatchMail{}
	_20.recvName = recvName
	_20.mailId = mailId
	err = _20.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_World2DBStub) delMail(recvName string, mailId int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(21))
	if err != nil {
		return err
	}
	_21 := SGE_World2DB_delMail{}
	_21.recvName = recvName
	_21.mailId = mailId
	err = _21.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_World2DBStub) updateMail(mail COM_Mail) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(22))
	if err != nil {
		return err
	}
	_22 := SGE_World2DB_updateMail{}
	_22.mail = mail
	err = _22.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_World2DBStub) insertGuild(guild COM_Guild, guildMember COM_GuildMember) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(23))
	if err != nil {
		return err
	}
	_23 := SGE_World2DB_insertGuild{}
	_23.guild = guild
	_23.guildMember = guildMember
	err = _23.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_World2DBStub) updateGuildRequestList(guildId uint32, data []COM_GuildRequestData) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(24))
	if err != nil {
		return err
	}
	_24 := SGE_World2DB_updateGuildRequestList{}
	_24.guildId = guildId
	_24.data = data
	err = _24.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_World2DBStub) createGuildMember(guildMember COM_GuildMember) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(25))
	if err != nil {
		return err
	}
	_25 := SGE_World2DB_createGuildMember{}
	_25.guildMember = guildMember
	err = _25.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_World2DBStub) delGuild(guildId int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(26))
	if err != nil {
		return err
	}
	_26 := SGE_World2DB_delGuild{}
	_26.guildId = guildId
	err = _26.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_World2DBStub) updateGuildNotice(guildId uint32, notice string) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(27))
	if err != nil {
		return err
	}
	_27 := SGE_World2DB_updateGuildNotice{}
	_27.guildId = guildId
	_27.notice = notice
	err = _27.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_World2DBStub) updateGuild(guild COM_Guild) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(28))
	if err != nil {
		return err
	}
	_28 := SGE_World2DB_updateGuild{}
	_28.guild = guild
	err = _28.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_World2DBStub) updateMemberPosition(roleId int32, job int) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(29))
	if err != nil {
		return err
	}
	_29 := SGE_World2DB_updateMemberPosition{}
	_29.roleId = roleId
	_29.job = job
	err = _29.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_World2DBStub) updateMemberContribution(roleId int32, contribution int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(30))
	if err != nil {
		return err
	}
	_30 := SGE_World2DB_updateMemberContribution{}
	_30.roleId = roleId
	_30.contribution = contribution
	err = _30.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_World2DBStub) updateGuildStruction(guildId uint32, level int8, struction int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(31))
	if err != nil {
		return err
	}
	_31 := SGE_World2DB_updateGuildStruction{}
	_31.guildId = guildId
	_31.level = level
	_31.struction = struction
	err = _31.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_World2DBStub) deleteGuildMember(playerId int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(32))
	if err != nil {
		return err
	}
	_32 := SGE_World2DB_deleteGuildMember{}
	_32.playerId = playerId
	err = _32.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_World2DBStub) queryIdgen(playerName string, idgen string) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(33))
	if err != nil {
		return err
	}
	_33 := SGE_World2DB_queryIdgen{}
	_33.playerName = playerName
	_33.idgen = idgen
	err = _33.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_World2DBStub) updateKeyGift(giftdata COM_KeyContent) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(34))
	if err != nil {
		return err
	}
	_34 := SGE_World2DB_updateKeyGift{}
	_34.giftdata = giftdata
	err = _34.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_World2DBStub) insertActivity(adt int, date SGE_SysActivity) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(35))
	if err != nil {
		return err
	}
	_35 := SGE_World2DB_insertActivity{}
	_35.adt = adt
	_35.date = date
	err = _35.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_World2DBStub) insertLoseCharge(playerId int32, order SGE_OrderInfo) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(36))
	if err != nil {
		return err
	}
	_36 := SGE_World2DB_insertLoseCharge{}
	_36.playerId = playerId
	_36.order = order
	err = _36.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_World2DBStub) insertEmployeeQuest(playerId uint32, data SGE_PlayerEmployeeQuest) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(37))
	if err != nil {
		return err
	}
	_37 := SGE_World2DB_insertEmployeeQuest{}
	_37.playerId = playerId
	_37.data = data
	err = _37.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_World2DBStub) delEmployeeQuest(playerId uint32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(38))
	if err != nil {
		return err
	}
	_38 := SGE_World2DB_delEmployeeQuest{}
	_38.playerId = playerId
	err = _38.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func Bridging_SGE_World2DB_queryPlayer(buffer *bytes.Buffer, p SGE_World2DBProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_0 := SGE_World2DB_queryPlayer{}
	err := _0.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.queryPlayer(_0.username, _0.serverId)
}
func Bridging_SGE_World2DB_createPlayer(buffer *bytes.Buffer, p SGE_World2DBProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_1 := SGE_World2DB_createPlayer{}
	err := _1.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.createPlayer(_1.username, _1.inst, _1.serverId)
}
func Bridging_SGE_World2DB_updatePlayer(buffer *bytes.Buffer, p SGE_World2DBProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_2 := SGE_World2DB_updatePlayer{}
	err := _2.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.updatePlayer(_2.username, _2.inst)
}
func Bridging_SGE_World2DB_deletePlayer(buffer *bytes.Buffer, p SGE_World2DBProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_3 := SGE_World2DB_deletePlayer{}
	err := _3.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.deletePlayer(_3.playername)
}
func Bridging_SGE_World2DB_queryPlayerById(buffer *bytes.Buffer, p SGE_World2DBProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_4 := SGE_World2DB_queryPlayerById{}
	err := _4.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.queryPlayerById(_4.name, _4.instId, _4.isFlag)
}
func Bridging_SGE_World2DB_insertEndlessStair(buffer *bytes.Buffer, p SGE_World2DBProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_5 := SGE_World2DB_insertEndlessStair{}
	err := _5.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.insertEndlessStair(_5.rank, _5.name)
}
func Bridging_SGE_World2DB_updateEndlessStair(buffer *bytes.Buffer, p SGE_World2DBProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_6 := SGE_World2DB_updateEndlessStair{}
	err := _6.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.updateEndlessStair(_6.rank, _6.name)
}
func Bridging_SGE_World2DB_createBaby(buffer *bytes.Buffer, p SGE_World2DBProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_7 := SGE_World2DB_createBaby{}
	err := _7.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.createBaby(_7.playername, _7.inst, _7.isToStorage)
}
func Bridging_SGE_World2DB_deleteBaby(buffer *bytes.Buffer, p SGE_World2DBProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_8 := SGE_World2DB_deleteBaby{}
	err := _8.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.deleteBaby(_8.playername, _8.babyInstId)
}
func Bridging_SGE_World2DB_resetBabyOwner(buffer *bytes.Buffer, p SGE_World2DBProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_9 := SGE_World2DB_resetBabyOwner{}
	err := _9.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.resetBabyOwner(_9.playername, _9.babyInstId)
}
func Bridging_SGE_World2DB_updateBaby(buffer *bytes.Buffer, p SGE_World2DBProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_10 := SGE_World2DB_updateBaby{}
	err := _10.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.updateBaby(_10.inst)
}
func Bridging_SGE_World2DB_updateBabys(buffer *bytes.Buffer, p SGE_World2DBProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_11 := SGE_World2DB_updateBabys{}
	err := _11.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.updateBabys(_11.playername, _11.babys)
}
func Bridging_SGE_World2DB_queryBabyById(buffer *bytes.Buffer, p SGE_World2DBProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_12 := SGE_World2DB_queryBabyById{}
	err := _12.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.queryBabyById(_12.name, _12.instid)
}
func Bridging_SGE_World2DB_createEmployee(buffer *bytes.Buffer, p SGE_World2DBProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_13 := SGE_World2DB_createEmployee{}
	err := _13.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.createEmployee(_13.playername, _13.inst)
}
func Bridging_SGE_World2DB_deleteEmployee(buffer *bytes.Buffer, p SGE_World2DBProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_14 := SGE_World2DB_deleteEmployee{}
	err := _14.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.deleteEmployee(_14.playername, _14.instIds)
}
func Bridging_SGE_World2DB_updateEmployee(buffer *bytes.Buffer, p SGE_World2DBProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_15 := SGE_World2DB_updateEmployee{}
	err := _15.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.updateEmployee(_15.inst)
}
func Bridging_SGE_World2DB_queryEmployeeById(buffer *bytes.Buffer, p SGE_World2DBProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_16 := SGE_World2DB_queryEmployeeById{}
	err := _16.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.queryEmployeeById(_16.name, _16.instid)
}
func Bridging_SGE_World2DB_insertMail(buffer *bytes.Buffer, p SGE_World2DBProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_17 := SGE_World2DB_insertMail{}
	err := _17.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.insertMail(_17.mail)
}
func Bridging_SGE_World2DB_insertMailAll(buffer *bytes.Buffer, p SGE_World2DBProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_18 := SGE_World2DB_insertMailAll{}
	err := _18.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.insertMailAll(_18.mail)
}
func Bridging_SGE_World2DB_insertMailByRecvs(buffer *bytes.Buffer, p SGE_World2DBProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_19 := SGE_World2DB_insertMailByRecvs{}
	err := _19.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.insertMailByRecvs(_19.mail, _19.recvs)
}
func Bridging_SGE_World2DB_fatchMail(buffer *bytes.Buffer, p SGE_World2DBProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_20 := SGE_World2DB_fatchMail{}
	err := _20.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.fatchMail(_20.recvName, _20.mailId)
}
func Bridging_SGE_World2DB_delMail(buffer *bytes.Buffer, p SGE_World2DBProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_21 := SGE_World2DB_delMail{}
	err := _21.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.delMail(_21.recvName, _21.mailId)
}
func Bridging_SGE_World2DB_updateMail(buffer *bytes.Buffer, p SGE_World2DBProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_22 := SGE_World2DB_updateMail{}
	err := _22.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.updateMail(_22.mail)
}
func Bridging_SGE_World2DB_insertGuild(buffer *bytes.Buffer, p SGE_World2DBProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_23 := SGE_World2DB_insertGuild{}
	err := _23.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.insertGuild(_23.guild, _23.guildMember)
}
func Bridging_SGE_World2DB_updateGuildRequestList(buffer *bytes.Buffer, p SGE_World2DBProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_24 := SGE_World2DB_updateGuildRequestList{}
	err := _24.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.updateGuildRequestList(_24.guildId, _24.data)
}
func Bridging_SGE_World2DB_createGuildMember(buffer *bytes.Buffer, p SGE_World2DBProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_25 := SGE_World2DB_createGuildMember{}
	err := _25.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.createGuildMember(_25.guildMember)
}
func Bridging_SGE_World2DB_delGuild(buffer *bytes.Buffer, p SGE_World2DBProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_26 := SGE_World2DB_delGuild{}
	err := _26.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.delGuild(_26.guildId)
}
func Bridging_SGE_World2DB_updateGuildNotice(buffer *bytes.Buffer, p SGE_World2DBProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_27 := SGE_World2DB_updateGuildNotice{}
	err := _27.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.updateGuildNotice(_27.guildId, _27.notice)
}
func Bridging_SGE_World2DB_updateGuild(buffer *bytes.Buffer, p SGE_World2DBProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_28 := SGE_World2DB_updateGuild{}
	err := _28.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.updateGuild(_28.guild)
}
func Bridging_SGE_World2DB_updateMemberPosition(buffer *bytes.Buffer, p SGE_World2DBProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_29 := SGE_World2DB_updateMemberPosition{}
	err := _29.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.updateMemberPosition(_29.roleId, _29.job)
}
func Bridging_SGE_World2DB_updateMemberContribution(buffer *bytes.Buffer, p SGE_World2DBProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_30 := SGE_World2DB_updateMemberContribution{}
	err := _30.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.updateMemberContribution(_30.roleId, _30.contribution)
}
func Bridging_SGE_World2DB_updateGuildStruction(buffer *bytes.Buffer, p SGE_World2DBProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_31 := SGE_World2DB_updateGuildStruction{}
	err := _31.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.updateGuildStruction(_31.guildId, _31.level, _31.struction)
}
func Bridging_SGE_World2DB_deleteGuildMember(buffer *bytes.Buffer, p SGE_World2DBProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_32 := SGE_World2DB_deleteGuildMember{}
	err := _32.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.deleteGuildMember(_32.playerId)
}
func Bridging_SGE_World2DB_queryIdgen(buffer *bytes.Buffer, p SGE_World2DBProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_33 := SGE_World2DB_queryIdgen{}
	err := _33.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.queryIdgen(_33.playerName, _33.idgen)
}
func Bridging_SGE_World2DB_updateKeyGift(buffer *bytes.Buffer, p SGE_World2DBProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_34 := SGE_World2DB_updateKeyGift{}
	err := _34.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.updateKeyGift(_34.giftdata)
}
func Bridging_SGE_World2DB_insertActivity(buffer *bytes.Buffer, p SGE_World2DBProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_35 := SGE_World2DB_insertActivity{}
	err := _35.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.insertActivity(_35.adt, _35.date)
}
func Bridging_SGE_World2DB_insertLoseCharge(buffer *bytes.Buffer, p SGE_World2DBProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_36 := SGE_World2DB_insertLoseCharge{}
	err := _36.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.insertLoseCharge(_36.playerId, _36.order)
}
func Bridging_SGE_World2DB_insertEmployeeQuest(buffer *bytes.Buffer, p SGE_World2DBProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_37 := SGE_World2DB_insertEmployeeQuest{}
	err := _37.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.insertEmployeeQuest(_37.playerId, _37.data)
}
func Bridging_SGE_World2DB_delEmployeeQuest(buffer *bytes.Buffer, p SGE_World2DBProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_38 := SGE_World2DB_delEmployeeQuest{}
	err := _38.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.delEmployeeQuest(_38.playerId)
}
func SGE_World2DBDispatch(buffer *bytes.Buffer, p SGE_World2DBProxy) error {
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
		return Bridging_SGE_World2DB_queryPlayer(buffer, p)
	case 1:
		return Bridging_SGE_World2DB_createPlayer(buffer, p)
	case 2:
		return Bridging_SGE_World2DB_updatePlayer(buffer, p)
	case 3:
		return Bridging_SGE_World2DB_deletePlayer(buffer, p)
	case 4:
		return Bridging_SGE_World2DB_queryPlayerById(buffer, p)
	case 5:
		return Bridging_SGE_World2DB_insertEndlessStair(buffer, p)
	case 6:
		return Bridging_SGE_World2DB_updateEndlessStair(buffer, p)
	case 7:
		return Bridging_SGE_World2DB_createBaby(buffer, p)
	case 8:
		return Bridging_SGE_World2DB_deleteBaby(buffer, p)
	case 9:
		return Bridging_SGE_World2DB_resetBabyOwner(buffer, p)
	case 10:
		return Bridging_SGE_World2DB_updateBaby(buffer, p)
	case 11:
		return Bridging_SGE_World2DB_updateBabys(buffer, p)
	case 12:
		return Bridging_SGE_World2DB_queryBabyById(buffer, p)
	case 13:
		return Bridging_SGE_World2DB_createEmployee(buffer, p)
	case 14:
		return Bridging_SGE_World2DB_deleteEmployee(buffer, p)
	case 15:
		return Bridging_SGE_World2DB_updateEmployee(buffer, p)
	case 16:
		return Bridging_SGE_World2DB_queryEmployeeById(buffer, p)
	case 17:
		return Bridging_SGE_World2DB_insertMail(buffer, p)
	case 18:
		return Bridging_SGE_World2DB_insertMailAll(buffer, p)
	case 19:
		return Bridging_SGE_World2DB_insertMailByRecvs(buffer, p)
	case 20:
		return Bridging_SGE_World2DB_fatchMail(buffer, p)
	case 21:
		return Bridging_SGE_World2DB_delMail(buffer, p)
	case 22:
		return Bridging_SGE_World2DB_updateMail(buffer, p)
	case 23:
		return Bridging_SGE_World2DB_insertGuild(buffer, p)
	case 24:
		return Bridging_SGE_World2DB_updateGuildRequestList(buffer, p)
	case 25:
		return Bridging_SGE_World2DB_createGuildMember(buffer, p)
	case 26:
		return Bridging_SGE_World2DB_delGuild(buffer, p)
	case 27:
		return Bridging_SGE_World2DB_updateGuildNotice(buffer, p)
	case 28:
		return Bridging_SGE_World2DB_updateGuild(buffer, p)
	case 29:
		return Bridging_SGE_World2DB_updateMemberPosition(buffer, p)
	case 30:
		return Bridging_SGE_World2DB_updateMemberContribution(buffer, p)
	case 31:
		return Bridging_SGE_World2DB_updateGuildStruction(buffer, p)
	case 32:
		return Bridging_SGE_World2DB_deleteGuildMember(buffer, p)
	case 33:
		return Bridging_SGE_World2DB_queryIdgen(buffer, p)
	case 34:
		return Bridging_SGE_World2DB_updateKeyGift(buffer, p)
	case 35:
		return Bridging_SGE_World2DB_insertActivity(buffer, p)
	case 36:
		return Bridging_SGE_World2DB_insertLoseCharge(buffer, p)
	case 37:
		return Bridging_SGE_World2DB_insertEmployeeQuest(buffer, p)
	case 38:
		return Bridging_SGE_World2DB_delEmployeeQuest(buffer, p)
	default:
		return errors.New(prpc.NoneDispatchMatchError)
	}
}
