package main

import "fmt"

/***

🔹 1. Arrays em Go
Um array em Go é uma coleção de tamanho fixo de elementos do mesmo tipo.
Arrays têm tamanho fixo e não podem ser redimensionados!

***/
func main() {
	var meuArrayOne [3]int
	meuArrayOne[0] = 10
	meuArrayOne[1] = 20
	meuArrayOne[2] = 30

	meuArrayTwo := [3]int{40, 50, 60}

	fmt.Println(meuArrayOne[len(meuArrayOne)-1])

	// Usando for tradicional
	fmt.Println("Usando for tradicional")
	for i := 0; i < len(meuArrayOne); i++ {
		fmt.Printf("O valor do índice é %d e o valor é %d\n", i, meuArrayOne[i])
	}

	// Usando for range
	fmt.Println("Usando for range")
	for i, v := range meuArrayTwo {
		fmt.Printf("O valor do índice é %d e o valor é %d\n", i, v)
	}
}
