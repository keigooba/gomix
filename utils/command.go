package utils

import (
	"fmt"
	"os/exec"
	"runtime"

	"golang.org/x/text/encoding/japanese"
)

// コマンドの実行
func Command() error {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("start.bat")
	} else {
		cmd = exec.Command("sh", "start.sh")
	}
	b, err := cmd.CombinedOutput() //標準出力・標準エラー出力両方取れる
	if err != nil {
		return err
	}
	// windowsでのコマンド実行の文字列はCP932のため、UTF-8に変換して受け取る
	if runtime.GOOS == "windows" {
		b, err = japanese.ShiftJIS.NewDecoder().Bytes(b)
		if err != nil {
			return err
		}
	}
	fmt.Println(string(b))
	fmt.Println("コマンドが正常に実行されました。")
	return nil
}
