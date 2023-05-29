package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID					primitive.ObjectID 	`json:"_id" bson:"_id"`
	Name				string				`json:"name" bson:"name"`
	LastName			string				`json:"last_name" bson:"last_name"`
	Username			string				`json:"username" bson:"username"`
	Email				string 				`json:"email" bson:"email" validate:"required,email"`
	Password			string				`json:"password" bson:"password" validate:"required"`
	PhoneNumber			string				`json:"phone_number" bson:"phone_number"`
	Photo				string				`json:"photo" bson:"photo" validate:"required"`
	TokenExpiresAt		time.Time			`json:"token_expires_at" bson:"token_expires_at"`
	CreatedAt        	string			 	`json:"created_at" bson:"created_at" validate:"required"`
	UpdatedAt        	string				`json:"updated_at" bson:"updated_at" validate:"required"`
}

type SignUpInput struct {
	Name				string				`json:"name" bson:"name"`
	LastName			string				`json:"last_name" bson:"last_name"`
	Username        	string 				`json:"username" bson:"username" binding:"required"`
	Email           	string 				`json:"email" bson:"email" binding:"required"`
	Password            string 				`json:"password" bson:"password" binding:"required,min=8"`
	PhoneNumber			string				`json:"phone_number" bson:"phone_number"`
}

type SignInInput struct {
	Username			string				`json:"username" bson:"username"`
	Email           	string 				`json:"email" bson:"email"`
	Password        	string 				`json:"password" bson:"password" binding:"required,min=8"`
}

type UserResponse struct {
	ID					primitive.ObjectID 	`json:"_id,omitempty" bson:"_id"`
	Name				string				`json:"name" bson:"name"`
	LastName			string				`json:"last_name" bson:"last_name"`
	Username  			string    			`json:"username,omitempty" bson:"username"`
	Email     			string    			`json:"email,omitempty" bson:"email"`
	Photo     			string    			`json:"photo,omitempty" bson:"photo"`
	CreatedAt 			string 				`json:"created_at" bson:"created_at"`
	UpdatedAt 			string 				`json:"updated_at" bson:"updated_at"`
}