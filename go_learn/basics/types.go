package main

import "fmt"

func main() {
	/* User specified types */
	const a rune = 12         //rune = int32 can use both to initialize var/const
	const b float32 = 20.5    // 32-bit float
	var c complex128 = 1 + 4i //128-bit complex number
	var d uint16 = 14         //16-bit unsigned integer

	/* Default types */
	n := 42              //int
	pi := 3.14           //float64
	x, y := true, false  //bool
	z := "Go is owesome" //string

	fmt.Printf("User-specified types: \n %T %T %T %T\n", a, b, c, d)
	fmt.Printf("Default types: \n %T %T %T %T %T \n", n, pi, x, y, z) //OUTPUT -> int float64 bool bool string
	//int can stand for either int32 or int64.
	// %T are called verbs. Type of the passed variable.
	/* There are plenty of verbs in Go.
	   | %d -> decimals
	   | %s -> strings
	   | %f -> floats
	   | %t -> booleans
	   | %v -> natural type
	*/
}
