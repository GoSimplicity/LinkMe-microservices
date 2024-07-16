package data

import "gorm.io/gorm"

func InitTables(db *gorm.DB) error {
	err := db.AutoMigrate(&Check{})
	if err != nil {
		return err
	}
	return nil
}
