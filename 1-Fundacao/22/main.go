package main

import (
	"fmt"

	"github.com/fullcycle/curso-go/matematica"
)

/***
O comando abaixo cria um arquivo go.mod que é o arquivo de configuração do módulo go

go mod init github.com/fullcycle/curso-go

***/
func main() {

	soma := matematica.Soma(10, 20)

	fmt.Println("Resultado: %v", soma)
}
