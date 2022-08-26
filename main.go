package main

import (
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// detaの取得
type Task struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Done bool   `json:"done"`
}

var tasks = make([]Task, 0)

func main() {
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(cors.New(config))

	r.GET("/tasks", func(c *gin.Context) {
		c.JSON(http.StatusOK, tasks)
	})

	r.POST("/tasks", func(c *gin.Context) {
		var task Task
		c.BindJSON(&task)

		task.ID = strconv.Itoa(rand.Int())

		tasks = append(tasks, task)

		c.Status(http.StatusCreated)
	})

	r.PUT("/tasks/:id", func(c *gin.Context) {
		var newTask Task
		c.BindJSON(&newTask)

		for i, task := range tasks {
			if task.ID == c.Param("id") {
				newTask.ID = task.ID
				tasks[i] = newTask
				break
			}
		}

		c.Status(http.StatusOK)
	})

	r.DELETE("/tasks/:id", func(c *gin.Context) {
		for i, task := range tasks {
			if task.ID == c.Param("id") {
				tasks = append(tasks[:i], tasks[i+1:]...)
				break
			}
		}

		c.Status(http.StatusOK)
	})

	r.Run()
}
