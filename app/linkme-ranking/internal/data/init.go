package data

import "gorm.io/gorm"

func InitTables(db *gorm.DB) error {
	err := db.AutoMigrate()
	if err != nil {
		return err
	}
	return nil
}
