package service

import (
	"MyGin/internal/repository/redis"
)

func LikeArticle(articleID string) error {
	likeKey := "article:" + articleID + ":like"
	return redis.IncrArticleLike(likeKey)
}

func GetArticleLikes(articleID string) (string, error) {
	likeKey := "article:" + articleID + ":like"
	return redis.GetArticleLikes(likeKey)
}
