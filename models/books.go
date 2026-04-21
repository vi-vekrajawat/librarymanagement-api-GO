package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Books struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	BookName        string             `bson:"book_name" json:"book_name"`
	BookAuthor      string             `bson:"book_author" json:"book_author"`
	BookPrice       float64            `bson:"book_price" json:"book_price"`
	BookDescription string             `bson:"book_description" json:"book_description"`
}