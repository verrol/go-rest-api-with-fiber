package handler

import (
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/gofiber/fiber/v2"
	"github.com/verrol/go-rest-api-with-fiber/model"
	"github.com/verrol/go-rest-api-with-fiber/respository"
	"xorm.io/xorm"
)

type ProductHandlers struct {
	repo respository.ProductRepository
}

func NewProductHandlers(db *xorm.Engine) *ProductHandlers {
	return &ProductHandlers{repo: respository.NewProductRepository(db)}
}

func (ph *ProductHandlers) GetAllProducts(c *fiber.Ctx) error {
	products, err := ph.repo.GetAllProducts()
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	// return products in JSON format
	return c.JSON(products)
}

func (ph *ProductHandlers) GetProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		log.Info("no id provided", "err", err)
		return c.SendStatus(http.StatusInternalServerError)
	}

	product, err := ph.repo.GetProduct(int64(id))

	if err != nil {
		log.Error("Error while trying to retrieve product", "id", id, "err", err)
		return c.SendStatus(http.StatusInternalServerError)
	} else if (product == model.Product{}) {
		log.Info("No record found!", "id", id, "err", err)
		return c.SendStatus(http.StatusNotFound)
	}

	return c.JSON(product)
}

func (ph *ProductHandlers) CreateProduct(c *fiber.Ctx) error {
	var p model.Product
	if err := c.BodyParser(&p); err != nil {
		log.Error("Error while parsing product", "err", err)
		return c.SendStatus(http.StatusInternalServerError)
	}

	id, err := ph.repo.CreateProduct(p)
	if err != nil {
		log.Error("Error while inserting product", "err", err)
		return c.SendStatus(http.StatusInternalServerError)
	}

	return c.JSON(id)
}

func (ph *ProductHandlers) DeleteProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		log.Info("no id provided", "err", err)
		return c.SendStatus(http.StatusInternalServerError)
	}

	err = ph.repo.DeleteProduct(int64(id))
	if err != nil {
		log.Error("Error while deleting product", "err", err, "id", id)
		return c.SendStatus(http.StatusInternalServerError)
	}

	log.Info("Product deleted", "id", id)
	return c.SendStatus(http.StatusOK)
}
