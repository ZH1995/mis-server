package config

import (
	"MyGin/global"
	"log"

	"github.com/go-redis/redis"
)

func initRedis() {
	RedisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	_, err := RedisClient.Ping().Result()
	if err != nil {
		log.Fatal("error connecting to redis: ", err)
	}

	global.RedisDB = RedisClient
}
