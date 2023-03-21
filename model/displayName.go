package model

import (
	"errors"
	"time"

	"github.com/smalake/kakebo-api/utils/logging"
)

type DisplayName struct {
	UID         string `json:"uid"`
	DisplayName string `json:"name"`
	UpdatedAt   time.Time
}

// UIDから表示名を取得する
func (n *DisplayName) GetDisplayName() error {
	db := ConnectDB()
	sqlDb, err := db.DB() //コネクションクローズ用
	if err != nil {
		return errors.New("DBとの接続に失敗しました。")
	}
	defer sqlDb.Close()

	err = db.Table("users").Where("uid = ?", n.UID).First(&n).Error
	if err != nil {
		logging.WriteErrorLog(err.Error(), true)
		return err
	}

	return nil
}

// 表示名を更新する
func (n *DisplayName) UpdateDisplayName() error {
	db := ConnectDB()
	sqlDb, err := db.DB() //コネクションクローズ用
	if err != nil {
		return errors.New("DBとの接続に失敗しました。")
	}
	defer sqlDb.Close()

	err = db.Table("users").Where("uid = ?", n.UID).Update("name", n.DisplayName).Error
	if err != nil {
		logging.WriteErrorLog(err.Error(), true)
		return err
	}
	return nil
}
