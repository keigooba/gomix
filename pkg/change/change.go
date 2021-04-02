package change

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"sync"
	"time"

	"golang.org/x/text/width"
)

// type Number struct {
// 	Num int `json:"number"`
// }

type Data struct {
	Numbers []Number
	Error   Err
}

type Number struct {
	S2  string
	S16 string
}

type Err struct {
	Number string
}

// 正規表現
var numReg = regexp.MustCompile(`[0-9０-９]`)

func Change(number int) (data Number) {
	s2 := fmt.Sprintf("「%v」2進数:%b", number, number)
	s16 := fmt.Sprintf("「%v」16進数:%x", number, number)
	data = Number{
		S2:  s2,
		S16: s16,
	}
	return data
}

// Index 2進数・18進数変換フォーム
func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// htmlファイルの読み込み
		generateHTML(w, nil, "change/index")

	} else if r.Method == "POST" {
		var data Data
		value := r.FormValue("number")
		// 正規表現を用いて判定
		if numReg.MatchString(value) {
			// 全角を半角に変換
			num := width.Narrow.String(value)
			// 数値に変換
			number, err := strconv.Atoi(num)
			if err != nil {
				log.Println(err)
			}
			workers := 5
			ch := make(chan int, workers)
			var numbers []Number
			var mutex = &sync.Mutex{}
			defer close(ch)

			// 数値+19の20個まで5つ並行で処理
			for i := 0; i < workers; i++ {
				go func() {
					for num := range ch {
						number := Change(num)
						// 処理をロックする
						mutex.Lock()
						numbers = append(numbers, number)
						mutex.Unlock()
					}
				}()
			}

			for i := 0; i < 20; i++ {
				ch <- number + i
			}

			// ゴルーチンの処理を1秒待つ
			time.Sleep(time.Second * 1)
			data.Numbers = numbers
			generateHTML(w, data, "change/index")
		} else {
			// 数値でなければerrorを返す エラーメッセージ作成
			err := Err{
				Number: "数字を入力して下さい",
			}
			// Numberのエラーメッセージをdataに格納
			data.Error = err
			generateHTML(w, data, "change/index")
		}

	}
}
