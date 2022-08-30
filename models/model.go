package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//House Database Models, Possibility of using multiple files
type ModelName struct {
	ID primitive.ObjectID `bson:"_id,  omitempty" json:"-"`
}
