package db

import (
	"gin-blog/internal/modules/user/models"
	"gorm.io/gorm"
)

type ImageSqlParam struct {
	Db    *gorm.DB
	Image models.ArticleImage
}

// InsertArticleImage 添加文章图片
func InsertArticleImage(i ImageSqlParam) error {
	if err := i.Db.Model(&models.ArticleImage{}).Create(&i.Image).Error; err != nil {
		return err
	}
	return nil
}
