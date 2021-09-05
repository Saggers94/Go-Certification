package main

import (
	"bufio"
	"fmt"
	"myapp/myLogger"
	"os"
	"strings"
	"time"
)

func main() {
	

	reader := bufio.NewReader(os.Stdin)

	// this channel would be used to inject the input inside the routine
	ch := make(chan string)

	// Here we are defining the go routine with the "go" keyword
	go myLogger.ListenForLog(ch)

	fmt.Println("Enter something")
	
	// read input from the user for 5 times and write it to a log
	for i:=0; i<5; i++ {
		fmt.Print("-> ")
		input, _ := reader.ReadString('\n')
		input = strings.Replace(input,"\r\n","",-1)
		ch <- input
		time.Sleep(time.Second)
	}
}