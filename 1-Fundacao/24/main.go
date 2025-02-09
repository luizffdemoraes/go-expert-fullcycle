package main

import "fmt"

/***
 For é uma estrutura de repetição que permite executar um bloco de código várias vezes.
***/
func main() {

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	numeros := []string{"zero", "um", "dois", "três", "quatro", "cinco", "seis", "sete", "oito", "nove"}
	// range é uma função que retorna um par de valores, o índice e o valor de um elemento em um array ou slice.
	for k, v := range numeros {
		fmt.Println(k, v)
	}

	// Se você não precisar do índice, pode ignorá-lo usando o caractere _.
	for _, v := range numeros {
		fmt.Println(v)
	}

	i := 0

	/***
	 	O for sem condição irá rodar repetidamente até que você use um break para sair do loop ou retorne de uma função.
		Semelhante ao while em outras linguagens.
	 ***/
	for i < 10 {
		fmt.Println(i)
		i++
	}

	// Loop infinito
	for {
		fmt.Println("loop infinito")
	}
}
