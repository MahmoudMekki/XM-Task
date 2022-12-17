package userDAL

import (
	"github.com/MahmoudMekki/XM-Task/database"
	"github.com/MahmoudMekki/XM-Task/pkg/models"
)

func IsEmailExists(email string) (exists bool, err error) {
	user := models.User{}
	dbConn, err := database.GetDatabaseConnection()
	if err != nil {
		return false, err
	}
	err = dbConn.Table(models.UsersTableName).Where("email=?", email).First(&user).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func IsUserNameExists(username string) (exists bool, err error) {
	user := models.User{}
	dbConn, err := database.GetDatabaseConnection()
	if err != nil {
		return false, err
	}
	err = dbConn.Table(models.UsersTableName).Where("user_name=?", username).First(&user).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func CreateUser(user models.User) (models.User, error) {
	dbConn, err := database.GetDatabaseConnection()
	if err != nil {
		return models.User{}, err
	}
	err = dbConn.Table(models.UsersTableName).Create(&user).Error
	return user, err
}

func GetUserByEmail(email string) (models.User, error) {
	user := models.User{}
	dbConn, err := database.GetDatabaseConnection()
	if err != nil {
		return models.User{}, err
	}
	err = dbConn.Table(models.UsersTableName).Where("email=?", email).First(&user).Error
	return user, err
}
