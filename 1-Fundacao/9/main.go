package main

import (
	"errors"
	"fmt"
)

func main() {
	valor, err := sum(50, 2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(valor)
}

// sum é uma função que soma dois números inteiros e retorna o resultado e um erro
func sum(a, b int) (int, error) {
	if a+b > 50 {
		return a + b, errors.New("O valor da soma é maior que 50")
	}
	// Se não houver erro, retorna o resultado da soma e nil (nulo)
	return a + b, nil
}
