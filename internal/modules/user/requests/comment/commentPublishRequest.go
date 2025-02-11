package comment

type PublishCommentRequest struct {
	ArticleID uint   `form:"article_id" binding:"required"`
	Content   string `form:"content" binding:"required"`
	UserID    uint
}
