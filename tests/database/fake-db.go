package database

import (
	"github.com/MahmoudMekki/XM-Task/database"
	"github.com/MahmoudMekki/XM-Task/migration"
	"github.com/MahmoudMekki/XM-Task/pkg/models"
	"gorm.io/gorm"
)

type Migrations struct {
	Id string
}

func DbSetup() {
	err := database.CreateDBConnection()
	if err != nil {
		panic(err.Error())
	}
	migration.RunMigration()
}
func CleanUpDb(db *gorm.DB) {
	err := db.Migrator().DropTable(&models.Company{}, &models.User{}, &models.Log{}, &Migrations{})
	if err != nil {
		panic(err.Error())
	}
	database.CloseDBConnection(db)
}

func CreateCompany(db *gorm.DB, company models.Company) models.Company {
	err := db.Create(&company).Error
	if err != nil {
		panic(err.Error())
	}
	return company
}
func GetCompany(db *gorm.DB, id string) (models.Company, error) {
	var company models.Company
	err := db.Where("id = ?", id).First(&company).Error
	return company, err
}

func CreateUser(db *gorm.DB, user models.User) models.User {
	err := db.Create(&user).Error
	if err != nil {
		panic(err.Error())
	}
	return user
}
