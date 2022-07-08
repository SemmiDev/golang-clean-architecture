package main

import (
	"golang-clean-architecture/config"
	"golang-clean-architecture/controller"
	"golang-clean-architecture/exception"
	"golang-clean-architecture/repository"
	"golang-clean-architecture/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	cfg, err := config.LoadConfig(".")
	exception.PanicIfNeeded(err)

	database := config.NewMongoDatabase(cfg, "prod")

	productRepository := repository.NewProductRepository(database)
	productService := service.NewProductService(productRepository)
	productController := controller.NewProductController(productService)

	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())
	app.Use(logger.New())

	productController.Route(app)

	err = app.Listen(":8080")
	exception.PanicIfNeeded(err)
}
