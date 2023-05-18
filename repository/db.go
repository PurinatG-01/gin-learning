package repository

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() (*gorm.DB, error) {

	dsn := fmt.Sprintf("%s", os.Getenv("DB_URL"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db, err

}
