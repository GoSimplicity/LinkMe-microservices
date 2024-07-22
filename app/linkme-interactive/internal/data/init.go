package data

import "gorm.io/gorm"

func InitTables(db *gorm.DB) error {
	err := db.AutoMigrate(&Interactive{}, &UserLike{}, &UserCollection{})
	if err != nil {
		return err
	}
	return nil
}
