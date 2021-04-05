package main

import (
	"gomix/config"
	"gomix/pkg"
	"gomix/pkg/change"
	"gomix/pkg/memo"
	"net/http"

	stats_api "github.com/fukata/golang-stats-api-handler"
)

func StartMainServer() error {
	// doc以下のファイル読み込み
	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))

	http.HandleFunc("/", pkg.Index)
	http.HandleFunc("/change", change.Index)
	http.HandleFunc("/memo", memo.Index)
	http.HandleFunc("/data/", memo.Open)
	http.HandleFunc("/stats", stats_api.Handler)
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
