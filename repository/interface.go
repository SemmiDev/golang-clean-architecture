package repository

import "golang-clean-architecture/entity"

type Creator interface {
	Create(product entity.Product)
}

type Finder interface {
	FindAll() <-chan []entity.Product
}

type Deleter interface {
	DeleteAll()
}
