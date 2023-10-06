// package utils

import (
	"github.com/joho/godotenv"
	//"gorm.io/gorm"
	//"gorm.io/driver/postgres"
	"onlymachiavelli/web-service-gin/src/helper"
	"os"
	"strconv"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


// func Connect() (*gorm.DB, error) {
// 	err := godotenv.Load()
// 	if err != nil {
// 		helper.ErrorHandler(err)
// 	}
// 	HOST := os.Getenv("HOST")
// 	USER := os.Getenv("USER")
// 	PORTStr := os.Getenv("DBPORT")
// 	PORT, err := strconv.Atoi(PORTStr)
// 	if err != nil {
// 		fmt.Println("Error converting PORT to integer, so the port by default is: 3000")
// 		PORT = 3000
// 	}
// 	DBNAME := os.Getenv("DBNAME")
// 	PASS := os.Getenv("PASSWORD")

// 	conn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d", HOST, USER, PASS, DBNAME, PORT)
// 	fmt.Println(conn)
// 	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})

// 	return db, err
// }



//connect to mongodb
func Connect() (*mongo.Client, error) {

	//get the URI
	err := godotenv.Load()
	if err != nil {
		helper.ErrorHandler(err)
	}
	URI := os.Getenv("URI")
	clientOptions := options.Client().ApplyURI(URI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}
	//check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	fmt.Println("Connected to MongoDB!")
	return client, nil	
	
}