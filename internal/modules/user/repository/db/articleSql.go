package db

import (
	"fmt"
	"gin-blog/internal/modules/user/models"
	"gin-blog/internal/modules/user/requests/article"
	"gorm.io/gorm"
)

type ArticleSqlParam struct {
	Db      *gorm.DB
	Alr     article.ArticleListRequest
	Article models.Article
}

// QueryArticleListByPage 模糊分页查询文章列表
func QueryArticleListByPage(a ArticleSqlParam) ([]models.Article, error) {
	articles := make([]models.Article, 0)
	if err := a.Db.
		InnerJoins("User").
		Where("User.is_owner = ?", true).
		Preload("ArticleClass").
		Preload("ArticleImages").
		Order("articles.created_at desc").Limit(a.Alr.Limit).
		Offset((a.Alr.Page-1)*a.Alr.Limit).
		Where("articles.title LIKE ?", fmt.Sprintf("%%%s%%", a.Alr.KeyWords)).
		Find(&articles).Error; err != nil {
		return nil, err
	}
	return articles, nil
}

func QueryArticle(a ArticleSqlParam) (models.Article, error) {
	if err := a.Db.
		Preload("User").
		Preload("ArticleImages").
		Preload("ArticleClass").
		Where("id = ?", a.Article.ID).
		First(&a.Article).Error; err != nil {
		return models.Article{}, err
	}
	return a.Article, nil
}

func QueryArticlesByClass(a ArticleSqlParam) ([]models.Article, error) {
	articles := make([]models.Article, 0)
	if err := a.Db.Model(&models.Article{}).Where("article_class_id = ?", a.Article.ArticleClassID).Find(&articles).Error; err != nil {
		return nil, err
	}
	return articles, nil
}

func InsertArticle(a ArticleSqlParam) (int, error) {
	if err := a.Db.Create(&a.Article).Error; err != nil {
		return -1, err
	}
	return int(a.Article.ID), nil
}

func UpdateArticleReadings(a ArticleSqlParam) error {
	if err := a.Db.Model(&models.Article{}).
		Where("id = ?", a.Article.ID).
		Update("readings", a.Article.Readings).Error; err != nil {
		return err
	}
	return nil
}

func UpdateArticleLikes(a ArticleSqlParam) error {
	if err := a.Db.Model(&models.Article{}).
		Where("id = ?", a.Article.ID).
		Update("likes", a.Article.Likes).Error; err != nil {
		return err
	}
	return nil
}

func DeleteArticle(a ArticleSqlParam) error {
	if err := a.Db.Delete(&models.Article{}, a.Article.ID).Error; err != nil {
		return err
	}
	return nil
}
