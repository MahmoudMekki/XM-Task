package migration

import (
	"github.com/MahmoudMekki/XM-Task/pkg/models"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func listUsers() []*gormigrate.Migration {
	return []*gormigrate.Migration{
		{
			ID: "create-Users-table-and-id-index",
			Migrate: func(db *gorm.DB) error {
				return db.AutoMigrate(&models.User{})
			},
			Rollback: func(db *gorm.DB) error {
				return db.Migrator().DropTable(models.UsersTableName)
			},
		},
	}
}
