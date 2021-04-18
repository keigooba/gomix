package c_memo

import (
	"flag"
	"fmt"
	cliCmd "gomix/cli"
	"gomix/config"
	"gomix/pkg/memo"
	"log"
	"os"
	"strconv"
	"time"

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
	return "コマンド [option...] を入力する"
}

func flagSet(args []string) (intvar int, strvar string) {

	// selectのオプションフラグの設定
	flags := flag.NewFlagSet(os.Args[2], flag.ContinueOnError)
	flags.IntVar(&intvar, "id", 0, "ID")
	flags.StringVar(&strvar, "name", "", "名前")
	flags.StringVar(&strvar, "text", "", "テキスト")
	if err := flags.Parse(args); err != nil {
		log.Println(err)
	}

	return intvar, strvar
}

func (c *AddCommand) Run(args []string) int {

	var id int
	var name, text string

	switch c.Cmd {
	case "add":
		// 作成
		if len(args) > 0 {
			name = args[0]
		}
		if len(args) > 1 {
			text = args[1]
		}
		memoEx := memo.Memo{}
		memoEx.Name = name
		memoEx.Text = text
		memoEx.CreatedAt = time.Now()
		config.Db.Create(&memoEx)
	case "select":
		// 検索

		var search string
		var data interface{}
		// オプションフラグを設定し、値をセット
		intvar, strvar := flagSet(args)

		if len(args) > 0 {
			switch args[0] {
			case "-id":
				data = intvar
				search = "id=?"
			case "-name":
				data = strvar
				search = "name=?"
			case "-text":
				data = strvar
				search = "text=?"
			}
		}

		memoEx := []memo.Memo{}
		config.Db.Find(&memoEx, search, data)
		fmt.Println("検索結果\n", memoEx)
	case "edit":
		// 編集
		if len(args) > 2 {
			// 数値に変換
			id, _ = strconv.Atoi(args[2])
		}

		// オプションフラグを設定し、値をセット
		_, strvar := flagSet(args)
		memoExBefore := memo.Memo{}
		memoExBefore.ID = id
		memoExAfter := memoExBefore
		config.Db.First(&memoExAfter)

		if len(args) > 2 {
			switch args[0] {
			case "-name":
				memoExAfter.Name = strvar
			case "-text":
				memoExAfter.Text = strvar
			}
		}
		config.Db.Save(&memoExAfter)
	case "delete":
		// 削除
		if len(args) > 0 {
			// 数値に変換
			id, _ = strconv.Atoi(args[0])
		}
		memoEx := memo.Memo{}
		memoEx.ID = id
		config.Db.First(&memoEx)
		config.Db.Delete(&memoEx)
	}

	fmt.Printf("%sを実行しました\n", c.Cmd)
	return 0 //正常終了
}

func CliMemo() {
	// CLI structを生成する
	// 以下ではこのstructに書く設定を追加していく
	c := cli.NewCLI("-memo", cliCmd.Version)

	var cmd string
	if len(os.Args) > 2 {
		// ユーザの引数を登録する
		c.Args = os.Args[2:]
		cmd = os.Args[2]
	}

	// サブコマンドを登録する
	// cli.CommandFactoryという関数である
	c.Commands = map[string]cli.CommandFactory{
		"add": func() (cli.Command, error) {
			return &AddCommand{
				Cmd: cmd,
			}, nil
		},
		"select": func() (cli.Command, error) {
			return &AddCommand{
				Cmd: cmd,
			}, nil
		},
		"edit": func() (cli.Command, error) {
			return &AddCommand{
				Cmd: cmd,
			}, nil
		},
		"delete": func() (cli.Command, error) {
			return &AddCommand{
				Cmd: cmd,
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
