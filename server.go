package main

import (
	"gomix/config"
	"gomix/pkg/change"
	"net/http"

	stats_api "github.com/fukata/golang-stats-api-handler"
)

func StartMainServer() error {
	// doc以下のファイル読み込み
	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))

	http.HandleFunc("/", change.Index)
	http.HandleFunc("/stats", stats_api.Handler)
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
