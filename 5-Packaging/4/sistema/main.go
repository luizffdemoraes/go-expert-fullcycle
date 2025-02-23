package main

import (
	"github.com/fullcycle/curso-go/5-Packaging/4/math"
	"github.com/google/uuid"
)

func main() {
	m := math.NewMath(10, 20)
	println(m.Add())
	println(uuid.New().String())
}
