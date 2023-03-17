package main

import (
	"net/http"

	"github.com/joho/godotenv"
	"github.com/smalake/kakebo-api/server"
	"github.com/smalake/kakebo-api/utils/firebase"
	"github.com/smalake/kakebo-api/utils/logging"
)

func main() {
	// .envからDB設定を読み込む
	err := godotenv.Load(".env")
	if err != nil {
		logging.WriteErrorLog(err.Error(), true)
	}
	// Firebase SDKの初期化処理
	firebase.InitFirebase()

	// ルーティングを呼び出す
	router := server.NewRouter()
	// HTTPサーバーを起動する
	if err := http.ListenAndServe(":8088", router); err != nil {
		logging.WriteErrorLog(err.Error(), true)
	}
}
