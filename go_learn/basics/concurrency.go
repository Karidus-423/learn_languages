package main

import (
	"fmt"
)

func main() {
	c := make(chan int) //Makes a channel to pass int
	for i := 0; i < 5; i++ {
		go cookingGopher(i, c) //Start a goroutine(thread)
	}

	for i := 0; i < 5; i++ {
		gopherID := <-c //Receive a value from a channel

		fmt.Println("gohpher", gopherID, "finished the dish")
	}
}

//Channel can be used as argument

func cookingGopher(id int, c chan int) {
	fmt.Println("gopher", id, "started cooking")
	c <- id // Sends the value to back to main
}
