//无尽之阶
#ifndef __ENDLESSSTAIR_H__
#define __ENDLESSSTAIR_H__

#include "config.h"
#include "Robot.h"

class Player;
class EndlessStair
{
public:
	EndlessStair();
public:
	static EndlessStair* instance() {static EndlessStair serv; return &serv;}

	static void init();
	static void fini();
	static void	robotInst();
	static int	findRankByName(std::string name);
public:
	
	void	getEndlessStairDate(std::vector<std::string> name);
	void	creatArena(Player* player,std::string name);
	Robot*	findRobotByName(std::string name);
	bool	findRivalName(std::string name);
	void	calcRival(int rank);
	void	checkRival(Player* pPlayer);
	void	getPlayerDBData(SGE_DBPlayerData &data);
	void	requestJJCData(Player* pPlayer);
	void	requestRank(Player* pPlayer);
	void	calcRank(Entity* pStart,std::string passName, bool isWin);
	void	checkMsg(Player* pPlayer, std::string name);
	void	rankrewardbytimes(std::string& sendName,std::string& title, std::string& content,U32 rank);
	void	rankrewardbyday(std::string& sendName,std::string& title, std::string& content);
	void	rankrewardbysenson(std::string& sendName,std::string& title, std::string& content);
	void	deleteplayerReCalcRank(std::string playname);
	void	firstWinreward(std::string& sendName,std::string& title, std::string& content,U32 rank,U32 dropid);
	bool	isFirstWin(U32 rank);
	static std::map<U32,Robot*>			robots_;
	static std::vector<std::string>		nameRate_;	//JJC名次序列
	std::vector<std::string>			rivalName_;
	std::vector<std::string>			battleName_;
	std::vector<std::string>			firstwin_;		//已发首胜奖励名单

	bool		isCreatBattle_;			//是否开启战斗

	U32		rivalPlayer_;			//当前挑战者
	U32		checkPlayer_;			//查看信息请求者
};

#endif