package comment

type DeleteCommentRequest struct {
	CommentID uint `json:"comment_id" binding:"required"`
}
