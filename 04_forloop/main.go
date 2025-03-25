package main

import "fmt"

func main() {
	loop()
	loopWhile()
	loopRange()
	loopInfinity()
}

func loop() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func loopWhile() {
	sum := 1
	for sum < 10 {
		sum += sum
	}
	fmt.Println(sum)
}

func loopRange() {
	arr := []string{"T", "u", "r", "b", "o", "w"}
	var name string

	for _, v := range arr {
		name += v
	}

	fmt.Println(name)
}

func loopInfinity() {
	for {
		fmt.Println("Infinity Loop")
	}
}
