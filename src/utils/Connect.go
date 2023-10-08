package utils 
import (
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"onlymachiavelli/web-service-gin/src/helper"
	"os"
	"context"
	"github.com/joho/godotenv"
)


func Connect() (*mongo.Client, error) {
	err := godotenv.Load(".env")
	if err != nil {
		helper.ErrorHandler(err)
	}
	
	URI := os.Getenv("URI")


	clientOptions := options.Client().ApplyURI(URI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		helper.ErrorHandler(err)
	}

	//pinging 
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		helper.ErrorHandler(err)
	}
	fmt.Println("Connected to MongoDB!")
	return client, nil


}