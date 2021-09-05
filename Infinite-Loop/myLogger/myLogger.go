package myLogger

import "log"

func ListenForLog(ch chan string) {
	// This is a go routine that is running in the background always like service
	for {
		msg := <-ch
		log.Println(msg)
	}
}