package main

import (
	"access-log-app/pkg/global"
	"access-log-app/pkg/initialize"
	log "github.com/xgtcode/log-demo"
)

//15011420771
func main() {
	var err error
	// 初始化数据库
	//global.GlobalDb, err = initialize.InitDB()
	global.GlobalDb, err = initialize.InitSqliteDB()
	if err != nil{
		log.Fatalf("initDb err: %+v", err)
	}
	log.Info("initialize db  success")
	engine := initialize.InitRouter()
	if err := engine.Run(":8080"); err != nil {
		log.Error("start server err ", err)
	}
}