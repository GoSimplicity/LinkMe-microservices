package data

import (
	"github.com/GoSimplicity/LinkMe-microservices/app/linkme-post/internal/biz"
	"gorm.io/gorm"
)

func InitTables(db *gorm.DB) error {
	err := db.AutoMigrate(&biz.Post{}, &biz.Plate{})
	if err != nil {
		return err
	}
	return nil
}
