package models

import "github.com/jinzhu/gorm"


type User struct {
	gorm.Model
	Name string `json:"name"`
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`	
}

func (User) TableName() string {
	return "users"
}