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

	// オプションコマンドの設定
	cli.CmdFlag()

	// コマンド入力の有無
	if len(os.Args) > 1 {

		// サーバー停止の通知設定
		go signalCall()

		// -memo,-mが入力された時
		if cli.ManageMemo {
			// サブコマンドの設定
			c_memo.CliMemo()
		}

	}

	// エントリーポイントの設定・サーバー起動
	err := StartMainServer()
	if err != nil {
		panic(err)
	}

}
