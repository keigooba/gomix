package main

import (
	"flag"
	"fmt"
	"gomix/config"
)

// Gitリポジトリのバージョン start.shで最新のバージョンに更新される
var version = "1.0.0"

func cmdFlag() {

	// サードパーティーkingpinでの実装の場合
	// port := kingpin.Flag("port", "ポート設定が可能").Default("8888").Short('p').Int()
	// kingpin.Parse()

	// ポート設定のオプション
	// envPort, _ := strconv.Atoi(os.Getenv("PORT")) //環境変数でも指定できる
	flag.IntVar(&config.FlagPort, "p", config.Config.Port, "ポート設定が可能(short)")
	flag.IntVar(&config.FlagPort, "port", config.Config.Port, "ポート設定が可能")

	// Gitリポジトリのバージョン確認
	var showVersion bool
	// -v -versionが指定された場合にshowVerionが真になるよう定義
	flag.BoolVar(&showVersion, "v", false, "バージョン確認(short)")
	flag.BoolVar(&showVersion, "version", false, "バージョン確認")
	flag.Parse() //引数からオプションをパースする
	if showVersion {
		// バージョン番号を表示する
		fmt.Println("version", version)
	}
	fmt.Println("port", config.FlagPort) //ポート確認

}
