package main

import (
	"fmt"
)

// Verify if num is a prime number. If yes, returns true, otherwise returns false
func isPrime(num int) bool {
	if num == 0 || num == 1 {
		return false
	}

	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var num int

	fmt.Print("Digit a number: ")
	_, err := fmt.Scanln(&num)
	check(err)

	if isPrime(num) {
		fmt.Printf("%d is prime.\n", num)
	} else {
		fmt.Printf("%d não é primo.\n", num)
	}
}
