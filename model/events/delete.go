package events

import (
	"errors"

	"github.com/smalake/kakebo-api/model"
	"github.com/smalake/kakebo-api/utils/logging"
)

type DeleteEvent struct {
	ID int `json:"id"`
}

// イベントを削除
func (e *DeleteEvent) DeleteOneEvent() error {
	db := model.ConnectDB()
	sqlDb, err := db.DB() //コネクションクローズ用
	if err != nil {
		return errors.New("DBとの接続に失敗しました。")
	}
	defer sqlDb.Close()

	err = db.Delete(&e).Error
	if err != nil {
		logging.WriteErrorLog(err.Error(), true)
		return err
	}

	return nil
}
