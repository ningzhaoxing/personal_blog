package comment

import (
	"gin-blog/internal/modules/user/models"
	db2 "gin-blog/internal/modules/user/repository/db"
	"gin-blog/internal/modules/user/requests/comment"
	"gin-blog/pkg/globals"
	"gin-blog/utils/time"
)

type GetCommentListLogic struct {
	err      error
	appCtx   *globals.AppCtx
	clr      comment.CommentListRequest
	comments []models.Comment
}

func NewGetCommentListLogic(appCtx *globals.AppCtx, clr comment.CommentListRequest) *GetCommentListLogic {
	return &GetCommentListLogic{
		appCtx: appCtx,
		clr:    clr,
	}
}

func (g *GetCommentListLogic) CommentList() ([]models.Comment, error) {
	g.GetCommentList().TimeTrans()
	return g.comments, g.err
}

func (g *GetCommentListLogic) GetCommentList() *GetCommentListLogic {
	if g.err != nil {
		return g
	}

	g.comments, g.err = db2.QueryCommentList(db2.CommentSqlParam{
		Db:      g.appCtx.GetDb(),
		Comment: models.Comment{ArticleID: g.clr.ArticleID},
	})
	return g
}

func (g *GetCommentListLogic) TimeTrans() *GetCommentListLogic {
	for i := range g.comments {
		g.comments[i].PublishAt, g.err = time.TimeAgo(g.comments[i].CreatedAt)
		if g.err != nil {
			return g
		}
	}
	return g
}
