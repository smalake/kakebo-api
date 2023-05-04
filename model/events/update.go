package events

import (
	"errors"
	"time"

	"github.com/smalake/kakebo-api/model"
	"github.com/smalake/kakebo-api/utils/logging"
)

type UpdateEvent struct {
	ID         int       `json:"id"`
	Category   int       `json:"category"`
	Amount     int       `json:"amount"`
	Date       time.Time `json:"date" time_format:"2006-01-02"`
	StoreName  string    `json:"storeName" gorm:"column:store_name"`
	UpdateUser string
	UpdatedAt  time.Time
}

// イベントを更新
func (e *UpdateEvent) UpdateEvent(uid string) error {
	db := model.ConnectDB()
	sqlDb, err := db.DB() //コネクションクローズ用
	if err != nil {
		return errors.New("DBとの接続に失敗しました。")
	}
	defer sqlDb.Close()

	err = db.Table("events").Where("id = ?", e.ID).Updates(UpdateEvent{
		Category:   e.Category,
		Amount:     e.Amount,
		Date:       e.Date,
		StoreName:  e.StoreName,
		UpdateUser: uid,
		UpdatedAt:  time.Now(),
	}).Error
	if err != nil {
		logging.WriteErrorLog(err.Error(), true)
		return err
	}
	return nil
}
