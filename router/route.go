package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/verrol/go-rest-api-with-fiber/handler"
	"github.com/verrol/go-rest-api-with-fiber/middleware"
)

func SetupRoutes(app *fiber.App) {
	// add auth middleware for the api routes
	api := app.Group("/api", requestid.New(), logger.New(), middleware.AuthReq())

	// add routes
	api.Get("/", handler.GetAllProducts)
	api.Get("/:id", handler.GetProduct)
	api.Post("/", handler.CreateProduct)
	api.Delete("/:id", handler.DeleteProduct)
}
