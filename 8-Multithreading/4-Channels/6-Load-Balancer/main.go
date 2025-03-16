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
	for x := range data { // Lê valores do canal até ser fechado
		fmt.Printf("Worker %d received %d\n", workerId, x)
		time.Sleep(time.Second) // Simula processamento demorado
	}
}

func main() {
	// Criação de um canal para enviar tarefas aos workers
	data := make(chan int)
	// Definição do número de workers que serão criados
	QtdWorkers := 100000

	// Criação de 100.000 goroutines workers
	for i := 0; i < QtdWorkers; i++ {
		go worker(i, data)
	}

	// Envio de 10.000 tarefas para o canal
	for i := 0; i < 10000; i++ {
		data <- i
	}
}
