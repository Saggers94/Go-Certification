package main

import (
	"bufio"
	"fmt"
	"myapp/doctor"
	"os"
	"strings"
)

// Main function cannot take any parameter
func main() {
	// Go won't allow a variable declaration if it isn't being used
	// var whatToSay string
	// whatToSay = "Hello World again!"
	
	//":=" will declare the variable and figure out what type to assign to the variable with a value
	// whatToSay := "Hello World again!"

	// sayHelloWorld(whatToSay)
	
	// "fmt" is a package available in the go standard library
	// fmt.Println("Hello World!")
	// fmt.Print("This is some text!")
	// fmt.Print("This is some more text!")

	// Reading input from the terminal through "bufio" package and "os.Stdin"
	reader := bufio.NewReader(os.Stdin)

	// var whatToSay string
	whatToSay := doctor.Intro()

	fmt.Println(whatToSay)

	for {
		fmt.Print("-> ")

		// Below "\n" means 'Enter key' on the keyword
		userInput, _ := reader.ReadString('\n')
		
		// "\r\n" is a returning with an Enter key which would be added to the input when
		// we return the value. and "-1" would replace it whereever you will find it in
		// the string
		// "\r\n" for windows and "\n" for mac and linux
		userInput = strings.Replace(userInput, "\r\n","",-1)
		userInput = strings.Replace(userInput,"\n","",-1)
		// fmt.Println(userInput)

		fmt.Println(doctor.Response(userInput))
		if userInput == "quit"{
			break
		}

	}

}

// func sayHelloWorld(whatToSay string){
// 	fmt.Println(whatToSay)
// }


// To build the executable in go
// this is for Mac and linux
// go build -o eliza main.go

// For windows
// go build -o eliza.exe main.go