package main

import "fmt"

func safe() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	panic("Something bad happened")
}

func main() {
	safe()
	fmt.Println("Program continues...")

	var arr = []int{1, 2, 3}
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	fmt.Println(arr[5])
}
