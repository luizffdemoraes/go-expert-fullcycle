package main

func main() {
	a := 1
	b := 2
	c := 3

	// if é uma estrutura de controle de fluxo que executa um bloco de código se uma condição for verdadeira.
	if a > b {
		println(a)
	} else {
		println(b)
	}

	// && é um operador lógico que representa a operação "e". Ele retorna verdadeiro se ambos os operandos forem verdadeiros.
	if a > b && c > a {
		println("a > b && c > a")
	}

	// || é um operador lógico que representa a operação "ou". Ele retorna verdadeiro se pelo menos um dos operandos for verdadeiro.
	if a > b || c > a {
		println("a > b && c > a")
	}

	switch a {
	case 1:
		println("a é 1")
	case 2:
		println("a é 2")
	case 3:
		println("a é 3")
	default:
		println("a não é 1, 2 ou 3")
	}

}
