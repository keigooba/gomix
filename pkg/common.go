package pkg

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
)

func Index(w http.ResponseWriter, _ *http.Request) {
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
