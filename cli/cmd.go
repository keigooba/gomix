package cli

import (
	"flag"
	"fmt"
	"gomix/config"
)

var ManageMemo bool

func CmdFlag() {

	// サードパーティーkingpinでの実装の場合
	// port := kingpin.Flag("port", "ポート設定が可能").Default("8888").Short('p').Int()
	// kingpin.Parse()

	// -v -versionが指定された場合にshowVerionが真になるよう定義
	flag.BoolVar(&ManageMemo, "memo", false, "メモの管理")
	flag.BoolVar(&ManageMemo, "m", false, "メモの管理(short)")

	// ポート設定のオプション
	// envPort, _ := strconv.Atoi(os.Getenv("PORT")) //環境変数でも指定できる
	flag.IntVar(&config.FlagPort, "port", config.Config.Port, "ポート設定が可能")
	flag.IntVar(&config.FlagPort, "p", config.Config.Port, "ポート設定が可能(short)")

	if ManageMemo {
		fmt.Println("ManageMemo")
	}
	fmt.Println("port", config.FlagPort) //ポート確認

}
