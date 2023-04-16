package events

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/smalake/kakebo-api/model"
	"github.com/smalake/kakebo-api/utils/logging"
)

// DBからデータを取得するための構造体
type GetEvent struct {
	ID             int       `json:"id"`
	Category       int       `json:"category"`
	Amount         int       `json:"amount"`
	Date           time.Time `json:"date" time_format:"2006-01-02"`
	StoreName      string    `json:"storeName" gorm:"column:store_name"`
	CreateUser     string    `json:"createUser"`
	UpdateUser     string    `json:"updateUser"`
	CreateUserName string
	UpdateUserName string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

// 該当ユーザの所属しているグループのイベントを全て取得する
func (e *GetEvent) GetEvents(uid string) ([]byte, error) {
	db := model.ConnectDB()
	sqlDb, err := db.DB() //コネクションクローズ用
	if err != nil {
		return nil, errors.New("DBとの接続に失敗しました。")
	}
	defer sqlDb.Close()

	var events []GetEvent
	err = db.Table("events").
		Joins("INNER JOIN users AS create_user ON events.create_user = create_user.uid").
		Joins("LEFT JOIN users AS update_user ON events.update_user = update_user.uid").
		Select("events.*, create_user.name AS create_user_name, update_user.name AS update_user_name").
		Where("create_user.group_id = events.group_id AND create_user.uid = ?", uid).Find(&events).Error
	if err != nil {
		logging.WriteErrorLog(err.Error(), true)
		return nil, err
	}
	// 送信用に変換
	eventMap := make(map[string][]map[string]interface{})
	for _, event := range events {
		date := event.Date.Format("2006-01-02")
		// マップ内の日付キーが存在しない場合は初期化する
		if _, ok := eventMap[date]; !ok {
			eventMap[date] = make([]map[string]interface{}, 0)
		}
		// イベントデータをマップに変換して追加
		eventMap[date] = append(eventMap[date], map[string]interface{}{
			"id":         event.ID,
			"amount":     event.Amount,
			"category":   event.Category,
			"storeName":  event.StoreName,
			"createUser": event.CreateUserName,
			"updateUser": event.UpdateUserName,
			"createDate": event.CreatedAt,
			"updateDate": event.UpdatedAt,
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
