package main

import "fmt"

var a int //Makes "a" an int variable

var (
	b bool
	c float32
	d string
)

func main() {
	a = 42                  //Assigns a single value
	b, c = true, 32.0       //Assign multiple values
	d = "string"            //Strings must contain double quotes " "
	fmt.Println(a, b, c, d) // OUTPUT -> 42 true 32 string
}
