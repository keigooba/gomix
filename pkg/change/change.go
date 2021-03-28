package change

import (
	"fmt"
	"net/http"
	"strconv"
)

// type Number struct {
// 	Num int `json:"number"`
// }

type Data struct {
	S2    string
	S16   string
	Error Err
}

type Err struct {
	Number string
}

func Change(number int) (data Data) {
	s2 := fmt.Sprintf("%b", number)
	s16 := fmt.Sprintf("%x", number)
	data = Data{
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
		// 数値に変換
		number, err := strconv.Atoi(r.FormValue("number"))
		if err != nil {
			// 数値でなければerrorを返す エラーメッセージ作成
			err := Err{
				Number: "数値で入力して下さい",
			}
			// エラーメッセージをdataに格納
			var data Data
			data.Error = err
			generateHTML(w, data, "change/index")
		} else {
			data := Change(number)
			generateHTML(w, data, "change/index")
		}
		// decoder := json.NewDecoder(r.Body)
		// fmt.Println(decoder)
		// var n Number
		// err := decoder.Decode(&n)
		// if err != nil {
		// 	fmt.Println(err)
		// }
		// fmt.Println(n.Num)

	}
}
