package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/verrol/go-rest-api-with-fiber/handler"
)

func setupAuthRoutes(router fiber.Router, authHandlers *handler.AuthHandlers) {
	router.Post("/signup", authHandlers.Register)
}
