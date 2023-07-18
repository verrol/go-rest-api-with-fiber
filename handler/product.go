package handler

import (
	"database/sql"
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/gofiber/fiber/v2"
	"github.com/verrol/go-rest-api-with-fiber/model"
)

type ProductHandlers struct {
	db *sql.DB
}

func NewProductHandlers(db *sql.DB) *ProductHandlers {
	return &ProductHandlers{db}
}

func (ph *ProductHandlers) GetAllProducts(c *fiber.Ctx) error {
	// query product table in the db
	query := `SELECT * FROM products order by name`
	rows, err := ph.db.Query(query)
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	defer rows.Close()
	result := model.Products{}
	for rows.Next() {
		product := model.Product{}
		err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Quantity,
			&product.Category, &product.Description)
		if err != nil {
			return c.SendStatus(http.StatusInternalServerError)
		}
		result = append(result, product)
	}

	log.Info("query successfull", "query", query, "records", len(result))

	// return products in JSON format
	return c.JSON(result)
}

func (ph *ProductHandlers) GetProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	product := model.Product{}
	query := `SELECT * FROM products WHERE id=$1`
	row, err := ph.db.Query(query, id)
	if err != nil {
		log.Info("query error", "id", id, "err", err)
		return c.SendStatus(http.StatusInternalServerError)
	}

	defer row.Close()
	if row.Next() {
		err = row.Scan(&product.ID, &product.Name, &product.Price, &product.Quantity,
			&product.Category, &product.Description)
	} else {
		log.Info("No record found!", "id", id, "err", err)
		return c.SendStatus(http.StatusNotFound)
	}

	if err != nil {
		log.Error("Error while scanning product", "id", id, "err", err)
		return c.SendStatus(http.StatusInternalServerError)
	}

	log.Info("Product found", "id", id, "name", product.Name)
	return c.JSON(product)
}

func (ph *ProductHandlers) CreateProduct(c *fiber.Ctx) error {
	p := new(model.Product)
	if err := c.BodyParser(p); err != nil {
		log.Error("Error while parsing product", "err", err)
		return c.SendStatus(http.StatusInternalServerError)
	}

	log.Info("Product parsed from request", "name", p.Name)
	query := `INSERT INTO products (name, price, quantity, category, description) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err := ph.db.QueryRow(query, p.Name, p.Price, p.Quantity, p.Category, p.Description).Scan(&p.ID)
	if err != nil {
		log.Error("Error while inserting product", "err", err)
		return c.SendStatus(http.StatusInternalServerError)
	}

	log.Info("Product inserted", "id", p.ID)
	return c.JSON(p)
}

func (ph *ProductHandlers) DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	query := `DELETE FROM products WHERE id=$1`
	_, err := ph.db.Exec(query, id)
	if err != nil {
		log.Error("Error while deleting product", "err", err, "id", id)
		return c.SendStatus(http.StatusInternalServerError)
	}

	log.Info("Product deleted", "id", id)
	return c.SendStatus(http.StatusOK)
}
