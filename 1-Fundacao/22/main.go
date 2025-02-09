package main

import (
	"fmt"

	"github.com/fullcycle/curso-go/matematica"
)

/*
**
O comando abaixo cria um arquivo go.mod que é o arquivo de configuração do módulo go

go mod init github.com/fullcycle/curso-go

# Tudo que começa com letra maiúscula é público e pode ser acessado por outros pacotes

**
*/
func main() {
	soma := matematica.Soma(matematica.A, 20)
	carro := matematica.Carro{Marca: "Fiat"}
	fmt.Println("Resultado: %v", soma)
	fmt.Printf("Carro: %v\n", carro)
	fmt.Println(carro.Andar())
}
