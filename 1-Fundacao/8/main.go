package main

import "fmt"

func main() {
	salarios := map[string]int{"Luiz": 1200, "Joana": 2000, "Maria": 1700}
	fmt.Println(salarios["Luiz"])
	// Percorrendo um map
	for nome, salario := range salarios {
		fmt.Printf("Nome: %s Salário: %d\n", nome, salario)
	}

	// Remover itens de um map
	delete(salarios, "Maria")
	// Verificando se o item foi removido
	for nome, salario := range salarios {
		fmt.Printf("Nome: %s Salário: %d\n", nome, salario)
	}

	// Adicionando itens a um map
	salarios["José"] = 1500
	// Verificando se o item foi adicionado
	for nome, salario := range salarios {
		fmt.Printf("Nome: %s Salário: %d\n", nome, salario)
	}

	// Criar um map vazio
	salOne := make(map[string]int)
	salTwo := map[string]int{}
	salOne["Mario"] = 1200
	salTwo["Joana"] = 2000

	// Percorrendo um map utilizando black identifier '_' para ignorar a chave
	for _, salario := range salOne {
		fmt.Printf("Salário: %d\n", salario)
	}
}
