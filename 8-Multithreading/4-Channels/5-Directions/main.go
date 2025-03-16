package main

import "fmt"

// Thread 1
func main() {
	hello := make(chan string)
	go recebe("Hello", hello)
	ler(hello)
}

// O canal tem o tipo chan<- string, o que significa que só pode enviar dados.
func recebe(nome string, hello chan<- string) {
	hello <- nome
}

// O canal tem o tipo <-chan string, o que significa que só pode receber dados.
func ler(data <-chan string) {
	fmt.Println(<-data)
}
