package random

import "github.com/gofiber/fiber/v2"

func RandomRoutes(a *fiber.App) {
	router := a.Group("/rand")

	router.Get("/", randomHandler)
}
