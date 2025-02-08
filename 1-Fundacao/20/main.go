package main

import "fmt"

//Type Assertions são usadas para acessar a interface de um tipo de valor e acessar seus valores subjacentes.
func main() {
	var minhaVar interface{} = "Hello"
	println(minhaVar.(string))
	res, ok := minhaVar.(int)
	fmt.Printf("O valor de res é %v e o resultado de ok é %v\n", res, ok)

	res2 := minhaVar.(int)
	fmt.Printf("O valor de res é %v\n", res2)
}
