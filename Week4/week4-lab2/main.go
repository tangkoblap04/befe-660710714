package main

import (
	"fmt"
)

//var email string = "nuttachot@hotmail.com"

func main(){
	var name string = "Thanin Tangkoblap"
	var age int = 10

	email := "tangkoblap_t@silpakorn.edu"
	gpa := 3.63

	firstname, lastname := "Thanin", "Tangkoblap"

	fmt.Printf(
		"Full Name : %s , age : %d , Email : %s , gpa : %.2f\nFirst name : %s Lastname : %s\n",
		name,age,email,gpa,firstname,lastname)
}