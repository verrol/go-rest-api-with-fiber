package database

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/lib/pq"

	"github.com/charmbracelet/log"
	"github.com/verrol/go-rest-api-with-fiber/config"
)

var DB *sql.DB

func Connect() error {
	var err error
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		return fmt.Errorf("error while parsing port: %w", err)
	}

	dbHost := config.Config("DB_HOST")
	dbUsername := config.Config("DB_USERNAME")
	dbPassword := config.Config("DB_PASSWORD")
	dbName := config.Config("DB_NAME")
	dbSqlMode := config.Config("DB_SSL_MODE")

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", dbHost, port, dbUsername, dbPassword, dbName, dbSqlMode)
	DB, err = sql.Open(config.Config("DB_DRIVER"), connStr)
	if err != nil {
		return fmt.Errorf("error while connecting to database: %w", err)
	}

	if err = DB.Ping(); err != nil {
		return fmt.Errorf("error while pinging database: %w", err)
	}

	CreateProductTable()
	log.Info("database connection established")

	return nil
}

func CreateProductTable() {
	createTableSQL := `CREATE TABLE IF NOT EXISTS products (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		price INT NOT NULL,
		quantity INT NOT NULL
		);`
	_, err := DB.Exec(createTableSQL)
	if err != nil {
		log.Error(err)
		return
	}

	log.Info("product table created")
}
