package main

import (
	"io"
	"net/http"
)

/***
Esse código faz uma **requisição HTTP GET** para `"http://google.com"` usando `http.Client{}`.
Ele cria a requisição com `http.NewRequest()`, define o cabeçalho `Accept: application/json` e a envia com `c.Do(req)`.
Se a requisição for bem-sucedida, o corpo da resposta é lido com `io.ReadAll(resp.Body)` e impresso no terminal.
Caso ocorra um erro em qualquer etapa, o programa **encerra com `panic(err)`**. 🚀
***/

func main() {
	c := http.Client{}
	req, err := http.NewRequest("GET", "http://google.com", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Accept", "application/json")
	resp, err := c.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	println(string(body))
}
