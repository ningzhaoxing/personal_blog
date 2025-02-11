package comment

import (
	"gin-blog/internal/modules/user/models"
	db2 "gin-blog/internal/modules/user/repository/db"
	"gin-blog/internal/modules/user/requests/comment"
	"gin-blog/pkg/globals"
)

type PublishLogic struct {
	err    error
	appCtx *globals.AppCtx
	cpr    comment.PublishCommentRequest
}

func NewPublishLogic(appCtx *globals.AppCtx, cpr comment.PublishCommentRequest) *PublishLogic {
	return &PublishLogic{
		appCtx: appCtx,
		cpr:    cpr,
	}
}

func (p *PublishLogic) Publish() error {
	p.SaveCommentInDb()
	return p.err
}

func (p *PublishLogic) SaveCommentInDb() *PublishLogic {
	if p.err != nil {
		return p
	}

	p.err = db2.InsertComment(db2.CommentSqlParam{
		Db: p.appCtx.GetDb(),
		Comment: models.Comment{
			Content:   p.cpr.Content,
			UserID:    p.cpr.UserID,
			ArticleID: p.cpr.ArticleID,
		},
	})
	return p
}
