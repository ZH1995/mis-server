package mysql

import (
	"MyGin/global"
	"MyGin/internal/model"
)

/*
CREATE TABLE `mygin_db`.`users`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  `create_time` bigint UNSIGNED NOT NULL,
  `update_time` bigint NOT NULL DEFAULT 0,
  `is_deleted` int NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `name_idx`(`username` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;
*/

func CreateUser(user *model.User) error {
	return global.Db.Create(user).Error
}

func GetUserByUsername(username string) (*model.User, error) {
	var user model.User
	err := global.Db.Where("username = ?", username).First(&user).Error
	return &user, err
}
