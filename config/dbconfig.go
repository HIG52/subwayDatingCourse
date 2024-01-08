package config

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func DbConnection() *gorm.DB {

	//로그
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // 콘솔에 로그를 출력
		logger.Config{
			SlowThreshold: time.Second, // 느린 쿼리 임계값
			LogLevel:      logger.Info, // 로그 레벨 설정
			Colorful:      true,        // 색상 있는 로그 출력
		},
	)

	dsn := "datero:datero1234@tcp(localhost:3306)/datero?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
