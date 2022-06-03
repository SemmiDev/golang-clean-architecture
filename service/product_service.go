package service

import (
	"golang-clean-architecture/entity"
	"golang-clean-architecture/model"
	"golang-clean-architecture/validation"
)

type ProductService struct {
	Insert    func(product entity.Product)
	FindAll   func() (products []entity.Product)
	DeleteAll func()
}

func (service *ProductService) Create(request model.CreateProductRequest) (response model.CreateProductResponse) {
	validation.Validate(request)

	product := entity.Product{
		Id:       request.Id,
		Name:     request.Name,
		Price:    request.Price,
		Quantity: request.Quantity,
	}

	service.Insert(product)

	response = model.CreateProductResponse{
		Id:       product.Id,
		Name:     product.Name,
		Price:    product.Price,
		Quantity: product.Quantity,
	}
	return response
}

func (service *ProductService) List() (responses []model.GetProductResponse) {
	products := service.FindAll()
	for _, product := range products {
		responses = append(responses, model.GetProductResponse{
			Id:       product.Id,
			Name:     product.Name,
			Price:    product.Price,
			Quantity: product.Quantity,
		})
	}
	return responses
}
