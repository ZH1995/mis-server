package service

import (
	"MyGin/internal/model"
	"MyGin/internal/repository/mysql"
)

func CreateArticle(article *model.Article) error {
	return mysql.CreateArticle(article)
}

func GetArticles() ([]model.Article, error) {
	return mysql.GetArticles()
}

func GetArticleByID(ID string) (*model.Article, error) {
	return mysql.GetArticleByID(ID)
}
