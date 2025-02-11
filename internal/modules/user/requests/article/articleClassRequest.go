package article

type ArticleClassRequest struct {
	ArticleClassID uint `json:"article_class_id" form:"article_class_id" binding:"required"`
}
