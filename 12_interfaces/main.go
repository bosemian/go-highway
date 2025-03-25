package main

import (
	"fmt"
)

func main() {
	// c := cart.Cart{"siwanon", 999}
	// c.Pay(1000.00)

	// cp := computer.Computer{Storage: 500}
	// fmt.Printf("My computer storage is %f\n", cp.Storage)
	// cp.AddStorage()
	// fmt.Printf("My computer storage is %f\n", cp.Storage)

	// clp := phone.CellPhone{"Nokia", 200}
	// fmt.Printf("My cellphone storage is %f\n", clp.Storage)
	// clp.AddStorage()
	// fmt.Printf("My cellphone storage is %f\n", clp.Storage)

	// emptyInterface()

	typeAssert()
}

func emptyInterface() {
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func typeAssert() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	// f = i.(float64) // panic
	// fmt.Println(f)
}
