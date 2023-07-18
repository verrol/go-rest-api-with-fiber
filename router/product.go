package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/verrol/go-rest-api-with-fiber/handler"
)

func setupProductRoutes(router fiber.Router, productHandlers *handler.ProductHandlers) {
	// add routes
	router.Get("/products", productHandlers.GetAllProducts)
	router.Get("/products/:id", productHandlers.GetProduct)
	router.Post("/products", productHandlers.CreateProduct)
	router.Delete("/products/:id", productHandlers.DeleteProduct)
}
