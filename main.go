package main
import (
	"fmt"
	//"helper"
	//"utils"
	"github.com/joho/godotenv"	
	"os"

)


func main () {
	fmt.Println("hello")
	err := godotenv.Load()
	if err != nil {
		fmt.Println("error loading .env file")
	}
	//printing TST data from dotenv 
	fmt.Println("TST: ", os.Getenv("TST"))
}