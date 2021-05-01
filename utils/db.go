package utils

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func GormConnect() *gorm.DB {
	DBMS := "mysql"
	USER := "root"
	PASS := "root"
	dbIp := os.Getenv("DB_IP")
	var PROTOCOL string
	if dbIp != "" {
		PROTOCOL = "tcp(" + dbIp + ":3306)"
	} else {
		PROTOCOL = "tcp(127.0.0.1:3306)"
	}
	DBNAME := "gomix_db"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true" //parseTimeで時間のScanが可能になる
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		log.Println(err)
	}

	return db
}
