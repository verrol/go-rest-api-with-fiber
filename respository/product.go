package respository

import (
	"database/sql"

	"github.com/charmbracelet/log"
	"github.com/verrol/go-rest-api-with-fiber/model"
)

// handle persistence database persistence by abtracting it away the sql DB details
type ProductRepository interface {
	GetAllProducts() (model.Products, error)
	GetProduct(id int) (model.Product, error)
	CreateProduct(product model.Product) (int, error)
	DeleteProduct(id int) error
}

type productRepoImpl struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepoImpl{db}
}

func (pr *productRepoImpl) GetAllProducts() (model.Products, error) {
	// query product table in the db
	query := `SELECT * FROM products order by name`
	rows, err := pr.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	result := model.Products{}
	for rows.Next() {
		product := model.Product{}
		err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Quantity,
			&product.Category, &product.Description)
		if err != nil {
			return nil, err
		}
		result = append(result, product)
	}

	log.Info("query successfull", "query", query, "records", len(result))

	return result, nil
}

func (pr *productRepoImpl) GetProduct(id int) (model.Product, error) {
	product := model.Product{}
	query := `SELECT * FROM products WHERE id=$1`
	row, err := pr.db.Query(query, id)
	if err != nil {
		log.Info("query error", "id", id, "err", err)
		return product, err
	}

	defer row.Close()
	if row.Next() {
		err = row.Scan(&product.ID, &product.Name, &product.Price, &product.Quantity,
			&product.Category, &product.Description)
	} else {
		log.Info("No record found!", "id", id, "err", err)
		return product, err
	}

	if err != nil {
		log.Error("Error while scanning product", "id", id, "err", err)
		return product, err
	}

	log.Info("Product found", "id", id, "name", product.Name)
	return product, nil
}

func (pr *productRepoImpl) CreateProduct(product model.Product) (int, error) {
	var id int
	p := product

	query := `INSERT INTO products (name, price, quantity, category, description) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err := pr.db.QueryRow(query, p.Name, p.Price, p.Quantity, p.Category, p.Description).Scan(&id)
	if err != nil {
		log.Error("Error while inserting product", "err", err)
		return id, err
	}

	log.Info("Product inserted", "id", id)
	return id, nil
}

func (pr *productRepoImpl) DeleteProduct(id int) error {
	query := `DELETE FROM products WHERE id=$1`
	_, err := pr.db.Exec(query, id)
	if err != nil {
		log.Error("Error while deleting product", "err", err, "id", id)
		return err
	}

	log.Info("Product deleted", "id", id)
	return nil
}
