package product

import (
	"context"
	"rarefinds-backend/common/date_db"
	"rarefinds-backend/common/errors"
	"rarefinds-backend/internal/product/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductsService interface {
	CreateProduct(domain.Product) (*domain.Product, *errors.Error)
	GetAll(context.Context) ([]domain.Product, *errors.Error)
	GetProduct(primitive.ObjectID, context.Context) (*domain.Product, *errors.Error)
	SearchProducts(string, context.Context) ([]domain.Product, *errors.Error)
}

type productsService struct {
	repository ProductsRep
}

func NewService(rep ProductsRep) ProductsService {
	return &productsService{
		repository: rep,
	}
}

func (s *productsService) CreateProduct(product domain.Product) (*domain.Product, *errors.Error) {
	product.ID = primitive.NewObjectID()
	product.CreatedAt = date_db.GetNowDBFormat()

	if err := s.repository.CreateProduct(product); err != nil {
		return nil, err
	}

	return &product, nil
}

func (s *productsService) GetAll(ctx context.Context) ([]domain.Product, *errors.Error) {
	products, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return products, err

}

func (s *productsService) GetProduct(productId primitive.ObjectID, ctx context.Context) (*domain.Product, *errors.Error) {
	user, err := s.repository.GetById(productId, ctx)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *productsService) SearchProducts(search string, ctx context.Context) ([]domain.Product, *errors.Error) {
	return s.repository.SearchProducts(search, ctx)
}