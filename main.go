package main

import (
	"gomix/cli/c_memo"
	"gomix/config"
	"gomix/pkg/memo"
)

func main() {

	// マイグレーション
	config.Db.AutoMigrate(&memo.Memo{})

	// メモのcliコマンドの設定
	c_memo.CliMemo()

	// サーバー停止の通知設定
	// go signalCall()

	// オプションコマンドの設定
	cmdFlag()

	// エントリーポイントの設定・サーバー起動
	err := StartMainServer()
	if err != nil {
		panic(err)
	}

}
