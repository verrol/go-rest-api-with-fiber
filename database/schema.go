package database

import "github.com/charmbracelet/log"

func CreateProductTable() {
	createTableSQL := `CREATE TABLE IF NOT EXISTS products (
		id SERIAL PRIMARY KEY,
		name text UNIQUE,
		price INT NOT NULL,
		quantity INT NOT NULL,
		category text NOT NULL,
		description text
		);`
	_, err := DB.Exec(createTableSQL)
	if err != nil {
		log.Error(err)
		return
	}

	log.Info("product table created")
}
