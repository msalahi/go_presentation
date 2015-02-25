package main

import (
	"fmt"
	"math/rand"
)

func search(words []string, searchWord string){
	nwords := len(searchWord)
	for i:=0; i<len(searchWord)
}

func main() {
	fmt.Println("For loop funs!")
	for i := 0; i < 2; i++ {
		fmt.Printf("%d\n", i)
	}

	fmt.Println("While loop funs!")
	limit := 2
	for limit > 0 {
		fmt.Printf("%d\n", limit)
		limit -= 1
	}

	fmt.Println("For loop MADNESS")

	var i int
	for i = 1; i%2 != 0; i = rand.Int() {
		fmt.Printf("%d\n", i)
	}
	fmt.Printf("%d is even\n", i)
}
