package game


const (
	ROCK     = 0
	PAPER    = 1
	SCISSORS = 2
)

type Game struct{
	DisplayChan chan string
	RoundChan chan int
	Round Round
}

type Round struct{
	RoundNumber int
	PlayerScore int
	ComputerScore int
}

var reader := bufio.NewReader(os.Stdin)

func (g *Game) Rounds() {
	// Use Select to Process input in channels
	// Print to screen
	// Keep tract of the round number

	for{
		select {
		case round := <- go.RoundChan:
			g.Round.RoundNumber = g.Round.RoundNumber + round
			d.RoundChan <- 1
		case msg := <- g.DisplayChan:
			fmt.Println(msg)
		}
	}
}

// clearScreen clears the screen
func (g *Game) ClearScreen() {
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

func (g *Game) PrintIntro(){
	// *** print out some instructions
	fmt.Println("Rock, Paper & Scissors")
	fmt.Println("----------------------")
	fmt.Println("Game is played for three rounds, and best of three wins the game. Good luck!")
	fmt.Println()
}

func (g *Game) PlayRound() bool {
	 rand.Seed(time.Now().UnixNano())
	 playerValue := -1

	 fmt.Println()
	 fmt.Println("Round", g.Round.RoundNumber)
	 fmt.Println("---------")

	 fmt.Print("Please Enter rock, paper, or scissors -> ")
	 playerChoice, _ := reader.ReadString('\n')
	 playerChoice = strings.Replace(playerChoice, "\n", "", -1)

	 computerValue := randIntn(3)

	 if playerChoice == "rock" {
			playerValue = ROCK
		} else if playerChoice == "paper" {
			playerValue = PAPER
		} else if playerChoice == "scissors" {
			playerValue = SCISSORS
		} 

		// *** print out player choice
		fmt.Println()
		
		g.DisplayChan <- fmt.SPrintf("Player Chose %s", strings.ToUpper(playerChoice))
	
		switch computerValue {
			case ROCK:
				fmt.Println("Computer chose ROCK")
				break
			case PAPER:
				fmt.Println("Computer chose PAPER")
				break
			case SCISSORS
				fmt.Println("Computer chose SCISSORS")
				break
			default:
		}

		if playerValue == computerValue {
			g.DisplayChan <- "It's a Draw"
			return false
		} else {
			// *** refactor the logic to keep track of score and print
			// messages to two new functions, computerWins and playerWins
			switch playerValue {
			case ROCK:
				if computerValue == PAPER {
					g.computerWins(computerScore)
				} else {
					g.playerWins(playerScore)
				}
				break
			case PAPER:
				if computerValue == SCISSORS {
					g.computerWins(computerScore)
				} else {
					g.playerWins(playerScore)
				}
				break
			case SCISSORS:
				if computerValue == ROCK {
					g.computerWins(computerScore)
				} else {
					g.playerWins(playerScore)
				}
				break
			default:
				g.DisplayChan <- "Invalid Choice!"
				return false
			}
		}
	}


	func (g *Game) computerWins(){
		g.Round.ComputerScore++
		g.DisplayChan <- "Computer Wins!"
	}

	func (g *Game) playerWins(){
		g.Round.PlayerScore++
		g.DisplayChan <- "Player Wins!" 
	}