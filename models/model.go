package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// House Database Models, Possibility of using multiple files
type User struct {
	ID       primitive.ObjectID `bson:"_id,  omitempty" json:"-"`
	Email    string
	Password string
}

// ERRORs
// for Returning Response, Define all Errors, then Errors can set Message directly
// Switch case containing all types of errors, if it matches with gotten from recover() {Test}
