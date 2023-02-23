package main

import (
	"net/http"

	"github.com/joho/godotenv"
	"github.com/smalake/kakebo-api/server"
	"github.com/smalake/kakebo-api/utils/logging"
)

func main() {
	// .envからDB設定を読み込む
	err := godotenv.Load(".env")
	if err != nil {
		logging.WriteErrorLog(err.Error(), true)
	}
	// ルーティングを呼び出す
	router := server.NewRouter()
	// HTTPサーバーを起動する
	if err := http.ListenAndServe(":8088", router); err != nil {
		logging.WriteErrorLog(err.Error(), true)
	}
}
