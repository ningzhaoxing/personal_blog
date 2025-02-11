package db

import (
	"gin-blog/internal/modules/user/models"
	"gin-blog/internal/modules/user/requests/class"
	"gorm.io/gorm"
)

type ClassSqlParam struct {
	Db    *gorm.DB
	Class class.ClassCreateRequest
}

func QueryClassList(cs ClassSqlParam) ([]models.ArticleClass, error) {
	var acs []models.ArticleClass
	if err := cs.Db.Find(&acs).Error; err != nil {
		return nil, err
	}
	return acs, nil
}

func CreateClass(cs ClassSqlParam) error {
	if err := cs.Db.Create(&models.ArticleClass{
		Name: cs.Class.Name,
	}).Error; err != nil {
		return err
	}
	return nil
}
