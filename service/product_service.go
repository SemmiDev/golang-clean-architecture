package service

import (
	"golang-clean-architecture/entity"
	"golang-clean-architecture/model"
	"golang-clean-architecture/repository"
	"golang-clean-architecture/validation"
)

type ProductServiceImpl struct {
	PR repository.ProductRepository
}

func NewProductService(pr repository.ProductRepository) *ProductServiceImpl {
	return &ProductServiceImpl{PR: pr}
}

func (service *ProductServiceImpl) Create(request model.CreateProductRequest) (response model.CreateProductResponse) {
	validation.Validate(request)

	product := entity.Product{
		Id:       request.Id,
		Name:     request.Name,
		Price:    request.Price,
		Quantity: request.Quantity,
	}

	service.PR.Insert(product)

	response = model.CreateProductResponse{
		Id:       product.Id,
		Name:     product.Name,
		Price:    product.Price,
		Quantity: product.Quantity,
	}
	return response
}

func (service *ProductServiceImpl) List() <-chan []model.GetProductResponse {
	responsesCh := make(chan []model.GetProductResponse)

	go func() {
		var responses []model.GetProductResponse
		products := <-service.PR.FindAll()
		for _, product := range products {
			responses = append(responses, model.GetProductResponse{
				Id:       product.Id,
				Name:     product.Name,
				Price:    product.Price,
				Quantity: product.Quantity,
			})
		}
		responsesCh <- responses
		close(responsesCh)
	}()

	return responsesCh
}
