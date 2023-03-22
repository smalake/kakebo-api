package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
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

	r := mux.NewRouter()
	// ルーティングを呼び出す
	server.Routing(r)

	// HTTPサーバーを起動する
	log.Println("Starting server on :8088...")
	if err := http.ListenAndServe(":8088", r); err != nil {
		logging.WriteErrorLog(err.Error(), true)
		log.Fatal(err)
	}
}
