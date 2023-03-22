package model

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/smalake/kakebo-api/utils/logging"
)

type DisplayName struct {
	Name      string `json:"name"`
	UpdatedAt time.Time
}

// UIDから表示名を取得する
func (n *DisplayName) GetDisplayName(uid string) ([]byte, error) {
	db := ConnectDB()
	sqlDb, err := db.DB() //コネクションクローズ用
	if err != nil {
		return nil, errors.New("DBとの接続に失敗しました。")
	}
	defer sqlDb.Close()

	err = db.Table("users").Select("name").Where("uid = ?", uid).First(&n).Error
	if err != nil {
		logging.WriteErrorLog(err.Error(), true)
		return nil, err
	}

	// JSON形式にしてからJSON変換
	jsonName, err := json.Marshal(map[string]string{"name": n.Name})
	if err != nil {
		logging.WriteErrorLog(err.Error(), true)
		return nil, err
	}
	return jsonName, nil
}

// 表示名を更新する
func (n *DisplayName) UpdateDisplayName(uid string) error {
	db := ConnectDB()
	sqlDb, err := db.DB() //コネクションクローズ用
	if err != nil {
		return errors.New("DBとの接続に失敗しました。")
	}
	defer sqlDb.Close()

	err = db.Table("users").Where("uid = ?", uid).Update("name", n.Name).Error
	if err != nil {
		logging.WriteErrorLog(err.Error(), true)
		return err
	}
	return nil
}
