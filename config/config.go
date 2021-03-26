package config

import (
	"gomix/utils"
	"log"
	"time"

	ini "gopkg.in/ini.v1"
)

type ConfigList struct {
	Port string
	// SQLDriver string
	// DbName    string
	LogFile string
	Static  string
	Url     string
}

var Config ConfigList

func init() {
	// Configの設定の読み込み
	LoadConfig()

	// 現在の日付
	now_date := time.Now().Format("20060102")
	utils.LoggingSettings(Config.LogFile + now_date + ".log")
}

func LoadConfig() {

	cfg, err := ini.Load("config/config.ini")
	if err != nil {
		log.Fatalln(err)
	}

	Config = ConfigList{
		Port: cfg.Section("web").Key("port").MustString("8080"),
		// 		SQLDriver: cfg.Section("db").Key("driver").String(),
		// DbName:    cfg.Section("db").Key("name").String(),
		LogFile: cfg.Section("web").Key("logfile").String(),
		Static:  cfg.Section("web").Key("static").String(),
		Url:     cfg.Section("web").Key("url").String(),
	}
}
