package main

import (
	"fmt"
	"strconv"
)

func main() {
	// fizzbuzz()
	printPrimes(20)
}

// Loops
func bulkSend(numMessages int) float64 {
	totalPrice := 0.0
	for i := 0; i < numMessages; i++ {
		fee := float64(i) * 0.01
		totalPrice += 1.0 + fee
	}

	return totalPrice
}

// Ommitting conditions
func maxMessages(thresh int) int {
	var totalCost int
	for i := 0; ; i++ {
		messageCost := 100 + i
		totalCost += messageCost
		if totalCost > thresh {
			return i
		}
	}
}

// No while loop
func whileLoop() {
	plantHeight := 1
	for plantHeight < 5 {
		fmt.Println("still growing! current height:", plantHeight)
		plantHeight++
	}
}

// Exercise
func fizzbuzz() {
	for i := 1; i < 101; i++ {
		var textToPrint string
		isThreeMultiple := i%3 == 0
		isFiveMultiple := i%5 == 0
		if isThreeMultiple {
			textToPrint += "fizz"
		}
		if isFiveMultiple {
			textToPrint += "buzz"
		}
		if textToPrint == "" {
			textToPrint = strconv.Itoa(i)
		}

		fmt.Println(textToPrint)
	}
}

// Exercise 2
func printPrimes(max int) {
	for n := 2; n < max; n++ {
		isTwo := n == 2
		if isTwo {
			fmt.Println(n)
			continue
		}

		isEven := n%2 == 0
		if isEven {
			continue
		}

		var isPrime bool = true
		// Checking up to sqrt(n) reduces the number of iterations from O(n) to O(√n).
		for i := 3; i*i <= n; i += 2 {
			iIsMultipledN := n%i == 0
			if iIsMultipledN {
				isPrime = false
				break
			}
		}

		if isPrime {
			fmt.Println(n)
		}
	}
}
