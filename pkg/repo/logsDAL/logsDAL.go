package logsDAL

import (
	"github.com/MahmoudMekki/XM-Task/database"
	"github.com/MahmoudMekki/XM-Task/pkg/models"
)

func CreateLog(log models.Log) error {
	dbConn, err := database.GetDatabaseConnection()
	if err != nil {
		return err
	}
	err = dbConn.Table(models.ActivityLogTableName).Create(&log).Error
	return err
}
