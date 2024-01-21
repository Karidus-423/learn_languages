package main

import "fmt"

func main() {
	var address *int  //Create a pointer to an int named address
	number := 42      // int value
	address = &number //address stores the memory address of number
	value := *address //dereferencing the value

	fmt.Printf("address: %v\n", address)
	fmt.Printf("value: %v\n", value)
}

// & binds a pointer to a value
// * dereferences the pointer. To access the thing to which the pointer points/
