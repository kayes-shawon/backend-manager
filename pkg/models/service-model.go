package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id       primitive.ObjectID `json:"id,omitempty"`
	Name     string             `json:"name,omitempty" validate:"required"`
	Location string             `json:"location,omitempty" validate:"required"`
	Title    string             `json:"title,omitempty" validate:"required"`
}

// type Service struct {
// 	ID          primitive.ObjectID `json:"id,omitempty"`
// 	Name        string             `json:"name,omitempty"`
// 	Statement   string             `json:"statement,omitempty"`
// 	Description string             `json:"description,omitempty"`
// 	Input       string             `json:"input,omitempty"`
// 	Output      string             `json:"output,omitempty"`
// 	Image       string             `json:"image,omitempty"`
// }

type Service struct {
	ID          primitive.ObjectID `json:"id,omitempty"`
	Name        string             `json:"name,omitempty"`
	Statement   string             `json:"statement,omitempty"`
	Description string             `json:"description,omitempty"`
	Input       string             `json:"input,omitempty"`
	Output      string             `json:"output,omitempty"`
	Image       string             `json:"image,omitempty"`
}

// type Service struct {
// 	ID          primitive.ObjectID `json:"id,omitempty" bson:"id,omitempty"`
// 	Name        string             `json:"name,omitempty" bson:"name,omitempty"`
// 	Statement   string             `json:"statement,omitempty" bson:"statement,omitempty"`
// 	Description string             `json:"description,omitempty" bson:"description,omitempty"`
// 	Input       string             `json:"input,omitempty" bson:"input,omitempty"`
// 	Output      string             `json:"output,omitempty" bson:"output,omitempty"`
// 	Image       string             `json:"image,omitempty" bson:"image,omitempty"`
// }
