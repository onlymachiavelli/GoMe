
package models 


import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {

	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FirstName string `json:"firstName,omitempty" bson:"firstName,omitempty"`
	LastName string `json:"lastName,omitempty" bson:"lastName,omitempty"`
	Email string `json:"email,omitempty" bson:"email,omitempty"`
	Password string `json:"password,omitempty" bson:"password,omitempty"`
	Picture string `json:"picture,omitempty" bson:"picture,omitempty"`
	CreatedAt string `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt string `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`

}