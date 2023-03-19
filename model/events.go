package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/smalake/kakebo-api/utils/logging"
)

// DBとやり取りを行うための構造体
type Event struct {
	ID        int       `json:"id"`
	UID       string    `json:"uid"`
	Category  int       `json:"category"`
	Amount    int       `json:"amount"`
	Date      time.Time `json:"date" time_format:"2006-01-02"`
	StoreName string    `json:"store_name" gorm:"column:store_name"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// フロント側から直接受け取るための構造体
type Events struct {
	ID        int       `json:"id"`
	UID       string    `json:"uid"`
	Category1 int       `json:"category1"`
	Category2 int       `json:"category2"`
	Amount1   int       `json:"amount1"`
	Amount2   int       `json:"amount2"`
	Date      time.Time `json:"date" time_format:"2006-01-02"`
	StoreName string    `json:"store_name"`
}

// 該当ユーザのイベントを全て取得する
func (e *Event) GetEvents(uid string) ([]byte, error) {
	db := ConnectDB()
	sqlDb, err := db.DB() //コネクションクローズ用
	if err != nil {
		return nil, errors.New("DBとの接続に失敗しました。")
	}
	defer sqlDb.Close()

	var events []Event
	err = db.Where("uid = ?", uid).Find(&events).Error
	if err != nil {
		logging.WriteErrorLog(err.Error(), true)
		return nil, err
	}
	// 送信用に変換
	eventMap := make(map[time.Time][]map[string]interface{})
	for _, event := range events {
		date := event.Date.UTC() // 日付をUTCに変換
		// マップ内の日付キーが存在しない場合は初期化する
		if _, ok := eventMap[date]; !ok {
			eventMap[date] = make([]map[string]interface{}, 0)
		}
		// イベントデータをマップに変換して追加
		eventMap[date] = append(eventMap[date], map[string]interface{}{
			"id":       event.ID,
			"spending": event.Amount,
			"category": event.Category,
			"store":    event.StoreName,
		})
	}

	// JSONへと変換
	jsonEvents, err := json.Marshal(eventMap)
	if err != nil {
		logging.WriteErrorLog(err.Error(), true)
		return nil, err
	}
	return jsonEvents, nil
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

	// 支出金額の合計
	amount := e.Amount1

	// 2つ目の支出が存在する場合、別途登録する
	if e.Amount2 != 0 {
		// 2つ目の金額分合計金額から引く
		amount = e.Amount1 - e.Amount2
		sub := &Event{
			UID:       e.UID,
			Category:  e.Category2,
			Amount:    e.Amount2,
			Date:      e.Date,
			StoreName: e.StoreName,
		}
		err = db.Table("events").Create(sub).Error
		if err != nil {
			return err
		}
	}

	// DB登録用に切り出して登録
	main := &Event{
		UID:       e.UID,
		Category:  e.Category1,
		Amount:    amount,
		Date:      e.Date,
		StoreName: e.StoreName,
	}
	err = db.Table("events").Create(main).Error
	if err != nil {
		return err
	}

	return nil
}

// イベントを編集
func (e *Event) UpdateEvent() error {
	db := ConnectDB()
	sqlDb, err := db.DB() //コネクションクローズ用
	if err != nil {
		return errors.New("DBとの接続に失敗しました。")
	}
	defer sqlDb.Close()
	fmt.Println(e.Date)

	err = db.Model(&e).Updates(Event{
		Category:  e.Category,
		Amount:    e.Amount,
		Date:      e.Date,
		StoreName: e.StoreName,
	}).Error
	if err != nil {
		logging.WriteErrorLog(err.Error(), true)
		return err
	}
	return nil
}
