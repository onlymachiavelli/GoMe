package helper


import ("fmt")


func ErrorHandler (err error) {
	if (err != nil) {
		fmt.Println("an error occured ! ")
		fmt.Println(err)
		fmt.Println("exiting ...")
		panic(err)
	}
}