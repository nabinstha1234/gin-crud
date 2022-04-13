package main

import (
	"BookCrud/models"
	"BookCrud/routes"
)

func main() {
	db := models.SetupDB()
	db.AutoMigrate(&models.Task{})
	r := routes.SetupRoutes(db)
	r.Run()
}
