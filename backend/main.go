package main

import (
	"log"

	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/phamtoan030222/test/backend/handlers"
)

func main() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
    AllowOrigins:     []string{"https://test-go-react.vercel.app"},
    AllowMethods:     []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"},
    AllowHeaders:     []string{"Origin", "Content-Type"},
    ExposeHeaders:    []string{"Content-Length"},
    AllowCredentials: true,
    MaxAge:           12 * time.Hour,
}))

	router.POST("/api/tasks", handlers.CreateTaskHandler)
	router.GET("/api/tasks", handlers.GetTaskHandler)
	router.PATCH("/api/tasks/:id", handlers.UpdateTaskHandler)
	router.DELETE("/api/tasks/:id", handlers.DeleteTaskHandler)

    router.NoRoute(func(c *gin.Context) {
        c.File("./client/dist/index.html")
    })
    router.Static("/static", "./client/dist/static")

	log.Println("Server running at http://localhost:4000")
	router.Run(":4000")
}
