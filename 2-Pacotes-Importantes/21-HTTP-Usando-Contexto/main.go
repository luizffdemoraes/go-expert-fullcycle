package main

import (
	"context"
	"io"
	"net/http"
	"time"
)

/***
Esse código faz uma requisição HTTP GET para "http://google.com", usando um contexto com timeout extremamente curto (time.Microsecond).
Ele cria um contexto base (context.Background()) e define um timeout com context.WithTimeout(). Se o tempo expirar antes da resposta, a requisição será cancelada automaticamente.
A requisição é criada com http.NewRequestWithContext(), enviada via http.DefaultClient.Do(req), e a resposta, se bem-sucedida, é lida e impressa no terminal.
***/

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Microsecond)
	// ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", "http://google.com", nil)
	if err != nil {
		panic(err)
	}
	resp, err := http.DefaultClient.Do(req)
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
