package model

type Article struct {
	ID         uint64 `json:"id" gorm:"primarykey"`
	Title      string `json:"title" binding:"required"`
	Content    string `json:"content" binding:"required"`
	Preview    string `json:"preview" binding:"required"`
	Likes      int    `json:"likes" gorm:"default:0"`
	CreateTime int64  `json:"create_time" gorm:"default:0"`
	UpdateTime int64  `json:"update_time" gorm:"default:0"`
	IsDeleted  int8   `json:"is_deleted" gorm:"default:0"`
}
