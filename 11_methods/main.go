package main

import (
	"11_methods/cart"
	"fmt"
	"log"
)

func main() {
	c := cart.Cart{Name: "siwanon", Ccv: 999}
	c.Display()
	c.Change("john", 888)
	fmt.Printf("name: %v, ccv: %v\n", c.Name, c.Ccv)
	p, err := c.TotalPrice()
	if err != nil {
		log.Fatal("Error calculating total price:", err)
		return
	}
	fmt.Println(p)
}
