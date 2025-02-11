package user

import (
	"context"
	"gin-blog/internal/modules/user/repository/redis"
	"gin-blog/internal/modules/user/requests/user"
	"gin-blog/pkg/globals"
	"strconv"
)

type LikeArticleLogic struct {
	err    error
	lar    user.LikeArticleRequest
	appCtx *globals.AppCtx
}

func NewLikeArticleLogic(lar user.LikeArticleRequest, appCtx *globals.AppCtx) *LikeArticleLogic {
	return &LikeArticleLogic{
		lar:    lar,
		appCtx: appCtx,
	}
}

func (l *LikeArticleLogic) Do() error {
	exist, err := l.UidExistInArticleSetOrNot()
	if err != nil {
		return err
	}

	if exist {
		l.CancelLike()
	} else {
		l.Like()
	}

	return l.err
}

func (l *LikeArticleLogic) Like() *LikeArticleLogic {
	if l.err != nil {
		return l
	}

	l.AddUidToArticleLikeSet().PlusOneLikeNum()

	return l
}

func (l *LikeArticleLogic) CancelLike() *LikeArticleLogic {
	if l.err != nil {
		return l
	}

	l.RemoveUidToArticleLikeSet().MinusOneLikeNum()

	return l
}

// AddUidToArticleLikeSet 将用户添加到文章点赞集合
func (l *LikeArticleLogic) AddUidToArticleLikeSet() *LikeArticleLogic {
	if l.err != nil {
		return l
	}

	l.err = redis.AddUserToLikeSet(redis.ArticleLikeParam{
		ArticleID: strconv.Itoa(l.lar.ArticleID),
		UserID:    strconv.Itoa(l.lar.UserID),
		Rdb:       l.appCtx.GetRdb(),
		Context:   context.Background(),
	})
	return l
}

// RemoveUidToArticleLikeSet 将用户移除文章点赞集合
func (l *LikeArticleLogic) RemoveUidToArticleLikeSet() *LikeArticleLogic {
	if l.err != nil {
		return l
	}

	l.err = redis.RemoveUserFromArticleLikeSet(redis.ArticleLikeParam{
		ArticleID: strconv.Itoa(l.lar.ArticleID),
		UserID:    strconv.Itoa(l.lar.UserID),
		Rdb:       l.appCtx.GetRdb(),
		Context:   context.Background(),
	})
	return l
}

// PlusOneLikeNum 点赞数量+1
func (l *LikeArticleLogic) PlusOneLikeNum() *LikeArticleLogic {
	if l.err != nil {
		return l
	}

	l.err = redis.IncreaseArticleLikeNum(redis.ArticleLikeParam{
		ArticleID: strconv.Itoa(l.lar.ArticleID),
		Rdb:       l.appCtx.GetRdb(),
		Context:   context.Background(),
	}, 1)
	return l
}

// MinusOneLikeNum 点赞数量-1
func (l *LikeArticleLogic) MinusOneLikeNum() *LikeArticleLogic {
	if l.err != nil {
		return l
	}

	l.err = redis.IncreaseArticleLikeNum(redis.ArticleLikeParam{
		ArticleID: strconv.Itoa(l.lar.ArticleID),
		Rdb:       l.appCtx.GetRdb(),
		Context:   context.Background(),
	}, -1)
	return l
}

// UidExistInArticleSetOrNot 检查当前用户是否存在于文章点赞集合
func (l *LikeArticleLogic) UidExistInArticleSetOrNot() (bool, error) {
	if l.err != nil {
		return false, l.err
	}

	exist, err := redis.IsUserInArticleLikeSet(redis.ArticleLikeParam{
		ArticleID: strconv.Itoa(l.lar.ArticleID),
		UserID:    strconv.Itoa(l.lar.UserID),
		Rdb:       l.appCtx.GetRdb(),
		Context:   context.Background(),
	})
	if err != nil {
		return false, err
	}

	return exist, nil
}
