package main

import (
	"flag"
	"fmt"
	"gomix/cli"
	"gomix/cli/c_memo"
	"gomix/config"
	"gomix/pkg/memo"

	"os"
)

// Gitリポジトリのバージョン start.shで最新のバージョンに更新される
var version = "1.0.0"

func main() {

	// Gitリポジトリのバージョン確認
	var showVersion bool
	// -v -versionが指定された場合にshowVerionが真になるよう定義
	flag.BoolVar(&showVersion, "version", false, "バージョン確認")
	flag.BoolVar(&showVersion, "v", false, "バージョン確認(short)")
	flag.Parse() //引数からオプションをパースする
	if showVersion {
		// バージョン番号を表示する
		fmt.Println("version", version)
	}

	// マイグレーション
	config.Db.AutoMigrate(&memo.Memo{})

	// コマンド入力の有無
	if len(os.Args) > 1 {

		// サーバー停止の通知設定
		go signalCall()

		// オプションコマンドの設定
		cli.CmdFlag()

		// メモのcliコマンドの設定
		c_memo.CliMemo()

	}

	// エントリーポイントの設定・サーバー起動
	err := StartMainServer()
	if err != nil {
		panic(err)
	}

}
