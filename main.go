package main

import (
	"golang-clean-architecture/config"
	"golang-clean-architecture/controller"
	"golang-clean-architecture/exception"
	"golang-clean-architecture/repository"
	"golang-clean-architecture/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	// Setup Configuration
	configuration := config.New()
	database := config.NewMongoDatabase(configuration)

	// Setup Repository
	productRepository := repository.NewProductRepository(database)

	// Setup Service
	productService := &service.ProductService{
		Insert:    productRepository.Insert,
		FindAll:   productRepository.FindAll,
		DeleteAll: productRepository.DeleteAll,
	}

	// Setup Controller
	productController := &controller.ProductController{
		CreateProduct: productService.Create,
		ListProducts:  productService.List,
	}

	// Setup Fiber
	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())

	// Setup Routing
	productController.Route(app)

	// Start App
	err := app.Listen(":3000")
	exception.PanicIfNeeded(err)
}
