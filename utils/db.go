package utils

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func GormConnect() *gorm.DB {
	DBMS := "mysql"
	USER := "root"
	PASS := "root"
	PROTOCOL := "tcp(35.232.147.156:3306)" //127.0.0.1
	DBNAME := "gomix_db"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME +"?parseTime=true" //parseTimeで時間のScanが可能になる
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		log.Println(err)
	}

	return db
}
