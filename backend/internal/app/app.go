package app

import (
	"github.com/gofiber/fiber/v2"
)

func Run(port string){
	app := fiber.New()
	app.Listen(":" + port)
}