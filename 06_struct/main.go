package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

type Person struct {
	name string
}

func main() {
	fmt.Println(Vertex{1, 2})
	fmt.Println(Person{"siwanon"})
}
