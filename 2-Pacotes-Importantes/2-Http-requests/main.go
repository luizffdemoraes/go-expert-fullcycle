package main

import (
	"io"
	"net/http"
)

func main() {
	// Realizar chamada http para o Google
	req, err := http.Get("https://google.com")
	if err != nil {
		panic(err)
	}

	// Ler o corpo da resposta
	res, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	// Imprimir o corpo da resposta
	println(string(res))
	req.Body.Close()
}
