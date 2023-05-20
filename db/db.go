package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	Name string `sql:"name`
	Age  int    `sql:"age"`
}

func ConnectDatabase() (*gorm.DB, error) {

	dsn := fmt.Sprintf("%s", os.Getenv("DB_URL"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db, err

}
