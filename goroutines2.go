package main

import (
	"fmt"
	"time"
)

func readStrings(strings <-chan string) {
	duration, _ := time.ParseDuration("800ms")
	for {
		time.Sleep(duration)
		fmt.Println("-> " + <-strings)
	}
}

func writeStrings(strings chan<- string) {
	for {
		fmt.Println("<- cat")
		strings <- "cat"
		fmt.Println("<- dog")
		strings <- "dog"
	}
}

func main() {
	strings := make(chan string, 10)
	go writeStrings(strings)
	go readStrings(strings)

	var input string
	fmt.Scanln(&input)
}
