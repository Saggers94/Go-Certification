package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

const (
	ROCK     = 0
	PAPER    = 1
	SCISSORS = 2
)

func main() {
	rand.Seed(time.Now().UnixNano())
	playerChoice := ""
	playerValue := -1

	

	reader := bufio.NewReader(os.Stdin)

	clearScreen()

	
	
	count := 0
	computerWinningCount := 0
	playerWinningCount := 0

	for count <3{
		count ++
		
		computerValue := rand.Intn(2)

		fmt.Print("Please enter rock, paper, or scissors ->")
		playerChoice, _ = reader.ReadString('\n')
		playerChoice = strings.Replace(playerChoice, "\r\n", "",-1)


		if playerChoice == "rock"{
			playerValue =ROCK
		}else if playerChoice == "paper"{
			playerValue = PAPER
		}else if playerChoice == "scissors"{
			playerValue = SCISSORS
		}else{
			count--
			fmt.Println("Invalid Choice")
			continue
		}

		switch computerValue {
			case ROCK:
				fmt.Println("Computer Chose Rock")
				break
			case PAPER:
				fmt.Println("Computer Chose Paper")
				break
			case SCISSORS:
				fmt.Println("Computer Chose Scissors")
				break
			default:
		}

		fmt.Println()

		if playerValue == computerValue {
			fmt.Println("It's a draw.")
		}else{
			switch playerValue{
			case ROCK:
				if computerValue == PAPER{
					fmt.Println("Computer wins")
					computerWinningCount++
				}else{
					fmt.Println("Player wins")
					playerWinningCount++
				}
				break
			case PAPER:
				if computerValue == SCISSORS{
					fmt.Println("Computer wins")
					computerWinningCount++
				}else{
					fmt.Println("Player wins")
					playerWinningCount++
				}
				break
			case SCISSORS:
				if computerValue == ROCK{
					fmt.Println("Computer wins")
					computerWinningCount++
				}else{
					fmt.Println("Player wins")
					playerWinningCount++
				}
				break
			default:
				fmt.Println("Invalid Choice")
			}
		}


		fmt.Println("Player value is:", playerValue)

		
	}

	if computerWinningCount > playerWinningCount {
		fmt.Println()
		fmt.Println(computerWinningCount,"/3")
		fmt.Println("Computer Wins!")
	}else if computerWinningCount == playerWinningCount {
		fmt.Println("It's a Draw")
	}else{
		fmt.Println()
		fmt.Println(playerWinningCount,"/3")
		fmt.Println("Player Wins!")
	}
		
	// fmt.Println()
	// fmt.Println("Player Chose", playerChoice, "and value is", playerValue)
	// fmt.Println("Computer Chose", computerValue)
}

// clearScreen clears the screen
func clearScreen() {
	if strings.Contains(runtime.GOOS, "windows") {
		// windows
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		// linux or mac
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
