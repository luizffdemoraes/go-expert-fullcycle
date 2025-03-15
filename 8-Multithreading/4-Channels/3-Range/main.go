package main

import (
	"fmt"
	"sync"
)

/**
🔍 O que esse código faz?
Este código implementa um sistema de comunicação concorrente entre duas goroutines usando canais (chan) 
e um WaitGroup (wg) para sincronização.

1️⃣ A main goroutine cria um canal (ch)
Esse canal permite a comunicação entre as goroutines publish e reader.

2️⃣ A função publish(ch) é iniciada em uma goroutine
Ela envia 10 números (0 a 9) para o canal ch.
Após isso, fecha o canal com close(ch), indicando que não haverá mais envios.

3️⃣ A função reader(ch, &wg) é iniciada em outra goroutine
Ela lê os valores do canal e imprime cada um.
A cada leitura, chama wg.Done() para indicar que um item foi processado.

4️⃣ A main goroutine aguarda (wg.Wait())

Ela espera até que todas as 10 chamadas de wg.Done() ocorram, garantindo que todas as mensagens foram processadas antes de encerrar o programa.
**/
// Thread 1
func main() {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(10)

	go publish(ch)
	go reader(ch, &wg)

	wg.Wait()
}

func reader(ch chan int, wg *sync.WaitGroup) {
	for x := range ch {
		fmt.Printf("Received %d\n", x)
		wg.Done()
	}
}

func publish(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}
