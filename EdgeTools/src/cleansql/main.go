package main

import (
	"bytes"
	"cleansql/prpc"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"time"
)

var debugLog *os.File

func OpenDB(dsn string) *sql.DB {
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	return db
}

func cleanGO(from string) {
	db := OpenDB(from)
	if db == nil {
		return
	}
	defer db.Close()

	records, err := db.Query("SELECT `BinData` FROM `Player`")

	if err != nil {
		fmt.Println(err)
		return
	}

	debugLog, err = os.Create( fmt.Sprintf("%d.log",int(time.Now().Unix())))
	if err != nil {
		fmt.Println(err)
	}


	var instIds []int
	var instNames []string
	//查询符合条件的角色
	for records.Next() {
		var bs []byte
		err := records.Scan(&bs)
		if err != nil {
			fmt.Println(err)
			continue
		}
		b := bytes.NewBuffer(bs)
		inst := prpc.SGE_DBPlayerData{}
		err = inst.Deserialize(b)
		if err != nil {
			fmt.Println(err)
			continue
		}

		playerlevel := inst.GetPlayerLevel()
		tm := time.Unix(int64(inst.GetLogoutTime()), 0)
		cd := time.Now().Sub(tm)
		days := cd.Hours() / 24

		if playerlevel <= 5 && days > 14 {
			playerId := inst.GetPlayerInstId()
			playerName := inst.GetPlayerName()
			magic := inst.GetM()
			instIds = append(instIds, playerId)
			instNames = append(instNames, playerName)
			fmt.Printf("SELECT PLAYER TABLE %d   %s \n", playerId, playerName)

			debugLog.WriteString(fmt.Sprintf("DELETE PLAYER NAME[%s]  LogoutTime[%d] Level[%d] Magic[%d]\n", playerName, int(days), playerlevel, magic))

		}
	}
	debugLog.Sync()
	debugLog.Close()
	//干掉相关表中符合条件的角色
	for i := 0; i < len(instNames); i++ {
		db.Exec("DELETE FROM `Player` WHERE `PlayerName`= ? ", instNames[i])
		db.Exec("DELETE FROM `Baby` WHERE `OwnerName`= ? ", instNames[i])
		db.Exec("DELETE FROM `Employee` WHERE `OwnerName`= ? ", instNames[i])
		db.Exec("DELETE FROM `Mail` WHERE `RecvName`= ? ", instNames[i])
		db.Exec("DELETE FROM `EndlessStair` WHERE `PlayerName`= ? ", instNames[i])
		db.Exec("DELETE FROM `GuildMember` WHERE `RoleName`= ? ", instNames[i])

		fmt.Printf("DELETE TABLE PLAYER BABY EMPLOYEE MAIL ENDLESSSTAIR GUILDMEMBER BY %s ===NUM[%d] ===INDEX[%d]\n", instNames[i], len(instNames), i)

	}

	for i := 0; i < len(instIds); i++ {
		db.Exec("DELETE FROM `EmployeeQuestTable` WHERE `PlayerId`= ? ", instIds[i])
		db.Exec("DELETE FROM `MallTable` WHERE `PlayerId`= ? ", instIds[i])
		db.Exec("DELETE FROM `MallSelledTable` WHERE `PlayerId`= ? ", instIds[i])

		fmt.Printf("DELETE TABLE EmployeeQuestTable MallTable MallSelledTable BY %d ===NUM[%d] ===INDEX[%d]\n", instIds[i], len(instIds), i)
	}

	fmt.Printf("CLEAN END !!!! \n")
}

func main() {

	conf, _ := NewConfigerFile("clean.conf")

	froms := conf.GetStrings("froms")

	for _, from := range froms {
		fmt.Println("Begin clean ", from)
		cleanGO(from)
	}



}
