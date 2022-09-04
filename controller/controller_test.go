package controller

import (
	"golang-clean-architecture/config"
	"golang-clean-architecture/repository"
	"golang-clean-architecture/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func createTestApp() *fiber.App {
	var app = fiber.New(config.NewFiberConfig())
	app.Use(recover.New())
	productController.Route(app)
	return app
}

var configuration, _ = config.LoadConfig("../")

var database = config.NewMongoDatabase(configuration, "test")
var productRepository = repository.NewProductRepository(database)
var productService = service.NewProductService(productRepository)

var productController = NewProductController(productService)

var app = createTestApp()
