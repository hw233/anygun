package prpc

import (
	"bytes"
	"errors"
	"suzuki/prpc"
)

type SGE_GMT2World_addExp struct {
	playerId uint32 //0
	exp      uint32 //1
}
type SGE_GMT2World_addMoney struct {
	playerId uint32 //0
	money    int32  //1
}
type SGE_GMT2World_addDiamond struct {
	playerId uint32 //0
	diamond  int32  //1
}
type SGE_GMT2World_openGM struct {
	playerId uint32 //0
}
type SGE_GMT2World_closeGM struct {
	playerId uint32 //0
}
type SGE_GMT2World_noTalkPlayer struct {
	playerId uint32 //0
	time     uint32 //1
}
type SGE_GMT2World_sealPlayer struct {
	playerId uint32 //0
}
type SGE_GMT2World_kickPlayer struct {
	playerId uint32 //0
}
type SGE_GMT2World_openTalkPlayer struct {
	playerId uint32 //0
}
type SGE_GMT2World_unsealPlayer struct {
	playerId uint32 //0
}
type SGE_GMT2World_sendMailAllOnline struct {
	mail        COM_Mail //0
	lowLevel    int32    //1
	highLevel   int32    //2
	lowRegTime  int64    //3
	highRegTime int64    //4
}
type SGE_GMT2World_gmtNotice struct {
	bType   int    //0
	note    string //1
	thetime uint64 //2
	itvtime int64  //3
}
type SGE_GMT2World_setChargeTotal struct {
	data COM_ADChargeTotal //0
}
type SGE_GMT2World_setChargeEvery struct {
	data COM_ADChargeEvery //0
}
type SGE_GMT2World_setDiscountStore struct {
	data COM_ADDiscountStore //0
}
type SGE_GMT2World_setLoginTotal struct {
	data COM_ADLoginTotal //0
}
type SGE_GMT2World_setHotRole struct {
	data COM_ADHotRole //0
}
type SGE_GMT2World_setEmployeeActivity struct {
	data COM_ADEmployeeTotal //0
}
type SGE_GMT2World_setMinGiftBagActivity struct {
	data COM_ADGiftBag //0
}
type SGE_GMT2World_setZhuanpanActivity struct {
	data COM_ZhuanpanData //0
}
type SGE_GMT2World_setIntegralshop struct {
	data COM_IntegralData //0
}
type SGE_GMT2World_makeOrder struct {
	playerId uint32       //0
	order    SGE_GmtOrder //1
}
type SGE_GMT2World_doScript struct {
	script string //0
}
type SGE_GMT2World_playerDoScript struct {
	playerId uint32 //0
	script   string //1
}
type SGE_GMT2WorldStub struct {
	Sender prpc.StubSender
}
type SGE_GMT2WorldProxy interface {
	addExp(playerId uint32, exp uint32) error                                                                    // 0
	addMoney(playerId uint32, money int32) error                                                                 // 1
	addDiamond(playerId uint32, diamond int32) error                                                             // 2
	openGM(playerId uint32) error                                                                                // 3
	closeGM(playerId uint32) error                                                                               // 4
	noTalkPlayer(playerId uint32, time uint32) error                                                             // 5
	sealPlayer(playerId uint32) error                                                                            // 6
	kickPlayer(playerId uint32) error                                                                            // 7
	openTalkPlayer(playerId uint32) error                                                                        // 8
	unsealPlayer(playerId uint32) error                                                                          // 9
	sendMailAllOnline(mail COM_Mail, lowLevel int32, highLevel int32, lowRegTime int64, highRegTime int64) error // 10
	gmtNotice(bType int, note string, thetime uint64, itvtime int64) error                                       // 11
	setChargeTotal(data COM_ADChargeTotal) error                                                                 // 12
	setChargeEvery(data COM_ADChargeEvery) error                                                                 // 13
	setDiscountStore(data COM_ADDiscountStore) error                                                             // 14
	setLoginTotal(data COM_ADLoginTotal) error                                                                   // 15
	setHotRole(data COM_ADHotRole) error                                                                         // 16
	setEmployeeActivity(data COM_ADEmployeeTotal) error                                                          // 17
	setMinGiftBagActivity(data COM_ADGiftBag) error                                                              // 18
	setZhuanpanActivity(data COM_ZhuanpanData) error                                                             // 19
	setIntegralshop(data COM_IntegralData) error                                                                 // 20
	makeOrder(playerId uint32, order SGE_GmtOrder) error                                                         // 21
	doScript(script string) error                                                                                // 22
	playerDoScript(playerId uint32, script string) error                                                         // 23
}

func (this *SGE_GMT2World_addExp) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.playerId != 0)
	mask.WriteBit(this.exp != 0)
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
	// serialize exp
	{
		if this.exp != 0 {
			err := prpc.Write(buffer, this.exp)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *SGE_GMT2World_addExp) Deserialize(buffer *bytes.Buffer) error {
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
	// deserialize exp
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.exp)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_GMT2World_addMoney) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.playerId != 0)
	mask.WriteBit(this.money != 0)
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
	// serialize money
	{
		if this.money != 0 {
			err := prpc.Write(buffer, this.money)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *SGE_GMT2World_addMoney) Deserialize(buffer *bytes.Buffer) error {
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
	// deserialize money
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.money)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_GMT2World_addDiamond) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.playerId != 0)
	mask.WriteBit(this.diamond != 0)
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
	// serialize diamond
	{
		if this.diamond != 0 {
			err := prpc.Write(buffer, this.diamond)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *SGE_GMT2World_addDiamond) Deserialize(buffer *bytes.Buffer) error {
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
	// deserialize diamond
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.diamond)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_GMT2World_openGM) Serialize(buffer *bytes.Buffer) error {
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
func (this *SGE_GMT2World_openGM) Deserialize(buffer *bytes.Buffer) error {
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
func (this *SGE_GMT2World_closeGM) Serialize(buffer *bytes.Buffer) error {
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
func (this *SGE_GMT2World_closeGM) Deserialize(buffer *bytes.Buffer) error {
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
func (this *SGE_GMT2World_noTalkPlayer) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.playerId != 0)
	mask.WriteBit(this.time != 0)
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
	// serialize time
	{
		if this.time != 0 {
			err := prpc.Write(buffer, this.time)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *SGE_GMT2World_noTalkPlayer) Deserialize(buffer *bytes.Buffer) error {
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
	// deserialize time
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.time)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_GMT2World_sealPlayer) Serialize(buffer *bytes.Buffer) error {
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
func (this *SGE_GMT2World_sealPlayer) Deserialize(buffer *bytes.Buffer) error {
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
func (this *SGE_GMT2World_kickPlayer) Serialize(buffer *bytes.Buffer) error {
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
func (this *SGE_GMT2World_kickPlayer) Deserialize(buffer *bytes.Buffer) error {
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
func (this *SGE_GMT2World_openTalkPlayer) Serialize(buffer *bytes.Buffer) error {
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
func (this *SGE_GMT2World_openTalkPlayer) Deserialize(buffer *bytes.Buffer) error {
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
func (this *SGE_GMT2World_unsealPlayer) Serialize(buffer *bytes.Buffer) error {
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
func (this *SGE_GMT2World_unsealPlayer) Deserialize(buffer *bytes.Buffer) error {
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
func (this *SGE_GMT2World_sendMailAllOnline) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //mail
	mask.WriteBit(this.lowLevel != 0)
	mask.WriteBit(this.highLevel != 0)
	mask.WriteBit(this.lowRegTime != 0)
	mask.WriteBit(this.highRegTime != 0)
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
	// serialize lowLevel
	{
		if this.lowLevel != 0 {
			err := prpc.Write(buffer, this.lowLevel)
			if err != nil {
				return err
			}
		}
	}
	// serialize highLevel
	{
		if this.highLevel != 0 {
			err := prpc.Write(buffer, this.highLevel)
			if err != nil {
				return err
			}
		}
	}
	// serialize lowRegTime
	{
		if this.lowRegTime != 0 {
			err := prpc.Write(buffer, this.lowRegTime)
			if err != nil {
				return err
			}
		}
	}
	// serialize highRegTime
	{
		if this.highRegTime != 0 {
			err := prpc.Write(buffer, this.highRegTime)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *SGE_GMT2World_sendMailAllOnline) Deserialize(buffer *bytes.Buffer) error {
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
	// deserialize lowLevel
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.lowLevel)
		if err != nil {
			return err
		}
	}
	// deserialize highLevel
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.highLevel)
		if err != nil {
			return err
		}
	}
	// deserialize lowRegTime
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.lowRegTime)
		if err != nil {
			return err
		}
	}
	// deserialize highRegTime
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.highRegTime)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_GMT2World_gmtNotice) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.bType != 0)
	mask.WriteBit(len(this.note) != 0)
	mask.WriteBit(this.thetime != 0)
	mask.WriteBit(this.itvtime != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize bType
	{
		if this.bType != 0 {
			err := prpc.Write(buffer, this.bType)
			if err != nil {
				return err
			}
		}
	}
	// serialize note
	if len(this.note) != 0 {
		err := prpc.Write(buffer, this.note)
		if err != nil {
			return err
		}
	}
	// serialize thetime
	{
		if this.thetime != 0 {
			err := prpc.Write(buffer, this.thetime)
			if err != nil {
				return err
			}
		}
	}
	// serialize itvtime
	{
		if this.itvtime != 0 {
			err := prpc.Write(buffer, this.itvtime)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *SGE_GMT2World_gmtNotice) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize bType
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.bType)
		if err != nil {
			return err
		}
	}
	// deserialize note
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.note)
		if err != nil {
			return err
		}
	}
	// deserialize thetime
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.thetime)
		if err != nil {
			return err
		}
	}
	// deserialize itvtime
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.itvtime)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_GMT2World_setChargeTotal) Serialize(buffer *bytes.Buffer) error {
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
func (this *SGE_GMT2World_setChargeTotal) Deserialize(buffer *bytes.Buffer) error {
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
func (this *SGE_GMT2World_setChargeEvery) Serialize(buffer *bytes.Buffer) error {
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
func (this *SGE_GMT2World_setChargeEvery) Deserialize(buffer *bytes.Buffer) error {
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
func (this *SGE_GMT2World_setDiscountStore) Serialize(buffer *bytes.Buffer) error {
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
func (this *SGE_GMT2World_setDiscountStore) Deserialize(buffer *bytes.Buffer) error {
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
func (this *SGE_GMT2World_setLoginTotal) Serialize(buffer *bytes.Buffer) error {
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
func (this *SGE_GMT2World_setLoginTotal) Deserialize(buffer *bytes.Buffer) error {
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
func (this *SGE_GMT2World_setHotRole) Serialize(buffer *bytes.Buffer) error {
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
func (this *SGE_GMT2World_setHotRole) Deserialize(buffer *bytes.Buffer) error {
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
func (this *SGE_GMT2World_setEmployeeActivity) Serialize(buffer *bytes.Buffer) error {
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
func (this *SGE_GMT2World_setEmployeeActivity) Deserialize(buffer *bytes.Buffer) error {
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
func (this *SGE_GMT2World_setMinGiftBagActivity) Serialize(buffer *bytes.Buffer) error {
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
func (this *SGE_GMT2World_setMinGiftBagActivity) Deserialize(buffer *bytes.Buffer) error {
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
func (this *SGE_GMT2World_setZhuanpanActivity) Serialize(buffer *bytes.Buffer) error {
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
func (this *SGE_GMT2World_setZhuanpanActivity) Deserialize(buffer *bytes.Buffer) error {
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
func (this *SGE_GMT2World_setIntegralshop) Serialize(buffer *bytes.Buffer) error {
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
func (this *SGE_GMT2World_setIntegralshop) Deserialize(buffer *bytes.Buffer) error {
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
func (this *SGE_GMT2World_makeOrder) Serialize(buffer *bytes.Buffer) error {
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
func (this *SGE_GMT2World_makeOrder) Deserialize(buffer *bytes.Buffer) error {
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
func (this *SGE_GMT2World_doScript) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.script) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize script
	if len(this.script) != 0 {
		err := prpc.Write(buffer, this.script)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_GMT2World_doScript) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize script
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.script)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_GMT2World_playerDoScript) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.playerId != 0)
	mask.WriteBit(len(this.script) != 0)
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
	// serialize script
	if len(this.script) != 0 {
		err := prpc.Write(buffer, this.script)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_GMT2World_playerDoScript) Deserialize(buffer *bytes.Buffer) error {
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
	// deserialize script
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.script)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *SGE_GMT2WorldStub) addExp(playerId uint32, exp uint32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(0))
	if err != nil {
		return err
	}
	_0 := SGE_GMT2World_addExp{}
	_0.playerId = playerId
	_0.exp = exp
	err = _0.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_GMT2WorldStub) addMoney(playerId uint32, money int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(1))
	if err != nil {
		return err
	}
	_1 := SGE_GMT2World_addMoney{}
	_1.playerId = playerId
	_1.money = money
	err = _1.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_GMT2WorldStub) addDiamond(playerId uint32, diamond int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(2))
	if err != nil {
		return err
	}
	_2 := SGE_GMT2World_addDiamond{}
	_2.playerId = playerId
	_2.diamond = diamond
	err = _2.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_GMT2WorldStub) openGM(playerId uint32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(3))
	if err != nil {
		return err
	}
	_3 := SGE_GMT2World_openGM{}
	_3.playerId = playerId
	err = _3.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_GMT2WorldStub) closeGM(playerId uint32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(4))
	if err != nil {
		return err
	}
	_4 := SGE_GMT2World_closeGM{}
	_4.playerId = playerId
	err = _4.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_GMT2WorldStub) noTalkPlayer(playerId uint32, time uint32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(5))
	if err != nil {
		return err
	}
	_5 := SGE_GMT2World_noTalkPlayer{}
	_5.playerId = playerId
	_5.time = time
	err = _5.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_GMT2WorldStub) sealPlayer(playerId uint32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(6))
	if err != nil {
		return err
	}
	_6 := SGE_GMT2World_sealPlayer{}
	_6.playerId = playerId
	err = _6.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_GMT2WorldStub) kickPlayer(playerId uint32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(7))
	if err != nil {
		return err
	}
	_7 := SGE_GMT2World_kickPlayer{}
	_7.playerId = playerId
	err = _7.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_GMT2WorldStub) openTalkPlayer(playerId uint32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(8))
	if err != nil {
		return err
	}
	_8 := SGE_GMT2World_openTalkPlayer{}
	_8.playerId = playerId
	err = _8.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_GMT2WorldStub) unsealPlayer(playerId uint32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(9))
	if err != nil {
		return err
	}
	_9 := SGE_GMT2World_unsealPlayer{}
	_9.playerId = playerId
	err = _9.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_GMT2WorldStub) sendMailAllOnline(mail COM_Mail, lowLevel int32, highLevel int32, lowRegTime int64, highRegTime int64) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(10))
	if err != nil {
		return err
	}
	_10 := SGE_GMT2World_sendMailAllOnline{}
	_10.mail = mail
	_10.lowLevel = lowLevel
	_10.highLevel = highLevel
	_10.lowRegTime = lowRegTime
	_10.highRegTime = highRegTime
	err = _10.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_GMT2WorldStub) gmtNotice(bType int, note string, thetime uint64, itvtime int64) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(11))
	if err != nil {
		return err
	}
	_11 := SGE_GMT2World_gmtNotice{}
	_11.bType = bType
	_11.note = note
	_11.thetime = thetime
	_11.itvtime = itvtime
	err = _11.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_GMT2WorldStub) setChargeTotal(data COM_ADChargeTotal) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(12))
	if err != nil {
		return err
	}
	_12 := SGE_GMT2World_setChargeTotal{}
	_12.data = data
	err = _12.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_GMT2WorldStub) setChargeEvery(data COM_ADChargeEvery) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(13))
	if err != nil {
		return err
	}
	_13 := SGE_GMT2World_setChargeEvery{}
	_13.data = data
	err = _13.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_GMT2WorldStub) setDiscountStore(data COM_ADDiscountStore) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(14))
	if err != nil {
		return err
	}
	_14 := SGE_GMT2World_setDiscountStore{}
	_14.data = data
	err = _14.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_GMT2WorldStub) setLoginTotal(data COM_ADLoginTotal) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(15))
	if err != nil {
		return err
	}
	_15 := SGE_GMT2World_setLoginTotal{}
	_15.data = data
	err = _15.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_GMT2WorldStub) setHotRole(data COM_ADHotRole) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(16))
	if err != nil {
		return err
	}
	_16 := SGE_GMT2World_setHotRole{}
	_16.data = data
	err = _16.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_GMT2WorldStub) setEmployeeActivity(data COM_ADEmployeeTotal) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(17))
	if err != nil {
		return err
	}
	_17 := SGE_GMT2World_setEmployeeActivity{}
	_17.data = data
	err = _17.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_GMT2WorldStub) setMinGiftBagActivity(data COM_ADGiftBag) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(18))
	if err != nil {
		return err
	}
	_18 := SGE_GMT2World_setMinGiftBagActivity{}
	_18.data = data
	err = _18.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_GMT2WorldStub) setZhuanpanActivity(data COM_ZhuanpanData) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(19))
	if err != nil {
		return err
	}
	_19 := SGE_GMT2World_setZhuanpanActivity{}
	_19.data = data
	err = _19.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_GMT2WorldStub) setIntegralshop(data COM_IntegralData) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(20))
	if err != nil {
		return err
	}
	_20 := SGE_GMT2World_setIntegralshop{}
	_20.data = data
	err = _20.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_GMT2WorldStub) makeOrder(playerId uint32, order SGE_GmtOrder) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(21))
	if err != nil {
		return err
	}
	_21 := SGE_GMT2World_makeOrder{}
	_21.playerId = playerId
	_21.order = order
	err = _21.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_GMT2WorldStub) doScript(script string) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(22))
	if err != nil {
		return err
	}
	_22 := SGE_GMT2World_doScript{}
	_22.script = script
	err = _22.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *SGE_GMT2WorldStub) playerDoScript(playerId uint32, script string) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(23))
	if err != nil {
		return err
	}
	_23 := SGE_GMT2World_playerDoScript{}
	_23.playerId = playerId
	_23.script = script
	err = _23.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func Bridging_SGE_GMT2World_addExp(buffer *bytes.Buffer, p SGE_GMT2WorldProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_0 := SGE_GMT2World_addExp{}
	err := _0.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.addExp(_0.playerId, _0.exp)
}
func Bridging_SGE_GMT2World_addMoney(buffer *bytes.Buffer, p SGE_GMT2WorldProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_1 := SGE_GMT2World_addMoney{}
	err := _1.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.addMoney(_1.playerId, _1.money)
}
func Bridging_SGE_GMT2World_addDiamond(buffer *bytes.Buffer, p SGE_GMT2WorldProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_2 := SGE_GMT2World_addDiamond{}
	err := _2.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.addDiamond(_2.playerId, _2.diamond)
}
func Bridging_SGE_GMT2World_openGM(buffer *bytes.Buffer, p SGE_GMT2WorldProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_3 := SGE_GMT2World_openGM{}
	err := _3.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.openGM(_3.playerId)
}
func Bridging_SGE_GMT2World_closeGM(buffer *bytes.Buffer, p SGE_GMT2WorldProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_4 := SGE_GMT2World_closeGM{}
	err := _4.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.closeGM(_4.playerId)
}
func Bridging_SGE_GMT2World_noTalkPlayer(buffer *bytes.Buffer, p SGE_GMT2WorldProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_5 := SGE_GMT2World_noTalkPlayer{}
	err := _5.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.noTalkPlayer(_5.playerId, _5.time)
}
func Bridging_SGE_GMT2World_sealPlayer(buffer *bytes.Buffer, p SGE_GMT2WorldProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_6 := SGE_GMT2World_sealPlayer{}
	err := _6.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.sealPlayer(_6.playerId)
}
func Bridging_SGE_GMT2World_kickPlayer(buffer *bytes.Buffer, p SGE_GMT2WorldProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_7 := SGE_GMT2World_kickPlayer{}
	err := _7.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.kickPlayer(_7.playerId)
}
func Bridging_SGE_GMT2World_openTalkPlayer(buffer *bytes.Buffer, p SGE_GMT2WorldProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_8 := SGE_GMT2World_openTalkPlayer{}
	err := _8.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.openTalkPlayer(_8.playerId)
}
func Bridging_SGE_GMT2World_unsealPlayer(buffer *bytes.Buffer, p SGE_GMT2WorldProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_9 := SGE_GMT2World_unsealPlayer{}
	err := _9.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.unsealPlayer(_9.playerId)
}
func Bridging_SGE_GMT2World_sendMailAllOnline(buffer *bytes.Buffer, p SGE_GMT2WorldProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_10 := SGE_GMT2World_sendMailAllOnline{}
	err := _10.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.sendMailAllOnline(_10.mail, _10.lowLevel, _10.highLevel, _10.lowRegTime, _10.highRegTime)
}
func Bridging_SGE_GMT2World_gmtNotice(buffer *bytes.Buffer, p SGE_GMT2WorldProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_11 := SGE_GMT2World_gmtNotice{}
	err := _11.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.gmtNotice(_11.bType, _11.note, _11.thetime, _11.itvtime)
}
func Bridging_SGE_GMT2World_setChargeTotal(buffer *bytes.Buffer, p SGE_GMT2WorldProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_12 := SGE_GMT2World_setChargeTotal{}
	err := _12.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.setChargeTotal(_12.data)
}
func Bridging_SGE_GMT2World_setChargeEvery(buffer *bytes.Buffer, p SGE_GMT2WorldProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_13 := SGE_GMT2World_setChargeEvery{}
	err := _13.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.setChargeEvery(_13.data)
}
func Bridging_SGE_GMT2World_setDiscountStore(buffer *bytes.Buffer, p SGE_GMT2WorldProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_14 := SGE_GMT2World_setDiscountStore{}
	err := _14.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.setDiscountStore(_14.data)
}
func Bridging_SGE_GMT2World_setLoginTotal(buffer *bytes.Buffer, p SGE_GMT2WorldProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_15 := SGE_GMT2World_setLoginTotal{}
	err := _15.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.setLoginTotal(_15.data)
}
func Bridging_SGE_GMT2World_setHotRole(buffer *bytes.Buffer, p SGE_GMT2WorldProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_16 := SGE_GMT2World_setHotRole{}
	err := _16.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.setHotRole(_16.data)
}
func Bridging_SGE_GMT2World_setEmployeeActivity(buffer *bytes.Buffer, p SGE_GMT2WorldProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_17 := SGE_GMT2World_setEmployeeActivity{}
	err := _17.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.setEmployeeActivity(_17.data)
}
func Bridging_SGE_GMT2World_setMinGiftBagActivity(buffer *bytes.Buffer, p SGE_GMT2WorldProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_18 := SGE_GMT2World_setMinGiftBagActivity{}
	err := _18.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.setMinGiftBagActivity(_18.data)
}
func Bridging_SGE_GMT2World_setZhuanpanActivity(buffer *bytes.Buffer, p SGE_GMT2WorldProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_19 := SGE_GMT2World_setZhuanpanActivity{}
	err := _19.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.setZhuanpanActivity(_19.data)
}
func Bridging_SGE_GMT2World_setIntegralshop(buffer *bytes.Buffer, p SGE_GMT2WorldProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_20 := SGE_GMT2World_setIntegralshop{}
	err := _20.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.setIntegralshop(_20.data)
}
func Bridging_SGE_GMT2World_makeOrder(buffer *bytes.Buffer, p SGE_GMT2WorldProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_21 := SGE_GMT2World_makeOrder{}
	err := _21.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.makeOrder(_21.playerId, _21.order)
}
func Bridging_SGE_GMT2World_doScript(buffer *bytes.Buffer, p SGE_GMT2WorldProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_22 := SGE_GMT2World_doScript{}
	err := _22.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.doScript(_22.script)
}
func Bridging_SGE_GMT2World_playerDoScript(buffer *bytes.Buffer, p SGE_GMT2WorldProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_23 := SGE_GMT2World_playerDoScript{}
	err := _23.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.playerDoScript(_23.playerId, _23.script)
}
func SGE_GMT2WorldDispatch(buffer *bytes.Buffer, p SGE_GMT2WorldProxy) error {
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
		return Bridging_SGE_GMT2World_addExp(buffer, p)
	case 1:
		return Bridging_SGE_GMT2World_addMoney(buffer, p)
	case 2:
		return Bridging_SGE_GMT2World_addDiamond(buffer, p)
	case 3:
		return Bridging_SGE_GMT2World_openGM(buffer, p)
	case 4:
		return Bridging_SGE_GMT2World_closeGM(buffer, p)
	case 5:
		return Bridging_SGE_GMT2World_noTalkPlayer(buffer, p)
	case 6:
		return Bridging_SGE_GMT2World_sealPlayer(buffer, p)
	case 7:
		return Bridging_SGE_GMT2World_kickPlayer(buffer, p)
	case 8:
		return Bridging_SGE_GMT2World_openTalkPlayer(buffer, p)
	case 9:
		return Bridging_SGE_GMT2World_unsealPlayer(buffer, p)
	case 10:
		return Bridging_SGE_GMT2World_sendMailAllOnline(buffer, p)
	case 11:
		return Bridging_SGE_GMT2World_gmtNotice(buffer, p)
	case 12:
		return Bridging_SGE_GMT2World_setChargeTotal(buffer, p)
	case 13:
		return Bridging_SGE_GMT2World_setChargeEvery(buffer, p)
	case 14:
		return Bridging_SGE_GMT2World_setDiscountStore(buffer, p)
	case 15:
		return Bridging_SGE_GMT2World_setLoginTotal(buffer, p)
	case 16:
		return Bridging_SGE_GMT2World_setHotRole(buffer, p)
	case 17:
		return Bridging_SGE_GMT2World_setEmployeeActivity(buffer, p)
	case 18:
		return Bridging_SGE_GMT2World_setMinGiftBagActivity(buffer, p)
	case 19:
		return Bridging_SGE_GMT2World_setZhuanpanActivity(buffer, p)
	case 20:
		return Bridging_SGE_GMT2World_setIntegralshop(buffer, p)
	case 21:
		return Bridging_SGE_GMT2World_makeOrder(buffer, p)
	case 22:
		return Bridging_SGE_GMT2World_doScript(buffer, p)
	case 23:
		return Bridging_SGE_GMT2World_playerDoScript(buffer, p)
	default:
		return errors.New(prpc.NoneDispatchMatchError)
	}
}
