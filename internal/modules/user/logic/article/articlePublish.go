package article

import (
	"context"
	"gin-blog/internal/modules/user/models"
	"gin-blog/internal/modules/user/repository/db"
	redis2 "gin-blog/internal/modules/user/repository/redis"
	"gin-blog/internal/modules/user/requests/article"
	"gin-blog/pkg/globals"
	"gorm.io/gorm"
	"strconv"
)

type PublishLogic struct {
	err    error
	appCtx *globals.AppCtx
	par    article.PublishArticleRequest
}

func NewPublishLogic(appCtx *globals.AppCtx, par article.PublishArticleRequest) *PublishLogic {
	return &PublishLogic{
		appCtx: appCtx,
		par:    par,
	}
}

func (p *PublishLogic) Publish() (int, error) {
	tx := p.appCtx.GetDb().Begin()
	p.SaveArticleInDb(tx)
	if len(p.par.Images) != 0 {
		p.SaveImagesInDb(tx)
	}
	tx.Commit()
	p.SetArticleLikes().SetArticleReadings()
	return p.par.ArticleID, p.err
}

func (p *PublishLogic) SaveArticleInDb(tx *gorm.DB) *PublishLogic {
	if p.err != nil {
		tx.Rollback()
		return p
	}

	p.par.ArticleID, p.err = db.InsertArticle(db.ArticleSqlParam{
		Db: tx,
		Article: models.Article{
			ArticleClassID: p.par.ClassID,
			Title:          p.par.Title,
			Content:        p.par.Content,
			UserID:         p.par.User.ID,
		},
	})
	return p
}

func (p *PublishLogic) SaveImagesInDb(tx *gorm.DB) *PublishLogic {
	if p.err != nil {
		tx.Rollback()
		return p
	}

	for _, image := range p.par.Images {
		err := db.InsertArticleImage(db.ImageSqlParam{
			Db:    tx,
			Image: models.ArticleImage{Path: image, ArticleID: uint(p.par.ArticleID)},
		})
		if err != nil {
			tx.Rollback()
			p.err = err
			return p
		}
	}
	return p
}

// SetArticleLikes 设置文章点赞数
func (p *PublishLogic) SetArticleLikes() *PublishLogic {
	if p.err != nil {
		return p
	}

	p.err = redis2.SetArticleLikes(redis2.ArticleLikeParam{
		ArticleID: strconv.Itoa(p.par.ArticleID),
		Rdb:       p.appCtx.GetRdb(),
		Context:   context.Background(),
	}, 0)
	return p
}

// SetArticleReadings 设置文章阅读量
func (p *PublishLogic) SetArticleReadings() *PublishLogic {
	if p.err != nil {
		return p
	}

	p.err = redis2.SetArticleReadings(redis2.ArticleReadingsParam{
		ArticleID: strconv.Itoa(p.par.ArticleID),
		Rdb:       p.appCtx.GetRdb(),
		Context:   context.Background(),
	}, 0)
	return p
}
