package migrations

import (
	"github.com/abnerdsilva/webapi_mvc_with_go/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Book{})
}
