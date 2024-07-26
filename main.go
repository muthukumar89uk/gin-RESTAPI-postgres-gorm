package main

import (
	"fmt"

	"github.com/golang-mitrah/gin-RestAPI-postgres-orm/database"
	"github.com/golang-mitrah/gin-RestAPI-postgres-orm/models"
	"github.com/golang-mitrah/gin-RestAPI-postgres-orm/routes"
)

func main() {
	if err := database.ConnectToDB(); err != nil {
		fmt.Println("Failed to connect to DB", err)
		return
	}

	database.DB.AutoMigrate(&models.User{})

	r := routes.SetupRouter()

	//running
	r.Run()
}
