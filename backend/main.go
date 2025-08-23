package main

import (
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/phamtoan030222/test/backend/handlers"
)

func main() {
	router := gin.Default()

	// Fix cảnh báo proxy (không trust proxy nào)
	if err := router.SetTrustedProxies(nil); err != nil {
		log.Fatal(err)
	}

	// Lấy domain client từ biến môi trường (nếu cần cho CORS)
	frontendURL := os.Getenv("FRONTEND_URL")
	if frontendURL == "" {
		frontendURL = "http://localhost:5173" // dev local Vite
	}

	// CORS config
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			frontendURL,
		},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// API routes
	api := router.Group("/api")
	{
		api.POST("/tasks", handlers.CreateTaskHandler)
		api.GET("/tasks", handlers.GetTaskHandler)
		api.PATCH("/tasks/:id", handlers.UpdateTaskHandler)
		api.DELETE("/tasks/:id", handlers.DeleteTaskHandler)
	}

	// Serve client tĩnh (Vite build)
	// Giả sử build client vào folder ./client/dist
	router.Static("/", "./client/dist")
	router.NoRoute(func(c *gin.Context) {
		c.File("./client/dist/index.html")
	})

	// Lấy PORT từ Railway env
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // fallback local
	}

	log.Println("✅ Server running on port", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("❌ Failed to start server: ", err)
	}
}
