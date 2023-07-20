package respository

import (
	"github.com/charmbracelet/log"
	"github.com/verrol/go-rest-api-with-fiber/model"
	"xorm.io/xorm"
)

// handle persistence database persistence by abtracting it away the sql DB details
type ProductRepository interface {
	GetAllProducts() (model.Products, error)
	GetProduct(id int64) (model.Product, error)
	CreateProduct(product model.Product) (int64, error)
	DeleteProduct(id int64) error
}

type productRepoImpl struct {
	db *xorm.Engine
}

func NewProductRepository(db *xorm.Engine) ProductRepository {
	return &productRepoImpl{db}
}

func (pr *productRepoImpl) GetAllProducts() (model.Products, error) {
	products := model.Products{}
	if err := pr.db.Find(&products); err != nil {
		log.Error("Error while getting products", "err", err)
		return nil, err
	}

	return products, nil
}

func (pr *productRepoImpl) GetProduct(id int64) (model.Product, error) {
	product := model.Product{}

	if has, err := pr.db.ID(id).Get(&product); err != nil {
		log.Info("error while getting product", "id", id, "err", err)
		return product, err
	} else if !has {
		log.Info("No record found!", "id", id)
		return product, nil
	}

	return product, nil
}

func (pr *productRepoImpl) CreateProduct(product model.Product) (int64, error) {
	product.Id = 0 // to avoid id conflict and let db generate id

	if _, err := pr.db.Insert(&product); err != nil {
		log.Error("Error while inserting product", "err", err)
		return 0, err
	}

	return product.Id, nil
}

func (pr *productRepoImpl) DeleteProduct(id int64) error {
	if _, err := pr.db.ID(id).Delete(new(model.Product));err != nil {
		log.Error("Error while deleting product", "err", err, "id", id)
		return err
	}

	return nil
}
