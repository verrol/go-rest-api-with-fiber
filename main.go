package main

import (
	"github.com/charmbracelet/log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/verrol/go-rest-api-with-fiber/config"
	"github.com/verrol/go-rest-api-with-fiber/database"
	"github.com/verrol/go-rest-api-with-fiber/handler"
	"github.com/verrol/go-rest-api-with-fiber/router"
)

func main() {
	// connect to database
	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}

	// create a new Fiber application
	app := fiber.New()

	app.Use("/metrics", monitor.New())
	app.Use(logger.New())

	db := database.DB
	ph := handler.NewProductHandlers(db)
	router.SetupRoutes(app, ph)

	serverPort := config.Config("SERVER_PORT")
	log.Info("Starting server", "port", serverPort)
	app.Listen(":" + serverPort)

}
