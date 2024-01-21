package main

import "fmt"

func main() {
	languages := [9]string{
		"C", "Lisp", "C++", "Java", "Python",
		"JavaScript", "Ruby", "Go", "Rust", //Slices need a trailing comma.
	}

	/* Define Slices */
	classics := languages[0:3]  //or languages [:3]
	modern := make([]string, 4) // len(modern) = 4
	modern = languages[3:7]     // include 3 exclude 7
	new := languages[7:9]       // or [7:]

	fmt.Printf("classic languages: %v\n", classics) // classic languages: [C Lisp C++]
	fmt.Printf("modern languages: %v\n", modern)    // modern languages: [Java Python JavaScript Ruby]
	fmt.Printf("new languages: %v\n", new)          // new languages: [Go Rust]
}
