package main

import (
	"net/http"

	"github.com/smalake/kakebo-api/lib"
	"github.com/smalake/kakebo-api/web"
)

func main() {
	// ルーティングを呼び出す
	router := web.NewRouter()
	// HTTPサーバーを起動する
	if err := http.ListenAndServe(":8088", router); err != nil {
		lib.WriteErrorLog(err.Error(), true)
	}
}
