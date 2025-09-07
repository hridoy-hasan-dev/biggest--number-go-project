package main

import "fmt"

func main() {
	a := 10
	b := 25
	c := 7

	biggest := a
	if b > biggest {
		biggest = b
	}
	if c > biggest {
		biggest = c
	}

	fmt.Printf("The biggest number is: %d\n", biggest)
}
