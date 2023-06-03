package auth

import (
	"context"
	"rarefinds-backend/common/crypto_utils"
	"rarefinds-backend/common/date_db"
	"rarefinds-backend/common/errors"
	"rarefinds-backend/common/jwt_token"
	"rarefinds-backend/internal/auth/domain"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UsersService interface {
	CreateUser(*domain.SignUpInput, context.Context) *errors.Error
	LoginUser(*domain.SignInInput, context.Context) (string, *errors.Error)
	GetUser(primitive.ObjectID, context.Context) (*domain.UserResponse, *errors.Error)
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

func (s *usersService) LoginUser(payload *domain.SignInInput, ctx context.Context) (string, *errors.Error) {
	user := &domain.User{
		Email: strings.ToLower(payload.Email),
		Username: payload.Username,
		Password: crypto_utils.GetMd5(payload.Password),
	}

	if err := s.repository.GetLogin(user, ctx); err != nil {
		return "", err
	}

	token, err := jwt_token.GenerateToken(60*time.Minute, user.ID, "UbjV&dxxc3dk6wTU")
	if err != nil {
		return "", errors.NewBadRequestError(err.Error())
	}

	return token, nil
}

func (s *usersService) GetUser(userId primitive.ObjectID, ctx context.Context) (*domain.UserResponse, *errors.Error) {
	user, err := s.repository.GetById(userId, ctx)
	if err != nil {
		return nil, err
	}

	return user, nil
}
