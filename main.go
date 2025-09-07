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

	for i := 1; i <= 100; i++ {
		fmt.Println(i)
	}

	// Easy problem: Check if a number is even or odd
	num := 15
	if num%2 == 0 {
		fmt.Printf("%d is even\n", num)
	} else {
		fmt.Printf("%d is odd\n", num)
	}
}
