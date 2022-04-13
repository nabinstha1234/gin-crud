package controllers

import (
	"BookCrud/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"time"
)

type CreateTaskInput struct {
	AssignedTo string `json:"assignedTo"`
	Task       string `json:"task"`
	Deadline   string `json:"deadline"`
}

// Get Task
// Get all the task
func FindTasks(context *gin.Context) {
	db := context.MustGet("db").(*gorm.DB)
	var tasks []models.Task
	db.Find(&tasks)
	context.JSON(http.StatusOK, gin.H{"data": tasks})
}

// Create task

func CreateTask(c *gin.Context) {
	var input CreateTaskInput
	fmt.Println("hello ia m hereSS1")
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("hello ia m hereSS")

	date := "2006-01-02"
	deadline, _ := time.Parse(date, input.Deadline)

	// create task
	task := models.Task{AssingedTo: input.AssignedTo, Task: input.Task, Deadline: deadline}

	db := c.MustGet("db").(*gorm.DB)
	db.Create(&task)

	c.JSON(http.StatusOK, gin.H{"task": task})
}

//GET task/:id

func FindTask(c *gin.Context) {
	var task models.Task
	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("id=?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not found !"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": task})
}

// PATCH /task/:id

func UpdateTask(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var task models.Task

	if err := db.Where("id=?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not fould !"})
		return
	}

	// validate inpute
	var input CreateTaskInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	date := "2006-01-02"
	deadline, _ := time.Parse(date, input.Deadline)

	var updatedInput models.Task
	updatedInput.Deadline = deadline
	updatedInput.AssingedTo = input.AssignedTo
	updatedInput.Task = input.Task

	db.Model(&task).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": task})
}

// Delete /task/:id

func DeleteTask(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var book models.Task

	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&book)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
