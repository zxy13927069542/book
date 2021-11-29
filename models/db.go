package models

import (
	"book/config"
	"book/models/system"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

/**
进行数据库初始化和获取数据库链接
*/

var db *gorm.DB

func SetupDataBase() {
	//获取配置文件
	config := config.GetConfig()

	//从配置文件读取数据库参数
	dbMysql, err := gorm.Open(mysql.Open(config.Mysql.DataSource), &gorm.Config{})
	if err != nil {
		fmt.Println("连接数据库失败", err)
		os.Exit(0)
	}
	db = dbMysql

	//数据库迁移
	//并不会导致数据库数据消失
	err = db.AutoMigrate(&system.Book{})
	if err != nil {
		fmt.Println("数据库迁移失败", err)
		os.Exit(0)
	}

}

func GetDB() *gorm.DB {
	return db
}
