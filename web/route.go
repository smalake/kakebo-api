package web

import (
	"net/http"

	"github.com/smalake/kakebo-api/handler"
	"github.com/smalake/kakebo-api/lib"
)

func NewRouter() *http.ServeMux {
	// Webサーバの起動
	mux := http.NewServeMux()
	println("Server Start Port:8088")
	lib.WriteErrorLog("Server Start Port:8088", false)
	// ルーティング
	mux.HandleFunc("/hello", handler.HelloHandler)
	return mux
}
