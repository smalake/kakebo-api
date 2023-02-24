package model

import "github.com/smalake/kakebo-api/utils/logging"

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

// ユーザを新規登録する
func (u *User) RegisterUser() error {
	db := ConnectDB()
	sqlDb, err := db.DB() //コネクションクローズ用
	if err != nil {
		logging.WriteErrorLog(err.Error(), true)
		return err
	}
	defer sqlDb.Close()

	err = db.Table("users").Create(u).Error
	if err != nil {
		logging.WriteErrorLog(err.Error(), true)
		return err
	}
	return nil
}
