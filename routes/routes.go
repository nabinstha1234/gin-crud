package routes

import (
	"BookCrud/controllers"
	"BookCrud/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	public := r.Group("/api")
	public.Use(func(context *gin.Context) {
		context.Set("db", db)
	})

	public.GET("/tasks", controllers.FindTasks)
	public.POST("/task", controllers.CreateTask)
	public.GET("/task/:id", controllers.FindTask)
	public.PATCH("/task/:id", controllers.UpdateTask)
	public.DELETE("/task/:id", controllers.DeleteTask)

	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)

	protected := r.Group("/api/admin")

	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/user", controllers.CurrentUser)

	return r
}
