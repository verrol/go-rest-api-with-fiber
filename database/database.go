package database

import (
	"fmt"
	"strconv"

	// load SQL drivers that we want to support
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"xorm.io/xorm"

	"github.com/charmbracelet/log"
	"github.com/verrol/go-rest-api-with-fiber/config"
)

var db *xorm.Engine

func GetDbConnection() *xorm.Engine {
	return db
}

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
	dbDriver := config.Config("DB_DRIVER")

	var connStr string
	switch dbDriver {
	case "sqlite3":
		connStr = dbName
	case "mysql":
		connStr = fmt.Sprintf("mysql://%s:%s@%s:%d/%s?parseTime=true", dbUsername, dbPassword, dbHost, port, dbName)
	case "postgres":
		connStr = fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", dbUsername, dbPassword, dbHost, port, dbName, dbSqlMode)
	default:
		return fmt.Errorf("unsupported database driver")
	}

	db, err = xorm.NewEngine(dbDriver, connStr)
	if err != nil {
		return fmt.Errorf("error while connecting to database: %w", err)
	}

	if err = db.Ping(); err != nil {
		return fmt.Errorf("error while pinging database: %w", err)
	}
	log.Info("database connection established")

	db.ShowSQL(true)

	err = CreateProductTable(db)
	if err != nil {
		return err
	}

	return nil
}
