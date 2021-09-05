package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	// Infinite for loop
	// for {

	// }
	//range loop over map or slices
	// for _, x := range myMap {

	// }

	// Three component for loop, it is for set number of times
	for i := 0; i <= 10; i++ {
		fmt.Println("i is" , i)
	}

	i:= 1000
	//execute a loop while i > 100
	for i > 100{
		// get a random number 1 and 1001
		i = rand.Intn(1000) + 1
		fmt.Println("is is",i)
		if i > 100 {
			fmt.Println("i is", i, "so loop keeps going")
		}
	}

	fmt.Println("Got", i, "and broke out of loop")
}