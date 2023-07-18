package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/verrol/go-rest-api-with-fiber/handler"
	"github.com/verrol/go-rest-api-with-fiber/middleware"
)

func SetupRoutes(app *fiber.App, productHandlers *handler.ProductHandlers) {
	// add auth middleware for the api routes
	router := app.Group("/api", requestid.New(), middleware.AuthReq())

	setupProductRoutes(router, productHandlers)
}
