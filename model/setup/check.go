package setup

import (
	"encoding/json"
	"errors"

	"github.com/smalake/kakebo-api/model"
	"github.com/smalake/kakebo-api/utils/logging"
)

type Group struct {
	ID      int    `json:"id"`
	UID     string `json:"uid"`
	GroupID int    `json:"group_id" gorm:"default:0"`
}

func (g *Group) GetGroupID(uid string) ([]byte, error) {
	db := model.ConnectDB()
	sqlDb, err := db.DB() //コネクションクローズ用
	if err != nil {
		logging.WriteErrorLog(err.Error(), true)
		return nil, errors.New("DBとの接続に失敗しました。")
	}
	defer sqlDb.Close()

	err = db.Table("users").Where("uid = ?", uid).First(&g).Error
	if err != nil {
		logging.WriteErrorLog(err.Error(), true)
		return nil, err
	}

	var resData = map[string]int{"groupId": g.GroupID}
	jsonData, err := json.Marshal(resData)
	if err != nil {
		logging.WriteErrorLog(err.Error(), true)
		return nil, err
	}

	return jsonData, nil
}
