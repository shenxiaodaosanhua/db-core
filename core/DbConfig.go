package core

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var GormDB *gorm.DB

func InitDB() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second * 2, // 慢 SQL 阈值
			LogLevel:                  logger.Info,     // 日志级别
			IgnoreRecordNotFoundError: true,            // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  false,           // 禁用彩色打印
		},
	)
	db, err := gorm.Open(mysql.Open(SysConfig.DBConfig.DSN), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatal(err)
	}
	sqldb, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	sqldb.SetMaxOpenConns(SysConfig.DBConfig.MaxOpenConn)
	sqldb.SetMaxIdleConns(SysConfig.DBConfig.MaxIdleConn)
	sqldb.SetConnMaxLifetime(time.Duration(SysConfig.DBConfig.MaxLifeTime) * time.Second)
	GormDB = db
}
