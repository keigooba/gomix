package change

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
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
