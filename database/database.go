package database

import (
	"fmt"
	"os"

	"github.com/golang-mitrah/gin-RestAPI-postgres-orm/helper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() error {
	err := helper.Configure(".env")
	if err != nil {
		fmt.Println("error is loading env file")
		return err
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")

	connectionURI := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", user, password, host, port, dbname)

	//connecting to postgres-SQL
	connection, err := gorm.Open(postgres.Open(connectionURI), &gorm.Config{})
	if err != nil {
		panic("could not connect to DB!")
	}

	DB = connection

	return nil
}
