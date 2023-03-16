package repository

import (
	"go_gin_mini_ecommerce/models"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB -> connection database
func DB() *gorm.DB {
	db, err := gorm.Open("mysql", os.Getenv("DB_USERNAME")+":"+os.Getenv("DB_PASSWORD")+"@/"+os.Getenv("DB_NAME")+"?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		log.Fatal("Error connecting to database")
		return nil
	}
	db.AutoMigrate(&models.User{})
	return db
}