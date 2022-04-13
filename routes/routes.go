package routes

import (
	"BookCrud/controllers"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(func(context *gin.Context) {
		context.Set("db", db)
	})

	r.GET("/tasks", controllers.FindTasks)
	r.POST("/task", controllers.CreateTask)
	r.GET("/task/:id", controllers.FindTask)
	r.PATCH("/task/:id", controllers.UpdateTask)
	r.DELETE("/task/:id", controllers.DeleteTask)

	return r
}
