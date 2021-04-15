package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"
	"sync"
	"time"

	"github.com/mattn/go-shellwords"
)

// func tr(src io.Reader, dst io.Writer, errDst io.Writer) error {
// 	cmd := exec.Command("tr", "a-z", "A-Z")
// 	// 実行するコマンド tr a-z A-Z
// 	stdin, _ := cmd.StdinPipe()
// 	stdout, _ := cmd.StdoutPipe()
// 	stderr, _ := cmd.StderrPipe()
// 	err := cmd.Start() //コマンドの実行を開始する
// 	if err != nil {
// 		return err
// 	}
// 	var wg sync.WaitGroup
// 	wg.Add(3)
// 	go func() {
// 		// コマンドの標準入力にsrcからコピーする
// 		_, err := io.Copy(stdin, src)
// 		if e, ok := err.(*os.PathError); ok && e.Err == syscall.EPIPE {
// 			// ignore EPIPE
// 		} else if err != nil {
// 			log.Println("failed to write to STDIN", err)
// 		}
// 		stdin.Close()
// 		wg.Done()
// 	}()
// 	go func() {
// 		io.Copy(dst, stdout)
// 		stdout.Close()
// 		wg.Done()
// 	}()
// 	go func() {
// 		io.Copy(errDst, stderr)
// 		fmt.Println("stderr")
// 		stderr.Close()
// 		wg.Done()
// 	}()
// 	wg.Wait()
// 	//標準入出力のI/Oを行うgoroutineが全て終わるまで待つ
// 	return cmd.Wait()
// 	// コマンドの終了を待つ
// }

// パーミッション確認
func execLs() {
	out, _ := exec.Command("ls", "-l", "config").Output()
	fmt.Println(string(out))
}

// go-shellwords 一つずつ出力を確認しながら実行したいとき使用
func shell() {
	aregs, err := shellwords.Parse("ls -l config")
	// argsは["ls", "-l", "config"]となる
	if err != nil {
		log.Println(err)
	}

	out, err := exec.Command(aregs[0], aregs[1:]...).Output()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(out))
}

func getHTTP(url string, dst io.Writer) error {
	// 10秒でタイムアウトするContextを作る
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client := &http.Client{}
	// client := &http.Client{
	// 	// 10秒でタイムアウトする
	// 	Timeout: 10 * time.Second,
	// }
	req, _ := http.NewRequest("GET", url, nil)
	// contextを与えたリクエストを使って実行
	resp, err := client.Do(req.WithContext(ctx)) // 呼び出す度に10秒待つ
	// resp, err := client.Do(req)
	if err != nil {
		// レスポンスヘッダーの取得までにエラー
		return err
	}
	defer resp.Body.Close()
	_, err = io.Copy(dst, resp.Body)
	// ボディ取得完了までにエラー
	return err
}

// ゴルーチンの処理
var wg sync.WaitGroup

func goroutine() {
	queue := make(chan string)

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go fetchURL(queue)
	}

	queue <- "https://www.example.com"
	queue <- "https://www.example.net"
	queue <- "https://www.example.net/foo"
	queue <- "https://www.example.net/bar"

	close(queue)
	wg.Wait()
}

func fetchURL(queue chan string) {
	for url := range queue {
		fmt.Println("fetching", url)
	}
	fmt.Println("worker exit")
	wg.Done()
	return
}
