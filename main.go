package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
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

	// CORS設定
	corsOpts := cors.New(cors.Options{
		AllowedOrigins:   []string{os.Getenv("FRONT_URL")},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})
	c := corsOpts.Handler(r)

	// HTTPサーバーを起動する
	log.Println("Starting server on :8088...")
	if err := http.ListenAndServe(":8088", c); err != nil {
		logging.WriteErrorLog(err.Error(), true)
		log.Fatal(err)
	}
}
