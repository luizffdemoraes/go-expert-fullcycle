package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func main() {

	/***
	Esse código faz uma requisição HTTP POST usando http.Client{} para enviar um JSON ({"name": "John Doe"}) para "http://google.com".
	Ele cria um buffer com os dados JSON, envia a requisição com o Content-Type: application/json, e imprime a resposta no terminal usando io.CopyBuffer(os.Stdout, resp.Body, nil).
	No entanto, a URL usada (http://google.com) não aceita POST, então a requisição provavelmente resultará em erro HTTP 405 (Method Not Allowed).
	***/
	c := http.Client{}
	jsonVar := bytes.NewBuffer([]byte(`{"name": "John Doe"}`))
	resp, err := c.Post("http://google.com", "application/json", jsonVar)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	io.CopyBuffer(os.Stdout, resp.Body, nil)
}
