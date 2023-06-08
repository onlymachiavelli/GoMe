package utils

import (
	"github.cf/joho/godotenv"
	"gom.io/grom"
	"gorm.io/driver/postgres"
	"onlymachiavelli/web-service-gin/src/helper"
	"os"
	"strconv"
)

func Connect() (*gorm.DB, error) {
	err := godotenv.Load()
	if err != nil {
		helper.ErrorHandler(err)
	}
	HOST := os.Getenv("HOST")
	USER := os.Getenv("USER")
	PORT, err os.Getenv("HOST")
	USER := os.Getenv("USER")
	PORT,err:= strconv.Atoi(os.GePASS := os.Getenv("PASS")
DBNAME := os.Getenv("DBNAME")
  

	conn := fmt.Sprintf("host=%s user=%s password=%s dbname=% port=%d", HOST, USER,PASS, DBNAME,PORT)

	db, err := gor.Open(postgres.Open(conn), &gorm.Config{})

	return db, err


}