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
	"time"

	"gomix/pkg"
)

type Data struct {
	File []string
}

// File.Closeのエラーチェックを行う為、定義
func Close(f *os.File) {
	err := f.Close()
	if err != nil {
		log.Println(err)
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

		var data Data
		// 作成済みファイルを読み込む
		if files, err := ioutil.ReadDir("doc/memo/data"); err == nil {
			for _, file := range files {
				data.File = append(data.File, file.Name())
			}
		}
		// htmlファイルの読み込み
		pkg.GenerateHTML(w, data, "memo/index")

	} else if r.Method == "POST" {

		// "/"でパス文字列を結合しない 物理パスを操作する場合filepathを使う
		dir := filepath.Join(pkg.Getpath(), "doc", "memo", "data")
		// dirのディレクトリを作成する MkdirAll 必要な親ディレクトリ全てを作成する
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			log.Println(err)
		}

		// 以降はパスを物理パスとして扱うのでfilepathパッケージを使う filepath.Base ファイル名を得る
		time := time.Now().Unix()
		stringTime := strconv.FormatInt(time, 10)
		docPath := "/doc/memo/data/" + "memo_" + stringTime + ".html"
		name := filepath.Join(pkg.Getpath(), "doc", "memo", "data", filepath.Base(docPath))
		memo := r.FormValue("memo")

		//ファイルに書き込む
		err = ioutil.WriteFile(name, []byte(memo), 0664)
		if err != nil {
			log.Println(err)
		}

		// ファイルを開く
		f, err := os.Open(name)
		if err != nil {
			log.Println(err)
		}
		defer Close(f)

		_, err = io.Copy(w, f) //Writerにファイルを書き出す
		if err != nil {
			log.Println("ファイルの書き出しに失敗しました。")
		}
	}
}

func Open(w http.ResponseWriter, r *http.Request) {
	// httpリクエストは論理パスなのでpathを使う
	if ok, err := path.Match("/data/memo_*.html", r.URL.Path); err != nil || !ok {
		http.NotFound(w, r)
		return
	}

	// 指定したファイルを開く
	name := filepath.Join(pkg.Getpath(), "doc", "memo", "data", filepath.Base(r.URL.Path))
	f, err := os.Open(name)
	if err != nil {
		log.Println(err)
	}
	defer Close(f)

	_, err = io.Copy(w, f) //Writerにファイルを書き出す
	if err != nil {
		log.Println("ファイルの書き出しに失敗しました。")
	}
}

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
	defer Close(f)

	return nil
}

// ファイルの作成・名前変更・deferの操作
func MytemFile() (*os.File, error) {
	file, err := ioutil.TempFile("", "temp") //適当なディレクトリ/tempランダム文字列 ファイルの作成
	if err != nil {
		return nil, err
	}
	// defer file.Close() //Closeが遅い deferはfunc()の呼び出し形式を取る 引数にはdeferを呼び出した時点の値が入る
	Close(file) //Renameを実行するため、すぐ閉じる

	// file.Close()するとwindowsではファイルが開かれていると認識され、Renameできない
	if err = os.Rename(file.Name(), file.Name()+".go"); err != nil {
		return nil, err
	}
	return file, nil
}
