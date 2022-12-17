package database

import (
	"fmt"
	"github.com/MahmoudMekki/XM-Task/config"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var dbConn *gorm.DB

func dsn() (dsn string) {
	dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True",
		config.GetEnvVar("DB_USER"),
		config.GetEnvVar("DB_PASSWORD"),
		config.GetEnvVar("DB_HOST"),
		config.GetEnvVar("DB_PORT"),
		config.GetEnvVar("DB_NAME"),
	)
	return dsn
}

func CreateDBConnection() error {
	if dbConn != nil {
		CloseDBConnection(dbConn)
	}
	dsn := dsn()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	sqlDB, err := db.DB()
	sqlDB.SetConnMaxIdleTime(time.Minute * 5)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	dbConn = db
	return err
}
func CloseDBConnection(conn *gorm.DB) {
	sqlDB, err := conn.DB()
	if err != nil {
		log.Err(err).Msg("Error occurred while closing a DB connection")
	}
	defer sqlDB.Close()
}
func GetDatabaseConnection() (*gorm.DB, error) {
	sqlDB, err := dbConn.DB()
	if err != nil {
		return dbConn, err
	}
	if err := sqlDB.Ping(); err != nil {
		return dbConn, err
	}
	return dbConn, nil
}
