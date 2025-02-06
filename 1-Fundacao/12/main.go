package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

// main é a função principal que será executada ao rodar o programa que exibirá o nome, idade e ativo do cliente
func main() {

	wesley := Cliente{
		Nome:  "Wesley",
		Idade: 30,
		Ativo: true,
	}

	// %s = string, %d = inteiro, %t = boolean
	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t\n", wesley.Nome, wesley.Idade, wesley.Ativo)
}
