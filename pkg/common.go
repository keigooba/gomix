package pkg

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/tcnksm/go-latest"
)

func Index(w http.ResponseWriter, _ *http.Request) {
	// 最新版バージョンチェック
	upVersion := "0.1.0" //buildされるアプリケーションのバージョン
	metaTag := &latest.HTMLMeta{
		URL:  "http://localhost:8080/",
		Name: "gomix",
	}
	res, _ := latest.Check(metaTag, upVersion)
	if res.Outdated {
		fmt.Printf("%s is not latest, you should upgrade to %s: %s\n", upVersion, res.Current, res.Meta.Message)
	}
	GenerateHTML(w, nil, "index")
}

func Getpath() (cwd string) {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return cwd
}

func GenerateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("doc/%s.html", file))
	}
	// ヘッダー・フッターを追加
	files = append(files, "doc/_header.html", "doc/_footer.html")

	templates := template.Must(template.ParseFiles(files...))
	err := templates.Execute(w, data)
	if err != nil {
		log.Println(err)
	}
}
