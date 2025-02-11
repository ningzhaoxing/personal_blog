package models

import "gorm.io/gorm"

type ArticleImage struct {
	gorm.Model
	ArticleID uint `json:"article_id"`
	Article   Article
	Path      string `gorm:"not null"`
}
