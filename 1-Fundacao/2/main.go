package main

/***
No gol quando você faz uma declaração de variável sem atribuir um valor a ela,
o compilador atribui um valor zero a ela, que é o valor padrão para o tipo da variável.
Por exemplo,
para inteiros, o valor zero é 0,
para strings é ""
para floats é 0.0
e para booleanos é false.
O código abaixo mostra como o compilador atribui valores zero a variáveis de diferentes tipos.
***/

// Declaração de escopo global e possível declarar e atribuir valores a variáveis

const a = "Hello, World!"

// := Ele só pode ser usado na primeira declaração de uma variável
var (
	b bool    = true
	c int     = 10
	d string  = "Luiz"
	e float64 = 1.2
)

func main() {
	// Declaração de escopo local
	var a string = "X"
	println(a)
}
