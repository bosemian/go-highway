package main

import "fmt"

func main() {
	sliceGrow()
}

func basic() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)

	b := []byte{72, 101, 108, 108, 111, 32, 87, 111, 114, 108, 100}
	fmt.Println(string(b))

	h := "สวัสดีครับ"

	for i, r := range h {
		fmt.Printf("Index: %d, byte: %d (Unicode: %U)\n", i, r, r)
	}
}

func sliceGrow() {
	sl := make([]int, 2)
	printSlice((sl))
	sl = append(sl, 1)
	printSlice((sl))
	sl = append(sl, 2)
	printSlice((sl))

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
