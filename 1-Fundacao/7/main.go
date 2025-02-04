package main

import "fmt"

/***
🔍 Entendendo Slices em Go
Um slice em Go é uma estrutura flexível e eficiente para lidar com sequências de elementos.
Diferente de um array, um slice não tem um tamanho fixo e pode crescer ou diminuir dinamicamente.
***/

func main() {
	s := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	fmt.Printf("len = %d, cap = %d   %v\n", len(s), cap(s), s)
	fmt.Printf("len = %d, cap = %d   %v\n", len(s[:0]), cap(s[:0]), s[:0])
	fmt.Printf("len = %d, cap = %d   %v\n", len(s[:4]), cap(s[:4]), s[:4])
	fmt.Printf("len = %d, cap = %d   %v\n", len(s[2:]), cap(s[2:]), s[2:])

	// Adicionando um elemento ao slice
	s = append(s, 110)
	fmt.Printf("len = %d, cap = %d   %v\n", len(s[2:]), cap(s[2:]), s[2:])

}
