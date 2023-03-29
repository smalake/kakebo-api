package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/smalake/kakebo-api/handlers"
	"github.com/smalake/kakebo-api/middleware"
)

// ルーティング
func Routing(r *mux.Router) {
	// 認証
	r.HandleFunc("/login", handlers.LoginHandler).Methods(http.MethodPost)
	r.HandleFunc("/register", handlers.RegisterHandler).Methods(http.MethodPost)

	// イベント
	r.HandleFunc("/events", middleware.AuthCheck(handlers.GetEvents)).Methods(http.MethodGet)
	r.HandleFunc("/events", middleware.AuthCheck(handlers.CreateEvent)).Methods(http.MethodPost)
	r.HandleFunc("/events", middleware.AuthCheck(handlers.EditEvent)).Methods(http.MethodPut)
	r.HandleFunc("/events", middleware.AuthCheck(handlers.DeleteEvent)).Methods(http.MethodDelete)

	// 表示名
	r.HandleFunc("/display-name", middleware.AuthCheck(handlers.GetName)).Methods(http.MethodGet)
	r.HandleFunc("/display-name", middleware.AuthCheck(handlers.EditName)).Methods(http.MethodPut)

	// セットアップ
	r.HandleFunc("/setup", middleware.AuthCheck(handlers.CheckSetup)).Methods(http.MethodGet)
	r.HandleFunc("/setup", middleware.AuthCheck(handlers.CreateGroup)).Methods(http.MethodPost)

}
