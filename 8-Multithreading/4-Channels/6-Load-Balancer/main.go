package main

import (
	"fmt"
	"time"
)

/*
**
Esse código distribui a carga de trabalho entre múltiplos "workers" concorrentes
✅ Cria um canal data para enviar números inteiros.
✅ Cria 100.000 workers concorrentes via goroutines (worker(i, data)).
✅ Cada worker escuta o canal data e processa valores.
✅ Envia 1.000.000 de tarefas para o canal, que serão processadas pelos workers.
**
*/
func worker(workerId int, data chan int) {
	for x := range data {
		fmt.Printf("Worker %d received %d\n", workerId, x)
		time.Sleep(time.Second)
	}
}

func main() {
	data := make(chan int)
	QtdWorkers := 100000

	for i := 0; i < QtdWorkers; i++ {
		go worker(i, data)
	}

	for i := 0; i < 10000; i++ {
		data <- i
	}
}
