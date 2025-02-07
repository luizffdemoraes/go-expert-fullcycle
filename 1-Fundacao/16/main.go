package main

func main() {

	// Memória -> Endereço -> Valor
	a := 10
	// & -> endereço
	println(&a)
	/***
	 A variável ponteiro contém o endereço de memória de outra variável,
	e esse endereço aponta para o local onde um valor está armazenado.
	***/
	var ponteiro *int = &a
	*ponteiro = 20
	b := &a
	*b = 30
	println(a)
}
