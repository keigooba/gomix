package c_memo

import (
	"fmt"
	"gomix/pkg/memo"
	"os"
	"time"

	"github.com/mitchellh/cli"
)

func CliMemo() {
	// CLI structを生成する
	// 以下ではこのstructに書く設定を追加していく
	c := cli.NewCLI("memo", "1.0.0")

	// ユーザの引数を登録する
	c.Args = os.Args[1:]

	// サブコマンドを登録する
	// cli.CommandFactoryという関数である
	c.Commands = map[string]cli.CommandFactory{
		"add": func() (cli.Command, error) {
			return &memo.Memo{
				Name:      "ここ",
				Text:      "ここだよ",
				CreatedAt: time.Now(),
			}, nil
		},
		"delete": func() (cli.Command, error) {
			return &memo.Memo{
				Name:      "",
				Text:      "",
				CreatedAt: time.Time{},
			}, nil
		},
	}

	// コマンドを実行する
	exitCode, err := c.Run()
	if err != nil {
		fmt.Printf("cilコマンドの実行に失敗しました: %s\n", err)
	}

	// コマンド入力されたときのみexit
	if exitCode != 127 {
		os.Exit(exitCode)
	}
}
