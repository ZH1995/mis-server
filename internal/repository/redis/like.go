package redis

import (
	"MyGin/global"
)

func IncrArticleLike(likeKey string) error {
	err := global.RedisDB.Incr(likeKey).Err()
	return err
}

func GetArticleLikes(likeKey string) (string, error) {
	likes, err := global.RedisDB.Get(likeKey).Result()
	return likes, err
}
