package core

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"os"
	"sync"
	"time"
	"web/config"
)

var db *xorm.EngineGroup
var dbOnce sync.Once
var oldFile *os.File

func init() {
	initDB()
}

func DB() *xorm.EngineGroup {
	if db == nil {
		initDB()
	}
	return db
}

func initDB() {
	dbOnce.Do(func() {
		if db == nil {
			cfg := config.GetConfig()
			dbCfg := cfg.Db
			dbType := dbCfg.DBType
			database := dbCfg.DataBase
			user := dbCfg.User
			password := dbCfg.Password
			charset := dbCfg.Charset
			host := dbCfg.Host
			port := dbCfg.Port
			maxIdleConns := dbCfg.MaxIdleConns
			maxOpenConns := dbCfg.MaxOpenConns
			idleTime := dbCfg.IdleTime
			slaves := dbCfg.Slaves
			//初始化主库
			dataSourceName := user + ":" + password + "@" + "tcp(" + host + ":" + port + ")" + "/" + database + "?charset" + charset
			masterDb, err := xorm.NewEngine(dbType, dataSourceName)
			if err != nil {
				panic("MasterDB init Fail. " + string(err.Error()))
			}
			masterDb.SetMaxIdleConns(maxIdleConns)
			masterDb.SetMaxOpenConns(maxOpenConns)
			masterDb.SetConnMaxLifetime(time.Duration(idleTime))
			err = masterDb.Ping()
			if err != nil {
				panic("masterDB ping err " + string(err.Error()))
			}
			//初始化从库
			slaveArray := make([]*xorm.Engine, 0)
			for _, slave := range slaves {
				slave := slave
				slaveSourceName := slave.User + ":" + slave.Password + "@" + "tcp(" + slave.Host + ":" + slave.Port + ")" + "/" + database + "?charset" + charset
				slaveDB, err := xorm.NewEngine(dbType, slaveSourceName)
				if err != nil {
					panic("one of SlaveDB init Fail. host:" + slave.Host + " port:" + slave.Port + " " + string(err.Error()))
				}
				slaveDB.SetMaxIdleConns(maxIdleConns)
				slaveDB.SetMaxOpenConns(maxOpenConns)
				slaveDB.SetConnMaxLifetime(time.Duration(idleTime))
				err = slaveDB.Ping()
				if err != nil {
					panic("slaveDB ping err. " + "host:" + slave.Host + " port:" + slave.Port + " " + string(err.Error()))
				}
				slaveArray = append(slaveArray, slaveDB)
			}
			db, err = xorm.NewEngineGroup(masterDb, slaveArray, xorm.LeastConnPolicy())
			if err != nil {
				panic("DB cluster init fail! please check " + string(err.Error()))
			}
			fmt.Println("DB Start Success!")
			//开发环境日志记录设置
			if cfg.Server.Environment == "dev" || cfg.Server.Environment == "" {
				db.ShowSQL(true)
				db.Logger().SetLevel(core.LOG_DEBUG)
				f, err := os.Create("./sql.log")
				if f != oldFile {
					oldFile.Close()
					oldFile = f
				}
				if err != nil {
					panic(string(err.Error()))
				}
				db.SetLogger(xorm.NewSimpleLogger(f))
			}
		}
	})
}
