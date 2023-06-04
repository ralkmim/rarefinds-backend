package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ID 					primitive.ObjectID		`json:"_id" bson:"_id"`
	Name				string					`json:"name" bson:"name" validate:"required"`
	Description			string					`json:"description" bson:"description"`
	Price 				float64					`json:"price" bson:"price" validate:"required"`
	Image				string					`json:"image" bson:"image"`
	Certified			bool					`json:"certified" bson:"certified"`
	Type 				string					`json:"type" bson:"type" validate:"required"`		
	Condition			string					`json:"condition" bson:"condition" validate:"required"`
	From				string					`json:"from" bson:"from"`
	LimitedEdition		bool					`json:"limited_edition" bson:"limited_edition" validate:"required"`
	OwnerID				string					`json:"owner_id" bson:"owner_id"`
	CreatedAt			string					`json:"created_at" bson:"created_at"`
	UpdatedAt			string					`json:"updated_at" bson:"updated_at"`
}