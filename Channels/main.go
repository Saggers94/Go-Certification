package main

import (
	"fmt"

	"github.com/eiannone/keyboard"
)

// You can only pass one type to a "Channel" which is the below case is "rune"
// "rune" is a single character that is being used to build a string
var keyPressChan chan rune

func main() {
	// this is running at the same time as the main function is running
	// we can have a go routine running concurrently and that's what the keyword "go"
	// is doing
	// we communicate to the "go" routine with channels
	// go doSomething("Hello, World")
	// fmt.Println("This is another message.")

	// for {
		// Do nothing
	// }

	keyPressChan = make(chan rune)

	go listenForAKeyPress()

	fmt.Println("Press any key or q to quit")

	_ = keyboard.Open()

	defer func(){
		keyboard.Close()
	}()

	for {
		char, _, _  := keyboard.GetSingleKey() 
		if char == 'q' || char == 'Q' {
			break
		}

		keyPressChan <- char
	}

}

func doSomething(s string) {
	until := 0
	for {
		fmt.Println("s is", s)
		until = until + 1
		if until >= 5 {
			break
		}
	}
}

func listenForAKeyPress(){
	for {
		key := <-keyPressChan
		fmt.Println("You pressed", string(key))
	}
}

// routines are kind of like services that runs in the background and through channels
// we can inject the data into the routines and than it would execute and give us what we want
// for example, we can have the "email" service running through routine and than have a
// channel to use that routine to send out an email