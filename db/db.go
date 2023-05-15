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

func ConnectDatabase() {

	dsn := fmt.Sprintf("%s", os.Getenv("DB_URL"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// fmt.Println()

	// user := User{Name: "Jinzhu", Age: 18}

	// result := db.Create(&user)

	// logger := log.InitLog(log.Logger{Name: "DB"})
	// logger.Log("Init DB")

	// fmt.Printf("> %v %v \n", db, err)
	// fmt.Printf(">>>>>>> result : %s", result)

}
