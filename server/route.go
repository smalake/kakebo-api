package server

import (
	"net/http"

	"github.com/smalake/kakebo-api/handlers"
	"github.com/smalake/kakebo-api/middleware"
	"github.com/smalake/kakebo-api/utils/logging"
)

func NewRouter() *http.ServeMux {
	// Webサーバの起動
	mux := http.NewServeMux()
	println("Server Start Port:8088")
	logging.WriteErrorLog("Server Start Port:8088", false)
	// ルーティング
	mux.HandleFunc("/hello", handlers.HelloHandler)
	mux.HandleFunc("/user", handlers.UserHandler)
	mux.HandleFunc("/login", handlers.LoginHandler)
	mux.HandleFunc("/user-data", middleware.AuthCheck(handlers.UserDataHandler))
	return mux
}
