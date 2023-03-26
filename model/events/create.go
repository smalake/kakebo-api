package events

import (
	"errors"
	"time"

	"github.com/smalake/kakebo-api/model"
	"github.com/smalake/kakebo-api/model/users"
	"gorm.io/gorm"
)

// イベントを新規登録するための構造体
type InsertEvent struct {
	Category   int       `json:"category"`
	Amount     int       `json:"amount"`
	Date       time.Time `json:"date" time_format:"2006-01-02"`
	StoreName  string    `json:"store_name" gorm:"column:store_name"`
	CreateUser string    `json:"create_user"`
	GroupID    int
	CreatedAt  time.Time
}

// フロント側から直接受け取るための構造体
type CreateEvent struct {
	ID         int       `json:"id"`
	UID        string    `json:"uid"`
	Category1  int       `json:"category1"`
	Category2  int       `json:"category2"`
	Amount1    int       `json:"amount1"`
	Amount2    int       `json:"amount2"`
	Date       time.Time `json:"date" time_format:"2006-01-02"`
	StoreName  string    `json:"store_name"`
	CreateUser string    `json:"create_user"`
	GroupID    int
}

// 登録するユーザが所属しているグループのIDを取得
func (e *CreateEvent) getGroupID(db *gorm.DB) error {
	var user users.User
	err := db.Where("uid = ?", e.CreateUser).First(&user).Error
	if err != nil {
		return err
	}
	e.GroupID = user.GroupID
	return nil
}

// イベントを新規作成
func (e *CreateEvent) InsertEvent(uid string) error {
	e.UID = uid
	db := model.ConnectDB()
	sqlDb, err := db.DB() //コネクションクローズ用
	if err != nil {
		return errors.New("DBとの接続に失敗しました。")
	}
	defer sqlDb.Close()

	// グループIDの取得
	err = e.getGroupID(db)
	if err != nil {
		return err
	}

	// 支出金額の合計
	amount := e.Amount1

	// 2つ目の支出が存在する場合、別途登録する
	if e.Amount2 != 0 {
		// 2つ目の金額分合計金額から引く
		amount = e.Amount1 - e.Amount2
		sub := &InsertEvent{
			Category:   e.Category2,
			Amount:     e.Amount2,
			Date:       e.Date,
			StoreName:  e.StoreName,
			CreateUser: e.CreateUser,
			GroupID:    e.GroupID,
		}
		err = db.Table("events").Create(sub).Error
		if err != nil {
			return err
		}
	}

	// DB登録用に切り出して登録
	main := &InsertEvent{
		Category:   e.Category1,
		Amount:     amount,
		Date:       e.Date,
		StoreName:  e.StoreName,
		CreateUser: e.CreateUser,
		GroupID:    e.GroupID,
	}
	err = db.Table("events").Create(main).Error
	if err != nil {
		return err
	}

	return nil
}
