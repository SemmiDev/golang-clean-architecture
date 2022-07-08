package controller

import (
	"golang-clean-architecture/exception"
	"golang-clean-architecture/model"
	"golang-clean-architecture/service"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ProductController struct {
	PS *service.ProductServiceImpl
}

func NewProductController(ps *service.ProductServiceImpl) *ProductController {
	return &ProductController{PS: ps}
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
	response := controller.PS.Create(request)
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

func (controller *ProductController) List(c *fiber.Ctx) error {
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   <-controller.PS.List(),
	})
}
