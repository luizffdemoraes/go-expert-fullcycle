package main

import (
	"fmt"
	"sync"
	"time"
)

func task(name string, wg *sync.WaitGroup) {
	defer wg.Done() // Garante que Done será chamado ao final

	for i := 0; i < 10; i++ {
		fmt.Printf("%d: Task %s is running\n", i, name)
		time.Sleep(1 * time.Second)
	}
}

/*
*
📌 Como o Código Usa Concorrência?
O código cria três goroutines (task("A"), task("B") e uma anônima).
Essas goroutines compartilham o mesmo espaço de memória e são executadas de forma concorrente pelo Go Scheduler.
A função waitGroup.Wait() garante que a main goroutine aguarde a finalização de todas as outras antes de encerrar.

✅ Concorrência ocorre porque várias goroutines são executadas simultaneamente, 
mesmo que em um único núcleo de CPU. O Go Scheduler alterna a execução entre elas, garantindo que cada uma tenha tempo de processamento.

*
*/
func main() {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(10)
	// Thread 2
	go task("A", &waitGroup)
	// Thread 3
	go task("B", &waitGroup)
	// Thread 4
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("%d: Task %s is running\n", i, "anonymous")
			time.Sleep(1 * time.Second)
			waitGroup.Done()
		}
	}()


	// for _, name := range tasks {
	// 	wg.Add(1) // Adiciona dinamicamente antes de criar a goroutine
	// 	go task(name, &wg)
	// }

	waitGroup.Wait()
}
