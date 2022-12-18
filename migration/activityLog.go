package migration

import (
	"github.com/MahmoudMekki/XM-Task/pkg/models"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func listLogs() []*gormigrate.Migration {
	return []*gormigrate.Migration{
		{
			ID: "create-logs-table-and-id-index",
			Migrate: func(db *gorm.DB) error {
				return db.AutoMigrate(&models.Log{})
			},
			Rollback: func(db *gorm.DB) error {
				return db.Migrator().DropTable(models.ActivityLogTableName)
			},
		},
	}
}
