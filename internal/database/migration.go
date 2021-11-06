package database

import (
	"github.com/jinzhu/gorm"
	"github.com/xvbnm48/rest-api-golang/internal/comment"
)

// migrateDB  - migrates our database and  create our comment table
func MigrateDB(db *gorm.DB) error {
	if result := db.AutoMigrate(&comment.Comment{}); result.Error != nil {
		return result.Error
	}

	return nil
}
