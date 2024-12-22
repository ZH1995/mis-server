package model

type User struct {
	ID         uint64 `json:"id" gorm:"primarykey"`
	Username   string `json:"username" gorm:"uniqueIndex;not null"`
	Password   string `json:"password" gorm:"not null"`
	CreateTime int64  `json:"create_time" gorm:"default:0"`
	UpdateTime int64  `json:"update_time" gorm:"default:0"`
	IsDeleted  int8   `json:"is_deleted" gorm:"default:0"`
}
