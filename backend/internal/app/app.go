package app

import (
	"memorandum-backend/internal/entities"

	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"memorandum-backend/internal/repository"
    "memorandum-backend/internal/usecases"
    "memorandum-backend/internal/controllers"
    "memorandum-backend/internal/service"
)

func Run(dsn string, port string) {
    log.Println("Connecting to database...")
    Database, err := gorm.Open(postgres.Open(dsn))
    if err != nil {
        log.Println("couldn't connect to database", err)
        return
    }

    log.Println("Applying migrations...")
    err = Database.AutoMigrate(&entities.User{})
    if err != nil {
        log.Fatalf("failed to migrate database: %v", err)
    }

    userRepo := repository.NewGormUserRepository(Database)
    tokenService := service.NewJWTService("secret") 
    userService := usecase.NewUserService(userRepo, tokenService)
    userHandler := controllers.NewUserHandler(userService)

    app := fiber.New()

    app.Post("/register", userHandler.Register)
    app.Post("/login", userHandler.Login)
    log.Println("Starting server on port", port)

    err = app.Listen(":" + port)
    if err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}
