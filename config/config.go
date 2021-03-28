package config

import (
	"fmt"
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
	URL     string
}

// Config Configの定義
var Config ConfigList

func init() {
	// Configの設定の読み込み
	err := LoadConfig()
	if err != nil {
		log.Println(err)
	}

	// 現在の日付
	nowDate := time.Now().Format("200601")
	utils.LoggingSettings(Config.LogFile + nowDate + ".log")
}

// LoadConfig Configの設定
func LoadConfig() error {

	cfg, err := ini.Load("config/config.ini")
	if err != nil {
		return err
	}

	Config = ConfigList{
		Port:    cfg.Section("web").Key("port").String(),
		LogFile: cfg.Section("web").Key("logfile").String(),
		Static:  cfg.Section("web").Key("static").String(),
		URL:     cfg.Section("web").Key("url").String(),
	}

	// 環境変数の値の判定
	format := "Port: %s\nLogFile: %s\nStatic: %s\nURL: %s\n"
	_, err = fmt.Printf(format, Config.Port, Config.LogFile, Config.Static, Config.URL)
	if err != nil {
		return err
	}
	return nil //自明であればnilにする
}
