package main

import (
	"memorandum-backend/internal/app"
	"os"
	"github.com/joho/godotenv"
	"log"
)


func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dsn := "host=" + os.Getenv("POSTGRES_HOST") +
		" user=" + os.Getenv("POSTGRES_USER") +
		" password=" + os.Getenv("POSTGRES_PASSWORD") +
		" dbname=" + os.Getenv("POSTGRES_DB") +
		" port=" + os.Getenv("POSTGRES_PORT") +
		" sslmode=disable" 

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080" 
	}

	app.Run(dsn, port)
}
