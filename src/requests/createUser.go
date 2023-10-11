package requests 



import (
	//gin
	"github.com/gin-gonic/gin"
	//mongo
	"go.mongodb.org/mongo-driver/bson"
)

//CreateUserRequest is the struct that will be used to receive the request body
type CreateUserRequest struct {
	FirstName string `json:"firstName,omitempty" bson:"firstName,omitempty"`
	LastName string `json:"lastName,omitempty" bson:"lastName,omitempty"`
	Email string `json:"email,omitempty" bson:"email,omitempty"`
	Password string `json:"password,omitempty" bson:"password,omitempty"`
	Picture string `json:"picture,omitempty" bson:"picture,omitempty"`
}
