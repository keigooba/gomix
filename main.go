package main

import (
	"net/http"

	stats_api "github.com/fukata/golang-stats-api-handler"
)

func main() {
	http.HandleFunc("/api/stats", stats_api.Handler)
	http.ListenAndServe(":8080", nil)
}
