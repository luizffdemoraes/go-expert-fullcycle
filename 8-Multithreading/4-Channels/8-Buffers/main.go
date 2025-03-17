package main


/**
📌 O que esse código faz?
Cria um canal ch com buffer de tamanho 2 → make(chan string, 2).
Envia duas mensagens para o canal → "Hello" e "World".
Lê e imprime as mensagens na ordem em que foram enviadas.

✅ Qual erro ele impede?
O código evita um deadlock que ocorreria se o canal 
fosse sem buffer e não houvesse uma goroutine para consumir 
os valores imediatamente.
**/
func main() {
	ch := make(chan string, 2)

	ch <- "Hello"
	ch <- "World"

	println(<-ch)
	println(<-ch)
}
