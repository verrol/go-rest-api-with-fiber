package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/verrol/go-rest-api-with-fiber/handler"
	"github.com/verrol/go-rest-api-with-fiber/middleware"
)

func SetupRoutes(app *fiber.App, authHandlers *handler.AuthHandlers, productHandlers *handler.ProductHandlers) {
	// auth routes don't need security
	setupAuthRoutes(app, authHandlers)
	// add auth middleware for the api routes
	api := app.Group("/api", requestid.New(), middleware.AuthReq())

	setupProductRoutes(api, productHandlers)
}
