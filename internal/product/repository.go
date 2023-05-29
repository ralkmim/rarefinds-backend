package product

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
	GetAll(context.Context) ([]domain.Product, *errors.Error)
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

func (r *productsRep) GetAll(ctx context.Context) ([]domain.Product, *errors.Error) {
	filter := bson.M{}

	cursor, err := database.Products.Find(ctx, filter)
	if err != nil {
		fmt.Println(err)
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer cursor.Close(ctx)

	var products []domain.Product

	for cursor.Next(ctx) {
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