package app

import (
	"log"

	"gorm.io/driver/postgres"
  	"gorm.io/gorm"
	"github.com/gofiber/fiber/v2"
)

func Run(dsn string, port string){
	Database, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Println("couldn't connect to database", err)
		return
	}
	log.Println(Database)
	app := fiber.New()
	app.Listen(":" + port)
}