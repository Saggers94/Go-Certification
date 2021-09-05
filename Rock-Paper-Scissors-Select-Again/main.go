package main

import "myapp/game"

// Select can only work with channels
// Channels are used for communicating with packages


func main() {
	displayChan := make(chan string)
	roundChan := make(chan int)

	game := game.Game{
		DisplayChan: displayChan,
		RoundChan: roundChan,
		Round: game.Round{
			RoundNumber: 0,
			PlayerScore: 0,
			ComputerScore: 0,
		},
	}

	go game.Rounds()
	game.ClearScreen()
	game.PrintIntro()

	// *** added the for loop
	for {
		game.RoundChan <- 1
		// The below syntax would wait for something to happen
		<- game.RoundChan

		if game.Round.RoundNumber > 3{
			break
		}
		if !game.PlayRound(){
			game.RoundChan <- -1
			<-game.RoundChan
		}
	}

	
	game.PrintSummary()
}



