package c_memo

import (
	"fmt"
	cliCmd "gomix/cli"
	"log"
	"os"

	"github.com/mitchellh/cli"
)

type AddCommand struct {
	Cmd string
}

// 簡単なコマンドの説明を記述
func (c *AddCommand) Synopsis() string {
	return "メモを作成・検索・編集・削除を行う"
}

// 使い方 詳細なヘルプメッセージを返す
func (c *AddCommand) Help() string {
	return "コマンドを入力する"
}

func (c *AddCommand) Run(args []string) int {

	// 作成
	// config.Db.Create(&m)

	return 0 //正常終了
}

func CliMemo() {
	// CLI structを生成する
	// 以下ではこのstructに書く設定を追加していく
	c := cli.NewCLI("memo", "1.0.0")

	// ユーザの引数を登録する
	c.Args = os.Args[1:]

	if cliCmd.ManageMemo {
		log.Println("ManageMemo")
	}

	// サブコマンドを登録する
	// cli.CommandFactoryという関数である
	c.Commands = map[string]cli.CommandFactory{
		"add": func() (cli.Command, error) {
			return &AddCommand{
				Cmd: os.Args[1],
			}, nil
		},
		"select": func() (cli.Command, error) {
			return &AddCommand{
				Cmd: os.Args[1],
			}, nil
		},
		"edit": func() (cli.Command, error) {
			return &AddCommand{
				Cmd: os.Args[1],
			}, nil
		},
		"delete": func() (cli.Command, error) {
			return &AddCommand{
				Cmd: os.Args[1],
			}, nil
		},
	}

	// コマンドを実行する
	exitCode, err := c.Run()
	if err != nil {
		fmt.Printf("cliコマンドの実行に失敗しました: %s\n", err)
	}

	// コマンド入力されたときのみexit
	if exitCode != 127 {
		os.Exit(exitCode)
	}
}
