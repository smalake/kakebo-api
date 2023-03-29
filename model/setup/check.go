package setup

import (
	"encoding/json"
	"errors"

	"github.com/smalake/kakebo-api/model"
	"github.com/smalake/kakebo-api/model/users"
	"github.com/smalake/kakebo-api/utils/logging"
	"gorm.io/gorm"
)

func GetGroupID(uid string) ([]byte, error) {
	db := model.ConnectDB()
	sqlDb, err := db.DB() //コネクションクローズ用
	if err != nil {
		logging.WriteErrorLog(err.Error(), true)
		return nil, errors.New("DBとの接続に失敗しました。")
	}
	defer sqlDb.Close()

	var user users.User

	err = db.Where("uid = ?", uid).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// レコードが存在しない場合はグループIDを-1にする
			var resData = map[string]int{"groupId": -1}
			jsonData, err := json.Marshal(resData)
			if err != nil {
				logging.WriteErrorLog(err.Error(), true)
				return nil, err
			}
			return jsonData, nil
		} else {
			logging.WriteErrorLog(err.Error(), true)
			return nil, err
		}
	}

	var resData = map[string]int{"groupId": user.GroupID}
	jsonData, err := json.Marshal(resData)
	if err != nil {
		logging.WriteErrorLog(err.Error(), true)
		return nil, err
	}

	return jsonData, nil
}
