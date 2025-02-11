package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Name      string    `json:"name"`
	Age       int       `json:"age"`
	HeadPhoto string    `json:"head_photo" gorm:"default:http://localhost:8899/uploads/default/head_photo.jpg"`
	IsOwner   bool      `json:"is_owner" gorm:"default:false" column:"is_owner"`
	Articles  []Article `json:"article"`
	Comments  []Comment `json:"comments"`
}
