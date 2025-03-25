package main

import (
	"fmt"
	"strconv"
)

const ENDPOINT_URI = "http://192.168.1.1:8080"
const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

var name string

func main() {
	declareVariables()
	// ENDPOINT_URI = "http://192.168.1.1:8081

}

func declareVariables() {
	var isActive bool
	name = "Siwanon"
	isActive = true

	fmt.Printf("Name: %s\nIs Active: %t\n", name, isActive)
}

func shortDeclareInitValue() {
	var isActive bool
	name = "Siwanon"
	isActive = true

	fmt.Printf("Name: %s\nIs Active: %t\n", name, isActive)
}

func zeroValue() {
	var i int
	var f float64
	var b bool
	var s string

	fmt.Printf("i = %v\nf = %v\nb = %v\ns = %q\n", i, f, b, s)
	fmt.Printf("%T %T %T %T\n", i, f, b, s)
}

func typeBoolean() {
	isActive := true

	fmt.Printf("Type: %T\n", isActive)
}

func typeNumber() {
	// Signed integers
	var i int = -42
	var i8 int8 = -128
	var i16 int16 = -32768
	var i32 int32 = -2147483648
	var i64 int64 = -9223372036854775808

	// Unsigned integers
	var u uint = 42
	var u8 uint8 = 255
	var u16 uint16 = 65535
	var u32 uint32 = 4294967295
	var u64 uint64 = 18446744073709551615

	// Pointer-sized integer
	var up uintptr = 123456

	// Floating point numbers
	var f32 float32 = 3.1415927
	var f64 float64 = 3.141592653589793

	// Complex numbers
	var c64 complex64 = complex(1.5, 2.5)
	var c128 complex128 = complex(3.5, 4.5)

	// Aliases
	var b byte = 255 // alias for uint8
	var r rune = 'A' // alias for int32

	// Print values
	fmt.Println("Signed Integers:", i, i8, i16, i32, i64)
	fmt.Println("Unsigned Integers:", u, u8, u16, u32, u64)
	fmt.Println("Pointer-sized Integer:", up)
	fmt.Println("Floating-point Numbers:", f32, f64)
	fmt.Println("Complex Numbers:", c64, c128)
	fmt.Println("Aliases (byte and rune):", b, r)
}

func convertType() {
	var x, y int = 3, 4
	var f float64 = float64(x*x + y*y)
	fmt.Println(x, y, f)

	hundred := "100"

	num, err := strconv.Atoi(hundred)
	if err != nil {
		fmt.Println("Error converting string to int:", err)
		return
	}
	fmt.Println("Converted to int:", num)

	str := strconv.Itoa(num)
	fmt.Println("Converted to string:", str)
}
