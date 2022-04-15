package main

import (
	"BookCrud/models"
	"BookCrud/routes"
)

func main() {
	db := models.SetupDB()
	db.AutoMigrate(&models.Task{})
	db.AutoMigrate(&models.User{})
	r := routes.SetupRoutes(db)
	r.Run()
}
