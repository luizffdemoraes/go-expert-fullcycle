package main

import (
	"fmt"

	"github.com/fullcycle/curso-go/5-Packaging/1/math"
)

func main() {
	m := math.Math{A: 1, B: 2}
	fmt.Println(m.Add())

	// fmt.Println("Hello, World!")
}
