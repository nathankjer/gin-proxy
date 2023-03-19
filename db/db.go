package db

import (
	"github.com/nathankjer/gin-proxy/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func getDatabase(db_connection string) (database *gorm.DB, err error) {
	// Connect to database
	database, err = gorm.Open(postgres.Open(db_connection), &gorm.Config{})
	if err != nil {
		return database, err
	}

	// Migrate database
	err = database.AutoMigrate(
		models.Request{},
	)
	if err != nil {
		return nil, err
	}
	return database, nil

}

func Connect() error {
	database, err := getDatabase("host=localhost dbname=requests port=5432 sslmode=disable")
	if err != nil {
		return err
	}
	DB = database
	return nil
}
