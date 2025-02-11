package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Content   string `json:"content"`
	PublishAt string `gorm:"-"`
	UserID    uint   `json:"user_id"`
	User      User
	ArticleID uint `json:"article_id"`
	Article   Article
}
