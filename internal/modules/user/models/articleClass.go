package models

import "gorm.io/gorm"

type ArticleClass struct {
	gorm.Model
	Name     string `json:"name"`
	Articles []Article
}
