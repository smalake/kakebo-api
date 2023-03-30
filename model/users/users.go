package users

import (
	"errors"
	"time"

	"github.com/smalake/kakebo-api/model"
)

type User struct {
	ID        int    `json:"id"`
	UID       string `json:"uid"`
	Name      string `json:"name"`
	Type      int    `json:"type"`
	GroupID   int    `json:"group_id"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// ユーザを新規登録する
func (u *User) CreateUser() error {
	db := model.ConnectDB()
	sqlDb, err := db.DB() //コネクションクローズ用
	if err != nil {
		return errors.New("DBとの接続に失敗しました。")
	}
	defer sqlDb.Close()

	// UIDがすでに登録されていないかチェック
	db.Where("uid = ?", u.UID).First(&u)
	if u.ID != 0 {
		err := errors.New("すでに登録されています。")
		return err
	}

	// DBへと登録
	err = db.Table("users").Create(u).Error
	if err != nil {
		return err
	}
	return nil
}

// ログイン
func (u *User) LoginUser() error {
	db := model.ConnectDB()
	sqlDb, err := db.DB() //コネクションクローズ用
	if err != nil {
		return errors.New("DBとの接続に失敗しました。")
	}
	defer sqlDb.Close()
	// UIDから登録されているユーザか判別
	db.Where("uid = ?", u.UID).First(&u)
	if u.ID == 0 {
		err := errors.New("登録されていないユーザーです。")
		return err
	}
	return nil
}
