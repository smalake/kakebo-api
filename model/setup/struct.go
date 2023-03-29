package setup

type Group struct {
	ID  int    `json:"group_id" gorm:"default:0"`
	UID string `json:"uid" gorm:"column:manage_uid"`
}
