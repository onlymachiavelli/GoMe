package connector


import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
	"helper"
	"gorm.io/driver/postgres"
  	"gorm.io/gorm"	

)

func Connect() {

	err := godotenv.Load()
	if (err != nil) {
		helper.errorHandler(err)
	}
	HOST := os.Getenv("HOST")
	USER := os.Getenv("USER")
	PORT:= os.Getenv("DBPORT")
	PASS := os.Getenv("PASS")
	DBNAME := os.Getenv("DBNAME")


	conn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d", HOST, USER,PASS, DBNAME,PORT)

	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})

	return db, err

}