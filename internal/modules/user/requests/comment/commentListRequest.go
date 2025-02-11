package comment

type CommentListRequest struct {
	ArticleID uint `form:"article_id" binding:"required"`
}
