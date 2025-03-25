package main

import (
	"fmt"
)

func sayHello() {
	fmt.Println("Hello from goroutine")
}

func main() {
	go sayHello()
	fmt.Println("Hello from main")
	// time.Sleep(time.Second) // Give goroutine time to finish
}
