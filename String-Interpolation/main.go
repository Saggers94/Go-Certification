package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/eiannone/keyboard"
)

var reader *bufio.Reader

// Defining our own variable type that is based on the Primitive types of GO
type User struct {
	UserName string
	Age int
	FavouriteNumber float64
	// Even though we have defined the "OwnsADog" field here, Go doesn't care if we are
	// not using it
	OwnsADog bool
}

func main() {
	reader = bufio.NewReader(os.Stdin)
	strOwnADog := "You don't have a dog."
	
	var user User

	user.UserName = readString("What is your name?")
	user.Age = readInt("How old are you?")
	user.FavouriteNumber = readFloat("What is your Favourite number?")
	user.OwnsADog = readBool("Do you have a dog? y/n")

	if user.OwnsADog {
		strOwnADog = "You Own a dog"
	}
	
	// '+' concatination
	// fmt.Println("Your name is" + userName + ". You are" + age + "years old.")
	// fmt.Println("Your name is" , userName , ".You are" , age , "years old.")

	// This is a better way than the concatination
	// fmt.Println(fmt.Sprintf("Your name is %s. You are %d years old.",userName, age))

	// This is a string Interpolation
	// This is a better way than the above way. This is more efficient
	fmt.Printf("Your name is %s. You are %d years old. Your favourite Number is %.2f. %s \n", user.UserName, user.Age, user.FavouriteNumber, strOwnADog)
}

func prompt(){
	fmt.Print("-> ")
}

func readString(s string) string{
	for{
		fmt.Println(s)
		prompt()

		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput,"\r\n","",-1)
		userInput = strings.Replace(userInput,"\n","",-1)

		if userInput == "" {
			fmt.Println("Please Enter a value")
		}else{
			return userInput
		}
	}
}

func readInt(s string) int{
	for{
		fmt.Println(s)
		prompt()

		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput,"\r\n","",-1)
		userInput = strings.Replace(userInput,"\n","",-1)

		num, err := strconv.Atoi(userInput)
		if err != nil {
			fmt.Println("Please Enter a whole number")
		}else {
			return num
		}

	}
}

func readFloat(s string) float64{
	for{
		fmt.Println(s)
		prompt()

		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput,"\r\n","",-1)
		userInput = strings.Replace(userInput,"\n","",-1)

		num, err := strconv.ParseFloat(userInput,64)
		if err != nil {
			fmt.Println("Please Enter a number")
		}else {
			return num
		}

	}
}

func readBool (s string) bool{
	for{
		fmt.Println(s)
		prompt()

		char, _ , err := keyboard.GetSingleKey()

		if err!=nil{
			log.Fatal(err)
		}

		if char == 'y' || char == 'Y'{
			return true
		}

		if char == 'n' || char == 'N'{
			return false
		}
	}
}
	
	
