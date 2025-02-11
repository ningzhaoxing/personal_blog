package db

import (
	"gin-blog/internal/modules/user/models"
	"gorm.io/gorm"
)

type CommentSqlParam struct {
	Db      *gorm.DB
	Comment models.Comment
}

func QueryCommentList(c CommentSqlParam) ([]models.Comment, error) {
	var comments []models.Comment
	if err := c.Db.
		Preload("User").
		Where("article_id = ?", c.Comment.ArticleID).Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}

func QueryArticleCommentNum(c CommentSqlParam, aid uint) (int, error) {
	var num int64
	if err := c.Db.Model(&models.Comment{}).Where("article_id = ?", aid).Count(&num).Error; err != nil {
		return -1, err
	}
	return int(num), nil
}

func InsertComment(c CommentSqlParam) error {
	if err := c.Db.Create(&c.Comment).Error; err != nil {
		return err
	}
	return nil
}

func DeleteComment(c CommentSqlParam) error {
	if err := c.Db.Delete(&models.Comment{}, c.Comment.ID).Error; err != nil {
		return err
	}
	return nil
}
