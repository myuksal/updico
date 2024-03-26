package database

import (
	"time"

	"github.com/myuksal/updico/entities"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func CreateConnection() *gorm.DB {

	dsn := "root:root@tcp(127.0.0.1:3306)/updico?charset=utf8mb4"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Db 연결에 실패하였습니다.")
	}
	db.AutoMigrate(
		// &entities.BlockEntity{},
		&entities.TransactionEntity{},
	)

	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(2)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(10 * time.Minute)

	return db
}
