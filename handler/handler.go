package handler

import "github.com/gofiber/fiber/v2"

func GetAllProducts(c *fiber.Ctx) error {
	return c.SendString("got all products")
}

func GetProduct(c *fiber.Ctx) error {
	return c.SendString("got product with id " + c.Params("id"))
}

func CreateProduct(c *fiber.Ctx) error {
	return c.SendString("created product")
}

func DeleteProduct(c *fiber.Ctx) error {
	return c.SendString("deleted product with id " + c.Params("id"))
}
