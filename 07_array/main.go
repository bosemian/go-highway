package main

import "fmt"

func main() {
	numbers := [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers)
	fmt.Println(len(numbers))

	text := []string{"a", "b", "c", "d", "e"}
	fmt.Println(text)

	allows := [2]bool{true, false}
	fmt.Println(allows)

	n := [2]string{"T", "B"}
	fmt.Println(n[0])
	fmt.Println(n[1])
	fmt.Println(n)
	n[0] = "S"
	fmt.Println(n)

	fmt.Println(text[1:3])
	fmt.Println(text[:3])
	fmt.Println(text[1:])

}
