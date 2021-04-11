// +build darwin,amd64 windows linux,!android
// +build go1.1

package config

import (
	"encoding/json"
	"fmt"
	"gomix/utils"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/markbates/pkger"
)

type ConfigList struct {
	Port    string `json:"port"`
	LogFile string `json:"log_file"`
	Static  string `json:"static"`
	URL     string `json:"url"` //本番時up_urlに変更
}

// Config Configの定義
var Config ConfigList

var Db *gorm.DB

func init() {
	// Configの設定の読み込み
	err := LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	// 現在の日付
	nowDate := time.Now().Format("200601")
	// ログファイルの設定
	err = utils.LoggingSettings(Config.LogFile + nowDate + ".log")
	if err != nil {
		log.Fatal(err)
	}
	// コマンドの実行
	err = utils.Command()
	if err != nil {
		log.Fatal(err)
	}

	// DB接続
	Db = utils.GormConnect()

}

// LoadConfig Configの設定
func LoadConfig() error {

	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	fname := filepath.Join(cwd, "config", "config.json")
	f, err := pkger.Open(fname)
	if err != nil {
		return err
	}

	defer func() {
		err := f.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	//Configにconfig.jsonを読み込む
	err = json.NewDecoder(f).Decode(&Config)
	if err != nil {
		return err
	}

	// 環境変数の値の判定
	format := "Port: %s\nLogFile: %s\nStatic: %s\nURL: %s\n"
	_, err = fmt.Printf(format, Config.Port, Config.LogFile, Config.Static, Config.URL)
	if err != nil {
		return err
	}
	return nil //自明であればnilにする
}
