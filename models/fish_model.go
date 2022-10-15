package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Fish struct {
	Id            primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name          string             `bson:"name" json:"name" validate:"required"`
	CookingMethod []string           `bson:"cookingMethod" json:"cookingMethod" validate:"required"`
	OkToEat       []string           `bson:"okToEat" json:"okToEat" validate:"required"`
	MostDelicious []string           `bson:"mostDelicious" json:"mostDelicious" validate:"required"`
}
