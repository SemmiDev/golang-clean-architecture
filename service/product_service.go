package service

import (
	"github.com/Rhymond/go-money"
	"github.com/google/uuid"
	"golang-clean-architecture/entity"
	"golang-clean-architecture/model"
	"golang-clean-architecture/repository"
	"golang-clean-architecture/validation"
)

type ProductServiceImpl struct {
	productRepository repository.ProductRepository
}

func NewProductService(productRepository repository.ProductRepository) *ProductServiceImpl {
	return &ProductServiceImpl{productRepository: productRepository}
}

func (service *ProductServiceImpl) Create(request model.CreateProductRequest) (response model.CreateProductResponse) {
	validation.Validate(request)

	ID := uuid.New().String()
	price := money.New(request.Price, money.IDR)

	product := entity.Product{
		Id:       ID,
		Name:     request.Name,
		Price:    price.Amount(),
		Quantity: request.Quantity,
	}

	service.productRepository.Create(product)

	response = model.CreateProductResponse{
		Id:       product.Id,
		Name:     product.Name,
		Price:    product.Price,
		Quantity: product.Quantity,
	}
	return response
}

func (service *ProductServiceImpl) List() <-chan []model.GetProductResponse {
	responsesChannel := make(chan []model.GetProductResponse)

	go func() {
		var responses []model.GetProductResponse
		products := <-service.productRepository.FindAll()
		for _, product := range products {
			responses = append(responses, model.GetProductResponse{
				Id:       product.Id,
				Name:     product.Name,
				Price:    product.Price,
				Quantity: product.Quantity,
			})
		}
		responsesChannel <- responses
		close(responsesChannel)
	}()

	return responsesChannel
}
