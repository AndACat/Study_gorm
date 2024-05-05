package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

// 定义mysql层面全局DB
var (
	DB *gorm.DB
)

func InitMySql() (*gorm.DB, error) {
	// 数据库配置等信息
	// golang:4_wvcLDCfhw.U4P12345@tcp(sh-cynosdbmysql-grp-o11thlf4.sql.tencentcdb.com:22547)/golang?charset=utf8mb4&parseTime=True&loc=Local
	username := "golang"
	password := "4_wvcLDCfhw.U4P12345@tcp"
	host := "sh-cynosdbmysql-grp-o11thlf4.sql.tencentcdb.com"
	port := "22547"
	dbName := "golang"
	url := fmt.Sprintf("%s:%s(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, dbName)
	log.Println("建立数据库连接url: " + url)
	// 这里无法直接赋值给DB
	Db, err := gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		log.Println("open mysql connect error:", err)
		panic("open mysql connect error:" + err.Error())
	}

	// 打开DB日志
	Db.Debug()
	sqlDb, _ := Db.DB()
	// 设置连接池的最大闲置数量
	sqlDb.SetConnMaxLifetime(10)
	// 设置连接池的最大连接数量
	sqlDb.SetMaxOpenConns(100)
	// 设置连接的最大复用时间 100秒
	sqlDb.SetConnMaxLifetime(time.Second * 100)
	// 赋值  DB: 全局db   Db:建立连接返回的db
	DB = Db
	return DB, nil
}
