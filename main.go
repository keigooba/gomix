package main

import (
	"flag"
	"fmt"
	"gomix/config"
	"gomix/pkg/memo"
)

// Gitリポジトリの最新バージョン start.shで更新される
var version = "1.0.0"

func main() {

	var showVersion bool
	// -v -versionが指定された場合にshowVerionが真になるよう定義
	flag.BoolVar(&showVersion, "v", false, "show version")
	flag.BoolVar(&showVersion, "version", false, "show version")
	flag.Parse() //引数からオプションをパースする
	if showVersion {
		// バージョン番号を表示する
		fmt.Println("version", version)
	}

	// マイグレーション
	config.Db.AutoMigrate(&memo.Memo{})

	// サーバー停止の通知
	go signalCall()

	// エントリーポイントの設定・サーバー起動
	err := StartMainServer()
	if err != nil {
		panic(err)
	}

}
