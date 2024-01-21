package main

import "fmt"

func main() {
	/* Define a map containing the release year of several languages. */
	firstRelease := map[string]int{
		"C": 1972, "C++": 1985, "Java": 1996,
		"Python": 1991, "JavaScript": 1996, "Go": 2012,
	}

	/* Loop through each entry and output the (k)name and release (v)year*/
	for k, v := range firstRelease {
		fmt.Printf("%s wass first released in %d\n", k, v)
	}
}
