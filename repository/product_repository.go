package repository

import (
	"golang-clean-architecture/config"
	"golang-clean-architecture/entity"
	"golang-clean-architecture/exception"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewProductRepository(database *mongo.Database) *productRepositoryImpl {
	return &productRepositoryImpl{
		collection: database.Collection("products"),
	}
}

type productRepositoryImpl struct {
	collection *mongo.Collection
}

func (repository *productRepositoryImpl) Insert(product entity.Product) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	_, err := repository.collection.InsertOne(ctx, bson.M{
		"_id":      product.Id,
		"name":     product.Name,
		"price":    product.Price,
		"quantity": product.Quantity,
	})
	exception.PanicIfNeeded(err)
}

func (repository *productRepositoryImpl) FindAll() (products []entity.Product) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	cursor, err := repository.collection.Find(ctx, bson.M{})
	exception.PanicIfNeeded(err)

	var documents []bson.M
	err = cursor.All(ctx, &documents)
	exception.PanicIfNeeded(err)

	for _, document := range documents {
		products = append(products, entity.Product{
			Id:       document["_id"].(string),
			Name:     document["name"].(string),
			Price:    document["price"].(int64),
			Quantity: document["quantity"].(int32),
		})
	}

	return products
}

func (repository *productRepositoryImpl) DeleteAll() {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	_, err := repository.collection.DeleteMany(ctx, bson.M{})
	exception.PanicIfNeeded(err)
}
