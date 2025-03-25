package main

import (
	"fmt"
	"strconv"
)

func main() {
	doSomething()
	f := func() {
		fmt.Print("func declare and assign\n")
	}
	f()

	func() {
		fmt.Print("func declare and execute immediately\n")
	}()
	rs := concat("Turbow")
	fmt.Println(rs)
	_, err := convert(rs)
	if err != nil {
		fmt.Println("Error converting string to int:", err)
		return
	}
}

func doSomething() {
	fmt.Println("Do something !!")

}

func concat(n string) string {
	return n + " is a Go Developer"
}

func convert(s string) (int, error) {
	return strconv.Atoi(s)
}
