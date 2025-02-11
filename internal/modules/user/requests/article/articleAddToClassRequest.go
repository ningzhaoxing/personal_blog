package article

type ArticleAddToClassRequest struct {
	ArticleID uint `form:"article_id" binding:"required"`
	ClassID   uint `form:"class_id" binding:"required"`
}
