package user

type LikeArticleRequest struct {
	ArticleID int `form:"article_id" json:"article_id" binding:"required"`
	UserID    int
}
