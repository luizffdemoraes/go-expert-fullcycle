package main

import (
	"fmt"
	"os"
)

func main() {
	i := 0
	for i < 5000 {
		f, err := os.Create(fmt.Sprintf("./tmp/file_%d.txt", i))
		if err != nil {
			panic(err)
		}
		f.WriteString("Hello, World!")
		f.Close()
		i++
	}
}
