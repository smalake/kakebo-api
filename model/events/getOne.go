package events

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/smalake/kakebo-api/model"
	"github.com/smalake/kakebo-api/utils/logging"
)

// DBからデータを取得するための構造体
type GetEventOne struct {
	ID             int       `json:"id"`
	Category       int       `json:"category"`
	Amount         int       `json:"amount"`
	Date           time.Time `json:"date" time_format:"2006-01-02"`
	StoreName      string    `json:"storeName" gorm:"column:store_name"`
	CreateUser     string
	UpdateUser     string
	CreateUserName string
	UpdateUserName string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

// 該当ユーザの所属しているグループのイベントを全て取得する
func (e *GetEventOne) GetEvent(uid string, id int) ([]byte, error) {
	db := model.ConnectDB()
	sqlDb, err := db.DB() //コネクションクローズ用
	if err != nil {
		return nil, errors.New("DBとの接続に失敗しました。")
	}
	defer sqlDb.Close()

	err = db.Table("events").
		Joins("INNER JOIN users AS create_user ON events.create_user = create_user.uid").
		Joins("LEFT JOIN users AS update_user ON events.update_user = update_user.uid").
		Select("events.*, create_user.name AS create_user_name, update_user.name AS update_user_name").
		Where("create_user.group_id = events.group_id AND create_user.uid = ? AND events.id = ?", uid, id).Find(&e).Error
	if err != nil {
		logging.WriteErrorLog(err.Error(), true)
		return nil, err
	}

	// 送信用に変換

	if e.UpdateUserName == "" {
		// 更新ユーザ名が設定されてない場合、作成ユーザ名を適用する
		e.UpdateUserName = e.CreateUserName
	}
	eventMap := make(map[string]interface{})
	eventMap["id"] = e.ID
	eventMap["category"] = e.Category
	eventMap["amount"] = e.Amount
	eventMap["storeName"] = e.StoreName
	eventMap["date"] = e.Date.Format("2006-01-02")
	eventMap["createUser"] = e.CreateUserName
	eventMap["updateUser"] = e.UpdateUserName
	eventMap["createdAt"] = e.CreatedAt.Format("2006-01-02 15:04:05")
	eventMap["updatedAt"] = e.UpdatedAt.Format("2006-01-02 15:04:05")

	// JSONへと変換
	jsonEvents, err := json.Marshal(eventMap)
	if err != nil {
		logging.WriteErrorLog(err.Error(), true)
		return nil, err
	}
	return jsonEvents, nil
}
