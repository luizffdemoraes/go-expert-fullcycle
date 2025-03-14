package main

import "fmt"

// Thread 1
func main() {
	canal := make(chan string)

	// Thread 2
	go func() {
		canal <- "Olá Mundo!"
	}()

	// Thread 1
	msg := <-canal
	fmt.Println(msg)
}
