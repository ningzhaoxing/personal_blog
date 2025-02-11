package models

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Title          string         `json:"title" gorm:"type:varchar(50);"`
	Content        string         `json:"content" gorm:"type:text"`
	Likes          int            `json:"likes" gorm:"default:0;column:likes"`
	Readings       int            `json:"readings" gorm:"default:0;column:readings"`
	CommentNum     int            `json:"comment_num" gorm:"-"`
	IsLiked        bool           `json:"is_liked" gorm:"-"`
	PublishAt      string         `json:"publish_at" gorm:"-"`
	Comments       []Comment      `json:"comments"`
	ArticleImages  []ArticleImage `json:"image"`
	UserID         uint           `json:"user_id" gorm:"not null"`
	User           User           `json:"user"`
	ArticleClassID uint           `json:"class_id" gorm:"default:4"`
	ArticleClass   ArticleClass   `json:"class"`
}
