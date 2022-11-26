package initialize

import (
	"access-log-app/pkg/model"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error){
	dsn := "root:root@tcp(mysqlhost:3306)/access-log?charset=utf8mb4&parseTime=True&loc=Local"
	mysqlDb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = mysqlDb.AutoMigrate(&model.AccessLog{})
	if err != nil {
		return nil, err
	}
	return mysqlDb, nil
}

func InitSqliteDB() (*gorm.DB, error) {
	sqlitDb, err := gorm.Open(sqlite.Open("./gorm.db"), &gorm.Config{})
	//sqlitDb, err := gorm.Open(sqlite.Open("gorm.io/driver/sqlserver"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = sqlitDb.AutoMigrate(&model.AccessLog{})
	if err != nil {
		return nil, err
	}
	return sqlitDb, nil
}
