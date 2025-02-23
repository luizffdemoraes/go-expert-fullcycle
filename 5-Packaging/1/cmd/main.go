package main

import (
	"fmt"

	"github.com/fullcycle/curso-go/5-Packaging/1/math"
)

func main() {
	m := math.NewMath(1, 2)
	// m := math.Math{a: 1, b: 2}
	fmt.Println(m)
	fmt.Println(m.Add())
	fmt.Println(math.X)

	// fmt.Println("Hello, World!")
}
