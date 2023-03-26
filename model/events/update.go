package events

import (
	"errors"
	"time"

	"github.com/smalake/kakebo-api/model"
	"github.com/smalake/kakebo-api/utils/logging"
)

type UpdateEvent struct {
	Category   int       `json:"category"`
	Amount     int       `json:"amount"`
	Date       time.Time `json:"date" time_format:"2006-01-02"`
	StoreName  string    `json:"store_name" gorm:"column:store_name"`
	UpdateUser string    `json:"update_user"`
	CreatedAt  time.Time
}

// イベントを更新
func (e *UpdateEvent) UpdateEvent() error {
	db := model.ConnectDB()
	sqlDb, err := db.DB() //コネクションクローズ用
	if err != nil {
		return errors.New("DBとの接続に失敗しました。")
	}
	defer sqlDb.Close()

	err = db.Model(&e).Updates(UpdateEvent{
		Category:   e.Category,
		Amount:     e.Amount,
		Date:       e.Date,
		StoreName:  e.StoreName,
		UpdateUser: e.UpdateUser,
	}).Error
	if err != nil {
		logging.WriteErrorLog(err.Error(), true)
		return err
	}
	return nil
}
