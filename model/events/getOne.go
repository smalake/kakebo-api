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
	CreateUser     string    `json:"createUser"`
	UpdateUser     string    `json:"updateUser"`
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

	// JSONへと変換
	jsonEvents, err := json.Marshal(e)
	if err != nil {
		logging.WriteErrorLog(err.Error(), true)
		return nil, err
	}
	return jsonEvents, nil
}
