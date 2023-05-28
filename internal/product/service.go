package product

import (
	"rarefinds-backend/common/date_db"
	"rarefinds-backend/common/errors"
	"rarefinds-backend/internal/product/domain"
	"rarefinds-backend/internal/product/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductsService interface {
	CreateProduct(domain.Product) (*domain.Product, *errors.Error)
	GetAll() ([]domain.Product, *errors.Error)
}

type productsService struct {
	repository repository.ProductsRep
}

func NewService(rep repository.ProductsRep) ProductsService {
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

func (s *productsService) GetAll() ([]domain.Product, *errors.Error) {
	products, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return products, err

}