package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"strconv"
)

/*
点赞模块：
	set存文章的点赞用户列表 (key:"article_like+id"-value:userid)
	hash存文章点赞数量 (key:"article_like" filed:"articleid" value:"likeNum")
*/

type ArticleLikeParam struct {
	ArticleID string
	UserID    string
	Rdb       *redis.Client
	Context   context.Context
}

func articleLikeHashKey() string {
	return "article_like"
}

func articleLikeSetKey(aid string) string {
	return "article_like" + aid
}

// AddUserToLikeSet 将用户id添加到文章点赞集合
func AddUserToLikeSet(al ArticleLikeParam) error {
	if err := al.Rdb.SAdd(al.Context, articleLikeSetKey(al.ArticleID), al.UserID).Err(); err != nil {
		return err
	}
	return nil
}

// IsUserInArticleLikeSet 用户是否存在文章点赞集合
func IsUserInArticleLikeSet(al ArticleLikeParam) (bool, error) {
	isLike, err := al.Rdb.SIsMember(al.Context, articleLikeSetKey(al.ArticleID), al.UserID).Result()
	if err != nil {
		return false, err
	}
	return isLike, nil
}

// IncreaseArticleLikeNum 增加文章点赞数
func IncreaseArticleLikeNum(al ArticleLikeParam, increment int64) error {
	if err := al.Rdb.HIncrBy(al.Context, articleLikeHashKey(), al.ArticleID, increment).Err(); err != nil {
		return err
	}
	return nil
}

// GetArticleLikeNum 获取文章点赞数
func GetArticleLikeNum(al ArticleLikeParam) (int, error) {
	likeNumString, err := al.Rdb.HGet(al.Context, articleLikeHashKey(), al.ArticleID).Result()
	if err != nil {
		return -1, err
	}

	likeNum, err := strconv.Atoi(likeNumString)
	if err != nil {
		return -1, err
	}
	return likeNum, nil
}

// GetArticleLikeSet 获取文章点赞用户ID集合
func GetArticleLikeSet(al ArticleLikeParam) ([]string, error) {
	uids, err := al.Rdb.SMembers(al.Context, articleLikeSetKey(al.ArticleID)).Result()
	if err != nil {
		return nil, err
	}
	return uids, nil
}

// RemoveUserFromArticleLikeSet 将用户ID移除文章点赞集合
func RemoveUserFromArticleLikeSet(al ArticleLikeParam) error {
	if err := al.Rdb.SRem(al.Context, articleLikeSetKey(al.ArticleID), al.UserID).Err(); err != nil {
		return err
	}
	return nil
}

// GetArticleLikeHash 获取文章点赞hash
func GetArticleLikeHash(al ArticleLikeParam) (map[string]string, error) {
	res, err := al.Rdb.HGetAll(al.Context, articleLikeHashKey()).Result()
	if err != nil {
		return nil, err
	}
	return res, nil
}

// SetArticleLikes 设置文章点赞量
func SetArticleLikes(a ArticleLikeParam, likes int) error {
	if err := a.Rdb.HSet(a.Context, articleLikeHashKey(), a.ArticleID, likes).Err(); err != nil {
		return err
	}
	return nil
}
