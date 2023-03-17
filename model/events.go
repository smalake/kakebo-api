package model

import (
	"errors"
	"fmt"
	"time"

	"github.com/smalake/kakebo-api/utils/logging"
)

type Events struct {
	ID        int       `json:"id"`
	UID       string    `json:"uid"`
	Category1 int       `json:"category1"`
	Category2 int       `json:"category2"`
	Amount1   int       `json:"amount1"`
	Amount2   int       `json:"amount2"`
	Date      time.Time `json:"date" time_format:"2006-01-02"`
	StoreName string    `json:"store_name"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// 該当ユーザのイベントを全て取得する
func (e *Events) GetEvents(uid string) error {
	db := ConnectDB()
	sqlDb, err := db.DB() //コネクションクローズ用
	if err != nil {
		return errors.New("DBとの接続に失敗しました。")
	}
	defer sqlDb.Close()

	err = db.Where("uid = ?", uid).Scan(&e).Error
	if err != nil {
		logging.WriteErrorLog(err.Error(), true)
		return err
	}
	// TODO
	fmt.Println(e)

	return nil
}

// イベントを新規作成
func (e *Events) CreateEvent(uid string) error {
	e.UID = uid
	db := ConnectDB()
	sqlDb, err := db.DB() //コネクションクローズ用
	if err != nil {
		return errors.New("DBとの接続に失敗しました。")
	}
	defer sqlDb.Close()

	err = db.Table("events").Create(e).Error
	if err != nil {
		return err
	}
	return nil
}
