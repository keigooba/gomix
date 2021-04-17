package main

import (
	"gomix/pkg/memo"
	"os"

	"github.com/mitchellh/cli"
)

func main() {
	// CLI structを生成する
	// 以下ではこのstructに書く設定を追加していく
	c := cli.NewCLI("memo", "1.0.0")

	// ユーザの引数を登録する
	c.Args = os.Args[1:]

	// サブコマンドを登録する
	// cli.CommandFactoryという関数である
	c.Commands = map[string]cli.CommandFactory{
		"add": func() (cli.Command, error) {
			return &memo.Memo{}, nil
		},
	}

	// コマンドを実行する
	// exitCode, err := c.Run()
	// if err != nil {
	// 	fmt.Printf("Failed to execute: %s\n", err)
	// }

	// os.Exit(exitCode)
}
