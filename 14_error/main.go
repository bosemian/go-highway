package main

import (
	"errors"
	"fmt"
	"os"
)

var AnyError = errors.New("Something went wrong")

type AppError struct {
	Code int
	Msg  string
}

func (e AppError) Error() string {
	return fmt.Sprintf("Code %d: %s", e.Code, e.Msg)
}

func doSomething() error {
	return AppError{Code: 404, Msg: "resource not found"}
}

func readFile(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return AppError{Code: 404, Msg: "file not found"}
	}
	fmt.Println(string(data))
	return nil
}

func main() {
	// result, err := divide(10, 0)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	// fmt.Println(AnyError)
	// } else {
	// 	fmt.Println("Result:", result)
	// }

	// if err := doSomething(); err != nil {
	// 	fmt.Println("Error:", err)
	// }

	if err := readFile("example.txt"); err != nil {
		fmt.Println("Error:", err)
	}
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}
