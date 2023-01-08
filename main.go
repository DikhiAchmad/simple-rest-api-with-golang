package main

import (
	"simple-crud-golang/config"
	"simple-crud-golang/models"
	"simple-crud-golang/routes"
)

func main() {
	db:= config.SetupDB()

	db.AutoMigrate(&models.User{})

	r := routes.SetupRoutes(db)
	r.Run(":5000")
}