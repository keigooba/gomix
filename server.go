package main

//go:generate statik -src=doc

import (
	"gomix/config"
	"gomix/pkg"
	"gomix/pkg/change"
	"gomix/pkg/memo"
	"net/http"

	stats_api "github.com/fukata/golang-stats-api-handler"
	"github.com/markbates/pkger"
)

func StartMainServer() error {
	// doc以下のファイル読み込み
	dir := pkger.Dir(config.Config.Static) //バイナリファイルに静的ファイルを埋め込める
	files := http.FileServer(http.Dir(dir))
	http.Handle("/static/", http.StripPrefix("/static/", files))
	http.HandleFunc("/json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "config/info.json") //ファイルにアクセス
	})

	http.HandleFunc("/", pkg.Index)
	http.HandleFunc("/change", change.Index)
	http.HandleFunc("/memo", memo.Index)
	http.HandleFunc("/data/", memo.Open)
	http.HandleFunc("/stats", stats_api.Handler)
	return http.ListenAndServe(config.Config.Port, nil)
}
