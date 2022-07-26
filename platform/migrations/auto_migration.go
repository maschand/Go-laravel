package migrations

import (
	"github.com/chand19-af/digitels-template/app/models"
	"gorm.io/gorm"
)

func AutoMigration(db gorm.DB) {
	db.AutoMigrate(&models.Hotel{})
	db.AutoMigrate(&models.User{})
}
