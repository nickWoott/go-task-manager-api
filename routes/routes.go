package routes

import 	"github.com/gin-gonic/gin"


 var Router = gin.Default()

	// here are the tasks with their methods, I assume
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/tasks", handlers.GetTasks(tasks))
	router.GET("/tasks/:id", handlers.GetTask(tasks))
	router.PUT("/tasks/:id", handlers.UpdateTask(tasks))
	router.DELETE("/tasks/:id", handlers.RemoveTask(tasks))
	router.POST("/tasks", handlers.AddTask(tasks))

	func SetupRouter(tasks  []data.tas)