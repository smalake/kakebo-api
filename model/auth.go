package model

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	"github.com/smalake/kakebo-api/utils/logging"
)

// ユーザを新規登録する
func (u *User) RegisterUser() error {
	db := ConnectDB()
	sqlDb, err := db.DB() //コネクションクローズ用
	if err != nil {
		return errors.New("DBとの接続に失敗しました。")
	}
	defer sqlDb.Close()

	// メールアドレスが使われていないかチェック
	db.Where("email = ?", u.Email).First(&u)
	if u.ID != 0 {
		err := errors.New("すでに使用されているメールアドレスです。")
		return err
	}

	// パスワードをハッシュ化
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

	// DBへと登録
	err = db.Table("users").Create(u).Error
	if err != nil {
		return err
	}
	return nil
}

// ログイン用にパスワードを取得する
func (u *User) GetPassword() (string, error) {
	db := ConnectDB()
	sqlDb, err := db.DB() //コネクションクローズ用
	if err != nil {
		logging.WriteErrorLog(err.Error(), true)
	}
	defer sqlDb.Close()

	var password string
	err = db.Table("users").Select("password").Where("email = ?", u.Email).Scan(&password).Error
	if err != nil {
		logging.WriteErrorLog(err.Error(), true)
		return "login error", err
	}

	return password, nil
}

// ミドルウェアによる認証チェック時にユーザIDを取得する
func GetUserId(email string) (int, error) {
	db := ConnectDB()
	sqlDb, err := db.DB() //コネクションクローズ用
	if err != nil {
		logging.WriteErrorLog(err.Error(), true)
	}
	defer sqlDb.Close()

	var id int
	err = db.Table("users").Select("id").Where("email = ?", email).Scan(&id).Error
	if err != nil {
		logging.WriteErrorLog(err.Error(), true)
		return 0, err
	}

	return id, nil
}
