package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/eiannone/keyboard"
)

func main() {
	err := keyboard.Open()

	if err != nil {
		log.Fatal(err)
	} 

	// "defer" keyword before the function would make the function run 
	// after the execution of all of the current functions lines
	defer func(){
		_ = keyboard.Close()
	}()

	coffees := make(map[int]string)
	coffees[1] = "Cappucino"
	coffees[2] = "Latte"
	coffees[3] = "Americano"
	coffees[4] = "Mocha"
	coffees[5] = "Macchiato"
	coffees[6] = "Espresso"

	// fmt.Println("Press Any Key on the Keyboard. Press Esc to quit.")
	fmt.Println("MENU")
	fmt.Println("----")
	fmt.Println("1 - Cappucino")
	fmt.Println("2 - Latte")
	fmt.Println("3 - Americano")
	fmt.Println("4 - Mocha")
	fmt.Println("5 - Macchiato")
	fmt.Println("6 - Espresso")
	fmt.Println("q - Quit the Program")

	char := ' '

	for char != 'q' && char!='Q'{
		// "rune" is a single character its a lower level than the type String
		// below the "char" variable is a type of "rune"
		// char, key, err := keyboard.GetSingleKey()
		
		// "_" is being used to discard the result
		char, _, err = keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}

		// if char == 'q' || char == 'Q' {
		// 	break
		// }

		i, _ := strconv.Atoi(string(char))
		// "%q" is a placeholder for rune type which is similar to character
		// "%d" is a placeholder for int
		// "%s" is a placeholder for string
		// t := fmt.Sprintf("You chose %q",char)
		
		// Chcking the Entry "i" exists in the "coffees"
		_,ok := coffees[i]

		if ok {
			fmt.Println(fmt.Sprintf("You chose %s",coffees[i]))
		}
		

		// if key != 0 {
		// 	fmt.Println("You Pressed", char, key)
		// }else{
		// 	fmt.Println("You Pressed", char)
		// }

		// if key == keyboard.KeyEsc {
		// 	break
		// }

	}

	fmt.Println("Program Exiting.")
	// reader := bufio.NewReader(os.Stdin)

	// for {
	// 	fmt.Print("-> ")
		// '_' is a second thing that is being returned from the ReadString function
		// in this case we are descarding it by not saving it in any variable and rather
		// than saving it in the "_"
	// 	userInput, _ := reader.ReadString('\n') 

	// 	userInput = strings.Replace(userInput,"\r\n","",-1)

	// 	if userInput == "quit"{
	// 		break
	// 	}else{
	// 		fmt.Println(userInput)
	// 	}
	// }

	// Do while loop
	for {
		//do some work

		//Check a condition and break
	}

	for ok :=true; ok; ok = char != 'q' && char != 'Q'{
		
	}
}