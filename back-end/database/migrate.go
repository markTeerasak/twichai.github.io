package database

import (
	"github.com/twichai/twichai.github.io/entities"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.Migrator().CreateTable(&entities.Language{})
}
