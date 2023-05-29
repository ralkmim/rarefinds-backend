package auth

import (
	"context"
	"rarefinds-backend/common/crypto_utils"
	"rarefinds-backend/common/date_db"
	"rarefinds-backend/common/errors"
	"rarefinds-backend/internal/auth/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UsersService interface {
	CreateUser(*domain.SignUpInput, context.Context) *errors.Error
}

type usersService struct {
	repository UsersRep
}

func NewService(repo UsersRep) UsersService {
	return &usersService{
		repository: repo,
	}
}

func (s *usersService) CreateUser(payload *domain.SignUpInput, ctx context.Context) *errors.Error {
	newUser := domain.User{
		ID: 			primitive.NewObjectID(),
		Name: 			payload.Name,
		LastName: 		payload.LastName,
		Username: 		payload.Username,
		Email: 			payload.Email,
		Password: 		crypto_utils.GetMd5(payload.Password),
		PhoneNumber: 	payload.PhoneNumber,
		CreatedAt: 		date_db.GetNowDBFormat(),
		UpdatedAt: 		date_db.GetNowDBFormat(),	
	}

	if err := s.repository.CreateUser(newUser, ctx); err != nil {
		return err
	}

	return nil
}