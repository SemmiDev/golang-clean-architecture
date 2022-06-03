package controller

import (
	"golang-clean-architecture/exception"
	"golang-clean-architecture/model"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ProductController struct {
	CreateProduct func(request model.CreateProductRequest) (response model.CreateProductResponse)
	ListProducts  func() (responses []model.GetProductResponse)
}

func (controller *ProductController) Route(app *fiber.App) {
	app.Post("/api/products", controller.Create)
	app.Get("/api/products", controller.List)
}

func (controller *ProductController) Create(c *fiber.Ctx) error {
	var request model.CreateProductRequest
	err := c.BodyParser(&request)
	exception.PanicIfNeeded(err)

	request.Id = uuid.New().String()
	response := controller.CreateProduct(request)
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

func (controller *ProductController) List(c *fiber.Ctx) error {
	responses := controller.ListProducts()
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   responses,
	})
}
