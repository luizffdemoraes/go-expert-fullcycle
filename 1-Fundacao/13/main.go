package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

// main é a função principal que será executada ao rodar o programa utilizara de composição para criar um endereço para o cliente
func main() {
	wesley := Cliente{
		Nome:  "Wesley",
		Idade: 30,
		Ativo: true,
		Endereco: Endereco{
			Logradouro: "Rua dos Bobos",
			Numero:     0,
			Cidade:     "São Paulo",
			Estado:     "SP",
		},
	}
	wesley.Ativo = false
	wesley.Endereco.Cidade = "São José dos Campos"

	// %s = string, %d = inteiro, %t = boolean
	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t, Logradouro: %s, Numero: %d, Cidade: %s, Estado: %s\n", wesley.Nome, wesley.Idade, wesley.Ativo, wesley.Endereco.Logradouro, wesley.Endereco.Numero, wesley.Endereco.Cidade, wesley.Endereco.Estado)
}
