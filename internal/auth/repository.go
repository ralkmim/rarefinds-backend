package auth

import (
	"context"
	"rarefinds-backend/common/database"
	"rarefinds-backend/common/errors"
	"rarefinds-backend/common/logger"
	"rarefinds-backend/internal/auth/domain"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UsersRep interface {
	CreateUser(domain.User, context.Context) *errors.Error
	GetLogin(*domain.User, context.Context) *errors.Error
}

type usersRep struct {}

func NewRepository() UsersRep {
	return &usersRep{}
}

func (r *usersRep) CreateUser(user domain.User, ctx context.Context) *errors.Error {
	checkEmail := database.Users.FindOne(ctx, bson.M{"email": user.Email}).Decode(&user)
	if checkEmail == nil {
		return errors.NewStatusConflictError("user with that email already exists")
	}

	checkUsername := database.Users.FindOne(ctx, bson.M{"username": user.Username}).Decode(&user)
	if checkUsername == nil {
		return errors.NewStatusConflictError("user with that username already exists")
	}

	result, err := database.Users.InsertOne(ctx, user)
	if err != nil {
		logger.Error("error when trying to insert user in database", err)
		return errors.NewInternalServerError("database error")
	}

	userId := result.InsertedID.(primitive.ObjectID)
	user.ID = userId
	
	return nil
}

func (r *usersRep) GetLogin(user *domain.User, ctx context.Context) *errors.Error {
	filter := bson.M{
		"$and": []bson.M{
			{"password": user.Password},
			{"$or": []bson.M{
				{"email": user.Email},
				{"username": user.Username},
			}},
		},
	}

	err := database.Users.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return errors.NewBadRequestError("invalid credentials")
	}

	return nil
} 

func (r *usersRep) GetByEmail(user *domain.User, ctx context.Context) *errors.Error {
	err := database.Users.FindOne(ctx, bson.M{"email": strings.ToLower(user.Email)}).Decode(&user)
	if err != nil {
		logger.Error("error when trying to find user by email in database", err)
		return errors.NewNotFoundError("users not found")
	}

	return nil
}