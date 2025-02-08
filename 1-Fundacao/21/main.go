package main

import "fmt"

type MyNumber int

// O ~ é um operador de tipo que representa um tipo que é um dos tipos listados
type Number interface {
	~int | float64
}

/***
 Constraints em Go são feitos através de interfaces para garantir que o tipo passado atenda a um determinado contrato
 Generics em Go são feitos através de interfaces e funções
 https://pkg.go.dev/golang.org/x/exp/constraints
***/
func Soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

// comparable é uma constraint que garante que o tipo passado seja comparável com o operador == do mesmo tipo
func Compara[T comparable](a, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {
	m := map[string]int{"Wesley": 1000, "Calebe": 2000, "Lucas": 3000}
	m2 := map[string]float64{"Wesley": 1000.0, "Calebe": 2000.0, "Lucas": 3000.0}
	m3 := map[string]MyNumber{"Wesley": 1000.0, "Calebe": 2000.0, "Lucas": 3000.0}

	fmt.Println(Soma(m))
	fmt.Println(Soma(m2))
	fmt.Println(Soma(m3))
	fmt.Println(Compara(10, 10.0))
}
