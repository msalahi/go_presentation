package main

import (
	"fmt"
)

func printAdd(num1 int, num2 int) {
	sum := num1 + num2
	fmt.Printf("%d + %d = %d\n", num1, num2, sum)
}

func main() {
	var a int
	var b = 2

	a = 2
	printAdd(a, b)
}
