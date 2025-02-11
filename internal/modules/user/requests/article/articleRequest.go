package article

import "gin-blog/internal/modules/user/models"

type ArticleRequest struct {
	ID      uint `form:"ID" binding:"required"`
	UserID  uint `form:"uid"`
	Article models.Article
}
