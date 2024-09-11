package main

import (
	"memorandum-backend/internal/app"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	dsn := "host=" + os.Getenv("POSTGRES_HOST") +
		" user=" + os.Getenv("POSTGRES_USER") +
		" password=" + os.Getenv("POSTGRES_PASSWORD") +
		" dbname=" + os.Getenv("POSTGRES_DB") +
		" port=" + os.Getenv("POSTGRES_PORT")

	app.Run(dsn, "8080")
}
