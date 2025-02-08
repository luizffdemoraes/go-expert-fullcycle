package main

import "fmt"

// Interfaces vazias são usadas por exemplo para criar funções que aceitam qualquer tipo de parâmetro
func main() {

	var x interface{} = 10
	var y interface{} = "Hello, World!"
	showType(x)
	showType(y)

}

func showType(t interface{}) {
	fmt.Printf("O tipo da variável é %T, e o valor é %v\n", t, t)
}
