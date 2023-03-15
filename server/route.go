package server

import (
	"net/http"

	"github.com/smalake/kakebo-api/handlers"
	"github.com/smalake/kakebo-api/utils/logging"
)

func NewRouter() *http.ServeMux {
	// Webサーバの起動
	mux := http.NewServeMux()
	println("Server Start Port:8088")
	logging.WriteErrorLog("Server Start Port:8088", false)
	// ルーティング
	mux.HandleFunc("/login", handlers.LoginHandler)
	// mux.HandleFunc("/register", middleware.AuthCheck(handlers.RegisterHandler))
	mux.HandleFunc("/register", handlers.RegisterHandler)
	return mux
}
