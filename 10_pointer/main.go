package main

import "fmt"

type Coffee struct {
	sugar int
	cream int
}

func main() {
	deRef()
	// c := Coffee{}
	// passByValue(c)
	// passByReference(&c)
	// fmt.Printf("Sugar: %d, Cream: %d\n", c.sugar, c.cream)
}

func deRef() {
	a := 10
	p := &a         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 20         // set i through the pointer
	fmt.Println(a)
}

func passByValue(c Coffee) {
	c.sugar = 2
	c.cream = 3
	fmt.Printf("Sugar: %d, Cream: %d\n", c.sugar, c.cream)
}

func passByReference(c *Coffee) {
	c.sugar = 2
	c.cream = 3
}
