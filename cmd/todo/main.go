package main

import (
	"fmt"
	"os"

	"github.com/mitchellh/cli"
)

type AddCommand struct {
	Debug string
}

// 簡単なコマンドの説明を記述
func (c *AddCommand) Synopsis() string {
	return "Add todo task to list"
}

// 使い方 詳細なヘルプメッセージを返す
func (c *AddCommand) Help() string {
	return "data"
}

func (c *AddCommand) Run(args []string) int {
	// TODOを追加するコード
	fmt.Println("ここ")
	return 0
}

func main() {
	// CLI structを生成する
	// 以下ではこのstructに書く設定を追加していく
	c := cli.NewCLI("todo", "0.1.0")

	// ユーザの引数を登録する
	c.Args = os.Args[1:]

	// サブコマンドを登録する
	// cli.CommandFactoryという関数である
	var debug = "koko"
	c.Commands = map[string]cli.CommandFactory{
		"add": func() (cli.Command, error) {
			return &AddCommand{
				Debug: debug,
			}, nil
		},
	}

	// コマンドを実行する
	exitCode, err := c.Run()
	if err != nil {
		fmt.Printf("Failed to execute: %s\n", err)
	}

	os.Exit(exitCode)
}
