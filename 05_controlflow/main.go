package main

import (
	"fmt"
	"time"
)

func main() {
	if true {
		fmt.Println("true")
	}

	num := 10

	if num%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	if sum := (num + 1) % 2; sum != 0 {
		fmt.Println("Sum is Odd")
	}

	if num == 0 {
		fmt.Println("Zero")
	} else if num <= 10 {
		fmt.Println("Less than or equal to 10")
	} else {
		fmt.Println("Greater than 10")
	}

	switch num {
	case 0:
		fmt.Println("Zero")
	case 1, 3, 5, 7, 9:
		fmt.Println("Odd")
	default:
		fmt.Println("Even")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
