package article

import "gin-blog/internal/modules/user/models"

type PublishArticleRequest struct {
	Title     string   `json:"title"  binding:"required,min=1,max=20"`
	Content   string   `json:"content"  binding:"required,min=5"`
	Images    []string `json:"images"`
	ClassID   uint     `json:"class_id"  binding:"required"`
	ArticleID int
	User      models.User
}
