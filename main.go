package main

import (
	"flag"
	"fmt"
)

// Gitリポジトリの最新バージョン
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

	// エントリーポイントの設定・サーバー起動
	err := StartMainServer()
	if err != nil {
		panic(err)
	}
}
