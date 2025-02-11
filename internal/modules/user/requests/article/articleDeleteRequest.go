package article

type DeleteArticleRequest struct {
	ArticleID uint `json:"article_id" binding:"required"`
}
