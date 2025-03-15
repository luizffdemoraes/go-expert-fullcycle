package main

/**
🔥 Deadlock no Código
O código abaixo gera um deadlock porque a goroutine principal fica bloqueada indefinidamente,
 esperando um valor do canal forever, mas nenhuma outra goroutine envia um valor para ele.

**/
func main() {
	forever := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			println(i)
		}
		forever <- true
	}()

	// forever <- true

	<-forever
}
