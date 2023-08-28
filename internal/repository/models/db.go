package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"tiktok/conf"
)

var DB *gorm.DB

func InitDB() {
	var err error
	config := conf.NewConfig()
	DB, err = gorm.Open(mysql.Open(conf.DBConnectString(&config.DB)), &gorm.Config{
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
		//Logger:                 logger.Default.LogMode(logger.Global),
	})
	if err != nil {
		panic(err)
	}
	err = DB.AutoMigrate(&UserInfo{}, &Comment{}, &UserLogin{}, &Video{})
	if err != nil {
		panic(err)
	}
}
