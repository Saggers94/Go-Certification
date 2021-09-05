package main

import (
	"fmt"
	"myapp/packageone"
)

// this is the package level variable
var one = "Love One"

func main() {
	
	// The block level variable has the presidence
	// The below is the shadow variable, don't do this
	// don't give the same name
	// var one = "This is a block level variable"
	var somethingElse = "This is a block level variable"
	fmt.Println(somethingElse)
	
	myFunc()

	newString := packageone.PublicVar
	fmt.Println("From Packageone", newString)
	
	packageone.Exported()

}

func myFunc(){
	// this "var" variables are blocked variables
	// var one = "The number one"
	fmt.Println(one)
}