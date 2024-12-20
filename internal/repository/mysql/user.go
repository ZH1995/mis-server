package mysql

import (
	"MyGin/global"
	"MyGin/internal/model"
)

func CreateUser(user *model.User) error {
	return global.Db.Create(user).Error
}

func GetUserByUsername(username string) (*model.User, error) {
	var user model.User
	err := global.Db.Where("username = ?", username).First(&user).Error
	return &user, err
}
