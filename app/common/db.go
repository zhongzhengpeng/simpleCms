package common

import (
	// mysql driver 驱动包
	"log"

	"gorm.io/driver/mysql"

	// gorm 包
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	dsn := "root:secret@tcp(127.0.0.1:3316)/simpleCms?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	DB = db
}
