package main

import (
	"fmt"

	"github.com/google/uuid"
)

/*
**
O comando go get github.com/google/uuid instala a dependência do pacote uuid no projeto.
O comando go mod tidy remove qualquer dependência que não esteja sendo utilizada no projeto.
**
*/
func main() {

	fmt.Println(uuid.New())
}
