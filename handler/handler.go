package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/verrol/go-rest-api-with-fiber/database"
	"github.com/verrol/go-rest-api-with-fiber/model"
)

func GetAllProducts(c *fiber.Ctx) error {
	// query product table in the db
	query := `SELECT * FROM products order by name`
	rows, err := database.DB.Query(query)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"status": false,
			"err":    err,
		})
	}

	defer rows.Close()
	result := model.Products{}
	for rows.Next() {
		product := model.Product{}
		err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Quantity,
			&product.Category, &product.Description)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
				"status": false,
				"err":    err,
			})
		}
		result = append(result, product)
	}

	// return products in JSON format
	err = c.JSON(&fiber.Map{
		"success": true,
		"data":    result,
		"message": "All products returned successfully",
	})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"status": false,
			"err":    err,
		})
	}

	return err
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
