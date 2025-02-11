package db

import (
	"gin-blog/internal/modules/user/models"
	"gorm.io/gorm"
)

type UserSqlParam struct {
	Db   *gorm.DB
	User models.User
}

// QueryUserById 根据id查找用户
func QueryUserById(u UserSqlParam) (models.User, error) {
	if err := u.Db.Take(&u.User, u.User.ID).Error; err != nil {
		return models.User{}, err
	}
	return u.User, nil
}

// QueryUserByEmail 根据邮箱查找用户
func QueryUserByEmail(u UserSqlParam) (models.User, error) {
	if err := u.Db.Take(&u.User, "email = ?", u.User.Email).Error; err != nil {
		return models.User{}, err
	}
	return u.User, nil
}

// InsertUser 添加用户信息
func InsertUser(u UserSqlParam) error {
	if err := u.Db.Create(&u.User).Error; err != nil {
		return err
	}
	return nil
}

// QueryOwnerUser 查询我的信息
func QueryOwnerUser(u UserSqlParam) (models.User, error) {
	var user models.User
	if err := u.Db.Where("is_owner = ?", true).First(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

// UpdateHeadPhoto 更新用户头像
func UpdateHeadPhoto(u UserSqlParam) error {
	if err := u.Db.Model(&models.User{}).Where("id = ?", u.User.ID).Update("head_photo", u.User.HeadPhoto).Error; err != nil {
		return err
	}
	return nil
}
