package main

import (
	"apigen/services/databases"
	"apigen/services/random"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	databases.InitDB()

	if err != nil {
		fmt.Println("[!] .env not loaded")
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to APIGEN")
	})

	random.RandomRoutes(app)

	app.Listen(":" + os.Getenv("PORT"))
}
