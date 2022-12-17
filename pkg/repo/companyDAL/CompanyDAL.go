package companyDAL

import (
	"github.com/MahmoudMekki/XM-Task/database"
	"github.com/MahmoudMekki/XM-Task/pkg/models"
)

func CreateCompany(company models.Company) (models.Company, error) {
	dbConn, err := database.GetDatabaseConnection()
	if err != nil {
		return models.Company{}, err
	}
	err = dbConn.Table(models.CompaniesTableName).Create(&company).Error
	return company, err
}

func GetCompany(id string) (company models.Company, err error) {
	dbConn, err := database.GetDatabaseConnection()
	if err != nil {
		return models.Company{}, err
	}
	err = dbConn.Table(models.CompaniesTableName).Where("id=?", id).Find(&company).Error
	return company, err
}

func DeleteCompany(id string) (err error) {
	dbConn, err := database.GetDatabaseConnection()
	if err != nil {
		return err
	}
	err = dbConn.Where("id =?", id).Delete(&models.Company{}).Error
	return err
}
func UpdateCompany(company models.Company) (err error) {
	dbConn, err := database.GetDatabaseConnection()
	if err != nil {
		return err
	}
	err = dbConn.Updates(company).Error
	return err
}
