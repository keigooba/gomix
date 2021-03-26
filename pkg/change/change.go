package change

import (
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {

	// htmlファイルの読み込み
	generateHTML(w, nil, "change/index")
}
