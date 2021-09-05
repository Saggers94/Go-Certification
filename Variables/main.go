package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

// constant can never change
const prompt = "and Press Enter when Ready."

func main() {
	// one way declare and than assign
	// var firstNumber int
	// firstNumber = 2

	// another way declare type and name and assign the value
	// var secondNumber = 5

	// one step variable, declare name, assign a value and let "Go" figure out the type
	// substraction := 7


	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	var firstNumber = rand.Intn(8) + 2
	var secondNumber = rand.Intn(8) + 2
	var substraction = rand.Intn(8) + 2
	var answer = firstNumber * secondNumber - substraction

	gameFunction(firstNumber,secondNumber,substraction,answer)
}

// One function, One purpose
func gameFunction(firstNumber, secondNumber, substraction, answer int ){
	reader := bufio.NewReader(os.Stdin)

	//display a welcome/instructions
	fmt.Println("Guess the number Game")
	fmt.Println("---------------------")
	fmt.Println("")

	fmt.Println("Think of a number between 1 and 10", prompt)
	// In the "ReadString" function argument should be in a single quote

	reader.ReadString('\n')


	//take them through a game

	// Println can take any number of arguments
	fmt.Println("Multiply your number by", firstNumber,prompt)
	reader.ReadString('\n')

	fmt.Println("Now Multiply the result by", secondNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Divide the Result by the number you originally thought of", prompt)
	reader.ReadString('\n')

	fmt.Println("Now substract", substraction, prompt)
	reader.ReadString('\n')

	//i will give them the answer
	fmt.Println("The answer is", answer)
}