#include "config.h"
#include "tinyxml/tinyxml.h"
#include "CSVParser.h"
#include "TableSystem.h"
#include "tmptable.h"
#include "skilltable.h"
#include "sttable.h"
#include "TokenParser.h"
#include "itemtable.h"
#include "DropTable.h"
#include "npctable.h"
#include "monstertable.h"
#include "scenetable.h"
#include "employeeTable.h"
#include "Quest.h"
#include "LiveSkilltable.h"
#include "LotteryTable.h"
#include "randName.h"
#include "robotTable.h"
#include "profession.h"
#include "BattleData.h"
#include "Shop.h"
#include "GameEvent.h"
#include "titleTable.h"
#include "achievementTable.h"
#include "challengeTable.h"
#include "DailyReward.h"
#include "PVPrunkTable.h"
#include "EmployeeConfig.h"
#include "DebrisTable.h"
#include "ArtifactLevelTable.h"
#include "ArtifactConfigTable.h"
#include "ArtifactChangeData.h"
#include "FilterWord.h"
#include "Activities.h"
#include "GuildData.h"
#include "copyscene.h"
#include "lastOrder.h"
#include "EmployeeQuestTable.h"
bool TableSystem::Load(){
	///load
	FUNCTION_PROBE;
	if(false == PlayerTmpTable::load(DESIGN_TABLE_ROOT"PlayerData.csv"))
		return false;
	if(false == SkillTable::load(DESIGN_TABLE_ROOT"SkillData.csv"))
		return false;
	if(false == SkillTable::loadBabyEquipSkillGroup(DESIGN_TABLE_ROOT"BabySuitSkill.csv"))
		return false;
	if(false == StateTable::load(DESIGN_TABLE_ROOT"State.csv"))
		return false;
	if(false == ItemTable::load(DESIGN_TABLE_ROOT"ItemData.csv"))
		return false;
	if(false == DropTable::load(DESIGN_TABLE_ROOT"Drop.csv"))
		return false;
	if(false == NpcTable::load(DESIGN_TABLE_ROOT"Npc.csv"))
		return false;
	if(false == MonsterTable::load(DESIGN_TABLE_ROOT"Monster.csv"))
		return false;
	if(false == MonsterClass::load(DESIGN_TABLE_ROOT"AIclass.csv"))
		return false;
	if(false == Monster2Table::load(DESIGN_TABLE_ROOT"Monster2.csv"))
		return false;
	if(false == SceneTable::load(DESIGN_TABLE_ROOT"Scene.csv"))
		return false;
	if(false == EmployeeTable::load(DESIGN_TABLE_ROOT"EmployeeData.csv"))
		return false;
	if(false == ExpTable::load(DESIGN_TABLE_ROOT"Exp.csv"))
		return false;
	if(false == Quest::load(DESIGN_TABLE_ROOT"Quest.csv"))
		return false;
	if(false == GatherTable::load(DESIGN_TABLE_ROOT"Gather.csv"))
		return false;
	if(false == MakeTable::load(DESIGN_TABLE_ROOT"Make.csv"))
		return false;
	if(false == LotteryTable::load(DESIGN_TABLE_ROOT"Lottery.csv"))
		return false;
	if(false == RandNameTable::load(DESIGN_TABLE_ROOT"Randname.csv"))
		return false;
	if(false == RobotTab::load(DESIGN_TABLE_ROOT"RobotData.csv"))
		return false;
	if(false == PlayerAI::load(DESIGN_TABLE_ROOT"ArenaAiData.csv"))
		return false;
	if(false == Profession::load(DESIGN_TABLE_ROOT"ProfessionData.csv"))
		return false;
	if(false == BattleData::load(DESIGN_TABLE_ROOT"BattleData.csv"))
		return false;
	if(false == Shop::load(DESIGN_TABLE_ROOT"ShopData.csv"))
		return false;
	if(false == GameEvent::load(DESIGN_TABLE_ROOT"GameEvent.xml"))
		return false;
	if(false == TitleTable::load(DESIGN_TABLE_ROOT"Title.csv"))
		return false;
	if(false == AchievementTable::load(DESIGN_TABLE_ROOT"AchieveData.csv"))
		return false;
	if(false == ChallengeTable::load(DESIGN_TABLE_ROOT"ChallengeData.csv"))
		return false;
	if(false == DailyReward::load(DESIGN_TABLE_ROOT"DailyReward.csv"))
		return false;
	if(false == PvpRunkTable::load(DESIGN_TABLE_ROOT"PVPrunk.csv"))
		return false;
	if(false == PvRrewardTable::load(DESIGN_TABLE_ROOT"PVRreward.csv"))
		return false;
	if(false == EmployeeConfigTable::load(DESIGN_TABLE_ROOT"EmployeeConfig.csv"))
		return false;
	if(false == DiamondsConfig::load(DESIGN_TABLE_ROOT"Diamondsconfig.csv"))
		return false;
	if(false == DebrisTable::load(DESIGN_TABLE_ROOT"DebrisConfig.csv"))
		return false;
	if(false == ArtifactLevelTable::load(DESIGN_TABLE_ROOT"ArtifactLevel.csv"))
		return false;
	if(false == ArtifactConfigTable::load(DESIGN_TABLE_ROOT"ArtifactConfig.csv"))
		return false;
	if(false == FilterWord::load(DESIGN_TABLE_ROOT"filterword.csv"))
		return false;
	if(false == ArtifactChangeTable::load(DESIGN_TABLE_ROOT"ArtifactChange.csv"))
		return false;
	if(false == PetActivity::load(DESIGN_TABLE_ROOT"PetActivityData.csv"))
		return false;
	if(false == GuildShopItemData::load(DESIGN_TABLE_ROOT"HomeShopData.csv"))
		return false;
	if(false == GuildBuildingData::load(DESIGN_TABLE_ROOT"family.csv"))
		return false;
	if(false == GuildBlessingData::load(DESIGN_TABLE_ROOT"Blessing.csv"))
		return false;
	if(false == PetIntensive::load(DESIGN_TABLE_ROOT"PetIntensive.csv"))
		return false;
	if(false == CopyScene::load(DESIGN_TABLE_ROOT"Copy.csv"))
		return false;
	if(false == OnlineTimeClass::load(DESIGN_TABLE_ROOT"timereward.csv"))
		return false;
	if(false == GrowFundTable::load(DESIGN_TABLE_ROOT"foundation.csv"))
		return false;
	if(false == RuneTable::load(DESIGN_TABLE_ROOT"Runes.csv"))
		return false;
	if(false == LastOrderTable::load(DESIGN_TABLE_ROOT"lastOrder.csv"))
		return false;
	if(false == RobotActionTable::load(DESIGN_TABLE_ROOT"AiNt.csv"))
		return false;
	if(false == EmployeeQuest::load(DESIGN_TABLE_ROOT"EmployeeQuest.csv"))
		return false;
	if(false == EmployeeSkill::load(DESIGN_TABLE_ROOT"EmployeeSkill.csv"))
		return false;
	if(false == EmployeeMonster::load(DESIGN_TABLE_ROOT"EmployeeMonster.csv"))
		return false;
	if(false == CrystalTable::load(DESIGN_TABLE_ROOT"Crystal.csv"))
		return false;
	if(false == CrystalUpTable::load(DESIGN_TABLE_ROOT"CrystalUp.csv"))
		return false;
	return true;
}
bool TableSystem::Check(){
	///check
	if(false == PlayerTmpTable::check())
		return false;
	if(false == SkillTable::check())
		return false;
	if(false == StateTable::check())
		return false;
	if(false == ItemTable::check())
		return false;
	if(false == DropTable::check())
		return false;
	if(false == NpcTable::check())
		return false;
	if(false == MonsterTable::check())
		return false;
	if(false == Monster2Table::check())
		return false;
	if(false == MonsterClass::check())
		return false;
	if(false == SceneTable::check())
		return false;
	if(false == EmployeeTable::check())
		return false;
	if(false == ExpTable::check())
		return false;
	if(false == GatherTable::check())
		return false;
	if(false == MakeTable::check())
		return false;
	if(false == LotteryTable::check())
		return false;
	if(false == RandNameTable::check())
		return false;
	if(false == RobotTab::check())
		return false;
	if(false == PlayerAI::check())
		return false;
	if(false == Profession::check())
		return false;
	if(false == BattleData::check())
		return false;
	if(false == Shop::check())
		return false;
	if(false == GameEvent::check())
		return false;
	if(false == TitleTable::check())
		return false;
	if(false == AchievementTable::check())
		return false;
	if(false == ChallengeTable::check())
		return false;
	if(false == DailyReward::check())
		return false;
	if(false == PvpRunkTable::check())
		return false;
	if(false == EmployeeConfigTable::check())
		return false;
	if(false == DebrisTable::check())
		return false;
	if(false == PetActivity::check())
		return false;
	if(false == CopyScene::check())
		return false;
	if(false == ArtifactLevelTable::check())
		return false;
	if(false == ArtifactConfigTable::check())
		return false;
	if(false == GuildBlessingData::check())
		return false;
	if(false == OnlineTimeClass::check())
		return false;
	if(false == GrowFundTable::check())
		return false;
	if(false == RuneTable::check())
		return false;
	if(false == Quest::check())
		return false;
	if(false == RobotActionTable::check())
		return false;
	if(false == EmployeeQuest::check())
		return false;
	if(false == CrystalTable::check())
		return false;
	if(false == CrystalUpTable::check())
		return false;
	return true;
}

