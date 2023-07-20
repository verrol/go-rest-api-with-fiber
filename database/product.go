package database

import (
	"github.com/charmbracelet/log"
	"xorm.io/xorm"
)

type Product struct {
	Id          int64
	Name        string `xorm:"unique"`
	Price       int 
	Quantity    int `xorm:"default 0"`
	Category    string `xorm:"not null"`
	Description string

	CreatedAt string `xorm:"created"`
	UpdatedAt string `xorm:"updated"`
}

func CreateProductTable(db *xorm.Engine) error {

	err := db.Sync(new(Product))
	if err != nil {
		log.Error("unable to sync entity table", "entity", "product", "error", err)
		return err
	}

	log.Info("product table created")
	return nil
}
