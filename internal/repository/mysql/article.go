package mysql

import (
	"MyGin/global"
	"MyGin/internal/model"
)

func CreateArticle(article *model.Article) error {
	return global.Db.Create(article).Error
}

func GetArticleByID(ID string) (*model.Article, error) {
	var article model.Article
	err := global.Db.Where("id = ?", ID).First(&article).Error
	return &article, err
}

func GetArticles() ([]model.Article, error) {
	var articles []model.Article
	err := global.Db.Find(&articles).Error
	return articles, err
}
