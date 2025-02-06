package main

import "fmt"

// Closure é uma função anônima que é declarada dentro de outra função
func main() {
	total := func() int {
		return sum(1, 3, 45, 6, 34, 654, 654, 7645, 534, 543, 543, 543)
	}()

	fmt.Println(total)
}

// sum é uma função que soma uma quantidade indefinida de números inteiros e retorna o resultado
func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
