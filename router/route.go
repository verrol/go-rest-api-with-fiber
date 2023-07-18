package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/verrol/go-rest-api-with-fiber/handler"
	"github.com/verrol/go-rest-api-with-fiber/middleware"
)

func SetupRoutes(app *fiber.App, productHandlers* handler.ProductHandlers) {
	// add auth middleware for the api routes
	api := app.Group("/api", requestid.New(), middleware.AuthReq())

	// add routes
	api.Get("/", productHandlers.GetAllProducts)
	api.Get("/:id", productHandlers.GetProduct)
	api.Post("/", productHandlers.CreateProduct)
	api.Delete("/:id", productHandlers.DeleteProduct)
}
