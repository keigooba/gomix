package memo

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"gomix/config"
	"gomix/pkg"
)

type Data struct {
	Json []string
	Txt  []string
}

func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

		var data Data
		// 作成済みファイルを読み込む
		slice := []string{"json", "txt"}
		for _, extension := range slice {
			if files, err := ioutil.ReadDir("doc/memo/data/" + extension); err == nil {
				for _, file := range files {
					if extension == "json" {
						data.Json = append(data.Json, file.Name())
					} else if extension == "txt" {
						data.Txt = append(data.Txt, file.Name())
					}
				}
			}
		}

		// ファイルの読み込み
		pkg.GenerateHTML(w, data, "memo/index")

	} else if r.Method == "POST" {

		memo := r.PostFormValue("memo")
		var extension string
		if string(memo[0]) == "[" {
			extension = "json"
		} else {
			extension = "txt"
		}

		// "/"でパス文字列を結合しない 物理パスを操作する場合filepathを使う
		dir := filepath.Join(pkg.Getpath(), "doc", "memo", "data", extension)
		// dirのディレクトリを作成する MkdirAll 必要な親ディレクトリ全てを作成する
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			log.Println(err)
		}

		// 以降はパスを物理パスとして扱うのでfilepathパッケージを使う filepath.Base ファイル名を得る
		time := time.Now().Unix()
		stringTime := strconv.FormatInt(time, 10)
		docPath := "/doc/memo/data/" + "memo_" + stringTime + "." + extension
		name := filepath.Join(pkg.Getpath(), "doc", "memo", "data", extension, filepath.Base(docPath))
		err = ioutil.WriteFile(name, []byte(memo), 0664)
		if err != nil {
			log.Println(err)
		}
		url := config.Config.URL + r.URL.Path
		http.Redirect(w, r, url, http.StatusSeeOther) //キャッシュを残したくないので、303指定
	}
}

func Open(w http.ResponseWriter, r *http.Request) {
	// httpリクエストは論理パスなのでpathを使う
	var extension string
	if strings.Contains(r.URL.Path, "json") {
		extension = "json"
	} else if strings.Contains(r.URL.Path, "txt") {
		extension = "txt"
	}

	if ok, err := path.Match("/data/"+extension+"/memo_*."+extension, r.URL.Path); err != nil || !ok {
		http.NotFound(w, r)
		return
	}

	// 指定したファイルを開く
	name := filepath.Join(pkg.Getpath(), "doc", "memo", "data", extension, filepath.Base(r.URL.Path))
	f, err := os.Open(name)
	if err != nil {
		log.Println(err)
	}
	defer pkg.Close(f)

	_, err = io.Copy(w, f) //Writerにファイルを書き出す
	if err != nil {
		log.Println("ファイルの書き出しに失敗しました。")
	}
}

//以下機能は主要機能に導入なし

// 一時ディレクトリの作成・削除
func Dosomething() error {
	err := os.MkdirAll("newdir", 0755)
	if err != nil {
		log.Println(err)
	}
	//  (2)ディレクトリ削除
	defer func() {
		err := os.RemoveAll("newdir")
		if err != nil {
			log.Println(err)
		}
	}()

	f, err := os.Create("newdir/newfile")
	if err != nil {
		log.Println(err)
	}
	// (1)ファイルハンドルが閉じられる
	defer pkg.Close(f)

	return nil
}

// ファイルの作成・名前変更・deferの操作
func MytemFile() (*os.File, error) {
	file, err := ioutil.TempFile("", "temp") //適当なディレクトリ/tempランダム文字列 ファイルの作成
	if err != nil {
		return nil, err
	}
	// defer file.Close() //Closeが遅い deferはfunc()の呼び出し形式を取る 引数にはdeferを呼び出した時点の値が入る
	pkg.Close(file) //Renameを実行するため、すぐ閉じる

	// defer file.Close()するとwindowsではファイルが開かれていると認識され、Renameできない
	if err = os.Rename(file.Name(), file.Name()+".go"); err != nil {
		return nil, err
	}
	return file, nil
}
