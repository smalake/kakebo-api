package setup

import (
	"errors"
	"fmt"
	"time"

	"github.com/smalake/kakebo-api/model"
	"github.com/smalake/kakebo-api/model/users"
	"github.com/smalake/kakebo-api/utils/logging"
)

func (g *Group) CreateGroup(uid string) error {
	fmt.Println("CreateGroup実行")
	db := model.ConnectDB()
	tx := db.Begin()      // トランザクション開始
	sqlDb, err := db.DB() //コネクションクローズ用
	if err != nil {
		logging.WriteErrorLog(err.Error(), true)
		return errors.New("DBとの接続に失敗しました。")
	}
	defer sqlDb.Close()

	g.UID = uid

	// グループの新規作成
	err = tx.Debug().Create(&g).Error
	if err != nil {
		tx.Rollback()
		logging.WriteErrorLog(err.Error(), true)
		return err
	}

	// ユーザにグループIDを反映

	err = tx.Where("uid = ?", uid).Updates(users.User{
		GroupID:   g.ID,
		UpdatedAt: time.Now(),
	}).Error
	if err != nil {
		tx.Rollback()
		logging.WriteErrorLog(err.Error(), true)
		return err
	}
	tx.Commit()
	return nil
}
