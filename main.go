package main

import (
	"gomix/cli"
	"gomix/cli/c_memo"
	"gomix/config"
	"gomix/pkg/memo"

	"os"
)

func main() {

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
