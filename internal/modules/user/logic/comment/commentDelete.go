package comment

import (
	"gin-blog/internal/modules/user/models"
	"gin-blog/internal/modules/user/repository/db"
	"gin-blog/internal/modules/user/requests/comment"
	"gin-blog/pkg/globals"
)

type DeleteLogic struct {
	err    error
	appCtx *globals.AppCtx
	cdr    comment.DeleteCommentRequest
}

func NewDeleteLogic(appCtx *globals.AppCtx, cdr comment.DeleteCommentRequest) *DeleteLogic {
	return &DeleteLogic{
		appCtx: appCtx,
		cdr:    cdr,
	}
}

func (d *DeleteLogic) Delete() error {
	d.DeleteInDb()
	return d.err
}

func (d *DeleteLogic) DeleteInDb() *DeleteLogic {
	if d.err != nil {
		return d
	}
	var comment models.Comment
	comment.ID = d.cdr.CommentID
	d.err = db.DeleteComment(db.CommentSqlParam{
		Db:      d.appCtx.GetDb(),
		Comment: comment,
	})
	return d
}
