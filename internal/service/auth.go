package service

import (
	"MyGin/internal/model"
	"MyGin/internal/repository/mysql"
	"MyGin/pkg/util"
	"time"
)

func Register(user *model.User) error {
	hashedPwd, err := util.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPwd
	user.CreateTime = time.Now().Unix()
	return mysql.CreateUser(user)
}

func Login(username, password string) (string, error) {
	user, err := mysql.GetUserByUsername(username)
	if err != nil {
		return "", err
	}

	if !util.CheckPassword(password, user.Password) {
		return "", err
	}

	return util.GenerateJWT(username)
}
