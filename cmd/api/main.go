package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"

	"cinema-ticket-system/internal/controllers"
	"cinema-ticket-system/internal/middleware"
	"cinema-ticket-system/internal/repositories"
	"cinema-ticket-system/internal/services"
	"cinema-ticket-system/pkg/config"
	"cinema-ticket-system/pkg/utils"
)

func main() {
	// Initialize database
	db, err := config.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Initialize JWT utility
	jwtUtil := utils.NewJWTUtil("your-secret-key", 24*time.Hour)

	// Initialize repositories
	userRepo := repositories.NewUserRepository(db)
	scheduleRepo := repositories.NewScheduleRepository(db)

	// Initialize services
	// authService tidak digunakan, jadi kita hapus deklarasinya
	scheduleService := services.NewScheduleService(scheduleRepo)

	// Initialize controllers
	authController := controllers.NewAuthController(userRepo, jwtUtil)
	scheduleController := controllers.NewScheduleController(scheduleService)

	// Setup router
	router := gin.Default()

	// Auth routes
	router.POST("/login", authController.Login)

	// API group with auth middleware
	api := router.Group("/api")
	api.Use(middleware.AuthMiddleware(jwtUtil))
	{
		// Schedule routes
		api.GET("/schedules", scheduleController.GetAllSchedules)
		api.POST("/schedules", scheduleController.CreateSchedule)
		api.GET("/schedules/:id", scheduleController.GetSchedule)
		api.PUT("/schedules/:id", scheduleController.UpdateSchedule)
		api.DELETE("/schedules/:id", scheduleController.DeleteSchedule)
	}

	// Start server
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}