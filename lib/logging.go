package lib

import (
	"log"
	"os"
	"time"
)

// エラーログをログファイルへと書き込む
func WriteErrorLog(message string, isError bool) {
	// ログファイル名用に今日の日付を取得する
	location, _ := time.LoadLocation("Asia/Tokyo")
	date := time.Now().In(location)
	today := date.Format("2006-01-02") //年月日を取得するフォーマットへと変換
	file, err := os.OpenFile("log/"+today+".log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 日時とファイル名を出力させる
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)

	if isError {
		// エラーログを出力
		logger := log.New(file, "Error: ", log.LstdFlags)
		logger.Println(message)

	} else {
		// エラー以外のログを出力
		logger := log.New(file, "Info: ", log.LstdFlags)
		logger.Println(message)

	}
}
