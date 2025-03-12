package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: Task %s is running\n", i, name)
		time.Sleep(1 * time.Second)
	}
}

/*
*
Este exemplo demonstra como criar e executar tarefas concorrentes em Go usando **goroutines**
Thread 1
*
*/
func main() {
	// Thread 2
	go task("A")
	// Thread 3
	go task("B")
	// Thread 4
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("%d: Task %s is running\n", i, "anonymous")
			time.Sleep(1 * time.Second)
		}
	}()
	// Sair
	time.Sleep(15 * time.Second)
}
