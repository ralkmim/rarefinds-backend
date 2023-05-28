package repository

import (
	"context"
	"fmt"
	"rarefinds-backend/common/database"
	"rarefinds-backend/common/errors"
	"rarefinds-backend/internal/product/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductsRep interface {
	CreateProduct(domain.Product) *errors.Error
	GetAll() ([]domain.Product, *errors.Error)
}

type productsRep struct {}

func NewRepository() ProductsRep {
	return &productsRep{}
}

func (r *productsRep) CreateProduct(product domain.Product) *errors.Error {
	insertProduct, err := database.Products.InsertOne(context.TODO(), product)
	if err != nil {
		return errors.NewInternalServerError("database error")
	}

	productId := insertProduct.InsertedID.(primitive.ObjectID)
	product.ID = productId
	return nil
}

func (r *productsRep) GetAll() ([]domain.Product, *errors.Error) {
	fmt.Println("eu estou aqui")
	filter := bson.M{}

	cursor, err := database.Products.Find(context.TODO(), filter)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer cursor.Close(context.TODO())

	var products []domain.Product

	for cursor.Next(context.TODO()) {
		var product domain.Product
		if err := cursor.Decode(&product); err != nil {
			return nil, errors.NewInternalServerError(err.Error())
		}
		products = append(products, product)
	}

	if err := cursor.Err(); err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}

	return products, nil
}