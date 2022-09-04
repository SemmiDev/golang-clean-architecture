package controller

import (
	"golang-clean-architecture/exception"
	"golang-clean-architecture/model"
	"golang-clean-architecture/service"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type ProductController struct {
	productService *service.ProductServiceImpl
}

func NewProductController(productService *service.ProductServiceImpl) *ProductController {
	return &ProductController{productService: productService}
}

func (controller *ProductController) Route(app *fiber.App) {
	app.Post("/api/products", controller.Create)
	app.Get("/api/products", controller.List)
}

func (controller *ProductController) Create(c *fiber.Ctx) error {
	var request model.CreateProductRequest
	err := c.BodyParser(&request)
	exception.PanicIfNeeded(err)

	response := controller.productService.Create(request)
	return c.JSON(model.WebResponse{
		Code:   http.StatusCreated,
		Status: http.StatusText(http.StatusCreated),
		Data:   response,
	})
}

func (controller *ProductController) List(c *fiber.Ctx) error {
	return c.JSON(model.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   <-controller.productService.List(),
	})
}
